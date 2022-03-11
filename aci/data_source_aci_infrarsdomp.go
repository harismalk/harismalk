package aci

import (
	"context"
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAciInfraRsDomP() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceAciInfraRsDomPRead,
		SchemaVersion: 1,
		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"attachable_access_entity_profile_dn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"annotation": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"t_dn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func dataSourceAciInfraRsDomPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	aciClient := m.(*client.Client)
	tDn := d.Get("tDn").(string)
	AttachableAccessEntityProfileDn := d.Get("attachable_access_entity_profile_dn").(string)
	rn := fmt.Sprintf("rsdomP-[%s]", tDn)
	dn := fmt.Sprintf("%s/%s", AttachableAccessEntityProfileDn, rn)

	infraRsDomP, err := getRemoteInfraRsDomP(aciClient, dn)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dn)

	_, err = setInfraRsDomPAttributes(infraRsDomP, d)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
