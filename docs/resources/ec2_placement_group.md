---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_placement_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::EC2::PlacementGroup
---

# awscc_ec2_placement_group (Resource)

Resource Type definition for AWS::EC2::PlacementGroup



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `strategy` (String) The placement strategy.

### Read-Only

- `group_name` (String) The Group Name of Placement Group.
- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_ec2_placement_group.example <resource ID>
```
