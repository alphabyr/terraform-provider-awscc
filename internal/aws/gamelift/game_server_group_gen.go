// Code generated by generators/resource/main.go; DO NOT EDIT.

package gamelift

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_gamelift_game_server_group", gameServerGroupResourceType)
}

// gameServerGroupResourceType returns the Terraform awscc_gamelift_game_server_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::GameLift::GameServerGroup resource type.
func gameServerGroupResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"auto_scaling_group_arn": {
			// Property: AutoScalingGroupArn
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A generated unique ID for the EC2 Auto Scaling group that is associated with this game server group.",
			//   "maxLength": 256,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A generated unique ID for the EC2 Auto Scaling group that is associated with this game server group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"auto_scaling_policy": {
			// Property: AutoScalingPolicy
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Configuration settings to define a scaling policy for the Auto Scaling group that is optimized for game hosting",
			//   "properties": {
			//     "EstimatedInstanceWarmup": {
			//       "additionalProperties": false,
			//       "description": "Length of time, in seconds, it takes for a new instance to start new game server processes and register with GameLift FleetIQ.",
			//       "type": "number"
			//     },
			//     "TargetTrackingConfiguration": {
			//       "additionalProperties": false,
			//       "description": "Settings for a target-based scaling policy applied to Auto Scaling group.",
			//       "properties": {
			//         "TargetValue": {
			//           "additionalProperties": false,
			//           "description": "Desired value to use with a game server group target-based scaling policy.",
			//           "type": "number"
			//         }
			//       },
			//       "required": [
			//         "TargetValue"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "TargetTrackingConfiguration"
			//   ],
			//   "type": "object"
			// }
			Description: "Configuration settings to define a scaling policy for the Auto Scaling group that is optimized for game hosting",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"estimated_instance_warmup": {
						// Property: EstimatedInstanceWarmup
						Description: "Length of time, in seconds, it takes for a new instance to start new game server processes and register with GameLift FleetIQ.",
						Type:        types.NumberType,
						Optional:    true,
					},
					"target_tracking_configuration": {
						// Property: TargetTrackingConfiguration
						Description: "Settings for a target-based scaling policy applied to Auto Scaling group.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"target_value": {
									// Property: TargetValue
									Description: "Desired value to use with a game server group target-based scaling policy.",
									Type:        types.NumberType,
									Required:    true,
								},
							},
						),
						Required: true,
					},
				},
			),
			Optional: true,
		},
		"balancing_strategy": {
			// Property: BalancingStrategy
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The fallback balancing method to use for the game server group when Spot Instances in a Region become unavailable or are not viable for game hosting.",
			//   "enum": [
			//     "SPOT_ONLY",
			//     "SPOT_PREFERRED",
			//     "ON_DEMAND_ONLY"
			//   ],
			//   "type": "string"
			// }
			Description: "The fallback balancing method to use for the game server group when Spot Instances in a Region become unavailable or are not viable for game hosting.",
			Type:        types.StringType,
			Optional:    true,
		},
		"delete_option": {
			// Property: DeleteOption
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The type of delete to perform.",
			//   "enum": [
			//     "SAFE_DELETE",
			//     "FORCE_DELETE",
			//     "RETAIN"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of delete to perform.",
			Type:        types.StringType,
			Optional:    true,
			// DeleteOption is a write-only attribute.
		},
		"game_server_group_arn": {
			// Property: GameServerGroupArn
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A generated unique ID for the game server group.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A generated unique ID for the game server group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"game_server_group_name": {
			// Property: GameServerGroupName
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "An identifier for the new game server group.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "An identifier for the new game server group.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
			},
		},
		"game_server_protection_policy": {
			// Property: GameServerProtectionPolicy
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A flag that indicates whether instances in the game server group are protected from early termination.",
			//   "enum": [
			//     "NO_PROTECTION",
			//     "FULL_PROTECTION"
			//   ],
			//   "type": "string"
			// }
			Description: "A flag that indicates whether instances in the game server group are protected from early termination.",
			Type:        types.StringType,
			Optional:    true,
		},
		"instance_definitions": {
			// Property: InstanceDefinitions
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A set of EC2 instance types to use when creating instances in the group.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "An allowed instance type for your game server group.",
			//     "properties": {
			//       "InstanceType": {
			//         "additionalProperties": false,
			//         "description": "An EC2 instance type designation.",
			//         "type": "string"
			//       },
			//       "WeightedCapacity": {
			//         "additionalProperties": false,
			//         "description": "Instance weighting that indicates how much this instance type contributes to the total capacity of a game server group.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "InstanceType"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 20,
			//   "minItems": 2,
			//   "type": "array"
			// }
			Description: "A set of EC2 instance types to use when creating instances in the group.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"instance_type": {
						// Property: InstanceType
						Description: "An EC2 instance type designation.",
						Type:        types.StringType,
						Required:    true,
					},
					"weighted_capacity": {
						// Property: WeightedCapacity
						Description: "Instance weighting that indicates how much this instance type contributes to the total capacity of a game server group.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 2,
					MaxItems: 20,
				},
			),
			Required: true,
		},
		"launch_template": {
			// Property: LaunchTemplate
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The EC2 launch template that contains configuration settings and game server code to be deployed to all instances in the game server group.",
			//   "properties": {
			//     "LaunchTemplateId": {
			//       "additionalProperties": false,
			//       "description": "A unique identifier for an existing EC2 launch template.",
			//       "type": "string"
			//     },
			//     "LaunchTemplateName": {
			//       "additionalProperties": false,
			//       "description": "A readable identifier for an existing EC2 launch template.",
			//       "type": "string"
			//     },
			//     "Version": {
			//       "additionalProperties": false,
			//       "description": "The version of the EC2 launch template to use.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The EC2 launch template that contains configuration settings and game server code to be deployed to all instances in the game server group.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"launch_template_id": {
						// Property: LaunchTemplateId
						Description: "A unique identifier for an existing EC2 launch template.",
						Type:        types.StringType,
						Optional:    true,
					},
					"launch_template_name": {
						// Property: LaunchTemplateName
						Description: "A readable identifier for an existing EC2 launch template.",
						Type:        types.StringType,
						Optional:    true,
					},
					"version": {
						// Property: Version
						Description: "The version of the EC2 launch template to use.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Required: true,
		},
		"max_size": {
			// Property: MaxSize
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The maximum number of instances allowed in the EC2 Auto Scaling group.",
			//   "minimum": 1,
			//   "type": "number"
			// }
			Description: "The maximum number of instances allowed in the EC2 Auto Scaling group.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"min_size": {
			// Property: MinSize
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The minimum number of instances allowed in the EC2 Auto Scaling group.",
			//   "minimum": 0,
			//   "type": "number"
			// }
			Description: "The minimum number of instances allowed in the EC2 Auto Scaling group.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The Amazon Resource Name (ARN) for an IAM role that allows Amazon GameLift to access your EC2 Auto Scaling groups.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) for an IAM role that allows Amazon GameLift to access your EC2 Auto Scaling groups.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A list of labels to assign to the new game server group resource.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "The key for a developer-defined key:value pair for tagging an AWS resource.",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for a developer-defined key:value pair for tagging an AWS resource.",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "A list of labels to assign to the new game server group resource.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key for a developer-defined key:value pair for tagging an AWS resource.",
						Type:        types.StringType,
						Optional:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for a developer-defined key:value pair for tagging an AWS resource.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 0,
					MaxItems: 200,
				},
			),
			Optional: true,
		},
		"vpc_subnets": {
			// Property: VpcSubnets
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A list of virtual private cloud (VPC) subnets to use with instances in the game server group.",
			//   "items": {
			//     "maxLength": 24,
			//     "minLength": 15,
			//     "pattern": "",
			//     "type": "string"
			//   },
			//   "maxItems": 20,
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "A list of virtual private cloud (VPC) subnets to use with instances in the game server group.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(1, 20),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "The AWS::GameLift::GameServerGroup resource creates an Amazon GameLift (GameLift) GameServerGroup.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::GameLift::GameServerGroup").WithTerraformTypeName("awscc_gamelift_game_server_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auto_scaling_group_arn":        "AutoScalingGroupArn",
		"auto_scaling_policy":           "AutoScalingPolicy",
		"balancing_strategy":            "BalancingStrategy",
		"delete_option":                 "DeleteOption",
		"estimated_instance_warmup":     "EstimatedInstanceWarmup",
		"game_server_group_arn":         "GameServerGroupArn",
		"game_server_group_name":        "GameServerGroupName",
		"game_server_protection_policy": "GameServerProtectionPolicy",
		"instance_definitions":          "InstanceDefinitions",
		"instance_type":                 "InstanceType",
		"key":                           "Key",
		"launch_template":               "LaunchTemplate",
		"launch_template_id":            "LaunchTemplateId",
		"launch_template_name":          "LaunchTemplateName",
		"max_size":                      "MaxSize",
		"min_size":                      "MinSize",
		"role_arn":                      "RoleArn",
		"tags":                          "Tags",
		"target_tracking_configuration": "TargetTrackingConfiguration",
		"target_value":                  "TargetValue",
		"value":                         "Value",
		"version":                       "Version",
		"vpc_subnets":                   "VpcSubnets",
		"weighted_capacity":             "WeightedCapacity",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/DeleteOption",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_gamelift_game_server_group", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
