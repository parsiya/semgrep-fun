package fun

import (
	_ "embed"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/olekukonko/tablewriter"
	"github.com/parsiya/semgrep_go/run"
)

// Embedding the rule as a string for demonstration purposes.
//
//go:embed 03-go-test-coverage.yaml
var rule string

// FuncInfo contains information about a function.
type FuncInfo struct {
	Package string
	Name    string
	Path    string
}

// FuncMap is a map where key is function name and value is FuncInfo struct.
type FuncMap map[string]FuncInfo

// FuncList is a map where key is the package name and the value is FuncMap.
type FuncList map[string]FuncMap

func GoTestCoverage(path string) error {

	// Create an empty .semgrepignore file.
	file, err := os.Create(".semgrepignore")
	if err != nil {
		panic(err)
	}
	file.Close()
	// Delete the file after the function returns.
	defer os.Remove(".semgrepignore")

	// Setup Semgrep options.
	opts := run.Options{
		Output:    run.JSON, // Output format is JSON.
		Paths:     []string{path},
		Rules:     []string{rule}, // Add the rule as a string.
		Verbosity: run.Debug,
	}
	// Concat all the rules together in a temp file and pass to Semgrep.
	opts.StringRule()

	log.Print("Running Semgrep, this might take a minute.")
	// Run Semgrep and get the deserialized output.
	out, err := opts.RunJSON()
	if err != nil {
		return err
	}

	funcList := make(FuncList)

	// Loop through the results.
	for _, hit := range out.Results {
		// $PKG - $FUNC
		msg := strings.Split(hit.Message(), " - ")
		// msg[0]: $PKG
		// msg[1]: $FUNC

		// Note we're not doing a lot of error checking here.
		if len(msg) != 2 {
			log.Printf("Wrong message, got: %s", hit.Message())
			continue
		}
		// Store the function info in a FuncInfo struct.
		fn := FuncInfo{
			Package: msg[0],
			Name:    msg[1],
			Path:    hit.FilePath(),
		}
		if _, ok := funcList[fn.Package]; !ok {
			funcList[fn.Package] = make(FuncMap)
		}
		funcList[fn.Package][fn.Name] = fn
	}

	data := make([][]string, 0)

	// Now, we have a map of all functions in code, we can go through the
	// functions in each package and check if they have a test.
	for _, funcs := range funcList {

		// It's easier to create a slice of all functions in a package for
		// searching.
		var funcNames []string
		for _, fn := range funcs {
			funcNames = append(funcNames, fn.Name)
		}

		for _, fn := range funcs {
			// Skip functions in `*_test.go` files.
			if strings.HasSuffix(fn.Path, "_test.go") {
				continue
			}

			// Check if the function has a test. AKA "Test"+fn.Name is in the
			// package.
			if slices.Contains(funcNames, "Test"+fn.Name) {
				continue
			}
			// Add the functions with missing tests to the data slice.
			data = append(data, []string{fn.Name, fn.Package, fn.Path})
		}
	}

	// Create the table.

	// String builder to hold the result.
	var final strings.Builder
	// Create the table writer and set the destination to the string builder.
	table := tablewriter.NewWriter(&final)
	// Set the headers.
	table.SetHeader([]string{"Name", "Package", "Path"})
	// Append the data.
	table.AppendBulk(data)
	// Render the table.
	table.Render()
	// Print the table
	log.Println(final.String())

	return nil
}
