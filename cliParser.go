package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
	"path"
)

func parseCommandLineArgs() (ArgsParserResult, bool) {
	parser := argparse.NewParser("node", "ProgpJS impersonate for NodeJs compatibility")

	paramVersion := parser.Flag("v", "version", &argparse.Options{
		Required: false,
		Help:     "Show the version number",
	})

	paramInspect := parser.Flag("", "inspect", &argparse.Options{
		Required: false,
		Help:     "Enabled the debugger",

		Validate: func(args []string) error {
			// Allows extracting params if doing "--inspect=options".
			// Here "options" is returned, event if there is something after.
			//
			return nil
		},
	})

	paramDebug := parser.Flag("", "debug", &argparse.Options{
		Required: false,
		Help:     "Enabled the debugger",
	})

	paramInspectBreak := parser.Flag("", "inspect-brk", &argparse.Options{
		Required: false,
		Help:     "Enabled the debugger",
	})

	paramScriptToRun := parser.StringPositional(&argparse.Options{Required: false})

	// **********************************
	// **********************************
	// **********************************

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return ArgsParserResult{}, false
	}

	if *paramVersion {
		_, _ = fmt.Fprint(os.Stdout, "50.50.50\n")
		return ArgsParserResult{}, true
	}

	scriptToRun := *paramScriptToRun

	if scriptToRun == "" {
		scriptToRun = "index.js"
	}

	mustDebug := false
	cwd, _ := os.Getwd()

	if !path.IsAbs(scriptToRun) {
		scriptToRun = path.Join(cwd, scriptToRun)
	}

	if *paramInspect || *paramInspectBreak || *paramDebug {
		// Note: for jetbrains IDE, a environnement variable is set.
		// Ex: 	NODE_OPTIONS=--require /Applications/PhpStorm.app/Contents/plugins/javascript-debugger/debugConnector.js
		// It will directly use the "inspect" package, which isn't supported today.
		//
		mustDebug = true
	}

	return ArgsParserResult{
		Debug:       mustDebug,
		ScriptToRun: scriptToRun,
	}, false
}

type ArgsParserResult struct {
	Debug       bool
	ScriptToRun string
}
