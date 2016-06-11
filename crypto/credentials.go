package credentials

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"syscall"
)

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
