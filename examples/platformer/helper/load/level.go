package load

import (
	"regexp"
	"strings"
)

// Level converts the level written in a string to another string consisting
// of only non-whitespace characters, which can be accessed by coordinates.
// It also returns the dimensions of the level.
func Level(str string) (w, h int, level string) {
	trimmer := regexp.MustCompile(`\n[\t\s]+`)
	cleared := strings.TrimLeft(trimmer.ReplaceAllString(str, "\n"), "\n")

	w, h = strings.Index(cleared, "\n"), strings.Count(cleared, "\n")
	level = strings.Replace(cleared, "\n", "", -1)

	return
}
