package server

import (
	"os"
	"path/filepath"
	"net/http"
	"errors"

	// log "github.com/sirupsen/logrus"
)

type VideoFile struct {
	Name        string `json:"name"`
	DownloadURL string `json:"downloadPath"`
	Size        int64  `json:"size"`
}

func videoFiles(output string) *Response {
	folder, err := os.Open(output)
	toReturn := make([]VideoFile, 0)

	if err != nil {
		return NewErrorResponse("Error opening folder", 500, err)
	}
	files, err := folder.Readdir(0)
	if err != nil {
		return NewErrorResponse("Error opening files in folder", 500, err)
	}

	for _, f := range files {
		if !f.IsDir() {
			name := f.Name()
			if _, err := contentType(name); err == nil {
				toReturn = append(toReturn, VideoFile{Name: name, DownloadURL: "/api/v1/recordings/" + name, Size: f.Size() })
			}
		}
	}
	return NewSuccessJsonResponse(toReturn, http.StatusOK)
}

func videoFile(output string, filename string) *Response {
	name := filepath.Join(output, filename)
	contentType, err := contentType(name)
	if err != nil {
		return NewErrorResponse("Unknown file", 404, nil)
	}

	fileBytes, err := os.ReadFile(name)
	if errors.Is(err, os.ErrNotExist) {
		return NewErrorResponse("Unknown file", 404, err)
	} else if err != nil {
		return NewErrorResponse("Error getting file", 500, err)
	}

	write := func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Content-Disposition", "attachment; filename=" + filename)
		w.WriteHeader(http.StatusOK)
		w.Write(fileBytes)
	}
	return &Response{Write: write}
}

func deleteVideoFile(output string, filenane string) *Response {
	name := filepath.Join(output, filenane)
	_, err := contentType(name)
	if err != nil {
		return NewErrorResponse("Unknown file", 404, err)
	}

	err = os.Remove(name)
	if errors.Is(err, os.ErrNotExist) {
		return NewErrorResponse("Unknown file", 404, err)
	} else if err != nil {
		return NewErrorResponse("Error getting file", 500, err)
	}

	return NewSuccessJsonResponse(map[string]string{"satus":"ok"}, http.StatusOK)
}

func contentType(name string) (string, error) {
	switch filepath.Ext(name) {
	case ".avi":
		return "video/x-msvideo", nil
	default:
		return "", errors.New("Unknown filetype")
	}
}
