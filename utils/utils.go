package utils

import "os/exec"

// LastString returns the last element of the given slice of strings.
func LastString(str []string) string {
	return str[len(str)-1]
}

// runShellCommand executes a shell command and returns an error if it fails
func RunShellCommand(cmd string) error {
	command := exec.Command("bash", "-c", cmd)
	return command.Run()
}
