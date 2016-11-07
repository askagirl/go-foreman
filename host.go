package foreman

/*
import (
	"net/http"
)
*/

//These are actual compute instance attributes
type Compute_attributes struct {
	Cpus      string `json:"cpus,omitempty"`
	Start     string `json:"start,omitempty"`
	Cluster   string `json:"cluster,omitempty"`
	Memory_mb string `json:"memory_mb,omitempty"`
	Guest_id  string `json:"guest_id,omitempty"`
	//This needs to be a struct map or the foreman API will kick back the JSON
	Volumes_attributes_map map[string]Volumes_attributes `json:"volumes_attributes,omitempty"`
}

//These are things like which datastore or cluster the virtual disks need to live on
type Volumes_attributes struct {
	Name    string `json:"name,omitempty"`
	Size_gb int    `json:"size_gb,omitempty"`
	//_delete   string `json:",omitempty"`
	Datastore string `json:"datastore,omitempty"`
}

//setup param archetype for host params attributes
type Params_archetype struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

//interfaces_attributes parameters
type Interfaces_attributes struct {
	Mac              string   `json:"mac,omitempty"`
	Ip               string   `json:"ip,omitempty"`
	Type             string   `json:"type,omitempty"`
	Name             string   `json:"name,omitempty"`
	Subnet_id        int      `json:"subnet_id,omitempty"`
	Domain_id        int      `json:"domain_id,omitempty"`
	Identifier       string   `json:"identifier,omitempty"`
	Managed          bool     `json:"managed,omitempty"`
	Primary          bool     `json:"primary,omitempty"`
	Provision        bool     `json:"provision,omitempty"`
	Username         string   `json:"username,omitempty"` //only for bmc
	Password         string   `json:"password,omitempty"` //only for bmc
	Provider         string   `json:"provider,omitempty"` //only accepted IPMI
	Virtual          bool     `json:"virtual,omitempty"`
	Tag              string   `json:"tag,omitempty"`
	Attached_to      string   `json:"attached_to,omitempty"`
	Mode             string   `json:"mode,omitempty"` // with validations
	Attached_devices []string `json:"attached_devices,omitempty"`
	Bond_options     string   `json:"bond_options,omitempty"`
	//These are special attributes for hypervisor like vlan and whatnot
	Compute_attributes Ifcompute_attributes `json:"compute_attributes,omitempty"`
}

//struct for nested interface compute attributes
type Ifcompute_attributes struct {
	Network string `json:"network,omitempty"`
	Type    string `json:"type,omitempty"`
}

//This is the main host struct instance that later gets wrapped in reqHost for JSON/foreman API reasons
type Host struct {
	Name                string             `json:"name,omitempty"`
	Environment_id      string             `json:"environment_id,omitempty"`
	Ip                  string             `json:"ip,omitempty"`
	Mac                 string             `json:"mac,omitempty"`
	Architecture_id     int                `json:"architecture_id,omitempty"`
	Domain_id           int                `json:"domain_id,omitempty"`
	Realm_id            int                `json:"realm_id,omitempty"`
	Puppet_proxy_id     int                `json:"puppet_proxy_id,omitempty"`
	Puppetclass_ids     []int              `json:"puppetclass_ids,omitempty"`
	Operatingsystem_id  string             `json:"operatingsystem_id,omitempty"`
	Medium_id           string             `json:"medium_id,omitempty"`
	Ptable_id           int                `json:"ptable_id,omitempty"`
	Subnet_id           int                `json:"subnet_id,omitempty"`
	Compute_resource_id int                `json:"compute_resource_id,omitempty"`
	Root_pass           string             `json:"root_pass,omitempty"`
	Model_id            int                `json:"model_id,omitempty"`
	Hostgroup_id        int                `json:"hostgroup_id,omitempty"`
	Puppet_ca_proxy_id  int                `json:"puppet_ca_proxy_id,omitempty"`
	Image_id            int                `json:"image_id,omitempty"`
	Build               bool               `json:"build,omitempty"`
	Enabled             bool               `json:"enabled,omitempty"`
	Provision_method    string             `json:"provision_method,omitempty"`
	Managed             bool               `json:"managed,omitempty"`
	Compute_attributes  Compute_attributes `json:"compute_attributes,omitempty"`
	Owner_id            int                `json:"owner_id,omitempty"`
	Owner_type          string             `json:"owner_type,omitempty"` // must be either User or Usergroup
	Progress_report_id  string             `json:"progress_report_id,omitempty"`
	Comment             string             `json:"comment,omitempty"`
	Capabilities        string             `json:"capabilities,omitempty"`
	Compute_profile_id  int                `json:"compute_profile_id,omitempty"`
	//mapped struct array for host parameters
	Host_parameters_attributes_map map[string]Params_archetype `json:"host_parameters_attributes,omitempty"`
	//struct array for multiple interfaces
	Interfaces_attributes_array []Interfaces_attributes `json:"interfaces_attributes,omitempty"`
}
