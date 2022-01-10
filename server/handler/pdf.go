package handler

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Festivals-App/festivals-fileserver/server/config"
	"github.com/Festivals-App/festivals-fileserver/server/manipulate"
	"github.com/go-chi/chi/v5"
)

var kMaxPDFSize int64 = 10 << 20

// GET functions

func MultipartPDFUpload(conf *config.Config, w http.ResponseWriter, r *http.Request) {

	// limit the request to kMaxPDFSize
	r.Body = http.MaxBytesReader(w, r.Body, kMaxPDFSize+512)
	// Parse our multipart form, kMaxPDFSize specifies a maximum
	// upload of 10 MB files.
	err := r.ParseMultipartForm(kMaxFileSize)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	// FormFile returns the first file for the given key `pdf`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, _, err := r.FormFile("pdf")
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	defer file.Close()

	// create intermidiate dirs if needed
	err = os.MkdirAll(conf.StorageURL, os.ModePerm)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile(conf.StorageURL, "upload-*.pdf")
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	// write this byte array to our temporary file
	_, err = tempFile.Write(fileBytes)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	// return that we have successfully uploaded our file!
	path := tempFile.Name()
	_, fileName := filepath.Split(path)
	respondJSON(w, 201, fileName)
}

func DownloadPDF(conf *config.Config, w http.ResponseWriter, r *http.Request) {

	// get image file name
	objectID := chi.URLParam(r, "pdfIdentifier")
	// create path to original file and check if it exists
	pdfpath := filepath.Join(conf.StorageURL, objectID)
	if !manipulate.FileExists(pdfpath) {
		respondError(w, 404, "File does not exist.")
		return
	}

	pdf, err := os.Open(pdfpath)
	// we assume the pdf does not exist
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}

	respondFile(w, pdf)
}

func UpdatePDF(conf *config.Config, w http.ResponseWriter, r *http.Request) {

	// get image file name
	objectID := chi.URLParam(r, "pdfIdentifier")
	// create path to original file and check if it exists
	pdfpath := filepath.Join(conf.StorageURL, objectID)
	if !manipulate.FileExists(pdfpath) {
		respondError(w, 404, "File does not exist.")
		return
	}
	// limit the request to kMaxFileSize
	r.Body = http.MaxBytesReader(w, r.Body, kMaxPDFSize+512)
	// Parse our multipart form, kMaxPDFSize specifies a maximum
	// upload of 10 MB files.
	err := r.ParseMultipartForm(kMaxPDFSize)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, _, err := r.FormFile("pdf")
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	defer file.Close()

	// create intermediate dirs if needed
	err = os.MkdirAll(conf.StorageURL, os.ModePerm)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	err = ioutil.WriteFile(pdfpath, fileBytes, os.ModePerm)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	defer file.Close()

	// return that we have successfully uploaded our file!
	respondJSON(w, 201, objectID)
}
