package internal

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"os"
	"net/http"
	"github.com/alexflint/go-arg"
	"github.com/cavaliergopher/grab/v3"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

func Init() error {
	arg.MustParse(&Args)
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(&Args)
	CheckIfError(err)
	var logLevel string
	if Args.LogLevel == ""{
		logLevel = "Debug"
	} else {
		logLevel = Args.LogLevel
	}
	parsedLogLevel, err := log.ParseLevel(logLevel)
	CheckIfError(err)
	log.SetLevel(parsedLogLevel)
	return nil
}


func DownloadAndConvertAvatar(imgUrl string) (body io.Reader) {
	resp, err := grab.Get("/tmp", imgUrl)
	CheckIfError(err)
	image, err := os.ReadFile(resp.Filename)
	CheckIfError(err)
	avatarb64 := base64.StdEncoding.EncodeToString(image)
	CheckIfError(err)
	bodyMap := map[string]string{"image":avatarb64}
	bodyBytes, _ := json.Marshal(bodyMap)
	body = bytes.NewBuffer(bodyBytes)
	return body
}

func ProcessAvatar (orgName string, avatarUrl string) error {
		client := &http.Client{}
		req, err := http.NewRequest("POST", GiteaLogin.GiteaUrl + "/api/v1/orgs/" + orgName + "/avatar", DownloadAndConvertAvatar(avatarUrl))
		CheckIfError(err)
	
		req.Header.Set("accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "token " + GiteaLogin.Password)
		resp, err := client.Do(req)
		CheckIfError(err)
		defer resp.Body.Close()
		return err
}