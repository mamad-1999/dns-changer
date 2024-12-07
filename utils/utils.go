package utils

import (
	"os"

	"github.com/fatih/color"
)

func HandleError(err error, msg string, fatal bool) {
	if err != nil {
		color.Red("%s: %s", msg, err)
		if fatal {
			os.Exit(1) // Exit the program on fatal errors
		}
	}
}
