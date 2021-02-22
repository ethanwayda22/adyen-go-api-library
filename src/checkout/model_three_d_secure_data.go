/*
 * Adyen Checkout API
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including 3D Secure), mobile wallets, and local payment methods (for example, iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [Checkout documentation](https://docs.adyen.com/checkout).  ## Authentication Each request to the Checkout API must be signed with an API key. For this, obtain an API Key from your Customer Area, as described in [How to get the API key](https://docs.adyen.com/user-management/how-to-get-the-api-key). Then set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: Your_Checkout_API_key\" \\ ... ``` Note that when going live, you need to generate a new API Key to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Checkout API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v67/payments ```
 *
 * API version: 65
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package checkout
// ThreeDSecureData struct for ThreeDSecureData
type ThreeDSecureData struct {
	// In 3D Secure 1, the authentication response if the shopper was redirected.  In 3D Secure 2, this is the `transStatus` from the challenge result. If the transaction was frictionless, omit this parameter.
	AuthenticationResponse string `json:"authenticationResponse,omitempty"`
	// The cardholder authentication value (base64 encoded, 20 bytes in a decoded form).
	Cavv string `json:"cavv,omitempty"`
	// The CAVV algorithm used. Include this only for 3D Secure 1.
	CavvAlgorithm string `json:"cavvAlgorithm,omitempty"`
	// In 3D Secure 1, this is the enrollment response from the 3D directory server.  In 3D Secure 2, this is the `transStatus` from the `ARes`.
	DirectoryResponse string `json:"directoryResponse,omitempty"`
	// Supported for 3D Secure 2. The unique transaction identifier assigned by the Directory Server (DS) to identify a single transaction.
	DsTransID string `json:"dsTransID,omitempty"`
	// The electronic commerce indicator.
	Eci string `json:"eci,omitempty"`
	// The version of the 3D Secure protocol.
	ThreeDSVersion string `json:"threeDSVersion,omitempty"`
	// Supported for 3D Secure 1. The transaction identifier (Base64-encoded, 20 bytes in a decoded form).
	Xid string `json:"xid,omitempty"`
}
