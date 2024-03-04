package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/term"

	"github.com/kubeclipper/kubeclipper/pkg/cli/logger"
)

func AskForConfirmation() bool {
	var response string

	_, err := fmt.Scanln(&response)
	if err != nil {
		logger.Fatal(err)
	}

	switch strings.ToLower(response) {
	case "y", "yes":
		return true
	case "n", "no":
		return false
	default:
		fmt.Println("I'm sorry but I didn't get what you meant, please type (y)es or (n)o and then press enter:")
		return AskForConfirmation()
	}
}

func WaitInputPasswd() (string, error) {
	terminal := term.IsTerminal(int(os.Stdin.Fd()))
	if !terminal {
		return "", errors.New("operation not supported by device")
	}
	pBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	return string(pBytes), nil
}
