package main

import (
	"context"
	"fmt"
	"log"

	"github.com/machinebox/graphql"
)

type ResponseStruct struct {
	AllFilms struct {
		Films []struct {
			Title       string
			Director    string
			ReleaseDate string
		}
	}
}

type SchemaStruct struct {
	Schema struct {
		QueryType struct {
			Name   string
			Fields []struct {
				Name string
			}
		}
	} `json:"__schema"`
}

func main() {
	client := graphql.NewClient("https://swapi-graphql.netlify.app/.netlify/functions/index")

	ctx := context.Background()

	req := graphql.NewRequest(`
		query Query {
			__schema {
				queryType {
					name
					fields {
						name
					}

				}
			}
		}
	`)

	// run it and capture the response

	//var respData ResponseStruct
	//var respData interface{}
	var schemaData SchemaStruct
	if err := client.Run(ctx, req, &schemaData); err != nil {
		log.Fatal(err)
	}

	//fmt.Println(respData)
	fmt.Println(schemaData)

}
