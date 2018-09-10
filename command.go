package bitbar

import (
	"fmt"
	"strconv"
	"strings"
)

type Command struct {
	path     string
	params   []string
	terminal bool
	refresh  bool
}

func NewCommand(path string, params []string) *Command {
	return &Command{path: path, params: params}
}

func (s *Command) string() string {

	params := []string{}
	for idx, param := range s.params {
		params = append(params, fmt.Sprintf("param%d=\"%s\"", idx+1, param))
	}

	return fmt.Sprintf(
		"bash=\"%s\" %s terminal=%s refresh=%s",
		s.path,
		strings.Join(params, " "),
		strconv.FormatBool(s.terminal),
		strconv.FormatBool(s.refresh),
	)
}
