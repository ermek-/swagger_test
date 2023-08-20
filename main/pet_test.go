package main

import (
	"context"
	"fmt"
	"github.com/ozontech/cute"
	"github.com/ozontech/cute/asserts/json"
	"io"
	"net/http"
	"testing"
	"time"
)

func Test_FindByStatus(t *testing.T) {
	url := host + "/v2/pet/findByStatus"
	cute.NewTestBuilder().
		Epic("Swagger Petstore").
		Story("Pet").
		Feature("GET /pet/findByStatus").
		Title("Find pets by status").
		Tags("regress", "smoke").
		CreateStep("Find pets by status = available").
		RequestBuilder(
			cute.WithURI(url),
			cute.WithMethod(http.MethodGet),
			cute.WithQueryKV("status", "available"),
		).
		ExpectExecuteTimeout(5*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			//TODO: научиться проверять статусы всех записей
			json.Equal("$[0].status", "available"),
		).
		NextTest().
		CreateStep("Find pets by status = sold").
		RequestBuilder(
			cute.WithURI(url),
			cute.WithMethod(http.MethodGet),
			cute.WithQueryKV("status", "sold"),
		).
		ExpectExecuteTimeout(5*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Equal("$[0].status", "sold"),
		).
		NextTest().
		CreateStep("Find pets by status = pending").
		RequestBuilder(
			cute.WithURI(url),
			cute.WithMethod(http.MethodGet),
			cute.WithQueryKV("status", "pending"),
		).
		ExpectExecuteTimeout(5*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Equal("$[0].status", "pending"),
		).
		ExecuteTest(context.Background(), t)
}

func Test_AddPet(t *testing.T) {
	var dRequest = "null"
	cute.NewTestBuilder().
		Epic("Swagger Petstore").
		Story("Pet").
		Feature("POST /pet").
		Title("Add new pet with only required params").
		Tags("regress", "smoke").
		CreateStep("Add new pet").
		RequestBuilder(
			cute.WithURI(petUrl),
			cute.WithMethod(http.MethodPost),
			cute.WithHeadersKV("accept", "application/json"),
			cute.WithHeadersKV("Content-Type", "application/json"),
			cute.WithBody([]byte(petData)),
		).
		ExpectExecuteTimeout(5*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.NotEmpty("id"),
			json.Equal("$.name", petName),
			json.Equal("$.photoUrls", "["+photoUrls+"]"),
		).
		NextTest().
		AfterTestExecute(
			func(response *http.Response, errors []error) error {
				b, err := io.ReadAll(response.Body)
				if err != nil {
					return err
				}

				temp, err := json.GetValueFromJSON(b, "id")
				if err != nil {
					return err
				}

				dRequest = fmt.Sprint(temp)

				return nil
			},
		).
		CreateStep("Find pets by id").
		RequestBuilder(
			cute.WithURI(petUrl+dRequest),
			cute.WithMethod(http.MethodGet),
		).
		ExpectExecuteTimeout(5*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
		//json.Equal("$[0].status", "sold"),
		).
		ExecuteTest(context.Background(), t)
}
