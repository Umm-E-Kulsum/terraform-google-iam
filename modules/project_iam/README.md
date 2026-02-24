# Project IAM Submodule

This module updates the IAM policy for a single project with given bindings.
It is an ADC-compliant wrapper for the projects_iam module.

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| bindings | Map of role (key) and list of members (value) to add the IAM policies/bindings | `map(list(string))` | `{}` | no |
| conditional\_bindings | List of maps of role and respective conditions, and the members to add the IAM policies/bindings | <pre>list(object({<br>    role        = string<br>    title       = string<br>    description = string<br>    expression  = string<br>    members     = list(string)<br>  }))</pre> | `[]` | no |
| mode | Mode for adding the IAM policies/bindings, additive and authoritative | `string` | `"additive"` | no |
| project\_id | Project to add the IAM policies/bindings | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| members | The list of members that were bound to the project. |
| project\_id | The ID of the project where IAM bindings were applied. |
| roles | The list of roles that were assigned to the members. |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
