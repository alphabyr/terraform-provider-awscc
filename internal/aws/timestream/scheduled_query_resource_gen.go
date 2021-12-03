// Code generated by generators/resource/main.go; DO NOT EDIT.

package timestream

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_timestream_scheduled_query", scheduledQueryResourceType)
}

// scheduledQueryResourceType returns the Terraform awscc_timestream_scheduled_query resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Timestream::ScheduledQuery resource type.
func scheduledQueryResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name of the scheduled query that is generated upon creation.",
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Amazon Resource Name of the scheduled query that is generated upon creation.",
			Type:        types.StringType,
			Computed:    true,
		},
		"client_token": {
			// Property: ClientToken
			// CloudFormation resource type schema:
			// {
			//   "description": "Token provided to ensure idempotency when creating scheduled queries.",
			//   "maxLength": 128,
			//   "minLength": 32,
			//   "type": "string"
			// }
			Description: "Token provided to ensure idempotency when creating scheduled queries.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(32, 128),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"error_report_configuration": {
			// Property: ErrorReportConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Configuration for where error reports will be placed, how they will be named, and how they will be encrypted.",
			//   "properties": {
			//     "S3Configuration": {
			//       "additionalProperties": false,
			//       "description": "S3 configuration for where error reports will be placed, how they will be named, and how they will be encrypted.",
			//       "properties": {
			//         "BucketName": {
			//           "description": "S3 bucket where error reports will be placed.",
			//           "maxLength": 63,
			//           "minLength": 3,
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "EncryptionOption": {
			//           "description": "How error reports will be encrypted.",
			//           "enum": [
			//             "SSE_S3",
			//             "SSE_KMS"
			//           ],
			//           "type": "string"
			//         },
			//         "ObjectKeyPrefix": {
			//           "description": "Prefix for error report names.",
			//           "maxLength": 896,
			//           "minLength": 1,
			//           "pattern": "",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "BucketName"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "S3Configuration"
			//   ],
			//   "type": "object"
			// }
			Description: "Configuration for where error reports will be placed, how they will be named, and how they will be encrypted.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"s3_configuration": {
						// Property: S3Configuration
						Description: "S3 configuration for where error reports will be placed, how they will be named, and how they will be encrypted.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"bucket_name": {
									// Property: BucketName
									Description: "S3 bucket where error reports will be placed.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(3, 63),
									},
								},
								"encryption_option": {
									// Property: EncryptionOption
									Description: "How error reports will be encrypted.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"SSE_S3",
											"SSE_KMS",
										}),
									},
								},
								"object_key_prefix": {
									// Property: ObjectKeyPrefix
									Description: "Prefix for error report names.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 896),
									},
								},
							},
						),
						Required: true,
					},
				},
			),
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "The KMS key for the scheduled query. If the KMS key is not specified, the database will be encrypted with a Timestream managed KMS key located in your account.",
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The KMS key for the scheduled query. If the KMS key is not specified, the database will be encrypted with a Timestream managed KMS key located in your account.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 2048),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"notification_configuration": {
			// Property: NotificationConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Configuration for notification upon scheduled query execution.",
			//   "properties": {
			//     "SnsConfiguration": {
			//       "additionalProperties": false,
			//       "description": "SNS configuration for notification upon scheduled query execution.",
			//       "properties": {
			//         "TopicArn": {
			//           "description": "SNS Topic to be notified upon scheduled query execution.",
			//           "maxLength": 2048,
			//           "minLength": 1,
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "TopicArn"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "SnsConfiguration"
			//   ],
			//   "type": "object"
			// }
			Description: "Configuration for notification upon scheduled query execution.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"sns_configuration": {
						// Property: SnsConfiguration
						Description: "SNS configuration for notification upon scheduled query execution.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"topic_arn": {
									// Property: TopicArn
									Description: "SNS Topic to be notified upon scheduled query execution.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 2048),
									},
								},
							},
						),
						Required: true,
					},
				},
			),
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"query_string": {
			// Property: QueryString
			// CloudFormation resource type schema:
			// {
			//   "description": "The query scheduled to be executed.",
			//   "maxLength": 262144,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The query scheduled to be executed.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 262144),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"sq_error_report_configuration": {
			// Property: SQErrorReportConfiguration
			// CloudFormation resource type schema:
			// {
			//   "description": "Configuration for where error reports will be placed, how they will be named, and how they will be encrypted.",
			//   "type": "string"
			// }
			Description: "Configuration for where error reports will be placed, how they will be named, and how they will be encrypted.",
			Type:        types.StringType,
			Computed:    true,
		},
		"sq_kms_key_id": {
			// Property: SQKmsKeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "The KMS key for the scheduled query. If the KMS key is not specified, the database will be encrypted with a Timestream owned KMS key located in the Timestream account.",
			//   "type": "string"
			// }
			Description: "The KMS key for the scheduled query. If the KMS key is not specified, the database will be encrypted with a Timestream owned KMS key located in the Timestream account.",
			Type:        types.StringType,
			Computed:    true,
		},
		"sq_name": {
			// Property: SQName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name for the scheduled query.",
			//   "type": "string"
			// }
			Description: "The name for the scheduled query.",
			Type:        types.StringType,
			Computed:    true,
		},
		"sq_notification_configuration": {
			// Property: SQNotificationConfiguration
			// CloudFormation resource type schema:
			// {
			//   "description": "Configuration for notification upon scheduled query execution.",
			//   "type": "string"
			// }
			Description: "Configuration for notification upon scheduled query execution.",
			Type:        types.StringType,
			Computed:    true,
		},
		"sq_query_string": {
			// Property: SQQueryString
			// CloudFormation resource type schema:
			// {
			//   "description": "The query scheduled to be executed.",
			//   "type": "string"
			// }
			Description: "The query scheduled to be executed.",
			Type:        types.StringType,
			Computed:    true,
		},
		"sq_schedule_configuration": {
			// Property: SQScheduleConfiguration
			// CloudFormation resource type schema:
			// {
			//   "description": "Configuration that indicates when the scheduled query is executed.",
			//   "type": "string"
			// }
			Description: "Configuration that indicates when the scheduled query is executed.",
			Type:        types.StringType,
			Computed:    true,
		},
		"sq_scheduled_query_execution_role_arn": {
			// Property: SQScheduledQueryExecutionRoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Role that AWS Timestream will assume to execute scheduled query.",
			//   "type": "string"
			// }
			Description: "Role that AWS Timestream will assume to execute scheduled query.",
			Type:        types.StringType,
			Computed:    true,
		},
		"sq_target_configuration": {
			// Property: SQTargetConfiguration
			// CloudFormation resource type schema:
			// {
			//   "description": "Target data source to export query results from.",
			//   "type": "string"
			// }
			Description: "Target data source to export query results from.",
			Type:        types.StringType,
			Computed:    true,
		},
		"schedule_configuration": {
			// Property: ScheduleConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Configuration that indicates when the scheduled query is executed.",
			//   "properties": {
			//     "ScheduleExpression": {
			//       "description": "The cron expression that indicates when the scheduled query is executed.",
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "ScheduleExpression"
			//   ],
			//   "type": "object"
			// }
			Description: "Configuration that indicates when the scheduled query is executed.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"schedule_expression": {
						// Property: ScheduleExpression
						Description: "The cron expression that indicates when the scheduled query is executed.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
						},
					},
				},
			),
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"scheduled_query_execution_role_arn": {
			// Property: ScheduledQueryExecutionRoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Role that AWS Timestream will assume to execute scheduled query.",
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Role that AWS Timestream will assume to execute scheduled query.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 2048),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"scheduled_query_name": {
			// Property: ScheduledQueryName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name for the scheduled query.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name for the scheduled query.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "type": "array"
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(200),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
		"target_configuration": {
			// Property: TargetConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Configuration of target destination table to query.",
			//   "properties": {
			//     "TimestreamConfiguration": {
			//       "additionalProperties": false,
			//       "description": "Timestream configuration of destination table to query.",
			//       "properties": {
			//         "DatabaseName": {
			//           "description": "The source database to query.",
			//           "type": "string"
			//         },
			//         "DimensionMappings": {
			//           "description": "Mappings of dimension names to dimension value types.",
			//           "insertionOrder": false,
			//           "items": {
			//             "additionalProperties": false,
			//             "description": "Mapping of dimension column name to dimension column value type.",
			//             "properties": {
			//               "DimensionValueType": {
			//                 "description": "Value type of the dimension column.",
			//                 "enum": [
			//                   "VARCHAR"
			//                 ],
			//                 "type": "string"
			//               },
			//               "Name": {
			//                 "description": "Name of the dimension column.",
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "Name",
			//               "DimensionValueType"
			//             ],
			//             "type": "object"
			//           },
			//           "type": "array"
			//         },
			//         "MeasureNameColumn": {
			//           "description": "Name of the source measure names column.",
			//           "type": "string"
			//         },
			//         "MixedMeasureMappings": {
			//           "description": "Mapping of measure names and measure value columns from the source table to the destination table.",
			//           "insertionOrder": false,
			//           "items": {
			//             "additionalProperties": false,
			//             "description": "Mapping of measure names and measure value columns from the source table to the destination table.",
			//             "properties": {
			//               "MeasureName": {
			//                 "description": "Name of the measure in source table.",
			//                 "type": "string"
			//               },
			//               "MeasureValueType": {
			//                 "description": "Value type of the measure value column in the destination table.",
			//                 "enum": [
			//                   "BIGINT",
			//                   "BOOLEAN",
			//                   "DOUBLE",
			//                   "VARCHAR",
			//                   "MULTI"
			//                 ],
			//                 "type": "string"
			//               },
			//               "MultiMeasureAttributeMappings": {
			//                 "description": "List of multi-measure value column mappings.",
			//                 "insertionOrder": false,
			//                 "items": {
			//                   "additionalProperties": false,
			//                   "description": "Mapping of measure value columns from the source table to the destination table.",
			//                   "properties": {
			//                     "MeasureValueType": {
			//                       "description": "Value type of the measure value column in the destination table.",
			//                       "enum": [
			//                         "BIGINT",
			//                         "BOOLEAN",
			//                         "DOUBLE",
			//                         "VARCHAR"
			//                       ],
			//                       "type": "string"
			//                     },
			//                     "SourceColumn": {
			//                       "description": "Name of the measure value column in the source table.",
			//                       "type": "string"
			//                     },
			//                     "TargetMultiMeasureAttributeName": {
			//                       "description": "Name of the measure value column in the destination table.",
			//                       "type": "string"
			//                     }
			//                   },
			//                   "required": [
			//                     "SourceColumn",
			//                     "MeasureValueType"
			//                   ],
			//                   "type": "object"
			//                 },
			//                 "minItems": 1,
			//                 "type": "array"
			//               },
			//               "SourceColumn": {
			//                 "description": "Name of the measure value column in the source table.",
			//                 "type": "string"
			//               },
			//               "TargetMeasureName": {
			//                 "description": "Name of the measure in the destination table.",
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "MeasureValueType"
			//             ],
			//             "type": "object"
			//           },
			//           "minItems": 1,
			//           "type": "array"
			//         },
			//         "MultiMeasureMappings": {
			//           "additionalProperties": false,
			//           "description": "Mapping of measure value columns from the source table to the destination table.",
			//           "properties": {
			//             "MultiMeasureAttributeMappings": {
			//               "description": "List of multi-measure value column mappings.",
			//               "insertionOrder": false,
			//               "items": {
			//                 "additionalProperties": false,
			//                 "description": "Mapping of measure value columns from the source table to the destination table.",
			//                 "properties": {
			//                   "MeasureValueType": {
			//                     "description": "Value type of the measure value column in the destination table.",
			//                     "enum": [
			//                       "BIGINT",
			//                       "BOOLEAN",
			//                       "DOUBLE",
			//                       "VARCHAR"
			//                     ],
			//                     "type": "string"
			//                   },
			//                   "SourceColumn": {
			//                     "description": "Name of the measure value column in the source table.",
			//                     "type": "string"
			//                   },
			//                   "TargetMultiMeasureAttributeName": {
			//                     "description": "Name of the measure value column in the destination table.",
			//                     "type": "string"
			//                   }
			//                 },
			//                 "required": [
			//                   "SourceColumn",
			//                   "MeasureValueType"
			//                 ],
			//                 "type": "object"
			//               },
			//               "minItems": 1,
			//               "type": "array"
			//             },
			//             "TargetMultiMeasureName": {
			//               "description": "Name of the multi-measure in the destination table.",
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "MultiMeasureAttributeMappings"
			//           ],
			//           "type": "object"
			//         },
			//         "TableName": {
			//           "description": "The source table to query.",
			//           "type": "string"
			//         },
			//         "TimeColumn": {
			//           "description": "Name of the \"time\" column.",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "DatabaseName",
			//         "TableName",
			//         "TimeColumn",
			//         "DimensionMappings"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "TimestreamConfiguration"
			//   ],
			//   "type": "object"
			// }
			Description: "Configuration of target destination table to query.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"timestream_configuration": {
						// Property: TimestreamConfiguration
						Description: "Timestream configuration of destination table to query.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"database_name": {
									// Property: DatabaseName
									Description: "The source database to query.",
									Type:        types.StringType,
									Required:    true,
								},
								"dimension_mappings": {
									// Property: DimensionMappings
									Description: "Mappings of dimension names to dimension value types.",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"dimension_value_type": {
												// Property: DimensionValueType
												Description: "Value type of the dimension column.",
												Type:        types.StringType,
												Required:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringInSlice([]string{
														"VARCHAR",
													}),
												},
											},
											"name": {
												// Property: Name
												Description: "Name of the dimension column.",
												Type:        types.StringType,
												Required:    true,
											},
										},
										tfsdk.ListNestedAttributesOptions{},
									),
									Required: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										Multiset(),
									},
								},
								"measure_name_column": {
									// Property: MeasureNameColumn
									Description: "Name of the source measure names column.",
									Type:        types.StringType,
									Optional:    true,
								},
								"mixed_measure_mappings": {
									// Property: MixedMeasureMappings
									Description: "Mapping of measure names and measure value columns from the source table to the destination table.",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"measure_name": {
												// Property: MeasureName
												Description: "Name of the measure in source table.",
												Type:        types.StringType,
												Optional:    true,
											},
											"measure_value_type": {
												// Property: MeasureValueType
												Description: "Value type of the measure value column in the destination table.",
												Type:        types.StringType,
												Required:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringInSlice([]string{
														"BIGINT",
														"BOOLEAN",
														"DOUBLE",
														"VARCHAR",
														"MULTI",
													}),
												},
											},
											"multi_measure_attribute_mappings": {
												// Property: MultiMeasureAttributeMappings
												Description: "List of multi-measure value column mappings.",
												Attributes: tfsdk.ListNestedAttributes(
													map[string]tfsdk.Attribute{
														"measure_value_type": {
															// Property: MeasureValueType
															Description: "Value type of the measure value column in the destination table.",
															Type:        types.StringType,
															Required:    true,
															Validators: []tfsdk.AttributeValidator{
																validate.StringInSlice([]string{
																	"BIGINT",
																	"BOOLEAN",
																	"DOUBLE",
																	"VARCHAR",
																}),
															},
														},
														"source_column": {
															// Property: SourceColumn
															Description: "Name of the measure value column in the source table.",
															Type:        types.StringType,
															Required:    true,
														},
														"target_multi_measure_attribute_name": {
															// Property: TargetMultiMeasureAttributeName
															Description: "Name of the measure value column in the destination table.",
															Type:        types.StringType,
															Optional:    true,
														},
													},
													tfsdk.ListNestedAttributesOptions{},
												),
												Optional: true,
												Validators: []tfsdk.AttributeValidator{
													validate.ArrayLenAtLeast(1),
												},
												PlanModifiers: []tfsdk.AttributePlanModifier{
													Multiset(),
												},
											},
											"source_column": {
												// Property: SourceColumn
												Description: "Name of the measure value column in the source table.",
												Type:        types.StringType,
												Optional:    true,
											},
											"target_measure_name": {
												// Property: TargetMeasureName
												Description: "Name of the measure in the destination table.",
												Type:        types.StringType,
												Optional:    true,
											},
										},
										tfsdk.ListNestedAttributesOptions{},
									),
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLenAtLeast(1),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										Multiset(),
									},
								},
								"multi_measure_mappings": {
									// Property: MultiMeasureMappings
									Description: "Mapping of measure value columns from the source table to the destination table.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"multi_measure_attribute_mappings": {
												// Property: MultiMeasureAttributeMappings
												Description: "List of multi-measure value column mappings.",
												Attributes: tfsdk.ListNestedAttributes(
													map[string]tfsdk.Attribute{
														"measure_value_type": {
															// Property: MeasureValueType
															Description: "Value type of the measure value column in the destination table.",
															Type:        types.StringType,
															Required:    true,
															Validators: []tfsdk.AttributeValidator{
																validate.StringInSlice([]string{
																	"BIGINT",
																	"BOOLEAN",
																	"DOUBLE",
																	"VARCHAR",
																}),
															},
														},
														"source_column": {
															// Property: SourceColumn
															Description: "Name of the measure value column in the source table.",
															Type:        types.StringType,
															Required:    true,
														},
														"target_multi_measure_attribute_name": {
															// Property: TargetMultiMeasureAttributeName
															Description: "Name of the measure value column in the destination table.",
															Type:        types.StringType,
															Optional:    true,
														},
													},
													tfsdk.ListNestedAttributesOptions{},
												),
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.ArrayLenAtLeast(1),
												},
												PlanModifiers: []tfsdk.AttributePlanModifier{
													Multiset(),
												},
											},
											"target_multi_measure_name": {
												// Property: TargetMultiMeasureName
												Description: "Name of the multi-measure in the destination table.",
												Type:        types.StringType,
												Optional:    true,
											},
										},
									),
									Optional: true,
								},
								"table_name": {
									// Property: TableName
									Description: "The source table to query.",
									Type:        types.StringType,
									Required:    true,
								},
								"time_column": {
									// Property: TimeColumn
									Description: "Name of the \"time\" column.",
									Type:        types.StringType,
									Required:    true,
								},
							},
						),
						Required: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "The AWS::Timestream::ScheduledQuery resource creates a Timestream Scheduled Query.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Timestream::ScheduledQuery").WithTerraformTypeName("awscc_timestream_scheduled_query")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                   "Arn",
		"bucket_name":                           "BucketName",
		"client_token":                          "ClientToken",
		"database_name":                         "DatabaseName",
		"dimension_mappings":                    "DimensionMappings",
		"dimension_value_type":                  "DimensionValueType",
		"encryption_option":                     "EncryptionOption",
		"error_report_configuration":            "ErrorReportConfiguration",
		"key":                                   "Key",
		"kms_key_id":                            "KmsKeyId",
		"measure_name":                          "MeasureName",
		"measure_name_column":                   "MeasureNameColumn",
		"measure_value_type":                    "MeasureValueType",
		"mixed_measure_mappings":                "MixedMeasureMappings",
		"multi_measure_attribute_mappings":      "MultiMeasureAttributeMappings",
		"multi_measure_mappings":                "MultiMeasureMappings",
		"name":                                  "Name",
		"notification_configuration":            "NotificationConfiguration",
		"object_key_prefix":                     "ObjectKeyPrefix",
		"query_string":                          "QueryString",
		"s3_configuration":                      "S3Configuration",
		"schedule_configuration":                "ScheduleConfiguration",
		"schedule_expression":                   "ScheduleExpression",
		"scheduled_query_execution_role_arn":    "ScheduledQueryExecutionRoleArn",
		"scheduled_query_name":                  "ScheduledQueryName",
		"sns_configuration":                     "SnsConfiguration",
		"source_column":                         "SourceColumn",
		"sq_error_report_configuration":         "SQErrorReportConfiguration",
		"sq_kms_key_id":                         "SQKmsKeyId",
		"sq_name":                               "SQName",
		"sq_notification_configuration":         "SQNotificationConfiguration",
		"sq_query_string":                       "SQQueryString",
		"sq_schedule_configuration":             "SQScheduleConfiguration",
		"sq_scheduled_query_execution_role_arn": "SQScheduledQueryExecutionRoleArn",
		"sq_target_configuration":               "SQTargetConfiguration",
		"table_name":                            "TableName",
		"tags":                                  "Tags",
		"target_configuration":                  "TargetConfiguration",
		"target_measure_name":                   "TargetMeasureName",
		"target_multi_measure_attribute_name":   "TargetMultiMeasureAttributeName",
		"target_multi_measure_name":             "TargetMultiMeasureName",
		"time_column":                           "TimeColumn",
		"timestream_configuration":              "TimestreamConfiguration",
		"topic_arn":                             "TopicArn",
		"value":                                 "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}