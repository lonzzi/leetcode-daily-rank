package models

type GraphQLRequest struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables"`
}