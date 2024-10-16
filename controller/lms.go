package controller

import (
	"net/http"

	"github.com/AkuLaper/akulaperbe/config"
	"github.com/AkuLaper/akulaperbe/helper/at"
	"github.com/AkuLaper/akulaperbe/helper/lms"
	"github.com/AkuLaper/akulaperbe/model"
)

func GetCountDocUser(w http.ResponseWriter, r *http.Request) {
	var resp model.Response
	rkp, err := lms.GetRekapPendaftaranUsers(config.Mongoconn)
	if err != nil {
		resp.Response = err.Error()
		at.WriteJSON(w, http.StatusConflict, resp)
		return
	}
	at.WriteJSON(w, http.StatusOK, rkp)
}

func RefreshLMSCookie(respw http.ResponseWriter, req *http.Request) {
	var resp model.Response
	err := lms.RefreshCookie(config.Mongoconn)
	if err != nil {
		resp.Response = err.Error()
		at.WriteJSON(respw, http.StatusBadRequest, resp)
		return
	}
	resp.Info = "ok"
	at.WriteJSON(respw, http.StatusOK, resp)
}
