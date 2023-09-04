# Golang API test suite for Konveyor

[![API tests on Quay](https://quay.io/repository/konveyor/go-konveyor-tests/status "API tests Repository on Quay")](https://quay.io/repository/konveyor/go-konveyor-tests)
[![License](http://img.shields.io/:license-apache-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0.html) [![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/konveyor/go-konveyor-tests/pulls)

[![API CI Test Suite](https://github.com/konveyor/go-konveyor-tests/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/konveyor/go-konveyor-tests/actions/workflows/main.yml)
[![End-To-End Test Suite](https://github.com/konveyor/go-konveyor-tests/actions/workflows/test-nightly.yml/badge.svg?branch=main)](https://github.com/konveyor/go-konveyor-tests/actions/workflows/test-nightly.yml)

This repository contains application-level tests for Konveyor. That means test focusing on integration of multiple components and real-world Koveyor use-cases. Basic components tests should be placed and executed in their own repositories.

Test are organized in packages/directories by the high-level Konveyor features.

Tests require running Konveyor/MTA installation (e.g. Minikube works great for test development purposes).

## Contribution guidelines

Background for this test suite come from Hub API tests [https://github.com/konveyor/tackle2-hub/pull/268](https://github.com/konveyor/tackle2-hub/tree/main/test).

There is a ```binding``` package providing API client methods https://github.com/konveyor/tackle2-hub/tree/main/binding.

Feel free to follow [application analysis integration test directory](https://github.com/konveyor/go-konveyor-tests/tree/main/analysis) as an example.

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
$ make test-all
```

Run test manually example:

```
$ export HUB_BASE_URL="http://`minikube ip`/hub"
$ go test -count=1 -v ./analysis/
```

## Test tiers

To provide maximum information about the project functionality, tests were separated into three tiers. From core functionality to nice to haves.

### Tier 0

Very basic and core functionality. A bug here would lead to mostly useless project. This stage should never fail. Examples: Hub API or a basic application analysis flow.

```
$ make test-tier0
```

### Tier 1

Features of the project expected to work to satifly most of end-users expectations. Examples: real-world use cases of application analysis, Jira integrations or metrics.

```
$ make test-tier1
```

### Tier 2

More advanced features like complex application analysis or some edge cases. This stage should be kept green, but a failure should not affect most of users.

```
$ make test-tier2
```

## Test execution options

#### PARALLEL

For parallel test execution, set ```export PARALLEL=1```.

#### KEEP

For keep data created by test e.g. for debugging purposes, set ```export KEEP=1```.

## Konveyor CI status

See https://github.com/konveyor/ci

## Code of Conduct
Refer to Konveyor's Code of Conduct [here](https://github.com/konveyor/community/blob/main/CODE_OF_CONDUCT.md).
