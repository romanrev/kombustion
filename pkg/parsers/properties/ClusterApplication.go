package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ClusterApplication Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-application.html
type ClusterApplication struct {
	Name           interface{} `yaml:"Name,omitempty"`
	Version        interface{} `yaml:"Version,omitempty"`
	AdditionalInfo interface{} `yaml:"AdditionalInfo,omitempty"`
	Args           interface{} `yaml:"Args,omitempty"`
}

// ClusterApplication validation
func (resource ClusterApplication) Validate() []error {
	errs := []error{}

	return errs
}
