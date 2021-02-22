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
// AdditionalDataCarRental struct for AdditionalDataCarRental
type AdditionalDataCarRental struct {
	// Pick-up date. * Date format: `yyyyMMdd`
	CarRentalCheckOutDate string `json:"carRental.checkOutDate,omitempty"`
	// The customer service phone number of the car rental company. * Format: Alphanumeric * maxLength: 17
	CarRentalCustomerServiceTollFreeNumber string `json:"carRental.customerServiceTollFreeNumber,omitempty"`
	// Number of days for which the car is being rented. * Format: Numeric * maxLength: 19
	CarRentalDaysRented string `json:"carRental.daysRented,omitempty"`
	// Any fuel charges associated with the rental. * Format: Numeric * maxLength: 12
	CarRentalFuelCharges string `json:"carRental.fuelCharges,omitempty"`
	// Any insurance charges associated with the rental. * Format: Numeric * maxLength: 12
	CarRentalInsuranceCharges string `json:"carRental.insuranceCharges,omitempty"`
	// The city from which the car is rented. * Format: Alphanumeric * maxLength: 18
	CarRentalLocationCity string `json:"carRental.locationCity,omitempty"`
	// The country from which the car is rented. * Format: Alphanumeric * maxLength: 2
	CarRentalLocationCountry string `json:"carRental.locationCountry,omitempty"`
	// The state or province from where the car is rented. * Format: Alphanumeric * maxLength: 3
	CarRentalLocationStateProvince string `json:"carRental.locationStateProvince,omitempty"`
	// Indicates if the customer was a \"no-show\" (neither keeps nor cancels their booking). * Y - Customer was a no show. * N - Not applicable.
	CarRentalNoShowIndicator string `json:"carRental.noShowIndicator,omitempty"`
	// Charge associated with not returning a vehicle to the original rental location.
	CarRentalOneWayDropOffCharges string `json:"carRental.oneWayDropOffCharges,omitempty"`
	// Daily rental rate. * Format: Alphanumeric * maxLength: 12
	CarRentalRate string `json:"carRental.rate,omitempty"`
	// Specifies whether the given rate is applied daily or weekly. * D - Daily rate. * W - Weekly rate.
	CarRentalRateIndicator string `json:"carRental.rateIndicator,omitempty"`
	// The rental agreement number associated with this car rental. * Format: Alphanumeric * maxLength: 9
	CarRentalRentalAgreementNumber string `json:"carRental.rentalAgreementNumber,omitempty"`
	// Daily rental rate. * Format: Alphanumeric * maxLength: 12
	CarRentalRentalClassId string `json:"carRental.rentalClassId,omitempty"`
	// The name of the person renting the car. * Format: Alphanumeric * maxLength: 26
	CarRentalRenterName string `json:"carRental.renterName,omitempty"`
	// The city where the car must be returned. * Format: Alphanumeric * maxLength: 18
	CarRentalReturnCity string `json:"carRental.returnCity,omitempty"`
	// The country where the car must be returned. * Format: Alphanumeric * maxLength: 2
	CarRentalReturnCountry string `json:"carRental.returnCountry,omitempty"`
	// The last date to return the car by. * Date format: `yyyyMMdd`
	CarRentalReturnDate string `json:"carRental.returnDate,omitempty"`
	// Agency code, phone number, or address abbreviation * Format: Alphanumeric * maxLength: 10
	CarRentalReturnLocationId string `json:"carRental.returnLocationId,omitempty"`
	// The state or province where the car must be returned. * Format: Alphanumeric * maxLength: 3
	CarRentalReturnStateProvince string `json:"carRental.returnStateProvince,omitempty"`
	// Indicates whether the goods or services were tax-exempt, or tax was not collected.  Values: * Y - Goods or services were tax exempt * N - Tax was not collected
	CarRentalTaxExemptIndicator string `json:"carRental.taxExemptIndicator,omitempty"`
	// Number of nights.  This should be included in the auth message. * Format: Numeric * maxLength: 2
	TravelEntertainmentAuthDataDuration string `json:"travelEntertainmentAuthData.duration,omitempty"`
	// Indicates what market-specific dataset will be submitted or is being submitted. Value should be \"A\" for Car rental. This should be included in the auth message. * Format: Alphanumeric * maxLength: 1
	TravelEntertainmentAuthDataMarket string `json:"travelEntertainmentAuthData.market,omitempty"`
}
