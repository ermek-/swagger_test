package main

import "time"

const host = "https://petstore.swagger.io"

var (
	userUrl  = host + "/v2/user"
	petUrl   = host + "/v2/pet/"
	orderUrl = host + "/v2/store/order"
)

var (
	petId     = "123"
	petName   = "pet"
	photoUrls = "https://google.com/pet"
)

var (
	id         = "123"
	username   = "testing_user"
	firstName  = "test"
	lastName   = "test"
	email      = "test@test.com"
	password   = "Strong_Password"
	phone      = "+1234567890"
	userStatus = "1"
)
var (
	quantity       = "1"
	shipDate       = time.Now().Format("2006-01-02T15:04:05.000Z")
	shipDateFormat = time.Now().Format("2006-01-02T15:04:05.000+0000")
	status         = "placed"
)

var (
	userData  = "{\n  \"id\": " + id + ",\n  \"username\": \"" + username + "\",\n  \"firstName\": \"" + firstName + "\",\n  \"lastName\": \"" + lastName + "\",\n  \"email\": \"" + email + "\",\n  \"password\": \"" + password + "\",\n  \"phone\": \"" + phone + "\",\n  \"userStatus\": " + userStatus + "\n}"
	petData   = "{\n  \"name\": \"" + petName + "\",\n\"photoUrls\": [\"" + photoUrls + "\"]\n}"
	orderData = "{\n  \"id\": " + id + ",\n  \"petId\": " + petId + ",\n  \"quantity\": " + quantity + ",\n  \"shipDate\": \"" + shipDate + "\",\n \"status\": \"" + status + "\",\n  \"complete\": true\n}"
)
