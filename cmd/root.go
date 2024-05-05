/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "diff",
	Short: "Check if two files are equal",
	RunE: func(cmd *cobra.Command, args []string) error {
		expected, err := readFiles(args[0])
		if err != nil {
			return err
		}

		actual, err := readFiles(args[1])
		if err != nil {
			return err
		}

		shortage := make([]string, 0)

		// deleteによってexpectedのインデックスがずれている？
		for _, v := range expected {
			if j := slices.Index(actual, v); j != -1 {
				actual = deleteItem(actual, j)
			} else {
				shortage = append(shortage, v)
			}
		}

		if len(shortage) == 0 && len(actual) == 0 {
			fmt.Println("Two files are equal")
		} else {
			fmt.Println("Two files are not equal")
			fmt.Printf("Shortage: %s\n", shortage)
			fmt.Printf("Extra: %s\n", actual)
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//rootCmd.Flags().StringP("expected", "e", "", "Expected file path")
	//rootCmd.Flags().StringP("actual", "a", "", "Actual file path")
}

func readFiles(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("%s : %v\n", filePath, err)
		return nil, err
	}

	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("%s : %v\n", filePath, err)
		return nil, err
	}

	return lines, nil
}

func deleteItem(list []string, index int) []string {
	list[index] = list[len(list)-1]
	return list[:len(list)-1]
}
