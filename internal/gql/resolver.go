package gql

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/hvchien216/golang-ecommerce/internal/modules/product/producttransport/gqlproduct"
)

// NewSchema returns the GraphQL schema
func NewSchema(
	productResolver gqlproduct.Resolver,
) graphql.ExecutableSchema {
	cfg := Config{
		Resolvers: &resolver{
			productResolver: productResolver,
		},
	}

	//cfg.Directives.HasPermission = hasPermission

	return NewExecutableSchema(cfg)
}

type resolver struct {
	productResolver gqlproduct.Resolver
}

// Query returns the QueryResolver
//func (r *resolver) Query() QueryResolver {
//	return &queryResolver{r}
//}

// Mutation returns the MutationResolver
func (r *resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type queryResolver struct {
	*resolver
}

type mutationResolver struct {
	*resolver
}
