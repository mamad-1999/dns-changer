package utils

import (
	"log"

	"github.com/fatih/color"
)

// HandleError logs the error message and prints a user-friendly message
func HandleError(err error, msg string) {
	if err != nil {
		color.Red("%s: %s", msg, err)
		log.Println(err) // Log the error for debugging
		return
	}
}
