package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BucketAccessControlTranslation Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-accesscontroltranslation.html
type BucketAccessControlTranslation struct {
	Owner interface{} `yaml:"Owner"`
}

// BucketAccessControlTranslation validation
func (resource BucketAccessControlTranslation) Validate() []error {
	errs := []error{}

	if resource.Owner == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Owner'"))
	}
	return errs
}
