package cmd

import (
	"bufio"
	"crypto"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ayushsherpa111/goBrute/hasher"
	"github.com/cheggaaa/pb/v3"
	"github.com/spf13/cobra"
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

func printResults(found map[string]string) {
	fmt.Println("\nFound hashes:")
	var tableLen int
	var indexLen int
	for i, v := range found {
		tableLen = len(v)
		indexLen = len(i)
		break
	}

	if indexLen < 10 {
		indexLen = 11
	}
	fmt.Printf("%-4v* Format: %v\n", "", hashFormat)
	fmt.Printf("%*v%-*v|%*v%-*v|\n", (indexLen/2)-2, "", indexLen+1, "Plain Text", (tableLen/2)-3, "", (tableLen/2)+5, "Hash")
	fmt.Printf("%v+%v+\n", strings.Repeat("=", indexLen+4), strings.Repeat("=", tableLen+2))

	for hashKey, plainText := range found {
		fmt.Printf("%*v%-*v|%1v%-*v|\n", (indexLen/2)-2, "", indexLen+1, hashKey, "", tableLen+1, plainText)
	}

}

func startBrute(bruteSettings *hashTarget) {
	// create a map of the wordlist for comparing with the password list
	var parsedHashFormat, _ = strconv.ParseInt(hashFormat, 10, 8)
	var totalHash hasher.HashBrute = hasher.Distinguish(bruteSettings.wFile, crypto.Hash(parsedHashFormat))
	// Start iterating over the map while comparing with the password list hash
	found := make(map[string]string)
	allFound := false

	// keep track of wheather all the passwords have been cracked
	firstPass := true
	var totalPassList int

	// progress bar
	bar := pb.StartNew(totalHash.GetCount())
	bar.SetRefreshRate(time.Millisecond)
	bar.Set(pb.Bytes, true)
	for key, value := range totalHash.GetList() {
		bruteSettings.pFile.Split(bufio.ScanLines)
		if allFound {
			break
		}
		bar.Increment()
		for bruteSettings.pFile.Scan() {

			if firstPass {
				totalPassList += 1
			}

			if !firstPass && len(found) == totalPassList {
				allFound = true
			}

			if found[key] != "" {
				continue
			}

			var passHash = bruteSettings.pFile.Text()
			var err = bruteSettings.pFile.Err()
			if err != nil {
				fmt.Println(err.Error())
				allFound = true
				break
			}

			if passHash == value {
				// fmt.Printf("Found hash %v %v\n", key, value)
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
		printResults(found)
	} else {
		fmt.Println("¯\\_(ツ)_/¯ no hashes found")
	}
	bar.Finish()
}
