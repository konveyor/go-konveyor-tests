package centralconfig

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/konveyor/go-konveyor-tests/utils"
	"github.com/konveyor/tackle2-hub/api"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Analysis Profile Bundle", func() {
	Context("Bundle Download", func() {
		It("Should include uploaded custom rules in downloaded bundle", func() {
			By("Creating a custom rule file in analyzer-lsp format")
			ruleContent := `- category: mandatory
  description: Custom test rule for centralconfig bundle test
  effort: 1
  labels:
  - konveyor.io/target=cloud-readiness
  - test
  - centralconfig
  message: This is a test rule for verifying bundle downloads
  ruleID: centralconfig-test-001
  when:
    java.referenced:
      location: PACKAGE
      pattern: com.example.test
`
			tmpFile, err := os.CreateTemp("", "custom-rule-*.yaml")
			Expect(err).NotTo(HaveOccurred())
			defer os.Remove(tmpFile.Name())

			_, err = tmpFile.WriteString(ruleContent)
			Expect(err).NotTo(HaveOccurred())
			tmpFile.Close()

			By("Uploading the custom rule file")
			uploadedFile, err := utils.RichClient.File.Post(tmpFile.Name())
			Expect(err).NotTo(HaveOccurred())
			Expect(uploadedFile.ID).NotTo(BeZero())
			defer utils.RichClient.File.Delete(uploadedFile.ID)

			By("Getting an existing target")
			targets, err := utils.RichClient.Target.List()
			Expect(err).NotTo(HaveOccurred())

			if len(targets) == 0 {
				Skip("No targets available in the system")
			}

			// Use first available target
			target := &targets[0]

			By("Creating a profile with custom rule file")
			bundleProfile := api.AnalysisProfile{
				Name: "test-bundle-profile",
				Rules: api.ApRules{
					Targets: []api.Ref{
						{ID: target.ID, Name: target.Name},
					},
					Files: []api.Ref{
						{ID: uploadedFile.ID, Name: uploadedFile.Name},
					},
				},
			}

			err = utils.RichClient.AnalysisProfile.Create(&bundleProfile)
			Expect(err).NotTo(HaveOccurred())
			Expect(bundleProfile.ID).NotTo(BeZero())
			defer utils.RichClient.AnalysisProfile.Delete(bundleProfile.ID)

			By("Downloading the profile bundle")
			bundleDir, err := os.MkdirTemp("", "profile-bundle-*")
			Expect(err).NotTo(HaveOccurred())
			defer os.RemoveAll(bundleDir)

			bundlePath := filepath.Join(bundleDir, "bundle.tar.gz")

			// Download bundle using the bundle endpoint
			path := fmt.Sprintf("/analysis/profiles/%d/bundle", bundleProfile.ID)
			err = utils.Client.FileGet(path, bundlePath)
			Expect(err).NotTo(HaveOccurred())

			By("Verifying the bundle was downloaded")
			_, err = os.Stat(bundlePath)
			Expect(err).NotTo(HaveOccurred())

			By("Extracting the bundle")
			extractDir := filepath.Join(bundleDir, "extracted")
			err = os.MkdirAll(extractDir, 0755)
			Expect(err).NotTo(HaveOccurred())

			cmd := exec.Command("tar", "-xzf", bundlePath, "-C", extractDir)
			output, err := cmd.CombinedOutput()
			Expect(err).NotTo(HaveOccurred(), "Failed to extract bundle: %s", string(output))

			By("Verifying bundle structure")
			// Check profile.yaml exists
			profilePath := filepath.Join(extractDir, "profile.yaml")
			_, err = os.Stat(profilePath)
			Expect(err).NotTo(HaveOccurred())

			// Check rules directory exists
			rulesPath := filepath.Join(extractDir, "rules")
			rulesInfo, err := os.Stat(rulesPath)
			if err == nil {
				Expect(rulesInfo.IsDir()).To(BeTrue(), "rules should be a directory")

				// Find all YAML files in rules directory
				var ruleFiles []string
				entries, err := os.ReadDir(rulesPath)
				Expect(err).NotTo(HaveOccurred())
				for _, entry := range entries {
					if !entry.IsDir() && filepath.Ext(entry.Name()) == ".yaml" {
						ruleFiles = append(ruleFiles, filepath.Join(rulesPath, entry.Name()))
					}
				}

				// Check if our custom rule content is in one of the files
				foundCustomRule := false
				for _, ruleFile := range ruleFiles {
					content, err := os.ReadFile(ruleFile)
					if err == nil && string(content) == ruleContent {
						foundCustomRule = true
						break
					}
				}
				Expect(foundCustomRule).To(BeTrue(), "Custom rule content should be found in the bundle")
			}
		})
	})
})
