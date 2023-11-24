package fun

import (
	_ "embed"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/parsiya/semgrep_go/run"
)

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

func GoTestCoverage() error {

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
		Paths:     []string{"code/lo"},
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
			log.Printf("Wrong messgae, got: %s", hit.Message())
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

	// Now, we have a map of all functions in code, we can go through the
	// functions in each package and check if they have a test.
	for pkg, funcs := range funcList {

		// Easier to create a string of all functions in a package.
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
			log.Printf("Function %s in package %s has no test", fn.Name, pkg)
		}
	}

	return nil
}
