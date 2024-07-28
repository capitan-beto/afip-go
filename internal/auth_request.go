package internal

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type AuthRequest struct {
	XMLName xml.Name `xml:"loginTicketRequest"`
	Version string   `xml:"version,attr"`
	Header  *Header
	Service string `xml:"service"`
}

type Header struct {
	XMLName        xml.Name `xml:"Header"`
	UniqueID       uint32   `xml:"uniqueId"`
	GenerationTime string   `xml:"generationTime"`
	ExpirationTime string   `xml:"expirationTime"`
}

func RequestAuth() {
	uid := rand.Uint32()
	dt := time.Now()
	gt := dt.Format("2006-01-02") + "T" + time.Now().Format("15:04:16")
	timein := time.Now().Add(time.Hour * 18)
	ed := timein.Format("2006-01-02") + "T" + timein.Format("15:04:16")

	newHead := &Header{UniqueID: uid, GenerationTime: gt, ExpirationTime: ed}
	newReq := &AuthRequest{Version: "1.0", Header: newHead, Service: "wtcmxa"}

	out, _ := xml.MarshalIndent(newReq, " ", " ")
	fmt.Println(xml.Header + string(out))
	err := os.WriteFile("MiLoginTicketRequest.xml", out, 0777)
	if err != nil {
		log.Error(err)
		return
	}

}
