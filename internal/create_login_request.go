package internal

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Envelope struct {
	XMLName   xml.Name  `xml:"soapenv:Envelope"`
	XMLNSs    string    `xml:"xmlns:soapenv,attr"`
	XMLNSwsaa string    `xml:"xmlns:wsaa,attr"`
	Header    string    `xml:"soapenv:Header"`
	Body      *WSAABody `xml:"soapenv:Body"`
}

type WSAABody struct {
	LoginCms *LoginCms `xml:"wsaa:loginCms"`
}

type LoginCms struct {
	Cms xml.CharData `xml:"wsaa:in0"`
}

func CreateLoginRequest() {
	soapenv := "http://schemas.xmlsoap.org/soap/envelope/"
	wsaa := "http://wsaa.view.sua.dvadac.desein.afip.gov"

	encodedCms, err := GetCMS()
	if err != nil {
		log.Error(err)
		return
	}

	newCmsField := LoginCms{Cms: []byte(encodedCms)}
	newWSAABody := WSAABody{LoginCms: &newCmsField}
	newReq := Envelope{XMLNSs: soapenv, XMLNSwsaa: wsaa, Body: &newWSAABody}

	out, _ := xml.MarshalIndent(newReq, " ", "  ")
	res := EscapeXML(string(out))
	fmt.Println(res)
	err = os.WriteFile("requestLoginCms.xml", []byte(res), 0777)
	if err != nil {
		log.Error(err)
		return
	}
}

func GetCMS() (string, error) {
	fp := "MiLoginTicketRequest.xml.cms"
	file, err := os.Open(fp)
	if err != nil {
		log.Error(err)
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Error(err)
		return "", err
	}

	res := string(content)
	return res[20 : len(res)-19], nil
}

func EscapeXML(s string) string {
	s = strings.ReplaceAll(s, "&#xA;", "")
	return s
}
