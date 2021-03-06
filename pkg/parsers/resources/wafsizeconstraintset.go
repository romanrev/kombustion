package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// WAFSizeConstraintSet Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sizeconstraintset.html
type WAFSizeConstraintSet struct {
	Type       string                         `yaml:"Type"`
	Properties WAFSizeConstraintSetProperties `yaml:"Properties"`
	Condition  interface{}                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                    `yaml:"DependsOn,omitempty"`
}

// WAFSizeConstraintSet Properties
type WAFSizeConstraintSetProperties struct {
	Name            interface{} `yaml:"Name"`
	SizeConstraints interface{} `yaml:"SizeConstraints"`
}

// NewWAFSizeConstraintSet constructor creates a new WAFSizeConstraintSet
func NewWAFSizeConstraintSet(properties WAFSizeConstraintSetProperties, deps ...interface{}) WAFSizeConstraintSet {
	return WAFSizeConstraintSet{
		Type:       "AWS::WAF::SizeConstraintSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseWAFSizeConstraintSet parses WAFSizeConstraintSet
func ParseWAFSizeConstraintSet(name string, data string) (cf types.TemplateObject, err error) {
	var resource WAFSizeConstraintSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFSizeConstraintSet - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseWAFSizeConstraintSet validator
func (resource WAFSizeConstraintSet) Validate() []error {
	return resource.Properties.Validate()
}

// ParseWAFSizeConstraintSetProperties validator
func (resource WAFSizeConstraintSetProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.SizeConstraints == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SizeConstraints'"))
	}
	return errs
}
