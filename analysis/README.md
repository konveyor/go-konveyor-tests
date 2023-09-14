# Application analysis tests

Run with ```$ make test-analysis```

## Stages

Due to fragile nature of some analysis parameters for more complex applications, analysis test execution supports multiple stages.

### Stage 0

Very basic Application->Task->Analyze->Result flow. Should never fail.

### Stage 1

Common example application analysis. Should work.

```
STAGE1=1 make test-analysis
```

### Stage 2

More advanced analysis parameters. Great if works.

```
STAGE2=1 make test-analysis
```

## Options

[HUB_BASE_URL](https://github.com/konveyor/tackle2-hub/tree/main/test#rest-api) is required in the same way as in API test.

### KEEP

Don't delete created applications and analysis, helpful for debugging test failures.

```
$ KEEP=1 make test-analysis
```

### PARALLEL

Run analysis tests in paralel

```
$ PARALLEL=1 make test-analysis
```

## Development

Test code itself is in ```analysis_test.go```.

List of test cases and analyzer setup is in ```test_cases.go```.

```tc_...``` files contain all test cases definitions (Application and expected analysis output).
