package foreman

/*
import (
	"net/http"
)
*/

// Media definition
type Media struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
	// Os family valid values: AIX, Altlinux, Archlinux, Coreos, Debian, Freebsd, Gentoo, Junos, Redhat, Solaris, Suse, Windows
	Os_family           string `json:"os_family,omitempty"`
	Operatingsystem_ids []int  `json:"operatingsystem_ids,omitempty"`
	Location_ids        []int  `json:"location_ids,omitempty"`
	Organization_ids    []int  `json:"organization_ids,omitempty"`
}
