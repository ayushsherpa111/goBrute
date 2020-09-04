package hasher

import (
	"bufio"
	"crypto"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"strings"

	"golang.org/x/crypto/md4"
)

func hashList(wordList *bufio.Scanner, hashSetting *HashBrute) {
	wordList.Split(bufio.ScanLines)
	for wordList.Scan() {
		(*hashSetting).IncrementCount()
		currLine := wordList.Text()
		currLine = strings.Trim(currLine, "\n ")
		(*hashSetting).SetKey(currLine, Brute(currLine, (*hashSetting).GetHash()))
		(*hashSetting).Reset()
		if len(currLine) == 0 {
			break
		}
	}
}

func Brute(target string, hasher hash.Hash) string {
	io.WriteString(hasher, target)
	// hasher.Write([]byte(target))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Distinguish(wordlist *bufio.Scanner, hashFormat crypto.Hash) HashBrute {
	var performHash HashBrute
	switch hashFormat {
	case 1:
		fmt.Println("Cracking MD4 hash")
		performHash = &MD4Brute{0, make(map[string]string), md4.New()}
	case 2:
		fmt.Println("Cracking MD5 hash")
		performHash = &MD5Brute{0, make(map[string]string), md5.New()}
		// hashedData = hashList(wordlist, BeginMd5Hash, &count)
	case 3:
		fmt.Println("Cracking SHA-1")
		performHash = &SHA1Brute{0, make(map[string]string), sha1.New()}
		// hashedData = hashList(wordlist, BeginSHA1Hash, &count)
	case 4:
		fmt.Println("Cracking SHA-224")
		performHash = &SHA224Brute{0, make(map[string]string), sha256.New224()}
	}
	hashList(wordlist, &performHash)
	return performHash
}
