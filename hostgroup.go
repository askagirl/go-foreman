package foreman

/*
import (
	"net/http"
)
*/

// Hostgroup definition
type Hostgroup struct {
	Id                 int    `json:"id,omitempty"`
	Name               string `json:"name,omitempty"`
	Parent_id          int    `json:"parent_id,omitempty"`
	Environment_id     int    `json:"environment_id,omitempty"`
	Operatingsystem_id int    `json:"operatingsystem_id,omitempty"`
	Architecture_id    int    `json:"architecture_id,omitempty"`
	Medium_id          int    `json:"medium_id,omitempty"`
	Ptable_id          int    `json:"ptable_id,omitempty"`
	Puppet_ca_proxy_id int    `json:"puppet_ca_proxy_id,omitempty"`
	Subnet_id          string `json:"subnet_id,omitempty"`
	Domain_id          int    `json:"domain_id,omitempty"`
	Realm_id           int    `json:"realm_id,omitempty"`
	Puppet_proxy_id    int    `json:"puppet_proxy_id,omitempty"`
	Location_ids       []int  `json:"location_ids,omitempty"`
	Organization_ids   []int  `json:"organization_ids,omitempty"`
}
