package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// DistributionOrigin Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origin.html
type DistributionOrigin struct {
	DomainName          interface{}                     `yaml:"DomainName"`
	Id                  interface{}                     `yaml:"Id"`
	OriginPath          interface{}                     `yaml:"OriginPath,omitempty"`
	S3OriginConfig      *DistributionS3OriginConfig     `yaml:"S3OriginConfig,omitempty"`
	OriginCustomHeaders interface{}                     `yaml:"OriginCustomHeaders,omitempty"`
	CustomOriginConfig  *DistributionCustomOriginConfig `yaml:"CustomOriginConfig,omitempty"`
}

// DistributionOrigin validation
func (resource DistributionOrigin) Validate() []error {
	errs := []error{}

	if resource.DomainName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DomainName'"))
	}
	if resource.Id == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Id'"))
	}
	return errs
}
