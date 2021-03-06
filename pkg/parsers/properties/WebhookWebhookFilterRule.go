package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// WebhookWebhookFilterRule Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-webhook-webhookfilterrule.html
type WebhookWebhookFilterRule struct {
	JsonPath    interface{} `yaml:"JsonPath"`
	MatchEquals interface{} `yaml:"MatchEquals,omitempty"`
}

// WebhookWebhookFilterRule validation
func (resource WebhookWebhookFilterRule) Validate() []error {
	errs := []error{}

	if resource.JsonPath == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'JsonPath'"))
	}
	return errs
}
