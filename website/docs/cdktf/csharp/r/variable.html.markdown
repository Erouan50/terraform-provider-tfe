---
layout: "tfe"
page_title: "Terraform Enterprise: tfe_variable"
description: |-
  Manages variables.
---

# tfe_variable

Creates, updates and destroys variables.

## Example Usage

Basic usage for workspaces:

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
        var tfeOrganizationTest = new Organization.Organization(this, "test", new OrganizationConfig {
            Email = "admin@company.com",
            Name = "my-org-name"
        });
        var tfeWorkspaceTest = new Workspace.Workspace(this, "test_1", new WorkspaceConfig {
            Name = "my-workspace-name",
            Organization = Token.AsString(tfeOrganizationTest.Name)
        });
        /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
        tfeWorkspaceTest.OverrideLogicalId("test");
        var tfeVariableTest = new Variable.Variable(this, "test_2", new VariableConfig {
            Category = "terraform",
            Description = "a useful description",
            Key = "my_key_name",
            Value = "my_value_name",
            WorkspaceId = Token.AsString(tfeWorkspaceTest.Id)
        });
        /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
        tfeVariableTest.OverrideLogicalId("test");
    }
}
```

Basic usage for variable sets:

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
        var tfeOrganizationTest = new Organization.Organization(this, "test", new OrganizationConfig {
            Email = "admin@company.com",
            Name = "my-org-name"
        });
        var tfeVariableSetTest = new VariableSet.VariableSet(this, "test_1", new VariableSetConfig {
            Description = "Some description.",
            Global = false,
            Name = "Test Varset",
            Organization = Token.AsString(tfeOrganizationTest.Name)
        });
        /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
        tfeVariableSetTest.OverrideLogicalId("test");
        new Variable.Variable(this, "test-a", new VariableConfig {
            Category = "terraform",
            Description = "a useful description",
            Key = "seperate_variable",
            Value = "my_value_name",
            VariableSetId = Token.AsString(tfeVariableSetTest.Id)
        });
        new Variable.Variable(this, "test-b", new VariableConfig {
            Category = "env",
            Description = "an environment variable",
            Key = "another_variable",
            Value = "my_value_name",
            VariableSetId = Token.AsString(tfeVariableSetTest.Id)
        });
    }
}
```

## Argument Reference

The following arguments are supported:

* `Key` - (Required) Name of the variable.
* `Value` - (Required) Value of the variable.
* `Category` - (Required) Whether this is a Terraform or environment variable.
  Valid values are `Terraform` or `Env`.
* `Description` - (Optional) Description of the variable.
* `Hcl` - (Optional) Whether to evaluate the value of the variable as a string
  of HCL code. Has no effect for environment variables. Defaults to `False`.
* `Sensitive` - (Optional) Whether the value is sensitive. If true then the
variable is written once and not visible thereafter. Defaults to `False`.
* One of the following (Required)
    * `WorkspaceId` - ID of the workspace that owns the variable.
    * `VariableSetId` - ID of the variable set that owns the variable.

~> **NOTE:** When `Sensitive` is set to true, Terraform cannot detect and repair
drift if `Value` is later changed out-of-band via the HCP Terraform UI.
Terraform will only change the value for a sensitive variable if you change
`Value` in the configuration, so that it no longer matches the last known value
in the state.

## Attributes Reference

* `Id` - The ID of the variable.
* `ReadableValue` - Only present if the variable is non-sensitive. A copy of the value which will not be marked as sensitive in plan outputs.

### Using readable_value

While the `Value` field may be referenced in other resources, for safety it is always treated as sensitive. This means that it will always be redacted from plan outputs, and any other resource attributes which depend on it will also be redacted.

The `ReadableValue` attribute is not sensitive, and will not be redacted; instead, it will be null if the variable is sensitive. This allows other resources to reference it, while keeping their plan outputs readable.

For example:
```
resource "tfe_variable" "sensitive_var" {
  key          = "sensitive_key"
  value        = "sensitive_value" // this will be redacted from plan outputs
  category     = "terraform"
  workspace_id = tfe_workspace.workspace.id
  sensitive    = true
}

resource "tfe_variable" "visible_var" {
  key          = "visible_key"
  value        = "visible_value" // this will be redacted from plan outputs
  category     = "terraform"
  workspace_id = tfe_workspace.workspace.id
  sensitive    = false
}

resource "tfe_workspace" "sensitive_workspace" {
  name = "workspace-${tfe_variable.sensitive_var.value}" // this will be redacted from plan outputs
  organization = "organization name"
}

resource "tfe_workspace" "visible_workspace" {
  name = "workspace-${tfe_variable.visible_var.readable_value}" // this will not be redacted from plan outputs
  organization = "organization name"
}

```

`ReadableValue` will be null if the variable is sensitive. `ReadableValue` may not be set explicitly in the resource configuration.


## Import

Variables can be imported.

To import a variable that's part of a workspace, use
`<ORGANIZATION NAME>/<WORKSPACE NAME>/<VARIABLE ID>` as the import ID. For
example:

```shell
terraform import tfe_variable.test my-org-name/my-workspace-name/var-5rTwnSaRPogw6apb
```

To import a variable that's part of a variable set, use
`<ORGANIZATION NAME>/<VARIABLE SET ID>/<VARIABLE ID>` as the import ID. For
example:

```shell
terraform import tfe_variable.test my-org-name/varset-47qC3LmA47piVan7/var-5rTwnSaRPogw6apb
```

<!-- cache-key: cdktf-0.17.0-pre.15 input-c3bad9baf89e2964d9e69299417eec51d0597f0fb81bde0ba7db8d5ce2d6617e -->