package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/d5kx/go-autotests/internal/random"
)

var tempfileCmd = cmd{
	name:      "tempfile",
	shortHelp: "generates path to random temporary file",
	do:        generateTempfile,
}

func generateTempfile() {
	tempdir := os.TempDir()
	filename := random.ASCIIString(5, 8)
	path := filepath.Join(tempdir, filename)
	fmt.Print(path)
}
