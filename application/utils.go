package application

import (
	"fmt"

	"github.com/konveyor/go-konveyor-tests/hack/uniq"
	"github.com/konveyor/tackle2-hub/api"
)

func CreateMultipleApplications(numberOfApplications int) []api.Application {
	var appSlice []api.Application
	for i := 0; i < numberOfApplications; i++ {
		// Create new application
		application := api.Application{Name: fmt.Sprintf("test-app-%s", uniq.RandString(10))}

		// TODO: dealing with create returns error
		Application.Create(&application)
		appSlice = append(appSlice, application)
	}
	return appSlice
}
