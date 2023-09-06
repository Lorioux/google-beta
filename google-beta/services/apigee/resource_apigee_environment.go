// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package apigee

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/lorioux/google-beta/google-beta/tpgresource"
	transport_tpg "github.com/lorioux/google-beta/google-beta/transport"
	"github.com/lorioux/google-beta/google-beta/verify"
)

func ResourceApigeeEnvironment() *schema.Resource {
	return &schema.Resource{
		Create: resourceApigeeEnvironmentCreate,
		Read:   resourceApigeeEnvironmentRead,
		Update: resourceApigeeEnvironmentUpdate,
		Delete: resourceApigeeEnvironmentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApigeeEnvironmentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The resource ID of the environment.`,
			},
			"org_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The Apigee Organization associated with the Apigee environment,
in the format 'organizations/{{org_name}}'.`,
			},
			"api_proxy_type": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"API_PROXY_TYPE_UNSPECIFIED", "PROGRAMMABLE", "CONFIGURABLE", ""}),
				Description: `Optional. API Proxy type supported by the environment. The type can be set when creating
the Environment and cannot be changed. Possible values: ["API_PROXY_TYPE_UNSPECIFIED", "PROGRAMMABLE", "CONFIGURABLE"]`,
			},
			"deployment_type": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"DEPLOYMENT_TYPE_UNSPECIFIED", "PROXY", "ARCHIVE", ""}),
				Description: `Optional. Deployment type supported by the environment. The deployment type can be
set when creating the environment and cannot be changed. When you enable archive
deployment, you will be prevented from performing a subset of actions within the
environment, including:
Managing the deployment of API proxy or shared flow revisions;
Creating, updating, or deleting resource files;
Creating, updating, or deleting target servers. Possible values: ["DEPLOYMENT_TYPE_UNSPECIFIED", "PROXY", "ARCHIVE"]`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Description of the environment.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Display name of the environment.`,
			},
			"node_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: `NodeConfig for setting the min/max number of nodes associated with the environment.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_node_count": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `The maximum total number of gateway nodes that the is reserved for all instances that
has the specified environment. If not specified, the default is determined by the
recommended maximum number of nodes for that gateway.`,
						},
						"min_node_count": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `The minimum total number of gateway nodes that the is reserved for all instances that
has the specified environment. If not specified, the default is determined by the
recommended minimum number of nodes for that gateway.`,
						},
						"current_aggregate_node_count": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `The current total number of gateway nodes that each environment currently has across
