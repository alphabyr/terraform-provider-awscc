// Code generated by generators/resource/main.go; DO NOT EDIT.

package workspaces

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
	registry.AddResourceTypeFactory("awscc_workspaces_connection_alias", connectionAliasResourceType)
}

// connectionAliasResourceType returns the Terraform awscc_workspaces_connection_alias resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::WorkSpaces::ConnectionAlias resource type.
func connectionAliasResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"alias_id": {
			// Property: AliasId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 68,
			//   "minLength": 13,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"associations": {
			// Property: Associations
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "AssociatedAccountId": {
			//         "type": "string"
			//       },
			//       "AssociationStatus": {
			//         "enum": [
			//           "NOT_ASSOCIATED",
			//           "PENDING_ASSOCIATION",
			//           "ASSOCIATED_WITH_OWNER_ACCOUNT",
			//           "ASSOCIATED_WITH_SHARED_ACCOUNT",
			//           "PENDING_DISASSOCIATION"
			//         ],
			//         "type": "string"
			//       },
			//       "ConnectionIdentifier": {
			//         "maxLength": 20,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "ResourceId": {
			//         "maxLength": 1000,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxLength": 25,
			//   "minLength": 1,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"associated_account_id": {
						// Property: AssociatedAccountId
						Type:     types.StringType,
						Optional: true,
					},
					"association_status": {
						// Property: AssociationStatus
						Type:     types.StringType,
						Optional: true,
					},
					"connection_identifier": {
						// Property: ConnectionIdentifier
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 20),
						},
					},
					"resource_id": {
						// Property: ResourceId
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 1000),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"connection_alias_state": {
			// Property: ConnectionAliasState
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "CREATING",
			//     "CREATED",
			//     "DELETING"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"connection_string": {
			// Property: ConnectionString
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
			},
			// ConnectionString is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			Computed: true,
			// Tags is a force-new attribute.
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::WorkSpaces::ConnectionAlias",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::WorkSpaces::ConnectionAlias").WithTerraformTypeName("awscc_workspaces_connection_alias")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alias_id":               "AliasId",
		"associated_account_id":  "AssociatedAccountId",
		"association_status":     "AssociationStatus",
		"associations":           "Associations",
		"connection_alias_state": "ConnectionAliasState",
		"connection_identifier":  "ConnectionIdentifier",
		"connection_string":      "ConnectionString",
		"key":                    "Key",
		"resource_id":            "ResourceId",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_workspaces_connection_alias", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
