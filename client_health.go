package meilisearch

import "net/http"

type clientHealth struct {
	client *Client
}

func newClientHealth(client *Client) clientHealth {
	return clientHealth{client: client}
}

func (c clientHealth) Get() error {
	req := internalRequest{
		endpoint:            "/health",
		method:              http.MethodGet,
		withRequest:         nil,
		withResponse:        nil,
		acceptedStatusCodes: []int{http.StatusNoContent},
		functionName:        "Get",
		apiName:             "Health",
	}

	return c.client.executeRequest(req)
}
