package hasher

import (
	"bufio"
	"crypto"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
	"hash"
	"io"
	"os"
	"strings"

	"golang.org/x/crypto/md4"
)

func hashList(wordList *bufio.Scanner, hashStore *map[string]string, hashSetting *HashStruct, counter int) int {
	// break the buffer up in to lines
	wordList.Split(bufio.ScanLines)
	// iterate over each line
	for wordList.Scan() {
		counter += 1
		// Trim whitespace or any trailing new lines
		currLine := wordList.Text()
		currLine = strings.Trim(currLine, "\n ")

		(*hashStore)[currLine] = Brute(currLine, (*hashSetting).GetHash())
		(*hashSetting).Reset()
		if len(currLine) == 0 {
			break
		}
	}
	return counter
}

func Brute(target string, hasher hash.Hash) string {
	io.WriteString(hasher, target)
	// hasher.Write([]byte(target))
	return hex.EncodeToString(hasher.Sum(nil))
}

func StartCracking(wordlist *bufio.Scanner, hashFormat crypto.Hash) (*map[string]string, int) {
	wdlst := make(map[string]string)
	hshFormat := Distinguish(hashFormat)
	count := hashList(wordlist, &wdlst, hshFormat, 0)
	return &wdlst, count
}

func Distinguish(hashFormat crypto.Hash) *HashStruct {
	var performHash *HashStruct
	switch hashFormat {
	case 1:
		fmt.Println("Cracking MD4 hash")
		performHash = &HashStruct{md4.New()}
	case 2:
		fmt.Println("Cracking MD5 hash")
		performHash = &HashStruct{md5.New()}
	case 3:
		fmt.Println("Cracking SHA-1")
		performHash = &HashStruct{sha1.New()}
	case 4:
		fmt.Println("Cracking SHA-224")
		performHash = &HashStruct{sha256.New224()}
	case 5:
		fmt.Println("Cracking SHA-256")
		performHash = &HashStruct{sha256.New()}
	case 6:
		fmt.Println("Cracking SHA-384")
		performHash = &HashStruct{sha512.New384()}
	case 7:
		fmt.Println("Cracking SHA-512")
		performHash = &HashStruct{sha512.New()}
	case 8:
		fmt.Println("MD5+SHA1 unavailable")
		os.Exit(1)
	case 9:
		fmt.Println("Cracking RIPEMD-160")
		performHash = &HashStruct{ripemd160.New()}
	case 10:
		fmt.Println("Cracking SHA3-224")
		performHash = &HashStruct{sha3.New224()}
	case 11:
		fmt.Println("Cracking SHA3-256")
		performHash = &HashStruct{sha3.New256()}
	case 12:
		fmt.Println("Cracking SHA3-384")
		performHash = &HashStruct{sha3.New384()}
	case 13:
		fmt.Println("Cracking SHA3-512")
		performHash = &HashStruct{sha3.New512()}
	case 14:
		fmt.Println("Cracking SHA512-224")
		performHash = &HashStruct{sha512.New512_224()}
	case 15:
		fmt.Println("Cracking SHA512-256")
		performHash = &HashStruct{sha512.New512_256()}
	case 16:
		fmt.Println("Cracking SHA512-256")
		performHash = &HashStruct{sha512.New()}
	}
	return performHash
}
