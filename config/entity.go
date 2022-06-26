/*
Copyright © 2022 BROCHIER MAXENCE maxence@brochier.xyz

*/
package config

type JsonConfig struct {
	Key       string `json:"key"`
	Hash      string `json:"hash"`
	MasterUrl string `json:"master-url"`
}

type Config interface {
	Load()
}
