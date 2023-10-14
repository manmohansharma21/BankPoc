package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) login(payload *KLoginPayload) (*KLoginRes, error) {
	// Pinging Keycloak, it takes form-data; we use "net/url" to parse form-data
	formData := url.Values{
		"client_id":     {payload.clientId},
		"client_secret": {payload.clientSecret},
		"grant_type":    {payload.grantType},
		"username":      {payload.username},
		"password":      {payload.password},
	}

	encodedFormData := formData.Encode()
	url := "http://localhost:8080/realms/bankpoc/protocol/openid-connect/token"

	req, err := http.NewRequest("POST", url, strings.NewReader(encodedFormData)) // To ping Keycloak
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req) // Do sends the HTTP request by making a request using the HTTP client.
	if err != nil {                   // The response (resp) and any potential error (err) are returned.
		return nil, err
	}
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("something went wrong while talking to Keycloak login")
	}

	kLoginRes := &KLoginRes{}
	json.NewDecoder(resp.Body).Decode(kLoginRes)

	return kLoginRes, nil
}

func (c *Client) introspect(payload *KIntrospectPayload) (*introspectRes, error) {
	// Pinging Keycloak, it takes form-data; we use "net/url" to parse form-data
	formData := url.Values{
		"client_id":     {payload.clientId},
		"client_secret": {payload.clientSecret},
		"token":         {payload.token},
	}

	encodedFormData := formData.Encode()
	introspectUrl := "http://localhost:8080/realms/bankpoc/protocol/openid-connect/token/introspect"

	req, err := http.NewRequest("POST", introspectUrl, strings.NewReader(encodedFormData)) // To ping Keycloak
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req) // Do sends the HTTP request by making a request using the HTTP client.
	if err != nil {                   // The response (resp) and any potential error (err) are returned.
		return nil, err
	}
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("something went wrong while talking to Keycloak login")
	}

	// Read the response boy
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println((string)(body))

	introspectRes := &introspectRes{}
	json.NewDecoder(resp.Body).Decode(introspectRes)
	return introspectRes, nil
}

/*
 sending a request to the Keycloak token endpoint for authentication.
 Authentication typically involves sending sensitive information like a username and password,
 which should not be exposed in the URL. Therefore, a "POST" request is the appropriate choice,
 as it allows you to send this sensitive data in the request body.

You determine whether to use "POST" or "GET" based on the API documentation or the specific requirements
of the endpoint you are accessing. In the case of authentication or data submission, "POST" is commonly
used for security and data privacy reasons.
*/
