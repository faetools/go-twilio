// Package proxy provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package proxy

import (
	"time"
)

const (
	AccountSid_authTokenScopes = "accountSid_authToken.Scopes"
)

// Defines values for ProxyV1ServiceGeoMatchLevel.
const (
	ProxyV1ServiceGeoMatchLevelAreaCode ProxyV1ServiceGeoMatchLevel = "area-code"
	ProxyV1ServiceGeoMatchLevelCountry  ProxyV1ServiceGeoMatchLevel = "country"
	ProxyV1ServiceGeoMatchLevelOverlay  ProxyV1ServiceGeoMatchLevel = "overlay"
	ProxyV1ServiceGeoMatchLevelRadius   ProxyV1ServiceGeoMatchLevel = "radius"
)

// Defines values for ProxyV1ServiceNumberSelectionBehavior.
const (
	ProxyV1ServiceNumberSelectionBehaviorAvoidSticky  ProxyV1ServiceNumberSelectionBehavior = "avoid-sticky"
	ProxyV1ServiceNumberSelectionBehaviorPreferSticky ProxyV1ServiceNumberSelectionBehavior = "prefer-sticky"
)

// Defines values for ProxyV1ServiceSessionMode.
const (
	ProxyV1ServiceSessionModeMessageOnly     ProxyV1ServiceSessionMode = "message-only"
	ProxyV1ServiceSessionModeVoiceAndMessage ProxyV1ServiceSessionMode = "voice-and-message"
	ProxyV1ServiceSessionModeVoiceOnly       ProxyV1ServiceSessionMode = "voice-only"
)

// Defines values for ProxyV1ServiceSessionStatus.
const (
	ProxyV1ServiceSessionStatusClosed     ProxyV1ServiceSessionStatus = "closed"
	ProxyV1ServiceSessionStatusFailed     ProxyV1ServiceSessionStatus = "failed"
	ProxyV1ServiceSessionStatusInProgress ProxyV1ServiceSessionStatus = "in-progress"
	ProxyV1ServiceSessionStatusOpen       ProxyV1ServiceSessionStatus = "open"
	ProxyV1ServiceSessionStatusUnknown    ProxyV1ServiceSessionStatus = "unknown"
)

// Defines values for ProxyV1ServiceSessionInteractionInboundResourceStatus.
const (
	ProxyV1ServiceSessionInteractionInboundResourceStatusAccepted        ProxyV1ServiceSessionInteractionInboundResourceStatus = "accepted"
	ProxyV1ServiceSessionInteractionInboundResourceStatusAnswered        ProxyV1ServiceSessionInteractionInboundResourceStatus = "answered"
	ProxyV1ServiceSessionInteractionInboundResourceStatusBusy            ProxyV1ServiceSessionInteractionInboundResourceStatus = "busy"
	ProxyV1ServiceSessionInteractionInboundResourceStatusCanceled        ProxyV1ServiceSessionInteractionInboundResourceStatus = "canceled"
	ProxyV1ServiceSessionInteractionInboundResourceStatusCompleted       ProxyV1ServiceSessionInteractionInboundResourceStatus = "completed"
	ProxyV1ServiceSessionInteractionInboundResourceStatusDeleted         ProxyV1ServiceSessionInteractionInboundResourceStatus = "deleted"
	ProxyV1ServiceSessionInteractionInboundResourceStatusDelivered       ProxyV1ServiceSessionInteractionInboundResourceStatus = "delivered"
	ProxyV1ServiceSessionInteractionInboundResourceStatusDeliveryUnknown ProxyV1ServiceSessionInteractionInboundResourceStatus = "delivery-unknown"
	ProxyV1ServiceSessionInteractionInboundResourceStatusFailed          ProxyV1ServiceSessionInteractionInboundResourceStatus = "failed"
	ProxyV1ServiceSessionInteractionInboundResourceStatusInProgress      ProxyV1ServiceSessionInteractionInboundResourceStatus = "in-progress"
	ProxyV1ServiceSessionInteractionInboundResourceStatusInitiated       ProxyV1ServiceSessionInteractionInboundResourceStatus = "initiated"
	ProxyV1ServiceSessionInteractionInboundResourceStatusNoAnswer        ProxyV1ServiceSessionInteractionInboundResourceStatus = "no-answer"
	ProxyV1ServiceSessionInteractionInboundResourceStatusQueued          ProxyV1ServiceSessionInteractionInboundResourceStatus = "queued"
	ProxyV1ServiceSessionInteractionInboundResourceStatusReceived        ProxyV1ServiceSessionInteractionInboundResourceStatus = "received"
	ProxyV1ServiceSessionInteractionInboundResourceStatusReceiving       ProxyV1ServiceSessionInteractionInboundResourceStatus = "receiving"
	ProxyV1ServiceSessionInteractionInboundResourceStatusRinging         ProxyV1ServiceSessionInteractionInboundResourceStatus = "ringing"
	ProxyV1ServiceSessionInteractionInboundResourceStatusScheduled       ProxyV1ServiceSessionInteractionInboundResourceStatus = "scheduled"
	ProxyV1ServiceSessionInteractionInboundResourceStatusSending         ProxyV1ServiceSessionInteractionInboundResourceStatus = "sending"
	ProxyV1ServiceSessionInteractionInboundResourceStatusSent            ProxyV1ServiceSessionInteractionInboundResourceStatus = "sent"
	ProxyV1ServiceSessionInteractionInboundResourceStatusUndelivered     ProxyV1ServiceSessionInteractionInboundResourceStatus = "undelivered"
	ProxyV1ServiceSessionInteractionInboundResourceStatusUnknown         ProxyV1ServiceSessionInteractionInboundResourceStatus = "unknown"
)

