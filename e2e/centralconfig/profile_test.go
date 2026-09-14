package centralconfig

import (
	"github.com/konveyor/go-konveyor-tests/utils"
	"github.com/konveyor/tackle2-hub/api"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Analysis Profile API", func() {
	var testProfile api.AnalysisProfile

	Context("Basic CRUD Operations", func() {
		It("Should create, get, and delete a profile", func() {
			By("Creating a new analysis profile")
			testProfile = api.AnalysisProfile{
				Name: "test-profile-basic",
			}

			// Create the profile
			err := utils.RichClient.AnalysisProfile.Create(&testProfile)
			Expect(err).NotTo(HaveOccurred())
			Expect(testProfile.ID).NotTo(BeZero())

			By("Getting the profile by ID")
			retrievedProfile, err := utils.RichClient.AnalysisProfile.Get(testProfile.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(retrievedProfile).NotTo(BeNil())
			Expect(retrievedProfile.Name).To(Equal(testProfile.Name))

			By("Verifying profile appears in list")
			profiles, err := utils.RichClient.AnalysisProfile.List()
			Expect(err).NotTo(HaveOccurred())
			Expect(profiles).To(ContainElement(HaveField("ID", testProfile.ID)))

			By("Deleting the profile")
			err = utils.RichClient.AnalysisProfile.Delete(testProfile.ID)
			Expect(err).NotTo(HaveOccurred())

			By("Verifying profile is deleted")
			_, err = utils.RichClient.AnalysisProfile.Get(testProfile.ID)
			Expect(err).To(HaveOccurred(), "Getting deleted profile should return an error")
		})

		It("Should update an existing profile", func() {
			By("Creating a profile to update")
			testProfile = api.AnalysisProfile{
				Name:        "test-profile-update",
				Description: "Original description",
				Rules: api.ApRules{
					Labels: api.InExList{
						Included: []string{"original"},
					},
				},
			}

			err := utils.RichClient.AnalysisProfile.Create(&testProfile)
			Expect(err).NotTo(HaveOccurred())
			originalID := testProfile.ID

			By("Updating the profile")
			testProfile.Description = "Updated description"
			testProfile.Rules.Labels.Included = []string{"updated", "test"}

			err = utils.RichClient.AnalysisProfile.Update(&testProfile)
			Expect(err).NotTo(HaveOccurred())

			By("Verifying the updates")
			updatedProfile, err := utils.RichClient.AnalysisProfile.Get(originalID)
			Expect(err).NotTo(HaveOccurred())
			Expect(updatedProfile.Description).To(Equal("Updated description"))
			Expect(updatedProfile.Rules.Labels.Included).To(ConsistOf("updated", "test"))

			By("Cleaning up")
			err = utils.RichClient.AnalysisProfile.Delete(testProfile.ID)
			Expect(err).NotTo(HaveOccurred())
		})

		It("Should handle label include/exclude rules", func() {
			By("Creating a profile with included and excluded labels")
			labelProfile := api.AnalysisProfile{
				Name: "test-profile-labels",
				Rules: api.ApRules{
					Labels: api.InExList{
						Included: []string{"cloud-readiness", "jakarta-ee"},
						Excluded: []string{"experimental"},
					},
				},
			}

			err := utils.RichClient.AnalysisProfile.Create(&labelProfile)
			Expect(err).NotTo(HaveOccurred())
			defer utils.RichClient.AnalysisProfile.Delete(labelProfile.ID)

			By("Verifying labels are persisted")
			retrieved, err := utils.RichClient.AnalysisProfile.Get(labelProfile.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(retrieved.Rules.Labels.Included).To(ConsistOf("cloud-readiness", "jakarta-ee"))
			Expect(retrieved.Rules.Labels.Excluded).To(ConsistOf("experimental"))
		})

		It("Should handle target references", func() {
			By("Getting available targets")
			targets, err := utils.RichClient.Target.List()
			Expect(err).NotTo(HaveOccurred())

			if len(targets) > 0 {
				By("Creating a profile with target references")
				targetProfile := api.AnalysisProfile{
					Name: "test-profile-targets",
					Rules: api.ApRules{
						Targets: []api.Ref{
							{ID: targets[0].ID, Name: targets[0].Name},
						},
					},
				}

				err = utils.RichClient.AnalysisProfile.Create(&targetProfile)
				Expect(err).NotTo(HaveOccurred())
				Expect(targetProfile.ID).NotTo(BeZero())

				By("Retrieving and verifying target references")
				retrieved, err := utils.RichClient.AnalysisProfile.Get(targetProfile.ID)
				Expect(err).NotTo(HaveOccurred())

				// Verify target was persisted
				Expect(retrieved.Rules.Targets).To(HaveLen(1))

				By("Cleaning up")
				err = utils.RichClient.AnalysisProfile.Delete(targetProfile.ID)
				Expect(err).NotTo(HaveOccurred())
			} else {
				Skip("No targets available in the system to test with")
			}
		})
	})
})
