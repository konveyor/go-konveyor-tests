package utils

import (
	"math/rand"

	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/hack/uniq"
	"github.com/konveyor/tackle2-hub/api"
)

func CreateMultipleApplications(numberOfApplications int) []api.Application {
	var appSlice []api.Application
	for i := 0; i < numberOfApplications; i++ {
		randomIndex := rand.Intn(len(data.ApplicationSamples))

		// Create random application
		//// Create a local copy
		application := data.ApplicationSamples[randomIndex]
		uniq.ApplicationName(&application)
		//// TODO: dealing with create returns error
		Application.Create(&application)
		appSlice = append(appSlice, application)
	}
	return appSlice
}

func DeleteApplicationsBySlice(apps []api.Application) {
	for i := 0; i < len(apps); i++ {
		Application.Delete(apps[i].ID)
	}
}
