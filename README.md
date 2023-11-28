# Golang API test suite for Konveyor

[![License](https://img.shields.io/:license-apache-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0.html) [![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/konveyor/go-konveyor-tests/pulls)

[![Test TIER0](https://github.com/konveyor/go-konveyor-tests/actions/workflows/main-tier0.yml/badge.svg)](https://github.com/konveyor/go-konveyor-tests/actions/workflows/main-tier0.yml)
[![Test TIER1](https://github.com/konveyor/go-konveyor-tests/actions/workflows/main-tier1.yml/badge.svg)](https://github.com/konveyor/go-konveyor-tests/actions/workflows/main-tier1.yml)
[![Test TIER2](https://github.com/konveyor/go-konveyor-tests/actions/workflows/main-tier2.yml/badge.svg)](https://github.com/konveyor/go-konveyor-tests/actions/workflows/main-tier2.yml)

[![Test nightly TIER0](https://github.com/konveyor/go-konveyor-tests/actions/workflows/nightly-tier0.yml/badge.svg)](https://github.com/konveyor/go-konveyor-tests/actions/workflows/nightly-tier0.yml)
[![Test nightly TIER1](https://github.com/konveyor/go-konveyor-tests/actions/workflows/nightly-tier1.yml/badge.svg)](https://github.com/konveyor/go-konveyor-tests/actions/workflows/nightly-tier1.yml)
[![Test nightly TIER2](https://github.com/konveyor/go-konveyor-tests/actions/workflows/nightly-tier2.yml/badge.svg)](https://github.com/konveyor/go-konveyor-tests/actions/workflows/nightly-tier2.yml)

This repository contains application-level tests for Konveyor. That means test focusing on integration of multiple components and real-world Koveyor use-cases. Basic components tests should be placed and executed in their own repositories.

Test are organized in packages/directories by the high-level Konveyor features.

Tests require running Konveyor/MTA installation (e.g. Minikube works great for test development purposes).

## Contribution guidelines

Background for this test suite come from Hub API tests [https://github.com/konveyor/tackle2-hub/pull/268](https://github.com/konveyor/tackle2-hub/tree/main/test).

There is a `binding` package providing API client methods https://github.com/konveyor/tackle2-hub/tree/main/binding.

Feel free to follow [application analysis integration test directory](https://github.com/konveyor/go-konveyor-tests/tree/main/analysis) as an example.

## Konveyor CI status

[![Run Konveyor main nightly tests](https://github.com/konveyor/ci/actions/workflows/nightly-main.yaml/badge.svg?branch=main)](https://github.com/konveyor/ci/actions/workflows/nightly-main.yaml)

More details at https://github.com/konveyor/ci

## Local test suite execution

### Clone this repo

```
git clone https://github.com/aufi/go-konveyor-tests && cd go-konveyor-tests
```

### Prepare environment

```
$ make setup # start minikube&tackle using David's scripts - local env only
```

### OpenShift cluster

These tests can be executed against OpenShift cluster with Konveyor installed by setting `KUBECONFIG` variable:

```
KUBECONFIG=<kubeconfig file>
```

> **_NOTE:_** You might be required to download and import the certificate chain. Please see [Hub API test README](https://github.com/konveyor/tackle2-hub/tree/main/test#https) for more information.

### Run test suite

Set `$HUB_BASE_URL` environment variable to point to Konveyor installation before running tests. More options could be found in [Hub API test README](https://github.com/konveyor/tackle2-hub/tree/main/test#rest-api).

```
$ HUB_BASE_URL="http://<KONVEYOR_HOST>/hub" make test-tier0
```

Run test manually example:

```
$ export HUB_BASE_URL="http://`minikube ip`/hub"
$ go test -count=1 -v ./analysis/
```

## Test tiers

To provide maximum information about the project functionality, tests were separated into three tiers. From core functionality to nice to haves.

### Tier 0

Very basic and core functionality. A bug here would lead to mostly useless project. This tier should never fail. Examples: Hub API or a basic application analysis flow.

```
$ make test-tier0
```

### Tier 1

Features of the project expected to work to satifly most of end-users expectations. Examples: real-world use cases of application analysis, Jira integrations or metrics.

```
$ make test-tier1
```

### Tier 2

More advanced features like complex application analysis or some edge cases. This tier should be kept green, but a failure should not affect most of users.

```
$ make test-tier2
```

## Test execution options

### DEBUG

For debug output like printing full analysis results, set `export DEBUG=1`.

### KEEP

For keep data created by test e.g. for debugging purposes, set `export KEEP=1`.

### PARALLEL

For parallel test execution, set `export PARALLEL=1`.

## Configuration

**Note:** Some tests utilize environment variables for configuration. Refer to the `Configuration` struct in `config/configuration.go` file for details on supported variables.

Ensure that the required variables are set as environment variables or added to `config/config.json` file.

### Test-Specific Configuration

Refer to the `README.md` files in each folder for test-specific configuration details.

- [e2e/jiraintegration](https://github.com/konveyor/go-konveyor-tests/blob/main/e2e/jiraintegration.README.md)

## Code of Conduct

Refer to Konveyor's Code of Conduct [here](https://github.com/konveyor/community/blob/main/CODE_OF_CONDUCT.md).
