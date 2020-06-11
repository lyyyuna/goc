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

package work

import (
	"flag"

	"github.com/qiniu/goc/pkg/flag/internal/cfg"
	"github.com/qiniu/goc/pkg/flag/internal/load"
)

// buildCompiler implements flag.Var.
// we only need buildCompiler to consume the flag
// don't care about the value, so just implement empty method
type buildCompiler struct{}

func (c buildCompiler) Set(value string) error {
	return nil
}

func (c buildCompiler) String() string {
	return ""
}

type BuildFlagMask int

const (
	DefaultBuildFlags BuildFlagMask = 0
	OmitModFlag       BuildFlagMask = 1 << iota
	OmitModCommonFlags
)

func AddBuildFlags(cmdset *flag.FlagSet, mask BuildFlagMask) {
	cmdset.BoolVar(&cfg.BuildA, "a", false, "")
	cmdset.BoolVar(&cfg.BuildN, "n", false, "")
	cmdset.IntVar(&cfg.BuildP, "p", cfg.BuildP, "")
	cmdset.BoolVar(&cfg.BuildV, "v", false, "")
	cmdset.BoolVar(&cfg.BuildX, "x", false, "")

	cmdset.Var(&load.BuildAsmflags, "asmflags", "")
	cmdset.Var(buildCompiler{}, "compiler", "")
	cmdset.StringVar(&cfg.BuildBuildmode, "buildmode", "default", "")
	cmdset.Var(&load.BuildGcflags, "gcflags", "")
	cmdset.Var(&load.BuildGccgoflags, "gccgoflags", "")
	if mask&OmitModFlag == 0 {
		cmdset.StringVar(&cfg.BuildMod, "mod", "", "")
	}
	if mask&OmitModCommonFlags == 0 {
		AddModCommonFlags(cmdset)
	}
	cmdset.StringVar(&cfg.BuildContext.InstallSuffix, "installsuffix", "", "")
	cmdset.Var(&load.BuildLdflags, "ldflags", "")
	cmdset.BoolVar(&cfg.BuildLinkshared, "linkshared", false, "")
	cmdset.StringVar(&cfg.BuildPkgdir, "pkgdir", "", "")
	cmdset.BoolVar(&cfg.BuildRace, "race", false, "")
	cmdset.BoolVar(&cfg.BuildMSan, "msan", false, "")
	cmdset.Var((*tagsFlag)(&cfg.BuildContext.BuildTags), "tags", "")
	cmdset.Var((*tagsFlag)(&cfg.BuildToolexec), "toolexec", "")
	cmdset.BoolVar(&cfg.BuildTrimpath, "trimpath", false, "")
	cmdset.BoolVar(&cfg.BuildWork, "work", false, "")

	// Undocumented, unstable debugging flags.
	cmdset.StringVar(&cfg.DebugActiongraph, "debug-actiongraph", "", "")

}

// AddModCommonFlags adds the module-related flags common to build commands
// and 'go mod' subcommands.
func AddModCommonFlags(cmdset *flag.FlagSet) {
	cmdset.BoolVar(&cfg.ModCacheRW, "modcacherw", false, "")
	cmdset.StringVar(&cfg.ModFile, "modfile", "", "")
}

// tagsFlag is the implementation of the -tags flag.
type tagsFlag []string

func (v *tagsFlag) Set(s string) error {
	return nil
}

func (v *tagsFlag) String() string {
	return "<TagsFlag>"
}
