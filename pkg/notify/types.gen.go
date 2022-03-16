// Package notify provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package notify

import (
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

const (
	AccountSid_authTokenScopes = "accountSid_authToken.Scopes"
)

// Defines values for NotifyV1CredentialType.
const (
	NotifyV1CredentialTypeApn NotifyV1CredentialType = "apn"
	NotifyV1CredentialTypeFcm NotifyV1CredentialType = "fcm"
	NotifyV1CredentialTypeGcm NotifyV1CredentialType = "gcm"
)

// Defines values for NotifyV1ServiceNotificationPriority.
const (
	NotifyV1ServiceNotificationPriorityHigh NotifyV1ServiceNotificationPriority = "high"
	NotifyV1ServiceNotificationPriorityLow  NotifyV1ServiceNotificationPriority = "low"
)

// NotifyV1Credential defines model for notify.v1.credential.
type NotifyV1Credential struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name"`

	// [APN only] Whether to send the credential to sandbox APNs
	Sandbox *string `json:"sandbox"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The Credential type, one of `gcm`, `fcm`, or `apn`
	Type *NotifyV1CredentialType `json:"type"`

	// The absolute URL of the Credential resource
	Url *string `json:"url"`
}

// The Credential type, one of `gcm`, `fcm`, or `apn`
type NotifyV1CredentialType string

// NotifyV1Service defines model for notify.v1.service.
type NotifyV1Service struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// Deprecated
	AlexaSkillId *string `json:"alexa_skill_id"`

	// The SID of the Credential to use for APN Bindings
	ApnCredentialSid *string `json:"apn_credential_sid"`

	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// Deprecated
	DefaultAlexaNotificationProtocolVersion *string `json:"default_alexa_notification_protocol_version"`

	// The protocol version to use for sending APNS notifications
	DefaultApnNotificationProtocolVersion *string `json:"default_apn_notification_protocol_version"`

	// The protocol version to use for sending FCM notifications
	DefaultFcmNotificationProtocolVersion *string `json:"default_fcm_notification_protocol_version"`

	// The protocol version to use for sending GCM notifications
	DefaultGcmNotificationProtocolVersion *string `json:"default_gcm_notification_protocol_version"`

	// Enable delivery callbacks
	DeliveryCallbackEnabled *bool `json:"delivery_callback_enabled"`

	// Webhook URL
	DeliveryCallbackUrl *string `json:"delivery_callback_url"`

	// Deprecated
	FacebookMessengerPageId *string `json:"facebook_messenger_page_id"`

	// The SID of the Credential to use for FCM Bindings
	FcmCredentialSid *string `json:"fcm_credential_sid"`

	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name"`

	// The SID of the Credential to use for GCM Bindings
	GcmCredentialSid *string `json:"gcm_credential_sid"`

	// The URLs of the resources related to the service
	Links *map[string]interface{} `json:"links"`

	// Whether to log notifications
	LogEnabled *bool `json:"log_enabled"`

	// The SID of the Messaging Service to use for SMS Bindings
	MessagingServiceSid *string `json:"messaging_service_sid"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The absolute URL of the Service resource
	Url *string `json:"url"`
}

// NotifyV1ServiceBinding defines model for notify.v1.service.binding.
type NotifyV1ServiceBinding struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// The channel-specific address
	Address *string `json:"address"`

	// The type of the Binding
	BindingType *string `json:"binding_type"`

	// The SID of the Credential resource to be used to send notifications to this Binding
	CredentialSid *string `json:"credential_sid"`

	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// Deprecated
	Endpoint *string `json:"endpoint"`

	// The `identity` value that identifies the new resource's User
	Identity *string `json:"identity"`

	// The URLs of related resources
	Links *map[string]interface{} `json:"links"`

	// The protocol version to use to send the notification
	NotificationProtocolVersion *string `json:"notification_protocol_version"`

	// The SID of the Service that the resource is associated with
	ServiceSid *string `json:"service_sid"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The list of tags associated with this Binding
	Tags *[]string `json:"tags"`

	// The absolute URL of the Binding resource
	Url *string `json:"url"`
}

// NotifyV1ServiceNotification defines model for notify.v1.service.notification.
type NotifyV1ServiceNotification struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// The actions to display for the notification
	Action *string `json:"action"`

	// Deprecated
	Alexa *map[string]interface{} `json:"alexa"`

	// The APNS-specific payload that overrides corresponding attributes in a generic payload for APNS Bindings
	Apn *map[string]interface{} `json:"apn"`

	// The notification body text
	Body *string `json:"body"`

	// The custom key-value pairs of the notification's payload
	Data *map[string]interface{} `json:"data"`

	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// Deprecated
	FacebookMessenger *map[string]interface{} `json:"facebook_messenger"`

	// The FCM-specific payload that overrides corresponding attributes in generic payload for FCM Bindings
	Fcm *map[string]interface{} `json:"fcm"`

	// The GCM-specific payload that overrides corresponding attributes in generic payload for GCM Bindings
	Gcm *map[string]interface{} `json:"gcm"`

	// The list of identity values of the Users to notify
	Identities *[]string `json:"identities"`

	// The priority of the notification
	Priority *NotifyV1ServiceNotificationPriority `json:"priority"`

	// The list of Segments to notify
	Segments *[]string `json:"segments"`

	// The SID of the Service that the resource is associated with
	ServiceSid *string `json:"service_sid"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The SMS-specific payload that overrides corresponding attributes in generic payload for SMS Bindings
	Sms *map[string]interface{} `json:"sms"`

	// The name of the sound to be played for the notification
	Sound *string `json:"sound"`

	// The tags that select the Bindings to notify
	Tags *[]string `json:"tags"`

	// The notification title
	Title *string `json:"title"`

	// How long, in seconds, the notification is valid
	Ttl *int `json:"ttl"`
}

// The priority of the notification
type NotifyV1ServiceNotificationPriority string

// ListCredentialParams defines parameters for ListCredential.
type ListCredentialParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListServiceParams defines parameters for ListService.
type ListServiceParams struct {
	// The string that identifies the Service resources to read.
	FriendlyName *string `json:"FriendlyName,omitempty"`

	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListBindingParams defines parameters for ListBinding.
type ListBindingParams struct {
	// Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`.
	StartDate *openapi_types.Date `json:"StartDate,omitempty"`

	// Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.
	EndDate *openapi_types.Date `json:"EndDate,omitempty"`

	// The [User](https://www.twilio.com/docs/chat/rest/user-resource)'s `identity` value of the resources to read.
	Identity *[]string `json:"Identity,omitempty"`

	// Only list Bindings that have all of the specified Tags. The following implicit tags are available: `all`, `apn`, `fcm`, `gcm`, `sms`, `facebook-messenger`. Up to 5 tags are allowed.
	Tag *[]string `json:"Tag,omitempty"`

	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}
