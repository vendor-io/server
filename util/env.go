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

	fmt.Println(string(rootPath))

	err := godotenv.Load(string(rootPath) + "/.env")
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
		panic("Failed loading .env file")
	}
}