// Defines values for ProxyV1ServiceSessionInteractionOutboundResourceStatus.
const (
	ProxyV1ServiceSessionInteractionOutboundResourceStatusAccepted        ProxyV1ServiceSessionInteractionOutboundResourceStatus = "accepted"
	ProxyV1ServiceSessionInteractionOutboundResourceStatusAnswered        ProxyV1ServiceSessionInteractionOutboundResourceStatus = "answered"
	ProxyV1ServiceSessionInteractionOutboundResourceStatusBusy            ProxyV1ServiceSessionInteractionOutboundResourceStatus = "busy"
	ProxyV1ServiceSessionInteractionOutboundResourceStatusCanceled        ProxyV1ServiceSessionInteractionOutboundResourceStatus = "canceled"
	ProxyV1ServiceSessionInteractionOutboundResourceStatusCompleted       ProxyV1ServiceSessionInteractionOutboundResourceStatus = "completed"
	ProxyV1ServiceSessionInteractionOutboundResourceStatusDeleted         ProxyV1ServiceSessionInteractionOutboundResourceStatus = "deleted"
	ProxyV1ServiceSessionInteractionOutboundResourceStatusDelivered       ProxyV1ServiceSessionInteractionOutboundResourceStatus = "delivered"
	ProxyV1ServiceSessionInteractionOutboundResourceStatusDeliveryUnknown ProxyV1ServiceSessionInteractionOutboundResourceStatus = "delivery-unknown"
	ProxyV1ServiceSessionInteractionOutboundResourceStatusFailed          ProxyV1ServiceSessionInteractionOutboundResourceStatus = "failed"
	ProxyV1ServiceSessionInteractionOutboundResourceStatusInProgress      ProxyV1ServiceSessionInteractionOutboundResourceStatus = "in-progress"
	ProxyV1ServiceSessionInteractionOutboundResourceStatusInitiated       ProxyV1ServiceSessionInteractionOutboundResourceStatus = "initiated"
	ProxyV1ServiceSessionInteractionOutboundResourceStatusNoAnswer        ProxyV1ServiceSessionInteractionOutboundResourceStatus = "no-answer"
	ProxyV1ServiceSessionInteractionOutboundResourceStatusQueued          ProxyV1ServiceSessionInteractionOutboundResourceStatus = "queued"
	ProxyV1ServiceSessionInteractionOutboundResourceStatusReceived        ProxyV1ServiceSessionInteractionOutboundResourceStatus = "received"
	ProxyV1ServiceSessionInteractionOutboundResourceStatusReceiving       ProxyV1ServiceSessionInteractionOutboundResourceStatus = "receiving"
	ProxyV1ServiceSessionInteractionOutboundResourceStatusRinging         ProxyV1ServiceSessionInteractionOutboundResourceStatus = "ringing"
	ProxyV1ServiceSessionInteractionOutboundResourceStatusScheduled       ProxyV1ServiceSessionInteractionOutboundResourceStatus = "scheduled"
	ProxyV1ServiceSessionInteractionOutboundResourceStatusSending         ProxyV1ServiceSessionInteractionOutboundResourceStatus = "sending"
	ProxyV1ServiceSessionInteractionOutboundResourceStatusSent            ProxyV1ServiceSessionInteractionOutboundResourceStatus = "sent"
	ProxyV1ServiceSessionInteractionOutboundResourceStatusUndelivered     ProxyV1ServiceSessionInteractionOutboundResourceStatus = "undelivered"
	ProxyV1ServiceSessionInteractionOutboundResourceStatusUnknown         ProxyV1ServiceSessionInteractionOutboundResourceStatus = "unknown"
)

