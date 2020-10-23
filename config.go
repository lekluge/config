package config

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"strings"
)

var (
	ErrDoesNotExist = errors.New("File does not exist")
	ErrEmptyArray   = errors.New("Empty file array, make sure to SetFile first!")
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
		return nil, ErrDoesNotExist
	}
	return config{path: f, file: file}, nil
}

func (c config) Parse() (map[string]string, error) {
	if c.file == nil {
		return nil, ErrEmptyArray
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
