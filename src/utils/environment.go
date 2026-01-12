package utils

import "os"

func Environment() {
	os.Setenv("PORT", "3000")
}
