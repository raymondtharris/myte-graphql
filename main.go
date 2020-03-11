package main

import (
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

func mytegraphql() (string, error) {
	return "Myte GraphQL API", nil
}

func main() {
	lambda.Start(mytegraphql)
}
