package cryptoutils

import (
	"bufio"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"hash"
	"os"
	"syscall"
)

// Returns the base64-encoded sha-512 hash of the given string.
func Sha512(hashMe string) (string, error) {
	hasher := sha512.New()
	hasher.Write([]byte(hashMe))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return sha, nil
}

func hashSha512(hashMe string) (string, error) {
	return hashImpl(nil, base64.URLEncoding, hashMe)
}

func hashImpl(hash hash.Hash, encoding *base64.Encoding, hashMe string) (string, error) {
	hasher := hash
	if hasher == nil {
		// default to sha512
		hasher = sha512.New()
	}

	_, err := hasher.Write([]byte(hashMe))
	if err != nil {
		return "", errors.New("Couldn't write to byte-array")
	}

	return encoding.EncodeToString(hasher.Sum(nil)), nil
}

// Returns the user-supplied username, CLEARTEXT password, and any error
// that may have occurred.
func Credentials() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	return username, password
}
