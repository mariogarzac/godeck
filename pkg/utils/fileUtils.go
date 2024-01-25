package utils

import (
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func UploadFile(file *multipart.FileHeader) (string, string, int64, error) {
    var fileType, fileName string
    var fileSize int64

    // open file
    src, err := file.Open()
    if err != nil {
        log.Println("error opening file", err)
        return fileType, fileName, fileSize, err
    }

    defer src.Close()

    // read file and get content type
    fileByte, err := io.ReadAll(src)
    fileType = http.DetectContentType(fileByte)
    fileName = file.Filename 

    if err != nil {
        log.Println("no file to read", err)
        return fileType, fileName, fileSize , err
    }

    var filePath string

    if fileType != "audio/mpeg" {
        return fileType, fileName, fileSize , errors.New("File is not an mp3")
    }

    filePath = "static/uploads/" + fileName

    err = os.WriteFile(filePath, fileByte, 0644)

    if err != nil {
        log.Println("Could not write file:",err)
        return fileType, fileName, fileSize, err
    }

    fileSize = file.Size

    return fileType, fileName, fileSize, nil
}

