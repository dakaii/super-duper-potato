package controller

import (
	"context"
	"fmt"
	"graphyy/controller/auth"
	"graphyy/model"
	"net/http"
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Schema builds a graphql schema and returns it
func Schema(controllers *Controllers) graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    getRootQuery(controllers),
		Mutation: getRootMutation(controllers),
	})
	if err != nil {
		panic(err)
	}

	return schema
}

// GraphqlHandlfunc is a handler for the graphql endpoint.
func GraphqlHandlfunc(schema graphql.Schema) *handler.Handler {
	return handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
		RootObjectFn: func(ctx context.Context, req *http.Request) map[string]interface{} {
			token := req.Header.Get("Authorization")
			user, _ := verifyToken(token)
			return map[string]interface{}{
				"currentUser": user,
			}
		},
	})
}

func verifyToken(token string) (model.User, error) {
	var user model.User
	var err error
	splitToken := strings.Split(token, "Bearer ")
	if len(splitToken) >= 2 {
		token := splitToken[1]
		user, err = auth.VerifyJWT(token)
		fmt.Printf("%+v\n", err)
	}
	return user, err
}
