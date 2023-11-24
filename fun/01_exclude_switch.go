package fun

import (
	"fmt"
	"log"
	"strings"

	"github.com/parsiya/semgrep_go/run"
)

func ExcludeSwitch() error {

	// In the real world we will get a long list from somewhere.
	excludedRules := []string{
		"javascript.audit.detect-replaceall-sanitization.detect-replaceall-sanitization",
	}
	var extra []string
	// Add all the excluded rules like a `--exclude-rule=[ID]` argument.
	for _, r := range excludedRules {
		extra = append(extra, "--exclude-rule="+r)
	}

	log.Printf("Excluding results for: %s", excludedRules)

	// Setup Semgrep switches.
	opts := run.Options{
		Output:    run.JSON, // Output format is JSON.
		Paths:     []string{"code/juice-shop"},
		Rules:     []string{"p/default"},
		Verbosity: run.Debug,
		Extra:     extra, // Items in Extra will be added to the CLI as-is.
	}

	log.Print("Running Semgrep, this might take a minute.")
	// Run Semgrep and get the deserialized output.
	out, err := opts.RunJSON()
	if err != nil {
		return err
	}

	// Check if any of the ruleIDs match what we wanted to exclude.
	for _, hit := range out.Results {
		if strings.Contains(hit.RuleID(), "detect-replaceall-sanitization") {
			return fmt.Errorf("Found a rule that should have been excluded.")
		}
	}
	// Print a table of ruleIDs and hits to show we have filtered the results.
	log.Print("Results:")
	log.Print(out.RuleIDStringTable(true))

	return nil
}
