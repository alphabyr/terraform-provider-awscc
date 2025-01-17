---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_redshift_scheduled_action Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::Redshift::ScheduledAction resource creates an Amazon Redshift Scheduled Action.
---

# awscc_redshift_scheduled_action (Resource)

The `AWS::Redshift::ScheduledAction` resource creates an Amazon Redshift Scheduled Action.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `scheduled_action_name` (String) The name of the scheduled action. The name must be unique within an account.

### Optional

- `enable` (Boolean) If true, the schedule is enabled. If false, the scheduled action does not trigger.
- `end_time` (String) The end time in UTC of the scheduled action. After this time, the scheduled action does not trigger.
- `iam_role` (String) The IAM role to assume to run the target action.
- `schedule` (String) The schedule in `at( )` or `cron( )` format.
- `scheduled_action_description` (String) The description of the scheduled action.
- `start_time` (String) The start time in UTC of the scheduled action. Before this time, the scheduled action does not trigger.
- `target_action` (Attributes) A JSON format string of the Amazon Redshift API operation with input parameters. (see [below for nested schema](#nestedatt--target_action))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `next_invocations` (List of String) List of times when the scheduled action will run.
- `state` (String) The state of the scheduled action.

<a id="nestedatt--target_action"></a>
### Nested Schema for `target_action`

Optional:

- `pause_cluster` (Attributes) Describes a pause cluster operation. For example, a scheduled action to run the `PauseCluster` API operation. (see [below for nested schema](#nestedatt--target_action--pause_cluster))
- `resize_cluster` (Attributes) Describes a resize cluster operation. For example, a scheduled action to run the `ResizeCluster` API operation. (see [below for nested schema](#nestedatt--target_action--resize_cluster))
- `resume_cluster` (Attributes) Describes a resume cluster operation. For example, a scheduled action to run the `ResumeCluster` API operation. (see [below for nested schema](#nestedatt--target_action--resume_cluster))

<a id="nestedatt--target_action--pause_cluster"></a>
### Nested Schema for `target_action.pause_cluster`

Optional:

- `cluster_identifier` (String)


<a id="nestedatt--target_action--resize_cluster"></a>
### Nested Schema for `target_action.resize_cluster`

Optional:

- `classic` (Boolean)
- `cluster_identifier` (String)
- `cluster_type` (String)
- `node_type` (String)
- `number_of_nodes` (Number)


<a id="nestedatt--target_action--resume_cluster"></a>
### Nested Schema for `target_action.resume_cluster`

Optional:

- `cluster_identifier` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_redshift_scheduled_action.example <resource ID>
```
