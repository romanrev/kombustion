package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ServiceCatalogPortfolio Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolio.html
type ServiceCatalogPortfolio struct {
	Type       string                            `yaml:"Type"`
	Properties ServiceCatalogPortfolioProperties `yaml:"Properties"`
	Condition  interface{}                       `yaml:"Condition,omitempty"`
	Metadata   interface{}                       `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                       `yaml:"DependsOn,omitempty"`
}

// ServiceCatalogPortfolio Properties
type ServiceCatalogPortfolioProperties struct {
	AcceptLanguage interface{} `yaml:"AcceptLanguage,omitempty"`
	Description    interface{} `yaml:"Description,omitempty"`
	DisplayName    interface{} `yaml:"DisplayName"`
	ProviderName   interface{} `yaml:"ProviderName"`
	Tags           interface{} `yaml:"Tags,omitempty"`
}

// NewServiceCatalogPortfolio constructor creates a new ServiceCatalogPortfolio
func NewServiceCatalogPortfolio(properties ServiceCatalogPortfolioProperties, deps ...interface{}) ServiceCatalogPortfolio {
	return ServiceCatalogPortfolio{
		Type:       "AWS::ServiceCatalog::Portfolio",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseServiceCatalogPortfolio parses ServiceCatalogPortfolio
func ParseServiceCatalogPortfolio(name string, data string) (cf types.TemplateObject, err error) {
	var resource ServiceCatalogPortfolio
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ServiceCatalogPortfolio - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseServiceCatalogPortfolio validator
func (resource ServiceCatalogPortfolio) Validate() []error {
	return resource.Properties.Validate()
}

// ParseServiceCatalogPortfolioProperties validator
func (resource ServiceCatalogPortfolioProperties) Validate() []error {
	errs := []error{}
	if resource.DisplayName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DisplayName'"))
	}
	if resource.ProviderName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ProviderName'"))
	}
	return errs
}
