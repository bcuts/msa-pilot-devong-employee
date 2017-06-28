package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"msa-pilot-devong-employee/router"
	"encoding/json"
	"msa-pilot-devong-employee/model"
)

func TestGetEmployees(t *testing.T) {
	request, _ := http.NewRequest("GET", "/employees", nil)
	response := excuteRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)
	var employees model.Employees

	if error := json.Unmarshal(response.Body.Bytes(), &employees); error != nil {
		panic(error)
		t.Errorf("Excepted not empty arrays")
	}

	assert.Equal(t, 2, len(employees), "unexpected length")
	assert.Equal(t, "JIMMY", employees[0].FirstName, "unexpected employee information")
}

func checkResponseCode(t *testing.T, expected int, result int) {
	assert.Equal(t, expected, result, "statusCode is Not ok")
}

func excuteRequest(request *http.Request) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	router.NewRouter().ServeHTTP(responseRecorder, request)

	return responseRecorder
}

