package main

import (
	"fmt"
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

// builds javascript files into a single bundle
// and writes it to <root>/web/public/index.js
func main() {

	result := api.Build(api.BuildOptions{
		EntryPoints: []string{"./frontend/index.js"},
		Bundle:      true,
		Write:       true,
		Outfile:     "./public/index.js",
	})
	// fmt.Printf("result: %v\n", result.OutputFiles[0].Path)
	if len(result.Errors) != 0 {
		fmt.Printf("errors: %v\n", result.Errors)
		os.Exit(1)
	}
}
