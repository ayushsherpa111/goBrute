/*
Copyright © 2020 NAME HERE ayush20100@hotmail.com

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
	"crypto"
	"fmt"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

// hashesCmd represents the hashes command
var hashesCmd = &cobra.Command{
	Use:   "hashes",
	Short: "Get All Available hashes and their corresponding code",
	Run: func(cmd *cobra.Command, args []string) {
		myFig := figure.NewColorFigure("GOBRUTE", "speed", "purple", true)
		myFig.Print()
		fmt.Println("* Specify one of the available format (-f) codes.")
		fmt.Println(getHashTable())
	},
}

func getHashTable() string {
	var tableStart = fmt.Sprintf("\n%3v%-3v|%3v%-12v|\n", "", "#", "", "Hash")
	tableStart += strings.Repeat("=", 6)
	tableStart += "+"
	tableStart += strings.Repeat("=", 15)
	tableStart += "+"
	for start := 1; start < 20; start++ {
		hashString := crypto.Hash(start)
		tableStart += fmt.Sprintf("\n%3v%-3v|%3v%-12v|", "", start, "", hashString)
	}

	return tableStart
}

func init() {
	rootCmd.AddCommand(hashesCmd)
}
