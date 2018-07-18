package outputs

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ParseAmazonMQConfiguration Documentation http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html
func ParseAmazonMQConfiguration(name string, data string) (cf types.TemplateObject, err error) {

	var resource, output types.TemplateObject
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}

	cf = types.TemplateObject{
		name: types.TemplateObject{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-AmazonMQConfiguration-" + name,
				},
			},
		},
	}

	output = types.TemplateObject{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "Arn"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-AmazonMQConfiguration-" + name + "-Arn",
			},
		},
	}
	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}
	cf[name+"Arn"] = output

	output = types.TemplateObject{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "Id"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-AmazonMQConfiguration-" + name + "-Id",
			},
		},
	}
	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}
	cf[name+"Id"] = output

	output = types.TemplateObject{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "Revision"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-AmazonMQConfiguration-" + name + "-Revision",
			},
		},
	}
	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}
	cf[name+"Revision"] = output

	return
}