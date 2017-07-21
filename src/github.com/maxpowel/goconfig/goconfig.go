package goconfig


import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"regexp"
	"bytes"
	"text/template"
	"fmt"
	"log"
	"reflect"
	"errors"
	"strings"
	"github.com/fatih/color"
)

type ConfigurationParameters struct {
	Key string
	Value    string
}

type Config struct {
	Mapping map[string]interface{}
	ConfigFilePath string
	ParametersFilePath string
}

func (c *Config) Load() {
	_, err := load(c.ConfigFilePath, c.ParametersFilePath, c.Mapping)
	if err != nil{
		log.Fatal(err)
	}
}


func NewConfig(configPath, parametersPath string) *Config {
	c := Config{Mapping: make(map[string]interface{})}
	c.ConfigFilePath = configPath
	c.ParametersFilePath = parametersPath
	return &c
}

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

func FillStruct(s interface{}, m map[interface{}]interface{}) error {
	for k, v := range m {
		err := SetField(s, strings.Title(k.(string)), v)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadParameters(path string) (map[string]interface{}, error){

	var m map[string]interface{}
	m = make(map[string]interface{})

	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(fileContent, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func loadConfig(path string, parameters map[string]interface{}) (bytes.Buffer, error) {
	// Extract all variables from config file
	var tpl bytes.Buffer
	r := regexp.MustCompile(`\{\{([a-zA-Z0-9_]+)}}`)
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		return tpl, err
	}

	matches := r.FindAllStringSubmatch(string(configFile), -1)
	for _, match := range matches {
		_, ok := parameters[match[1]]
		if !ok {
			return tpl, fmt.Errorf("Parameter %v not found", match[1])
			//log.Println("Un campo no esta")
			//log.Println(match[1])
		} /*else {
			log.Println("OTODO OK")
			log.Println(match[1])
			log.Println(res)
		}*/

	}

	// Prepare the configuration file to be used with template. Basically convert {{var}} for {{.var}}
	re := regexp.MustCompile(`\{\{([a-zA-Z0-9_]+)}}`)
	configForTemplate := re.ReplaceAllString(string(configFile), `{{.$1}}`)

	configTemplate, templateErr := template.New("test").Parse(configForTemplate)
	if templateErr != nil { return tpl, templateErr }

	templateErr = configTemplate.Execute(&tpl, parameters)
	if templateErr != nil { return tpl, templateErr }

	return tpl, nil
}

func load(configurationPath string, parametersPath string, mapping map[string]interface{}) (map[string]interface{}, error){
	parameters, err := loadParameters(parametersPath)
	if err != nil {
		return nil, err
	}

	config, err := loadConfig(configurationPath, parameters)
	if err != nil {
		return nil, err
	}


	var conf map[string]interface{}
	err = yaml.Unmarshal(config.Bytes(), &conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for k, v := range conf {

		var moduleConf, moduleErr = mapping[k]
		if !moduleErr {
			color.Yellow("Ignorando configuracion de %v", k)
			//return nil, err
		} else {

			FillStruct(moduleConf, v.(map[interface{}]interface{}))
		}
	}

	return conf, nil

}
