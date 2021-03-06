package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ServiceCatalogAcceptedPortfolioShare Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-acceptedportfolioshare.html
type ServiceCatalogAcceptedPortfolioShare struct {
	Type       string                                         `yaml:"Type"`
	Properties ServiceCatalogAcceptedPortfolioShareProperties `yaml:"Properties"`
	Condition  interface{}                                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                    `yaml:"DependsOn,omitempty"`
}

// ServiceCatalogAcceptedPortfolioShare Properties
type ServiceCatalogAcceptedPortfolioShareProperties struct {
	AcceptLanguage interface{} `yaml:"AcceptLanguage,omitempty"`
	PortfolioId    interface{} `yaml:"PortfolioId"`
}

// NewServiceCatalogAcceptedPortfolioShare constructor creates a new ServiceCatalogAcceptedPortfolioShare
func NewServiceCatalogAcceptedPortfolioShare(properties ServiceCatalogAcceptedPortfolioShareProperties, deps ...interface{}) ServiceCatalogAcceptedPortfolioShare {
	return ServiceCatalogAcceptedPortfolioShare{
		Type:       "AWS::ServiceCatalog::AcceptedPortfolioShare",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseServiceCatalogAcceptedPortfolioShare parses ServiceCatalogAcceptedPortfolioShare
func ParseServiceCatalogAcceptedPortfolioShare(name string, data string) (cf types.TemplateObject, err error) {
	var resource ServiceCatalogAcceptedPortfolioShare
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ServiceCatalogAcceptedPortfolioShare - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseServiceCatalogAcceptedPortfolioShare validator
func (resource ServiceCatalogAcceptedPortfolioShare) Validate() []error {
	return resource.Properties.Validate()
}

// ParseServiceCatalogAcceptedPortfolioShareProperties validator
func (resource ServiceCatalogAcceptedPortfolioShareProperties) Validate() []error {
	errs := []error{}
	if resource.PortfolioId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PortfolioId'"))
	}
	return errs
}
