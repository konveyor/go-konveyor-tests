# Application analysis tests

Run with ```$ make test-analysis```

## Tiers

Due to fragile nature of some analysis parameters for more complex applications, analysis test execution supports multiple tiers.

### Tier 0

Very basic Application->Task->Analyze->Result flow. Default for ```test-analysis``` target.

Test cases: https://github.com/konveyor/go-konveyor-tests/blob/main/analysis/test_cases.go#L6

### Tier 1

Common example application analysis. Should work.

Test cases: https://github.com/konveyor/go-konveyor-tests/blob/main/analysis/test_cases.go#L13

```
TIER1=1 make test-analysis
```

### Tier 2

More advanced analysis parameters. Great if works.

Test cases: https://github.com/konveyor/go-konveyor-tests/blob/main/analysis/test_cases.go#L19

```
TIER2=1 make test-analysis
```

### Tier 3

Application analysis with credentials, which are supplied as part of the test configuration. Should work.

Test cases: https://github.com/konveyor/go-konveyor-tests/blob/main/analysis/test_cases.go#L31

#### Test Requirements
You should provide `GITHUB_USER` and `GITHUB_TOKEN` with write access to [tackle-testapp](https://github.com/konveyor/tackle-testapp) repository:

```
export GITHUB_USER=""
export GITHUB_TOKEN=""
TIER3=1 MAVEN_TESTAPP_USER=$(GITHUB_USER) MAVEN_TESTAPP_TOKEN=$(GITHUB_TOKEN) make test-analysis
```

## Options

[HUB_BASE_URL](https://github.com/konveyor/tackle2-hub/tree/main/test#rest-api) is required in the same way as in API test.

Analysis tests accept ```DEBUG```, ```KEEP``` and ```PARALLEL``` flags.

## Development

Test code itself is in ```analysis_test.go```.

List of test cases and analyzer setup is in ```test_cases.go```.

```tc_...``` files contain all test cases definitions (Application and expected analysis output).
