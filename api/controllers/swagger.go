package controllers

import (
	"encoding/json"

	"github.com/tlake/go-dungeonbot/api/models"
	"github.com/zpatrick/fireball"
	swagger "github.com/zpatrick/go-plugin-swagger"
)

type SwaggerController struct {
	host string
}

func NewSwaggerController(host string) *SwaggerController {
	return &SwaggerController{
		host: host,
	}
}

func (s *SwaggerController) Routes() []*fireball.Route {
	routes := []*fireball.Route{
		{
			Path: "/swagger.json",
			Handlers: fireball.Handlers{
				"GET": s.ServeSwaggerJSON,
			},
		},
	}

	return routes
}

func (s *SwaggerController) ServeSwaggerJSON(c *fireball.Context) (fireball.Response, error) {
	spec := swagger.Spec{
		SwaggerVersion: "2.0",
		Host:           s.host,
		Schemes:        []string{"http"},
		Info: &swagger.Info{
			Title:   "DungeonBot",
			Version: "5.something, I think?",
		},
		Tags: []swagger.Tag{
			{
				Name:        "Initiative",
				Description: "Methods related to the initiative tracker",
			},
		},
		Paths: map[string]swagger.Path{
			"/init": map[string]swagger.Method{
				"get": {
					Summary: "Print help text",
					Tags:    []string{"Initiative"},
					Responses: map[string]swagger.Response{
						"200": {
							Description: "The help text",
						},
					},
				},
				"post": {
					Summary: "Add an entity to the initiative order",
					Tags:    []string{"Initiative"},
					Parameters: []swagger.Parameter{
						swagger.NewBodyParam("InitiativeAddEntityRequest", "Entity to add", true),
					},
					Responses: map[string]swagger.Response{
						"202": {
							Description: "The added entity",
							Schema:      swagger.NewObjectSchema("InitiativeAddEntityResponse"),
						},
					},
				},
			},
			"/init/entity/{entity}": map[string]swagger.Method{
				"get": {
					Summary: "Describe an initiative entity",
					Tags:    []string{"Initiative"},
					Parameters: []swagger.Parameter{
						swagger.NewStringPathParam("InitiativeGetEntityRequest", "Entity to describe", true),
					},
					Responses: map[string]swagger.Response{
						"200": {
							Description: "The desired entity",
							Schema:      swagger.NewObjectSchema("InitiativeGetEntityResponse"),
						},
					},
				},
				"delete": {
					Summary: "Remove an entity from the initiative order",
					Tags:    []string{"Initiative"},
					Parameters: []swagger.Parameter{
						swagger.NewStringPathParam("InitiativeRemoveEntityRequest", "Entity to remove", true),
					},
					Responses: map[string]swagger.Response{
						"202": {
							Description: "The deleted entity",
							Schema:      swagger.NewObjectSchema("InitiativeRemoveEntityResponse"),
						},
					},
				},
			},
			"/init/list": map[string]swagger.Method{
				"get": {
					Summary: "Get the entire current initiative order",
					Tags:    []string{"Initiative"},
					Responses: map[string]swagger.Response{
						"200": {
							Description: "The current initiative order",
							Schema:      swagger.NewObjectSchema("InitiativeListEntitiesResponse"),
						},
					},
				},
			},
			"/init/clear": map[string]swagger.Method{
				"delete": {
					Summary: "Clear the current initiative order",
					Tags:    []string{"Initiative"},
					Responses: map[string]swagger.Response{
						"202": {
							Description: "Success",
						},
					},
				},
			},
		},
		Definitions: map[string]swagger.Definition{
			"InitiativeAddEntityRequest":     models.InitiativeAddEntityRequest{}.Definition(),
			"InitiativeAddEntityResponse":    models.InitiativeAddEntityResponse{}.Definition(),
			"InitiativeGetEntityRequest":     models.InitiativeGetEntityRequest{}.Definition(),
			"InitiativeGetEntityResponse":    models.InitiativeGetEntityResponse{}.Definition(),
			"InitiativeRemoveEntityRequest":  models.InitiativeRemoveEntityRequest{}.Definition(),
			"InitiativeRemoveEntityResponse": models.InitiativeRemoveEntityResponse{}.Definition(),
			"InitiativeListEntitiesResponse": models.InitiativeListEntitiesResponse{}.Definition(),
		},
	}

	bytes, err := json.MarshalIndent(spec, "", "	")
	if err != nil {
		return nil, err
	}

	return fireball.NewResponse(200, bytes, fireball.JSONHeaders), nil
}
