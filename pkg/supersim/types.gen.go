// Package supersim provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package supersim

import (
	"time"
)

const (
	AccountSid_authTokenScopes = "accountSid_authToken.Scopes"
)

// Defines values for SupersimV1CommandDirection.
const (
	SupersimV1CommandDirectionFromSim SupersimV1CommandDirection = "from_sim"
	SupersimV1CommandDirectionToSim   SupersimV1CommandDirection = "to_sim"
)

// Defines values for SupersimV1CommandStatus.
const (
	SupersimV1CommandStatusDelivered SupersimV1CommandStatus = "delivered"
	SupersimV1CommandStatusFailed    SupersimV1CommandStatus = "failed"
	SupersimV1CommandStatusQueued    SupersimV1CommandStatus = "queued"
	SupersimV1CommandStatusReceived  SupersimV1CommandStatus = "received"
	SupersimV1CommandStatusSent      SupersimV1CommandStatus = "sent"
)

// Defines values for SupersimV1EsimProfileStatus.
const (
	SupersimV1EsimProfileStatusAvailable  SupersimV1EsimProfileStatus = "available"
	SupersimV1EsimProfileStatusDownloaded SupersimV1EsimProfileStatus = "downloaded"
	SupersimV1EsimProfileStatusFailed     SupersimV1EsimProfileStatus = "failed"
	SupersimV1EsimProfileStatusInstalled  SupersimV1EsimProfileStatus = "installed"
	SupersimV1EsimProfileStatusNew        SupersimV1EsimProfileStatus = "new"
	SupersimV1EsimProfileStatusReserving  SupersimV1EsimProfileStatus = "reserving"
)

// Defines values for SupersimV1FleetCommandsMethod.
const (
	SupersimV1FleetCommandsMethodDELETE SupersimV1FleetCommandsMethod = "DELETE"
	SupersimV1FleetCommandsMethodGET    SupersimV1FleetCommandsMethod = "GET"
	SupersimV1FleetCommandsMethodHEAD   SupersimV1FleetCommandsMethod = "HEAD"
	SupersimV1FleetCommandsMethodPATCH  SupersimV1FleetCommandsMethod = "PATCH"
	SupersimV1FleetCommandsMethodPOST   SupersimV1FleetCommandsMethod = "POST"
	SupersimV1FleetCommandsMethodPUT    SupersimV1FleetCommandsMethod = "PUT"
)

// Defines values for SupersimV1FleetDataMetering.
const (
	SupersimV1FleetDataMeteringPayg SupersimV1FleetDataMetering = "payg"
)

// Defines values for SupersimV1FleetIpCommandsMethod.
const (
	SupersimV1FleetIpCommandsMethodDELETE SupersimV1FleetIpCommandsMethod = "DELETE"
	SupersimV1FleetIpCommandsMethodGET    SupersimV1FleetIpCommandsMethod = "GET"
	SupersimV1FleetIpCommandsMethodHEAD   SupersimV1FleetIpCommandsMethod = "HEAD"
	SupersimV1FleetIpCommandsMethodPATCH  SupersimV1FleetIpCommandsMethod = "PATCH"
	SupersimV1FleetIpCommandsMethodPOST   SupersimV1FleetIpCommandsMethod = "POST"
	SupersimV1FleetIpCommandsMethodPUT    SupersimV1FleetIpCommandsMethod = "PUT"
)

// Defines values for SupersimV1FleetSmsCommandsMethod.
const (
	SupersimV1FleetSmsCommandsMethodDELETE SupersimV1FleetSmsCommandsMethod = "DELETE"
	SupersimV1FleetSmsCommandsMethodGET    SupersimV1FleetSmsCommandsMethod = "GET"
	SupersimV1FleetSmsCommandsMethodHEAD   SupersimV1FleetSmsCommandsMethod = "HEAD"
	SupersimV1FleetSmsCommandsMethodPATCH  SupersimV1FleetSmsCommandsMethod = "PATCH"
	SupersimV1FleetSmsCommandsMethodPOST   SupersimV1FleetSmsCommandsMethod = "POST"
	SupersimV1FleetSmsCommandsMethodPUT    SupersimV1FleetSmsCommandsMethod = "PUT"
)

