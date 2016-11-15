package foreman

/*
import (
	"net/http"
)
*/

// Operatingsystem definition
type Operatingsystem struct {
	Id                  int    `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	Major               string `json:"major,omitempty"`
	Minor               string `json:"minor,omitempty"`
	Description         string `json:"description,omitempty"`
	Family              string `json:"family,omitempty"`
	Release_name        string `json:"release_name,omitempty"`
	Password_hash       string `json:"password_hash,omitempty"`
	Architecture_ids    []int  `json:"architecture_ids,omitempty"`
	Config_template_ids []int  `json:"config_template_ids,omitempty"`
	Medium_ids          []int  `json:"medium_ids,omitempty"`
	Ptable_ids          []int  `json:"ptable_ids,omitempty"`
}
