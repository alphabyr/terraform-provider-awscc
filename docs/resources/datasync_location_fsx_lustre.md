---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_datasync_location_fsx_lustre Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::DataSync::LocationFSxLustre.
---

# awscc_datasync_location_fsx_lustre (Resource)

Resource schema for AWS::DataSync::LocationFSxLustre.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `fsx_filesystem_arn` (String) The Amazon Resource Name (ARN) for the FSx for Lustre file system.
- `security_group_arns` (List of String) The ARNs of the security groups that are to use to configure the FSx for Lustre file system.

### Optional

- `subdirectory` (String) A subdirectory in the location's path.
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `location_arn` (String) The Amazon Resource Name (ARN) of the Amazon FSx for Lustre file system location that is created.
- `location_uri` (String) The URL of the FSx for Lustre location that was described.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key for an AWS resource tag.
- `value` (String) The value for an AWS resource tag.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_datasync_location_fsx_lustre.example <resource ID>
```
