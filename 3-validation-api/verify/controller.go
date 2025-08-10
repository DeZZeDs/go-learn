package verify

import (
	"main/config"
	"net/http"
)

type VerifyHandlerController struct {
	verifyService *VerifyService
}

func VerifyHandler(router *http.ServeMux, config *config.AppConfig) {
	handler := &VerifyHandlerController{
		verifyService: &VerifyService{
			config: config,
		},
	}
	router.HandleFunc("POST /send", handler.VerifySendHandler())
	router.HandleFunc("/verify/{hash}", handler.VerifyGetHash())
}

func (handler *VerifyHandlerController) VerifySendHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.verifyService.SendEmail()
	}
}

func (handler *VerifyHandlerController) VerifyGetHash() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		w.Write([]byte("your hash " + hash + ": true"))
	}
}
