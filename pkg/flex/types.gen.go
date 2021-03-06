// Package flex provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package flex

import (
	"time"
)

const (
	AccountSid_authTokenScopes = "accountSid_authToken.Scopes"
)

// Defines values for FlexV1ConfigurationStatus.
const (
	FlexV1ConfigurationStatusInprogress FlexV1ConfigurationStatus = "inprogress"
	FlexV1ConfigurationStatusNotstarted FlexV1ConfigurationStatus = "notstarted"
	FlexV1ConfigurationStatusOk         FlexV1ConfigurationStatus = "ok"
)

// Defines values for FlexV1FlexFlowChannelType.
const (
	FlexV1FlexFlowChannelTypeCustom   FlexV1FlexFlowChannelType = "custom"
	FlexV1FlexFlowChannelTypeFacebook FlexV1FlexFlowChannelType = "facebook"
	FlexV1FlexFlowChannelTypeLine     FlexV1FlexFlowChannelType = "line"
	FlexV1FlexFlowChannelTypeSms      FlexV1FlexFlowChannelType = "sms"
	FlexV1FlexFlowChannelTypeWeb      FlexV1FlexFlowChannelType = "web"
	FlexV1FlexFlowChannelTypeWhatsapp FlexV1FlexFlowChannelType = "whatsapp"
)

// Defines values for FlexV1FlexFlowIntegrationType.
const (
	FlexV1FlexFlowIntegrationTypeExternal FlexV1FlexFlowIntegrationType = "external"
	FlexV1FlexFlowIntegrationTypeStudio   FlexV1FlexFlowIntegrationType = "studio"
	FlexV1FlexFlowIntegrationTypeTask     FlexV1FlexFlowIntegrationType = "task"
)

