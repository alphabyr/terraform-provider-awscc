// Code generated by generators/resource/main.go; DO NOT EDIT.

package config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_config_organization_conformance_pack", organizationConformancePackResourceType)
}

// organizationConformancePackResourceType returns the Terraform awscc_config_organization_conformance_pack resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Config::OrganizationConformancePack resource type.
func organizationConformancePackResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"conformance_pack_input_parameters": {
			// Property: ConformancePackInputParameters
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of ConformancePackInputParameter objects.",
			//   "items": {
			//     "description": "Input parameters in the form of key-value pairs for the conformance pack.",
			//     "properties": {
			//       "ParameterName": {
			//         "maxLength": 255,
			//         "minLength": 0,
			//         "type": "string"
			//       },
			//       "ParameterValue": {
			//         "maxLength": 4096,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "ParameterName",
			//       "ParameterValue"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 60,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "A list of ConformancePackInputParameter objects.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"parameter_name": {
						// Property: ParameterName
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 255),
						},
					},
					"parameter_value": {
						// Property: ParameterValue
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 4096),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 0,
					MaxItems: 60,
				},
			),
			Optional: true,
		},
		"delivery_s3_bucket": {
			// Property: DeliveryS3Bucket
			// CloudFormation resource type schema:
			// {
			//   "description": "AWS Config stores intermediate files while processing conformance pack template.",
			//   "maxLength": 63,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "AWS Config stores intermediate files while processing conformance pack template.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 63),
			},
		},
		"delivery_s3_key_prefix": {
			// Property: DeliveryS3KeyPrefix
			// CloudFormation resource type schema:
			// {
			//   "description": "The prefix for the delivery S3 bucket.",
			//   "maxLength": 1024,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The prefix for the delivery S3 bucket.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 1024),
			},
		},
		"excluded_accounts": {
			// Property: ExcludedAccounts
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "maxItems": 1000,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "A list of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 1000),
			},
		},
		"organization_conformance_pack_name": {
			// Property: OrganizationConformancePackName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the organization conformance pack.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the organization conformance pack.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"template_body": {
			// Property: TemplateBody
			// CloudFormation resource type schema:
			// {
			//   "description": "A string containing full conformance pack template body.",
			//   "maxLength": 51200,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "A string containing full conformance pack template body.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 51200),
			},
			// TemplateBody is a write-only property.
		},
		"template_s3_uri": {
			// Property: TemplateS3Uri
			// CloudFormation resource type schema:
			// {
			//   "description": "Location of file containing the template body.",
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Location of file containing the template body.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 1024),
			},
			// TemplateS3Uri is a write-only property.
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::Config::OrganizationConformancePack.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Config::OrganizationConformancePack").WithTerraformTypeName("awscc_config_organization_conformance_pack")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"conformance_pack_input_parameters":  "ConformancePackInputParameters",
		"delivery_s3_bucket":                 "DeliveryS3Bucket",
		"delivery_s3_key_prefix":             "DeliveryS3KeyPrefix",
		"excluded_accounts":                  "ExcludedAccounts",
		"organization_conformance_pack_name": "OrganizationConformancePackName",
		"parameter_name":                     "ParameterName",
		"parameter_value":                    "ParameterValue",
		"template_body":                      "TemplateBody",
		"template_s3_uri":                    "TemplateS3Uri",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/TemplateBody",
		"/properties/TemplateS3Uri",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
