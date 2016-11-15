package foreman

/*
import (
	"net/http"
)
*/

// Environment definition
type Environment struct {
	id               int    `json:"id,omitempty"`
	name             string `json:"name,omitempty"`
	location_ids     []int  `json:"location_ids,omitempty"`
	organization_ids []int  `json:"organization_ids,omitempty"`
}
