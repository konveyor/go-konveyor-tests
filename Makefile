VENDOR_DIR ?= /tmp/konveyor-vendor
ARCH ?= amd64
MINIKUBE_IP ?= `minikube ip`

# Setup local minikube with tackle - work in progress (TODO: enable auth)
# This is for local setup, CI uses shared github actions
setup:
	mkdir -p ${VENDOR_DIR}
	rm -rf ${VENDOR_DIR}/start-minikube.sh
	rm -rf ${VENDOR_DIR}/install-tackle.sh
	curl https://raw.githubusercontent.com/konveyor/tackle2-operator/main/hack/start-minikube.sh -Lo ${VENDOR_DIR}/start-minikube.sh  && chmod +x ${VENDOR_DIR}/start-minikube.sh
	curl https://raw.githubusercontent.com/konveyor/tackle2-operator/main/hack/install-tackle.sh -Lo ${VENDOR_DIR}/install-tackle.sh && chmod +x ${VENDOR_DIR}/install-tackle.sh
	${VENDOR_DIR}/start-minikube.sh && \
	${VENDOR_DIR}/install-tackle.sh
	kubectl apply -f hack/windup-addon.yml || true	# Temporary add also windup addon

# Clean local minikube with tackle
clean:
	minikube delete || true

# Execute application analysis tests.
test-analysis:
	HUB_BASE_URL="http://${MINIKUBE_IP}/hub" go test -count=1 -timeout 7200s -v ./analysis/...

# Execute all tests.
test-all: test-analysis

# Update Hub dependency with latest binding and api.
update-hub:
	go get -u github.com/konveyor/tackle2-hub@main
