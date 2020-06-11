package flag

import (
	"flag"
	"os"
	"strings"

	"github.com/qiniu/goc/pkg/flag/internal/work"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// ExtractGocFlags parses all GOC flags, and left only go build/run/install flags/args
func ExtractGocFlags(cmd *cobra.Command, args []string) ([]string, []string) {
	var help bool
	gocFlags := make(map[string]*pflag.Flag)
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		gocFlags[f.Name] = f
	})

	newfset := flag.NewFlagSet("goc", flag.ContinueOnError)
	for k, _ := range gocFlags {
		if k == "help" {
			// delete(gocFlags, k)
			continue
		}
		newfset.Var(gocFlags[k].Value, k, "")
	}
	newfset.BoolVar(&help, "help", false, "")
	newfset.BoolVar(&help, "h", false, "")
	work.AddBuildFlags(newfset, work.DefaultBuildFlags)

	//newfset.ParseErrorsWhitelist = pflag.ParseErrorsWhitelist{UnknownFlags: true}
	newfset.Usage = func() { cmd.Help() }
	err := newfset.Parse(args)
	if err != nil {
		os.Exit(0)
	}
	if help {
		cmd.Help()
		os.Exit(0)
	}

	goArgs := make([]string, len(newfset.Args()))
	copy(goArgs, newfset.Args())
	// fmt.Println(goArgs)
	argsWithoutGoc := removeGocFlags(gocFlags, args[:len(args)-len(goArgs)])
	argsWithoutGoc = append(argsWithoutGoc, goArgs...)
	// fmt.Println(argsWithoutGoc)

	return argsWithoutGoc, goArgs
}

func removeGocFlags(gocFlags map[string]*pflag.Flag, args []string) []string {
	newArgs := make([]string, 0)
	for i := len(args) - 1; i >= 0; i-- {
		for k, v := range gocFlags {
			// if there is goc flags in the arguments, "-xxx" like
			if strings.Contains(args[i], "-"+k) {
				switch v.Value.Type() {
				case "bool":
					args[i] = ""
				default:
					if strings.Contains(args[i], "-"+k+"=") {
						args[i] = ""
					} else {
						args[i] = ""
						if i+1 < len(args) {
							args[i+1] = ""
						}
					}
				}
			}
		}
	}

	for i := 0; i < len(args); i++ {
		if args[i] != "" {
			newArgs = append(newArgs, args[i])
		}
	}

	return newArgs
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
