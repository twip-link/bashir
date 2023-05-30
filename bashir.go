// Package bashir calls bash with exec.Command for one or more commands, and returns the output.
package bashir

import (
	"fmt"
	"os/exec"
)

func Bash(command ...string) []string {
	result := []string{}

	for i := 0; i < len(command); i++ {
		result = append(result, bash(command[i]))
	}

	return result
}

func bash(s string) string {
	out, err := exec.Command("bash", "-c", s).Output()
	if err != nil {
		return fmt.Sprintf("insufficiently bashed: %s\n%s", s, err.Error())
	}
	return string(out)
}
