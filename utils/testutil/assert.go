package testutil

import (
	"testing"
)

// Should checks error and if present, fails the test case.
// Example usage: testutil.Should(t, task.Create(&r))
func Should(t *testing.T, err error) {
	if err != nil {
		t.Errorf(err.Error())
	}
}

// Must checks error and if present, fails and stops the test suite.
// Example usage: testutil.Must(t, task.Create(&r))
func Must(t *testing.T, err error) {
	if err != nil {
		t.Fatalf(err.Error())
	}
}
