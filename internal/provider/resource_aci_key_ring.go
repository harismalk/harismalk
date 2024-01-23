// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/ciscoecosystem/aci-go-client/v2/client"
	"github.com/ciscoecosystem/aci-go-client/v2/container"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PkiKeyRingResource{}
var _ resource.ResourceWithImportState = &PkiKeyRingResource{}

func NewPkiKeyRingResource() resource.Resource {
	return &PkiKeyRingResource{}
}

// PkiKeyRingResource defines the resource implementation.
type PkiKeyRingResource struct {
	client *client.Client
}

// PkiKeyRingResourceModel describes the resource data model.
type PkiKeyRingResourceModel struct {
	Id            types.String `tfsdk:"id"`
	ParentDn      types.String `tfsdk:"parent_dn"`
	AdminState    types.String `tfsdk:"admin_state"`
	Annotation    types.String `tfsdk:"annotation"`
	Cert          types.String `tfsdk:"certificate"`
	Descr         types.String `tfsdk:"description"`
	EccCurve      types.String `tfsdk:"elliptic_curve"`
	Key           types.String `tfsdk:"key"`
	KeyType       types.String `tfsdk:"key_type"`
	Modulus       types.String `tfsdk:"modulus"`
	Name          types.String `tfsdk:"name"`
	NameAlias     types.String `tfsdk:"name_alias"`
	OwnerKey      types.String `tfsdk:"owner_key"`
	OwnerTag      types.String `tfsdk:"owner_tag"`
	Regen         types.String `tfsdk:"regenerate"`
	Tp            types.String `tfsdk:"certificate_authority"`
	TagAnnotation types.Set    `tfsdk:"annotations"`
	TagTag        types.Set    `tfsdk:"tags"`
}

// TagAnnotationPkiKeyRingResourceModel describes the resource data model for the children without relation ships.
type TagAnnotationPkiKeyRingResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

// TagTagPkiKeyRingResourceModel describes the resource data model for the children without relation ships.
type TagTagPkiKeyRingResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

type PkiKeyRingIdentifier struct {
	Name types.String
}

func (r *PkiKeyRingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of resource: aci_key_ring")
	resp.TypeName = req.ProviderTypeName + "_key_ring"
	tflog.Debug(ctx, "End metadata of resource: aci_key_ring")
}

func (r *PkiKeyRingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of resource: aci_key_ring")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The key_ring resource for the 'pkiKeyRing' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the Key Ring object.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"parent_dn": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("uni/userext/pkiext"),
				MarkdownDescription: "The distinguished name (DN) of the parent object.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
			},
			"admin_state": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("completed", "created", "reqCreated", "started", "tpSet"),
				},
				MarkdownDescription: `The current administrative state of the certificate request process.`,
			},
			"annotation": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Default:             stringdefault.StaticString(globalAnnotation),
				MarkdownDescription: `The annotation of the Key Ring object.`,
			},
			"certificate": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `A certificate is a file containing a device's public key along with signed information verifying the identity of the device.`,
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The description of the Key Ring object.`,
			},
			"elliptic_curve": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("none", "prime256v1", "secp384r1", "secp521r1"),
				},
				MarkdownDescription: `The elliptic curve used by the provided key.`,
			},
			"key": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The private key of the certificate. This sensitive value is excluded from the resource's lifecycle configuration and is not tracked by Terraform.`,
			},
			"key_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("ECC", "RSA", "indeterminate"),
				},
				MarkdownDescription: `The type used by the provided key.`,
			},
			"modulus": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("mod1024", "mod1536", "mod2048", "mod3072", "mod4096", "mod512", "none"),
				},
				MarkdownDescription: `The length of the encryption keys. A longer key length increases the difficulty of breaking the key.`,
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: `The name of the Key Ring object.`,
			},
			"name_alias": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The name alias of the Key Ring object.`,
			},
			"owner_key": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The key for enabling clients to own their data for entity correlation.`,
			},
			"owner_tag": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `A tag for enabling clients to add their own data. For example, to indicate who created this object.`,
			},
			"regenerate": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("no", "yes"),
				},
				MarkdownDescription: `Forces regeneration of the keypair. Each PKI device holds a pair of asymmetric Rivest-Shamir-Adleman (RSA) or Elliptic Curve Cryptography (ECC) encryption keys, one kept private and one made public, stored in an internal key ring.`,
			},
			"certificate_authority": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `A third-party certificate from a trusted source, or trusted point, that affirms the identity of your device. The third-party certificate is signed by the issuing certificate authority (CA or trustpoint), which can be a root CA, an intermediate CA, or a trust anchor that is part of a trust chain that leads to a root CA.`,
			},
			"annotations": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
								SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
							},
							MarkdownDescription: `The key used to uniquely identify this configuration object.`,
						},
						"value": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
								SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
							},
							MarkdownDescription: `The value of the property.`,
						},
					},
				},
			},
			"tags": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
								SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
							},
							MarkdownDescription: `The key used to uniquely identify this configuration object.`,
						},
						"value": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
								SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
							},
							MarkdownDescription: `The value of the property.`,
						},
					},
				},
			},
		},
	}
	tflog.Debug(ctx, "End schema of resource: aci_key_ring")
}

