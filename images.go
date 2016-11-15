package foreman

/*
import (
	"net/http"
)
*/

// Image definition
type Image struct {
	Id                  int    `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	Username            string `json:"username,omitempty"`
	Uuid                string `json:"uuid,omitempty"`
	Compute_resource_id string `json:"compute_resource_id,omitempty"`
	Architecture_id     string `json:"architecture_id,omitempty"`
	Operatingsystem_id  string `json:"operatingsystem_id,omitempty"`
}