// Defines values for ProxyV1ServiceSessionInteractionType.
const (
	ProxyV1ServiceSessionInteractionTypeMessage ProxyV1ServiceSessionInteractionType = "message"
	ProxyV1ServiceSessionInteractionTypeUnknown ProxyV1ServiceSessionInteractionType = "unknown"
	ProxyV1ServiceSessionInteractionTypeVoice   ProxyV1ServiceSessionInteractionType = "voice"
)

// Defines values for ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus.
const (
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusAccepted        ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "accepted"
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusAnswered        ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "answered"
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusBusy            ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "busy"
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusCanceled        ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "canceled"
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusCompleted       ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "completed"
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusDeleted         ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "deleted"
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusDelivered       ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "delivered"
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusDeliveryUnknown ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "delivery-unknown"
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusFailed          ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "failed"
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusInProgress      ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "in-progress"
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusInitiated       ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "initiated"
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusNoAnswer        ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "no-answer"
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusQueued          ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "queued"
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusReceived        ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "received"
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusReceiving       ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "receiving"
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusRinging         ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "ringing"
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusScheduled       ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "scheduled"
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusSending         ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "sending"
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusSent            ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "sent"
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusUndelivered     ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "undelivered"
	ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatusUnknown         ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus = "unknown"
)

// Defines values for ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus.
const (
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusAccepted        ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "accepted"
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusAnswered        ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "answered"
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusBusy            ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "busy"
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusCanceled        ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "canceled"
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusCompleted       ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "completed"
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusDeleted         ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "deleted"
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusDelivered       ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "delivered"
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusDeliveryUnknown ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "delivery-unknown"
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusFailed          ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "failed"
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusInProgress      ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "in-progress"
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusInitiated       ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "initiated"
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusNoAnswer        ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "no-answer"
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusQueued          ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "queued"
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusReceived        ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "received"
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusReceiving       ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "receiving"
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusRinging         ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "ringing"
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusScheduled       ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "scheduled"
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusSending         ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "sending"
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusSent            ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "sent"
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusUndelivered     ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "undelivered"
	ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatusUnknown         ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus = "unknown"
)

// Defines values for ProxyV1ServiceSessionParticipantMessageInteractionType.
const (
	ProxyV1ServiceSessionParticipantMessageInteractionTypeMessage ProxyV1ServiceSessionParticipantMessageInteractionType = "message"
	ProxyV1ServiceSessionParticipantMessageInteractionTypeUnknown ProxyV1ServiceSessionParticipantMessageInteractionType = "unknown"
	ProxyV1ServiceSessionParticipantMessageInteractionTypeVoice   ProxyV1ServiceSessionParticipantMessageInteractionType = "voice"
)

// ProxyV1Service defines model for proxy.v1.service.
type ProxyV1Service struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// The URL we call when the interaction status changes
	CallbackUrl *string `json:"callback_url"`

	// The SID of the Chat Service Instance
	ChatInstanceSid *string `json:"chat_instance_sid"`

	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// Default TTL for a Session, in seconds
	DefaultTtl *int `json:"default_ttl"`

	// Where a proxy number must be located relative to the participant identifier
	GeoMatchLevel *ProxyV1ServiceGeoMatchLevel `json:"geo_match_level"`

	// The URL we call on each interaction
	InterceptCallbackUrl *string `json:"intercept_callback_url"`

	// The URLs of resources related to the Service
	Links *map[string]interface{} `json:"links"`

	// The preference for Proxy Number selection for the Service instance
	NumberSelectionBehavior *ProxyV1ServiceNumberSelectionBehavior `json:"number_selection_behavior"`

	// The URL we call when an inbound call or SMS action occurs on a closed or non-existent Session
	OutOfSessionCallbackUrl *string `json:"out_of_session_callback_url"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"unique_name"`

	// The absolute URL of the Service resource
	Url *string `json:"url"`
}

// Where a proxy number must be located relative to the participant identifier
type ProxyV1ServiceGeoMatchLevel string

// The preference for Proxy Number selection for the Service instance
type ProxyV1ServiceNumberSelectionBehavior string

