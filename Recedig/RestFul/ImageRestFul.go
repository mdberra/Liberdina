package RestFul

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const maxUploadSize = 2097152 // 2 mb
type ImageRestFul struct {
}

func (rest *ImageRestFul) PostImage(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		renderError(w, "Archivo demasiado grande", http.StatusBadRequest)
		return
	}
	// Armo el nombre
	dni := r.FormValue("dni")
	tipoFile := r.FormValue("tipoFile")
	aux := "image" + tipoFile
	log.Println("PostImage: " + dni + " " + aux)

	file, handle, err := r.FormFile(aux)
	if err != nil {
		renderError(w, "Archivo no encontrado", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileName := dni + "-" + tipoFile

	mimeType := handle.Header.Get("Content-Type")
	switch mimeType {
	case "application/octet-stream":
		fileName += ".jpeg"
		break
	case "image/jpeg":
		fileName += ".jpeg"
		break
	case "image/jpg":
		fileName += ".jpg"
		break
	case "image/gif":
		fileName += ".gif"
		break
	case "image/png":
		fileName += ".png"
		break
	case "application/pdf":
		fileName += ".pdf"
		break
	default:
		renderError(w, "Tipo de archivo invalido", http.StatusBadRequest)
		return
	}
	if err := cloudStorageService.Connect(); err != nil {
		log.Fatalf("No se puede contectar: %v", err)
	}
	// Post
	if err := cloudStorageService.UploadFile(fileName, file); err != nil {
		log.Fatalf("No puede UploadObject : %v", err)
	}
	w.Header().Set("Content-Type", "application/json") // standard del protocolo http
	w.WriteHeader(http.StatusCreated)                  // escribe el 201 Created
}

func (rest *ImageRestFul) GetImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")
	msg := []byte("")

	vars := mux.Vars(r)
	fileName := vars["fileName"] + ".jpeg"

	if err := cloudStorageService.Connect(); err != nil {
		msg = []byte("ImageRestFul.GetImageDni - No se puede contectar: " + err.Error())
		w.WriteHeader(http.StatusNotFound)
		w.Write(msg)
	} else {

		image, err := cloudStorageService.Read(fileName)

		if err != nil {
			msg = []byte("No se encontro la imagen " + fileName + " " + err.Error())
			w.WriteHeader(http.StatusNotFound)
			w.Write(msg)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(image)
		}
	}
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json") // standard del protocolo http
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}
