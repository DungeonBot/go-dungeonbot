package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tlake/go-dungeonbot/api/config"
	"github.com/tlake/go-dungeonbot/api/controllers"
	"github.com/tlake/go-dungeonbot/api/logging"
	"github.com/urfave/cli"
	"github.com/zpatrick/fireball"
)

const (
	SWAGGER_URL     = "/api/"
	SWAGGER_UI_PATH = "static/swagger-ui/dist"
)

func serveSwaggerUI(w http.ResponseWriter, r *http.Request) {
	dir := http.Dir(SWAGGER_UI_PATH)
	fileServer := http.FileServer(dir)
	http.StripPrefix(SWAGGER_URL, fileServer).ServeHTTP(w, r)
}

func main() {
	app := cli.NewApp()
	app.Name = "dungeonbot"
	app.Flags = []cli.Flag{
		cli.BoolTFlag{
			Name:   "d, debug",
			EnvVar: config.ENVVAR_DEBUG,
		},
	}

	app.Before = func(c *cli.Context) error {
		if err := validateConfig(c); err != nil {
			return err
		}

		log.SetOutput(logging.NewLogWriter(c.Bool("debug")))
		return nil
	}

	app.Action = func(c *cli.Context) error {

		rootController := controllers.NewRootController()
		swaggerController := controllers.NewSwaggerController(c.String("swagger-host"))
		// attributeController := controllers.NewAttributeController()
		// highlightController := controllers.NewHighlightController()
		initiativeController := controllers.NewInitiativeController()
		// karmaController := controllers.NewKarmaController()
		// questController := controllers.NewQuestController()
		rollController := controllers.NewRollController()

		routes := rootController.Routes()
		routes = append(routes, swaggerController.Routes()...)
		// routes = append(routes, attributeController.Routes()...)
		routes = append(routes, initiativeController.Routes()...)
		// routes = append(routes, karmaController.Routes()...)
		// routes = append(routes, questController.Routes()...)
		routes = append(routes, rollController.Routes()...)
		routes = fireball.Decorate(routes, fireball.LogDecorator())

		fb := fireball.NewApp(routes)

		log.Println("Running on port 9090")

		http.Handle("/", fb)
		http.HandleFunc(SWAGGER_URL, serveSwaggerUI)

		return http.ListenAndServe(":9090", nil)
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func validateConfig(c *cli.Context) error {
	vars := map[string]error{}

	for name, err := range vars {
		if c.String(name) == "" {
			return err
		}
	}

	return nil
}
