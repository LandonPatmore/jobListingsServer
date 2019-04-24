package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/landonp1203/goUtils/aws"
	"github.com/landonp1203/jobListingsWorker/types"
	"net/http"
)

var dynamoClient *dynamodb.DynamoDB

const TableName = "Job-Listings"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Go to "+IDPath+"/all or "+IDPath+"/status")
}

func GetAllDBRowsHandler(w http.ResponseWriter, r *http.Request) {
	err := CreateDynamoClient()

	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}

	var jobs [] types.GithubJob
	err = aws.GetAllItems(dynamoClient, TableName, &jobs)

	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}

	respondWithJSON(w, http.StatusOK, jobs)
}

func GetDBStatusHandler(w http.ResponseWriter, r *http.Request) {
	err := CreateDynamoClient()

	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}

	type status struct {
		Table       string
		RecordCount int64
	}

	count, err := aws.GetRowCount(dynamoClient, TableName)

	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}

	respondWithJSON(w, http.StatusOK, status{Table: TableName, RecordCount: count})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.MarshalIndent(payload, "", "    ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func CreateDynamoClient() error {
	if dynamoClient == nil {
		client, err := aws.CreateDynamoClient()

		if err != nil {
			return err
		}

		dynamoClient = client
	}

	return nil
}
