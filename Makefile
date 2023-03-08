TESTS_ROOT ?= ./e2e-api-tests
VENDOR_DIR ?= ${TESTS_ROOT}/vendor
ARCH ?= amd64
MINIKUBE_IP ?= `minikube ip`

# Setup local minikube with tackle - work in progress (TODO: enable auth)
# This is for local setup, CI uses shared github actions
setup:
	curl https://raw.githubusercontent.com/konveyor/tackle2-operator/main/hack/start-minikube.sh -Lo ${VENDOR_DIR}/start-minikube.sh  && chmod +x ${VENDOR_DIR}/start-minikube.sh
	curl https://raw.githubusercontent.com/konveyor/tackle2-operator/main/hack/install-tackle.sh -Lo ${VENDOR_DIR}/install-tackle.sh && chmod +x ${VENDOR_DIR}/install-tackle.sh
	${VENDOR_DIR}/start-minikube.sh && \
	${VENDOR_DIR}/install-tackle.sh

# Clean local minikube with tackle
clean:
	minikube delete || true

# Execute end-to-end testsuite
test:
	export HUB_ENDPOINT="http://${MINIKUBE_IP}/hub"
	go test -v ./...
