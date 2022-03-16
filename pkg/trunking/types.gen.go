// Package trunking provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package trunking

import (
	"time"
)

const (
	AccountSid_authTokenScopes = "accountSid_authToken.Scopes"
)

// Defines values for TrunkingV1TrunkDisasterRecoveryMethod.
const (
	TrunkingV1TrunkDisasterRecoveryMethodDELETE TrunkingV1TrunkDisasterRecoveryMethod = "DELETE"
	TrunkingV1TrunkDisasterRecoveryMethodGET    TrunkingV1TrunkDisasterRecoveryMethod = "GET"
	TrunkingV1TrunkDisasterRecoveryMethodHEAD   TrunkingV1TrunkDisasterRecoveryMethod = "HEAD"
	TrunkingV1TrunkDisasterRecoveryMethodPATCH  TrunkingV1TrunkDisasterRecoveryMethod = "PATCH"
	TrunkingV1TrunkDisasterRecoveryMethodPOST   TrunkingV1TrunkDisasterRecoveryMethod = "POST"
	TrunkingV1TrunkDisasterRecoveryMethodPUT    TrunkingV1TrunkDisasterRecoveryMethod = "PUT"
)

// Defines values for TrunkingV1TrunkTransferCallerId.
const (
	TrunkingV1TrunkTransferCallerIdFromTransferee TrunkingV1TrunkTransferCallerId = "from-transferee"
	TrunkingV1TrunkTransferCallerIdFromTransferor TrunkingV1TrunkTransferCallerId = "from-transferor"
)

// Defines values for TrunkingV1TrunkTransferMode.
const (
	TrunkingV1TrunkTransferModeDisableAll TrunkingV1TrunkTransferMode = "disable-all"
	TrunkingV1TrunkTransferModeEnableAll  TrunkingV1TrunkTransferMode = "enable-all"
	TrunkingV1TrunkTransferModeSipOnly    TrunkingV1TrunkTransferMode = "sip-only"
)

// Defines values for TrunkingV1TrunkPhoneNumberAddressRequirements.
const (
	TrunkingV1TrunkPhoneNumberAddressRequirementsAny     TrunkingV1TrunkPhoneNumberAddressRequirements = "any"
	TrunkingV1TrunkPhoneNumberAddressRequirementsForeign TrunkingV1TrunkPhoneNumberAddressRequirements = "foreign"
	TrunkingV1TrunkPhoneNumberAddressRequirementsLocal   TrunkingV1TrunkPhoneNumberAddressRequirements = "local"
	TrunkingV1TrunkPhoneNumberAddressRequirementsNone    TrunkingV1TrunkPhoneNumberAddressRequirements = "none"
)

// Defines values for TrunkingV1TrunkPhoneNumberSmsFallbackMethod.
const (
	TrunkingV1TrunkPhoneNumberSmsFallbackMethodDELETE TrunkingV1TrunkPhoneNumberSmsFallbackMethod = "DELETE"
	TrunkingV1TrunkPhoneNumberSmsFallbackMethodGET    TrunkingV1TrunkPhoneNumberSmsFallbackMethod = "GET"
	TrunkingV1TrunkPhoneNumberSmsFallbackMethodHEAD   TrunkingV1TrunkPhoneNumberSmsFallbackMethod = "HEAD"
	TrunkingV1TrunkPhoneNumberSmsFallbackMethodPATCH  TrunkingV1TrunkPhoneNumberSmsFallbackMethod = "PATCH"
	TrunkingV1TrunkPhoneNumberSmsFallbackMethodPOST   TrunkingV1TrunkPhoneNumberSmsFallbackMethod = "POST"
	TrunkingV1TrunkPhoneNumberSmsFallbackMethodPUT    TrunkingV1TrunkPhoneNumberSmsFallbackMethod = "PUT"
)

