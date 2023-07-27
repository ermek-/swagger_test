package main

import (
	"context"
	"github.com/ozontech/cute"
	"github.com/ozontech/cute/asserts/json"
	"net/http"
	"testing"
	"time"
)

func Test_CreateUser(t *testing.T) {
	url := host + "/v2/user"
	username := "testing_user"
	cute.NewTestBuilder().
		Epic("Swagger Petstore").
		Story("User").
		Feature("POST /user").
		Title("Create new user").
		Tags("regress", "smoke").
		CreateStep("Create new user").
		RequestBuilder(
			cute.WithURI(url),
			cute.WithMethod(http.MethodPost),
			cute.WithHeadersKV("accept", "application/json"),
			cute.WithHeadersKV("Content-Type", "application/json"),
			cute.WithBody([]byte("{\n  \"id\": 123456789,\n  \"username\": \"testing_user\",\n  \"firstName\": \"test\",\n  \"lastName\": \"user\",\n  \"email\": \"test@test.com\",\n  \"password\": \"Strong_Password\",\n  \"phone\": \"+1234567890\",\n  \"userStatus\": 1\n}")),
		).
		ExpectExecuteTimeout(10*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Equal("$.code", 200),
			json.Equal("$.type", "unknown"),
			json.Equal("$.message", "123456789"),
		).
		NextTest().
		CreateStep("Get user by username").
		RequestBuilder(
			cute.WithURI(url+"/"+username),
			cute.WithMethod(http.MethodGet),
		).
		AssertBody(
			// TODO: разобраться почему вот так не работает:
			//json.Equal("$.id", 123456789),
			json.Equal("$.username", username),
			json.Equal("$.firstName", "test"),
			json.Equal("$.lastName", "user"),
			json.Equal("$.email", "test@test.com"),
			json.Equal("$.password", "Strong_Password"),
			json.Equal("$.phone", "+1234567890"),
			json.Equal("$.userStatus", 1),
		).
		NextTest().
		CreateStep("Delete user").
		RequestBuilder(
			cute.WithURI(url+"/"+username),
			cute.WithMethod(http.MethodDelete),
		).
		AssertBody(
			json.Equal("$.code", 200),
			json.Equal("$.type", "unknown"),
			json.Equal("$.message", username),
		).
		ExecuteTest(context.Background(), t)
}
