//
// Hexgame Server
//
package main

import (
	"fmt"
//	"os"
)

// ---------------------------------------------------------------------------
// Configuration and environment variables
// ---------------------------------------------------------------------------


// arguments take precedence. Set those first.
func process_arguments() {
	fmt.Println("Processing arguments")
}

// Set from environment if not set in args
func process_environment() {
	fmt.Println("Processing envronment")
}

// andy required value that has not been set and has a default gets set here
func process_defaults() {
	fmt.Println("Processing defaults")
}
// ---------------------------------------------------------------------------
// MAIN
// ---------------------------------------------------------------------------


func main() {
	fmt.Println("-- Start Program --");
	process_arguments()
	process_environment()
	process_defaults()
	fmt.Println("-- End Program --");
}