// ProxyV1ServicePhoneNumber defines model for proxy.v1.service.phone_number.
type ProxyV1ServicePhoneNumber struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// The capabilities of the phone number
	Capabilities *struct {
		Fax   *bool `json:"fax,omitempty"`
		Mms   *bool `json:"mms,omitempty"`
		Sms   *bool `json:"sms,omitempty"`
		Voice *bool `json:"voice,omitempty"`
	} `json:"capabilities"`

	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name"`

	// The number of open session assigned to the number.
	InUse *int `json:"in_use"`

	// Reserve the phone number for manual assignment to participants only
	IsReserved *bool `json:"is_reserved"`

	// The ISO Country Code
	IsoCountry *string `json:"iso_country"`

	// The phone number in E.164 format
	PhoneNumber *string `json:"phone_number"`

	// The SID of the PhoneNumber resource's parent Service resource
	ServiceSid *string `json:"service_sid"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The absolute URL of the PhoneNumber resource
	Url *string `json:"url"`
}

// ProxyV1ServiceSession defines model for proxy.v1.service.session.
type ProxyV1ServiceSession struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// The reason the Session ended
	ClosedReason *string `json:"closed_reason"`

	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The ISO 8601 date when the Session ended
	DateEnded *time.Time `json:"date_ended"`

	// The ISO 8601 date when the Session should expire
	DateExpiry *time.Time `json:"date_expiry"`

	// The ISO 8601 date when the Session last had an interaction
	DateLastInteraction *time.Time `json:"date_last_interaction"`

	// The ISO 8601 date when the Session started
	DateStarted *time.Time `json:"date_started"`

	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// The URLs of resources related to the Session
	Links *map[string]interface{} `json:"links"`

	// The Mode of the Session
	Mode *ProxyV1ServiceSessionMode `json:"mode"`

	// The SID of the resource's parent Service
	ServiceSid *string `json:"service_sid"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The status of the Session
	Status *ProxyV1ServiceSessionStatus `json:"status"`

	// When the session will expire
	Ttl *int `json:"ttl"`

	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"unique_name"`

	// The absolute URL of the Session resource
	Url *string `json:"url"`
}

// The Mode of the Session
type ProxyV1ServiceSessionMode string

// The status of the Session
type ProxyV1ServiceSessionStatus string

// ProxyV1ServiceSessionInteraction defines model for proxy.v1.service.session.interaction.
type ProxyV1ServiceSessionInteraction struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// A JSON string that includes the message body of message interactions
	Data *string `json:"data"`

	// The ISO 8601 date and time in GMT when the Interaction was created
	DateCreated *time.Time `json:"date_created"`

	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// The SID of the inbound Participant resource
	InboundParticipantSid *string `json:"inbound_participant_sid"`

	// The SID of the inbound resource
	InboundResourceSid *string `json:"inbound_resource_sid"`

	// The inbound resource status of the Interaction
	InboundResourceStatus *ProxyV1ServiceSessionInteractionInboundResourceStatus `json:"inbound_resource_status"`

	// The inbound resource type
	InboundResourceType *string `json:"inbound_resource_type"`

	// The URL of the Twilio inbound resource
	InboundResourceUrl *string `json:"inbound_resource_url"`

	// The SID of the outbound Participant
	OutboundParticipantSid *string `json:"outbound_participant_sid"`

	// The SID of the outbound resource
	OutboundResourceSid *string `json:"outbound_resource_sid"`

	// The outbound resource status of the Interaction
	OutboundResourceStatus *ProxyV1ServiceSessionInteractionOutboundResourceStatus `json:"outbound_resource_status"`

	// The outbound resource type
	OutboundResourceType *string `json:"outbound_resource_type"`

	// The URL of the Twilio outbound resource
	OutboundResourceUrl *string `json:"outbound_resource_url"`

	// The SID of the resource's parent Service
	ServiceSid *string `json:"service_sid"`

	// The SID of the resource's parent Session
	SessionSid *string `json:"session_sid"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The Type of the Interaction
	Type *ProxyV1ServiceSessionInteractionType `json:"type"`

	// The absolute URL of the Interaction resource
	Url *string `json:"url"`
}

// The inbound resource status of the Interaction
type ProxyV1ServiceSessionInteractionInboundResourceStatus string

// The outbound resource status of the Interaction
type ProxyV1ServiceSessionInteractionOutboundResourceStatus string

// The Type of the Interaction
type ProxyV1ServiceSessionInteractionType string

