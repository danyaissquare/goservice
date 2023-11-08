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

	err := request.ParseForm()
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	fileName := request.FormValue("file_name")

	filePath := filepath.Join("src/uploads", fileName)

	err = os.Remove(filePath)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusNoContent)
}