// Defines values for SupersimV1IpCommandDirection.
const (
	SupersimV1IpCommandDirectionFromSim SupersimV1IpCommandDirection = "from_sim"
	SupersimV1IpCommandDirectionToSim   SupersimV1IpCommandDirection = "to_sim"
)

// Defines values for SupersimV1IpCommandPayloadType.
const (
	SupersimV1IpCommandPayloadTypeBinary SupersimV1IpCommandPayloadType = "binary"
	SupersimV1IpCommandPayloadTypeText   SupersimV1IpCommandPayloadType = "text"
)

// Defines values for SupersimV1IpCommandStatus.
const (
	SupersimV1IpCommandStatusFailed   SupersimV1IpCommandStatus = "failed"
	SupersimV1IpCommandStatusQueued   SupersimV1IpCommandStatus = "queued"
	SupersimV1IpCommandStatusReceived SupersimV1IpCommandStatus = "received"
	SupersimV1IpCommandStatusSent     SupersimV1IpCommandStatus = "sent"
)

// Defines values for SupersimV1SimStatus.
const (
	SupersimV1SimStatusActive    SupersimV1SimStatus = "active"
	SupersimV1SimStatusInactive  SupersimV1SimStatus = "inactive"
	SupersimV1SimStatusNew       SupersimV1SimStatus = "new"
	SupersimV1SimStatusReady     SupersimV1SimStatus = "ready"
	SupersimV1SimStatusScheduled SupersimV1SimStatus = "scheduled"
)

// Defines values for SupersimV1SimBillingPeriodPeriodType.
const (
	SupersimV1SimBillingPeriodPeriodTypeActive SupersimV1SimBillingPeriodPeriodType = "active"
	SupersimV1SimBillingPeriodPeriodTypeReady  SupersimV1SimBillingPeriodPeriodType = "ready"
)

// Defines values for SupersimV1SmsCommandDirection.
const (
	SupersimV1SmsCommandDirectionFromSim SupersimV1SmsCommandDirection = "from_sim"
	SupersimV1SmsCommandDirectionToSim   SupersimV1SmsCommandDirection = "to_sim"
)

// Defines values for SupersimV1SmsCommandStatus.
const (
	SupersimV1SmsCommandStatusDelivered SupersimV1SmsCommandStatus = "delivered"
	SupersimV1SmsCommandStatusFailed    SupersimV1SmsCommandStatus = "failed"
	SupersimV1SmsCommandStatusQueued    SupersimV1SmsCommandStatus = "queued"
	SupersimV1SmsCommandStatusReceived  SupersimV1SmsCommandStatus = "received"
	SupersimV1SmsCommandStatusSent      SupersimV1SmsCommandStatus = "sent"
)

// SupersimV1Command defines model for supersim.v1.command.
type SupersimV1Command struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// The message body of the command sent to or from the SIM
	Command *string `json:"command"`

	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// The direction of the Command
	Direction *SupersimV1CommandDirection `json:"direction"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The SID of the SIM that this Command was sent to or from
	SimSid *string `json:"sim_sid"`

	// The status of the Command
	Status *SupersimV1CommandStatus `json:"status"`

	// The absolute URL of the Command resource
	Url *string `json:"url"`
}

// The direction of the Command
type SupersimV1CommandDirection string

// The status of the Command
type SupersimV1CommandStatus string

// SupersimV1EsimProfile defines model for supersim.v1.esim_profile.
type SupersimV1EsimProfile struct {
	// The SID of the Account to which the eSIM Profile resource belongs
	AccountSid *string `json:"account_sid"`

	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// Identifier of the eUICC that can claim the eSIM Profile
	Eid *string `json:"eid"`

	// Code indicating the failure if the download of the SIM Profile failed and the eSIM Profile is in `failed` state
	ErrorCode *string `json:"error_code"`

	// Error message describing the failure if the download of the SIM Profile failed and the eSIM Profile is in `failed` state
	ErrorMessage *string `json:"error_message"`

	// The ICCID associated with the Sim resource
	Iccid *string `json:"iccid"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The SID of the Sim resource that this eSIM Profile controls
	SimSid *string `json:"sim_sid"`

	// Address of the SM-DP+ server from which the Profile will be downloaded
	SmdpPlusAddress *string `json:"smdp_plus_address"`

	// The status of the eSIM Profile
	Status *SupersimV1EsimProfileStatus `json:"status"`

	// The absolute URL of the eSIM Profile resource
	Url *string `json:"url"`
}

