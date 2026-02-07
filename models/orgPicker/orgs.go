// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    orgs, err := UnmarshalOrgs(bytes)
//    bytes, err = orgs.Marshal()

package orgPicker

import "time"

import "encoding/json"

func UnmarshalOrgs(data []byte) (Orgs, error) {
	var r Orgs
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Orgs) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Orgs struct {
	Status   int64         `json:"status"`
	Result   Result        `json:"result"`
	Warnings []any         `json:"warnings"`
}

type Result struct {
	Other          []OrgDescriptor         `json:"other"`
	Sandboxes      []OrgDescriptor         `json:"sandboxes"`
	NonScratchOrgs []OrgDescriptor         `json:"nonScratchOrgs"`
	DevHubs        []OrgDescriptor         `json:"devHubs"`
	ScratchOrgs    []OrgDescriptor         `json:"scratchOrgs"`
}

type OrgDescriptor struct {
	AccessToken                     string      `json:"accessToken"`
	InstanceURL                     string      `json:"instanceUrl"`
	OrgID                           string      `json:"orgId"`
	Username                        string      `json:"username"`
	LoginURL                        string      `json:"loginUrl"`
	ClientID                        string      `json:"clientId"`
	IsDevHub                        bool        `json:"isDevHub"`
	InstanceAPIVersion              string      `json:"instanceApiVersion"`
	InstanceAPIVersionLastRetrieved string      `json:"instanceApiVersionLastRetrieved"`
	Alias                           string      `json:"alias"`
	IsDefaultDevHubUsername         bool        `json:"isDefaultDevHubUsername"`
	IsDefaultUsername               bool        `json:"isDefaultUsername"`
	LastUsed                        time.Time   `json:"lastUsed"`
	ConnectedStatus                 string      `json:"connectedStatus"`
	Name                            *string     `json:"name,omitempty"`
	InstanceName                    *string     `json:"instanceName,omitempty"`
	NamespacePrefix                 *any 				`json:"namespacePrefix"`
	IsSandbox                       *bool       `json:"isSandbox,omitempty"`
	IsScratch                       *bool       `json:"isScratch,omitempty"`
	TrailExpirationDate             *any 				`json:"trailExpirationDate"`
	TracksSource                    *bool       `json:"tracksSource,omitempty"`
	DefaultMarker                   *string     `json:"defaultMarker,omitempty"`
}
