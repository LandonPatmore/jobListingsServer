package main

import (
	"encoding/json"
	"fmt"
	"github.com/landonp1203/jobListingsWorker/utils"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Go to "+IDPath+"/all or "+IDPath+"/status")
}

func GetAllDBRowsHandler(w http.ResponseWriter, r *http.Request) {
	jobs, err := utils.GetAllItems()

	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}

	respondWithJSON(w, http.StatusOK, jobs)
}

func GetDBStatusHandler(w http.ResponseWriter, r *http.Request) {
	type status struct {
		Table       string
		RecordCount int64
	}

	count, err := utils.GetRowCount()

	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}

	respondWithJSON(w, http.StatusOK, status{Table: utils.TableName, RecordCount: count})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.MarshalIndent(payload, "", "    ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
