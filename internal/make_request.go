package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func MakeRequest() {
	file, err := os.Open("requestLoginCms.xml")
	if err != nil {
		log.Error(err)
		return
	}
	defer file.Close()

	fmt.Println(file)

	fmt.Println("Creating http request to AFIP Web Services")
	req, err := http.NewRequest("POST", "https://wsaahomo.afip.gov.ar/ws/services/LoginCms", file)
	if err != nil {
		log.Error(err)
		return
	}

	req.Header.Add("Content-Type", "text/html;charset=UTF-8")
	req.Header.Add("SOAPAction", "urn:LoginCms")
	req.Header.Add("User-Agent", "Apache-HttpClient/4.1.1 (java 1.5)")

	transport := &http.Transport{
		Proxy: nil,
	}

	client := &http.Client{
		Transport: transport,
	}

	res, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Println("RESPONSE: ", string(body))
}
