package fun

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/parsiya/semgrep_go/run"
)

// Embed the rules.
//
//go:embed 06-go-func-chain.yaml
var rule06 string

// ImportMap is map of one file's imports.
//
// Key: Alias if it exists or the top-level package if not (e.g., "bar" for
// "github.com/foo/bar").
//
// Value: Complete package name.
type ImportMap map[string]string

// Imports is a map of file paths to their imports. Key is file path.
type Imports map[string]ImportMap

// Function represents one function.
type Function struct {
	Package  string
	Name     string
	FilePath string
}

// Stringer for Function that returns the struct as a JSON string using
// json.Marshal.
func (f Function) String() string {
	b, _ := json.Marshal(f)
	return string(b)
}

func FunctionChain(path string) error {

	// Setup Semgrep switches.
	opts := run.Options{
		Output:    run.JSON, // Output format is JSON.
		Paths:     []string{path},
		Rules:     []string{rule06}, // pass the string
		Verbosity: run.Debug,
		Extra:     []string{"--no-rewrite-rule-ids"}, // Do not rewrite rule IDs.
	}
	// Contents of the Rules field is a string that should be stored in a temp
	// file and then passed to Semgrep.
	opts.StringRule()

	// Run Semgrep.
	out, err := opts.RunJSON()
	if err != nil {
		return err
	}

	// Map to track file imports. Key is filepath.
	imports := make(Imports)

	// Map to track function chains.
	functions := make(map[Function][]Function)

	// Process the results by rule ID.
	for _, hit := range out.Results {
		switch hit.RuleID() {
		case "go-import-collection":
			// Message format: $ALIAS - $IMPORT
			msgBits := strings.Split(hit.Message(), " - ")

			// If alias exists (AKA it's not `$ALIAS`), set it.
			alias := ""
			if msgBits[0] != "$ALIAS" {
				alias = msgBits[0]
			} else {
				// Otherwise, set the alias to the top level package name.
				// Split the import by "/" and take the last element.
				alias = TopLevelPackage(msgBits[1])
			}

			// Check if the inner map (ImportMap) exists. If not, create it.
			if _, ok := imports[hit.FilePath()]; !ok {
				imports[hit.FilePath()] = make(ImportMap)
			}
			// Add the complete import to the inner map.
			imports[hit.FilePath()][alias] = msgBits[1]

		case "go-function-chain":
			// Message format: $PKG - $CALLER - $CALLEE - $IMP
			msgBits := strings.Split(hit.Message(), " - ")

			// If the import is empty (it's `$IMP`), it's in the same package.
			calleePackage := msgBits[0]
			if msgBits[3] != "$IMP" {
				// Otherwise, set the import.
				calleePackage = msgBits[3]
			}

			callee := Function{calleePackage, msgBits[2], hit.FilePath()}
			caller := Function{msgBits[0], msgBits[1], hit.FilePath()}

			functions[caller] = append(functions[caller], callee)
		}
	}

	// We're done parsing the rules, but, we need to go through all callees and
	// update their imports with the complete package name (and replace aliases
	// if any). We couldn't do this during rule parsing because the import rules
	// are not completely parsed so `imports` is incomplete.

	// This is a little bit better than parsing the imports first and then
	// doing the functions, because the number of callees are fewer than the
	// total number of hits. But ultimately, it really doesn't matter because
	// most of the processing time is used by Semgrep.

	for _, calleeList := range functions {
		for _, callee := range calleeList {
			// Get the imports for callee's file path.
			calleeImports, ok := imports[callee.FilePath]
			if !ok {
				fmt.Printf("callee imports not found for %s\n", callee.FilePath)
			}
			// Check if the import exists in the map.
			if calleeImport, ok := calleeImports[callee.Package]; ok {
				// If it does, replace the alias with the complete package name.
				callee.Package = calleeImport
			}
		}
	}

	// Now we have one level of checks. We have the info for populating the chains
	// for children too, but I need to think how to do it.

	// Let's print what we have for now.

	for caller, calleeList := range functions {
		fmt.Printf("Caller: %s\n", caller)
		for _, callee := range calleeList {
			fmt.Printf("  Callee: {\"Package\":\"%s\",\"Name\":\"%s\"}\n", callee.Package, callee.Name)
		}
		fmt.Println()
	}
	return nil
}
