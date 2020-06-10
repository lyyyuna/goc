package flag

import (
	"flag"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// ExtractGocFlags parses all GOC flags, and left only go build/run/install flags/args
func ExtractGocFlags(cmd *cobra.Command, args []string) []string {
	gocFlags := make(map[string]*pflag.Flag)
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		gocFlags[f.Name] = f
	})

	newfset := flag.NewFlagSet("goc", flag.ExitOnError)
	for k, _ := range gocFlags {
		if k == "help" {
			delete(gocFlags, k)
			continue
		}
		newfset.Var(gocFlags[k].Value, k, "")
	}
	//newfset.ParseErrorsWhitelist = pflag.ParseErrorsWhitelist{UnknownFlags: true}
	newfset.Parse(args)

	goArgs := make([]string, len(newfset.Args()))
	copy(goArgs, newfset.Args())
	argsWithoutGoc := removeGocFlags(gocFlags, args[:len(args)-len(goArgs)])
	argsWithoutGoc = append(argsWithoutGoc, goArgs...)
	//fmt.Println(argsWithoutGoc)

	return argsWithoutGoc
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
