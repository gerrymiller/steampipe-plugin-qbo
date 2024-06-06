<!-- markdownlint-disable-next-line MD041 -->
![image](https://www.cloudticity.com/hubfs/Cloudticity_Logo_2020%20(1).png#keepProtocol)

![CodeQL](https://github.com/Cloudticity/steampipe-plugin-qbo/actions/workflows/codeql.yml/badge.svg)

# Quickbooks Online Plugin for Steampipe

Use SQL to query accounting data from QuickBooks Online (QBO).

- **[Get started →](https://cloudticity.com)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/ansible/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/Cloudticity/steampipe-plugin-qbo/issues)

## Quick start

### Install

Download and install the latest QuickBooks Online plugin:

```bash
steampipe plugin install qbo
```

Configure your config file in `~/.steampipe/config/qbo.spc`:

```hcl
connection "qbo" {
  plugin = "qbo"

  # Sandbox configuration
  baseURL = "https://sandbox-quickbooks.api.intuit.com"
  discoveryDocumentURL = "https://developer.api.intuit.com/.well-known/openid_sandbox_configuration"
  clientId = "[Client ID]"
  clientSecret = "[Client Secret]"
  realmId = "[Realm ID]"
  accessToken = "[Access Token]"
  refreshToken = "[Refresh Token]"}
```

Run steampipe:

```shell
steampipe query
```

Get information about the QuickBooks Online company:

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

## Engines

This plugin is available for the following engines:

| Engine        | Description
|---------------|------------------------------------------
| [Steampipe](https://steampipe.io/docs) | The Steampipe CLI exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. Mods extend Steampipe's capabilities with dashboards, reports, and controls built with simple HCL. The Steampipe CLI is a turnkey solution that includes its own Postgres database, plugin management, and mod support.
| [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview) | Steampipe Postgres FDWs are native Postgres Foreign Data Wrappers that translate APIs to foreign tables. Unlike Steampipe CLI, which ships with its own Postgres server instance, the Steampipe Postgres FDWs can be installed in any supported Postgres database version.
| [SQLite Extension](https://steampipe.io/docs/steampipe_sqlite/overview) | Steampipe SQLite Extensions provide SQLite virtual tables that translate your queries into API calls, transparently fetching information from your API or service as you request it.
| [Export](https://steampipe.io/docs/steampipe_export/overview) | Steampipe Plugin Exporters provide a flexible mechanism for exporting information from cloud services and APIs. Each exporter is a stand-alone binary that allows you to extract data using Steampipe plugins without a database.
| [Turbot Pipes](https://turbot.com/pipes/docs) | Turbot Pipes is the only intelligence, automation & security platform built specifically for DevOps. Pipes provide hosted Steampipe database instances, shared dashboards, snapshots, and more.

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/Cloudticity/steampipe-plugin-qbo.git
cd steampipe-plugin-qbo
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/qbo.spc
```

Try it!

```
steampipe query
> .inspect qbo
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Open Source & Contributing

This repository is published under the [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0) (source code) and [CC BY-NC-ND](https://creativecommons.org/licenses/by-nc-nd/2.0/) (docs) licenses. Please see our [code of conduct](https://github.com/turbot/.github/blob/main/CODE_OF_CONDUCT.md). We look forward to collaborating with you!

[Steampipe](https://steampipe.io) is a product produced from this open source software, exclusively by [Turbot HQ, Inc](https://turbot.com). It is distributed under our commercial terms. Others are allowed to make their own distribution of the software, but cannot use any of the Turbot trademarks, cloud services, etc. You can learn more in our [Open Source FAQ](https://turbot.com/open-source).

## Get Involved

**[Join #steampipe on Slack →](https://turbot.com/community/join)**

Want to help but don't know where to start? Pick up one of the `help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Quickbooks Online Plugin](https://github.com/turbot/steampipe-plugin-qbo/labels/help%20wanted)
