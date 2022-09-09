/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

// stupidCmd represents the stupid command
var stupidCmd = &cobra.Command{
	Use:   "stupid [<number>]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		str := "田中さん、お元気ですか。今度の火曜日に鈴木さんと打ち合わせがあるのですがご都合いかがでしょうか。"
		rep := regexp.MustCompile("田中|鈴木")
		fmt.Println(rep.String())
		s := rep.ReplaceAllString(str, "XX")
		fmt.Println(s)
		// s := args[0]
		// fmt.Println(strings.ReplaceAll(s, "/3/", "3!"))
	},
}

func Hoge(num string) string {
	fmt.Println(num)
	return num
}

func init() {
	rootCmd.AddCommand(stupidCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stupidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stupidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
