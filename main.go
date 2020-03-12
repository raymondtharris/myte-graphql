package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/graph-gophers/graphql-go"
)

const Schema = `
	schema{
		query: Query
	}
	
	type Query{
		post(id:ID, posttype:String): Post!
	}

	type Post {
		id: ID!
		title: String!
		posttype: String!
		markdownpath: String!
		markdowpreview: String
	}
`

type post struct {
	ID    graphql.ID
	Title string
	//PublishedDate time.Date
	PostType        string
	MarkdownPath    string
	MarkdownPreview string
}

type Resolver struct{}

type Request struct {
	ID    float64 `json:"id"`
	Value string  `json:"value"`
}

//func (r *Resolver) Post(args struct{ID graphql.ID, PostType String}) *postResolver{

//}

//type postResolver {
//	post
//}

func mytegraphql(ctx context.Context, request Request) (string, error) {

	fmt.Printf("ID: %f Value %s ", request.ID, request.Value)
	return "Myte GraphQL API", nil
}

func main() {
	lambda.Start(mytegraphql)
}
