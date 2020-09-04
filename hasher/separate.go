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
		performHash = &HashStruct{0, make(map[string]string), md4.New()}
	case 2:
		fmt.Println("Cracking MD5 hash")
		performHash = &HashStruct{0, make(map[string]string), md5.New()}
	case 3:
		fmt.Println("Cracking SHA-1")
		performHash = &HashStruct{0, make(map[string]string), sha1.New()}
	case 4:
		fmt.Println("Cracking SHA-224")
		performHash = &HashStruct{0, make(map[string]string), sha256.New224()}
	case 5:
		fmt.Println("Cracking SHA-256")
		performHash = &HashStruct{0, make(map[string]string), sha256.New()}
	case 6:
		fmt.Println("Cracking SHA-384")
		performHash = &HashStruct{0, make(map[string]string), sha512.New384()}
	case 7:
		fmt.Println("Cracking SHA-512")
		performHash = &HashStruct{0, make(map[string]string), sha512.New()}
	case 8:
		fmt.Println("MD5+SHA1 unavailable")
		os.Exit(1)
	case 9:
		fmt.Println("Cracking RIPEMD-160")
		performHash = &HashStruct{0, make(map[string]string), ripemd160.New()}
	case 10:
		fmt.Println("Cracking SHA3-224")
		performHash = &HashStruct{0, make(map[string]string), sha3.New224()}
	case 11:
		fmt.Println("Cracking SHA3-256")
		performHash = &HashStruct{0, make(map[string]string), sha3.New256()}
	case 12:
		fmt.Println("Cracking SHA3-384")
		performHash = &HashStruct{0, make(map[string]string), sha3.New384()}
	case 13:
		fmt.Println("Cracking SHA3-512")
		performHash = &HashStruct{0, make(map[string]string), sha3.New512()}
	case 14:
		fmt.Println("Cracking SHA512-224")
		performHash = &HashStruct{0, make(map[string]string), sha512.New512_224()}
	case 15:
		fmt.Println("Cracking SHA512-256")
		performHash = &HashStruct{0, make(map[string]string), sha512.New512_256()}
	case 16:
		fmt.Println("Cracking SHA512-256")
		performHash = &HashStruct{0, make(map[string]string), sha512.New()}
	}
	hashList(wordlist, &performHash)
	return performHash
}
