// In Go it’s idiomatic to communicate errors via an explicit, separate return value.
// This contrasts with the exceptions used in languages like Java and Ruby and the overloaded single result / error value sometimes used in C.
// Go’s approach makes it easy to see which functions return errors and to handle them using the same language constructs employed for other, non-error tasks.

package main

import (
	"errors"
	"fmt"
)

// By convention, errors are the last return value and have type error, a built-in interface.
func f(arg int) (int, error) {

	// errors.New constructs a basic error value with the given error message.
	if 42 == arg {
		return -1, errors.New("can't work with 42")
	}

	// A nil value in the error position indicates that there was no error.
	return arg + 3, nil
}

// A sentinel error is a predeclared variable that is used to signify a specific error condition.
var errOutOfTea = fmt.Errorf("no more tea available")
var errPower = fmt.Errorf("can't boil water")

func MakeTea(arg int) error {
	if 2 == arg {
		return errOutOfTea
	} else if 4 == arg {
		// We can wrap errors with higher-level errors to add context.
		// The simplest way to do this is with the %w verb in fmt.Errorf.
		// Wrapped errors create a logical chain (A wraps B, which wraps C, etc.) that can be queried with functions like errors.Is and errors.As.
		return fmt.Errorf("making tea: %w", errPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {

		// It’s common to use an inline error check in the if line.
		if res, err := f(i); err != nil {
			fmt.Println("f failed: ", err)
		} else {
			fmt.Println("f worked: ", res)
		}
	}

	for i := range 5 {
		if err := MakeTea(i); err != nil {
			if errors.Is(err, errOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, errPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("Unknown error: %s\n", err)
			}

			continue
		}

		fmt.Println("Tea is ready!")
	}
}
