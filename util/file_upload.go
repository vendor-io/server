package util

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	uuid "github.com/nu7hatch/gouuid"
)

func MultipleFileUpload(c echo.Context, formFiles string, dir string) []string {
	form, err := c.MultipartForm()
	if err != nil {
		panic(err)
	}
	files := form.File[formFiles]
	var filesUrls []string

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			panic(err)
		}
		defer src.Close()

		id, err := uuid.NewV4()
		filename := fmt.Sprintf("%s-%s", id, file.Filename)

		dst, err := os.Create(filepath.Join("public", dir, filename))
		if err != nil {
			panic(err)
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			panic(err)
		}

		filesUrls = append(filesUrls, fmt.Sprintf("%s/api/public/%s/%s", os.Getenv("URL"), dir, filename))
	}

	return filesUrls
}
