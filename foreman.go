package foreman

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Foreman struct {
	Hostname  string
	Username  string
	password  string
	VerifySSL bool
	BaseURL   string
	client    *http.Client
	auth      string
}

func Client(HostName string, UserName string, Password string) (foreman *Foreman) {
	var tr *http.Transport

	foreman = new(Foreman)
	foreman.Hostname = HostName
	foreman.Username = UserName
	foreman.password = Password
	foreman.VerifySSL = false
	foreman.BaseURL = "https://" + foreman.Hostname + "/api/"
	foreman.auth = "Basic " + base64.StdEncoding.EncodeToString([]byte(UserName+":"+Password))

	if foreman.VerifySSL {
		tr = &http.Transport{}
	} else {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	foreman.client = &http.Client{Transport: tr}

	return foreman
}

func (foreman *Foreman) Post(endpoint string, jsonData []byte) (map[string]interface{}, error) {
	var target string
	var data interface{}
	var req *http.Request

	target = foreman.BaseURL + endpoint
	//fmt.Println("POST form " + target)
	req, err := http.NewRequest("POST", target, bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(jsonData)))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", foreman.auth)
	r, err := foreman.client.Do(req)
	defer r.Body.Close()
	if err != nil {
		fmt.Println("Error while posting")
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

func (foreman *Foreman) Get(endpoint string) (map[string]interface{}, error) {
	var target string
	var data interface{}

	target = foreman.BaseURL + endpoint
	req, err := http.NewRequest("GET", target, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", foreman.auth)
	r, err := foreman.client.Do(req)
	defer r.Body.Close()
	if err != nil {
		fmt.Println("Error while getting")
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
		fmt.Print("Error ")
		fmt.Println(err)
		return nil, err
	} else {
		fmt.Printf("Created Host: %s", data)
	}
	return nil, nil
	//return strconv.FormatFloat(data["id"].(float64), 'f', 0, 64), nil

}

func (foreman *Foreman) GetHost(Host *Host) (returnedHost *Host, err error) {
	fmt.Printf("Get host %s", Host.Name)
	/* TODO
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
	} */
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

func (foreman *Foreman) DeleteHost(HostID string) error {
	var err error

	_, err = foreman.Delete("hosts/" + HostID)
	return err
}
