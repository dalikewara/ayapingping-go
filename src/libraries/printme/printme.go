// Example library/helper package.

package printme

import "fmt"

// Print prints and returns the given text.
func Print(text string) string {
	fmt.Println(text)
	return text
}
