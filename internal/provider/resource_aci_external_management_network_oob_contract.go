// Code generated by "gen/generator.go"; DO NOT EDIT.

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
var _ resource.Resource = &MgmtRsOoBConsResource{}
var _ resource.ResourceWithImportState = &MgmtRsOoBConsResource{}

func NewMgmtRsOoBConsResource() resource.Resource {
	return &MgmtRsOoBConsResource{}
}

// MgmtRsOoBConsResource defines the resource implementation.
type MgmtRsOoBConsResource struct {
	client *client.Client
}

// MgmtRsOoBConsResourceModel describes the resource data model.
type MgmtRsOoBConsResourceModel struct {
	Id              types.String `tfsdk:"id"`
	ParentDn        types.String `tfsdk:"parent_dn"`
	Annotation      types.String `tfsdk:"annotation"`
	Prio            types.String `tfsdk:"priority"`
	TnVzOOBBrCPName types.String `tfsdk:"out_of_band_contract_name"`
	TagAnnotation   types.Set    `tfsdk:"annotations"`
}

// TagAnnotationMgmtRsOoBConsResourceModel describes the resource data model for the children without relation ships.
type TagAnnotationMgmtRsOoBConsResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

type MgmtRsOoBConsIdentifier struct {
	TnVzOOBBrCPName types.String
}

func (r *MgmtRsOoBConsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of resource: aci_external_management_network_oob_contract")
	resp.TypeName = req.ProviderTypeName + "_external_management_network_oob_contract"
	tflog.Debug(ctx, "End metadata of resource: aci_external_management_network_oob_contract")
}

func (r *MgmtRsOoBConsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of resource: aci_external_management_network_oob_contract")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The external_management_network_oob_contract resource for the 'mgmtRsOoBCons' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the External Management Network Oob Contract object.",
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
				},
				Default:             stringdefault.StaticString(globalAnnotation),
				MarkdownDescription: `The annotation of the External Management Network Oob Contract object.`,
			},
			"priority": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("level1", "level2", "level3", "level4", "level5", "level6", "unspecified"),
				},
				MarkdownDescription: `The Quality of service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.`,
			},
			"out_of_band_contract_name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: `An out-of-band management endpoint group contract consists of switches (leaves/spines) and APICs that are part of the associated out-of-band management zone. Each node in the group is assigned an IP address that is dynamically allocated from the address pool associated with the corresponding out-of-band management zone.`,
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
							},
							MarkdownDescription: `The key used to uniquely identify this configuration object.`,
						},
						"value": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: `The value of the property.`,
						},
					},
				},
			},
		},
	}
	tflog.Debug(ctx, "End schema of resource: aci_external_management_network_oob_contract")
}

func (r *MgmtRsOoBConsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of resource: aci_external_management_network_oob_contract")
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
	tflog.Debug(ctx, "End configure of resource: aci_external_management_network_oob_contract")
}

func (r *MgmtRsOoBConsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Start create of resource: aci_external_management_network_oob_contract")
	// On create retrieve information on current state prior to making any changes in order to determine child delete operations
	var stateData *MgmtRsOoBConsResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)
	setMgmtRsOoBConsId(ctx, stateData)
	setMgmtRsOoBConsAttributes(ctx, &resp.Diagnostics, r.client, stateData)

	var data *MgmtRsOoBConsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setMgmtRsOoBConsId(ctx, data)

	tflog.Debug(ctx, fmt.Sprintf("Create of resource aci_external_management_network_oob_contract with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationMgmtRsOoBConsResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	jsonPayload := getMgmtRsOoBConsCreateJsonPayload(ctx, &resp.Diagnostics, data, tagAnnotationPlan, tagAnnotationState)

	if resp.Diagnostics.HasError() {
		return
	}

	doMgmtRsOoBConsRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}

	setMgmtRsOoBConsAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, "End create of resource: aci_external_management_network_oob_contract")
}

func (r *MgmtRsOoBConsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Start read of resource: aci_external_management_network_oob_contract")
	var data *MgmtRsOoBConsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read of resource aci_external_management_network_oob_contract with id '%s'", data.Id.ValueString()))

	setMgmtRsOoBConsAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	if data.Id.IsNull() {
		var emptyData *MgmtRsOoBConsResourceModel
		resp.Diagnostics.Append(resp.State.Set(ctx, &emptyData)...)
	} else {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}

	tflog.Debug(ctx, "End read of resource: aci_external_management_network_oob_contract")
}

