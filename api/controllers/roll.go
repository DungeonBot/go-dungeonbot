package controllers

type RollController struct{}

func NewRollController() *RollController {
	return &RollController{}
}

func (r *RollController) Routes() []*fireball.Route {
	routes := []*fireball.Route{
		{
			Path: "/roll",
			Handlers: fireball.Handlers{
				"GET":  r.Help,
				"POST": r.RollFromBody,
			},
		},
		{
			Path: "/roll/:roll",
			Handlers: fireball.Handlers{
				"GET": r.RollFromURI,
			},
		},
		{
			Path: "/roll/savedroll/:rollname",
			Handlers: fireball.Handlers{
				"GET":    r.GetRoll,
				"POST":   r.SaveRoll,
				"DELETE": r.DeleteRoll,
			},
		},
	}

	return routes
}

func (r *RollController) Help(c *fireball.Context) (fireball.Response, error) {
	helpText := `Roll API Endpoints:

/roll
	GET:	Print this help text
	POST:	Make the roll(s) contained in the request's body

/roll/:roll
	GET:	Make a single roll found in the request's URI

/roll/savedroll/:rollname
	GET:	Retrieve a saved roll
	POST:	Save a roll
	DELETE:	Delete a saved roll
`

	return fireball.NewResponse(200, []byte(helpText), nil), nil
}

func (r *RollController) RollFromURI(c *fireball.Context) (fireball.Response, error) {
	roll := c.PathVariables["roll"]
	roll, err := r.processRoll(roll)
	if err != nil {
		return nil, err
	}

	result, err := r.calculateRoll(roll)
	if err != nil {
		return nil, err
	}

	return fireball.NewResponse(200, []byte(result), nil), nil
}

func (r *RollController) RollFromBody() (fireball.Response, error) {
	return nil, nil
}

func (r *RollController) GetRoll() (fireball.Response, error) {
	return nil, nil
}

func (r *RollController) SaveRoll() (fireball.Response, error) {
	return nil, nil
}

func (r *RollController) DeleteRoll() (fireball.Response, error) {
	return nil, nil
}

func (r *RollController) calculateRoll(req models.RollCalculatorRequest) (models.RollCalculatorResponse) {

