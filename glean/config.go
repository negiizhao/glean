package glean

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	urlBase = "https://mirror.sjtu.edu.cn/elan"
)

var (
	Command = flag.String("install", "", "available: `elan`, `lean`")
	version = flag.String("version", "", "example Elan version: `3.0.0` (without the `v` prefix); example Lean version: `4.1.0`, `4.2.0-rc2` (without the `v` prefix)")
)

var (
	dotElanBaseDir = getDotElanBaseDir()
)

func getDotElanBaseDir() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Failed to get `UserHomeDir`")
	}
	dotElanBaseDir := filepath.Join(userHomeDir, ".elan")
	log.Println("got dotElanBaseDir `" + dotElanBaseDir + "`")
	return dotElanBaseDir
}

func InitFlags() {
	flag.Usage = func() {
		fmt.Println("example usage:")
		fmt.Println("\tglean -install elan -version 3.0.0")
		fmt.Println("\tglean -install lean -version v4.2.0-rc2")
		fmt.Println("Please refer to `https://mirror.sjtu.edu.cn/elan/?mirror_intel_list` for available versions")
		_, err := fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		if err != nil {
			panic("")
		}
		flag.PrintDefaults()
	}
	flag.Parse()
	if *version == "" {
		fmt.Println("invalid version")
		flag.Usage()
	}
	log.Printf("flag.Parse: got command = `%s`, version = `%s`", *Command, *version)
}