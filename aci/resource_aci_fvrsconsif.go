package aci

import (
	"context"
	"fmt"
	"log"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAciContractInterfaceRelationship() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAciContractInterfaceRelationshipCreate,
		UpdateContext: resourceAciContractInterfaceRelationshipUpdate,
		ReadContext:   resourceAciContractInterfaceRelationshipRead,
		DeleteContext: resourceAciContractInterfaceRelationshipDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciContractInterfaceRelationshipImport,
		},

		SchemaVersion: 1,
		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"application_epg_dn": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"prio": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"level1",
					"level2",
					"level3",
					"level4",
					"level5",
					"level6",
					"unspecified",
				}, false),
			},
			"tn_vz_cp_if_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		}),
	}
}

func getRemoteContractInterfaceRelationship(client *client.Client, dn string) (*models.ContractInterfaceRelationship, error) {
	fvRsConsIfCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}
	fvRsConsIf := models.ContractInterfaceRelationshipFromContainer(fvRsConsIfCont)
	if fvRsConsIf.DistinguishedName == "" {
		return nil, fmt.Errorf("ContractInterfaceRelationship %s not found", fvRsConsIf.DistinguishedName)
	}
	return fvRsConsIf, nil
}

func setContractInterfaceRelationshipAttributes(fvRsConsIf *models.ContractInterfaceRelationship, d *schema.ResourceData) (*schema.ResourceData, error) {
	dn := d.Id()
	d.SetId(fvRsConsIf.DistinguishedName)
	if dn != fvRsConsIf.DistinguishedName {
		d.Set("application_epg_dn", "")
	}
	fvRsConsIfMap, err := fvRsConsIf.ToMap()
	if err != nil {
		return d, err
	}
	d.Set("prio", fvRsConsIfMap["prio"])
	d.Set("tn_vz_cp_if_name", fvRsConsIfMap["tnVzCPIfName"])
	return d, nil
}

func resourceAciContractInterfaceRelationshipImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	log.Printf("[DEBUG] %s: Beginning Import", d.Id())
	aciClient := m.(*client.Client)
	dn := d.Id()
	fvRsConsIf, err := getRemoteContractInterfaceRelationship(aciClient, dn)
	if err != nil {
		return nil, err
	}
	schemaFilled, err := setContractInterfaceRelationshipAttributes(fvRsConsIf, d)
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] %s: Import finished successfully", d.Id())
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciContractInterfaceRelationshipCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] ContractInterfaceRelationship: Beginning Creation")
	aciClient := m.(*client.Client)
	tnVzCPIfName := d.Get("tn_vz_cp_if_name").(string)
	ApplicationEPGDn := d.Get("application_epg_dn").(string)

	fvRsConsIfAttr := models.ContractInterfaceRelationshipAttributes{}

	if Annotation, ok := d.GetOk("annotation"); ok {
		fvRsConsIfAttr.Annotation = Annotation.(string)
	} else {
		fvRsConsIfAttr.Annotation = "{}"
	}

	if Prio, ok := d.GetOk("prio"); ok {
		fvRsConsIfAttr.Prio = Prio.(string)
	}

	if TnVzCPIfName, ok := d.GetOk("tn_vz_cp_if_name"); ok {
		fvRsConsIfAttr.TnVzCPIfName = TnVzCPIfName.(string)
	}
	fvRsConsIf := models.NewContractInterfaceRelationship(fmt.Sprintf(models.RnfvRsConsIf, tnVzCPIfName), ApplicationEPGDn, fvRsConsIfAttr)

	err := aciClient.Save(fvRsConsIf)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fvRsConsIf.DistinguishedName)
	log.Printf("[DEBUG] %s: Creation finished successfully", d.Id())
	return resourceAciContractInterfaceRelationshipRead(ctx, d, m)
}

func resourceAciContractInterfaceRelationshipUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] ContractInterfaceRelationship: Beginning Update")
	aciClient := m.(*client.Client)
	tnVzCPIfName := d.Get("tn_vz_cp_if_name").(string)
	ApplicationEPGDn := d.Get("application_epg_dn").(string)

	fvRsConsIfAttr := models.ContractInterfaceRelationshipAttributes{}

	if Annotation, ok := d.GetOk("annotation"); ok {
		fvRsConsIfAttr.Annotation = Annotation.(string)
	} else {
		fvRsConsIfAttr.Annotation = "{}"
	}

	if Prio, ok := d.GetOk("prio"); ok {
		fvRsConsIfAttr.Prio = Prio.(string)
	}

	if TnVzCPIfName, ok := d.GetOk("tn_vz_cp_if_name"); ok {
		fvRsConsIfAttr.TnVzCPIfName = TnVzCPIfName.(string)
	}
	fvRsConsIf := models.NewContractInterfaceRelationship(fmt.Sprintf("rsconsIf-%s", tnVzCPIfName), ApplicationEPGDn, fvRsConsIfAttr)

	fvRsConsIf.Status = "modified"

	err := aciClient.Save(fvRsConsIf)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fvRsConsIf.DistinguishedName)
	log.Printf("[DEBUG] %s: Update finished successfully", d.Id())
	return resourceAciContractInterfaceRelationshipRead(ctx, d, m)
}

func resourceAciContractInterfaceRelationshipRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning Read", d.Id())
	aciClient := m.(*client.Client)
	dn := d.Id()

	fvRsConsIf, err := getRemoteContractInterfaceRelationship(aciClient, dn)
	if err != nil {
		d.SetId("")
		return diag.FromErr(err)
	}

	_, err = setContractInterfaceRelationshipAttributes(fvRsConsIf, d)
	if err != nil {
		d.SetId("")
		return nil
	}

	log.Printf("[DEBUG] %s: Read finished successfully", d.Id())
	return nil
}

func resourceAciContractInterfaceRelationshipDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning Destroy", d.Id())
	aciClient := m.(*client.Client)
	dn := d.Id()

	err := aciClient.DeleteByDn(dn, "fvRsConsIf")
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] %s: Destroy finished successfully", d.Id())
	d.SetId("")
	return diag.FromErr(err)
}