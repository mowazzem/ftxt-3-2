package login

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

type loginHandler struct{}

func NewLoginHandler() *loginHandler {
	return &loginHandler{}
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (lh *loginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.Write([]byte("error occured: " + err.Error()))
		return
	}

	hash := sha1.New()
	hash.Write([]byte(credentials.Username))
	hash.Write([]byte(credentials.Password))

	shaChecksum := hex.EncodeToString(hash.Sum(nil))

	w.Header().Add("Content-type", "application/json")
	w.Write([]byte(shaChecksum))
}
