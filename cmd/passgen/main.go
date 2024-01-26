package main

import (
	"fmt"
	"os"

	"github.com/Apolisk/passgen"
	"github.com/spf13/cobra"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

var cmd = &cobra.Command{
	Use: "passgen",
	Run: runCmd,
}

func init() {
	cmd.PersistentFlags().IntP("count", "c", 1, "count of passwords to generate")
	cmd.PersistentFlags().IntP("length", "n", 10, "length of the each password")
	cmd.PersistentFlags().BoolP("letters", "l", false, "includes letters")
	cmd.PersistentFlags().BoolP("specials", "s", false, "includes specials")
	// TODO: Add "output" file flag.

	cmd.CompletionOptions.DisableDefaultCmd = true
}

func runCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		return
	}

	count, _ := cmd.Flags().GetInt("count")
	length, _ := cmd.Flags().GetInt("length")
	letters, _ := cmd.Flags().GetBool("letters")
	specials, _ := cmd.Flags().GetBool("specials")

	config := passgen.Config{
		Letters:  letters,
		Specials: specials,
	}

	pwds, err := passgen.Many(count, length, config)
	if err != nil {
		fmt.Println("Error generating a password:", err)
		return
	}

	if err := pwds.WriteFile("output.txt"); err != nil {
		fmt.Println("Error dumping passwords to file:", err)
	}
}
