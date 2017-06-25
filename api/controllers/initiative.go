package controllers

import (
	"encoding/json"

	"github.com/tlake/go-dungeonbot/api/models"
	"github.com/zpatrick/fireball"
)

type InitiativeController struct{}

func NewInitiativeController() *InitiativeController {
	return &InitiativeController{}
}

func (i *InitiativeController) Routes() []*fireball.Route {
	routes := []*fireball.Route{
		{
			Path: "/init",
			Handlers: fireball.Handlers{
				"GET":  i.Help,
				"POST": i.AddEntity,
			},
		},
		{
			Path: "/init/entity/:entity",
			Handlers: fireball.Handlers{
				"GET":    i.GetEntity,
				"DELETE": i.RemoveEntity,
			},
		},
		{
			Path: "/init/list",
			Handlers: fireball.Handlers{
				"GET": i.ListEntities,
			},
		},
		{
			Path: "/init/clear",
			Handlers: fireball.Handlers{
				"DELETE": i.ClearEntities,
			},
		},
	}

	return routes
}

func (i *InitiativeController) Help(c *fireball.Context) (fireball.Response, error) {
	helpText := `Initiative API Endpoints:

/init
	GET:	Print this help text
	POST:	Add entity (body: {"entity":"ENTITY_NAME","init":"INTEGER"}

/init/:entity
	GET:	View initiative for [:entity]
	DELETE: Remove [:entity] from initiative order

/init/list
	GET:	Print all entities in current initiative

/init/clear
	DELETE: Clears the current initiative
`
	return fireball.NewResponse(200, []byte(helpText), nil), nil
}

func (i *InitiativeController) AddEntity(c *fireball.Context) (fireball.Response, error) {
	var req models.InitiativeAddEntityRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		return nil, err
	}

	// todo: actually write to database
	// if _, err := db.InitiativeAddEntity(req); err != nil {
	//     return nil, err
	//  }

	resp := models.InitiativeAddEntityResponse{
		Entity:     req.Entity,
		Initiative: req.Initiative,
	}

	return fireball.NewJSONResponse(202, resp)
}

func (i *InitiativeController) RemoveEntity(c *fireball.Context) (fireball.Response, error) {
	entity := c.PathVariables["entity"]
	req := models.InitiativeRemoveEntityRequest{
		Entity: entity,
	}

	// todo: actually write to database
	// if _, err := db.InitiativeRemoveEntity(req); err != nil {
	//     return nil, err
	// }

	resp := models.InitiativeRemoveEntityResponse{
		Entity: req.Entity,
	}

	return fireball.NewJSONResponse(202, resp)
}

func (i *InitiativeController) GetEntity(c *fireball.Context) (fireball.Response, error) {
	entity := c.PathVariables["entity"]
	req := models.InitiativeGetEntityRequest{
		Entity: entity,
	}

	// todo: actually get from database
	// if entity, err := db.InitiativeGetEntity(req); err != nil {
	//     return nil, err
	// }

	resp := models.InitiativeGetEntityResponse{
		Entity:     req.Entity,
		Initiative: 15,
	}

	return fireball.NewJSONResponse(200, resp)
}

func (i *InitiativeController) ListEntities(c *fireball.Context) (fireball.Response, error) {
	entities := []models.InitiativeGetEntityResponse{}

	// todo: actually get from database
	// results, err := db.InitiativeListEntities()
	// if err != nil {
	//     return nil, err
	// }

	entities = append(entities, models.InitiativeGetEntityResponse{
		Entity:     "entity1",
		Initiative: 15,
	})
	entities = append(entities, models.InitiativeGetEntityResponse{
		Entity:     "entity2",
		Initiative: 16,
	})
	entities = append(entities, models.InitiativeGetEntityResponse{
		Entity:     "entity3",
		Initiative: 17,
	})

	resp := models.InitiativeListEntitiesResponse{
		Entities: entities,
	}

	return fireball.NewJSONResponse(200, resp)
}

func (i *InitiativeController) ClearEntities(c *fireball.Context) (fireball.Response, error) {
	// todo: actually delete from database
	// if err := db.InitiativeClearEntries(); err != nil {
	//     return nil, err
	// }

	return fireball.NewJSONResponse(202, nil)
}
