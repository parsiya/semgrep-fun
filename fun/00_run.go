package fun

import (
	"log"

	"github.com/parsiya/semgrep_go/run"
)

func RunSemgrep(path string) error {

	// Setup Semgrep switches.
	opts := run.Options{
		Output:    run.JSON,       // Output format is JSON.
		Paths:     []string{path}, // "code/juice-shop"
		Rules:     []string{"p/default"},
		Verbosity: run.Debug,
		Extra:     []string{"--output=output/juice-shop.json"},
	}

	log.Print("Running Semgrep, this might take a minute.")
	// Run Semgrep and ignore the output.
	_, err := opts.Run()
	return err
}
