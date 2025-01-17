---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_forecast_dataset Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type Definition for AWS::Forecast::Dataset
---

# awscc_forecast_dataset (Resource)

Resource Type Definition for AWS::Forecast::Dataset



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `dataset_name` (String) A name for the dataset
- `dataset_type` (String) The dataset type
- `domain` (String) The domain associated with the dataset
- `schema` (Attributes) (see [below for nested schema](#nestedatt--schema))

### Optional

- `data_frequency` (String) Frequency of data collection. This parameter is required for RELATED_TIME_SERIES
- `encryption_config` (Attributes) (see [below for nested schema](#nestedatt--encryption_config))
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--schema"></a>
### Nested Schema for `schema`

Required:

- `attributes` (Attributes List) (see [below for nested schema](#nestedatt--schema--attributes))

<a id="nestedatt--schema--attributes"></a>
### Nested Schema for `schema.attributes`

Required:

- `attribute_name` (String) Name of the dataset field
- `attribute_type` (String) Data type of the field



<a id="nestedatt--encryption_config"></a>
### Nested Schema for `encryption_config`

Optional:

- `kms_key_arn` (String) KMS key used to encrypt the Dataset data
- `role_arn` (String) The ARN of the IAM role that Amazon Forecast can assume to access the AWS KMS key.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_forecast_dataset.example <resource ID>
```
