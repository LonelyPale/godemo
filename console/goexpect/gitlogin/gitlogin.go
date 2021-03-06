package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/LonelyPale/goutils/crypto/aes"
	"github.com/google/goexpect"
	"github.com/google/goterm/term"
)

const (
	DefaultTimeout         = 10 * time.Second
	DefaultDownloadTimeout = 1 * time.Hour
)

var (
	userRE   = regexp.MustCompile("Username for")
	passRE   = regexp.MustCompile("Password for")
	promptRE = regexp.MustCompile("%")
)

// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gitlogin gitlogin.go
// 只支持Linux
func main() {
	if len(os.Args) == 1 {
		Print("gitlogin is a tool for git login.")
		Print("usage: gitlogin git-command")
		Print("example: ./gitlogin git -C ~/git/bsnportal/ pull")
		return
	}

	cmd := strings.Join(os.Args[1:], " ")

	user := "lonelypale@126.com"
	pass := ""

	key := []byte{0}
	ciphertext := []byte{0}

	bs, err := aes.Decrypt(ciphertext, key)
	if err != nil {
		Fatal(err)
	}
	pass = string(bs)

	//sh, _, err := expect.Spawn(cmd, -1, expect.Verbose(true))
	sh, res, err := expect.Spawn(cmd, -1)
	if err != nil {
		Fatal(err)
	}

	defer func() {
		if err := sh.Close(); err != nil {
			if !IsNotRunningError(err) {
				Fatal(err)
			}
		}
	}()

	resuser, _, err := sh.Expect(userRE, DefaultTimeout)
	if err != nil {
		Fatal(err)
	}
	Print(resuser)
	if err := sh.Send(user + "\n"); err != nil {
		Fatal(err)
	}

	respass, _, err := sh.Expect(passRE, DefaultTimeout)
	if err != nil {
		Fatal(err)
	}
	Print(respass)
	if err := sh.Send(pass + "\n"); err != nil {
		Fatal(err)
	}

	for {
		select {
		case <-res:
			return
		default:
			result, _, err := sh.Expect(promptRE, DefaultDownloadTimeout)
			if err != nil {
				if !IsNotRunningError(err) {
					Fatal(err)
				}
			}
			Success(result)
		}
	}

}

func Print(s string) {
	if len(s) > 0 {
		fmt.Println(term.Blue(s))
	}
}

func Fatal(err error) {
	fmt.Println(term.Red(err.Error()))
	os.Exit(-1)
}

func Success(s string) {
	if len(s) > 0 {
		fmt.Println(term.Green(s))
	}
}

func IsNotRunningError(err error) bool {
	switch err.Error() {
	case "expect: Process not running":
		return true
	case "os: process already finished":
		return true
	default:
		return false
	}
}

func ReadPIDFile(filepath string) (int, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return -1, err
	}

	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return -1, err
	}

	return pid, nil
}
