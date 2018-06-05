package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type EC2RouteTable struct {
	Type       string                  `yaml:"Type"`
	Properties EC2RouteTableProperties `yaml:"Properties"`
	Condition  interface{}             `yaml:"Condition,omitempty"`
	Metadata   interface{}             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}             `yaml:"DependsOn,omitempty"`
}

type EC2RouteTableProperties struct {
	VpcId interface{} `yaml:"VpcId"`
	Tags  interface{} `yaml:"Tags,omitempty"`
}

func NewEC2RouteTable(properties EC2RouteTableProperties, deps ...interface{}) EC2RouteTable {
	return EC2RouteTable{
		Type:       "AWS::EC2::RouteTable",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2RouteTable(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource EC2RouteTable
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2RouteTable - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource EC2RouteTable) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2RouteTableProperties) Validate() []error {
	errs := []error{}
	if resource.VpcId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpcId'"))
	}
	return errs
}