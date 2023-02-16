package application

import (
	"github.com/konveyor/tackle2-hub/api"
	"log"
)

func CreateApplication(app *api.Application) error {
	log.Printf("creating app: %v", app.Name)
	return nil
}
