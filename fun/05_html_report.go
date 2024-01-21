package fun

import (
	"bytes"
	_ "embed"
	"html/template"
	"log"
	"os"

	"github.com/parsiya/semgrep_go/output"
)

// embed the 05-report-html-tmpl.html in a string.
//
//go:embed 05-report-html-template.html
var tmpl string

// Report contains the information in the HTML report.
type Report struct {
	NumberOfFindings int
	ByRuleID         []output.HitMapRow
	ByFilePath       []output.HitMapRow
}

func HTMLReport(path string) error {

	// Instead of running Semgrep, we will use the output from example 00 in
	// `output/juice-shop.json`.
	dat, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	// Deserialize the data.
	out, err := output.Deserialize(dat)
	if err != nil {
		return err
	}

	// Create the report object.
	rep := Report{
		NumberOfFindings: len(out.Results),
		ByRuleID:         out.RuleIDHitMap(true),
		ByFilePath:       out.FilePathHitMap(true),
	}

	// Apply the template.
	t, err := template.New("report").Parse(tmpl)
	if err != nil {
		return err
	}
	var data bytes.Buffer
	if err = t.Execute(&data, rep); err != nil {
		return err
	}

	// Write the report to a file.

	// We're gonna hardcode the report file for this example instead of passing
	// a second argument.
	log.Print("Writing the output to output/05-report.html")
	return WriteFile("output/05-report.html", data.Bytes())
}
