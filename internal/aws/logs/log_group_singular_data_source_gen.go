// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package logs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_logs_log_group", logGroupDataSourceType)
}

// logGroupDataSourceType returns the Terraform awscc_logs_log_group data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Logs::LogGroup resource type.
func logGroupDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The CloudWatch log group ARN.",
			//   "type": "string"
			// }
			Description: "The CloudWatch log group ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
			//   "maxLength": 256,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
			Type:        types.StringType,
			Computed:    true,
		},
		"log_group_name": {
			// Property: LogGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.",
			//   "maxLength": 512,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"retention_in_days": {
			// Property: RetentionInDays
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653.",
			//   "enum": [
			//     1,
			//     3,
			//     5,
			//     7,
			//     14,
			//     30,
			//     60,
			//     90,
			//     120,
			//     150,
			//     180,
			//     365,
			//     400,
			//     545,
			//     731,
			//     1827,
			//     3653
			//   ],
			//   "type": "integer"
			// }
			Description: "The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653.",
			Type:        types.NumberType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Logs::LogGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Logs::LogGroup").WithTerraformTypeName("awscc_logs_log_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":               "Arn",
		"kms_key_id":        "KmsKeyId",
		"log_group_name":    "LogGroupName",
		"retention_in_days": "RetentionInDays",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}