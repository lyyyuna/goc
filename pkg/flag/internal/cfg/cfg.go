/*
 Copyright 2020 Qiniu Cloud (qiniu.com)

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package cfg

import (
	"github.com/qiniu/goc/pkg/flag/internal/build"
)

// These are general "build flags" used by build and other commands.
var (
	BuildA                 bool          // -a flag
	BuildBuildmode         string        // -buildmode flag
	BuildContext           build.Context //= defaultContext()
	BuildMod               string        // -mod flag
	BuildModReason         string        // reason -mod flag is set, if set by default
	BuildI                 bool          // -i flag
	BuildLinkshared        bool          // -linkshared flag
	BuildMSan              bool          // -msan flag
	BuildN                 bool          // -n flag
	BuildO                 string        // -o flag
	BuildP                 int           // = runtime.NumCPU() // -p flag
	BuildPkgdir            string        // -pkgdir flag
	BuildRace              bool          // -race flag
	BuildToolexec          []string      // -toolexec flag
	BuildToolchainName     string
	BuildToolchainCompiler func() string
	BuildToolchainLinker   func() string
	BuildTrimpath          bool // -trimpath flag
	BuildV                 bool // -v flag
	BuildWork              bool // -work flag
	BuildX                 bool // -x flag

	ModCacheRW bool   // -modcacherw flag
	ModFile    string // -modfile flag

	CmdName string // "build", "install", "list", "mod tidy", etc.

	DebugActiongraph string // -debug-actiongraph flag (undocumented, unstable)
)
