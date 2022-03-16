// Package frontline provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package frontline

const (
	AccountSid_authTokenScopes = "accountSid_authToken.Scopes"
)

// Defines values for FrontlineV1UserState.
const (
	FrontlineV1UserStateActive      FrontlineV1UserState = "active"
	FrontlineV1UserStateDeactivated FrontlineV1UserState = "deactivated"
)

// FrontlineV1User defines model for frontline.v1.user.
type FrontlineV1User struct {
	// The avatar URL which will be shown in Frontline application
	Avatar *string `json:"avatar"`

	// The string that you assigned to describe the User
	FriendlyName *string `json:"friendly_name"`

	// The string that identifies the resource's User
	Identity *string `json:"identity"`

	// Whether the User is available for new conversations
	IsAvailable *bool `json:"is_available"`

	// The unique string that identifies the resource
	Sid *string `json:"sid"`

	// Current state of this user
	State *FrontlineV1UserState `json:"state"`

	// An absolute URL for this user.
	Url *string `json:"url"`
}

// Current state of this user
type FrontlineV1UserState string
