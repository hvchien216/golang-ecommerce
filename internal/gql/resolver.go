package gql

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorytransport/gqlcategory"
)

// NewSchema returns the GraphQL schema
// productResolver gqlproduct.Resolver,
func NewSchema(
	categoryResolver gqlcategory.Resolver,
) graphql.ExecutableSchema {
	cfg := Config{
		Resolvers: &resolver{
			categoryResolver: categoryResolver,
		},
	}

	//cfg.Directives.HasPermission = hasPermission

	return NewExecutableSchema(cfg)
}

// productResolver  gqlproduct.Resolver
type resolver struct {
	categoryResolver gqlcategory.Resolver
}

// Query returns the QueryResolver
func (r *resolver) Query() QueryResolver {
	return &queryResolver{r}
}

// Mutation returns the MutationResolver
func (r *resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *resolver) Category() CategoryResolver {
	return &categoryResolver{r}
}

type queryResolver struct {
	*resolver
}

type mutationResolver struct {
	*resolver
}

type categoryResolver struct {
	*resolver
}
