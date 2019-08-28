package iteration

import (
	"strings"
)


// Takes a string c and replicates it length times
func Repeat(c string, length int) string {
	return strings.Repeat(c, length)
}