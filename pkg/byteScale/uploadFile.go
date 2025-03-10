package byteScale

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type filePath struct {
	Path string `json:"fileUrl"`
}

func UploadFile(file []byte, contentType string, fileName string) (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := "https://api.bytescale.com/v2/accounts/G22nhbT/uploads/binary?fileName=" + fileName
	token := fmt.Sprintf("Bearer %s", os.Getenv("TOKEN_BYTE_SCALE"))

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(file))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	var pathUploaded filePath
	err = json.NewDecoder(resp.Body).Decode(&pathUploaded)
	if err != nil {
		return "", err
	}

	slog.Info("status - " + resp.Status)

	return pathUploaded.Path, nil
}
