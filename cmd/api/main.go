package main

import "cmd/api/main.go/internal"

func main() {
	internal.RequestAuth()
	internal.GenerateCMS()
	internal.CreateLoginRequest()
	internal.MakeRequest()
}
