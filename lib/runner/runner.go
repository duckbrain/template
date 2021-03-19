package runner

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func Main(c *cobra.Command) {
	err := c.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Exec(e string, args ...string) func(*cobra.Command, []string) error {
	return func(*cobra.Command, []string) error {
		c := exec.CommandContext(context.Background(), e, args...)
		c.Stdin = os.Stdout
		c.Stdin = os.Stdin
		c.Stderr = os.Stderr
		return c.Run()
	}
}

func CmdScript(e string, args ...string) func(*cobra.Command, []string) error {
	return Exec("go", append([]string{"run", fmt.Sprintf("./cmd/%v", e)}, args...)...)
	// return Exec("go", "run", fmt.Sprintf("./cmd/%v", e))
}
