package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"reflect"
)

var (
	config map[string]interface{}
)

//LoadConfig loads the config and prints an error if one is encountered
func LoadConfig() {
	stagedRoutedHostnames := make(map[string]int)
	stagedSites := []Site{}

	var err error
	var files []os.FileInfo

	if err = os.MkdirAll(*sitesLoc, 0600); err != nil {
		log.Print(err)
	}

	if files, err = ioutil.ReadDir(*sitesLoc); err != nil {
		log.Print(err)
		return
	}

	var configFiles []os.FileInfo

	//Filter out non json files
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if path.Ext(f.Name()) != ".json" {
			continue
		}
		configFiles = append(configFiles, f)
	}

	//Process each site
	for _, f := range configFiles {
		var dat []byte
		var err error

		siteConf := make(map[string]interface{})

		siteIndex := len(stagedSites)
		stagedSites = append(stagedSites, Site{})

		if dat, err = ioutil.ReadFile(*sitesLoc + "/" + f.Name()); err != nil {
			log.Print(err)
		}

		if err = json.Unmarshal(dat, &siteConf); err != nil {
			log.Print(err)
		}

		switch siteConf["hostnames"].(type) {
		case []interface{}:
			for _, hostname := range siteConf["hostnames"].([]interface{}) {
				switch hostname.(type) {
				case string:
					log.Print("Adding hostname -> ", hostname)
					stagedRoutedHostnames[hostname.(string)] = siteIndex
				default:
					log.Print("Expected string but got ", reflect.TypeOf(hostname), " while parsing hostname in ", f.Name())
				}
			}
		default:
			log.Print("Expected array but got ", reflect.TypeOf(siteConf["hostnames"]), " while parsing hosts in ", f.Name())
		}

		switch siteConf["backends"].(type) {
		case []interface{}:
			for _, backend := range siteConf["backends"].([]interface{}) {
				switch backend.(type) {
				case string:
					log.Print("Adding backend -> ", backend)
					stagedSites[siteIndex].Backends = append(stagedSites[siteIndex].Backends, backend.(string))
				default:
					log.Print("Expected string but got ", reflect.TypeOf(backend), " while parsing backend in ", f.Name())
				}
			}
		default:
			log.Print("Expected array but got ", reflect.TypeOf(siteConf["backends"]), " while parsing backends in ", f.Name())
		}

		switch siteConf["secret"].(type) {
		case string:
			stagedSites[siteIndex].Secret = siteConf["secret"].(string)
		default:
			log.Print("Expected string but got ", reflect.TypeOf(siteConf["secret"]), " while parsing secret in ", f.Name())
		}

		routedHostnames = stagedRoutedHostnames
		Sites = stagedSites
	}
}
