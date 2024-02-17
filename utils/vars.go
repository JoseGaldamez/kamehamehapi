package utils

import "os"

var (
	ApiUrlPrefix   = "/api/v1"
	SecretKeyToken = []byte(os.Getenv("SECRET_KEY_JWT"))
)
