// Package lookups provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package lookups

const (
	AccountSid_authTokenScopes = "accountSid_authToken.Scopes"
)

// LookupsV1PhoneNumber defines model for lookups.v1.phone_number.
type LookupsV1PhoneNumber struct {
	// A JSON string with the results of the Add-ons you specified
	AddOns *map[string]interface{} `json:"add_ons"`

	// The name of the phone number's owner
	CallerName *map[string]interface{} `json:"caller_name"`

	// The telecom company that provides the phone number
	Carrier *map[string]interface{} `json:"carrier"`

	// The ISO country code for the phone number
	CountryCode *string `json:"country_code"`

	// The phone number, in national format
	NationalFormat *string `json:"national_format"`

	// The phone number in E.164 format
	PhoneNumber *string `json:"phone_number"`

	// The absolute URL of the resource
	Url *string `json:"url"`
}

// FetchPhoneNumberParams defines parameters for FetchPhoneNumber.
type FetchPhoneNumberParams struct {
	// The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the phone number to fetch. This is used to specify the country when the phone number is provided in a national format.
	CountryCode *string `json:"CountryCode,omitempty"`

	// The type of information to return. Can be: `carrier` or `caller-name`. The default is null.  Carrier information costs $0.005 per phone number looked up.  Caller Name information is currently available only in the US and costs $0.01 per phone number looked up.  To retrieve both types on information, specify this parameter twice; once with `carrier` and once with `caller-name` as the value.
	Type *[]string `json:"Type,omitempty"`

	// The `unique_name` of an Add-on you would like to invoke. Can be the `unique_name` of an Add-on that is installed on your account. You can specify multiple instances of this parameter to invoke multiple Add-ons. For more information about  Add-ons, see the [Add-ons documentation](https://www.twilio.com/docs/add-ons).
	AddOns *[]string `json:"AddOns,omitempty"`

	// Data specific to the add-on you would like to invoke. The content and format of this value depends on the add-on.
	AddOnsData *map[string]interface{} `json:"AddOnsData,omitempty"`
}
