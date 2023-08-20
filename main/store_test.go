package main

import (
	"context"
	"github.com/ozontech/cute"
	"github.com/ozontech/cute/asserts/json"
	"net/http"
	"testing"
	"time"
)

func Test_CreateOrder(t *testing.T) {
	cute.NewTestBuilder().
		Epic("Swagger Petstore").
		Story("Store").
		Feature("POST /store/order").
		Title("Create new order").
		Tags("regress", "smoke").
		CreateStep("Create new order").
		RequestBuilder(
			cute.WithURI(orderUrl),
			cute.WithMethod(http.MethodPost),
			cute.WithHeadersKV("accept", "application/json"),
			cute.WithHeadersKV("Content-Type", "application/json"),
			cute.WithBody([]byte(orderData)),
		).
		ExpectExecuteTimeout(5*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Equal("$.id", id),
			json.Equal("$.petId", petId),
			json.Equal("$.quantity", quantity),
			json.Equal("$.shipDate", shipDateFormat),
			json.Equal("$.status", status),
			json.Equal("$.complete", true),
		).
		NextTest().
		CreateStep("Get order by id").
		AfterExecute(func(resp *http.Response, errs []error) error {
			cute.NewTestBuilder().
				Title("Delete order").
				Create().
				RequestBuilder(
					cute.WithURI(orderUrl+"/"+id),
					cute.WithMethod(http.MethodDelete),
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
			cute.WithURI(orderUrl+"/"+id),
			cute.WithMethod(http.MethodGet),
		).
		ExpectExecuteTimeout(5*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Equal("$.id", id),
			json.Equal("$.petId", petId),
			json.Equal("$.quantity", quantity),
			json.Equal("$.shipDate", shipDateFormat),
			json.Equal("$.status", status),
			json.Equal("$.complete", true),
		).
		ExecuteTest(context.Background(), t)
}

func Test_GetOrder(t *testing.T) {
	cute.NewTestBuilder().
		Epic("Swagger Petstore").
		Story("Store").
		Feature("GET /store/order/{orderId}").
		Title("Get order by orderId").
		Tags("regress", "smoke").
		CreateStep("Get order").
		BeforeExecute(func(req *http.Request) error {
			cute.NewTestBuilder().
				Title("Create new order").
				Create().
				RequestBuilder(
					cute.WithURI(orderUrl),
					cute.WithMethod(http.MethodPost),
					cute.WithHeadersKV("accept", "application/json"),
					cute.WithHeadersKV("Content-Type", "application/json"),
					cute.WithBody([]byte(orderData)),
				).
				ExpectExecuteTimeout(5*time.Second).
				ExpectStatus(http.StatusOK).
				AssertBody(
					json.Equal("$.id", id),
					json.Equal("$.petId", petId),
					json.Equal("$.quantity", quantity),
					json.Equal("$.shipDate", shipDateFormat),
					json.Equal("$.status", status),
					json.Equal("$.complete", true),
				).
				ExecuteTest(context.Background(), t)
			return nil
		}).
		AfterExecute(func(resp *http.Response, errs []error) error {
			cute.NewTestBuilder().
				Title("Delete order").
				Create().
				RequestBuilder(
					cute.WithURI(orderUrl+"/"+id),
					cute.WithMethod(http.MethodDelete),
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
			cute.WithURI(orderUrl+"/"+id),
			cute.WithMethod(http.MethodGet),
		).
		ExpectExecuteTimeout(5*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Equal("$.id", id),
			json.Equal("$.petId", petId),
			json.Equal("$.quantity", quantity),
			json.Equal("$.shipDate", shipDateFormat),
			json.Equal("$.status", status),
			json.Equal("$.complete", true),
		).
		ExecuteTest(context.Background(), t)
}

func Test_DeleteOrder(t *testing.T) {
	cute.NewTestBuilder().
		Epic("Swagger Petstore").
		Story("Store").
		Feature("DELETE /store/order/{orderId}").
		Title("Delete order by orderId").
		Tags("regress", "smoke").
		CreateStep("Delete order").
		BeforeExecute(func(req *http.Request) error {
			cute.NewTestBuilder().
				Title("Create new order").
				Create().
				RequestBuilder(
					cute.WithURI(orderUrl),
					cute.WithMethod(http.MethodPost),
					cute.WithHeadersKV("accept", "application/json"),
					cute.WithHeadersKV("Content-Type", "application/json"),
					cute.WithBody([]byte(orderData)),
				).
				ExpectExecuteTimeout(5*time.Second).
				ExpectStatus(http.StatusOK).
				AssertBody(
					json.Equal("$.id", id),
					json.Equal("$.petId", petId),
					json.Equal("$.quantity", quantity),
					json.Equal("$.shipDate", shipDateFormat),
					json.Equal("$.status", status),
					json.Equal("$.complete", true),
				).
				ExecuteTest(context.Background(), t)
			return nil
		}).
		RequestBuilder(
			cute.WithURI(orderUrl+"/"+id),
			cute.WithMethod(http.MethodDelete),
		).
		ExpectExecuteTimeout(5*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Equal("$.code", 200),
			json.Equal("$.type", "unknown"),
			json.Equal("$.message", id),
		).
		ExecuteTest(context.Background(), t)
}

func Test_GetInventory(t *testing.T) {
	cute.NewTestBuilder().
		Epic("Swagger Petstore").
		Story("Store").
		Feature("GET /store/inventory").
		Title("Get inventory").
		Tags("regress", "smoke").
		CreateStep("Get inventory").
		RequestBuilder(
			cute.WithURI(host+"/v2/store/inventory"),
			cute.WithMethod(http.MethodGet),
		).
		ExpectExecuteTimeout(5*time.Second).
		ExpectStatus(http.StatusOK).
		ExpectJSONSchemaFile("file://./resources/inventory.json").
		ExecuteTest(context.Background(), t)
}
