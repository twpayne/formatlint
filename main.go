package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var ignoreDirs = map[string]bool{
	".git":   true,
	"vendor": true,
}

var extParseCheckers = map[string][]ParseChecker{
	".go": {
		goParseChecker{},
	},
	".json": {
		jsonParseChecker{},
	},
	".sh": {
		shParseChecker{},
	},
	".toml": {
		tomlParseChecker{},
	},
	".xml": {
		xmlParseChecker{},
	},
	".yaml": {
		yamlParseChecker{},
	},
	".yml": {
		yamlParseChecker{},
	},
}

func parseCheckFile(filename string) error {
	parseCheckers := extParseCheckers[filepath.Ext(filename)]
	if len(parseCheckers) == 0 {
		return nil
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	for _, parseChecker := range parseCheckers {
		parseCheckerPrefix := parseChecker.Name() + ": "
		if err := parseChecker.Parse(data); err != nil {
			s := err.Error()
			if !strings.HasPrefix(s, parseCheckerPrefix) {
				s = parseCheckerPrefix + s
			}
			return errors.New(s)
		}
	}

	return nil
}

func findAll(root string) ([]string, error) {
	var paths []string
	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && ignoreDirs[filepath.Base(path)] {
			return filepath.SkipDir
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		paths = append(paths, path)
		return nil
	}); err != nil {
		return nil, err
	}
	return paths, nil
}

func run() error {
	flag.Parse()

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	args := flag.Args()
	if len(args) == 0 {
		args = []string{wd}
	}

	uniquePaths := make(map[string]bool)
	for _, arg := range args {
		info, err := os.Stat(arg)
		if err != nil {
			return err
		}
		switch {
		case info.IsDir():
			paths, err := findAll(arg)
			if err != nil {
				return err
			}
			for _, path := range paths {
				uniquePaths[path] = true
			}
		case info.Mode().IsRegular():
			uniquePaths[arg] = true
		default:
			return fmt.Errorf("%s: not a regular file or directory", arg)
		}
	}
	paths := make([]string, 0, len(uniquePaths))
	for path := range uniquePaths {
		paths = append(paths, path)
	}
	sort.Strings(paths)

	errors := 0
	for _, path := range paths {
		if err := parseCheckFile(path); err != nil {
			errors++
			fmt.Printf("%s: %v\n", path, err)
		}
	}

	if errors > 0 {
		os.Exit(1)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
