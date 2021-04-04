---
page_title: "openpgp_key Resource - terraform-provider-openpgp"
subcategory: ""
description: |-
  
---

# Resource `openpgp_key`



## Example Usage

```terraform
resource "openpgp_key" "foo" {
  name  = "John Doe"
  email = "john.doe@example.com"
}
```

## Schema

### Required

- **email** (String) Email for this key

### Optional

- **comment** (String) Comment for this key
- **name** (String) Name for this key

### Read-only

- **id** (String) Fingerprint
- **private_key_armor** (String) Private key in armor format
- **public_key_armor** (String) Public key in armor format


