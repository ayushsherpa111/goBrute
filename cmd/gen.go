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
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/ayushsherpa111/goBrute/hasher"
	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var (
	genCmd = &cobra.Command{
		Use:   "gen",
		Short: "goBrute generates the password of the given pattern/length and try an brute force the hash",
		Long: `Give goBrute a pattern of what the target password looks like with a set of flags.
goBrute will then try and generate the sequence hash it and compare it to the supplied hash/hash file.
	*Format:
		?l : Lowercase Letter
		?u : Uppercase Letter
		?n : Number
		?s : Special Character
	Example:
		goBrute gen -p ?u?l?l?l?l?l?n 
		- generates a 7 character long password with 1 leading uppercase letter, 5 lowercase letters and 1 number 
	`,
		RunE: func(cmd *cobra.Command, args []string) error {

			// The pattern for the password is missing
			if len(pattern) == 0 {
				patternErr := paramError("pattern", "p", "Missing password pattern")
				patternErr += `
 * Format:
	?l : Lowercase Letter
	?u : Uppercase Letter
	?n : Number
	?s : Special Character
				`
				return fmt.Errorf(patternErr)
			}

			// check if the user has supplied a target hash to brute force
			if len(hshTarget) == 0 {
				return fmt.Errorf(paramError("hash", "j", "Target Hash is missing"))
			}

			// split the pattern into individual components for easier processing
			splts := strings.Split(pattern, "?")[1:]
			if !validatePattern(splts) {
				return fmt.Errorf("* Unknown pattern supplied")
			}
			startGen(len(splts), "", splts, splts[0], false)
			return nil
		},
	}

	hshTarget    string
	pattern      string
	targetFormat string
	// printable characters for the passsword
	printableNumbers string = "0123456789"
	printableLower   string = "abcdefghijklmnopqrstuvwxyz"
	printableUpper   string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	printableSpecial string = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ 	"
)

func getCharLen(flag string) (int, string) {
	switch flag {
	case "l":
		return len(printableLower), printableLower
	case "u":
		return len(printableUpper), printableUpper
	case "n":
		return len(printableNumbers), printableNumbers
	case "s":
		return len(printableSpecial), printableSpecial
	default:
		return 0, ""
	}
}

func startGen(chars int, start string, pattern []string, nextPr string, found bool) {
	// A recursive function that generates every possible password combination
	// the user supplies a `pattern`
	// the function goes through each pattern format

	// once the total required characters have been met return
	if chars == 0 || found {

		// start hashing and comparing
		fndHash := hasher.Brute(start, md5.New())
		if fndHash == hshTarget {
			fmt.Println("Found password", start, fndHash)
			found = true
		}
		return
	}

	// get the character set of the current pattern
	size, data := getCharLen(nextPr)

	// loop over each character in the character set
	for counter := 0; counter < size; counter++ {
		newPrefix := start + string(data[counter])
		var nextPtrn string
		// find the next pattern in the list. if its the last pattern ignore the data
		if len(start) < len(pattern)-1 {
			nextPtrn = pattern[len(start)+1]
		}
		// start iterating over the next pattern and reduce the required number of characters by 1
		startGen(chars-1, newPrefix, pattern, nextPtrn, found)
	}

}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.Flags().StringVarP(&hshTarget, "hash", "j", pattern, "The Target hash to brute force")
	genCmd.Flags().StringVarP(&pattern, "pattern", "p", pattern, "Supply a pattern of the target password.")
	genCmd.Flags().StringVarP(&targetFormat, "target", "t", targetFormat, "Target Hash Algorithim to use.")
}

func validatePattern(splts []string) bool {
	for _, v := range splts {
		if v != "l" && v != "u" && v != "n" && v != "s" {
			return false
		}
	}
	return true
}
