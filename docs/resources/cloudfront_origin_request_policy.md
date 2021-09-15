---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cloudfront_origin_request_policy Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::CloudFront::OriginRequestPolicy
---

# awscc_cloudfront_origin_request_policy (Resource)

Resource Type definition for AWS::CloudFront::OriginRequestPolicy



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **origin_request_policy_config** (Attributes) (see [below for nested schema](#nestedatt--origin_request_policy_config))

### Read-Only

- **id** (String) The ID of this resource.
- **last_modified_time** (String)

<a id="nestedatt--origin_request_policy_config"></a>
### Nested Schema for `origin_request_policy_config`

Required:

- **comment** (String)
- **cookies_config** (Attributes) (see [below for nested schema](#nestedatt--origin_request_policy_config--cookies_config))
- **headers_config** (Attributes) (see [below for nested schema](#nestedatt--origin_request_policy_config--headers_config))
- **name** (String)
- **query_strings_config** (Attributes) (see [below for nested schema](#nestedatt--origin_request_policy_config--query_strings_config))

<a id="nestedatt--origin_request_policy_config--cookies_config"></a>
### Nested Schema for `origin_request_policy_config.cookies_config`

Required:

- **cookie_behavior** (String)
- **cookies** (List of String)


<a id="nestedatt--origin_request_policy_config--headers_config"></a>
### Nested Schema for `origin_request_policy_config.headers_config`

Required:

- **header_behavior** (String)
- **headers** (List of String)


<a id="nestedatt--origin_request_policy_config--query_strings_config"></a>
### Nested Schema for `origin_request_policy_config.query_strings_config`

Required:

- **query_string_behavior** (String)
- **query_strings** (List of String)

