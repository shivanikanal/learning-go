package main

/*
   GO_TIP: go.sum file is generated/appended only when external modules/packages are used, and used to track external dependencies
*/
import (
	"fmt"
	"os"

	"rsc.io/quote"
)

/*
   GO_TIP: main function executes by default when you run main package
*/
func main() {
	fmt.Println(quote.Go())

	useFprintf()
}

func useFprintf() {
	const name, age = "Kim", 22

	/*
	   GO_TIP: `:=` is shortcut of declaring and initializing variables in go, value on right is used to determine type of variable
	*/
	n, err := fmt.Fprintf(os.Stdout, "%s is %d years old.\n", name, age)

	// The n and err return values from Fprint are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
	fmt.Print(n, " bytes written.\n")
}
