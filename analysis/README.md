# Application analysis tests

Run with `$ make test-analysis`

## Test Case Configuration

### YAML-based Configuration (New)

Test cases can be configured using YAML files in addition to or instead of hardcoded Go structs.

#### Configuration source

Configuration is loaded from a single aggregated YAML file in the dynamically cloned CI repository:

- `shared_tests/test_cases.yml`

Each entry in this file is keyed by the test case `Name` (as defined in `tc_*.go`) and may provide values for selected fields. Only empty/unset fields in the Go struct are populated from YAML.

#### Hybrid Configuration

The YAML configuration system works as a **hybrid approach**:

1. Test cases can be partially defined in Go code (`tc_*.go` files)
2. Missing fields are automatically populated from `shared_tests/test_cases.yml` during test execution
3. Only empty/unset fields in the Go struct are populated from YAML

This allows for:

- **Gradual Migration**: Existing test cases continue to work unchanged
- **Selective Override**: Choose which fields to externalize to YAML
- **Fallback Behavior**: Tests work even without YAML files (if fully defined in Go)

### CI Integration and Dynamic Repository Cloning

The YAML-based configuration system integrates with CI systems by dynamically cloning the [konveyor/ci](https://github.com/konveyor/ci) repository during test execution.

#### CI Repository Configuration

You can customize which CI repository and branch to clone using environment variables:

- `CI_REPO_AUTHOR`: GitHub username/organization (defaults to `konveyor`)
- `CI_REPO_BRANCH`: Branch to clone (defaults to `main`)

## Tiers

Due to fragile nature of some analysis parameters for more complex applications, analysis test execution supports multiple tiers.

### Tier 0

Very basic Application->Task->Analyze->Result flow. Default for `test-analysis` target.

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

## Options

[HUB_BASE_URL](https://github.com/konveyor/tackle2-hub/tree/main/test#rest-api) is required in the same way as in API test.

Analysis tests accept `DEBUG`, `KEEP` and `PARALLEL` flags.

## Development

Test code itself is in `analysis_test.go`.

List of test cases and analyzer setup is in `test_cases.go`.

`tc_...` files contain all test cases definitions (Application and expected analysis output).
