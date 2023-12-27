package models

type TemplateData struct {
    StringMap map[string]string
    IntMap map[int]int
    FloatMap map[float32]float32
    Data map[string]interface{}
}

type User struct {
    UserID string   `json:"userId"`
    Name string     `json:"name"`
    Email string    `json:"email"`
    Password string 
}

type File struct {
    FileName string `json:"filename"`
    FileType string `json:"filetype"`
    FileSize int64  `json:"filesize"`
}

type Response struct {
    Message string `json:"message"`
    Data  File     `json:"file_info"`
}

