package hasher

import (
	"bufio"
	"crypto"
	"fmt"
	"strings"
)

func hashList(wordList *bufio.Scanner, hashSetting HashBrute) {
	wordList.Split(bufio.ScanLines)
	for wordList.Scan() {
		hashSetting.IncrementCount()
		currLine := wordList.Text()
		currLine = strings.Trim(currLine, "\n ")
		hashSetting.SetKey(currLine, hashSetting.BeginHash(currLine))
		if len(currLine) == 0 {
			break
		}
	}
}

func create(hashFormat crypto.Hash) {

}

func Distinguish(wordlist *bufio.Scanner, hashFormat crypto.Hash) HashBrute {
	var performHash HashBrute
	switch hashFormat {
	case 1:
		fmt.Println("Cracking MD4 hash")
		performHash = &MD4Brute{0, make(map[string]string)}
	case 2:
		fmt.Println("Cracking MD5 hash")
		performHash = &MD5Brute{0, make(map[string]string)}
		// hashedData = hashList(wordlist, BeginMd5Hash, &count)
	case 3:
		fmt.Println("Cracking SHA-1")
		// hashedData = hashList(wordlist, BeginSHA1Hash, &count)
	}
	hashList(wordlist, performHash)
	return performHash
}
