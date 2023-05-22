# Golang API test suite for Konveyor

[![End-To-End Test Suite](https://github.com/konveyor/go-konveyor-tests/actions/workflows/test-nightly.yml/badge.svg?branch=main)](https://github.com/konveyor/go-konveyor-tests/actions/workflows/test-nightly.yml)

This repository contains application-level tests for Konveyor. That means test focusing on integration of multiple components and real-world Koveyor use-cases. Basic components tests should be placed and executed in their own repositories.

Test are organized in packages/directories by the high-level Konveyor features.

Tests require running Konveyor/MTA installation (e.g. Minikube works great for test development purposes).

## Contribution guidelines

Background for this test suite come from Hub API tests https://github.com/konveyor/tackle2-hub/pull/268.

Application Analysis integration test example: https://github.com/konveyor/tackle2-hub/blob/main/test/integration/applications-inventory/analysis/windup_basic_test.go

Full hub provided test "framework": [https://github.com/konveyor/tackle2-hub/pull/268/files#diff-fcd2bf711a00447192da2d46171f15d0bb6302397b523e0ba92ac9f61bbb8ff7](https://github.com/konveyor/tackle2-hub/tree/main/test)


## Local test suite execution

### Clone this repo

```
git clone https://github.com/aufi/go-konveyor-tests && cd go-konveyor-tests
```

### Prepare environment

```
$ make setup # start minikube&tackle using David's scripts - local env only
```

### Run test suite

```
$ make test
```

Run test manually example:

```
$ export HUB_BASE_URL="http://`minikube ip`/hub"
$ go test -v analysis/windup_basic_test.go
```
## Konveyor CI status


Component | CI (after merge) | Nightly (cron)
--|--|--
**Hub** | [![Hub main](https://github.com/konveyor/tackle2-hub/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/konveyor/tackle2-hub/actions/workflows/main.yml) | [![Hub Test Suite](https://github.com/konveyor/tackle2-hub/actions/workflows/test-nightly.yml/badge.svg?branch=main)](https://github.com/konveyor/tackle2-hub/actions/workflows/test-nightly.yml)
**UI** | [![UI test](https://github.com/konveyor/tackle2-ui/actions/workflows/ci-actions.yml/badge.svg?branch=main)](https://github.com/konveyor/tackle2-ui/actions/workflows/ci-actions.yml) | -
**Addon/Windup** | [![Test Windup Addon](https://github.com/konveyor/tackle2-addon-windup/actions/workflows/test-windup.yml/badge.svg?branch=main)](https://github.com/konveyor/tackle2-addon-windup/actions/workflows/test-windup.yml) | - |
**E2E API** | [![End-To-End Test Suite](https://github.com/konveyor/go-konveyor-tests/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/konveyor/go-konveyor-tests/actions/workflows/main.yml) | [![End-To-End Test Suite](https://github.com/konveyor/go-konveyor-tests/actions/workflows/test-nightly.yml/badge.svg?branch=main)](https://github.com/konveyor/go-konveyor-tests/actions/workflows/test-nightly.yml)
**E2E UI** | [![UI E2E test](https://github.com/konveyor/tackle-ui-tests/actions/workflows/k8s-ci.yml/badge.svg?branch=main)](https://github.com/konveyor/tackle-ui-tests/actions/workflows/k8s-ci.yml) | [![UI E2E test](https://github.com/konveyor/tackle-ui-tests/actions/workflows/k8s-cron.yml/badge.svg?branch=main)](https://github.com/konveyor/tackle-ui-tests/actions/workflows/k8s-cron.yml)


## Code of Conduct
Refer to Konveyor's Code of Conduct [here](https://github.com/konveyor/community/blob/main/CODE_OF_CONDUCT.md).