all instances.`,
						},
					},
				},
			},
		},
		UseJSONNumber: true,
	}
}

func resourceApigeeEnvironmentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandApigeeEnvironmentName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	displayNameProp, err := expandApigeeEnvironmentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandApigeeEnvironmentDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	deploymentTypeProp, err := expandApigeeEnvironmentDeploymentType(d.Get("deployment_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("deployment_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(deploymentTypeProp)) && (ok || !reflect.DeepEqual(v, deploymentTypeProp)) {
		obj["deploymentType"] = deploymentTypeProp
	}
	apiProxyTypeProp, err := expandApigeeEnvironmentApiProxyType(d.Get("api_proxy_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("api_proxy_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(apiProxyTypeProp)) && (ok || !reflect.DeepEqual(v, apiProxyTypeProp)) {
		obj["apiProxyType"] = apiProxyTypeProp
	}
	nodeConfigProp, err := expandApigeeEnvironmentNodeConfig(d.Get("node_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("node_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeConfigProp)) && (ok || !reflect.DeepEqual(v, nodeConfigProp)) {
		obj["nodeConfig"] = nodeConfigProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}{{org_id}}/environments")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Environment: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating Environment: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{org_id}}/environments/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = ApigeeOperationWaitTimeWithResponse(
		config, res, &opRes, "Creating Environment", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Environment: %s", err)
	}

	if err := d.Set("name", flattenApigeeEnvironmentName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "{{org_id}}/environments/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Environment %q: %#v", d.Id(), res)

	return resourceApigeeEnvironmentRead(d, meta)
}

func resourceApigeeEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}{{org_id}}/environments/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ApigeeEnvironment %q", d.Id()))
	}

	if err := d.Set("name", flattenApigeeEnvironmentName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Environment: %s", err)
	}
	if err := d.Set("display_name", flattenApigeeEnvironmentDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Environment: %s", err)
	}
	if err := d.Set("description", flattenApigeeEnvironmentDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Environment: %s", err)
	}
	if err := d.Set("deployment_type", flattenApigeeEnvironmentDeploymentType(res["deploymentType"], d, config)); err != nil {
		return fmt.Errorf("Error reading Environment: %s", err)
	}
	if err := d.Set("api_proxy_type", flattenApigeeEnvironmentApiProxyType(res["apiProxyType"], d, config)); err != nil {
		return fmt.Errorf("Error reading Environment: %s", err)
	}
	if err := d.Set("node_config", flattenApigeeEnvironmentNodeConfig(res["nodeConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Environment: %s", err)
	}

	return nil
}

func resourceApigeeEnvironmentUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	nodeConfigProp, err := expandApigeeEnvironmentNodeConfig(d.Get("node_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("node_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nodeConfigProp)) {
		obj["nodeConfig"] = nodeConfigProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}{{org_id}}/environments/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Environment %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("node_config") {
		updateMask = append(updateMask, "nodeConfig")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating Environment %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Environment %q: %#v", d.Id(), res)
	}

	err = ApigeeOperationWaitTime(
		config, res, "Updating Environment", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceApigeeEnvironmentRead(d, meta)
}

func resourceApigeeEnvironmentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}{{org_id}}/environments/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Environment %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Environment")
	}

	err = ApigeeOperationWaitTime(
		config, res, "Deleting Environment", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Environment %q: %#v", d.Id(), res)
	return nil
}

func resourceApigeeEnvironmentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats cannot import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	nameParts := strings.Split(d.Get("name").(string), "/")
	if len(nameParts) == 4 {
		// `organizations/{{org_name}}/environments/{{name}}`
		orgId := fmt.Sprintf("organizations/%s", nameParts[1])
		if err := d.Set("org_id", orgId); err != nil {
			return nil, fmt.Errorf("Error setting org_id: %s", err)
		}
		if err := d.Set("name", nameParts[3]); err != nil {
			return nil, fmt.Errorf("Error setting name: %s", err)
		}
	} else if len(nameParts) == 3 {
		// `organizations/{{org_name}}/{{name}}`
		orgId := fmt.Sprintf("organizations/%s", nameParts[1])
		if err := d.Set("org_id", orgId); err != nil {
			return nil, fmt.Errorf("Error setting org_id: %s", err)
		}
		if err := d.Set("name", nameParts[2]); err != nil {
			return nil, fmt.Errorf("Error setting name: %s", err)
		}
	} else {
		return nil, fmt.Errorf(
			"Saw %s when the name is expected to have shape %s or %s",
			d.Get("name"),
			"organizations/{{org_name}}/environments/{{name}}",
			"organizations/{{org_name}}/{{name}}")
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{org_id}}/environments/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenApigeeEnvironmentName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeEnvironmentDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeEnvironmentDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeEnvironmentDeploymentType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeEnvironmentApiProxyType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeEnvironmentNodeConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["min_node_count"] =
		flattenApigeeEnvironmentNodeConfigMinNodeCount(original["minNodeCount"], d, config)
	transformed["max_node_count"] =
		flattenApigeeEnvironmentNodeConfigMaxNodeCount(original["maxNodeCount"], d, config)
	transformed["current_aggregate_node_count"] =
		flattenApigeeEnvironmentNodeConfigCurrentAggregateNodeCount(original["currentAggregateNodeCount"], d, config)
	return []interface{}{transformed}
}
func flattenApigeeEnvironmentNodeConfigMinNodeCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeEnvironmentNodeConfigMaxNodeCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeEnvironmentNodeConfigCurrentAggregateNodeCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandApigeeEnvironmentName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentDeploymentType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentApiProxyType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentNodeConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinNodeCount, err := expandApigeeEnvironmentNodeConfigMinNodeCount(original["min_node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minNodeCount"] = transformedMinNodeCount
	}

	transformedMaxNodeCount, err := expandApigeeEnvironmentNodeConfigMaxNodeCount(original["max_node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxNodeCount"] = transformedMaxNodeCount
	}

	transformedCurrentAggregateNodeCount, err := expandApigeeEnvironmentNodeConfigCurrentAggregateNodeCount(original["current_aggregate_node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCurrentAggregateNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["currentAggregateNodeCount"] = transformedCurrentAggregateNodeCount
	}

	return transformed, nil
}

func expandApigeeEnvironmentNodeConfigMinNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentNodeConfigMaxNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentNodeConfigCurrentAggregateNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
