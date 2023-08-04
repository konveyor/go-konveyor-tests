# Application analysis tests

Run with ```$ make test-analysis```

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
