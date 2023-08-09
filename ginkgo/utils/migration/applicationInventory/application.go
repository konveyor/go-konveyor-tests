package applicationinventory

import (
	"fmt"

	"github.com/konveyor/tackle2-hub/api"
)

func CreateMultipleApplications(numberOfApplications int) []api.Application {
	var appSlice []api.Application
	for i := 0; i < numberOfApplications; i++ {
		// Create new application
		application := api.Application{Name: fmt.Sprintf("test-app%d", i+1)}

		// TODO: dealing with create returns error (e.g. duplicate app) OR random app name + cleanup
		Application.Create(&application)
		appSlice = append(appSlice, application)
	}
	return appSlice
}
