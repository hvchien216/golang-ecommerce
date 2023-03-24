package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hvchien216/golang-ecommerce/config"
	"github.com/hvchien216/golang-ecommerce/internal/gql"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorytransport/gqlcategory"
	"github.com/hvchien216/golang-ecommerce/internal/pkg/component/appctx"
	"github.com/hvchien216/golang-ecommerce/internal/pkg/db/pg"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()

	appConfig := config.NewConfigFromEnv()

	if err := run(ctx, appConfig); err != nil {
		log.Fatalln(err)
	}
}

func run(ctx context.Context, appConfig config.Config) error {
	dbConn, err := pg.NewPool(ctx, appConfig.DBUrl)

	if err != nil {
		return err
	}

	defer dbConn.Close()

	appCtx := appctx.NewAppContext(dbConn)

	r, err := initRouter(ctx, appCtx)

	if err != nil {
		return err
	}
	port := fmt.Sprintf(":%v", appConfig.Port)
	log.Printf("connect to http://localhost%s/ for GraphQL playground", port)

	boil.DebugMode = true

	return http.ListenAndServe(port, r)
}

func initRouter(
	ctx context.Context,
	appCtx appctx.AppContext,
) (http.Handler, error) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	const path = "/api/v1/graphql"
	r.Handle(path, RouteHandler(gql.NewSchema(gqlcategory.New(appCtx)), true))

	return r, nil
}

func RouteHandler(es graphql.ExecutableSchema, isIntrospectionEnabled bool) http.Handler {
	srv := handler.New(es)
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	//srv.SetErrorPresenter(errorPresenter(isIntrospectionEnabled))
	if isIntrospectionEnabled {
		srv.Use(extension.Introspection{})
	}
	return srv
}
