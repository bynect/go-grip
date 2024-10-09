package cmd

import (
	"os"

	"github.com/chrishrb/go-grip/pkg"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-grip [file]",
	Short: "Render markdown document as html",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		dark, _ := cmd.Flags().GetBool("dark")
		browser, _ := cmd.Flags().GetBool("browser")
		port, _ := cmd.Flags().GetInt("port")

		client := pkg.Client{Dark: dark, OpenBrowser: browser, Port: port}

		err := client.Serve(args[0])
		cobra.CheckErr(err)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("dark", "d", false, "Darkmode")
	rootCmd.Flags().BoolP("browser", "b", true, "Open new browser tab")
	rootCmd.Flags().IntP("port", "p", 6419, "Port to use")
}