func (r *PkiKeyRingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of resource: aci_key_ring")
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
	tflog.Debug(ctx, "End configure of resource: aci_key_ring")
}

func (r *PkiKeyRingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Start create of resource: aci_key_ring")
	// On create retrieve information on current state prior to making any changes in order to determine child delete operations
	var stateData *PkiKeyRingResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)
	setPkiKeyRingId(ctx, stateData)
	getAndSetPkiKeyRingAttributes(ctx, &resp.Diagnostics, r.client, stateData)

	var data *PkiKeyRingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setPkiKeyRingId(ctx, data)

	tflog.Debug(ctx, fmt.Sprintf("Create of resource aci_key_ring with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationPkiKeyRingResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagPkiKeyRingResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getPkiKeyRingCreateJsonPayload(ctx, &resp.Diagnostics, data, stateData, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	wrapperClassMap := map[string]string{"uni/userext/pkiext": "", "certstore": "cloudCertStore"}
	for rnPrepend, wrapperClass := range wrapperClassMap {
		if strings.Contains(data.Id.ValueString(), rnPrepend) && wrapperClass != "" {
			DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s%s.json", strings.Split(data.Id.ValueString(), rnPrepend)[0], rnPrepend), "POST", jsonPayload)
			break
		} else if strings.Contains(data.Id.ValueString(), rnPrepend) && wrapperClass == "" {
			DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetPkiKeyRingAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End create of resource aci_key_ring with id '%s'", data.Id.ValueString()))
}

func (r *PkiKeyRingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Start read of resource: aci_key_ring")
	var data *PkiKeyRingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read of resource aci_key_ring with id '%s'", data.Id.ValueString()))

	getAndSetPkiKeyRingAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	if data.Id.IsNull() {
		var emptyData *PkiKeyRingResourceModel
		resp.Diagnostics.Append(resp.State.Set(ctx, &emptyData)...)
	} else {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}

	tflog.Debug(ctx, fmt.Sprintf("End read of resource aci_key_ring with id '%s'", data.Id.ValueString()))
}

func (r *PkiKeyRingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Start update of resource: aci_key_ring")
	var data *PkiKeyRingResourceModel
	var stateData *PkiKeyRingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Update of resource aci_key_ring with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationPkiKeyRingResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagPkiKeyRingResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getPkiKeyRingCreateJsonPayload(ctx, &resp.Diagnostics, data, stateData, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	wrapperClassMap := map[string]string{
		"uni/userext/pkiext": "",
		"certstore":          "cloudCertStore",
	}
	for rnPrepend, wrapperClass := range wrapperClassMap {
		if strings.Contains(data.Id.ValueString(), rnPrepend) && wrapperClass != "" {
			DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s%s.json", strings.Split(data.Id.ValueString(), rnPrepend)[0], rnPrepend), "POST", jsonPayload)
			break
		} else if strings.Contains(data.Id.ValueString(), rnPrepend) && wrapperClass == "" {
			DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetPkiKeyRingAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End update of resource aci_key_ring with id '%s'", data.Id.ValueString()))
}

func (r *PkiKeyRingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Start delete of resource: aci_key_ring")
	var data *PkiKeyRingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Delete of resource aci_key_ring with id '%s'", data.Id.ValueString()))
	jsonPayload := GetDeleteJsonPayload(ctx, &resp.Diagnostics, "pkiKeyRing", data.Id.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}
	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("End delete of resource aci_key_ring with id '%s'", data.Id.ValueString()))
}

func (r *PkiKeyRingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Start import state of resource: aci_key_ring")
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	var stateData *PkiKeyRingResourceModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &stateData)...)
	tflog.Debug(ctx, fmt.Sprintf("Import state of resource aci_key_ring with id '%s'", stateData.Id.ValueString()))

	tflog.Debug(ctx, "End import of state resource: aci_key_ring")
}

func getAndSetPkiKeyRingAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *PkiKeyRingResourceModel) {
	requestData := DoRestRequest(ctx, diags, client, fmt.Sprintf("api/mo/%s.json?rsp-subtree=children&rsp-subtree-class=%s", data.Id.ValueString(), "pkiKeyRing,tagAnnotation,tagTag"), "GET", nil)

	if diags.HasError() {
		return
	}
	if requestData.Search("imdata").Search("pkiKeyRing").Data() != nil {
		classReadInfo := requestData.Search("imdata").Search("pkiKeyRing").Data().([]interface{})
		if len(classReadInfo) == 1 {
			attributes := classReadInfo[0].(map[string]interface{})["attributes"].(map[string]interface{})
			for attributeName, attributeValue := range attributes {
				if attributeName == "dn" {
					data.Id = basetypes.NewStringValue(attributeValue.(string))
					setPkiKeyRingParentDn(ctx, attributeValue.(string), data)
				}
				if attributeName == "adminState" {
					data.AdminState = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "annotation" {
					data.Annotation = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "cert" {
					data.Cert = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "descr" {
					data.Descr = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "eccCurve" && attributeValue.(string) == "" {
					data.EccCurve = basetypes.NewStringValue("none")
				} else if attributeName == "eccCurve" {
					data.EccCurve = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "key" {
					data.Key = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "keyType" {
					data.KeyType = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "modulus" && attributeValue.(string) == "" {
					data.Modulus = basetypes.NewStringValue("none")
				} else if attributeName == "modulus" {
					data.Modulus = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "name" {
					data.Name = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "nameAlias" {
					data.NameAlias = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "ownerKey" {
					data.OwnerKey = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "ownerTag" {
					data.OwnerTag = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "regen" {
					data.Regen = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "tp" {
					data.Tp = basetypes.NewStringValue(attributeValue.(string))
				}
			}
			if data.AdminState.IsUnknown() {
				data.AdminState = types.StringNull()
			}
			if data.Annotation.IsUnknown() {
				data.Annotation = types.StringNull()
			}
			if data.Cert.IsUnknown() {
				data.Cert = types.StringNull()
			}
			if data.Descr.IsUnknown() {
				data.Descr = types.StringNull()
			}
			if data.EccCurve.IsUnknown() {
				data.EccCurve = types.StringNull()
			}
			if data.Key.IsUnknown() {
				data.Key = types.StringNull()
			}
			if data.KeyType.IsUnknown() {
				data.KeyType = types.StringNull()
			}
			if data.Modulus.IsUnknown() {
				data.Modulus = types.StringNull()
			}
			if data.Name.IsUnknown() {
				data.Name = types.StringNull()
			}
			if data.NameAlias.IsUnknown() {
				data.NameAlias = types.StringNull()
			}
			if data.OwnerKey.IsUnknown() {
				data.OwnerKey = types.StringNull()
			}
			if data.OwnerTag.IsUnknown() {
				data.OwnerTag = types.StringNull()
			}
			if data.Regen.IsUnknown() {
				data.Regen = types.StringNull()
			}
			if data.Tp.IsUnknown() {
				data.Tp = types.StringNull()
			}
			TagAnnotationPkiKeyRingList := make([]TagAnnotationPkiKeyRingResourceModel, 0)
			TagTagPkiKeyRingList := make([]TagTagPkiKeyRingResourceModel, 0)
			_, ok := classReadInfo[0].(map[string]interface{})["children"]
			if ok {
				children := classReadInfo[0].(map[string]interface{})["children"].([]interface{})
				for _, child := range children {
					for childClassName, childClassDetails := range child.(map[string]interface{}) {
						childAttributes := childClassDetails.(map[string]interface{})["attributes"].(map[string]interface{})
						if childClassName == "tagAnnotation" {
							TagAnnotationPkiKeyRing := TagAnnotationPkiKeyRingResourceModel{}
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagAnnotationPkiKeyRing.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagAnnotationPkiKeyRing.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagAnnotationPkiKeyRingList = append(TagAnnotationPkiKeyRingList, TagAnnotationPkiKeyRing)
						}
						if childClassName == "tagTag" {
							TagTagPkiKeyRing := TagTagPkiKeyRingResourceModel{}
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagTagPkiKeyRing.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagTagPkiKeyRing.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagTagPkiKeyRingList = append(TagTagPkiKeyRingList, TagTagPkiKeyRing)
						}
					}
				}
			}
			tagAnnotationSet, _ := types.SetValueFrom(ctx, data.TagAnnotation.ElementType(ctx), TagAnnotationPkiKeyRingList)
			data.TagAnnotation = tagAnnotationSet
			tagTagSet, _ := types.SetValueFrom(ctx, data.TagTag.ElementType(ctx), TagTagPkiKeyRingList)
			data.TagTag = tagTagSet
		} else {
			diags.AddError(
				"too many results in response",
				fmt.Sprintf("%v matches returned for class 'pkiKeyRing'. Please report this issue to the provider developers.", len(classReadInfo)),
			)
		}
	} else {
		data.Id = basetypes.NewStringNull()
	}
}

func getPkiKeyRingRn(ctx context.Context, data *PkiKeyRingResourceModel) string {
	rn := "keyring-{name}"
	for _, identifier := range []string{"name"} {
		fieldName := fmt.Sprintf("%s%s", strings.ToUpper(identifier[:1]), identifier[1:])
		fieldValue := reflect.ValueOf(data).Elem().FieldByName(fieldName).Interface().(basetypes.StringValue).ValueString()
		rn = strings.ReplaceAll(rn, fmt.Sprintf("{%s}", identifier), fieldValue)
	}
	return rn
}

func setPkiKeyRingParentDn(ctx context.Context, dn string, data *PkiKeyRingResourceModel) {
	bracketIndex := 0
	rnIndex := 0
	for i := len(dn) - 1; i >= 0; i-- {
		if string(dn[i]) == "]" {
			bracketIndex = bracketIndex + 1
		} else if string(dn[i]) == "[" {
			bracketIndex = bracketIndex - 1
		} else if string(dn[i]) == "/" && bracketIndex == 0 {
			rnIndex = i
			break
		}
	}
	parentDn := dn[:rnIndex]
	rnMap := map[string]string{
		"tn": "certstore",
	}
	for parent_identifier, rn_prepend := range rnMap {
		if strings.Contains(parentDn, parent_identifier) {
			parentDn = parentDn[:strings.Index(parentDn, fmt.Sprintf("/%s", rn_prepend))]
			data.ParentDn = basetypes.NewStringValue(parentDn)
			break
		}
	}
}

func setPkiKeyRingId(ctx context.Context, data *PkiKeyRingResourceModel) {
	rn := getPkiKeyRingRn(ctx, data)
	rnMap := map[string]string{
		"tn": "certstore",
	}
	data.Id = types.StringValue(fmt.Sprintf("%s/%s", data.ParentDn.ValueString(), rn))
	for parent_identifier, rn_prepend := range rnMap {
		if strings.Contains(data.ParentDn.ValueString(), parent_identifier) {
			data.Id = types.StringValue(fmt.Sprintf("%s/%s/%s", data.ParentDn.ValueString(), rn_prepend, rn))
			break
		}
	}
}

func getPkiKeyRingTagAnnotationChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *PkiKeyRingResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationPkiKeyRingResourceModel) []map[string]interface{} {

	childPayloads := []map[string]interface{}{}
	if !data.TagAnnotation.IsUnknown() {
		tagAnnotationIdentifiers := []TagAnnotationIdentifier{}
		for _, tagAnnotation := range tagAnnotationPlan {
			childMap := map[string]map[string]interface{}{"attributes": {}}
			if !tagAnnotation.Key.IsUnknown() {
				childMap["attributes"]["key"] = tagAnnotation.Key.ValueString()
			}
			if !tagAnnotation.Value.IsUnknown() {
				childMap["attributes"]["value"] = tagAnnotation.Value.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"tagAnnotation": childMap})
			tagAnnotationIdentifier := TagAnnotationIdentifier{}
			tagAnnotationIdentifier.Key = tagAnnotation.Key
			tagAnnotationIdentifiers = append(tagAnnotationIdentifiers, tagAnnotationIdentifier)
		}
		for _, tagAnnotation := range tagAnnotationState {
			delete := true
			for _, tagAnnotationIdentifier := range tagAnnotationIdentifiers {
				if tagAnnotationIdentifier.Key == tagAnnotation.Key {
					delete = false
					break
				}
			}
			if delete {
				childMap := map[string]map[string]interface{}{"attributes": {}}
				childMap["attributes"]["status"] = "deleted"
				childMap["attributes"]["key"] = tagAnnotation.Key.ValueString()
				childPayloads = append(childPayloads, map[string]interface{}{"tagAnnotation": childMap})
			}
		}
	} else {
		data.TagAnnotation = types.SetNull(data.TagAnnotation.ElementType(ctx))
	}

	return childPayloads
}
func getPkiKeyRingTagTagChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *PkiKeyRingResourceModel, tagTagPlan, tagTagState []TagTagPkiKeyRingResourceModel) []map[string]interface{} {

	childPayloads := []map[string]interface{}{}
	if !data.TagTag.IsUnknown() {
		tagTagIdentifiers := []TagTagIdentifier{}
		for _, tagTag := range tagTagPlan {
			childMap := map[string]map[string]interface{}{"attributes": {}}
			if !tagTag.Key.IsUnknown() {
				childMap["attributes"]["key"] = tagTag.Key.ValueString()
			}
			if !tagTag.Value.IsUnknown() {
				childMap["attributes"]["value"] = tagTag.Value.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"tagTag": childMap})
			tagTagIdentifier := TagTagIdentifier{}
			tagTagIdentifier.Key = tagTag.Key
			tagTagIdentifiers = append(tagTagIdentifiers, tagTagIdentifier)
		}
		for _, tagTag := range tagTagState {
			delete := true
			for _, tagTagIdentifier := range tagTagIdentifiers {
				if tagTagIdentifier.Key == tagTag.Key {
					delete = false
					break
				}
			}
			if delete {
				childMap := map[string]map[string]interface{}{"attributes": {}}
				childMap["attributes"]["status"] = "deleted"
				childMap["attributes"]["key"] = tagTag.Key.ValueString()
				childPayloads = append(childPayloads, map[string]interface{}{"tagTag": childMap})
			}
		}
	} else {
		data.TagTag = types.SetNull(data.TagTag.ElementType(ctx))
	}

	return childPayloads
}

func getPkiKeyRingCreateJsonPayload(ctx context.Context, diags *diag.Diagnostics, data, stateData *PkiKeyRingResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationPkiKeyRingResourceModel, tagTagPlan, tagTagState []TagTagPkiKeyRingResourceModel) *container.Container {
	payloadMap := map[string]interface{}{
		"attributes": map[string]string{},
	}
	childPayloads := []map[string]interface{}{}
	TagAnnotationchildPayloads := getPkiKeyRingTagAnnotationChildPayloads(ctx, diags, data, tagAnnotationPlan, tagAnnotationState)
	if TagAnnotationchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagAnnotationchildPayloads...)
	TagTagchildPayloads := getPkiKeyRingTagTagChildPayloads(ctx, diags, data, tagTagPlan, tagTagState)
	if TagTagchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagTagchildPayloads...)
	payloadMap["children"] = childPayloads
	if !data.AdminState.IsNull() && !data.AdminState.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["adminState"] = data.AdminState.ValueString()
	}
	if !data.Annotation.IsNull() && !data.Annotation.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["annotation"] = data.Annotation.ValueString()
	}
	if !data.Cert.IsNull() && !data.Cert.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["cert"] = data.Cert.ValueString()
	}
	if !data.Descr.IsNull() && !data.Descr.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["descr"] = data.Descr.ValueString()
	}
	if !data.EccCurve.IsNull() && !data.EccCurve.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["eccCurve"] = data.EccCurve.ValueString()
	}
	if !data.Key.IsNull() && !data.Key.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["key"] = data.Key.ValueString()
	}
	if !data.KeyType.IsNull() && !data.KeyType.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["keyType"] = data.KeyType.ValueString()
	}
	if !data.Modulus.IsNull() && !data.Modulus.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["modulus"] = data.Modulus.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["name"] = data.Name.ValueString()
	}
	if !data.NameAlias.IsNull() && !data.NameAlias.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["nameAlias"] = data.NameAlias.ValueString()
	}
	if !data.OwnerKey.IsNull() && !data.OwnerKey.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["ownerKey"] = data.OwnerKey.ValueString()
	}
	if !data.OwnerTag.IsNull() && !data.OwnerTag.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["ownerTag"] = data.OwnerTag.ValueString()
	}
	if !data.Regen.IsNull() && !data.Regen.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["regen"] = data.Regen.ValueString()
	}
	if !data.Tp.IsNull() && !data.Tp.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["tp"] = data.Tp.ValueString()
	}

	wrapperClassMap := map[string]string{
		"uni/userext/pkiext": "",
		"certstore":          "cloudCertStore",
	}

	var payload []byte
	var err error
	for rnPrepend, wrapperClass := range wrapperClassMap {
		if strings.Contains(data.Id.ValueString(), rnPrepend) && wrapperClass != "" && stateData.Id.ValueString() == "" {
			wrapperPayloadMap := map[string]interface{}{
				wrapperClass: map[string]interface{}{
					"attributes": map[string]interface{}{},
					"children":   []interface{}{map[string]interface{}{"pkiKeyRing": payloadMap}},
				},
			}
			payload, err = json.Marshal(wrapperPayloadMap)
			break
		} else if (strings.Contains(data.Id.ValueString(), rnPrepend) && wrapperClass == "") || stateData.Id.ValueString() != "" {
			payload, err = json.Marshal(map[string]interface{}{"pkiKeyRing": payloadMap})
			break
		}
	}

	if err != nil {
		diags.AddError(
			"Marshalling of json payload failed",
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}

	jsonPayload, err := container.ParseJSON(payload)
	if err != nil {
		diags.AddError(
			"Construction of json payload failed",
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}

	return jsonPayload
}
