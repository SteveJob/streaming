package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"modernday.com/stevejob/streaming/model"
	"modernday.com/stevejob/streaming/util"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	reqModel := model.UserCredential{}

	if err := json.Unmarshal(reqBody, &reqModel); err != nil {
		util.FlushResponse(w, -1, "参数解析出错")
	}
}

func Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := params.ByName("id")
	_, _ = io.WriteString(w, userId)
}