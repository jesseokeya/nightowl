// Package services helps to create ftp configuration file
package services

import (
	"encoding/json"
	"html/template"
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

func handle(e error) {
	if e != nil {
		panic(e)
	}
}

func configTemplate() string {
	return `
	{
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
}

// CreateFTPConfiguration created ftp configuration file
func (c *UserFTPConfig) CreateFTPConfiguration() {
	path, err := os.Getwd()
	handle(err)

	configFilePath := path + "/ftpserver.json"
	ftpConfig := &FtpConfig{
		Version: 1,
		Accesses: []Access{
			{
				User: c.User,
				Pass: c.Pass,
				Parameters: Params{
					BasePath: path + os.Getenv("STORE"),
				},
			},
		},
		PassiveTransferPortRange: PortRange{
			Start: 2122,
			End:   2130,
		},
	}
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		temp := template.Must(template.New("").Parse(configTemplate()))
		config, err := json.Marshal(ftpConfig)
		handle(err)

		reference := map[string]interface{}{}
		err = json.Unmarshal(config, &reference)
		handle(err)

		createdFile, err := os.Create(configFilePath)
		handle(err)

		err = temp.Execute(createdFile, reference)
		handle(err)
	}
}
