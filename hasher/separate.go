package hasher

import (
	"bufio"
	"crypto"
	"fmt"
	"strings"
)

func hashList(wordList *bufio.Scanner, hashPass func(string) string, count *int) map[string]string {
	var passStore = make(map[string]string)
	wordList.Split(bufio.ScanLines)
	for wordList.Scan() {
		*count += 1
		currLine := wordList.Text()
		currLine = strings.Trim(currLine, "\n ")
		passStore[currLine] = hashPass(currLine)
		if len(currLine) == 0 {
			break
		}
	}
	return passStore
}

func Distinguish(wordlist *bufio.Scanner, hashFormat crypto.Hash) (map[string]string, int) {
	var hashedData = make(map[string]string)
	var count int
	switch hashFormat {
	case 1:
		fmt.Println("Cracking MD4 hash")
		hashedData = hashList(wordlist, BeginMd4Hash, &count)
	case 2:
		fmt.Println("Cracking MD5 hash")
		hashedData = hashList(wordlist, BeginMd5Hash, &count)
	case 3:
		fmt.Println("Cracking SHA-1")
		hashedData = hashList(wordlist, BeginSHA1Hash, &count)
	}
	return hashedData, count
}
