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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MgmtSubnetResource{}
var _ resource.ResourceWithImportState = &MgmtSubnetResource{}

func NewMgmtSubnetResource() resource.Resource {
	return &MgmtSubnetResource{}
}

// MgmtSubnetResource defines the resource implementation.
type MgmtSubnetResource struct {
	client *client.Client
}

// MgmtSubnetResourceModel describes the resource data model.
type MgmtSubnetResourceModel struct {
	Id            types.String `tfsdk:"id"`
	ParentDn      types.String `tfsdk:"parent_dn"`
	Annotation    types.String `tfsdk:"annotation"`
	Descr         types.String `tfsdk:"description"`
	Ip            types.String `tfsdk:"ip"`
	Name          types.String `tfsdk:"name"`
	NameAlias     types.String `tfsdk:"name_alias"`
	TagAnnotation types.Set    `tfsdk:"annotations"`
	TagTag        types.Set    `tfsdk:"tags"`
}

// TagAnnotationMgmtSubnetResourceModel describes the resource data model for the children without relation ships.
type TagAnnotationMgmtSubnetResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

func TagAnnotationMgmtSubnetResourceModelAttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"key":   types.StringType,
		"value": types.StringType,
	}
}

func TagAnnotationMgmtSubnetResourceModelElementType() attr.TypeWithAttributeTypes {
	return basetypes.ObjectType.WithAttributeTypes(basetypes.ObjectType{}, TagAnnotationMgmtSubnetResourceModelAttributeTypes())
}

// TagTagMgmtSubnetResourceModel describes the resource data model for the children without relation ships.
type TagTagMgmtSubnetResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

func TagTagMgmtSubnetResourceModelAttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"key":   types.StringType,
		"value": types.StringType,
	}
}

func TagTagMgmtSubnetResourceModelElementType() attr.TypeWithAttributeTypes {
	return basetypes.ObjectType.WithAttributeTypes(basetypes.ObjectType{}, TagTagMgmtSubnetResourceModelAttributeTypes())
}

type MgmtSubnetIdentifier struct {
	Ip types.String
}

func (r *MgmtSubnetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of resource: aci_external_management_network_subnet")
	resp.TypeName = req.ProviderTypeName + "_external_management_network_subnet"
	tflog.Debug(ctx, "End metadata of resource: aci_external_management_network_subnet")
}

