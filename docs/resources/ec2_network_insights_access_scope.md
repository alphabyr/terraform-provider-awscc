---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_network_insights_access_scope Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::EC2::NetworkInsightsAccessScope
---

# awscc_ec2_network_insights_access_scope (Resource)

Resource schema for AWS::EC2::NetworkInsightsAccessScope



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `exclude_paths` (Attributes List) (see [below for nested schema](#nestedatt--exclude_paths))
- `match_paths` (Attributes List) (see [below for nested schema](#nestedatt--match_paths))
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `created_date` (String)
- `id` (String) Uniquely identifies the resource.
- `network_insights_access_scope_arn` (String)
- `network_insights_access_scope_id` (String)
- `updated_date` (String)

<a id="nestedatt--exclude_paths"></a>
### Nested Schema for `exclude_paths`

Optional:

- `destination` (Attributes) (see [below for nested schema](#nestedatt--exclude_paths--destination))
- `source` (Attributes) (see [below for nested schema](#nestedatt--exclude_paths--source))
- `through_resources` (Attributes List) (see [below for nested schema](#nestedatt--exclude_paths--through_resources))

<a id="nestedatt--exclude_paths--destination"></a>
### Nested Schema for `exclude_paths.destination`

Optional:

- `packet_header_statement` (Attributes) (see [below for nested schema](#nestedatt--exclude_paths--destination--packet_header_statement))
- `resource_statement` (Attributes) (see [below for nested schema](#nestedatt--exclude_paths--destination--resource_statement))

<a id="nestedatt--exclude_paths--destination--packet_header_statement"></a>
### Nested Schema for `exclude_paths.destination.packet_header_statement`

Optional:

- `destination_addresses` (List of String)
- `destination_ports` (List of String)
- `destination_prefix_lists` (List of String)
- `protocols` (List of String)
- `source_addresses` (List of String)
- `source_ports` (List of String)
- `source_prefix_lists` (List of String)


<a id="nestedatt--exclude_paths--destination--resource_statement"></a>
### Nested Schema for `exclude_paths.destination.resource_statement`

Optional:

- `resource_types` (List of String)
- `resources` (List of String)



<a id="nestedatt--exclude_paths--source"></a>
### Nested Schema for `exclude_paths.source`

Optional:

- `packet_header_statement` (Attributes) (see [below for nested schema](#nestedatt--exclude_paths--source--packet_header_statement))
- `resource_statement` (Attributes) (see [below for nested schema](#nestedatt--exclude_paths--source--resource_statement))

<a id="nestedatt--exclude_paths--source--packet_header_statement"></a>
### Nested Schema for `exclude_paths.source.packet_header_statement`

Optional:

- `destination_addresses` (List of String)
- `destination_ports` (List of String)
- `destination_prefix_lists` (List of String)
- `protocols` (List of String)
- `source_addresses` (List of String)
- `source_ports` (List of String)
- `source_prefix_lists` (List of String)


<a id="nestedatt--exclude_paths--source--resource_statement"></a>
### Nested Schema for `exclude_paths.source.resource_statement`

Optional:

- `resource_types` (List of String)
- `resources` (List of String)



<a id="nestedatt--exclude_paths--through_resources"></a>
### Nested Schema for `exclude_paths.through_resources`

Optional:

- `resource_statement` (Attributes) (see [below for nested schema](#nestedatt--exclude_paths--through_resources--resource_statement))

<a id="nestedatt--exclude_paths--through_resources--resource_statement"></a>
### Nested Schema for `exclude_paths.through_resources.resource_statement`

Optional:

- `resource_types` (List of String)
- `resources` (List of String)




<a id="nestedatt--match_paths"></a>
### Nested Schema for `match_paths`

Optional:

- `destination` (Attributes) (see [below for nested schema](#nestedatt--match_paths--destination))
- `source` (Attributes) (see [below for nested schema](#nestedatt--match_paths--source))
- `through_resources` (Attributes List) (see [below for nested schema](#nestedatt--match_paths--through_resources))

<a id="nestedatt--match_paths--destination"></a>
### Nested Schema for `match_paths.destination`

Optional:

- `packet_header_statement` (Attributes) (see [below for nested schema](#nestedatt--match_paths--destination--packet_header_statement))
- `resource_statement` (Attributes) (see [below for nested schema](#nestedatt--match_paths--destination--resource_statement))

<a id="nestedatt--match_paths--destination--packet_header_statement"></a>
### Nested Schema for `match_paths.destination.packet_header_statement`

Optional:

- `destination_addresses` (List of String)
- `destination_ports` (List of String)
- `destination_prefix_lists` (List of String)
- `protocols` (List of String)
- `source_addresses` (List of String)
- `source_ports` (List of String)
- `source_prefix_lists` (List of String)


<a id="nestedatt--match_paths--destination--resource_statement"></a>
### Nested Schema for `match_paths.destination.resource_statement`

Optional:

- `resource_types` (List of String)
- `resources` (List of String)



<a id="nestedatt--match_paths--source"></a>
### Nested Schema for `match_paths.source`

Optional:

- `packet_header_statement` (Attributes) (see [below for nested schema](#nestedatt--match_paths--source--packet_header_statement))
- `resource_statement` (Attributes) (see [below for nested schema](#nestedatt--match_paths--source--resource_statement))

<a id="nestedatt--match_paths--source--packet_header_statement"></a>
### Nested Schema for `match_paths.source.packet_header_statement`

Optional:

- `destination_addresses` (List of String)
- `destination_ports` (List of String)
- `destination_prefix_lists` (List of String)
- `protocols` (List of String)
- `source_addresses` (List of String)
- `source_ports` (List of String)
- `source_prefix_lists` (List of String)


<a id="nestedatt--match_paths--source--resource_statement"></a>
### Nested Schema for `match_paths.source.resource_statement`

Optional:

- `resource_types` (List of String)
- `resources` (List of String)



<a id="nestedatt--match_paths--through_resources"></a>
### Nested Schema for `match_paths.through_resources`

Optional:

- `resource_statement` (Attributes) (see [below for nested schema](#nestedatt--match_paths--through_resources--resource_statement))

<a id="nestedatt--match_paths--through_resources--resource_statement"></a>
### Nested Schema for `match_paths.through_resources.resource_statement`

Optional:

- `resource_types` (List of String)
- `resources` (List of String)




<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_ec2_network_insights_access_scope.example <resource ID>
```
