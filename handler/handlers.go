package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type APIHandler interface {
	RegisterRoutes(chi.Router)
}

var ()

type APIResponseErrJson struct {
	ErrCode string      `json:"errcode"`
	ErrData interface{} `json:"errdata"`
}

type EmptyData struct{}

type APIResponse struct {
	Err  interface{} `json:"err"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func APIResponseOK(w http.ResponseWriter, r *http.Request, data interface{}, msg string) {
	responseobj := &APIResponse{
		Err:  nil,
		Data: data,
		Msg:  msg,
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, responseobj)
}

func APIResponseBadRequest(w http.ResponseWriter, r *http.Request, errorcode string, msg string, errData interface{}) {
	responseobj := &APIResponse{
		Err: &APIResponseErrJson{
			ErrCode: errorcode,
			ErrData: errData,
		},
		Data: nil,
		Msg:  msg,
	}
	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, responseobj)
}

func APIResponseUnauthorized(w http.ResponseWriter, r *http.Request, errorcode string, msg string, errData interface{}) {
	responseobj := &APIResponse{
		Err: &APIResponseErrJson{
			ErrCode: errorcode,
			ErrData: errData,
		},
		Data: nil,
		Msg:  msg,
	}
	render.Status(r, http.StatusUnauthorized)
	render.JSON(w, r, responseobj)
}

func APIResponseForbidden(w http.ResponseWriter, r *http.Request, errorcode string, msg string, errData interface{}) {
	responseobj := &APIResponse{
		Err: &APIResponseErrJson{
			ErrCode: errorcode,
			ErrData: errData,
		},
		Data: nil,
		Msg:  msg,
	}
	render.Status(r, http.StatusForbidden)
	render.JSON(w, r, responseobj)
}

func APIResponseConflict(w http.ResponseWriter, r *http.Request, errorcode string, msg string, errData interface{}) {
	responseobj := &APIResponse{
		Err: &APIResponseErrJson{
			ErrCode: errorcode,
			ErrData: errData,
		},
		Data: nil,
		Msg:  msg,
	}
	render.Status(r, http.StatusConflict)
	render.JSON(w, r, responseobj)
}

func APIResponseGone(w http.ResponseWriter, r *http.Request, errorcode string, msg string, errData interface{}) {
	responseobj := &APIResponse{
		Err: &APIResponseErrJson{
			ErrCode: errorcode,
			ErrData: errData,
		},
		Data: nil,
		Msg:  msg,
	}
	render.Status(r, http.StatusGone)
	render.JSON(w, r, responseobj)
}

func APIResponseUnprocessableEntity(w http.ResponseWriter, r *http.Request, errorcode string, msg string, errData interface{}) {
	responseobj := &APIResponse{
		Err: &APIResponseErrJson{
			ErrCode: errorcode,
			ErrData: errData,
		},
		Data: nil,
		Msg:  msg,
	}
	render.Status(r, http.StatusUnprocessableEntity)
	render.JSON(w, r, responseobj)
}

func APIResponseNotAcceptable(w http.ResponseWriter, r *http.Request, errorcode string, msg string, errData interface{}) {
	responseobj := &APIResponse{
		Err: &APIResponseErrJson{
			ErrCode: errorcode,
			ErrData: errData,
		},
		Data: nil,
		Msg:  msg,
	}
	render.Status(r, http.StatusNotAcceptable)
	render.JSON(w, r, responseobj)
}

func APIResponseInternalServerError(w http.ResponseWriter, r *http.Request, errorcode string, msg string, errData interface{}) {
	responseobj := &APIResponse{
		Err: &APIResponseErrJson{
			ErrCode: errorcode,
			ErrData: errData,
		},
		Data: nil,
		Msg:  msg,
	}
	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, responseobj)
}

func APIFailedInternalAPICall(w http.ResponseWriter, r *http.Request, errorcode string, msg string, statusCode int, errData interface{}) {
	responseobj := &APIResponse{
		Err: &APIResponseErrJson{
			ErrCode: errorcode,
			ErrData: errData,
		},
		Data: nil,
		Msg:  msg,
	}
	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, responseobj)
}
