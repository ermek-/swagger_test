package main

import (
	"context"
	"github.com/ozontech/cute"
	"github.com/ozontech/cute/asserts/json"
	"net/http"
	"testing"
	"time"
)

var url = host + "/v2/user"
var id = "123"
var username = "testing_user"
var firstName = "test"
var lastName = "test"
var email = "test@test.com"
var password = "Strong_Password"
var phone = "+1234567890"
var userStatus = "1"
var userData = "{\n  \"id\": " + id + ",\n  \"username\": \"" + username + "\",\n  \"firstName\": \"" + firstName + "\",\n  \"lastName\": \"" + lastName + "\",\n  \"email\": \"" + email + "\",\n  \"password\": \"" + password + "\",\n  \"phone\": \"" + phone + "\",\n  \"userStatus\": " + userStatus + "\n}"

func Test_CreateUser(t *testing.T) {
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
			json.Equal("$.message", id),
		).
		NextTest().
		CreateStep("Get user by username").
		RequestBuilder(
			cute.WithURI(url+"/"+username),
			cute.WithMethod(http.MethodGet),
		).
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
					json.Equal("$.message", id),
				).
				ExecuteTest(context.Background(), t)
			return nil
		}).
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
			json.Equal("$.message", id),
		).
		NextTest().
		CreateStep("Get user by username").
		AfterExecute(func(resp *http.Response, errs []error) error {
			cute.NewTestBuilder().
				Title("Delete user").
				Create().
				RequestBuilder(
					cute.WithURI(url+"/"+newUsername),
					cute.WithMethod(http.MethodDelete),
				).
				ExpectExecuteTimeout(10*time.Second).
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
			cute.WithURI(url+"/"+newUsername),
			cute.WithMethod(http.MethodGet),
		).
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
					cute.WithURI(url+"/"+username),
					cute.WithMethod(http.MethodDelete),
				).
				ExpectExecuteTimeout(10*time.Second).
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
			cute.WithURI(url+"/login"),
			cute.WithMethod(http.MethodGet),
			cute.WithQueryKV("username", username),
			cute.WithQueryKV("password", password),
		).
		AssertBody(
			json.Equal("$.code", 200),
			json.Equal("$.type", "unknown"),
			//TODO: подумать как сделать нормально:
			//json.Equal("$.message", "logged in user session:"+string(time.Now().Nanosecond())),
			json.NotEmpty("$.message"),
		).
		//NextTest().
		////TODO: подумать как вынести это из шага в postcondition
		//CreateStep("Logs out current logged in user session").
		//RequestBuilder(
		//	cute.WithURI(url+"/logout"),
		//	cute.WithMethod(http.MethodGet),
		//).
		//AssertBody(
		//	json.Equal("$.code", 200),
		//	json.Equal("$.type", "unknown"),
		//	json.Equal("$.message", "ok"),
		//).
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
					json.Equal("$.message", id),
				).
				NextTest().
				CreateStep("Login user").
				RequestBuilder(
					cute.WithURI(url+"/login"),
					cute.WithMethod(http.MethodGet),
					cute.WithQueryKV("username", username),
					cute.WithQueryKV("password", password),
				).
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
					cute.WithURI(url+"/"+username),
					cute.WithMethod(http.MethodDelete),
				).
				ExpectExecuteTimeout(10*time.Second).
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
			cute.WithURI(url+"/logout"),
			cute.WithMethod(http.MethodGet),
		).
		AssertBody(
			json.Equal("$.code", 200),
			json.Equal("$.type", "unknown"),
			json.Equal("$.message", "ok"),
		).
		ExecuteTest(context.Background(), t)
}
