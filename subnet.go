package foreman

/*
import (
	"net/http"
)
*/

// Subnet definition
type Subnet struct {
	Id               int    `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	Network          string `json:"network,omitempty"`
	Mask             string `json:"mask,omitempty"`
	Gateway          string `json:"gateway,omitempty"`
	Dns_primary      string `json:"dns_primary,omitempty"`
	Ipam             string `json:"ipam,omitempty"`
	From             string `json:"from,omitempty"`
	To               string `json:"to,omitempty"`
	Vlanid           string `json:"vlanid,omitempty"`
	Domain_ids       []int  `json:"domain_ids,omitempty"`
	Dhcp_id          int    `json:"dhcp_id,omitempty"`
	Tftp_id          int    `json:"tftp_id,omitempty"`
	Dns_id           int    `json:"dns_id,omitempty"`
	Boot_mode        string `json:"boot_mode,omitempty"`
	Location_ids     []int  `json:"location_ids,omitempty"`
	Organization_ids []int  `json:"organization_ids,omitempty"`
}
