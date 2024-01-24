/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "passgen",
}

// genpassCmd represents the genpass command
var genpassCmd = &cobra.Command{
	Use:   "genpass",
	Short: "A brief description of your command",
	//	Long: `A longer description that spans multiple lines and likely contains examples
	//and usage of using your command. For example:
	//
	//Cobra is a CLI library for Go that empowers applications.
	//This application is a tool to generate the needed files
	//to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		count, _ := cmd.Flags().GetInt("count")
		length, _ := cmd.Flags().GetInt("length")
		letters, _ := cmd.Flags().GetBool("letters")
		specials, _ := cmd.Flags().GetBool("specials")

		config := Config{
			Letters:  letters,
			Specials: specials,
		}

		pwds, err := Many(count, length, config)
		if err != nil {
			fmt.Println("Error generating a password:", err)
			return
		}

		if err := pwds.WriteFile("output.txt"); err != nil {
			fmt.Println("Error dumping passwords to file:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(genpassCmd)
	genpassCmd.PersistentFlags().IntP("count", "c", 0, "Count")
	genpassCmd.PersistentFlags().IntP("length", "n", 0, "Length")
	genpassCmd.PersistentFlags().BoolP("letters", "l", false, "Letters")
	genpassCmd.PersistentFlags().BoolP("specials", "s", false, "Specials")
}