// The status of the eSIM Profile
type SupersimV1EsimProfileStatus string

// SupersimV1Fleet defines model for supersim.v1.fleet.
type SupersimV1Fleet struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// Deprecated
	CommandsEnabled *bool `json:"commands_enabled"`

	// Deprecated
	CommandsMethod *SupersimV1FleetCommandsMethod `json:"commands_method"`

	// Deprecated
	CommandsUrl *string `json:"commands_url"`

	// Defines whether SIMs in the Fleet are capable of using data connectivity
	DataEnabled *bool `json:"data_enabled"`

	// The total data usage (download and upload combined) in Megabytes that each Sim resource assigned to the Fleet resource can consume
	DataLimit *int `json:"data_limit"`

	// The model by which a SIM is metered and billed
	DataMetering *SupersimV1FleetDataMetering `json:"data_metering"`

	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// A string representing the HTTP method to use when making a request to `ip_commands_url`
	IpCommandsMethod *SupersimV1FleetIpCommandsMethod `json:"ip_commands_method"`

	// The URL that will receive a webhook when a Super SIM in the Fleet is used to send an IP Command from your device
	IpCommandsUrl *string `json:"ip_commands_url"`

	// The SID of the Network Access Profile of the Fleet
	NetworkAccessProfileSid *string `json:"network_access_profile_sid"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands
	SmsCommandsEnabled *bool `json:"sms_commands_enabled"`

	// A string representing the HTTP method to use when making a request to `sms_commands_url`
	SmsCommandsMethod *SupersimV1FleetSmsCommandsMethod `json:"sms_commands_method"`

	// The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the SMS Commands number
	SmsCommandsUrl *string `json:"sms_commands_url"`

	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"unique_name"`

	// The absolute URL of the Fleet resource
	Url *string `json:"url"`
}

// Deprecated
type SupersimV1FleetCommandsMethod string

// The model by which a SIM is metered and billed
type SupersimV1FleetDataMetering string

// A string representing the HTTP method to use when making a request to `ip_commands_url`
type SupersimV1FleetIpCommandsMethod string

// A string representing the HTTP method to use when making a request to `sms_commands_url`
type SupersimV1FleetSmsCommandsMethod string

// SupersimV1IpCommand defines model for supersim.v1.ip_command.
type SupersimV1IpCommand struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// The IP address of the device that the IP Command was sent to or received from
	DeviceIp *string `json:"device_ip"`

	// The port that the IP Command either originated from or was sent to
	DevicePort *int `json:"device_port"`

	// The direction of the IP Command
	Direction *SupersimV1IpCommandDirection `json:"direction"`

	// The payload of the IP Command sent to or from the Super SIM
	Payload *string `json:"payload"`

	// The payload type of the IP Command
	PayloadType *SupersimV1IpCommandPayloadType `json:"payload_type"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The ICCID of the Super SIM that this IP Command was sent to or from
	SimIccid *string `json:"sim_iccid"`

	// The SID of the Super SIM that this IP Command was sent to or from
	SimSid *string `json:"sim_sid"`

	// The status of the IP Command
	Status *SupersimV1IpCommandStatus `json:"status"`

	// The absolute URL of the IP Command resource
	Url *string `json:"url"`
}

// The direction of the IP Command
type SupersimV1IpCommandDirection string

// The payload type of the IP Command
type SupersimV1IpCommandPayloadType string

// The status of the IP Command
type SupersimV1IpCommandStatus string

// SupersimV1Network defines model for supersim.v1.network.
type SupersimV1Network struct {
	// A human readable identifier of this resource
	FriendlyName *string `json:"friendly_name"`

	// The MCC/MNCs included in the Network resource
	Identifiers *[]map[string]interface{} `json:"identifiers"`

	// The ISO country code of the Network resource
	IsoCountry *string `json:"iso_country"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The absolute URL of the Network resource
	Url *string `json:"url"`
}