func (r *MgmtRsOoBConsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Start update of resource: aci_external_management_network_oob_contract")
	var data *MgmtRsOoBConsResourceModel
	var stateData *MgmtRsOoBConsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Update of resource aci_external_management_network_oob_contract with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationMgmtRsOoBConsResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	jsonPayload := getMgmtRsOoBConsCreateJsonPayload(ctx, &resp.Diagnostics, data, tagAnnotationPlan, tagAnnotationState)

	if resp.Diagnostics.HasError() {
		return
	}

	doMgmtRsOoBConsRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	setMgmtRsOoBConsAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, "End update of resource: aci_external_management_network_oob_contract")
}

func (r *MgmtRsOoBConsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Start delete of resource: aci_external_management_network_oob_contract")
	var data *MgmtRsOoBConsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Delete of resource aci_external_management_network_oob_contract with id '%s'", data.Id.ValueString()))
	jsonPayload := getMgmtRsOoBConsDeleteJsonPayload(ctx, &resp.Diagnostics, data)
	if resp.Diagnostics.HasError() {
		return
	}
	doMgmtRsOoBConsRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "End delete of resource: aci_external_management_network_oob_contract")
}

func (r *MgmtRsOoBConsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Start import state of resource: aci_external_management_network_oob_contract")
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	var stateData *MgmtRsOoBConsResourceModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &stateData)...)
	tflog.Debug(ctx, fmt.Sprintf("Import state of resource aci_external_management_network_oob_contract with id '%s'", stateData.Id.ValueString()))

	tflog.Debug(ctx, "End import of state resource: aci_external_management_network_oob_contract")
}

func setMgmtRsOoBConsAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *MgmtRsOoBConsResourceModel) {
	requestData := doMgmtRsOoBConsRequest(ctx, diags, client, fmt.Sprintf("api/mo/%s.json?rsp-subtree=children&rsp-subtree-class=%s", data.Id.ValueString(), "mgmtRsOoBCons,tagAnnotation"), "GET", nil)

	if diags.HasError() {
		return
	}
	if requestData.Search("imdata").Search("mgmtRsOoBCons").Data() != nil {
		classReadInfo := requestData.Search("imdata").Search("mgmtRsOoBCons").Data().([]interface{})
		if len(classReadInfo) == 1 {
			attributes := classReadInfo[0].(map[string]interface{})["attributes"].(map[string]interface{})
			for attributeName, attributeValue := range attributes {
				if attributeName == "dn" {
					data.Id = basetypes.NewStringValue(attributeValue.(string))
					setMgmtRsOoBConsParentDn(ctx, attributeValue.(string), data)
				}
				if attributeName == "annotation" {
					data.Annotation = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "prio" {
					data.Prio = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "tnVzOOBBrCPName" {
					data.TnVzOOBBrCPName = basetypes.NewStringValue(attributeValue.(string))
				}
			}
			TagAnnotationMgmtRsOoBConsList := make([]TagAnnotationMgmtRsOoBConsResourceModel, 0)
			_, ok := classReadInfo[0].(map[string]interface{})["children"]
			if ok {
				children := classReadInfo[0].(map[string]interface{})["children"].([]interface{})
				for _, child := range children {
					for childClassName, childClassDetails := range child.(map[string]interface{}) {
						childAttributes := childClassDetails.(map[string]interface{})["attributes"].(map[string]interface{})
						if childClassName == "tagAnnotation" {
							TagAnnotationMgmtRsOoBCons := TagAnnotationMgmtRsOoBConsResourceModel{}
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagAnnotationMgmtRsOoBCons.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagAnnotationMgmtRsOoBCons.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagAnnotationMgmtRsOoBConsList = append(TagAnnotationMgmtRsOoBConsList, TagAnnotationMgmtRsOoBCons)
						}
					}
				}
			}
			if len(TagAnnotationMgmtRsOoBConsList) > 0 {
				tagAnnotationSet, _ := types.SetValueFrom(ctx, data.TagAnnotation.ElementType(ctx), TagAnnotationMgmtRsOoBConsList)
				data.TagAnnotation = tagAnnotationSet
			}
		} else {
			diags.AddError(
				"too many results in response",
				fmt.Sprintf("%v matches returned for class 'mgmtRsOoBCons'. Please report this issue to the provider developers.", len(classReadInfo)),
			)
		}
	} else {
		data.Id = basetypes.NewStringNull()
	}
}

func getMgmtRsOoBConsRn(ctx context.Context, data *MgmtRsOoBConsResourceModel) string {
	rn := "rsooBCons-{tnVzOOBBrCPName}"
	for _, identifier := range []string{"tnVzOOBBrCPName"} {
		fieldName := fmt.Sprintf("%s%s", strings.ToUpper(identifier[:1]), identifier[1:])
		fieldValue := reflect.ValueOf(data).Elem().FieldByName(fieldName).Interface().(basetypes.StringValue).ValueString()
		rn = strings.ReplaceAll(rn, fmt.Sprintf("{%s}", identifier), fieldValue)
	}
	return rn
}

func setMgmtRsOoBConsParentDn(ctx context.Context, dn string, data *MgmtRsOoBConsResourceModel) {
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

func setMgmtRsOoBConsId(ctx context.Context, data *MgmtRsOoBConsResourceModel) {
	rn := getMgmtRsOoBConsRn(ctx, data)
	data.Id = types.StringValue(fmt.Sprintf("%s/%s", data.ParentDn.ValueString(), rn))
}

func getMgmtRsOoBConsTagAnnotationChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *MgmtRsOoBConsResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationMgmtRsOoBConsResourceModel) []map[string]interface{} {

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

func getMgmtRsOoBConsCreateJsonPayload(ctx context.Context, diags *diag.Diagnostics, data *MgmtRsOoBConsResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationMgmtRsOoBConsResourceModel) *container.Container {
	payloadMap := map[string]interface{}{}
	payloadMap["attributes"] = map[string]string{}
	childPayloads := []map[string]interface{}{}

	TagAnnotationchildPayloads := getMgmtRsOoBConsTagAnnotationChildPayloads(ctx, diags, data, tagAnnotationPlan, tagAnnotationState)
	if TagAnnotationchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagAnnotationchildPayloads...)

	payloadMap["children"] = childPayloads
	if !data.Annotation.IsNull() && !data.Annotation.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["annotation"] = data.Annotation.ValueString()
	}
	if !data.Prio.IsNull() && !data.Prio.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["prio"] = data.Prio.ValueString()
	}
	if !data.TnVzOOBBrCPName.IsNull() && !data.TnVzOOBBrCPName.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["tnVzOOBBrCPName"] = data.TnVzOOBBrCPName.ValueString()
	}

	payload, err := json.Marshal(map[string]interface{}{"mgmtRsOoBCons": payloadMap})
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

func getMgmtRsOoBConsDeleteJsonPayload(ctx context.Context, diags *diag.Diagnostics, data *MgmtRsOoBConsResourceModel) *container.Container {

	jsonString := fmt.Sprintf(`{"mgmtRsOoBCons":{"attributes":{"dn": "%s","status": "deleted"}}}`, data.Id.ValueString())
	jsonPayload, err := container.ParseJSON([]byte(jsonString))
	if err != nil {
		diags.AddError(
			"Construction of json payload failed",
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}
	return jsonPayload
}

func doMgmtRsOoBConsRequest(ctx context.Context, diags *diag.Diagnostics, client *client.Client, path, method string, payload *container.Container) *container.Container {

	restRequest, err := client.MakeRestRequest(method, path, payload, true)
	if err != nil {
		diags.AddError(
			"Creation of rest request failed",
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}

	cont, restResponse, err := client.Do(restRequest)

	if restResponse != nil && restResponse.StatusCode != 200 {
		diags.AddError(
			fmt.Sprintf("The %s rest request failed", strings.ToLower(method)),
			fmt.Sprintf("Response: %s, err: %s. Please report this issue to the provider developers.", cont.Data().(map[string]interface{})["imdata"], err),
		)
		return nil
	} else if err != nil {
		diags.AddError(
			fmt.Sprintf("The %s rest request failed", strings.ToLower(method)),
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}

	return cont
}
