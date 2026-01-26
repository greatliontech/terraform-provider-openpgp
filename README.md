# Terraform Provider OpenPGP

> **Note**: This project is archived. It works but is not actively maintained.

A Terraform provider for generating OpenPGP (GPG) key pairs.

## Use Cases

- Generate GPG keys for SOPS encryption in Terraform
- Create signing keys for packages or commits
- Bootstrap encryption keys as part of infrastructure provisioning

## Usage

```hcl
terraform {
  required_providers {
    openpgp = {
      source  = "greatliontech/openpgp"
      version = "~> 0.1"
    }
  }
}

provider "openpgp" {}

resource "openpgp_key" "sops" {
  name    = "SOPS Key"
  email   = "sops@example.com"
  comment = "Key for SOPS encryption"
}

output "public_key" {
  value = openpgp_key.sops.public_key_armor
}

output "private_key" {
  value     = openpgp_key.sops.private_key_armor
  sensitive = true
}
```

## Resource: openpgp_key

### Arguments

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `email` | string | yes | Email address for the key identity |
| `name` | string | no | Name for the key identity |
| `comment` | string | no | Comment for the key identity |

### Attributes

| Name | Description |
|------|-------------|
| `id` | Key fingerprint |
| `public_key_armor` | Public key in ASCII armor format |
| `private_key_armor` | Private key in ASCII armor format (sensitive) |

## Limitations

- Keys are generated with default settings (no passphrase, default algorithm)
- No support for key expiration or subkeys
- Keys exist only in Terraform state (no external storage)

## Building

```bash
go install
```

## License

Mozilla Public License 2.0
