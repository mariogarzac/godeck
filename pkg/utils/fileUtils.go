package utils

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

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

    if err != nil {
        log.Println("no file to read", err)
        return fileType, fileName, fileSize , err
    }

    switch fileType{
    case "application/pdf":
        fileName = "static/uploads/" + strconv.FormatInt(time.Now().Unix(), 10) + ".pdf"
    default :
        fileName = "static/uploads/" + strconv.FormatInt(time.Now().Unix(), 10) + ".jpg"
    }

    err = os.WriteFile(fileName, fileByte, 0777)

    if err != nil {
        log.Println("No file to write to",err)
        return fileType, fileName, fileSize, err
    }

    fileSize = file.Size

    return fileType, fileName, fileSize, nil
}

