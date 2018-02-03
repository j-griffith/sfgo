/*
Copyright 2018 John Griffith

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

package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/j-griffith/sfgo/pkg/provider"
)

type config struct {
	tenantName       string
	endPoint         string
	defaultVolSz     int64 //Default volume size in GiB
	svip             string
	initiatorIFace   string //iface to use of iSCSI initiator on this node
	accessGroups     []int64
	useCHAP          bool   //you can use CHAP or AccessGroups, CHAP is the default
	defaultBlockSize int64  //blocksize to use on create when not specified  (512|4096, 512 is default)
	apiVersion       string //API version of the SF endpoint to use, default is "9.0"
}

// SolidFire represents a minimal client for issuing API requests
type SolidFire struct {
	Endpoint   string
	APIVersion string
	TenantID   int64
	Transport  *provider.HTTP
	Cfg        *config
}

var SFClient = SolidFire{}

func processConfig(c string) (error, config) {
	content, err := ioutil.ReadFile(c)
	if err != nil {
		log.Fatal("error processing config file: ", err)
	}

	var conf config
	err = json.Unmarshal(content, &conf)
	if err != nil {
		log.Fatal("error parsing config file: ", err)
	}
	return err, conf
}

func init() {
	cfgFile := os.Getenv("SF_CONFIG_FILE")
	InitClient(cfgFile)
}

func GetSolidFireClient() SolidFire {
	return SFClient
}

// InitClient is a callable initializer with specified config file.  We have a default init that runs on startup, but we provide this for any overrides/restarts
func InitClient(cfgFile string) {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(cfgFile) == 0 {
		cfgFile = "/var/solidfire-csi/solidfire.json"
	}
	_, cfg := processConfig(cfgFile)
	SFClient.Endpoint = cfg.endPoint
	SFClient.APIVersion = cfg.apiVersion
	SFClient.TenantID = 1
	SFClient.Cfg = &cfg

	// Transport is broken out only to allow mocking
	// we'll always use provider.HTTP currently
	SFClient.Transport = &provider.HTTP{}
}

func (sf *SolidFire) ClientRequest(mName string, params interface{}, resType interface{}) (interface{}, error) {
	req := provider.Request{}
	req.Name = mName
	req.URL = sf.Endpoint
	req.ID = provider.NewReqID()

	body, err := provider.IssueRequest(req, sf.Transport)
	if err != nil {
		log.Println("Damn!")
	}
	return provider.DecodeResponse(body, resType)
}
