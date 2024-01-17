package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Fatalf(message string, args ...interface{}) {
	if len(args) > 0 {
		_, _ = fmt.Fprintf(os.Stderr, message+"\n", args...)
	} else {
		_, _ = fmt.Fprintln(os.Stderr, message)
	}
	os.Exit(1)
}
func ExactArgs(cmd *cobra.Command, args []string, l int) {
	if len(args) < l {
		Fatalf(`%sExpected exactly %d command line arguments but got %d.`, cmd.UsageString(), l, len(args))
	}
}
