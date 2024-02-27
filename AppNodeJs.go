/*
 * (C) Copyright 2024 Johan Michel PIQUET, France (https://johanpiquet.fr/).
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"github.com/progpjs/modules/v2/modCore"
	"github.com/progpjs/modules/v2/modHttp"
	"github.com/progpjs/modules/v2/modNodeJs"
	"github.com/progpjs/modules/v2/modReact"
	"os"
	"path"
	"strings"
)

func resolveScriptPath(scriptPath string) string {
	stat, err := os.Stat(scriptPath)

	if err != nil {
		if strings.HasSuffix(scriptPath, ".js") {
			scriptPath = scriptPath[0:len(scriptPath)-3] + ".ts"
			return resolveScriptPath(scriptPath)
		}

		return ""
	}

	if stat.IsDir() {
		scriptPath = path.Join(scriptPath, "index.js")
		return resolveScriptPath(scriptPath)
	}

	return scriptPath
}

func main() {
	runArgs, mustExit := parseCommandLineArgs()
	if mustExit {
		return
	}

	scriptPath := runArgs.ScriptToRun
	cwd, _ := os.Getwd()
	if !path.IsAbs(scriptPath) {
		scriptPath = path.Join(cwd, scriptPath)
	}

	scriptPath = resolveScriptPath(runArgs.ScriptToRun)
	if scriptPath == "" {
		println("Script not found: ", scriptPath)
		os.Exit(1)
		return
	}

	awaiter := bootstrapProgpJS(scriptPath, runArgs.Debug, nil, RegisterMyModules)

	// Will wait until the VM can exit.
	//
	// The VM can't exit if there is background task remaining (ex: a webserver).
	// If you don't call awaiter then your app will quit immediately and your server
	// will not be able to execute.
	//
	awaiter()
}

// RegisterMyModules registers our all ProgpJS modules.
func RegisterMyModules() {
	// Required core modules.
	//
	modCore.InstallProgpJsModule()
	modNodeJs.InstallProgpJsModule()

	// Optional core modules.
	//
	modReact.InstallProgpJsModule()
	modHttp.InstallProgpJsModule()
	//progpJsonDB.InstallProgpJsModule()
}
