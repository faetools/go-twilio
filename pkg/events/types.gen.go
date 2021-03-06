// Package events provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package events

import (
	"time"
)

const (
	AccountSid_authTokenScopes = "accountSid_authToken.Scopes"
)

// Defines values for EventsV1SinkSinkType.
const (
	EventsV1SinkSinkTypeKinesis EventsV1SinkSinkType = "kinesis"
	EventsV1SinkSinkTypeSegment EventsV1SinkSinkType = "segment"
	EventsV1SinkSinkTypeWebhook EventsV1SinkSinkType = "webhook"
)

// Defines values for EventsV1SinkStatus.
const (
	EventsV1SinkStatusActive      EventsV1SinkStatus = "active"
	EventsV1SinkStatusFailed      EventsV1SinkStatus = "failed"
	EventsV1SinkStatusInitialized EventsV1SinkStatus = "initialized"
	EventsV1SinkStatusValidating  EventsV1SinkStatus = "validating"
)

// EventsV1EventType defines model for events.v1.event_type.
type EventsV1EventType struct {
	// The date this Event Type was created.
	DateCreated *time.Time `json:"date_created"`

	// The date this Event Type was updated.
	DateUpdated *time.Time `json:"date_updated"`

	// Event Type description.
	Description *string                 `json:"description"`
	Links       *map[string]interface{} `json:"links"`

	// The Schema identifier for this Event Type.
	SchemaId *string `json:"schema_id"`

	// The Event Type identifier.
	Type *string `json:"type"`

	// The URL of this resource.
	Url *string `json:"url"`
}

// EventsV1Schema defines model for events.v1.schema.
type EventsV1Schema struct {
	// Schema Identifier.
	Id *string `json:"id"`

	// Latest schema version.
	LatestVersion *int `json:"latest_version"`

	// The date that the latest schema version was created.
	LatestVersionDateCreated *time.Time `json:"latest_version_date_created"`

	// Nested resource URLs.
	Links *map[string]interface{} `json:"links"`

	// The URL of this resource.
	Url *string `json:"url"`
}

// EventsV1SchemaSchemaVersion defines model for events.v1.schema.schema_version.
type EventsV1SchemaSchemaVersion struct {
	// The date the schema version was created.
	DateCreated *time.Time `json:"date_created"`

	// The unique identifier of the schema.
	Id  *string `json:"id"`
	Raw *string `json:"raw"`

	// The version of this schema.
	SchemaVersion *int `json:"schema_version"`

	// The URL of this resource.
	Url *string `json:"url"`
}

// EventsV1Sink defines model for events.v1.sink.
type EventsV1Sink struct {
	// The date this Sink was created
	DateCreated *time.Time `json:"date_created"`

	// The date this Sink was updated
	DateUpdated *time.Time `json:"date_updated"`

	// Sink Description
	Description *string `json:"description"`

	// Nested resource URLs.
	Links *map[string]interface{} `json:"links"`

	// A string that uniquely identifies this Sink.
	Sid *string `json:"sid"`

	// JSON Sink configuration.
	SinkConfiguration *map[string]interface{} `json:"sink_configuration"`

	// Sink type.
	SinkType *EventsV1SinkSinkType `json:"sink_type"`

	// The Status of this Sink
	Status *EventsV1SinkStatus `json:"status"`

	// The URL of this resource.
	Url *string `json:"url"`
}

// Sink type.
type EventsV1SinkSinkType string

// The Status of this Sink
type EventsV1SinkStatus string

// EventsV1SinkSinkTest defines model for events.v1.sink.sink_test.
type EventsV1SinkSinkTest struct {
	// Feedback indicating whether the test event was generated.
	Result *string `json:"result"`
}

// EventsV1SinkSinkValidate defines model for events.v1.sink.sink_validate.
type EventsV1SinkSinkValidate struct {
	// Feedback indicating whether the given Sink was validated.
	Result *string `json:"result"`
}

// EventsV1Subscription defines model for events.v1.subscription.
type EventsV1Subscription struct {
	// Account SID.
	AccountSid *string `json:"account_sid"`

	// The date this Subscription was created
	DateCreated *time.Time `json:"date_created"`

	// The date this Subscription was updated
	DateUpdated *time.Time `json:"date_updated"`

	// Subscription description
	Description *string `json:"description"`

	// Nested resource URLs.
	Links *map[string]interface{} `json:"links"`

	// A string that uniquely identifies this Subscription.
	Sid *string `json:"sid"`

	// Sink SID.
	SinkSid *string `json:"sink_sid"`

	// The URL of this resource.
	Url *string `json:"url"`
}

// EventsV1SubscriptionSubscribedEvent defines model for events.v1.subscription.subscribed_event.
type EventsV1SubscriptionSubscribedEvent struct {
	// Account SID.
	AccountSid *string `json:"account_sid"`

	// The schema version that the subscription should use.
	SchemaVersion *int `json:"schema_version"`

	// Subscription SID.
	SubscriptionSid *string `json:"subscription_sid"`

	// Type of event being subscribed to.
	Type *string `json:"type"`

	// The URL of this resource.
	Url *string `json:"url"`
}

// ListSchemaVersionParams defines parameters for ListSchemaVersion.
type ListSchemaVersionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListSinkParams defines parameters for ListSink.
type ListSinkParams struct {
	// A boolean query parameter filtering the results to return sinks used/not used by a subscription.
	InUse *bool `json:"InUse,omitempty"`

	// A String query parameter filtering the results by status `initialized`, `validating`, `active` or `failed`.
	Status *string `json:"Status,omitempty"`

	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListSubscriptionParams defines parameters for ListSubscription.
type ListSubscriptionParams struct {
	// The SID of the sink that the list of Subscriptions should be filtered by.
	SinkSid *string `json:"SinkSid,omitempty"`

	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListSubscribedEventParams defines parameters for ListSubscribedEvent.
type ListSubscribedEventParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListEventTypeParams defines parameters for ListEventType.
type ListEventTypeParams struct {
	// A string parameter filtering the results to return only the Event Types using a given schema.
	SchemaId *string `json:"SchemaId,omitempty"`

	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}
