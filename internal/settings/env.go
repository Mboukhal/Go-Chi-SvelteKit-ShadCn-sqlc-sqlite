package settings

import (
	"bytes"
	"database/sql"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Mboukhal/FactoryBase/cmd/ui"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pressly/goose/v3"
)

func DevelopmentSettings(r chi.Router) {
	// // setup cors for development
	// r.Use(middleware.SetHeader("Access-Control-Allow-Origin", "http://localhost:"+port))
	// r.Use(middleware.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS"))
	// r.Use(middleware.SetHeader("Access-Control-Allow-Headers", "Content-Type, Authorization"))

	svelte := "http://localhost:1337"

	// everything else → proxy to SvelteKit DEV
	r.NotFound(func(w http.ResponseWriter, req *http.Request) {
		if strings.HasPrefix(req.URL.String(), "/?token=") {
			http.NotFound(w, req)
			return
		}

		path := req.URL.Path

		proxyReq, _ := http.NewRequest(req.Method, svelte+path, req.Body)
		proxyReq.Header = req.Header

		resp, err := http.DefaultClient.Do(proxyReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		for k, v := range resp.Header {
			w.Header()[k] = v
		}
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	})

}

func ProductionSettings(r chi.Router) {
	// Apply gzip middleware to all responses
	r.Use(middleware.Compress(5))

	// Serve static files from /_/{path...}
	r.Get("/_app/*", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// path := req.PathValue("*")

		// // Add cache headers for static assets
		// if path != "" {
		// 	w.Header().Set("Cache-Control", "max-age=1209600, stale-while-revalidate=86400")
		// }

		// Add CSP header if not already set
		// if w.Header().Get("Content-Security-Policy") == "" {
		// 	w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' 'unsafe-inline'; img-src 'self' http://127.0.0.1:* https://tile.openstreetmap.org data: blob:; connect-src 'self' http://127.0.0.1:* https://nominatim.openstreetmap.org; script-src 'self' 'sha256-GRUzBA7PzKYug7pqxv5rJaec5bwDCw1Vo6/IXwvD3Tc='")
		// }

		http.FileServer(http.FS(ui.DistDirFS)).ServeHTTP(w, req)
	}))

	// SPA fallback - serve index.html for all other routes
	r.NotFound(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Try to serve the requested file if it exists

		path := req.URL.Path
		// log.Println("Requested path:", path)
		// fmt.Println("Requested path:", path)
		// Prevent serving files from .well-known directory
		// This is a security measure to avoid exposing sensitive files
		if strings.HasPrefix(path, "/.well-known/") {
			http.NotFound(w, req)
			return
		}

		cleanPath := filepath.Clean(req.URL.Path)
		// log.Println("cleanPath:", cleanPath)
		// If the path is not root, check if the file exists
		if cleanPath != "/" {
			if f, err := ui.DistDirFS.Open(cleanPath[1:]); err == nil {
				defer f.Close()
				stat, _ := f.Stat()
				if !stat.IsDir() {
					http.FileServer(http.FS(ui.DistDirFS)).ServeHTTP(w, req)
					return
				}
			}
		}

		// Fallback to index.html for SPA routing
		indexFile, err := ui.DistDirFS.Open("index.html")
		if err != nil {
			http.Error(w, "index.html not found", http.StatusNotFound)
			return
		}
		defer indexFile.Close()

		info, _ := indexFile.Stat()
		content, _ := io.ReadAll(indexFile)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeContent(w, req, "index.html", info.ModTime(), bytes.NewReader(content))
	}))

}

func OpenDB() (*sql.DB, error) {
	dbFile := os.Getenv("DATABASE_URL")
	if dbFile == "" {
		dbFile = "./migration/database.db"
	}

	// Extract directory from file path
	dbPath := ""
	for i := len(dbFile) - 1; i >= 0; i-- {
		if dbFile[i] == '/' {
			dbPath = dbFile[:i]
			break
		}
	}

	// Ensure the directory for the database file exists
	if dbPath != "" {
		if _, err := os.Stat(dbPath); os.IsNotExist(err) {
			if err := os.MkdirAll(dbPath, 0755); err != nil {
				return nil, err
			}
		}
	}

	// The DSN format for enabling WAL mode
	dsn := "file:" + dbFile + "?_journal_mode=WAL"
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	// Set goose dialect to sqlite3
	goose.SetDialect("sqlite3")

	// Run migrations using goose from GOOSE_MIGDIR environment variable
	migDir := os.Getenv("GOOSE_MIGDIR")
	if migDir == "" {
		migDir = "./internal/adapter/db/schema"
	}

	// log.Printf("Using database file: %s", dbFile)
	// log.Printf("Running migrations from: %s", migDir)

	if err := goose.Up(db, migDir); err != nil {
		// panic("Warning: Migration error: " + err.Error())
		return nil, err
	}

	// The setting is persistent for the database file itself after the first run.
	return db, nil
}
