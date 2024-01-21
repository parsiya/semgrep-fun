package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/parsiya/semgrep_fun/fun"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Wrong number of arguments.")
		fmt.Println("Usage: semgrep_fun <subcommand> <path>")
		os.Exit(1)
	}

	// Disable log timestamps.
	log.SetFlags(0)

	subCommand := strings.ToLower(os.Args[1])
	log.Printf("Running subcommand: %s", subCommand)

	path := os.Args[2]

	switch subCommand {

	case "00":
		if err := fun.RunSemgrep(path); err != nil {
			panic(err)
		}
	case "01":
		if err := fun.ExcludeSwitch(path); err != nil {
			panic(err)
		}
	case "02":
		if err := fun.ExcludeRuleID(path); err != nil {
			panic(err)
		}
	case "03":
		if err := fun.GoTestCoverage(path); err != nil {
			panic(err)
		}
	case "04":
		if err := fun.TextReport(path); err != nil {
			panic(err)
		}
	case "05":
		if err := fun.HTMLReport(path); err != nil {
			panic(err)
		}
	case "06":
		if err := fun.FunctionChain(path); err != nil {
			panic(err)
		}
	default:
		fmt.Println("use one of the subcommands: ZZZ list subcommands here?")
	}

}
