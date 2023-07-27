package main

import (
	"context"
	"github.com/ozontech/cute"
	"github.com/ozontech/cute/asserts/json"
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
		ExpectExecuteTimeout(10*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Equal("$[0].status", "available"),
		).
		NextTest().
		CreateStep("Find pets by status = sold").
		RequestBuilder(
			cute.WithURI(url),
			cute.WithMethod(http.MethodGet),
			cute.WithQueryKV("status", "sold"),
		).
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
		AssertBody(
			json.Equal("$[0].status", "pending"),
		).
		ExecuteTest(context.Background(), t)
}
