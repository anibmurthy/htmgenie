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
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/anibmurthy/htmgenie/pkg/logger"
	"github.com/anibmurthy/htmgenie/pkg/parser"
	"github.com/anibmurthy/htmgenie/pkg/util"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var fileArg string

func Validate(cmd *cobra.Command, args []string) error {
	// Validate that the passed input file path is valid
	ctx := cmd.Context()
	l := logger.FromCtx(ctx)
	var err error

	if fileArg == "" {
		l.Error("Required argument 'file' not found")
		err = fmt.Errorf("file is a required field.")
	} else if !strings.HasSuffix(fileArg, ".md") {
		l.Error("Only files with '.md' extensions are supported")
		err = fmt.Errorf("Only files with '.md' extensions are supported")
	}

	if info, err := os.Stat(fileArg); err == nil {
		// Check the size of the file.
		// TODO: This can be made configurable
		if size := info.Size(); size > 5242880 {
			err = fmt.Errorf("File size is more than the allowed limit (5MB): %v", size/1048576)
		}
	} else {

		// Check that the file exists
		l.Error("File not found in specified path")
	}
	return err
}

func GenerateCmdRun(cmd *cobra.Command, args []string) {
	start := time.Now()
	log := logger.FromCtx(cmd.Context())
	log.Sugar().Infof("Generating html for the supplied '%s' markdown file", fileArg)

	file, err := os.Open(fileArg)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not read the file in path: %s", fileArg),
			zap.String("Error:", err.Error()))
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(fmt.Sprintf("Error closing the markdown file after read at: %s", fileArg),
				zap.String("Error:", err.Error()))
		}
	}()

	// Output file is created in the path set through environment variable 'HTMGENIE_OPATH'
	opath := util.GetOutPath(fileArg)
	ofile, err := os.Create(opath)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not scan the loaded file from path: %s", opath),
			zap.String("Error:", err.Error()))
	}

	defer func() {
		if err := ofile.Close(); err != nil {
			log.Fatal(fmt.Sprintf("Error closing the markdown file after read at: %s", opath),
				zap.String("Error:", err.Error()))
		}
	}()

	parser := parser.New(file, ofile, log)
	parser.Generate()
	log.Sugar().Info("Markdown generation took: ", time.Since(start))
}

// generateCmd represents the generate command
func GenerateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate HTML equivalent of a Markdown file.",
		Long: `Converts a Markdown file to its HTML equivalent and saves the result to a new file. 
	This tool supports a subset of Markdown syntax and generates HTML tags accordingly.
	Example:
			htmgenie generate -f <filename_with_path>

			`,
		PreRunE: Validate,
		Run:     GenerateCmdRun,
	}

	GenerateCommandFlags(cmd)
	return cmd
}

func GenerateCommandFlags(generateCmd *cobra.Command) {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")
	generateCmd.PersistentFlags().StringVarP(&fileArg, "file", "f", "", "Input markdown file to be converted to html. Accepts only '.md' files")
	_ = generateCmd.MarkFlagRequired("file")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
