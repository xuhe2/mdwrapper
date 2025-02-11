package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/cobra"
	"github.com/xuhe2/mdwrapper/core"
)

var rootCmd = &cobra.Command{
	Use:   "mdwrapper",
	Short: "a wrapper for markdown files",
	Long:  `a wrapper for markdown files, it can wrap all resources in markdown file`,
	Run: func(cmd *cobra.Command, args []string) {
		// miss args
		if len(args) == 0 {
			fmt.Println("Please provide a markdown file")
			os.Exit(1)
		}
		// get zip file name
		zipFileName, err := cmd.Flags().GetString("output")
		if err != nil || zipFileName == "" {
			// get zip file name from first arg
			firstFileName := filepath.Base(args[0])
			zipFileName = strings.Split(firstFileName, ".")[0] + ".zip"
		}
		// create a file for wrapper
		file, err := os.Create(zipFileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()
		// wrap markdown file
		wrapper := core.NewWrapper().WithArchive(core.NewZipArchive(file))
		defer wrapper.Close()

		wg := &sync.WaitGroup{}
		wg.Add(len(args))
		for _, path := range args {
			go func() {
				defer wg.Done()
				// open the markdown file
				mdFile := core.NewMarkdownFile()
				if err := mdFile.Open(path); err != nil {
					fmt.Println(err)
					return
				}
				// TODO: add logic to wrap the markdown file
				if err := wrapper.Wrap(mdFile); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}()
		}
		wg.Wait()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("output", "o", "", "output zip file name")
}
