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

import "github.com/progpjs/progpjs/v2"

// ===  Plug & Play file  ===   Enable the use of javascript  ===
//
// - When enabled (file named "plugAndPlay_EnabledJavascript.go)
//		==> javascript is used and script are really executed.
// - When disable (file named "_plugAndPlay_EnabledJavascript.go)
//		==> javascript is disabled and script aren't executed.
//			The bootstrap step returns after enabled the minimal things.

func init() {
	gBootstrapHook = progpjs.Bootstrap
}
