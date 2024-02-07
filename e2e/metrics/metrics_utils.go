package metrics

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	. "github.com/onsi/gomega"

	"github.com/konveyor/go-konveyor-tests/utils"
)

var (
	operator_namespace string
	hub_pod_name       string
)

func getHubFullName() string {
	operator_namespace = "konveyor-tackle"
	pod := "tackle-hub"

	baseURL := os.Getenv("HUB_BASE_URL")
	if strings.Contains(baseURL, "mta") {
		operator_namespace = "openshift-mta"
		pod = "mta-hub"
	}
	cmd := fmt.Sprintf("oc get pod -n %s | grep %s | awk '{print $1}'", operator_namespace, pod)
	return runOpenshiftCmd(cmd)
}

func GetMetricValue(metricName string) int {
	hub_pod_name = getHubFullName()
	utils.Log.Info("Waiting for 30 seconds...")
	time.Sleep(30 * time.Second)
	cmd := fmt.Sprintf("oc exec -n %s %s -- curl -s http://localhost:2112/metrics | grep %s", operator_namespace, hub_pod_name, metricName)
	var outStr []string = strings.Split(runOpenshiftCmd(cmd), "\n")
	metricValue := outStr[len(outStr)-1]
	metricValue = utils.LastString(strings.Split(metricValue, " "))
	outInt, _ := strconv.Atoi(metricValue)
	return outInt
}

// runOpenshiftCmd runs an OpenShift command and returns its output as a string.
func runOpenshiftCmd(cmd string) string {
	os.Setenv("KUBECONFIG", os.Getenv("KUBECONFIG"))

	// Create a new command to execute in a Bash environment.
	command := exec.Command("bash", "-c", cmd)

	// Create a buffer to capture the standard output of the executed command.
	var out bytes.Buffer
	command.Stdout = &out

	// Execute the command and check for any errors.
	err := command.Run()
	Expect(err).ShouldNot(HaveOccurred())

	// Assert that the captured output is not empty.
	Expect(out.Len()).ShouldNot(BeZero())

	// Removing the trailing newline character ("\n")
	outStr := out.String()
	return outStr[:len(outStr)-1]
}
