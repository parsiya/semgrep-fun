package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/parsiya/semgrep_fun/fun"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("expecting a subcommand.")
		os.Exit(1)
	}

	// Disable log timestamps.
	log.SetFlags(0)

	subCommand := strings.ToLower(os.Args[1])
	log.Printf("Running subcommand: %s", subCommand)

	switch subCommand {
	case "01-exclude-switch":
		if err := fun.ExcludeSwitch(); err != nil {
			panic(err)
		}
	case "02-exclude-ruleid":
		if err := fun.ExcludeRuleID(); err != nil {
			panic(err)
		}
	case "03-go-test-coverage":
		if err := fun.GoTestCoverage(); err != nil {
			panic(err)
		}
	default:
		fmt.Println("use one of the subcommands: ZZZ list subcommands here?")
	}

}
