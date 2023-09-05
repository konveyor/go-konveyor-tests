package metrics

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/konveyor/go-konveyor-tests/utils"
)

var (
	operator_namespace string
	hub_pod_name       string
)

func init() {
	hub_pod_name = getHubFullName()
}

func getHubFullName() string {
	operator_namespace = "konveyor-tackle"
	pod := "tackle-hub"

	baseURL := os.Getenv("HUB_BASE_URL")
	if strings.Contains(baseURL, "mta") {
		operator_namespace = "openshift-mta"
		pod = "mta-hub"
	}
	cmd := fmt.Sprintf("oc get pod -n %s | grep %s | awk '{print $1}'", operator_namespace, pod)
	return runCmd(cmd)
}

func GetMetricValue(metricName string) int {
	time.Sleep(30 * time.Second)
	cmd := fmt.Sprintf("oc exec -n %s %s -- curl http://localhost:2112/metrics | grep %s", operator_namespace, hub_pod_name, metricName)
	var outStr []string = strings.Split(runCmd(cmd), "\n")
	metricValue := outStr[len(outStr)-1]
	metricValue = utils.LastString(strings.Split(metricValue, " "))
	outInt, _ := strconv.Atoi(metricValue)
	return outInt
}

func runCmd(cmd string) string {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return fmt.Sprintf("Failed to execute command: %s", cmd)
	}
	out = out[:len(out)-1]
	return string(out)
}
