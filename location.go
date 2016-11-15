package foreman

/*
import (
	"net/http"
)
*/

// Location definition
type Location struct {
	Id                   int    `json:"id,omitempty"`
	Name                 string `json:"name,omitempty"`
	Description          string `json:"description,omitempty"`
	User_ids             []int  `json:"user_ids,omitempty"`
	Smart_proxy_ids      []int  `json:"smart_proxy_ids,omitempty"`
	Compute_resource_ids []int  `json:"compute_resource_ids,omitempty"`
	Media_ids            []int  `json:"media_ids,omitempty"`
	Config_template_ids  []int  `json:"config_template_ids,omitempty"`
	Domain_ids           []int  `json:"domain_ids,omitempty"`
	Realm_ids            []int  `json:"realm_ids,omitempty"`
	Hostgroup_ids        []int  `json:"hostgroup_ids,omitempty"`
	Environment_ids      []int  `json:"environment_ids,omitempty"`
	Subnet_ids           []int  `json:"subnet_ids,omitempty"`
}
