package flag

import (
	"flag"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// ExtractGocFlags parses all GOC flags, and left only go build/run/install flags/args
func ExtractGocFlags(cmd *cobra.Command, args []string) {
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
	// newfset.ParseErrorsWhitelist = pflag.ParseErrorsWhitelist{UnknownFlags: true}
	newfset.Parse(args)

	argsWithoutGoc := args[:len(newfset.Args())]
	removeGocFlags(gocFlags, argsWithoutGoc)
	argsWithoutGoc = append(argsWithoutGoc, newfset.Args()...)
}

func removeGocFlags(gocFlags map[string]*pflag.Flag, args []string) {
	for k, v := range gocFlags {
		// fmt.Println(k, v.Value, v.Value.Type())
		for i := len(args) - 1; i >= 0; i-- {
			// if there is goc flags in the arguments, "-xxx" like
			if strings.Contains(args[i], "-"+k) {
				switch v.Value.Type() {
				case "bool":
					args = append(args[:i], args[i+1:]...)
				default:
					if strings.Contains(args[i], "-"+k+"=") {
						args = append(args[:i], args[i+1:]...)
					} else {
						args = append(args[:i], args[i+2:]...)
					}
				}
			}
		}
	}
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
