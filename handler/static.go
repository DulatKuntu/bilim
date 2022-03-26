package handler

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// RouterImageHandler helps to handle images
// /customExercises/thumbnail/
// /customExercises/image/
func (h *AppHandler) RouterImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("here")
	params := mux.Vars(r)
	filename := params["fileName"]
	locationType := params["locationType"]

	if filename == "" || locationType == "" {
		DefaultErrorHandler(errors.New("bad request"), w)
		return
	}

	if locationType != "profileImage" {
		DefaultErrorHandler(errors.New("bad request"), w)
		return
	}

	var locationPrefix string
	var exists bool

	switch locationType {
	case "profileImage":
		locationPrefix, exists = os.LookupEnv("LocationProfileImage")

		if !exists {
			DefaultErrorHandler(errors.New("enviroment variable is not set"), w)
			return
		}

	}
	h.ImageHandler(locationPrefix, filename, w, r)
}

// ImageHandler used to handle topic thread
func (h *AppHandler) ImageHandler(fileURI string, fileName string, w http.ResponseWriter, r *http.Request) {

	//Check if file exists and open
	Openfile, err := os.Open(fileURI + fileName)

	if err != nil {
		//File not found, send 404
		DefaultErrorHandler(errors.New("not found | "+fileName+" doesn't exist"), w)
		return
	}
	defer Openfile.Close() //Close after function return

	//File is found, create and send the correct headers
	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(w, Openfile) //'Copy' the file to the client
}
