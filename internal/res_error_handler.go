package internal

import (
	"encoding/xml"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Response struct {
	XMLName xml.Name
	Body    string
}

func ResponseErrorHandler(resBytes []byte) {
	res := &Response{}

	err := xml.Unmarshal(resBytes, &res)
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Println(res)
}
