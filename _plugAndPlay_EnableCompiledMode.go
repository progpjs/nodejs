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
	"github.com/progpjs/progpV8Engine/v2"
	"os"
	"path"
)

// ===  Plug & Play file  ===   Enable compile mode  ===
//
// - When enabled (file named "plugAndPlay_enableCompileMode.go)
//		==> ProgpJS uses "compiled mode" which mean that he will generate fast and optimized code
//			in order to bind your Go function to the V8 javascript engine.
// - When disable (file named "_plugAndPlay_enableCompileMode.go)
//		==> ProgpJS will uses "plugin mode" where he will load the project ProgpV8Engine as an external library.
//			It's something like a Windows "dll" but for Go.
//			-> When this mode is enable he will search the file insde "../plugins/progpV8.so".
//			-> To rebuild this file use this bash scripts:
//					rm ./_plugins/*
//					cp ./progpjs.progpV8Engine/_generatedBlank/* ./progpjs.progpV8Engine
//					go build -buildmode=plugin -gcflags='all=-N -l' -o ./_plugins/progpV8.so ./progpjs.progpV8Engine/asInstaller/installer.go
//
//		WARNING 1
//			 !!! When build or run your application you must use the Go parameters 	"	gcflags='all=-N -l'  "		!!!
//		     !!! For exemple :		go run gcflags='all=-N -l' .
//			 !!!
//			 !!! Without that Go will throw your lib sayinh that it don't use the same ABI.
//			 !!! The reason is the "gcflags" options added when building the plugin.
//			 !!! They have been added in order to allows Go debugger to executing, without that the debug doesn't accept our plugin.
//
//		WARNING 2
//			 !!! The plugin must be rebuild if your alter the projects "progpjs.progpV8Engine" or "progpjs.progpAPI".  !!!
//			 !!! You will known when rebuild is required since Go will not accept your plugin anymore.

func testEnginePath(toTest string) string {
	cwd, _ := os.Getwd()
	toTest = path.Join(cwd, toTest)

	stat, err := os.Stat(toTest)

	if (err == nil) && stat.IsDir() {
		return toTest
	}

	return ""
}

func init() {
	// Allows to link the progpV8Engine inside the executable.
	// If you comment this line (or delete this file) then V8 will not be embedded.
	// The engine will then try to load hil from the plugin searched at ../plugins/progpV8.so
	// See script createPlugin.sh to create this plugin.
	//
	progpV8Engine.RegisterEngine()

	// Allows hacking the path to the source code of the project progpjs.progpV8Engine
	// which is used when generating code.
	//
	currentEnvValue := os.Getenv("PROGPV8_DIR")

	if currentEnvValue != "off" {
		if currentEnvValue == "" {
			// For the core dev team, where the real version is installed in the parent folder.
			foundPath := testEnginePath("../../progpjs.progpV8Engine")

			if foundPath == "" {
				// For those using the samples directly from the repo.
				foundPath = testEnginePath("../../progpjs.progpV8Engine")
			}

			if foundPath != "" {
				_ = os.Setenv("PROGPV8_DIR", foundPath)
			}
		}
	} else {
		os.Setenv("PROGPV8_DIR", "")
	}
}
