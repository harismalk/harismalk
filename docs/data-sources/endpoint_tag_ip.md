---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Generic"
layout: "aci"
page_title: "ACI: aci_endpoint_tag_ip"
sidebar_current: "docs-aci-data-source-aci_endpoint_tag_ip"
description: |-
  Data source for Endpoint Tag Ip
---

# aci_endpoint_tag_ip #

Data source for Endpoint Tag Ip

## API Information ##

* Class: [fvEpIpTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvEpIpTag/overview)


* Distinguished Name Format: `uni/tn-{name}/eptags/epiptag-[{ip}]-{ctxName}`

## GUI Information ##

* Location: `Generic`

## Example Usage ##

```hcl

data "aci_endpoint_tag_ip" "example_tenant" {
  parent_dn = aci_tenant.example.id
  vrf_name  = "test_ctx_name"
  ip        = "10.0.0.2"
}

```

## Schema ##

### Required ###

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - [aci_tenant](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/tenant) ([fvTenant](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvTenant/overview))
* `vrf_name` (ctxName) - (string) The VRF name of the Endpoint Tag Ip object.
* `ip` (ip) - (string) The IP address of the Endpoint Tag Ip object.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the Endpoint Tag Ip object.
* `annotation` (annotation) - (string) The annotation of the Endpoint Tag Ip object.
* `id_attribute` (id) - (string) The identifier of the Endpoint Tag Ip object.
* `name` (name) - (string) The name of the Endpoint Tag Ip object.
* `name_alias` (nameAlias) - (string) The name alias of the Endpoint Tag Ip object.

* `annotations` - (list) A list of Annotations objects ([tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
  * `key` (key) - (string) The key used to uniquely identify this configuration object.
  * `value` (value) - (string) The value of the property.