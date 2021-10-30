package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"

	"github.com/uptrace/opentelemetry-go-extra/otelgraphql"
	"github.com/uptrace/opentelemetry-go-extra/otelplay"
)

const schemaString = `
	schema {
		query: Query
		mutation: Mutation
	}
	type User {
		userID: ID!
		fullName: String!
		username: String!
		organization: String!
	}
	input UserInput {
		fullName: String!
		username: String!
		organization: String!
	}
	type Query {
		user(username: String!): User
		users: [User!]!
		usersOfOrganization(organization: String!): [User!]!
	}
	type Mutation {
		createUser(userInput: UserInput!): User
	}
`

type RootResolver struct{}

type User struct {
	UserID       graphql.ID
	FullName     string
	Username     string
	Organization string
}

type UserInput struct {
	FullName     string
	Username     string
	Organization string
}

var users = []User{
	{graphql.ID("1"), "John Smith", "johnsmith", "HR"},
	{graphql.ID("2"), "Jone Doe", "jonedoe", "IT"},
	{graphql.ID("3"), "Jane Doe", "janedoe", "Marketing"},
}

func (*RootResolver) User(args struct{ Username string }) (*User, error) {
	for _, u := range users {
		if u.Username == args.Username {
			return &u, nil
		}
	}
	return nil, nil
}

func (*RootResolver) Users() ([]User, error) {
	return users, nil
}

func (*RootResolver) UsersOfOrganization(args struct{ Organization string }) ([]User, error) {
	return []User{}, errors.New("intentional error")
}

func (*RootResolver) CreateUser(args struct{ UserInput UserInput }) (*User, error) {
	user := User{
		UserID:       graphql.ID(uuid.NewString()),
		FullName:     args.UserInput.FullName,
		Username:     args.UserInput.Username,
		Organization: args.UserInput.Organization,
	}
	users = append(users, user)
	return &user, nil
}

var schema *graphql.Schema

func main() {
	ctx := context.Background()

	shutdown := otelplay.ConfigureOpentelemetry(ctx)
	defer shutdown()

	tracer := otelgraphql.NewTracer()
	opts := []graphql.SchemaOpt{
		graphql.Tracer(tracer),
		graphql.UseFieldResolvers(),
	}
	schema = graphql.MustParseSchema(schemaString, &RootResolver{}, opts...)

	http.Handle("/graphql", &relay.Handler{Schema: schema})

	fmt.Println("listening on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
