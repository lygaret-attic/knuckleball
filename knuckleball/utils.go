package knuckleball

import (
	"fmt"
	"strings"
)

func read_params(str string) map[string]string {

	params := make(map[string]string)
	allparts := strings.Split(str, " ")
	for _, p := range allparts {
		pparts := strings.Split(p, ":")
		if len(pparts) == 2 {
			params[trim_up(pparts[0])] = trim(pparts[1])
		}
	}

	return params
}

func read_cmd(line *string) (cmd string, rest string, err error) {

	parts := strings.SplitN(*line, " ", 2)
	if parts == nil || len(parts) == 0 {
		err = fmt.Errorf("Invalid command string found: %v", *line)
		return
	}

	cmd = trim_up(parts[0])
	if len(parts) == 2 {
		rest = trim(parts[1])
	}

	return
}

func trim(str string) string {
	return strings.TrimSpace(str)
}

func trim_up(str string) string {
	return strings.ToUpper(strings.TrimSpace(str))
}
