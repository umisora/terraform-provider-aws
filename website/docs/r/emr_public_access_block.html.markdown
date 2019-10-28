---
layout: "aws"
page_title: "AWS: aws_emr_public_access_block"
description: |-
  Provides a resource to manage AWS EMR Block Public Access Configuration
---

# Resource: emr_public_access_block

Provides a resource to manage AWS EMR Block Public Access Configuration

## Example Usage

```hcl
resource "aws_emr_block_public_access_configuration" "example" {
  block_public_security_group_rules = true

  permitted_public_security_group_rule_range {
    min_range = 22
    max_range = 22
  }

  permitted_public_security_group_rule_range {
    min_range = 100
    max_range = 101
  }
}
```

## Argument Reference

The following arguments are supported:

* `block_public_security_group_rules` - (Optional)  xxxx.
* `permitted_public_security_group_rule_range` - (Optional)  xxxx.

## permitted_public_security_group_rule_range

* `min_range` - (Optional) xxxx.
* `max_range` - (Optional) xxxx.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

## Import

EMR Block Public Access Configuration can be imported, e.g.

```
$ terraform import emr_public_access_block
```
