package models

type RollCalculatorRequest struct {
	NumDice      int  `json:"num_dice"`
	NumSides     int  `json:"num_sides"`
	DieMod       int  `json:"die_mod"`
	TimesToRoll  int  `json:"times_to_roll"`
	RollMod      int  `json:"roll_mod"`
	Advantage    bool `json:"advantage"`
	Disadvantage bool `json:"disadvantage"`
}

type RollCalculatorResponse struct {
	Result       int  `json:"result"`
	NumDice      int  `json:"num_dice"`
	NumSides     int  `json:"num_sides"`
	DieMod       int  `json:"die_mod"`
	TimesToRoll  int  `json:"times_to_roll"`
	RollMod      int  `json:"roll_mod"`
	Advantage    bool `json:"advantage"`
	Disadvantage bool `json:"disadvantage"`
}

func (r RollCalculatorRequest) Definition() swagger.Definition {
	return swagger.Definition{
		Type: "object",
		Properties: map[string]swagger.Property{
			"num_dice":      swagger.NewIntProperty(),
			"num_sides":     swagger.NewIntProperty(),
			"die_mod":       swagger.NewIntProperty(),
			"times_to_roll": swagger.NewIntProperty(),
			"roll_mod":      swagger.NewIntProperty(),
			"advantage":     swagger.NewBoolProperty(),
			"disadvantage":  swagger.NewBoolProperty(),
		},
	}
}

func (r RollCalculatorResponse) Definition() swagger.Definition {
	return swagger.Definition{
		Type: "object",
		Properties: map[string]swagger.Property{
			"result":        swagger.NewIntProperty(),
			"num_dice":      swagger.NewIntProperty(),
			"num_sides":     swagger.NewIntProperty(),
			"die_mod":       swagger.NewIntProperty(),
			"times_to_roll": swagger.NewIntProperty(),
			"roll_mod":      swagger.NewIntProperty(),
			"advantage":     swagger.NewBoolProperty(),
			"disadvantage":  swagger.NewBoolProperty(),
		},
	}
}
