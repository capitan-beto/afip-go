package internal

import (
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func GenerateCMS() {
	fmt.Println("Step 2: generating ticket request's cms")

	cmd := exec.Command("openssl", "cms",
		"-sign", "-in", "MiLoginTicketRequest.xml",
		"-out", "MiLoginTicketRequest.xml.cms",
		"-signer", "MiCertificado.pem",
		"-inkey", "MiClavePrivada.key",
		"-nodetach", "-outform", "PEM",
	)

	_, err := cmd.Output()
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Println("Step 2: cms generated successfully")
}