// Defines values for TrunkingV1TrunkPhoneNumberSmsMethod.
const (
	TrunkingV1TrunkPhoneNumberSmsMethodDELETE TrunkingV1TrunkPhoneNumberSmsMethod = "DELETE"
	TrunkingV1TrunkPhoneNumberSmsMethodGET    TrunkingV1TrunkPhoneNumberSmsMethod = "GET"
	TrunkingV1TrunkPhoneNumberSmsMethodHEAD   TrunkingV1TrunkPhoneNumberSmsMethod = "HEAD"
	TrunkingV1TrunkPhoneNumberSmsMethodPATCH  TrunkingV1TrunkPhoneNumberSmsMethod = "PATCH"
	TrunkingV1TrunkPhoneNumberSmsMethodPOST   TrunkingV1TrunkPhoneNumberSmsMethod = "POST"
	TrunkingV1TrunkPhoneNumberSmsMethodPUT    TrunkingV1TrunkPhoneNumberSmsMethod = "PUT"
)

// Defines values for TrunkingV1TrunkPhoneNumberStatusCallbackMethod.
const (
	TrunkingV1TrunkPhoneNumberStatusCallbackMethodDELETE TrunkingV1TrunkPhoneNumberStatusCallbackMethod = "DELETE"
	TrunkingV1TrunkPhoneNumberStatusCallbackMethodGET    TrunkingV1TrunkPhoneNumberStatusCallbackMethod = "GET"
	TrunkingV1TrunkPhoneNumberStatusCallbackMethodHEAD   TrunkingV1TrunkPhoneNumberStatusCallbackMethod = "HEAD"
	TrunkingV1TrunkPhoneNumberStatusCallbackMethodPATCH  TrunkingV1TrunkPhoneNumberStatusCallbackMethod = "PATCH"
	TrunkingV1TrunkPhoneNumberStatusCallbackMethodPOST   TrunkingV1TrunkPhoneNumberStatusCallbackMethod = "POST"
	TrunkingV1TrunkPhoneNumberStatusCallbackMethodPUT    TrunkingV1TrunkPhoneNumberStatusCallbackMethod = "PUT"
)

// Defines values for TrunkingV1TrunkPhoneNumberVoiceFallbackMethod.
const (
	TrunkingV1TrunkPhoneNumberVoiceFallbackMethodDELETE TrunkingV1TrunkPhoneNumberVoiceFallbackMethod = "DELETE"
	TrunkingV1TrunkPhoneNumberVoiceFallbackMethodGET    TrunkingV1TrunkPhoneNumberVoiceFallbackMethod = "GET"
	TrunkingV1TrunkPhoneNumberVoiceFallbackMethodHEAD   TrunkingV1TrunkPhoneNumberVoiceFallbackMethod = "HEAD"
	TrunkingV1TrunkPhoneNumberVoiceFallbackMethodPATCH  TrunkingV1TrunkPhoneNumberVoiceFallbackMethod = "PATCH"
	TrunkingV1TrunkPhoneNumberVoiceFallbackMethodPOST   TrunkingV1TrunkPhoneNumberVoiceFallbackMethod = "POST"
	TrunkingV1TrunkPhoneNumberVoiceFallbackMethodPUT    TrunkingV1TrunkPhoneNumberVoiceFallbackMethod = "PUT"
)

// Defines values for TrunkingV1TrunkPhoneNumberVoiceMethod.
const (
	TrunkingV1TrunkPhoneNumberVoiceMethodDELETE TrunkingV1TrunkPhoneNumberVoiceMethod = "DELETE"
	TrunkingV1TrunkPhoneNumberVoiceMethodGET    TrunkingV1TrunkPhoneNumberVoiceMethod = "GET"
	TrunkingV1TrunkPhoneNumberVoiceMethodHEAD   TrunkingV1TrunkPhoneNumberVoiceMethod = "HEAD"
	TrunkingV1TrunkPhoneNumberVoiceMethodPATCH  TrunkingV1TrunkPhoneNumberVoiceMethod = "PATCH"
	TrunkingV1TrunkPhoneNumberVoiceMethodPOST   TrunkingV1TrunkPhoneNumberVoiceMethod = "POST"
	TrunkingV1TrunkPhoneNumberVoiceMethodPUT    TrunkingV1TrunkPhoneNumberVoiceMethod = "PUT"
)

