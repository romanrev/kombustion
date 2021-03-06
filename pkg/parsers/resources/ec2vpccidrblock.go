package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// EC2VPCCidrBlock Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html
type EC2VPCCidrBlock struct {
	Type       string                    `yaml:"Type"`
	Properties EC2VPCCidrBlockProperties `yaml:"Properties"`
	Condition  interface{}               `yaml:"Condition,omitempty"`
	Metadata   interface{}               `yaml:"Metadata,omitempty"`
	DependsOn  interface{}               `yaml:"DependsOn,omitempty"`
}

// EC2VPCCidrBlock Properties
type EC2VPCCidrBlockProperties struct {
	AmazonProvidedIpv6CidrBlock interface{} `yaml:"AmazonProvidedIpv6CidrBlock,omitempty"`
	CidrBlock                   interface{} `yaml:"CidrBlock,omitempty"`
	VpcId                       interface{} `yaml:"VpcId"`
}

// NewEC2VPCCidrBlock constructor creates a new EC2VPCCidrBlock
func NewEC2VPCCidrBlock(properties EC2VPCCidrBlockProperties, deps ...interface{}) EC2VPCCidrBlock {
	return EC2VPCCidrBlock{
		Type:       "AWS::EC2::VPCCidrBlock",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2VPCCidrBlock parses EC2VPCCidrBlock
func ParseEC2VPCCidrBlock(name string, data string) (cf types.TemplateObject, err error) {
	var resource EC2VPCCidrBlock
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2VPCCidrBlock - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseEC2VPCCidrBlock validator
func (resource EC2VPCCidrBlock) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2VPCCidrBlockProperties validator
func (resource EC2VPCCidrBlockProperties) Validate() []error {
	errs := []error{}
	if resource.VpcId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpcId'"))
	}
	return errs
}
