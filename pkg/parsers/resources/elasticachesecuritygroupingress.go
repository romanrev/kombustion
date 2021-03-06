package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ElastiCacheSecurityGroupIngress Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html
type ElastiCacheSecurityGroupIngress struct {
	Type       string                                    `yaml:"Type"`
	Properties ElastiCacheSecurityGroupIngressProperties `yaml:"Properties"`
	Condition  interface{}                               `yaml:"Condition,omitempty"`
	Metadata   interface{}                               `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                               `yaml:"DependsOn,omitempty"`
}

// ElastiCacheSecurityGroupIngress Properties
type ElastiCacheSecurityGroupIngressProperties struct {
	CacheSecurityGroupName  interface{} `yaml:"CacheSecurityGroupName"`
	EC2SecurityGroupName    interface{} `yaml:"EC2SecurityGroupName"`
	EC2SecurityGroupOwnerId interface{} `yaml:"EC2SecurityGroupOwnerId,omitempty"`
}

// NewElastiCacheSecurityGroupIngress constructor creates a new ElastiCacheSecurityGroupIngress
func NewElastiCacheSecurityGroupIngress(properties ElastiCacheSecurityGroupIngressProperties, deps ...interface{}) ElastiCacheSecurityGroupIngress {
	return ElastiCacheSecurityGroupIngress{
		Type:       "AWS::ElastiCache::SecurityGroupIngress",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseElastiCacheSecurityGroupIngress parses ElastiCacheSecurityGroupIngress
func ParseElastiCacheSecurityGroupIngress(name string, data string) (cf types.TemplateObject, err error) {
	var resource ElastiCacheSecurityGroupIngress
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElastiCacheSecurityGroupIngress - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseElastiCacheSecurityGroupIngress validator
func (resource ElastiCacheSecurityGroupIngress) Validate() []error {
	return resource.Properties.Validate()
}

// ParseElastiCacheSecurityGroupIngressProperties validator
func (resource ElastiCacheSecurityGroupIngressProperties) Validate() []error {
	errs := []error{}
	if resource.CacheSecurityGroupName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CacheSecurityGroupName'"))
	}
	if resource.EC2SecurityGroupName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'EC2SecurityGroupName'"))
	}
	return errs
}
