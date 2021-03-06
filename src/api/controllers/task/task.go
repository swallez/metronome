package taskCtrl

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/runabove/metronome/src/api/core"
	"github.com/runabove/metronome/src/api/core/io/in"
	"github.com/runabove/metronome/src/api/core/io/out"
	authSrv "github.com/runabove/metronome/src/api/services/auth"
	taskSrv "github.com/runabove/metronome/src/api/services/task"
	"github.com/runabove/metronome/src/metronome/models"
)

// schedule regex: https://regex101.com/r/vyBrRd/3
func Create(w http.ResponseWriter, r *http.Request) {
	token := authSrv.GetToken(r.Header.Get("Authorization"))
	if token == nil {
		out.Unauthorized(w)
		return
	}

	var task models.Task

	body, err := in.JSON(r, &task)
	if err != nil {
		out.JSON(w, 400, err)
		return
	}

	result := core.ValidateJSON("task", "create", string(body))
	if !result.Valid {
		out.JSON(w, 422, result.Errors)
		return
	}

	task.UserId = authSrv.UserId(token)

	success := taskSrv.Create(&task)
	if !success {
		out.BadGateway(w)
		return
	}
	out.JSON(w, 200, task)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	token := authSrv.GetToken(r.Header.Get("Authorization"))
	if token == nil {
		out.Unauthorized(w)
		return
	}

	success := taskSrv.Delete(mux.Vars(r)["id"], authSrv.UserId(token))
	if !success {
		out.BadGateway(w)
		return
	}
	out.Success(w)
}
