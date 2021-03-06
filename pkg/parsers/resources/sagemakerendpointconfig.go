package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// SageMakerEndpointConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html
type SageMakerEndpointConfig struct {
	Type       string                            `yaml:"Type"`
	Properties SageMakerEndpointConfigProperties `yaml:"Properties"`
	Condition  interface{}                       `yaml:"Condition,omitempty"`
	Metadata   interface{}                       `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                       `yaml:"DependsOn,omitempty"`
}

// SageMakerEndpointConfig Properties
type SageMakerEndpointConfigProperties struct {
	EndpointConfigName interface{} `yaml:"EndpointConfigName,omitempty"`
	KmsKeyId           interface{} `yaml:"KmsKeyId,omitempty"`
	ProductionVariants interface{} `yaml:"ProductionVariants"`
	Tags               interface{} `yaml:"Tags,omitempty"`
}

// NewSageMakerEndpointConfig constructor creates a new SageMakerEndpointConfig
func NewSageMakerEndpointConfig(properties SageMakerEndpointConfigProperties, deps ...interface{}) SageMakerEndpointConfig {
	return SageMakerEndpointConfig{
		Type:       "AWS::SageMaker::EndpointConfig",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSageMakerEndpointConfig parses SageMakerEndpointConfig
func ParseSageMakerEndpointConfig(name string, data string) (cf types.TemplateObject, err error) {
	var resource SageMakerEndpointConfig
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SageMakerEndpointConfig - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseSageMakerEndpointConfig validator
func (resource SageMakerEndpointConfig) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSageMakerEndpointConfigProperties validator
func (resource SageMakerEndpointConfigProperties) Validate() []error {
	errs := []error{}
	if resource.ProductionVariants == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ProductionVariants'"))
	}
	return errs
}
