package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// EC2Host Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html
type EC2Host struct {
	Type       string            `yaml:"Type"`
	Properties EC2HostProperties `yaml:"Properties"`
	Condition  interface{}       `yaml:"Condition,omitempty"`
	Metadata   interface{}       `yaml:"Metadata,omitempty"`
	DependsOn  interface{}       `yaml:"DependsOn,omitempty"`
}

// EC2Host Properties
type EC2HostProperties struct {
	AutoPlacement    interface{} `yaml:"AutoPlacement,omitempty"`
	AvailabilityZone interface{} `yaml:"AvailabilityZone"`
	InstanceType     interface{} `yaml:"InstanceType"`
}

// NewEC2Host constructor creates a new EC2Host
func NewEC2Host(properties EC2HostProperties, deps ...interface{}) EC2Host {
	return EC2Host{
		Type:       "AWS::EC2::Host",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2Host parses EC2Host
func ParseEC2Host(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource EC2Host
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2Host - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseEC2Host validator
func (resource EC2Host) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2HostProperties validator
func (resource EC2HostProperties) Validate() []error {
	errs := []error{}
	if resource.AvailabilityZone == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AvailabilityZone'"))
	}
	if resource.InstanceType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceType'"))
	}
	return errs
}