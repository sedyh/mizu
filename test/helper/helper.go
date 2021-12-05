package helper

import (
	"fmt"
	"regexp"
	"runtime/debug"

	"mizu/pkg/engine"
)

type SingleEntityScene struct {
	Entity interface{}
}

func (s *SingleEntityScene) Setup(w engine.World) {
	w.AddEntities(s.Entity)
}

func FailMessage(err interface{}) string {
	message := ShortPanicMessage()
	if message == "" {
		return "Expected panic not to be nil"
	}

	return fmt.Sprintf(
		"Expected:\n\t\t%s\n\t%s\nto be other panic",
		err, message,
	)
}

func ShortPanicMessage() string {
	group := "trace"
	re := regexp.MustCompile(`panic.go.+\n.+\n(?P<` + group + `>.+)`)

	matches := re.FindStringSubmatch(string(debug.Stack()))
	if len(matches) == 0 {
		return ""
	}

	return matches[re.SubexpIndex(group)]
}
