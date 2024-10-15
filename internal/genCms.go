package internal

import (
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func GenerateCMS() {

	cmd := exec.Command("openssl", "cms",
		"-sign", "-in", "MiLoginTicketRequest.xml",
		"-out", "MiLoginTicketRequest.xml.cms",
		"-signer", "MiCertificado.pem",
		"-inkey", "MiClavePrivada.key",
		"-nodetach", "-outform", "PEM",
	)

	stdout, err := cmd.Output()
	if err != nil {
		logrus.Error(err)
		return
	}

	fmt.Println(string(stdout))
}
