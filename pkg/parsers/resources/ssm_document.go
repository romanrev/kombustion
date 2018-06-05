package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type SSMDocument struct {
	Type       string                `yaml:"Type"`
	Properties SSMDocumentProperties `yaml:"Properties"`
	Condition  interface{}           `yaml:"Condition,omitempty"`
	Metadata   interface{}           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}           `yaml:"DependsOn,omitempty"`
}

type SSMDocumentProperties struct {
	Content      interface{} `yaml:"Content"`
	DocumentType interface{} `yaml:"DocumentType,omitempty"`
	Tags         interface{} `yaml:"Tags,omitempty"`
}

func NewSSMDocument(properties SSMDocumentProperties, deps ...interface{}) SSMDocument {
	return SSMDocument{
		Type:       "AWS::SSM::Document",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseSSMDocument(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource SSMDocument
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SSMDocument - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource SSMDocument) Validate() []error {
	return resource.Properties.Validate()
}

func (resource SSMDocumentProperties) Validate() []error {
	errs := []error{}
	if resource.Content == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Content'"))
	}
	return errs
}