---
layout: "tfe"
page_title: "Terraform Enterprise: tfe_ip_ranges"
description: |-
  Get HCP Terraform and Terraform Enterprise's IP ranges of its services
---

# Data Source: tfe_ip_ranges

Use this data source to retrieve a list of HCP Terraform's IP ranges. For more information about these IP ranges, view our [documentation about HCP Terraform IP Ranges](https://developer.hashicorp.com/terraform/cloud-docs/architectural-details/ip-ranges).

## Example Usage

```csharp
using Constructs;
using HashiCorp.Cdktf;
/*Provider bindings are generated by running cdktf get.
See https://cdk.tf/provider-generation for more details.*/
using Gen.Providers.Tfe;
class MyConvertedCode : TerraformStack
{
    public MyConvertedCode(Construct scope, string name) : base(scope, name)
    {
        var dataTfeIpRangesAddresses = new DataTfeIpRanges.DataTfeIpRanges(this, "addresses", new DataTfeIpRangesConfig { });
        new TerraformOutput(this, "notifications_ips", new TerraformOutputConfig {
            Value = dataTfeIpRangesAddresses.Notifications
        });
    }
}
```

## Argument Reference

No arguments are required for this datasource.

## Attributes Reference

The following attributes are exported:

* `Api` - The list of IP ranges in CIDR notation used for connections from user site to HCP Terraform APIs.
* `Notifications` - The list of IP ranges in CIDR notation used for notifications.
* `Sentinel` - The list of IP ranges in CIDR notation used for outbound requests from Sentinel policies.
* `Vcs` - The list of IP ranges in CIDR notation used for connecting to VCS providers.


<!-- cache-key: cdktf-0.17.0-pre.15 input-a34895b98b3769147c0512ab50d1c1788daf43f5ff3206286266831c516fc66b -->