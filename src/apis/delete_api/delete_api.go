package deletes_api

import (
	"net/http"
	"os"
	"path/filepath"
)

func DeleteFile(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodDelete {
		http.Error(response, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	err := request.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	fileName := request.MultipartForm.Value["file_name"][0]

	absPath, err := filepath.Abs("src/uploads/" + fileName)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	//является ли путь файлом
	fileInfo, err := os.Stat(absPath)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	//удаляем если не директория
	if !fileInfo.IsDir() {
		err = os.Remove(absPath)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(response, "Указанный путь не является файлом", http.StatusBadRequest)
	}
}