func (r *MgmtSubnetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of resource: aci_external_management_network_subnet")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The external_management_network_subnet resource for the 'mgmtSubnet' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the External Management Network Subnet object.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"parent_dn": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The distinguished name (DN) of the parent object.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
			},
			"annotation": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Default:             stringdefault.StaticString(globalAnnotation),
				MarkdownDescription: `The annotation of the External Management Network Subnet object.`,
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The description of the External Management Network Subnet object.`,
			},
			"ip": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: `The external subnet IP address and subnet mask. This IP address is used for creating an external management entity. The subnet mask for the IP address to be imported from the outside into the fabric. The contracts associated with its parent instance profile (l3ext:InstP) are applied to the subnet.`,
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The name of the External Management Network Subnet object.`,
			},
			"name_alias": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The name alias of the External Management Network Subnet object.`,
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
	tflog.Debug(ctx, "End schema of resource: aci_external_management_network_subnet")
}

func (r *MgmtSubnetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of resource: aci_external_management_network_subnet")
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
	tflog.Debug(ctx, "End configure of resource: aci_external_management_network_subnet")
}

func (r *MgmtSubnetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Start create of resource: aci_external_management_network_subnet")
	// On create retrieve information on current state prior to making any changes in order to determine child delete operations
	var stateData *MgmtSubnetResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)
	setMgmtSubnetId(ctx, stateData)
	getAndSetMgmtSubnetAttributes(ctx, &resp.Diagnostics, r.client, stateData)

	var data *MgmtSubnetResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setMgmtSubnetId(ctx, data)

	tflog.Debug(ctx, fmt.Sprintf("Create of resource aci_external_management_network_subnet with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationMgmtSubnetResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagMgmtSubnetResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getMgmtSubnetCreateJsonPayload(ctx, &resp.Diagnostics, data, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetMgmtSubnetAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End create of resource aci_external_management_network_subnet with id '%s'", data.Id.ValueString()))
}

func (r *MgmtSubnetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Start read of resource: aci_external_management_network_subnet")
	var data *MgmtSubnetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read of resource aci_external_management_network_subnet with id '%s'", data.Id.ValueString()))

	getAndSetMgmtSubnetAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	if data.Id.IsNull() {
		var emptyData *MgmtSubnetResourceModel
		resp.Diagnostics.Append(resp.State.Set(ctx, &emptyData)...)
	} else {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}

	tflog.Debug(ctx, fmt.Sprintf("End read of resource aci_external_management_network_subnet with id '%s'", data.Id.ValueString()))
}

func (r *MgmtSubnetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Start update of resource: aci_external_management_network_subnet")
	var data *MgmtSubnetResourceModel
	var stateData *MgmtSubnetResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Update of resource aci_external_management_network_subnet with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationMgmtSubnetResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagMgmtSubnetResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getMgmtSubnetCreateJsonPayload(ctx, &resp.Diagnostics, data, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetMgmtSubnetAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End update of resource aci_external_management_network_subnet with id '%s'", data.Id.ValueString()))
}

func (r *MgmtSubnetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Start delete of resource: aci_external_management_network_subnet")
	var data *MgmtSubnetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Delete of resource aci_external_management_network_subnet with id '%s'", data.Id.ValueString()))
	jsonPayload := GetDeleteJsonPayload(ctx, &resp.Diagnostics, "mgmtSubnet", data.Id.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}
	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("End delete of resource aci_external_management_network_subnet with id '%s'", data.Id.ValueString()))
}

func (r *MgmtSubnetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Start import state of resource: aci_external_management_network_subnet")
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	var stateData *MgmtSubnetResourceModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &stateData)...)
	tflog.Debug(ctx, fmt.Sprintf("Import state of resource aci_external_management_network_subnet with id '%s'", stateData.Id.ValueString()))

	tflog.Debug(ctx, "End import of state resource: aci_external_management_network_subnet")
}

func getAndSetMgmtSubnetAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *MgmtSubnetResourceModel) {
	requestData := DoRestRequest(ctx, diags, client, fmt.Sprintf("api/mo/%s.json?rsp-subtree=full&rsp-subtree-class=%s", data.Id.ValueString(), "mgmtSubnet,tagAnnotation,tagTag"), "GET", nil)

	if diags.HasError() {
		return
	}
	if requestData.Search("imdata").Search("mgmtSubnet").Data() != nil {
		classReadInfo := requestData.Search("imdata").Search("mgmtSubnet").Data().([]interface{})
		if len(classReadInfo) == 1 {
			attributes := classReadInfo[0].(map[string]interface{})["attributes"].(map[string]interface{})
			for attributeName, attributeValue := range attributes {
				if attributeName == "dn" {
					data.Id = basetypes.NewStringValue(attributeValue.(string))
					setMgmtSubnetParentDn(ctx, attributeValue.(string), data)
				}
				if attributeName == "annotation" {
					data.Annotation = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "descr" {
					data.Descr = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "ip" {
					data.Ip = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "name" {
					data.Name = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "nameAlias" {
					data.NameAlias = basetypes.NewStringValue(attributeValue.(string))
				}
			}
			if data.Annotation.IsUnknown() {
				data.Annotation = types.StringNull()
			}
			if data.Descr.IsUnknown() {
				data.Descr = types.StringNull()
			}
			if data.Ip.IsUnknown() {
				data.Ip = types.StringNull()
			}
			if data.Name.IsUnknown() {
				data.Name = types.StringNull()
			}
			if data.NameAlias.IsUnknown() {
				data.NameAlias = types.StringNull()
			}
			TagAnnotationMgmtSubnet := TagAnnotationMgmtSubnetResourceModel{}
			TagAnnotationMgmtSubnetList := make([]TagAnnotationMgmtSubnetResourceModel, 0)
			TagTagMgmtSubnet := TagTagMgmtSubnetResourceModel{}
			TagTagMgmtSubnetList := make([]TagTagMgmtSubnetResourceModel, 0)
			_, ok := classReadInfo[0].(map[string]interface{})["children"]
			if ok {
				children := classReadInfo[0].(map[string]interface{})["children"].([]interface{})
				for _, child := range children {
					for childClassName, childClassDetails := range child.(map[string]interface{}) {
						childAttributes := childClassDetails.(map[string]interface{})["attributes"].(map[string]interface{})
						if childClassName == "tagAnnotation" {
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagAnnotationMgmtSubnet.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagAnnotationMgmtSubnet.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}

							}
							TagAnnotationMgmtSubnetList = append(TagAnnotationMgmtSubnetList, TagAnnotationMgmtSubnet)
						}
						if childClassName == "tagTag" {
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagTagMgmtSubnet.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagTagMgmtSubnet.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}

							}
							TagTagMgmtSubnetList = append(TagTagMgmtSubnetList, TagTagMgmtSubnet)
						}
					}
				}
			}
			tagAnnotationSet, _ := types.SetValueFrom(ctx, data.TagAnnotation.ElementType(ctx), TagAnnotationMgmtSubnetList)
			data.TagAnnotation = tagAnnotationSet
			tagTagSet, _ := types.SetValueFrom(ctx, data.TagTag.ElementType(ctx), TagTagMgmtSubnetList)
			data.TagTag = tagTagSet
		} else {
			diags.AddError(
				"too many results in response",
				fmt.Sprintf("%v matches returned for class 'mgmtSubnet'. Please report this issue to the provider developers.", len(classReadInfo)),
			)
		}
	} else {
		data.Id = basetypes.NewStringNull()
	}
}

func getMgmtSubnetRn(ctx context.Context, data *MgmtSubnetResourceModel) string {
	rn := "subnet-[{ip}]"
	for _, identifier := range []string{"ip"} {
		fieldName := fmt.Sprintf("%s%s", strings.ToUpper(identifier[:1]), identifier[1:])
		fieldValue := reflect.ValueOf(data).Elem().FieldByName(fieldName).Interface().(basetypes.StringValue).ValueString()
		rn = strings.ReplaceAll(rn, fmt.Sprintf("{%s}", identifier), fieldValue)
	}
	return rn
}

func setMgmtSubnetParentDn(ctx context.Context, dn string, data *MgmtSubnetResourceModel) {
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
	data.ParentDn = basetypes.NewStringValue(dn[:rnIndex])
}

func setMgmtSubnetId(ctx context.Context, data *MgmtSubnetResourceModel) {
	rn := getMgmtSubnetRn(ctx, data)
	data.Id = types.StringValue(fmt.Sprintf("%s/%s", data.ParentDn.ValueString(), rn))
}

func getMgmtSubnetTagAnnotationChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *MgmtSubnetResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationMgmtSubnetResourceModel) []map[string]interface{} {
	childMap := NewAciObject()
	childPayloads := []map[string]interface{}{}
	if !data.TagAnnotation.IsNull() && !data.TagAnnotation.IsUnknown() {
		tagAnnotationIdentifiers := []TagAnnotationIdentifier{}
		for _, tagAnnotation := range tagAnnotationPlan {
			if !tagAnnotation.Key.IsNull() && !tagAnnotation.Key.IsUnknown() {
				childMap.Attributes["key"] = tagAnnotation.Key.ValueString()
			}
			if !tagAnnotation.Value.IsNull() && !tagAnnotation.Value.IsUnknown() {
				childMap.Attributes["value"] = tagAnnotation.Value.ValueString()
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
				tagAnnotationChildMapForDelete := NewAciObject()
				tagAnnotationChildMapForDelete.Attributes["status"] = "deleted"
				tagAnnotationChildMapForDelete.Attributes["key"] = tagAnnotation.Key.ValueString()
				childPayloads = append(childPayloads, map[string]interface{}{"tagAnnotation": tagAnnotationChildMapForDelete})
			}
		}
	} else {
		data.TagAnnotation = types.SetNull(data.TagAnnotation.ElementType(ctx))
	}

	return childPayloads
}

func getMgmtSubnetTagTagChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *MgmtSubnetResourceModel, tagTagPlan, tagTagState []TagTagMgmtSubnetResourceModel) []map[string]interface{} {
	childMap := NewAciObject()
	childPayloads := []map[string]interface{}{}
	if !data.TagTag.IsNull() && !data.TagTag.IsUnknown() {
		tagTagIdentifiers := []TagTagIdentifier{}
		for _, tagTag := range tagTagPlan {
			if !tagTag.Key.IsNull() && !tagTag.Key.IsUnknown() {
				childMap.Attributes["key"] = tagTag.Key.ValueString()
			}
			if !tagTag.Value.IsNull() && !tagTag.Value.IsUnknown() {
				childMap.Attributes["value"] = tagTag.Value.ValueString()
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
				tagTagChildMapForDelete := NewAciObject()
				tagTagChildMapForDelete.Attributes["status"] = "deleted"
				tagTagChildMapForDelete.Attributes["key"] = tagTag.Key.ValueString()
				childPayloads = append(childPayloads, map[string]interface{}{"tagTag": tagTagChildMapForDelete})
			}
		}
	} else {
		data.TagTag = types.SetNull(data.TagTag.ElementType(ctx))
	}

	return childPayloads
}

func getMgmtSubnetCreateJsonPayload(ctx context.Context, diags *diag.Diagnostics, data *MgmtSubnetResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationMgmtSubnetResourceModel, tagTagPlan, tagTagState []TagTagMgmtSubnetResourceModel) *container.Container {
	payloadMap := map[string]interface{}{}
	payloadMap["attributes"] = map[string]string{}
	childPayloads := []map[string]interface{}{}

	TagAnnotationchildPayloads := getMgmtSubnetTagAnnotationChildPayloads(ctx, diags, data, tagAnnotationPlan, tagAnnotationState)
	if TagAnnotationchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagAnnotationchildPayloads...)

	TagTagchildPayloads := getMgmtSubnetTagTagChildPayloads(ctx, diags, data, tagTagPlan, tagTagState)
	if TagTagchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagTagchildPayloads...)

	payloadMap["children"] = childPayloads
	if !data.Annotation.IsNull() && !data.Annotation.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["annotation"] = data.Annotation.ValueString()
	}
	if !data.Descr.IsNull() && !data.Descr.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["descr"] = data.Descr.ValueString()
	}
	if !data.Ip.IsNull() && !data.Ip.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["ip"] = data.Ip.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["name"] = data.Name.ValueString()
	}
	if !data.NameAlias.IsNull() && !data.NameAlias.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["nameAlias"] = data.NameAlias.ValueString()
	}

	payload, err := json.Marshal(map[string]interface{}{"mgmtSubnet": payloadMap})
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
