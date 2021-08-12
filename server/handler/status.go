package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Festivals-App/festivals-fileserver/server/config"
)

type ServerStatus struct {
	Images                         int    `json:"num_images"`
	ResizedImages                  int    `json:"num_resized_images"`
	ImagesSize                     int64  `json:"size_images"`
	ImagesSizeHumanReadable        string `json:"size_images_human_readable"`
	ResizedImagesSize              int64  `json:"size_resized_images"`
	ResizedImagesSizeHumanReadable string `json:"size_resized_images_human_readable"`
	Comment                        string `json:"comment"`
}

type ServerFiles struct {
	Images []string `json:"images"`
	PDFs   []string `json:"pdfs"`
}

func Status(conf *config.Config, w http.ResponseWriter, _ *http.Request) {

	// fetch all original images
	imageSearchPattern := conf.StorageURL + "/" + "upload-*"
	images, err := filepath.Glob(imageSearchPattern)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var imagesSize int64
	for _, i := range images {
		ii, _ := os.Stat(i)
		imagesSize = imagesSize + ii.Size()
	}

	// fetch all resized images
	resizedSearchPattern := conf.ResizeStorageURL + "/" + "*_upload-*"
	resizedImages, err := filepath.Glob(resizedSearchPattern)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var resizedImagesSize int64
	for _, ri := range resizedImages {
		rii, _ := os.Stat(ri)
		resizedImagesSize = resizedImagesSize + rii.Size()
	}

	// create status struct
	status := ServerStatus{
		Images:                         len(images),
		ResizedImages:                  len(resizedImages),
		ImagesSize:                     imagesSize,
		ImagesSizeHumanReadable:        ByteCountSI(imagesSize),
		ResizedImagesSize:              resizedImagesSize,
		ResizedImagesSizeHumanReadable: ByteCountSI(resizedImagesSize),
		Comment:                        "",
	}

	respondJSON(w, 200, status)
}

func Files(conf *config.Config, w http.ResponseWriter, _ *http.Request) {

	// fetch all original images
	imageSearchPattern := conf.StorageURL + "/" + "upload-*"
	images, err := filepath.Glob(imageSearchPattern)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// fetch all pdfs
	pdfSearchPattern := conf.StorageURL + "/" + "upload-*"
	pdfs, err := filepath.Glob(pdfSearchPattern)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var imageNames []string
	for _, i := range images {
		_, fileName := filepath.Split(i)
		imageNames = append(imageNames, fileName)
	}

	var pdfNames []string
	for _, i := range pdfs {
		_, fileName := filepath.Split(i)
		pdfNames = append(pdfNames, fileName)
	}

	respondJSON(w, 200, &ServerFiles{
		Images: imageNames,
		PDFs:   pdfNames,
	})
}

func ByteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}
