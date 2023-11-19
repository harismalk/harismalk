---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
subcategory: "Multicast"
layout: "aci"
page_title: "ACI: pim_route_map_policy"
sidebar_current: "docs-aci-data-source-pim_route_map_policy"
description: |-
  Data source for Pim Route Map Policy
---

# aci_pim_route_map_policy #

Data source for Pim Route Map Policy

## API Information ##

* `Class` - [pimRouteMapPol](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/pimRouteMapPol/overview)

* `Distinguished Name Formats`
  - `uni/tn-{name}/rtmap-{name}`

## GUI Information ##

* `Location` - `Tenants -> Policies -> Protocol -> Route Maps for Multicast`

## Example Usage ##

```hcl

data "aci_pim_route_map_policy" "example" {
  parent_dn = aci_tenant.example.id
  name      = "test_name"
}

```

## Schema

### Required

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - [aci_tenant](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/tenant) ([fvTenant](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvTenant/overview))
* `name` - (string) The name of the Pim Route Map Policy object.

### Read-Only

* `id` - (string) The distinguished name (DN) of the Pim Route Map Policy object.
* `annotation` - (string) The annotation of the Pim Route Map Policy object.
* `description` - (string) The description of the Pim Route Map Policy object.
* `name_alias` - (string) The name alias of the Pim Route Map Policy object.
* `owner_key` - (string) The key for enabling clients to own their data for entity correlation.
* `owner_tag` - (string) A tag for enabling clients to add their own data. For example, to indicate who created this object.

* `annotations` - (list) A list of Annotations objects ([tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)).
  * `key` - (string) The key or password used to uniquely identify this configuration object.
  * `value` - (string) The value of the property.