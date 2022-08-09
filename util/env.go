package util

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "keyboardify-server"

func LoadEnv() {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	errenv := godotenv.Load(string(rootPath) + "/.env")
	if errenv != nil {
		fmt.Println(errenv.Error())
		log.Fatal(errenv)
		panic("Failed loading .env file")
	}
}
