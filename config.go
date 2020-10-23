package config

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	ErrDoesNotExist = "File does not exist"
	ErrEmptyArray   = "Empty file array, make sure to SetFile first!"
)

type (
	//Config is the main config interface
	Config interface {
		Parse() (map[string]string, error)
		Raw() string
	}
	config struct {
		path string
		file []byte
	}
)

//Raw string of selected config file
func (c config) Raw() string {
	return string(c.file)
}

//SetFile inits the file
func SetFile(f string) (Config, error) {
	var err error
	file, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf(ErrDoesNotExist)
	}
	return config{path: f, file: file}, nil
}

func (c config) Parse() (map[string]string, error) {
	if c.file == nil {
		return nil, fmt.Errorf(ErrEmptyArray)
	}
	r := bytes.NewReader(c.file)
	scanner := bufio.NewScanner(r)
	data := make(map[string]string)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "[") {
			continue
		}
		spl := strings.Split(scanner.Text(), "=")
		for i := range spl {
			spl[i] = strings.TrimSpace(spl[i])
		}
		data[spl[0]] = spl[1]
	}
	return data, nil
}
