package util

import (
	"fmt"
	"io/ioutil"
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

	filesDir, errfile := ioutil.ReadDir(string(rootPath))
	if errfile != nil {
		log.Fatal(errfile)
	}

	for _, file := range filesDir {
		fmt.Println(file.Name(), file.IsDir())
	}

	files, errfiles := ioutil.ReadDir("/")
	if errfiles != nil {
		log.Fatal(errfiles)
	}

	for _, filex := range files {
		fmt.Println(filex.Name(), filex.IsDir())
	}

	fmt.Println(string(rootPath))

	errenv := godotenv.Load(string(rootPath) + "/.env")
	if errenv != nil {
		fmt.Println(errenv.Error())
		log.Fatal(errenv)
		panic("Failed loading .env file")
	}
}
