package schema

import (
	"github.com/graphql-go/graphql"
)

var (
	isInitiated = false
	// ArtistType is the GraphQL schema for an artist
	ArtistType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Artist",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
	// SetSongType is the GraphQL schema for a set song
	SetSongType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "SetSong",
			Fields: graphql.Fields{
				"song": &graphql.Field{
					Type: SongType,
				},
				"tag": &graphql.Field{
					Type: TagType,
				},
				"transition": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
	// SetType is the GraphQL schema for a set
	SetType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Set",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"label": &graphql.Field{
					Type: graphql.String,
				},
				"songs": &graphql.Field{
					Type: graphql.NewList(SetSongType),
				},
			},
		},
	)
	// ShowType is the GraphQL schema for a show
	ShowType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Show",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"date": &graphql.Field{
					Type: graphql.DateTime,
				},
				"artist": &graphql.Field{
					Type: ArtistType,
				},
				"venue": &graphql.Field{
					Type: VenueType,
				},
				"tour": &graphql.Field{
					Type: TourType,
				},
				"notes": &graphql.Field{
					Type: graphql.String,
				},
				"soundcheck": &graphql.Field{
					Type: graphql.String,
				},
				"sets": &graphql.Field{
					Type: graphql.NewList(SetType),
				},
			},
		},
	)
	// SongType is the GraphQL schema for a song
	SongType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Song",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
	// TagType is the GraphQL schema for a tag
	TagType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Tag",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"text": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
	// TourType is the GraphQL schema for a tour
	TourType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Tour",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
	// VenueType is the GraphQL schema for a venue
	VenueType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Venue",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"city": &graphql.Field{
					Type: graphql.String,
				},
				"state": &graphql.Field{
					Type: graphql.String,
				},
				"country": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
	// ConditionType is the GraphQL schema for a condition
	ConditionType = graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "Condition",
			Fields: graphql.InputObjectConfigFieldMap{
				"base": &graphql.InputObjectFieldConfig{
					Type: BaseConditionType,
				},
			},
		},
	)
	// BaseConditionType is the GraphQL schema for a base condition
	BaseConditionType = graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "BaseCondition",
			Fields: graphql.InputObjectConfigFieldMap{
				"year": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"month": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"day": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"dayOfWeek": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"city": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"state": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"country": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"song": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		},
	)
)

// InitTypes adds the circular dependencies of ConditionType since they can't be defined in the constructor
func InitTypes() {
	if !isInitiated {
		ConditionType.AddFieldConfig("and", &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(ConditionType),
		})
		ConditionType.AddFieldConfig("or", &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(ConditionType),
		})

		isInitiated = true
	}
}
