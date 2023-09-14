package utils

import (
	"math/rand"

	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/hack/uniq"
	"github.com/konveyor/tackle2-hub/api"
	. "github.com/onsi/gomega"
)

func CreateMultipleApplications(numberOfApplications int) []api.Application {
	var appSlice []api.Application
	for i := 0; i < numberOfApplications; i++ {
		randomIndex := rand.Intn(len(data.ApplicationSamples))

		// Create random application
		//// Create a local copy
		application := data.ApplicationSamples[randomIndex]
		uniq.ApplicationName(&application)
		Expect(Application.Create(&application)).To(Succeed())
		appSlice = append(appSlice, application)
	}
	return appSlice
}

func DeleteApplicationsBySlice(apps []api.Application) {
	for i := 0; i < len(apps); i++ {
		Expect(Application.Delete(apps[i].ID)).To(Succeed())
	}
}
