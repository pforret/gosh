/*
Copyright © 2022 Peter Forret <peter@forret.com>

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
    "time"
    "flag"

	"github.com/spf13/cobra"
)

// microtimeCmd represents the microtime command
var microtimeCmd = &cobra.Command{
	Use:   "microtime",
	Short: "return time in 'epoch' style (seconds since 1 Jan 1970)",
	Long: `returns time as an integer or as a float
> microtime
1672080617
> microtime true
1672080617.458
`,
    Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        loc, _ := time.LoadLocation("UTC")
        now := time.Now().In(loc)
        if len(flag.Args()) <= 1 {
            // microtime -- without arguments
            fmt.Printf("%d\n",now.Unix())
        }
        if len(flag.Args()) >= 2 {
            // microtime true -or- microtime 1
            microSeconds := float64(now.Nanosecond()) / 1000000000
            fmt.Printf("%.6f\n",float64(now.Unix()) + microSeconds)
        }
	},
}

func init() {
	rootCmd.AddCommand(microtimeCmd)
}
