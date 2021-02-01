package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"msig/api/response"
	"msig/common/apperrors"
	"msig/conf"
	"msig/models"
	"msig/repos"
	"msig/services"
	"net/http"
)

func (api *API) AuthRequest(w http.ResponseWriter, r *http.Request) {
	net, err := ToNetwork(mux.Vars(r)["network"])
	if err != nil {
		response.JsonError(w, apperrors.New(apperrors.ErrBadParam, "network"))
		return
	}

	auth, err := api.provider.GetAuthProvider(net)
	if err != nil {
		response.JsonError(w, apperrors.New(apperrors.ErrBadParam))
		return
	}

	db, err := api.provider.GetDb(net)
	if err != nil {
		response.JsonError(w, apperrors.New(apperrors.ErrBadParam))
		return
	}

	var req models.AuthTokenReq
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.JsonError(w, apperrors.New(apperrors.ErrBadRequest))
		return
	}

	err = req.Validate()
	if err != nil {
		response.JsonError(w, apperrors.New(apperrors.ErrBadRequest, err.Error()))
		return
	}

	service := services.New(repos.New(db), nil, auth, net)

	resp, err := service.AuthRequest(req)
	if err != nil {
		response.JsonError(w, err)
		return
	}

	response.Json(w, resp)
}

func (api *API) Auth(w http.ResponseWriter, r *http.Request) {
	net, err := ToNetwork(mux.Vars(r)["network"])
	if err != nil {
		response.JsonError(w, apperrors.New(apperrors.ErrBadParam, "network"))
		return
	}

	auth, err := api.provider.GetAuthProvider(net)
	if err != nil {
		response.JsonError(w, apperrors.New(apperrors.ErrBadParam))
		return
	}

	db, err := api.provider.GetDb(net)
	if err != nil {
		response.JsonError(w, apperrors.New(apperrors.ErrBadParam))
		return
	}

	var req models.SignatureReq
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.JsonError(w, apperrors.New(apperrors.ErrBadRequest))
		return
	}

	err = req.Validate()
	if err != nil {
		response.JsonError(w, apperrors.New(apperrors.ErrBadRequest, err.Error()))
		return
	}

	service := services.New(repos.New(db), nil, auth, net)

	resp, err := service.Auth(req)
	if err != nil {
		response.JsonError(w, err)
		return
	}

	api.setCookie(net, resp.EncodedCookie, w)

	response.Json(w, resp)
}

func (api *API) RefreshAuth(w http.ResponseWriter, r *http.Request) {
	net, err := ToNetwork(mux.Vars(r)["network"])
	if err != nil {
		response.JsonError(w, apperrors.New(apperrors.ErrBadParam, "network"))
		return
	}

	auth, err := api.provider.GetAuthProvider(net)
	if err != nil {
		response.JsonError(w, apperrors.New(apperrors.ErrBadParam))
		return
	}

	db, err := api.provider.GetDb(net)
	if err != nil {
		response.JsonError(w, apperrors.New(apperrors.ErrBadParam))
		return
	}

	var data struct {
		RefreshToken string `json:"refresh_token"`
	}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response.JsonError(w, apperrors.New(apperrors.ErrBadRequest))
		return
	}

	if data.RefreshToken == "" {
		response.JsonError(w, apperrors.New(apperrors.ErrBadParam, "refresh_token"))
		return
	}

	service := services.New(repos.New(db), nil, auth, net)

	resp, err := service.RefreshAuthSession(data.RefreshToken)
	if err != nil {
		response.JsonError(w, err)
		return
	}

	api.setCookie(net, resp.EncodedCookie, w)

	response.Json(w, resp)
}

func (api *API) RestoreAuth(w http.ResponseWriter, r *http.Request) {
	net, err := ToNetwork(mux.Vars(r)["network"])
	if err != nil {
		response.JsonError(w, apperrors.New(apperrors.ErrBadParam, "network"))
		return
	}

	cookie, err := r.Cookie(getCookieName(net))
	if err != nil {
		response.JsonError(w, apperrors.New(apperrors.ErrBadAuth))
		return
	}

	//TODO move to service
	auth, err := api.provider.GetAuthProvider(net)
	if err != nil {
		response.JsonError(w, apperrors.New(apperrors.ErrBadParam))
		return
	}

	tokens, err := auth.DecodeSessionCookie(cookie.Value)
	if err != nil {
		response.JsonError(w, err)
		return
	}

	response.Json(w, map[string]interface{}{
		"tokens": tokens,
	})
}

func (api *API) Logout(w http.ResponseWriter, r *http.Request) {
	net, err := ToNetwork(mux.Vars(r)["network"])
	if err != nil {
		response.JsonError(w, apperrors.New(apperrors.ErrBadParam, "network"))
		return
	}

	cookie, err := r.Cookie(getCookieName(net))
	if err != nil || cookie.Value == "" {
		response.Json(w, map[string]interface{}{"message": "success"})
		return
	}

	auth, err := api.provider.GetAuthProvider(net)
	if err != nil {
		response.JsonError(w, apperrors.New(apperrors.ErrBadParam))
		return
	}

	db, err := api.provider.GetDb(net)
	if err != nil {
		response.JsonError(w, apperrors.New(apperrors.ErrBadParam))
		return
	}

	defer api.clearCookie(net, w)

	service := services.New(repos.New(db), nil, auth, net)

	err = service.Logout(cookie.Value)
	if err != nil {
		response.JsonError(w, err)
		return
	}

	response.Json(w, map[string]interface{}{"message": "success"})
}

func getCookieName(net models.Network) string {
	return fmt.Sprintf("%s_%s", "session", string(net))
}

func (api *API) setCookie(net models.Network, encodedCookie string, w http.ResponseWriter) {

	cookie := &http.Cookie{
		Name:     getCookieName(net),
		Value:    encodedCookie,
		Path:     "/",
		MaxAge:   conf.TtlCookie,
		Secure:   api.cfg.API.IsProtocolHttps,
		HttpOnly: true,

		//TODO remove after tests
		SameSite: http.SameSiteNoneMode,
	}

	http.SetCookie(w, cookie)
}

func (api *API) clearCookie(network models.Network, w http.ResponseWriter) {
	clearCookie := &http.Cookie{
		Name:     getCookieName(network),
		Value:    "{}",
		Path:     "/",
		MaxAge:   conf.TtlCookie,
		Secure:   api.cfg.API.IsProtocolHttps,
		HttpOnly: true,
	}

	http.SetCookie(w, clearCookie)
}