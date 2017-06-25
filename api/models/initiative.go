package models

import swagger "github.com/zpatrick/go-plugin-swagger"

type InitiativeAddEntityRequest struct {
	Entity     string `json:"entity"`
	Initiative int    `json:"initiative"`
}

func (r InitiativeAddEntityRequest) Definition() swagger.Definition {
	return swagger.Definition{
		Type: "object",
		Properties: map[string]swagger.Property{
			"entity":     swagger.NewStringProperty(),
			"initiative": swagger.NewIntProperty(),
		},
	}
}

type InitiativeAddEntityResponse struct {
	Entity     string `json:"entity"`
	Initiative int    `json:"initiative"`
}

func (r InitiativeAddEntityResponse) Definition() swagger.Definition {
	return swagger.Definition{
		Type: "object",
		Properties: map[string]swagger.Property{
			"entity":     swagger.NewStringProperty(),
			"initiative": swagger.NewIntProperty(),
		},
	}
}

type InitiativeGetEntityRequest struct {
	Entity string `json:"entity"`
}

func (r InitiativeGetEntityRequest) Definition() swagger.Definition {
	return swagger.Definition{
		Type: "object",
		Properties: map[string]swagger.Property{
			"entity": swagger.NewStringProperty(),
		},
	}
}

type InitiativeGetEntityResponse struct {
	Entity     string `json:"entity"`
	Initiative int    `json:"initiative"`
}

func (r InitiativeGetEntityResponse) Definition() swagger.Definition {
	return swagger.Definition{
		Type: "object",
		Properties: map[string]swagger.Property{
			"entity":     swagger.NewStringProperty(),
			"initiative": swagger.NewIntProperty(),
		},
	}
}

type InitiativeRemoveEntityRequest struct {
	Entity string `json:"entity"`
}

func (r InitiativeRemoveEntityRequest) Definition() swagger.Definition {
	return swagger.Definition{
		Type: "object",
		Properties: map[string]swagger.Property{
			"entity": swagger.NewStringProperty(),
		},
	}
}

type InitiativeRemoveEntityResponse struct {
	Entity string `json:"entity"`
}

func (r InitiativeRemoveEntityResponse) Definition() swagger.Definition {
	return swagger.Definition{
		Type: "object",
		Properties: map[string]swagger.Property{
			"entity": swagger.NewStringProperty(),
		},
	}
}

type InitiativeListEntitiesResponse struct {
	Entities []InitiativeGetEntityResponse `json:"entities"`
}

func (r InitiativeListEntitiesResponse) Definition() swagger.Definition {
	return swagger.Definition{
		Type: "object",
		Properties: map[string]swagger.Property{
			"entities": swagger.NewObjectSliceProperty("InitiativeGetEntityResponse"),
		},
	}
}
