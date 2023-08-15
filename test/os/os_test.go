package os

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/bgentry/speakeasy"
)

func TestName(t *testing.T) {
	cmdFunc("pwd")
}

func cmdFunc(cmdValue string, args ...string) {
	cmd := exec.Command(cmdValue, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

func TestFork(t *testing.T) {
	Fork()
}

func Fork() {
	password, err := speakeasy.Ask("Please enter a password: ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Password result: %q\n", password)
	fmt.Printf("Password len: %d\n", len(password))
}
