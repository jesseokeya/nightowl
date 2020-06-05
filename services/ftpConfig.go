// Package services helps to create ftp configuration file
package services

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
)

// Params is the ftp json object params
type Params struct {
	BasePath string `json:"basePath"`
}

// PortRange is the ftp json object config for port range
type PortRange struct {
	Start int32 `json:"start"`
	End   int32 `json:"end"`
}

// Access is the authentication fields for the ftp server json object config
type Access struct {
	User       string `json:"user"`
	Pass       string `json:"pass"`
	Parameters Params `json:"params"`
}

// FtpConfig is json ftp configuration
type FtpConfig struct {
	Version                  int32     `json:"version"`
	Accesses                 []Access  `json:"accesses"`
	PassiveTransferPortRange PortRange `json:"passive_transfer_port_range"`
}

// UserFTPConfig is user configuration override for ftp config file
type UserFTPConfig struct {
	User string
	Pass string
	Path string
}

var templ string = `{
  "version": 1,
  "accesses": [
    {
	  "user": "{{ (index .accesses 0).user }}",
      "pass": "{{ (index .accesses 0).pass }}",
      "fs": "os",
      "params": {
        "basePath": "{{ (index .accesses 0).params.basePath }}"
      }
    }
  ],
  "passive_transfer_port_range": {
    "start": 2122,
    "end": 2130
  }
}`

// CreateFTPConfiguration created ftp configuration file
func (c *UserFTPConfig) CreateFTPConfiguration() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	configFilePath := path + "/ftpserver.json"
	fmt.Println(configFilePath)
	ftpConfig := &FtpConfig{
		Version: 1,
		Accesses: []Access{
			Access{
				User: c.User,
				Pass: c.Pass,
				Parameters: Params{
					BasePath: c.Path,
				},
			},
		},
		PassiveTransferPortRange: PortRange{
			Start: 2122,
			End:   2130,
		},
	}
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		log.Printf("config file %s does exist", configFilePath)
	} else {
		temp := template.Must(template.New("").Parse(templ))
		config, err := json.Marshal(ftpConfig)
		if err != nil {
			panic(err)
		}

		reference := map[string]interface{}{}
		if err := json.Unmarshal(config, &reference); err != nil {
			panic(err)
		}

		createdFile, err := os.Create(configFilePath)
		if err != nil {
			panic(err)
		}

		if err := temp.Execute(createdFile, reference); err != nil {
			panic(err)
		}

		defer createdFile.Close()
	}
}