package gogpt

import (
	"net/http"
)

const (
	apiURLv1                       = "https://api.openai.com/v1"
	defaultEmptyMessagesLimit uint = 300
)

// ClientConfig is a configuration of a client.
type ClientConfig struct {
	AuthToken string

	HTTPClient *http.Client

	BaseURL string
	OrgID   string

	EmptyMessagesLimit uint
}

func DefaultConfig(authToken string) ClientConfig {
	return ClientConfig{
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				// set proxy from environment variables
				Proxy: http.ProxyFromEnvironment,
			},
		},
		BaseURL:   apiURLv1,
		OrgID:     "",
		AuthToken: authToken,

		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}
}
