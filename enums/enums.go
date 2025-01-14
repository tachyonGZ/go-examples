// Enumerated types (enums) are a special case of sum types.
// An enum is a type that has a fixed number of possible values, each with a distinct name.
// Go doesn’t have an enum type as a distinct language feature, but enums are simple to implement using existing language idioms.

package main

import "fmt"

// Our enum type ServerState has an underlying int type.
type ServerState int

// The possible values for ServerState are defined as constants.
// The special keyword iota generates successive constant values automatically;
// in this case 0, 1, 2 and so on.
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// By implementing the fmt.Stringer interface, values of ServerState can be printed out or converted to strings.
//
// This can get cumbersome if there are many possible values.
// In such cases the stringer tool can be used in conjunction with go:generate to automate the process.
// See this post for a longer explanation.
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

// transition emulates a state transition for a server;
// it takes the existing state and returns a new state.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:

		// Suppose we check some predicates here to determine the next state…
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

func main() {

	// If we have a value of type int, we cannot pass it to transition - the compiler will complain about type mismatch.
	// This provides some degree of compile-time type safety for enums.
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}
