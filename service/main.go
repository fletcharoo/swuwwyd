package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/hello", func(re *core.RequestEvent) error {
			return re.String(200, "Hello, World!")
		})

		return se.Next()
	})

	log.Println("Starting server on :8080")
	if err := app.Start(); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
