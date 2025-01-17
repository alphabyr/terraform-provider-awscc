---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_connect_contact_flow Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Connect::ContactFlow
---

# awscc_connect_contact_flow (Resource)

Resource Type definition for AWS::Connect::ContactFlow



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `content` (String) The content of the contact flow in JSON format.
- `instance_arn` (String) The identifier of the Amazon Connect instance (ARN).
- `name` (String) The name of the contact flow.

### Optional

- `description` (String) The description of the contact flow.
- `state` (String) The state of the contact flow.
- `tags` (Attributes Set) One or more tags. (see [below for nested schema](#nestedatt--tags))
- `type` (String) The type of the contact flow.

### Read-Only

- `contact_flow_arn` (String) The identifier of the contact flow (ARN).
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. . You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_connect_contact_flow.example <resource ID>
```
