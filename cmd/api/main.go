package main

import "cmd/api/main.go/internal"

func main() {

	//Creacion de la clave privada rsa de 2048 bits}
	//1er paso

	// filename := "MiClavePrivada"

	// keyPair, err := rsa.GenerateKey(rand.Reader, 2048)
	// if err != nil {
	// 	log.Error(err)
	// 	return
	// }

	// keyPEM := pem.EncodeToMemory(
	// 	&pem.Block{
	// 		Type:  "RSA PRIVATE KEY",
	// 		Bytes: x509.MarshalPKCS1PrivateKey(keyPair),
	// 	},
	// )

	// if err := os.WriteFile(filename+".key", keyPEM, 0700); err != nil {
	// 	log.Error(err)
	// 	return
	// }

	// fmt.Println("Private .key Created")

	// //create csr
	// //paso 2

	// privateKeyBytes, err := os.ReadFile("MiClavePrivada.key")
	// if err != nil {
	// 	log.Error(err)
	// 	return
	// }
	// block, _ := pem.Decode(privateKeyBytes)
	// privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	// if err != nil {
	// 	log.Error(err)
	// 	return
	// }

	// template := x509.CertificateRequest{
	// 	Subject: pkix.Name{
	// 		Country:      []string{"AR"},
	// 		Organization: []string{"MAC"},
	// 		CommonName:   "carlostest3",
	// 		ExtraNames: []pkix.AttributeTypeAndValue{
	// 			{Type: []int{2, 5, 4, 5}, Value: "CUIT 20406165281"},
	// 		},
	// 	},
	// }

	// crsBytes, err := x509.CreateCertificateRequest(rand.Reader, &template, privateKey)
	// if err != nil {
	// 	log.Error(err)
	// 	return
	// }

	// csrFile, err := os.Create("MiPedidoCSR.csr")
	// if err != nil {
	// 	log.Error(err)
	// }
	// defer csrFile.Close()

	// pem.Encode(csrFile, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: crsBytes})

	// fmt.Println("CSR generated successfully and saved to MiPedidoCSR.csr")

	internal.RequestAuth()
}