// Defines values for TrunkingV1TrunkRecordingMode.
const (
	TrunkingV1TrunkRecordingModeDoNotRecord           TrunkingV1TrunkRecordingMode = "do-not-record"
	TrunkingV1TrunkRecordingModeRecordFromAnswer      TrunkingV1TrunkRecordingMode = "record-from-answer"
	TrunkingV1TrunkRecordingModeRecordFromAnswerDual  TrunkingV1TrunkRecordingMode = "record-from-answer-dual"
	TrunkingV1TrunkRecordingModeRecordFromRinging     TrunkingV1TrunkRecordingMode = "record-from-ringing"
	TrunkingV1TrunkRecordingModeRecordFromRingingDual TrunkingV1TrunkRecordingMode = "record-from-ringing-dual"
)

// Defines values for TrunkingV1TrunkRecordingTrim.
const (
	TrunkingV1TrunkRecordingTrimDoNotTrim   TrunkingV1TrunkRecordingTrim = "do-not-trim"
	TrunkingV1TrunkRecordingTrimTrimSilence TrunkingV1TrunkRecordingTrim = "trim-silence"
)

// TrunkingV1Trunk defines model for trunking.v1.trunk.
type TrunkingV1Trunk struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// The types of authentication mapped to the domain
	AuthType *string `json:"auth_type"`

	// Reserved
	AuthTypeSet *[]string `json:"auth_type_set"`

	// Whether Caller ID Name (CNAM) lookup is enabled for the trunk
	CnamLookupEnabled *bool `json:"cnam_lookup_enabled"`

	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// The HTTP method we use to call the disaster_recovery_url
	DisasterRecoveryMethod *TrunkingV1TrunkDisasterRecoveryMethod `json:"disaster_recovery_method"`

	// The HTTP URL that we call if an error occurs while sending SIP traffic towards your configured Origination URL
	DisasterRecoveryUrl *string `json:"disaster_recovery_url"`

	// The unique address you reserve on Twilio to which you route your SIP traffic
	DomainName *string `json:"domain_name"`

	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name"`

	// The URLs of related resources
	Links *map[string]interface{} `json:"links"`

	// The recording settings for the trunk
	Recording *map[string]interface{} `json:"recording"`

	// Whether Secure Trunking is enabled for the trunk
	Secure *bool `json:"secure"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// Caller Id for transfer target
	TransferCallerId *TrunkingV1TrunkTransferCallerId `json:"transfer_caller_id"`

	// The call transfer settings for the trunk
	TransferMode *TrunkingV1TrunkTransferMode `json:"transfer_mode"`

	// The absolute URL of the resource
	Url *string `json:"url"`
}

// The HTTP method we use to call the disaster_recovery_url
type TrunkingV1TrunkDisasterRecoveryMethod string

// Caller Id for transfer target
type TrunkingV1TrunkTransferCallerId string

// The call transfer settings for the trunk
type TrunkingV1TrunkTransferMode string

// TrunkingV1TrunkCredentialList defines model for trunking.v1.trunk.credential_list.
type TrunkingV1TrunkCredentialList struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The SID of the Trunk the credential list in associated with
	TrunkSid *string `json:"trunk_sid"`

	// The absolute URL of the resource
	Url *string `json:"url"`
}

// TrunkingV1TrunkIpAccessControlList defines model for trunking.v1.trunk.ip_access_control_list.
type TrunkingV1TrunkIpAccessControlList struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The SID of the Trunk the resource is associated with
	TrunkSid *string `json:"trunk_sid"`

	// The absolute URL of the resource
	Url *string `json:"url"`
}

// TrunkingV1TrunkOriginationUrl defines model for trunking.v1.trunk.origination_url.
type TrunkingV1TrunkOriginationUrl struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// Whether the URL is enabled
	Enabled *bool `json:"enabled"`

	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name"`

	// The relative importance of the URI
	Priority *int `json:"priority"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The SIP address you want Twilio to route your Origination calls to
	SipUrl *string `json:"sip_url"`

	// The SID of the Trunk that owns the Origination URL
	TrunkSid *string `json:"trunk_sid"`

	// The absolute URL of the resource
	Url *string `json:"url"`

	// The value that determines the relative load the URI should receive compared to others with the same priority
	Weight *int `json:"weight"`
}

// TrunkingV1TrunkPhoneNumber defines model for trunking.v1.trunk.phone_number.
type TrunkingV1TrunkPhoneNumber struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// Whether the phone number requires an Address registered with Twilio
	AddressRequirements *TrunkingV1TrunkPhoneNumberAddressRequirements `json:"address_requirements"`

	// The API version used to start a new TwiML session
	ApiVersion *string `json:"api_version"`

	// Whether the phone number is new to the Twilio platform
	Beta *bool `json:"beta"`

	// Indicate if a phone can receive calls or messages
	Capabilities *map[string]interface{} `json:"capabilities"`

	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name"`

	// The URLs of related resources
	Links *map[string]interface{} `json:"links"`

	// The phone number in E.164 format
	PhoneNumber *string `json:"phone_number"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The SID of the application that handles SMS messages sent to the phone number
	SmsApplicationSid *string `json:"sms_application_sid"`

	// The HTTP method used with sms_fallback_url
	SmsFallbackMethod *TrunkingV1TrunkPhoneNumberSmsFallbackMethod `json:"sms_fallback_method"`

	// The URL that we call when an error occurs while retrieving or executing the TwiML
	SmsFallbackUrl *string `json:"sms_fallback_url"`

	// The HTTP method to use with sms_url
	SmsMethod *TrunkingV1TrunkPhoneNumberSmsMethod `json:"sms_method"`

	// The URL we call when the phone number receives an incoming SMS message
	SmsUrl *string `json:"sms_url"`

	// The URL to send status information to your application
	StatusCallback *string `json:"status_callback"`

	// The HTTP method we use to call status_callback
	StatusCallbackMethod *TrunkingV1TrunkPhoneNumberStatusCallbackMethod `json:"status_callback_method"`

	// The SID of the Trunk that handles calls to the phone number
	TrunkSid *string `json:"trunk_sid"`

	// The absolute URL of the resource
	Url *string `json:"url"`

	// The SID of the application that handles calls to the phone number
	VoiceApplicationSid *string `json:"voice_application_sid"`

	// Whether to lookup the caller's name
	VoiceCallerIdLookup *bool `json:"voice_caller_id_lookup"`

	// The HTTP method that we use to call voice_fallback_url
	VoiceFallbackMethod *TrunkingV1TrunkPhoneNumberVoiceFallbackMethod `json:"voice_fallback_method"`

	// The URL we call when an error occurs in TwiML
	VoiceFallbackUrl *string `json:"voice_fallback_url"`

	// The HTTP method used with the voice_url
	VoiceMethod *TrunkingV1TrunkPhoneNumberVoiceMethod `json:"voice_method"`

	// The URL we call when the phone number receives a call
	VoiceUrl *string `json:"voice_url"`
}

// Whether the phone number requires an Address registered with Twilio
type TrunkingV1TrunkPhoneNumberAddressRequirements string

// The HTTP method used with sms_fallback_url
type TrunkingV1TrunkPhoneNumberSmsFallbackMethod string

// The HTTP method to use with sms_url
type TrunkingV1TrunkPhoneNumberSmsMethod string

// The HTTP method we use to call status_callback
type TrunkingV1TrunkPhoneNumberStatusCallbackMethod string

// The HTTP method that we use to call voice_fallback_url
type TrunkingV1TrunkPhoneNumberVoiceFallbackMethod string

// The HTTP method used with the voice_url
type TrunkingV1TrunkPhoneNumberVoiceMethod string

// TrunkingV1TrunkRecording defines model for trunking.v1.trunk.recording.
type TrunkingV1TrunkRecording struct {
	// The recording mode for the trunk.
	Mode *TrunkingV1TrunkRecordingMode `json:"mode"`

	// The recording trim setting for the trunk.
	Trim *TrunkingV1TrunkRecordingTrim `json:"trim"`
}

// The recording mode for the trunk.
type TrunkingV1TrunkRecordingMode string

// The recording trim setting for the trunk.
type TrunkingV1TrunkRecordingTrim string

// ListTrunkParams defines parameters for ListTrunk.
type ListTrunkParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListCredentialListParams defines parameters for ListCredentialList.
type ListCredentialListParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListIpAccessControlListParams defines parameters for ListIpAccessControlList.
type ListIpAccessControlListParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListOriginationUrlParams defines parameters for ListOriginationUrl.
type ListOriginationUrlParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListPhoneNumberParams defines parameters for ListPhoneNumber.
type ListPhoneNumberParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}