// FlexV1Channel defines model for flex.v1.channel.
type FlexV1Channel struct {
	// The SID of the Account that created the resource and owns this Workflow
	AccountSid *string `json:"account_sid"`

	// The ISO 8601 date and time in GMT when the Flex chat channel was created
	DateCreated *time.Time `json:"date_created"`

	// The ISO 8601 date and time in GMT when the Flex chat channel was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// The SID of the Flex Flow
	FlexFlowSid *string `json:"flex_flow_sid"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The SID of the TaskRouter Task
	TaskSid *string `json:"task_sid"`

	// The absolute URL of the Flex chat channel resource
	Url *string `json:"url"`

	// The SID of the chat user
	UserSid *string `json:"user_sid"`
}

// FlexV1Configuration defines model for flex.v1.configuration.
type FlexV1Configuration struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// An object that contains application-specific data
	Attributes *map[string]interface{} `json:"attributes"`

	// Whether call recording is enabled
	CallRecordingEnabled *bool `json:"call_recording_enabled"`

	// The call recording webhook URL
	CallRecordingWebhookUrl *string `json:"call_recording_webhook_url"`

	// The SID of the chat service this user belongs to
	ChatServiceInstanceSid *string `json:"chat_service_instance_sid"`

	// An object that contains the CRM attributes
	CrmAttributes *map[string]interface{} `json:"crm_attributes"`

	// The CRM Callback URL
	CrmCallbackUrl *string `json:"crm_callback_url"`

	// Whether CRM is present for Flex
	CrmEnabled *bool `json:"crm_enabled"`

	// The CRM Fallback URL
	CrmFallbackUrl *string `json:"crm_fallback_url"`

	// The CRM Type
	CrmType *string `json:"crm_type"`

	// The ISO 8601 date and time in GMT when the Configuration resource was created
	DateCreated *time.Time `json:"date_created"`

	// The ISO 8601 date and time in GMT when the Configuration resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// Setting to enable Flex UI redirection
	FlexInsightsDrilldown *bool `json:"flex_insights_drilldown"`

	// Object that controls workspace reporting
	FlexInsightsHr *map[string]interface{} `json:"flex_insights_hr"`

	// The SID of the Flex service instance
	FlexServiceInstanceSid *string `json:"flex_service_instance_sid"`

	// URL to redirect to in case drilldown is enabled.
	FlexUrl *string `json:"flex_url"`

	// A list of objects that contain the configurations for the Integrations supported in this configuration
	Integrations *[]map[string]interface{} `json:"integrations"`

	// Configurable parameters for Markdown
	Markdown *map[string]interface{} `json:"markdown"`

	// The SID of the Messaging service instance
	MessagingServiceInstanceSid *string `json:"messaging_service_instance_sid"`

	// Configurable parameters for Notifications
	Notifications *map[string]interface{} `json:"notifications"`

	// The list of outbound call flows
	OutboundCallFlows *map[string]interface{} `json:"outbound_call_flows"`

	// The plugin service attributes
	PluginServiceAttributes *map[string]interface{} `json:"plugin_service_attributes"`

	// Whether the plugin service enabled
	PluginServiceEnabled *bool `json:"plugin_service_enabled"`

	// The list of public attributes
	PublicAttributes *map[string]interface{} `json:"public_attributes"`

	// Configurable parameters for Queues Statistics
	QueueStatsConfiguration *map[string]interface{} `json:"queue_stats_configuration"`

	// The URL where the Flex instance is hosted
	RuntimeDomain *string `json:"runtime_domain"`

	// The list of serverless service SIDs
	ServerlessServiceSids *[]string `json:"serverless_service_sids"`

	// The Flex Service version
	ServiceVersion *string `json:"service_version"`

	// The status of the Flex onboarding
	Status *FlexV1ConfigurationStatus `json:"status"`

	// The TaskRouter SID of the offline activity
	TaskrouterOfflineActivitySid *string `json:"taskrouter_offline_activity_sid"`

	// The Skill description for TaskRouter workers
	TaskrouterSkills *[]map[string]interface{} `json:"taskrouter_skills"`

	// The SID of the TaskRouter Target TaskQueue
	TaskrouterTargetTaskqueueSid *string `json:"taskrouter_target_taskqueue_sid"`

	// The SID of the TaskRouter target Workflow
	TaskrouterTargetWorkflowSid *string `json:"taskrouter_target_workflow_sid"`

	// The list of TaskRouter TaskQueues
	TaskrouterTaskqueues *[]map[string]interface{} `json:"taskrouter_taskqueues"`

	// The TaskRouter Worker attributes
	TaskrouterWorkerAttributes *map[string]interface{} `json:"taskrouter_worker_attributes"`

	// The TaskRouter default channel capacities and availability for workers
	TaskrouterWorkerChannels *map[string]interface{} `json:"taskrouter_worker_channels"`

	// The SID of the TaskRouter Workspace
	TaskrouterWorkspaceSid *string `json:"taskrouter_workspace_sid"`

	// The object that describes Flex UI characteristics and settings
	UiAttributes *map[string]interface{} `json:"ui_attributes"`

	// The object that defines the NPM packages and versions to be used in Hosted Flex
	UiDependencies *map[string]interface{} `json:"ui_dependencies"`

	// The primary language of the Flex UI
	UiLanguage *string `json:"ui_language"`

	// The Pinned UI version
	UiVersion *string `json:"ui_version"`

	// The absolute URL of the Configuration resource
	Url *string `json:"url"`
}

// The status of the Flex onboarding
type FlexV1ConfigurationStatus string

// FlexV1FlexFlow defines model for flex.v1.flex_flow.
type FlexV1FlexFlow struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// The channel type
	ChannelType *FlexV1FlexFlowChannelType `json:"channel_type"`

	// The SID of the chat service
	ChatServiceSid *string `json:"chat_service_sid"`

	// The channel contact's Identity
	ContactIdentity *string `json:"contact_identity"`

	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// Whether the Flex Flow is enabled
	Enabled *bool `json:"enabled"`

	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name"`

	// An object that contains specific parameters for the integration
	Integration *map[string]interface{} `json:"integration"`

	// The software that will handle inbound messages.
	IntegrationType *FlexV1FlexFlowIntegrationType `json:"integration_type"`

	// Remove active Proxy sessions if the corresponding Task is deleted.
	JanitorEnabled *bool `json:"janitor_enabled"`

	// Re-use this chat channel for future interactions with a contact
	LongLived *bool `json:"long_lived"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The absolute URL of the Flex Flow resource
	Url *string `json:"url"`
}

// The channel type
type FlexV1FlexFlowChannelType string

// The software that will handle inbound messages.
type FlexV1FlexFlowIntegrationType string

// FlexV1WebChannel defines model for flex.v1.web_channel.
type FlexV1WebChannel struct {
	// The SID of the Account that created the resource and owns this Workflow
	AccountSid *string `json:"account_sid"`

	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// The SID of the Flex Flow
	FlexFlowSid *string `json:"flex_flow_sid"`

	// The unique string that identifies the WebChannel resource
	Sid *string `json:"sid"`

	// The absolute URL of the WebChannel resource
	Url *string `json:"url"`
}

// ListChannelParams defines parameters for ListChannel.
type ListChannelParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// FetchConfigurationParams defines parameters for FetchConfiguration.
type FetchConfigurationParams struct {
	// The Pinned UI version of the Configuration resource to fetch.
	UiVersion *string `json:"UiVersion,omitempty"`
}

// ListFlexFlowParams defines parameters for ListFlexFlow.
type ListFlexFlowParams struct {
	// The `friendly_name` of the Flex Flow resources to read.
	FriendlyName *string `json:"FriendlyName,omitempty"`

	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListWebChannelParams defines parameters for ListWebChannel.
type ListWebChannelParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}
