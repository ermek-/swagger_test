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
	cute.NewTestBuilder().
		Epic("Swagger Petstore").
		Story("User").
		Feature("POST /user").
		Title("Create new user").
		Tags("regress", "smoke").
		CreateStep("Create new user").
		RequestBuilder(
			cute.WithURI(userUrl),
			cute.WithMethod(http.MethodPost),
			cute.WithHeadersKV("accept", "application/json"),
			cute.WithHeadersKV("Content-Type", "application/json"),
			cute.WithBody([]byte(userData)),
		).
		ExpectExecuteTimeout(5*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Equal("$.code", 200),
			json.Equal("$.type", "unknown"),
			json.Equal("$.message", id),
		).
		NextTest().
		CreateStep("Get user by username").
		AfterExecute(func(resp *http.Response, errs []error) error {
			cute.NewTestBuilder().
				Title("Delete user").
				Create().
				RequestBuilder(
					cute.WithURI(userUrl+"/"+username),
					cute.WithMethod(http.MethodDelete),
				).
				ExpectExecuteTimeout(5*time.Second).
				ExpectStatus(http.StatusOK).
				AssertBody(
					json.Equal("$.code", 200),
					json.Equal("$.type", "unknown"),
					json.Equal("$.message", username),
				).
				ExecuteTest(context.Background(), t)
			return nil
		}).
		RequestBuilder(
			cute.WithURI(userUrl+"/"+username),
			cute.WithMethod(http.MethodGet),
		).
		ExpectExecuteTimeout(5*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Equal("$.id", id),
			json.Equal("$.username", username),
			json.Equal("$.firstName", firstName),
			json.Equal("$.lastName", lastName),
			json.Equal("$.email", email),
			json.Equal("$.password", password),
			json.Equal("$.phone", phone),
			json.Equal("$.userStatus", userStatus),
		).
		ExecuteTest(context.Background(), t)
}

func Test_UpdateUser(t *testing.T) {
	newUsername := "updating_username"
	newFirstName := "updating_firstName"
	newLastName := "updating_lastName"
	newUserData := "{\n  \"id\": " + id + ",\n  \"username\": \"" + newUsername + "\",\n  \"firstName\": \"" + newFirstName + "\",\n  \"lastName\": \"" + newLastName + "\",\n  \"email\": \"" + email + "\",\n  \"password\": \"" + password + "\",\n  \"phone\": \"" + phone + "\",\n  \"userStatus\": " + userStatus + "\n}"
	cute.NewTestBuilder().
		Epic("Swagger Petstore").
		Story("User").
		Feature("PUT /user/{username}").
		Title("Update user").
		Tags("regress", "smoke").
		CreateStep("Update user").
		BeforeExecute(func(req *http.Request) error {
			cute.NewTestBuilder().
				Title("Create new user").
				Create().
				RequestBuilder(
					cute.WithURI(userUrl),
					cute.WithMethod(http.MethodPost),
					cute.WithHeadersKV("accept", "application/json"),
					cute.WithHeadersKV("Content-Type", "application/json"),
					cute.WithBody([]byte(userData)),
				).
				ExpectExecuteTimeout(5*time.Second).
				ExpectStatus(http.StatusOK).
				AssertBody(
					json.Equal("$.code", 200),
					json.Equal("$.type", "unknown"),
					json.Equal("$.message", id),
				).
				ExecuteTest(context.Background(), t)
			return nil
		}).
		RequestBuilder(
			cute.WithURI(userUrl+"/"+username),
			cute.WithMethod(http.MethodPut),
			cute.WithHeadersKV("accept", "application/json"),
			cute.WithHeadersKV("Content-Type", "application/json"),
			cute.WithBody([]byte(newUserData)),
		).
		ExpectExecuteTimeout(5*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Equal("$.code", 200),
			json.Equal("$.type", "unknown"),
			json.Equal("$.message", id),
		).
		NextTest().
		CreateStep("Get user by username").
		AfterExecute(func(resp *http.Response, errs []error) error {
			cute.NewTestBuilder().
				Title("Delete user").
				Create().
				RequestBuilder(
					cute.WithURI(userUrl+"/"+newUsername),
					cute.WithMethod(http.MethodDelete),
				).
				ExpectExecuteTimeout(5*time.Second).
				ExpectStatus(http.StatusOK).
				AssertBody(
					json.Equal("$.code", 200),
					json.Equal("$.type", "unknown"),
					json.Equal("$.message", newUsername),
				).
				ExecuteTest(context.Background(), t)
			return nil
		}).
		RequestBuilder(
			cute.WithURI(userUrl+"/"+newUsername),
			cute.WithMethod(http.MethodGet),
		).
		ExpectExecuteTimeout(5*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Equal("$.id", id),
			json.Equal("$.username", newUsername),
			json.Equal("$.firstName", newFirstName),
			json.Equal("$.lastName", newLastName),
			json.Equal("$.email", email),
			json.Equal("$.password", password),
			json.Equal("$.phone", phone),
			json.Equal("$.userStatus", userStatus),
		).
		ExecuteTest(context.Background(), t)
}

