package directions_results

import (
	"awesomeProject/pkg/byteScale"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"path/filepath"
	"time"
)

func LoadFile(file multipart.FileHeader) (string, error) {
	ext := filepath.Ext(file.Filename)

	var header string
	switch ext {
	case ".pdf":
		header = "application/pdf"
	case ".docx":
		header = "Content-Type: application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	default:
		return "", errors.New("Invalid file format")
	}

	slog.Info("Header: ", header)

	newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	fileOpened, err := file.Open()
	if err != nil {
		slog.Info("open error - " + err.Error())
		return "", err
	}
	defer fileOpened.Close()

	fileByte, err := io.ReadAll(fileOpened)
	if err != nil {
		slog.Info("file read error - " + err.Error())
		return "", err
	}

	fileUrl, err := byteScale.UploadFile(fileByte, header, newFileName)
	if err != nil {
		slog.Info("loadFile Error - " + err.Error())
		return "", err
	}

	return fileUrl, nil
}
