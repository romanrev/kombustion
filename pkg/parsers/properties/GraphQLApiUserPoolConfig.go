package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// GraphQLApiUserPoolConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-userpoolconfig.html
type GraphQLApiUserPoolConfig struct {
	AppIdClientRegex interface{} `yaml:"AppIdClientRegex,omitempty"`
	AwsRegion        interface{} `yaml:"AwsRegion,omitempty"`
	DefaultAction    interface{} `yaml:"DefaultAction,omitempty"`
	UserPoolId       interface{} `yaml:"UserPoolId,omitempty"`
}

// GraphQLApiUserPoolConfig validation
func (resource GraphQLApiUserPoolConfig) Validate() []error {
	errs := []error{}

	return errs
}
