/*
Copyright © 2024 ani4learning@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"

	"github.com/anibmurthy/htmgenie/pkg/logger"
	"github.com/rs/xid"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
func RootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "htmgenie",
		Short: "Handles Markdown to HTML equivalent conversion and validations",
		Long:  `Handles Markdown to HTML equivalent conversion and validations`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(rootCmd *cobra.Command) error {
	RootCmdFlags(rootCmd)
	rootCmd.AddCommand(GenerateCommand())

	// create a correlation ID for the request
	correlationID := xid.New().String()

	ctx := context.WithValue(
		context.Background(),
		logger.CtxCorrelationKey,
		correlationID,
	)

	// retrieve the standard logger instance
	l := logger.Get()
	ctx = logger.WithCtx(ctx, l)

	rootCmd.SetContext(ctx)
	return rootCmd.Execute()

}

func RootCmdFlags(rootCmd *cobra.Command) {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.htmgenie.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func init() {

}
