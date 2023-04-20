# Golang API test suite for Konveyor [WIP]

Background for this test suite come from Hub API tests https://github.com/konveyor/tackle2-hub/pull/268.

Application Analysis integration test example: https://github.com/konveyor/tackle2-hub/pull/268/files#diff-fcd2bf711a00447192da2d46171f15d0bb6302397b523e0ba92ac9f61bbb8ff7

## Test status

[![End-To-End API Test](https://github.com/aufi/go-konveyor-tests/actions/workflows/e2e-api-test.yml/badge.svg?branch=main)](https://github.com/aufi/go-konveyor-tests/actions/workflows/e2e-api-test.yml)

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
$ export HUB_ENDPOINT="http://`minikube ip`/hub"   # TODO: Allow set user&password variables
$ go test -v developer/applications-inventory/applications_create_test.go
```

## Code of Conduct
Refer to Konveyor's Code of Conduct [here](https://github.com/konveyor/community/blob/main/CODE_OF_CONDUCT.md).
