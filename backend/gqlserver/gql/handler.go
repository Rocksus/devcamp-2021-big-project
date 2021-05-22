package gql

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"log"
	"net/http"
)

type Handler struct {
	schemaWrapper *SchemaWrapper
}

func NewHandler(sw *SchemaWrapper) *Handler {
	return &Handler{
		schemaWrapper: sw,
	}
}

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func (h *Handler) Handle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var p postData
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			w.WriteHeader(400)
			return
		}

		result := graphql.Do(graphql.Params{
			Schema:         h.schemaWrapper.Schema,
			RequestString:  p.Query,
			VariableValues: p.Variables,
			OperationName:  p.Operation,
		})
		if len(result.Errors) > 0 {
			log.Println("[GQLHandler][Handle] there were some errors, errs: ", result.Errors)
		}

		json.NewEncoder(w).Encode(result)
	})
}
