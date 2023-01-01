/*
Copyright © 2023 Peter Forret <peter@forret.com>

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
	"bufio"
	"os"
    "regexp"
    "strings"
	"github.com/anyascii/go"
	"github.com/spf13/cobra"
)

// slugCmd represents the slug command
var slugCmd = &cobra.Command{
	Use:   "slug",
	Short: "Convert Unicode text into a slug",
	Long: `Convert Unicode text into a slug
> echo 'Îñtérnåtîönàlïsātìõń is $(AWESOME)!!' | gosh slug
internationalisation-is-awesome`,
	Run: func(cmd *cobra.Command, args []string) {
    scanner := bufio.NewScanner((os.Stdin))
    badchars := regexp.MustCompile(`[^\w]+`)
    whitespace := regexp.MustCompile(`\s+`)
    for {
        scanner.Scan()
        orig := scanner.Text()
        roman := anyascii.Transliterate(orig)
        clean := strings.Trim(badchars.ReplaceAllString(roman," ")," ")
        fmt.Println(strings.ToLower(whitespace.ReplaceAllString(clean,"-")))
        break
        }
	},
}

func init() {
	rootCmd.AddCommand(slugCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// slugCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// slugCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