// SupersimV1NetworkAccessProfile defines model for supersim.v1.network_access_profile.
type SupersimV1NetworkAccessProfile struct {
	// The SID of the Account that the Network Access Profile belongs to
	AccountSid *string `json:"account_sid"`

	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time              `json:"date_updated"`
	Links       *map[string]interface{} `json:"links"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"unique_name"`

	// The absolute URL of the resource
	Url *string `json:"url"`
}

// SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork defines model for supersim.v1.network_access_profile.network_access_profile_network.
type SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork struct {
	// A human readable identifier of this resource
	FriendlyName *string `json:"friendly_name"`

	// The MCC/MNCs included in the resource
	Identifiers *[]map[string]interface{} `json:"identifiers"`

	// The ISO country code of the Network resource
	IsoCountry *string `json:"iso_country"`

	// The unique string that identifies the Network Access Profile resource
	NetworkAccessProfileSid *string `json:"network_access_profile_sid"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The absolute URL of the resource
	Url *string `json:"url"`
}

// SupersimV1Sim defines model for supersim.v1.sim.
type SupersimV1Sim struct {
	// The SID of the Account that the Super SIM belongs to
	AccountSid *string `json:"account_sid"`

	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// The unique ID of the Fleet configured for this SIM
	FleetSid *string `json:"fleet_sid"`

	// The ICCID associated with the SIM
	Iccid *string                 `json:"iccid"`
	Links *map[string]interface{} `json:"links"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The status of the Super SIM
	Status *SupersimV1SimStatus `json:"status"`

	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"unique_name"`

	// The absolute URL of the Sim Resource
	Url *string `json:"url"`
}

// The status of the Super SIM
type SupersimV1SimStatus string

// SupersimV1SimBillingPeriod defines model for supersim.v1.sim.billing_period.
type SupersimV1SimBillingPeriod struct {
	// The SID of the Account the Super SIM belongs to
	AccountSid *string `json:"account_sid"`

	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// The end time of the Billing Period
	EndTime *time.Time `json:"end_time"`

	// The type of the Billing Period
	PeriodType *SupersimV1SimBillingPeriodPeriodType `json:"period_type"`

	// The SID of the Billing Period
	Sid *string `json:"sid"`

	// The SID of the Super SIM the Billing Period belongs to
	SimSid *string `json:"sim_sid"`

	// The start time of the Billing Period
	StartTime *time.Time `json:"start_time"`
}

// The type of the Billing Period
type SupersimV1SimBillingPeriodPeriodType string

// SupersimV1SmsCommand defines model for supersim.v1.sms_command.
type SupersimV1SmsCommand struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid"`

	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created"`

	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated"`

	// The direction of the SMS Command
	Direction *SupersimV1SmsCommandDirection `json:"direction"`

	// The message body of the SMS Command sent to or from the SIM
	Payload *string `json:"payload"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// The SID of the SIM that this SMS Command was sent to or from
	SimSid *string `json:"sim_sid"`

	// The status of the SMS Command
	Status *SupersimV1SmsCommandStatus `json:"status"`

	// The absolute URL of the SMS Command resource
	Url *string `json:"url"`
}

// The direction of the SMS Command
type SupersimV1SmsCommandDirection string

// The status of the SMS Command
type SupersimV1SmsCommandStatus string

// SupersimV1UsageRecord defines model for supersim.v1.usage_record.
type SupersimV1UsageRecord struct {
	// The SID of the Account that incurred the usage.
	AccountSid *string `json:"account_sid"`

	// Total data downloaded in bytes, aggregated by the query parameters.
	DataDownload *int `json:"data_download"`

	// Total of data_upload and data_download.
	DataTotal *int `json:"data_total"`

	// Total data uploaded in bytes, aggregated by the query parameters.
	DataUpload *int `json:"data_upload"`

	// SID of the Fleet resource on which the usage occurred.
	FleetSid *string `json:"fleet_sid"`

	// Alpha-2 ISO Country Code of the country the usage occurred in.
	IsoCountry *string `json:"iso_country"`

	// SID of the Network resource on which the usage occurred.
	NetworkSid *string `json:"network_sid"`

	// The time period for which the usage is reported.
	Period *map[string]interface{} `json:"period"`

	// SID of a Sim resource to which the UsageRecord belongs.
	SimSid *string `json:"sim_sid"`
}

// ListCommandParams defines parameters for ListCommand.
type ListCommandParams struct {
	// The SID or unique name of the Sim that Command was sent to or from.
	Sim *string `json:"Sim,omitempty"`

	// The status of the Command. Can be: `queued`, `sent`, `delivered`, `received` or `failed`. See the [Command Status Values](https://www.twilio.com/docs/wireless/api/command-resource#status-values) for a description of each.
	Status *ListCommandParamsStatus `json:"Status,omitempty"`

	// The direction of the Command. Can be `to_sim` or `from_sim`. The value of `to_sim` is synonymous with the term `mobile terminated`, and `from_sim` is synonymous with the term `mobile originated`.
	Direction *ListCommandParamsDirection `json:"Direction,omitempty"`

	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListCommandParamsStatus defines parameters for ListCommand.
type ListCommandParamsStatus string

// ListCommandParamsDirection defines parameters for ListCommand.
type ListCommandParamsDirection string

// ListEsimProfileParams defines parameters for ListEsimProfile.
type ListEsimProfileParams struct {
	// List the eSIM Profiles that have been associated with an EId.
	Eid *string `json:"Eid,omitempty"`

	// Find the eSIM Profile resource related to a [Sim](https://www.twilio.com/docs/wireless/api/sim-resource) resource by providing the SIM SID. Will always return an array with either 1 or 0 records.
	SimSid *string `json:"SimSid,omitempty"`

	// List the eSIM Profiles that are in a given status.
	Status *ListEsimProfileParamsStatus `json:"Status,omitempty"`

	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListEsimProfileParamsStatus defines parameters for ListEsimProfile.
type ListEsimProfileParamsStatus string

// ListFleetParams defines parameters for ListFleet.
type ListFleetParams struct {
	// The SID or unique name of the Network Access Profile that controls which cellular networks the Fleet's SIMs can connect to.
	NetworkAccessProfile *string `json:"NetworkAccessProfile,omitempty"`

	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListIpCommandParams defines parameters for ListIpCommand.
type ListIpCommandParams struct {
	// The SID or unique name of the Sim resource that IP Command was sent to or from.
	Sim *string `json:"Sim,omitempty"`

	// The ICCID of the Sim resource that IP Command was sent to or from.
	SimIccid *string `json:"SimIccid,omitempty"`

	// The status of the IP Command. Can be: `queued`, `sent`, `received` or `failed`. See the [IP Command Status Values](https://www.twilio.com/docs/wireless/api/ipcommand-resource#status-values) for a description of each.
	Status *ListIpCommandParamsStatus `json:"Status,omitempty"`

	// The direction of the IP Command. Can be `to_sim` or `from_sim`. The value of `to_sim` is synonymous with the term `mobile terminated`, and `from_sim` is synonymous with the term `mobile originated`.
	Direction *ListIpCommandParamsDirection `json:"Direction,omitempty"`

	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListIpCommandParamsStatus defines parameters for ListIpCommand.
type ListIpCommandParamsStatus string

// ListIpCommandParamsDirection defines parameters for ListIpCommand.
type ListIpCommandParamsDirection string

// ListNetworkAccessProfileParams defines parameters for ListNetworkAccessProfile.
type ListNetworkAccessProfileParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListNetworkAccessProfileNetworkParams defines parameters for ListNetworkAccessProfileNetwork.
type ListNetworkAccessProfileNetworkParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListNetworkParams defines parameters for ListNetwork.
type ListNetworkParams struct {
	// The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the Network resources to read.
	IsoCountry *string `json:"IsoCountry,omitempty"`

	// The 'mobile country code' of a country. Network resources with this `mcc` in their `identifiers` will be read.
	Mcc *string `json:"Mcc,omitempty"`

	// The 'mobile network code' of a mobile operator network. Network resources with this `mnc` in their `identifiers` will be read.
	Mnc *string `json:"Mnc,omitempty"`

	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListSimParams defines parameters for ListSim.
type ListSimParams struct {
	// The status of the Sim resources to read. Can be `new`, `ready`, `active`, `inactive`, or `scheduled`.
	Status *ListSimParamsStatus `json:"Status,omitempty"`

	// The SID or unique name of the Fleet to which a list of Sims are assigned.
	Fleet *string `json:"Fleet,omitempty"`

	// The [ICCID](https://en.wikipedia.org/wiki/Subscriber_identity_module#ICCID) associated with a Super SIM to filter the list by. Passing this parameter will always return a list containing zero or one SIMs.
	Iccid *string `json:"Iccid,omitempty"`

	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListSimParamsStatus defines parameters for ListSim.
type ListSimParamsStatus string

// ListBillingPeriodParams defines parameters for ListBillingPeriod.
type ListBillingPeriodParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListSmsCommandParams defines parameters for ListSmsCommand.
type ListSmsCommandParams struct {
	// The SID or unique name of the Sim resource that SMS Command was sent to or from.
	Sim *string `json:"Sim,omitempty"`

	// The status of the SMS Command. Can be: `queued`, `sent`, `delivered`, `received` or `failed`. See the [SMS Command Status Values](https://www.twilio.com/docs/wireless/api/smscommand-resource#status-values) for a description of each.
	Status *ListSmsCommandParamsStatus `json:"Status,omitempty"`

	// The direction of the SMS Command. Can be `to_sim` or `from_sim`. The value of `to_sim` is synonymous with the term `mobile terminated`, and `from_sim` is synonymous with the term `mobile originated`.
	Direction *ListSmsCommandParamsDirection `json:"Direction,omitempty"`

	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListSmsCommandParamsStatus defines parameters for ListSmsCommand.
type ListSmsCommandParamsStatus string

// ListSmsCommandParamsDirection defines parameters for ListSmsCommand.
type ListSmsCommandParamsDirection string

// ListUsageRecordParams defines parameters for ListUsageRecord.
type ListUsageRecordParams struct {
	// SID or unique name of a Sim resource. Only show UsageRecords representing usage incurred by this Super SIM.
	Sim *string `json:"Sim,omitempty"`

	// SID or unique name of a Fleet resource. Only show UsageRecords representing usage for Super SIMs belonging to this Fleet resource at the time the usage occurred.
	Fleet *string `json:"Fleet,omitempty"`

	// SID of a Network resource. Only show UsageRecords representing usage on this network.
	Network *string `json:"Network,omitempty"`

	// Alpha-2 ISO Country Code. Only show UsageRecords representing usage in this country.
	IsoCountry *string `json:"IsoCountry,omitempty"`

	// Dimension over which to aggregate usage records. Can be: `sim`, `fleet`, `network`, `isoCountry`. Default is to not aggregate across any of these dimensions, UsageRecords will be aggregated into the time buckets described by the `Granularity` parameter.
	Group *ListUsageRecordParamsGroup `json:"Group,omitempty"`

	// Time-based grouping that UsageRecords should be aggregated by. Can be: `hour`, `day`, or `all`. Default is `all`. `all` returns one UsageRecord that describes the usage for the entire period.
	Granularity *ListUsageRecordParamsGranularity `json:"Granularity,omitempty"`

	// Only include usage that occurred at or after this time, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Default is one month before the `end_time`.
	StartTime *time.Time `json:"StartTime,omitempty"`

	// Only include usage that occurred before this time (exclusive), specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Default is the current time.
	EndTime *time.Time `json:"EndTime,omitempty"`

	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

// ListUsageRecordParamsGroup defines parameters for ListUsageRecord.
type ListUsageRecordParamsGroup string

// ListUsageRecordParamsGranularity defines parameters for ListUsageRecord.
type ListUsageRecordParamsGranularity string
