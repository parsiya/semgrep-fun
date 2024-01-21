package fun

import (
	"log"
	"os"

	"golang.org/x/exp/slices"

	"github.com/parsiya/semgrep_go/output"
)

func ExcludeRuleID(path string) error {

	// In the real world we will get a long list from somewhere.
	excludedRules := []string{
		"javascript.audit.detect-replaceall-sanitization.detect-replaceall-sanitization",
	}

	// Instead of running Semgrep, we will use the output from example 00 in
	// `output/juice-shop.json`.
	data, err := os.ReadFile("output/juice-shop.json")
	if err != nil {
		return err
	}
	// Deserialize the data.
	out, err := output.Deserialize(data)
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
	log.Print(out.RuleIDTextReport(true))

	// Optionally, we can serialize the output back to JSON.
	js, err := out.Serialize(true)
	if err != nil {
		return err
	}
	// Do something with js which is a []byte.
	_ = js

	return nil
}
