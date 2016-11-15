package foreman

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type Foreman struct {
	Hostname  string
	Username  string
	password  string
	VerifySSL bool
	Proxy     string
	BaseURL   string
	client    *http.Client
	auth      string
}

func Client(HostName string, UserName string, Password string, Insecure bool, Proxy string) (foreman *Foreman) {
	var tr *http.Transport

	foreman = new(Foreman)
	foreman.Hostname = HostName
	foreman.Username = UserName
	foreman.password = Password
	foreman.VerifySSL = Insecure
	envProxy, proxySet := os.LookupEnv("HTTP_PROXY")
	if Proxy != "" {
		// Set if Proxy passed explicitly
		foreman.Proxy = Proxy
	} else if proxySet {
		foreman.Proxy = envProxy
	}
	proxyURL, err := url.Parse(foreman.Proxy)
	if err != nil {
		log.Printf("Failed to parse Proxy URL: %s\n")
		panic("Invalid proxy defined!")
	}
	foreman.BaseURL = "https://" + foreman.Hostname + "/api/"
	foreman.auth = "Basic " + base64.StdEncoding.EncodeToString([]byte(UserName+":"+Password))

	if foreman.VerifySSL {
		tr = &http.Transport{}
	} else {
		log.Println("Disabling API SSL verification!")

		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	if foreman.Proxy != "" {
		tr.Proxy = http.ProxyURL(proxyURL)
	}
	foreman.client = &http.Client{
		Transport: tr,
		Timeout:   15 * time.Second,
	}
	if foreman.client != nil {
		log.Println("Returning Foreman Client")
		return foreman
	} else {
		panic("Failed to attain client!")
		return nil
	}
}

func (foreman *Foreman) Post(endpoint string, jsonData []byte) (map[string]interface{}, error) {
	var target string
	var data interface{}
	var req *http.Request

	target = foreman.BaseURL + endpoint
	log.Println("POST Form: " + target)
	log.Println("POST Data: " + string(jsonData))

	req, err := http.NewRequest("POST", target, bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(jsonData)))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", foreman.auth)
	r, err := foreman.client.Do(req)
	if err != nil {
		log.Println("Error while posting")
		log.Println(err)
		return nil, err
	} else {
		defer r.Body.Close()
	}

	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error while reading body")
		log.Println(err)
		return nil, err
	} else {
		log.Printf("Received: %s", response)
	}
	err = json.Unmarshal(response, &data)
	if err != nil {
		log.Println("Error while processing JSON")
		log.Println(err)
		return nil, err
	}
	if r.StatusCode < 200 || r.StatusCode > 299 {
		log.Printf("Error Body: %s", data)
		return nil, errors.New("HTTP Error " + r.Status)
	}
	m := data.(map[string]interface{})
	log.Printf("POST Data returned: %s", m)
	return m, nil
}

func (foreman *Foreman) Get(endpoint string) (map[string]interface{}, error) {
	var target string
	var data interface{}

	target = foreman.BaseURL + endpoint
	log.Printf("GET %s", target)
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		panic("Failed to create GET request")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", foreman.auth)
	log.Printf("Foreman Auth token: %s", foreman.auth)
	r, err := foreman.client.Do(req)
	if err != nil {
		fmt.Printf("Failed to connect to %s\n", target)
		log.Println(err)
		panic("Failed to connect to Foreman API")
		return nil, err
	} else {
		log.Printf("Recvd response from %s\n", target)
		defer r.Body.Close()
		log.Printf("Recvd: %s", r)
	}
	if r.StatusCode < 200 || r.StatusCode > 299 {
		return nil, errors.New("HTTP Error " + r.Status)
	}

	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error while reading body")
		log.Println(err)
		return nil, err
	}
	err = json.Unmarshal(response, &data)
	if err != nil {
		log.Println("Error while processing JSON")
		log.Println(err)
		return nil, err
	}
	m := data.(map[string]interface{})
	log.Printf("Foreman GET: %v", m)
	return m, nil
}

func (foreman *Foreman) Delete(endpoint string) (map[string]interface{}, error) {
	var target string
	var data interface{}
	var req *http.Request

	target = foreman.BaseURL + endpoint
	req, err := http.NewRequest("DELETE", target, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", foreman.auth)
	r, err := foreman.client.Do(req)
	defer r.Body.Close()
	if err != nil {
		fmt.Println("Error while deleting")
		fmt.Println(err)
		return nil, err
	}
	if r.StatusCode < 200 || r.StatusCode > 299 {
		return nil, errors.New("HTTP Error " + r.Status)
	}
	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error while reading body")
		fmt.Println(err)
		return nil, err
	}
	err = json.Unmarshal(response, &data)
	if err != nil {
		fmt.Println("Error while processing JSON")
		fmt.Println(err)
		return nil, err
	}
	m := data.(map[string]interface{})
	return m, nil
}

type HostMap map[string]Host

/*
func (foreman *Foreman) CreateHost(HostGroupId int, Name string, Mac string) (string, error) {
	var hostMap map[string]Host
	var err error

	hostMap = make(HostMap)
	hostMap["host"] = Host{
		Hostgroup_id: HostGroupId,
		Name:         Name,
		Mac:          Mac,
		Build:        true,
	}
	jsonText, err := json.Marshal(hostMap)
	data, err := foreman.Post("hosts", jsonText)
	if err != nil {
		fmt.Print("Error ")
		fmt.Println(err)
		return "", err
	}
	return strconv.FormatFloat(data["id"].(float64), 'f', 0, 64), nil
}
*/
func (foreman *Foreman) CreateHost(Host *Host) (returnedHost *Host, err error) {
	jsonText, err := json.Marshal(&Host)
	data, err := foreman.Post("hosts", jsonText)
	if err != nil {
		log.Print("Error ")
		log.Println(err)
		return nil, err
	} else {
		log.Printf("Created Host: %s", data)
	}
	err = mapstructure.Decode(data, &returnedHost)
	if err != nil {
		log.Print("Error unmarshalling Host struct after create")
		panic(err)
	}
	return returnedHost, nil
	//return strconv.FormatFloat(data["id"].(float64), 'f', 0, 64), nil

}

func (foreman *Foreman) GetHost(Host *Host) (returnedHost *Host, err error) {
	fmt.Printf("Get host %s", Host.Name)
	//uri := fmt.Sprintf("hosts/%s.%s", Host.Name, Host.Domain.Name)
	uri := fmt.Sprintf("hosts/%s", Host.Name)
	data, err := foreman.Get(uri)
	if err != nil { //error
		fmt.Print("Error ")
		fmt.Println(err)
		return nil, err
	} else {
		// Populate the Host struct and return it
		fmt.Printf("GetHost returned: %s\n", data)
		return returnedHost, nil
		//return strconv.FormatFloat(data["id"].(float64), 'f', 0, 64), nil
	}
	return nil, nil
}

func (foreman *Foreman) UpdateHost(Host *Host) (returnedHost *Host, err error) {
	fmt.Printf("Update host %s", Host.Name)
	/* TODO
	domain	:=
	uri := fmt.Sprintf("hosts/%s.%s", Host.Name, Host.Domain.Name)
	data, err := foreman.Get("hosts")
	if err != nil { //error
		fmt.Print("Error ")
		fmt.Println(err)
		return nil, err
	} else {
		//TODO
		return returnedHost
		//return strconv.FormatFloat(data["id"].(float64), 'f', 0, 64), nil
	}
	*/
	return nil, nil
}

func (foreman *Foreman) DeleteHost(Host *Host) error {
	var err error

	_, err = foreman.Delete("hosts/" + Host.Name)
	return err
}
