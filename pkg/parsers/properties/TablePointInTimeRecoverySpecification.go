package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// TablePointInTimeRecoverySpecification Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-pointintimerecoveryspecification.html
type TablePointInTimeRecoverySpecification struct {
	PointInTimeRecoveryEnabled interface{} `yaml:"PointInTimeRecoveryEnabled,omitempty"`
}

// TablePointInTimeRecoverySpecification validation
func (resource TablePointInTimeRecoverySpecification) Validate() []error {
	errs := []error{}

	return errs
}
