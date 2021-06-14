//go:generate go run -mod=mod github.com/99designs/gqlgen --verbose

package graphql

import (
	"net/http"
	"time"

	"entexample/pkg/ent"
	"entexample/pkg/graphql/gqlgen"
	"entexample/pkg/graphql/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/websocket"
)

func NewHandler(client *ent.Client) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.DefaultLogger)

	gqlConfig := gqlgen.Config{
		Resolvers: resolver.New(client),
	}

	schema := gqlgen.NewExecutableSchema(gqlConfig)
	srv := handler.New(schema)

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.AutomaticPersistedQuery{Cache: lru.New(100)})
	playgroundHandler := playground.Handler("GraphQL", "/graphql")

	r.Handle("/graphql", srv)
	r.Get("/graphql/explorer", playgroundHandler)

	return r
}
