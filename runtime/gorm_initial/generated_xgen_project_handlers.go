// Code generated by github.com/goxgen/goxgen, DO NOT EDIT.

package gorm_initial

import(
    "github.com/99designs/gqlgen/graphql"
    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/urfave/cli/v2"
    "context"
    "embed"
    "github.com/goxgen/goxgen/runtime/gorm_initial/generated"
)

//go:embed tests/*
var TestsFS embed.FS

func Server(ctx *cli.Context) (*handler.Server, error){
    resolver, err := NewResolver(ctx)
    if err != nil {
        return nil, err
    }

    return handler.NewDefaultServer(
        generated.NewExecutableSchema(
            generated.Config{
                Resolvers:  resolver,
                Directives: generated.DirectiveRoot{
                    ListAction: func(ctx context.Context, obj interface{}, next graphql.Resolver, resource string, action generated.XgenResourceListActionType, route *string, pagination *bool, schemaFieldName *string) (res interface{}, err error) {
                        return next(ctx)
                    },
                    Action: func(ctx context.Context, obj interface{}, next graphql.Resolver, resource string, action generated.XgenResourceActionType, route *string, schemaFieldName *string) (res interface{}, err error) {
                        return next(ctx)
                    },
                    Resource: func(ctx context.Context, obj interface{}, next graphql.Resolver, name string, route *string, primary *bool, db *generated.XgenResourceDbConfigInput) (res interface{}, err error) {
                        return next(ctx)
                    },
                    ActionField: func(ctx context.Context, obj interface{}, next graphql.Resolver, label *string, description *string, resourceMap []string) (res interface{}, err error) {
                        return next(ctx)
                    },
                    Field: func(ctx context.Context, obj interface{}, next graphql.Resolver, label *string, description *string, db *generated.XgenResourceFieldDbConfigInput) (res interface{}, err error) {
                        return next(ctx)
                    },
                },
            },
        ),
    ), nil
}

