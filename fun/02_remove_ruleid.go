package fun

import (
	"log"

	"golang.org/x/exp/slices"

	"github.com/parsiya/semgrep_go/output"
	"github.com/parsiya/semgrep_go/run"
)

func ExcludeRuleID() error {

	// In the real world we will get a long list from somewhere.
	excludedRules := []string{
		"javascript.audit.detect-replaceall-sanitization.detect-replaceall-sanitization",
	}

	// Setup Semgrep switches. We're not adding anything to extras this time.
	opts := run.Options{
		Output:    run.JSON, // Output format is JSON.
		Paths:     []string{"code/juice-shop"},
		Rules:     []string{"p/default"},
		Verbosity: run.Debug,
	}

	log.Print("Running Semgrep, this might take a minute.")
	// Run Semgrep and get the deserialized output.
	out, err := opts.RunJSON()
	if err != nil {
		return err
	}

	// Go through the results and remove any that match the excluded rules.

	// Create a new slice to hold the modified results.
	var modifiedResults []output.CliMatch

	// Loop through the results.
	for _, hit := range out.Results {
		// Check if the hit's ruleID matches any of the excluded rules.
		if slices.Contains(excludedRules, hit.RuleID()) {
			// If it does, skip it.
			continue
		}
		// If the ruleID is not excluded, add it to modifiedResults.
		modifiedResults = append(modifiedResults, hit)
	}

	// Replace the results with the modified results.
	out.Results = modifiedResults
	log.Print("Results:")
	log.Print(out.RuleIDStringTable(true))

	// Optionally, we can serialize the output back to JSON.
	js, err := out.Serialize(true)
	if err != nil {
		return err
	}
	// Do something with js which is a []byte.
	_ = js

	return nil
}
