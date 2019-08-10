package schema

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/jloom6/phishql/graphql/resolver"
)

type Params struct {
	Resolver resolver.Interface
}

func NewHandleFunc(p Params) (func(w http.ResponseWriter, r *http.Request), error) {
	s, err := New(p)
	if err != nil {
		return nil, err
	}

	return func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:        s,
			RequestString: r.URL.Query().Get("query"),
			Context:       r.Context(),
		})

		if len(result.Errors) > 0 {
			log.Printf("wrong result, unexpected errors: %v", result.Errors)
			return
		}

		if err := json.NewEncoder(w).Encode(result); err != nil {
			log.Printf("failed to encode: %v", err)
		}
	}, nil
}

func New(p Params) (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "RootQuery",
				Fields: graphql.Fields{
					"artists": &graphql.Field{
						Type:    graphql.NewList(ArtistType),
						Resolve: p.Resolver.GetArtists,
					},
					"shows": &graphql.Field{
						Args: graphql.FieldConfigArgument{
							"condition": &graphql.ArgumentConfig{
								Type: ConditionType,
							},
						},
						Type:    graphql.NewList(ShowType),
						Resolve: p.Resolver.GetShows,
					},
					"songs": &graphql.Field{
						Type:    graphql.NewList(SongType),
						Resolve: p.Resolver.GetSongs,
					},
					"tags": &graphql.Field{
						Type:    graphql.NewList(TagType),
						Resolve: p.Resolver.GetTags,
					},
					"tours": &graphql.Field{
						Type:    graphql.NewList(TourType),
						Resolve: p.Resolver.GetTours,
					},
					"venues": &graphql.Field{
						Type:    graphql.NewList(VenueType),
						Resolve: p.Resolver.GetVenues,
					},
				},
			}),
	})
}
