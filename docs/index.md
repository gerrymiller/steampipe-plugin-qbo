---
organization: Cloudticity
category: ["software development"]
icon_url: "/images/plugins/turbot/ansible.svg"
brand_color: "#1A1918"
display_name: "Quickbooks Online"
short_name: "qbo"
description: "Steampipe plugin to query accounting information from QuickBooks Online."
og_description: "Query QuickBooks Online with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/ansible-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Ansible + Steampipe

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

[Ansible](https://www.ansible.com) offers open-source automation that is simple, flexible, and powerful.

The QBO plugin makes it simpler to query the configured QuickBooks Online account.

Get information about the Quickbooks Online company:

```sql
select
  id,
  company_name
from
  qbo_company_info;
```

```
+----+--------------+
| id | company_name | 
+----+--------------+
| 1  | Cloudticity  | 
+----+--------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/ansible/tables)**

## Quick start

### Install

Download and install the latest QBO plugin:

```sh
steampipe plugin install qbo
```

### Credentials

No credentials are required.

### Configuration

Installing the latest QBO plugin will create a config file (`~/.steampipe/config/qbo.spc`) with a single connection named `qbo`:

Configure your file paths in `~/.steampipe/config/qbo.spc`:

```hcl
connection "qbo" {
  plugin = "qbo"

}
```