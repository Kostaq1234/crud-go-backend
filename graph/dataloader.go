package graph

import (
	"context"
	"github.com/Kostaq1234/graphql/graph/models"
	"github.com/go-pg/pg"
	"net/http"
	"time"
)

const userLoaderKey = "userloader"

func DataloaderMiddleware(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userLoader := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Microsecond,
			fetch: func(ids []string) ([]*models.User, []error) {
				var users []*models.User
				err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()
				if err != nil {
					return nil, []error{err}
				}
				return users, nil
			},
		}
		ctx := context.WithValue(r.Context(), userLoaderKey, &userLoader)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
func getUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userLoaderKey).(*UserLoader)

}
