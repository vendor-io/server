package util

import (
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "keyboardify-server"

func LoadEnv() {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + "/.env")
	if err != nil {
		fmt.Println(err)
		panic("Failed loading .env file")
	}
}
