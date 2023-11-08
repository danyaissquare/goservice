package upload_api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

//http://localhost:5000/api/upload

func UploadFile(response http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 * 1024 * 1024)
	files := request.MultipartForm.File["myfiles"]
	for _, file := range files {
		fmt.Println("file info")
		fmt.Println("File name: ", file.Filename)
		fmt.Println("File size: ", file.Size, "bytes")
		fmt.Println("File type: ", file.Header.Get("Content-Type"))
		fmt.Println("-------------------------------")
		f, _ := file.Open()

		extension := filepath.Ext(file.Filename) // Получить расширение файла
		tempFile, err := ioutil.TempFile("src/uploads", "upload-*"+extension)
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()

		fileBytes, err2 := ioutil.ReadAll(f)
		if err2 != nil {
			fmt.Println(err2)
		}
		tempFile.Write(fileBytes)
	}
	fmt.Println("done")

	if request.Method == http.MethodPost {

	}
}
