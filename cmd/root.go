package cmd

import (
	"bufio"
	"crypto"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ayushsherpa111/goBrute/hasher"
	"github.com/spf13/cobra"
	// "github.com/cheggaaa/pb/v3"
)

type hashTarget struct {
	pFile   *bufio.Scanner
	pOsFile *os.File
	wFile   *bufio.Scanner
	wOsFile *os.File
}

const defaultWordlist = "/usr/share/wordlist/rockyou.txt"

var (
	rootCmd = &cobra.Command{
		Use:   "Brute",
		Short: "Brute force thru a hash | list of hashes",
		Long:  "Specify a hash or a list of hash in a file along with a wordlist to brute force",
		RunE: func(cmd *cobra.Command, args []string) error {
			if isStrEmpty(passwordFile) && isStrEmpty(passwordHash) || isStrEmpty(hashFormat) {
				errorMessage := "Error:"
				errorMessage += paramError("passwordFile", "P", "Missing Password file")
				errorMessage += "or"
				errorMessage += paramError("passwordHash", "p", "Missing Password Hash")
				errorMessage += paramError("hashFormat", "f", "Missing Hash Format. Use < goBrute hashes > to view all available hashes\n")
				fmt.Println(errorMessage)
				return nil
			}

			// check if the user has provided a wordlist to read from
			if !checkFile(wordlist) {
				// if not specified use the rock you wordlist which is the default path
				fmt.Println("Wordlist Not Specified. Using rockyou")
				wordlist = defaultWordlist
			}
			var wordErr error
			brute.wOsFile, wordErr = os.Open(wordlist)

			// If wordlist file doesnt exist throw an error
			if os.IsNotExist(wordErr) {
				return fmt.Errorf(paramError("wordlist", "w", "Couldn't locate wordlist file"))
			} else {
				// var delim = '\n'
				brute.wFile = bufio.NewScanner(brute.wOsFile)

				// check if the password File has been specified or the file exists
				if !isStrEmpty(passwordFile) {
					if checkFile(passwordFile) {
						// start doing stuff
						var passErr error
						brute.pOsFile, passErr = os.Open(passwordFile)
						if passErr == nil {
							brute.pFile = bufio.NewScanner(brute.pOsFile)
							startBrute(&brute)
						}
					} else {
						return fmt.Errorf(paramError("passwordFile", "P", "Couldn't locate passwordFile"))
					}
				} else if !isStrEmpty(passwordHash) {
					// password hash is defined
				} else {
					return fmt.Errorf("Missing arguments passwordFile (-P) or passwordHash (-p)")
				}
			}
			return nil
		},
	}
	brute        hashTarget
	wordlist     string
	passwordFile string
	passwordHash string
	hashFormat   string
)

func paramError(parameter, alias, message string) string {
	return fmt.Sprintf("\n * %s (-%s): %s", parameter, alias, message)
}

func init() {
	rootCmd.Flags().StringVarP(&wordlist, "wordlist", "w", wordlist, "Define a Wordlist for the application to hash.")
	rootCmd.Flags().StringVarP(&passwordFile, "passwordFile", "P", passwordFile, "Path to file with hashes.")
	rootCmd.Flags().StringVarP(&passwordHash, "passwordHash", "p", passwordHash, "Hash of the password to crack.")
	rootCmd.Flags().StringVarP(&hashFormat, "format", "f", hashFormat, "The format of the hash in question.")
	rootCmd.ParseFlags(os.Args[1:])
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(10)
	}
}

func checkFile(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func isStrEmpty(str string) bool {
	if len(str) > 0 {
		return false
	}
	return true
}

func printResults(results map[string]string) {

}

func startBrute(bruteSettings *hashTarget) {

	// create a map of the wordlist for comparing with the password list
	var parsedHashFormat, _ = strconv.ParseInt(hashFormat, 10, 8)
	var totalHash hasher.HashBrute = hasher.Distinguish(bruteSettings.wFile, crypto.Hash(parsedHashFormat))
	fmt.Println("number of lines ", totalHash.GetCount())
	// Start iterating over the map while comparing with the password list hash
	found := make(map[string]string)
	allFound := false
	for key, value := range totalHash.GetList() {
		bruteSettings.pFile.Split(bufio.ScanLines)
		for bruteSettings.pFile.Scan() {
			var passHash = bruteSettings.pFile.Text()
			var err = bruteSettings.pFile.Err()
			if err != nil {
				fmt.Println(err.Error())
				allFound = true
				break
			}

			if passHash == value {
				fmt.Printf("Found hash %v %v\n", key, value)
				found[key] = value
				break
			}
		}
		bruteSettings.pOsFile.Seek(0, os.SEEK_SET)
		bruteSettings.pFile = bufio.NewScanner(brute.pOsFile)
		if allFound {
			break
		}
	}
	if len(found) > 0 {
		fmt.Println("\nFound hashes:")
		fmt.Printf("%-4v* Format: %v\n", "", hashFormat)
		fmt.Printf("%5v%-15v|%15v%-19v|\n", "", "Plain Text", "", "Hash")
		fmt.Printf("%v+%v+\n", strings.Repeat("=", 20), strings.Repeat("=", 34))

		for hashKey, plainText := range found {
			fmt.Printf("%5v%-15v|%1v%-33v|\n", "", hashKey, "", plainText)
		}
	} else {
		fmt.Println("¯\\_(ツ)_/¯ no hashes found")
	}
}
