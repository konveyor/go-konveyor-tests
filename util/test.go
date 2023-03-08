package util

import "github.com/konveyor/tackle2-hub/api"

type Test struct {
	Name        string
	Subject     interface{}
	Application api.Application
	Task        api.Task
	TaskData    string

	ShouldError bool
}
