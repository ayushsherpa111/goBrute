/*
Copyright © 2020 ayush sherpa

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Let goBrute generate the password of the given pattern/length and try an brute force the hash",
	Long: `Give goBrute a pattern of what the target password looks like with a set of flags.
	goBrute will then try and generate the sequence hash it and compare it to the supplied hash/hash file.
		*Format:
			?l : Lowercase Letter
			?u : Uppercase Letter
			?n : Number
			?s : Special Character
		Example:
			goBrute gen -f ?u?l?l?l?l?l?n 
			- generates a 7 character long password with 1 leading uppercase letter, 5 lowercase letters and 1 number 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen called")
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
