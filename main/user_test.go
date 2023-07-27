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
	userData := "{\n  \"id\": 123456789,\n  \"username\": \"testing_user\",\n  \"firstName\": \"test\",\n  \"lastName\": \"user\",\n  \"email\": \"test@test.com\",\n  \"password\": \"Strong_Password\",\n  \"phone\": \"+1234567890\",\n  \"userStatus\": 1\n}"
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
			cute.WithBody([]byte(userData)),
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
		//TODO: подумать как вынести это из шага в postcondition
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

func Test_UpdateUser(t *testing.T) {
	url := host + "/v2/user"
	username := "testing_user"
	newUsername := "updating_username"
	userData := "{\n  \"id\": 123456789,\n  \"username\": \"testing_user\",\n  \"firstName\": \"test\",\n  \"lastName\": \"user\",\n  \"email\": \"test@test.com\",\n  \"password\": \"Strong_Password\",\n  \"phone\": \"+1234567890\",\n  \"userStatus\": 1\n}"
	newUserData := "{\n  \"id\": 123456789,\n  \"username\": \"updating_username\",\n  \"firstName\": \"updating_firstName\",\n  \"lastName\": \"updating_lastName\",\n  \"email\": \"test@test.com\",\n  \"password\": \"Strong_Password\",\n  \"phone\": \"+1234567890\",\n  \"userStatus\": 1\n}"
	cute.NewTestBuilder().
		Epic("Swagger Petstore").
		Story("User").
		Feature("PUT /user/{username}").
		Title("Update user").
		Tags("regress", "smoke").
		//TODO: подумать как вынести это в precondition
		CreateStep("Create new user").
		RequestBuilder(
			cute.WithURI(url),
			cute.WithMethod(http.MethodPost),
			cute.WithHeadersKV("accept", "application/json"),
			cute.WithHeadersKV("Content-Type", "application/json"),
			cute.WithBody([]byte(userData)),
		).
		ExpectExecuteTimeout(10*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Equal("$.code", 200),
			json.Equal("$.type", "unknown"),
			json.Equal("$.message", "123456789"),
		).
		NextTest().
		CreateStep("Update user").
		RequestBuilder(
			cute.WithURI(url+"/"+username),
			cute.WithMethod(http.MethodPut),
			cute.WithHeadersKV("accept", "application/json"),
			cute.WithHeadersKV("Content-Type", "application/json"),
			cute.WithBody([]byte(newUserData)),
		).
		AssertBody(
			json.Equal("$.code", 200),
			json.Equal("$.type", "unknown"),
			json.Equal("$.message", "123456789"),
		).
		NextTest().
		CreateStep("Get user by username").
		RequestBuilder(
			cute.WithURI(url+"/"+newUsername),
			cute.WithMethod(http.MethodGet),
		).
		AssertBody(
			// TODO: разобраться почему вот так не работает:
			//json.Equal("$.id", 123456789),
			json.Equal("$.username", "updating_username"),
			json.Equal("$.firstName", "updating_firstName"),
			json.Equal("$.lastName", "updating_lastName"),
			json.Equal("$.email", "test@test.com"),
			json.Equal("$.password", "Strong_Password"),
			json.Equal("$.phone", "+1234567890"),
			json.Equal("$.userStatus", 1),
		).
		NextTest().
		//TODO: подумать как вынести это из шага в postcondition
		CreateStep("Delete user").
		RequestBuilder(
			cute.WithURI(url+"/"+newUsername),
			cute.WithMethod(http.MethodDelete),
		).
		AssertBody(
			json.Equal("$.code", 200),
			json.Equal("$.type", "unknown"),
			json.Equal("$.message", newUsername),
		).
		ExecuteTest(context.Background(), t)
}
