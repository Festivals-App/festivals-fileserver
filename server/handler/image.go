package handler

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Festivals-App/festivals-fileserver/server/config"
	"github.com/Festivals-App/festivals-fileserver/server/manipulate"
	servertools "github.com/Festivals-App/festivals-server-tools"
	"github.com/go-chi/chi/v5"
)

var kMaxFileSize int64 = 10 << 20

// GET functions

func MultipartUpload(conf *config.Config, w http.ResponseWriter, r *http.Request) {

	// limit the request to kMaxFileSize
	r.Body = http.MaxBytesReader(w, r.Body, kMaxFileSize+512)
	// Parse our multipart form, kMacFileSize specifies a maximum
	// upload of 10 MB files.
	err := r.ParseMultipartForm(kMaxFileSize)
	if err != nil {
		servertools.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, _, err := r.FormFile("image")
	if err != nil {
		servertools.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	// create intermidiate dirs if needed
	err = os.MkdirAll(conf.StorageURL, os.ModePerm)
	if err != nil {
		servertools.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile(conf.StorageURL, "upload-*.jpg")
	if err != nil {
		servertools.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		servertools.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	// write this byte array to our temporary file
	_, err = tempFile.Write(fileBytes)
	if err != nil {
		servertools.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	// return that we have successfully uploaded our file!
	path := tempFile.Name()
	_, fileName := filepath.Split(path)
	servertools.RespondJSON(w, 201, fileName)
}

func Download(conf *config.Config, w http.ResponseWriter, r *http.Request) {

	// get image file name
	objectID := chi.URLParam(r, "imageIdentifier")
	// create path to original file and check if it exists
	imagepath := filepath.Join(conf.StorageURL, objectID)
	if !manipulate.FileExists(imagepath) {
		servertools.RespondError(w, http.StatusBadRequest, "File does not exist.")
		return
	}
	// get query values if the exist
	values := r.URL.Query()
	if len(values) == 0 {

		img, err := os.Open(imagepath)
		// we assume the image does not exist
		if err != nil {
			servertools.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}

		respondFile(w, img)
		return
	}

	resizedImage, err := manipulate.ResizeIfNeeded(conf, objectID, values)
	if err != nil {
		servertools.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondFile(w, resizedImage)
}

func Update(conf *config.Config, w http.ResponseWriter, r *http.Request) {

	// get image file name
	objectID := chi.URLParam(r, "imageIdentifier")
	// create path to original file and check if it exists
	imagepath := filepath.Join(conf.StorageURL, objectID)
	if !manipulate.FileExists(imagepath) {
		servertools.RespondError(w, http.StatusBadRequest, "File does not exist.")
		return
	}
	// limit the request to kMaxFileSize
	r.Body = http.MaxBytesReader(w, r.Body, kMaxFileSize+512)
	// Parse our multipart form, kMacFileSize specifies a maximum
	// upload of 10 MB files.
	err := r.ParseMultipartForm(kMaxFileSize)
	if err != nil {
		servertools.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, _, err := r.FormFile("image")
	if err != nil {
		servertools.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	// create intermediate dirs if needed
	err = os.MkdirAll(conf.StorageURL, os.ModePerm)
	if err != nil {
		servertools.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		servertools.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	err = ioutil.WriteFile(imagepath, fileBytes, os.ModePerm)
	if err != nil {
		servertools.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	// remove old scaled versions
	searchPatern := conf.ResizeStorageURL + "/" + "*_" + objectID
	files, err := filepath.Glob(searchPatern)
	if err != nil {
		servertools.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			servertools.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	// return that we have successfully uploaded our file!
	servertools.RespondJSON(w, 201, objectID)
}
