package foreman

/*
import (
	"net/http"
)
*/

// Domain definition
type Domain struct {
	id       int    `json:"id,omitempty"`
	name     string `json:"name,omitempty"`
	fullname string `json:"fullname,omitempty"`
	dns_id   int    `json:"dns_id,omitempty"`
	//domain_parameters_attributes   []interface{}    `json:"dns_id,omitempty"`
	location_ids     []int `json:"location_ids,omitempty"`
	organization_ids []int `json:"organization_ids,omitempty"`
}
