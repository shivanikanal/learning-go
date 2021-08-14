/*
	GO_TIP: Ending a file's name with _test.go tells the go test command that this file contains test functions.
*/

package greetings

import (
	"regexp"
	"testing"
)

// tests greetings.Hello for a valid return
/*
	GO_TIP: test functions takes a pointer to testing packages's testing.T type as param
*/
func TestHelloName(t *testing.T) {
	name := "Shivani"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Shivani")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Shivani") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// tests greetings.Hello failure case with an empty string
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
