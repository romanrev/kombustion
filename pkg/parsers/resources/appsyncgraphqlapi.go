package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// AppSyncGraphQLApi Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html
type AppSyncGraphQLApi struct {
	Type       string                      `yaml:"Type"`
	Properties AppSyncGraphQLApiProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

// AppSyncGraphQLApi Properties
type AppSyncGraphQLApiProperties struct {
	AuthenticationType  interface{}                               `yaml:"AuthenticationType"`
	Name                interface{}                               `yaml:"Name"`
	UserPoolConfig      *properties.GraphQLApiUserPoolConfig      `yaml:"UserPoolConfig,omitempty"`
	OpenIDConnectConfig *properties.GraphQLApiOpenIDConnectConfig `yaml:"OpenIDConnectConfig,omitempty"`
	LogConfig           *properties.GraphQLApiLogConfig           `yaml:"LogConfig,omitempty"`
}

// NewAppSyncGraphQLApi constructor creates a new AppSyncGraphQLApi
func NewAppSyncGraphQLApi(properties AppSyncGraphQLApiProperties, deps ...interface{}) AppSyncGraphQLApi {
	return AppSyncGraphQLApi{
		Type:       "AWS::AppSync::GraphQLApi",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseAppSyncGraphQLApi parses AppSyncGraphQLApi
func ParseAppSyncGraphQLApi(name string, data string) (cf types.TemplateObject, err error) {
	var resource AppSyncGraphQLApi
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: AppSyncGraphQLApi - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseAppSyncGraphQLApi validator
func (resource AppSyncGraphQLApi) Validate() []error {
	return resource.Properties.Validate()
}

// ParseAppSyncGraphQLApiProperties validator
func (resource AppSyncGraphQLApiProperties) Validate() []error {
	errs := []error{}
	if resource.AuthenticationType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AuthenticationType'"))
	}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
