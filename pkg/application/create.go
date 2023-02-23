package application

import (
	"fmt"
	"github.com/konveyor/tackle2-hub/addon"
	"github.com/konveyor/tackle2-hub/api"
)

func CreateApplication(client *addon.Client, app *api.Application) error {
	return client.Post(api.ApplicationsRoot, app)
}

func DeleteApplication(client *addon.Client, id uint) error {
	return client.Delete(fmt.Sprintf("%s/%d", api.ApplicationsRoot, id))
}
