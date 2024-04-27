---
title: "Steampipe Table: qbo_company_info - Query QuickBooks Online Company Info using SQL"
description: "Allows users to query QuickBooks Online Company Info."
---

# Table: qbo_company_info - Query QuickBooks Online Company Info using SQL

QuickBooks online is an online accounting package used by millions of businesses worldwide. QBO Company Info is a collection of information .

## Table Usage Guide

The `qbo_company_info` table provides insights into information about the Company that owns the QuickBooks Online account.

## Examples

### Execute a simple query
Discover information about the Company that owns the QBO account.

```sql
select
  id,
  name
from
  qbo_company_info;
```

```
+----+-------------+
| id | name        | 
+----+-------------+
| 1  | Cloudticity | 
+----+-------------+
```