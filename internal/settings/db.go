package settings

import (
	"context"
	"net/http"

	db "github.com/Mboukhal/FactoryBase/internal/db"
)

type contextKey string

const QueriesKey contextKey = "queries"

// WithQueries creates a middleware that injects queries into the request context
func WithQueries(queries *db.Queries) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), QueriesKey, queries)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetQueries retrieves queries from the request context
func GetQueries(ctx context.Context) *db.Queries {
	queries, ok := ctx.Value(QueriesKey).(*db.Queries)
	if !ok {
		return nil
	}
	return queries
}