func Test_DeleteUser(t *testing.T) {
	cute.NewTestBuilder().
		Epic("Swagger Petstore").
		Story("User").
		Feature("DELETE /user/{username}").
		Title("Delete user").
		Tags("regress", "smoke").
		CreateStep("Delete user").
		BeforeExecute(func(req *http.Request) error {
			cute.NewTestBuilder().
				Title("Create new user").
				Create().
				RequestBuilder(
					cute.WithURI(userUrl),
					cute.WithMethod(http.MethodPost),
					cute.WithHeadersKV("accept", "application/json"),
					cute.WithHeadersKV("Content-Type", "application/json"),
					cute.WithBody([]byte(userData)),
				).
				ExpectExecuteTimeout(5*time.Second).
				ExpectStatus(http.StatusOK).
				AssertBody(
					json.Equal("$.code", 200),
					json.Equal("$.type", "unknown"),
					json.Equal("$.message", id),
				).
				ExecuteTest(context.Background(), t)
			return nil
		}).
		RequestBuilder(
			cute.WithURI(userUrl+"/"+username),
			cute.WithMethod(http.MethodDelete),
		).
		ExpectExecuteTimeout(5*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Equal("$.code", 200),
			json.Equal("$.type", "unknown"),
			json.Equal("$.message", username),
		).
		NextTest().
		CreateStep("Get user by username").
		RequestBuilder(
			cute.WithURI(userUrl+"/"+username),
			cute.WithMethod(http.MethodGet),
		).
		ExpectStatus(http.StatusNotFound).
		ExpectExecuteTimeout(5*time.Second).
		AssertBody(
			json.Equal("$.code", 1),
			json.Equal("$.type", "error"),
			json.Equal("$.message", "User not found"),
		).
		ExecuteTest(context.Background(), t)
}

func Test_LoginUser(t *testing.T) {
	cute.NewTestBuilder().
		Epic("Swagger Petstore").
		Story("User").
		Feature("GET /user/login").
		Title("Login user").
		Tags("regress", "smoke").
		Create().
		BeforeExecute(func(req *http.Request) error {
			cute.NewTestBuilder().
				Title("Create new user").
				Create().
				RequestBuilder(
					cute.WithURI(userUrl),
					cute.WithMethod(http.MethodPost),
					cute.WithHeadersKV("accept", "application/json"),
					cute.WithHeadersKV("Content-Type", "application/json"),
					cute.WithBody([]byte(userData)),
				).
				ExpectExecuteTimeout(5*time.Second).
				ExpectStatus(http.StatusOK).
				AssertBody(
					json.Equal("$.code", 200),
					json.Equal("$.type", "unknown"),
					json.Equal("$.message", id),
				).
				ExecuteTest(context.Background(), t)
			return nil
		}).
		AfterExecute(func(resp *http.Response, errs []error) error {
			cute.NewTestBuilder().
				Title("Delete user").
				Create().
				RequestBuilder(
					cute.WithURI(userUrl+"/"+username),
					cute.WithMethod(http.MethodDelete),
				).
				ExpectExecuteTimeout(5*time.Second).
				ExpectStatus(http.StatusOK).
				AssertBody(
					json.Equal("$.code", 200),
					json.Equal("$.type", "unknown"),
					json.Equal("$.message", username),
				).
				ExecuteTest(context.Background(), t)
			return nil
		}).
		RequestBuilder(
			cute.WithURI(userUrl+"/login"),
			cute.WithMethod(http.MethodGet),
			cute.WithQueryKV("username", username),
			cute.WithQueryKV("password", password),
		).
		ExpectStatus(http.StatusOK).
		ExpectExecuteTimeout(5*time.Second).
		AssertBody(
			json.Equal("$.code", 200),
			json.Equal("$.type", "unknown"),
			//TODO: подумать как сделать нормально:
			//json.Equal("$.message", "logged in user session:"+string(time.Now().Nanosecond())),
			json.NotEmpty("$.message"),
		).
		ExecuteTest(context.Background(), t)
}

func Test_LogoutUser(t *testing.T) {
	cute.NewTestBuilder().
		Epic("Swagger Petstore").
		Story("User").
		Feature("GET /user/logout").
		Title("Logout user").
		Tags("regress", "smoke").
		Create().
		BeforeExecute(func(req *http.Request) error {
			cute.NewTestBuilder().
				CreateStep("Create new user").
				RequestBuilder(
					cute.WithURI(userUrl),
					cute.WithMethod(http.MethodPost),
					cute.WithHeadersKV("accept", "application/json"),
					cute.WithHeadersKV("Content-Type", "application/json"),
					cute.WithBody([]byte(userData)),
				).
				ExpectExecuteTimeout(5*time.Second).
				ExpectStatus(http.StatusOK).
				AssertBody(
					json.Equal("$.code", 200),
					json.Equal("$.type", "unknown"),
					json.Equal("$.message", id),
				).
				NextTest().
				CreateStep("Login user").
				RequestBuilder(
					cute.WithURI(userUrl+"/login"),
					cute.WithMethod(http.MethodGet),
					cute.WithQueryKV("username", username),
					cute.WithQueryKV("password", password),
				).
				ExpectExecuteTimeout(5*time.Second).
				ExpectStatus(http.StatusOK).
				AssertBody(
					json.Equal("$.code", 200),
					json.Equal("$.type", "unknown"),
					json.NotEmpty("$.message"),
				).
				ExecuteTest(context.Background(), t)
			return nil
		}).
		AfterExecute(func(resp *http.Response, errs []error) error {
			cute.NewTestBuilder().
				Title("Delete user").
				Create().
				RequestBuilder(
					cute.WithURI(userUrl+"/"+username),
					cute.WithMethod(http.MethodDelete),
				).
				ExpectExecuteTimeout(5*time.Second).
				ExpectStatus(http.StatusOK).
				AssertBody(
					json.Equal("$.code", 200),
					json.Equal("$.type", "unknown"),
					json.Equal("$.message", username),
				).
				ExecuteTest(context.Background(), t)
			return nil
		}).
		RequestBuilder(
			cute.WithURI(userUrl+"/logout"),
			cute.WithMethod(http.MethodGet),
		).
		ExpectExecuteTimeout(5*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Equal("$.code", 200),
			json.Equal("$.type", "unknown"),
			json.Equal("$.message", "ok"),
		).
		ExecuteTest(context.Background(), t)
}
