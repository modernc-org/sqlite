// https://gitlab.com/cznic/sqlite/-/issues/198#note_2232364348
package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"

	_ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	absDb, err := filepath.Abs("simple-web.db")
	if err != nil {
		panic(err)
	}
	dbSlash := "/" + strings.TrimPrefix(filepath.ToSlash(absDb), "/")
	connStr := "file://" + dbSlash + "?_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)&_pragma=busy_timeout(10000)"
	// connStr = "file:///" + dbSlash + "?_foreign_keys=1&_journal_mode=WAL&_synchronous=NORMAL&_busy_timeout=10000&_mutex=no"
	fmt.Printf("connecting to %s\n", connStr)
	db, err := sql.Open("sqlite", connStr)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(1 * time.Minute)
	defer db.Close()

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS data (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL DEFAULT ''
	)`); err != nil {
		panic(err)
	}
	db.Exec(`DELETE FROM data`)
	db.Exec(`INSERT INTO data (Id, Name) VALUES (1, 'A'),(2, 'B'),(3, 'C'),(4, 'D');`)

	http.HandleFunc("/items/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		type entry struct {
			Id   int64  `sql:"id"`
			Name string `sql:"name"`
		}

		scanEntry := func(ctx context.Context, id string) (entry, error) {
			ctx = context.Background()
			// modernc sqlite breaks connections when context gets cancelled
			dbConn, err := db.Conn(ctx)
			// var sqliteErr *sqlite.Error
			// for err != nil && errors.As(err, &sqliteErr) && sqliteErr.Code() == sqlite3.SQLITE_BUSY {
			// 	// fmt.Fprintf(os.Stderr, "failed to obtain connection. retrying %#v\n", err)
			// 	time.Sleep(time.Microsecond)
			// 	dbConn, err = db.Conn(ctx)
			// }
			if err != nil {
				return entry{}, fmt.Errorf("obtaining db conn: %w", err)
			}
			defer dbConn.Close()

			// modernc sqlite breaks connections when context gets cancelled
			// ctx = context.Background()
			row := dbConn.QueryRowContext(ctx, "SELECT Id,Name FROM data WHERE id = ? LIMIT 1", id)
			if err := row.Err(); err != nil {
				return entry{}, fmt.Errorf("retrieving row: %w", err)
			}

			e := entry{}
			if err := row.Scan(&e.Id, &e.Name); err != nil {
				return entry{}, fmt.Errorf("scanning entry: %w", err)
			}
			return e, nil
		}

		const HttpClientClosedRequest = 499

		e, err := scanEntry(r.Context(), id)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				w.WriteHeader(HttpClientClosedRequest)
				return
			}
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(os.Stderr, "%v - failed to retrieve row %v\n", time.Now().Format(time.RFC3339), err)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%#v", e)
	})

	server := http.Server{
		Addr:    "127.0.0.1:8082",
		Handler: http.DefaultServeMux,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Printf("server: %v\n", err)
		}
	}()
	go func() {
		RunClient(ctx)
	}()
	<-ctx.Done()
	server.Shutdown(context.Background())

	db.Close()
}

func RunClient(ctx context.Context) {
	eg, _ := errgroup.WithContext(ctx)

	limit := 20
	eg.SetLimit(limit)
	c := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:      6,
			MaxConnsPerHost:   6,
			DisableKeepAlives: false,
			IdleConnTimeout:   10 * time.Second,
		}, Timeout: 5 * time.Second}

	for i := 0; i < limit; i++ {
		eg.Go(func() error {

			for ctx.Err() == nil {
				res, err := c.Get("http://127.0.0.1:8082/items/2")
				if err != nil {
					fmt.Printf("err http.Do %v\n", err)
					time.Sleep(time.Second)
					continue
				}
				io.Copy(io.Discard, res.Body)
				res.Body.Close()
			}

			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		panic(err)
	}
}
