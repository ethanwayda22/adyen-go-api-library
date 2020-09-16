/*
 * Adyen Checkout API
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including One-Click and 3D Secure), mobile wallets, and local payment methods (e.g. iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [Checkout documentation](https://docs.adyen.com/checkout).  ## Authentication Each Request to the Checkout API must be signed with an API key. For this, obtain an API Key from your Customer Area, as described in [How to get the API key](https://docs.adyen.com/user-management/how-to-get-the-api-key). Then set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: Your_Checkout_API_key\" \\ ... ``` Note that when going live, you need to generate a new API Key to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Checkout API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v51/payments ```
 *
 * API version: 51
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package common

import (
	"fmt"
	"net/http"
	"time"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes an oauth2.TokenSource as authentication for the Request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the Request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the Request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the Request
	ContextAPIKey = contextKey("apikey")
)

// BasicAuth provides basic http authentication to a Request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a Request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

type Environment string

const (
	TestEnv Environment = "TEST"
	LiveEnv Environment = "LIVE"
)

const (
	LibName    = "adyen-go-api-library"
	LibVersion = "3.0.0"
)

// Config stores the configuration of the API client
type Config struct {
	Username                string        `json:"username,omitempty"`
	Password                string        `json:"password,omitempty"`
	MerchantAccount         string        `json:"merchantAccount,omitempty"`
	Environment             Environment   `json:"environment,omitempty"`
	Endpoint                string        `json:"endpoint,omitempty"`
	MarketPayEndpoint       string        `json:"marketPayEndpoint,omitempty"`
	ApiKey                  string        `json:"apiKey,omitempty"`
	ConnectionTimeoutMillis time.Duration `json:"connectionTimeoutMillis,omitempty"`
	CertificatePath         string        `json:"certificatePath,omitempty"`

	//Checkout Specific
	CheckoutEndpoint string `json:"checkoutEndpoint,omitempty"`

	//Terminal API Specific
	TerminalApiCloudEndpoint string `json:"terminalApiCloudEndpoint,omitempty"`
	TerminalApiLocalEndpoint string `json:"terminalApiLocalEndpoint,omitempty"`
	TerminalCertificatePath  string `json:"terminalCertificatePath,omitempty"`

	LiveEndpointURLPrefix string            `json:"liveEndpointURLPrefix,omitempty"`
	DefaultHeader         map[string]string `json:"defaultHeader,omitempty"`
	Debug                 bool              `json:"debug,omitempty"`
	UserAgent             string            `json:"userAgent,omitempty"`
	HTTPClient            *http.Client
}

func (c *Config) GetCheckoutEndpoint() (string, error) {
	if c.CheckoutEndpoint == "" {
		message := "Please provide your unique live url prefix on the SetEnvironment() call on the APIClient or provide checkoutEndpoint in your config object."
		return "", fmt.Errorf(message)
	}
	return c.CheckoutEndpoint, nil
}
