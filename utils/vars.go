package utils

import "os"

var (
	ApiUrlPrefix   = "/api/v1"
	ServerAddress  = "127.0.0.1:8000"
	SecretKeyToken = []byte(os.Getenv("SECRET_KEY_JWT"))
)
