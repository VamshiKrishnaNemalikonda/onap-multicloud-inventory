/*
Copyright 2020  Tech Mahindra.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	con "github.com/onap/multicloud-k8s/src/inventory/constants"
	log "github.com/onap/multicloud-k8s/src/inventory/logutils"
	utils "github.com/onap/multicloud-k8s/src/inventory/utils"
	//k8sint "github.com/onap/multicloud-k8s/src/k8splugin/internal/app"
	//k8scon "github.com/onap/multicloud-k8s/src/k8splugin/internal/connection"

        pkgerrors "github.com/pkg/errors"
	"encoding/json"
	"net/http"
	"os"
        "fmt"
        "crypto/tls"
)

func ListInstances() ([]string, error) {

	MK8S_URI := "https://10.211.1.20" //os.Getenv("onap-multicloud-k8s")
	MK8S_Port := "30283" //os.Getenv("multicloud-k8s-port")

	instanceapi := MK8S_URI + ":" + MK8S_Port + con.INSTANCE_EP
        fmt.Println(instanceapi)
	req, err := http.NewRequest(http.MethodGet, instanceapi, nil)
	if err != nil {

		log.Error("Something went wrong while listing resources - contructing request")
                return []string{}, pkgerrors.New("Something went wrong while listing resources - contructing request")
	}

	client := http.DefaultClient

        //Disabling security checks
        http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	res, err := client.Do(req)

	if err != nil {
		log.Error("Something went wrong while listing resources - executing request")
                fmt.Println(err)
                return []string{}, pkgerrors.New("Something went wrong while listing resources - executing request")
	}

	//defer res.Body.Close()
        fmt.Println("ListInstances : parsing the /instance api response")
	decoder := json.NewDecoder(res.Body)
	var rlist []con.InstanceMiniResponse
	err = decoder.Decode(&rlist)

        resourceList := utils.ParseListInstanceResponse(rlist)

	return resourceList, nil

}

func GetConnection(cregion string) (con.Connection, error) {

        fmt.Println("GetConnection: started")

	MK8S_URI := "https://10.211.1.20" //os.Getenv("onap-multicloud-k8s")
	MK8S_Port := "30283" //os.Getenv("multicloud-k8s-port")

	connection_url := MK8S_URI + ":" + MK8S_Port + con.CONNECTION_EP + "/"+  cregion

        fmt.Println(connection_url)
	req, err := http.NewRequest(http.MethodGet, connection_url, nil)
	if err != nil {

		log.Error("Something went wrong while getting Connection resource - contructing request")
                return con.Connection{}, pkgerrors.New("Something went wrong while listing resources - executing request")
	}

	client := http.DefaultClient
	res, err := client.Do(req)

	if err != nil {
		log.Error("Something went wrong while getting Connection resource - executing request")
                return con.Connection{}, pkgerrors.New("Something went wrong while listing resources - executing request")
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var connection con.Connection
	err = decoder.Decode(&connection)

	return connection, nil

}

func CheckStatusForEachInstance(instanceID string) (con.InstanceStatus, error) {

	MK8S_URI := os.Getenv("onap-multicloud-k8s")
	MK8S_Port := os.Getenv("multicloud-k8s-port")

	instancelist := MK8S_URI + ":" + MK8S_Port + con.INSTANCE_EP + "/" +instanceID + "/status"

	req, err := http.NewRequest(http.MethodGet, instancelist, nil)
	if err != nil {
		log.Error("Error while checking instance status - building http request")
                return con.InstanceStatus{}, pkgerrors.New("Something went wrong while listing resources - executing request")
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {

		log.Error("Error while checking instance status - making rest request")
                return con.InstanceStatus{}, pkgerrors.New("Something went wrong while listing resources - executing request")
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var instStatus con.InstanceStatus
	err = decoder.Decode(&instStatus)

	return instStatus, nil
}

func DummyStatusResponse(instanceID string) (con.DummyStatus, error) {

	dummyres := `{
		"id": "zen_goldberg",
		"request": {
			"rb-name": "testing",
			"rb-version": "1",
			"profile-name": "testing_profile",
			"cloud-region": "k8sregion",
			"labels": null,
			"override-values": null
		},
		"namespace": "default",
		"clusterip": "10.211.1.20",
		"port": "677"
	}`

        var status con.DummyStatus

	json.Unmarshal([]byte(dummyres), &status)

       return status, nil

}
