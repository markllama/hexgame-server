//
// Hexgame Server
//
package main

import (
	"fmt"
	"flag"
	//	"os"

	"github.com/markllama/hexgame-server/service"
)

var opts struct {
	debug bool
	verbose bool
	url string
	port int
}

// ---------------------------------------------------------------------------
// Configuration and environment variables
// ---------------------------------------------------------------------------

// verbose: HEXGAME_VERBOSE
// debug:   HEXGAME_DEBUG
// url:     HEXGAME_URL
// port:    HEXGAME_PORT
//

func init() {
	flag.BoolVar(&opts.verbose, "verbose", false, "provide verbose output")
	flag.BoolVar(&opts.debug, "debug", false, "provide internal debugging output")
	flag.StringVar(&opts.url, "url", "http://localhost", "server URL")
	flag.IntVar(&opts.port, "port", 3000, "server port")
}

// arguments take precedence. Set those first.
func process_arguments() {
	fmt.Println("Processing arguments")
	flag.Parse()
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

	if opts.debug {
		fmt.Println("Options: ", opts)
	}

	game_service := service.NewServer()
	game_service.Run(fmt.Sprintf(":%d", opts.port))
	
	fmt.Println("-- End Program --");
}
