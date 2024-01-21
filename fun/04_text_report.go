package fun

import (
	"log"
	"os"

	"github.com/parsiya/semgrep_go/output"
)

func TextReport(path string) error {

	// Instead of running Semgrep, we will use the output from example 00 in
	// `output/juice-shop.json`.
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	// Deserialize the data.
	out, err := output.Deserialize(data)
	if err != nil {
		return err
	}

	// Create the reports.
	ruleIDTextReport := out.RuleIDTextReport(true)
	filePathTextReport := out.FilePathTextReport(true)

	// Print the reports.
	log.Print("Rule ID report:")
	log.Print(ruleIDTextReport)

	log.Print("File Path Report:")
	log.Print(filePathTextReport)
	return nil
}
