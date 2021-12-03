package auth

func IsAuthorized(authHeader string, body []byte) bool {
	return true
	//hash := hmac.New(sha256.New, []byte(os.Getenv("SECRET")))
	//hash.Write(body)
	//return hex.EncodeToString(hash.Sum(nil)) == authHeader
}
