package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/anothermemory/server/pkg/commands"
	"github.com/spf13/cobra/doc"
)

func main() {
	const defaultDir = "./docs/commands"
	dir := flag.String("out", defaultDir, "Target directory where documenation will be generated")
	help := flag.Bool("help", false, "Display this help")

	md := flag.Bool("markdown", false, "Generate markdown docs")
	rst := flag.Bool("rst", false, "Generate RST docs")
	man := flag.Bool("man", false, "Generate MAN docs")
	yaml := flag.Bool("yaml", false, "Generate Yaml docs")

	flag.Parse()

	if *help == true {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Printf("Output dir is set to: %s\n", *dir)
	if *md {
		genMarkdown(*dir)
	}
	if *rst {
		genRST(*dir)
	}
	if *man {
		genMAN(*dir)
	}
	if *yaml {
		genYaml(*dir)
	}
}

func genMarkdown(baseDir string) {
	targetDir := path.Join(baseDir, "md")
	os.MkdirAll(targetDir, os.ModePerm)

	fmt.Printf("Generating markdown documents in: %s\n", targetDir)

	err := doc.GenMarkdownTree(commands.RootCmd, targetDir)
	if err != nil {
		fmt.Printf("Failed to generate markdown documents: %s", err)
		os.Exit(1)
	}
}

func genRST(baseDir string) {
	targetDir := path.Join(baseDir, "rst")
	os.MkdirAll(targetDir, os.ModePerm)

	fmt.Printf("Generating RST documents in: %s\n", targetDir)

	err := doc.GenReSTTree(commands.RootCmd, targetDir)
	if err != nil {
		fmt.Printf("Failed to generate RST documents: %s", err)
		os.Exit(1)
	}
}

func genMAN(baseDir string) {
	targetDir := path.Join(baseDir, "man")
	os.MkdirAll(targetDir, os.ModePerm)

	fmt.Printf("Generating MAN documents in: %s\n", targetDir)

	err := doc.GenManTree(commands.RootCmd, &doc.GenManHeader{}, targetDir)
	if err != nil {
		fmt.Printf("Failed to generate MAN documents: %s", err)
		os.Exit(1)
	}
}

func genYaml(baseDir string) {
	targetDir := path.Join(baseDir, "yaml")
	os.MkdirAll(targetDir, os.ModePerm)

	fmt.Printf("Generating Yaml documents in: %s\n", targetDir)

	err := doc.GenYamlTree(commands.RootCmd, targetDir)
	if err != nil {
		fmt.Printf("Failed to generate Yaml documents: %s", err)
		os.Exit(1)
	}
}
