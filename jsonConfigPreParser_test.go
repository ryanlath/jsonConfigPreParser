package jsonConfigPreParser

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

type Config struct {
	AppName string
	Debug   bool
	Server  struct {
		Port            int
		Host            string
		StaticDirectory string
		CacheDirectory  string
	}
	Database struct {
		MariaDNS string
	}
	Log struct {
		Level string
		Color bool
	}
}

var configFile = "test/json.config"

var expected string = `{"Debug":true,"Server":{"Port":6969,"Host":"localhost","StaticDirectory":"www"},"Database":{"MariaDNS":"user:password@/db?parseTime=true"},"Log":{"Level":"debug","Color":true,"FakeTest":"#EEccEE","#Fake3":" /*whatever! ","Fake4":"\\Not a comment"}}`

var output string

func TestComplete(t *testing.T) {

	var bytes []byte
	var bytesMin []byte
	var err error
	var config *Config

	if bytes, err = ioutil.ReadFile(configFile); err == nil {
		if bytesMin, err = Parse(bytes); err == nil {
			output = string(bytesMin)
			if err = json.Unmarshal(bytesMin, &config); err == nil {
			}
		}
	}

	if err != nil {
		t.Errorf("Test failed with error:\n%#v", err)
	}

	if expected != output {
		t.Errorf("Output not as expected: \n%#v\n-------------\n%#v", output, expected)
	}
}