// ProxyV1ServiceSessionParticipant defines model for proxy.v1.service.session.participant.
type ProxyV1ServiceSessionParticipant struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The ISO 8601 date the Participant was removed
	DateDeleted *time.Time `json:"date_deleted"`

	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// The string that you assigned to describe the participant
	FriendlyName *string `json:"friendly_name"`

	// The phone number or channel identifier of the Participant
	Identifier *string `json:"identifier"`

	// The URLs to resources related the participant
	Links *map[string]interface{} `json:"links"`

	// The phone number or short code of the participant's partner
	ProxyIdentifier *string `json:"proxy_identifier"`

	// The SID of the Proxy Identifier assigned to the Participant
	ProxyIdentifierSid *string `json:"proxy_identifier_sid"`

	// The SID of the resource's parent Service
	ServiceSid *string `json:"service_sid"`

	// The SID of the resource's parent Session
	SessionSid *string `json:"session_sid"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The absolute URL of the Participant resource
	Url *string `json:"url"`
}

// ProxyV1ServiceSessionParticipantMessageInteraction defines model for proxy.v1.service.session.participant.message_interaction.
type ProxyV1ServiceSessionParticipantMessageInteraction struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// A JSON string that includes the message body sent to the participant
	Data *string `json:"data"`

	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// Always empty for Message Interactions
	InboundParticipantSid *string `json:"inbound_participant_sid"`

	// Always empty for Message Interactions
	InboundResourceSid *string `json:"inbound_resource_sid"`

	// Always empty for Message Interactions
	InboundResourceStatus *ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus `json:"inbound_resource_status"`

	// Always empty for Message Interactions
	InboundResourceType *string `json:"inbound_resource_type"`

	// Always empty for Message Interactions
	InboundResourceUrl *string `json:"inbound_resource_url"`

	// The SID of the outbound Participant resource
	OutboundParticipantSid *string `json:"outbound_participant_sid"`

	// The SID of the outbound Message resource
	OutboundResourceSid *string `json:"outbound_resource_sid"`

	// The outbound resource status
	OutboundResourceStatus *ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus `json:"outbound_resource_status"`

	// The outbound resource type
	OutboundResourceType *string `json:"outbound_resource_type"`

	// The URL of the Twilio message resource
	OutboundResourceUrl *string `json:"outbound_resource_url"`

	// The SID of the Participant resource
	ParticipantSid *string `json:"participant_sid"`

	// The SID of the resource's parent Service
	ServiceSid *string `json:"service_sid"`

	// The SID of the resource's parent Session
	SessionSid *string `json:"session_sid"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The Type of Message Interaction
	Type *ProxyV1ServiceSessionParticipantMessageInteractionType `json:"type"`

	// The absolute URL of the MessageInteraction resource
	Url *string `json:"url"`
}

// Always empty for Message Interactions
type ProxyV1ServiceSessionParticipantMessageInteractionInboundResourceStatus string

// The outbound resource status
type ProxyV1ServiceSessionParticipantMessageInteractionOutboundResourceStatus string

// The Type of Message Interaction
type ProxyV1ServiceSessionParticipantMessageInteractionType string

// ProxyV1ServiceShortCode defines model for proxy.v1.service.short_code.
type ProxyV1ServiceShortCode struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// The capabilities of the short code
	Capabilities *struct {
		Fax   *bool `json:"fax,omitempty"`
		Mms   *bool `json:"mms,omitempty"`
		Sms   *bool `json:"sms,omitempty"`
		Voice *bool `json:"voice,omitempty"`
	} `json:"capabilities"`

	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// Whether the short code should be reserved for manual assignment to participants only
	IsReserved *bool `json:"is_reserved"`

	// The ISO Country Code
	IsoCountry *string `json:"iso_country"`

	// The SID of the resource's parent Service
	ServiceSid *string `json:"service_sid"`

	// The short code's number
	ShortCode *string `json:"short_code"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The absolute URL of the ShortCode resource
	Url *string `json:"url"`
}

// ListServiceParams defines parameters for ListService.
type ListServiceParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListPhoneNumberParams defines parameters for ListPhoneNumber.
type ListPhoneNumberParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListSessionParams defines parameters for ListSession.
type ListSessionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListInteractionParams defines parameters for ListInteraction.
type ListInteractionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListParticipantParams defines parameters for ListParticipant.
type ListParticipantParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListMessageInteractionParams defines parameters for ListMessageInteraction.
type ListMessageInteractionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListShortCodeParams defines parameters for ListShortCode.
type ListShortCodeParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}