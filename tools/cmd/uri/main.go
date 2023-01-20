package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/struct-ure/kg/tools/util"

	"github.com/pkg/errors"
)

// Converts a path to the struct-ure URI format
func main() {
	if len(os.Args) != 2 {
		panic("Usage: uri <path-to-convert>")
	}

	path := os.Args[1]
	uri := util.URIFromPath(path)
	if !strings.HasPrefix(uri, util.URIDomainPrefix()) {
		panic(errors.Errorf("Invalid uri from path %s", os.Args[1]))
	}
	fmt.Println(uri)
}
