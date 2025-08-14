VENDOR_DIR ?= /tmp/konveyor-vendor
ARCH ?= amd64
JUNIT_REPORT_DIR ?= /tmp/junit-report

# Setup local minikube with tackle - work in progress (TODO: enable auth)
# This is for local setup, CI uses shared github actions
setup:
	mkdir -p ${VENDOR_DIR}
	rm -rf ${VENDOR_DIR}/start-minikube.sh
	rm -rf ${VENDOR_DIR}/install-tackle.sh
	curl https://raw.githubusercontent.com/konveyor/operator/main/hack/start-minikube.sh -Lo ${VENDOR_DIR}/start-minikube.sh  && chmod +x ${VENDOR_DIR}/start-minikube.sh
	curl https://raw.githubusercontent.com/konveyor/operator/main/hack/install-tackle.sh -Lo ${VENDOR_DIR}/install-tackle.sh && chmod +x ${VENDOR_DIR}/install-tackle.sh
	${VENDOR_DIR}/start-minikube.sh && \
	${VENDOR_DIR}/install-tackle.sh

# Clean local minikube with tackle
clean:
	minikube delete || true

# Update Hub dependency with latest binding and api.
update-hub:
	go get -u github.com/konveyor/tackle2-hub@main

#
# Test tiers.
#
# Set HUB_BASE_URL to point to Konveyor installation, e.g. export HUB_BASE_URL="http://192.168.39.16/hub"
#

# TIER0 - a core functionality, should never fail, Konveyor would be fully broken.
test-tier0:
	$(MAKE) test-analysis

# TIER1 - all normal features expected to work.
test-tier1:
	${MAKE} test-hub-api
	$(MAKE) test-metrics
	TIER1=1 $(MAKE) test-analysis

# TIER2 - advanced features and nice-to-haves.
test-tier2:
	TIER2=1 $(MAKE) test-analysis

# TIER3
test-tier3:
	$(MAKE) test-jira
	$(MAKE) test-migrationwave
	TIER3=1 $(MAKE) test-analysis

#
# Feature tests.
#

# Application analysis tests.
test-analysis:
	mkdir -pv ${JUNIT_REPORT_DIR}
	go test -count=1 -p=1 -timeout 7200s -v ./analysis/... 2>&1 | go-junit-report -iocopy -set-exit-code -out ${JUNIT_REPORT_DIR}/analysis-report.xml

# Metrics.
test-metrics:
	cd e2e/metrics/ && ginkgo -v --junit-report=metrics-report.xml --output-dir=${JUNIT_REPORT_DIR}

# Jira Integration.
test-jira:
	cd e2e/jiraintegration/ && ginkgo -v --junit-report=jiraintegration-report.xml --output-dir=${JUNIT_REPORT_DIR}

# Migration wave
test-migrationwave:
	cd e2e/migrationwave/ && ginkgo -v --junit-report=migrationwave-report.xml --output-dir=${JUNIT_REPORT_DIR}

# Hub API remote tests.
test-hub-api:
	./hub-api/run-tests.sh ${HUB_TESTS_REF}

# Add next features tests here and call the target from appropriate stage.

# Execute all tests.
test-all: test-tier0 test-tier1 test-tier2

# Merge Junit reports
merge-report:
	go install github.com/nezorflame/junit-merger@latest
	cd ${JUNIT_REPORT_DIR} && rm -f merged.xml && junit-merger -o merged.xml *

clean-report-dir:
	rm -f ${JUNIT_REPORT_DIR}/*
