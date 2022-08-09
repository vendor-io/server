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

	data, err := os.ReadFile(string(rootPath) + "/.env")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)
	os.Stdout.Close()

	errenv := godotenv.Load(string(rootPath) + "/.env")
	if errenv != nil {
		fmt.Println(errenv.Error())
		log.Fatal(errenv)
		panic("Failed loading .env file")
	}
}
