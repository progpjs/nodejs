package main

import (
	"github.com/progpjs/progpjs/v2"
	"os"
)

var gBootstrapHook func(scriptPath string, enableDebug bool, options *progpjs.EngineOptions, installMods func()) progpjs.BootstrapExitAwaiterF

// bootstrapProgpJS will start ProgpJS engine.
//
// How this function is build allows to easily enable/disable support for javascript,
// the main interest being to make compilation much faster, also than having
// an executable starting faster. It's very interesting when developing things requiring
// a lot of small test / debug / ...
func bootstrapProgpJS(scriptPath string, enableDebug bool, options *progpjs.EngineOptions, installMods func()) progpjs.BootstrapExitAwaiterF {
	if gBootstrapHook == nil {
		return bootstrapWithoutJavascript(scriptPath, enableDebug, options, installMods)
	} else {
		return gBootstrapHook(scriptPath, enableDebug, options, installMods)
	}
}

func bootstrapWithoutJavascript(_ string, _ bool, _ *progpjs.EngineOptions, installMods func()) progpjs.BootstrapExitAwaiterF {
	installMods()

	// Rebuild the generated code this the project "progpjs.progpV8Engine" is found.
	//
	progpV8Path := os.Getenv("PROGPV8_DIR")
	if progpV8Path == "" {
		progpV8Path = os.Getenv("DEV_PROGPV8_DIR")
	}

	stat, err := os.Stat(progpV8Path)
	if (err == nil) && (stat.IsDir()) {
		progpjs.GenerateSourceCode(progpV8Path)
	}

	return func() {}
}
