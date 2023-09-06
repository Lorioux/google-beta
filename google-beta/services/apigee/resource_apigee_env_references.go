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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/lorioux/google-beta/google-beta/tpgresource"
	transport_tpg "github.com/lorioux/google-beta/google-beta/transport"
)

func ResourceApigeeEnvReferences() *schema.Resource {
	return &schema.Resource{
		Create: resourceApigeeEnvReferencesCreate,
		Read:   resourceApigeeEnvReferencesRead,
		Delete: resourceApigeeEnvReferencesDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApigeeEnvReferencesImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
			Delete: schema.DefaultTimeout(1 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"env_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The Apigee environment group associated with the Apigee environment,
in the format 'organizations/{{org_name}}/environments/{{env_name}}'.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Required. The resource id of this reference. Values must match the regular expression [\w\s-.]+.`,
			},
			"refers": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Required. The id of the resource to which this reference refers. Must be the id of a resource that exists in the parent environment and is of the given resourceType.`,
			},
			"resource_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The type of resource referred to by this reference. Valid values are 'KeyStore' or 'TrustStore'.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Optional. A human-readable description of this reference.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceApigeeEnvReferencesCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandApigeeEnvReferencesName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandApigeeEnvReferencesDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	resourceTypeProp, err := expandApigeeEnvReferencesResourceType(d.Get("resource_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("resource_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(resourceTypeProp)) && (ok || !reflect.DeepEqual(v, resourceTypeProp)) {
		obj["resourceType"] = resourceTypeProp
	}
	refersProp, err := expandApigeeEnvReferencesRefers(d.Get("refers"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("refers"); !tpgresource.IsEmptyValue(reflect.ValueOf(refersProp)) && (ok || !reflect.DeepEqual(v, refersProp)) {
		obj["refers"] = refersProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}{{env_id}}/references/")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new EnvReferences: %#v", obj)
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
		return fmt.Errorf("Error creating EnvReferences: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{env_id}}/references/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating EnvReferences %q: %#v", d.Id(), res)

	return resourceApigeeEnvReferencesRead(d, meta)
}

func resourceApigeeEnvReferencesRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}{{env_id}}/references/{{name}}")
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ApigeeEnvReferences %q", d.Id()))
	}

	if err := d.Set("name", flattenApigeeEnvReferencesName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading EnvReferences: %s", err)
	}
	if err := d.Set("description", flattenApigeeEnvReferencesDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading EnvReferences: %s", err)
	}
	if err := d.Set("resource_type", flattenApigeeEnvReferencesResourceType(res["resourceType"], d, config)); err != nil {
		return fmt.Errorf("Error reading EnvReferences: %s", err)
	}
	if err := d.Set("refers", flattenApigeeEnvReferencesRefers(res["refers"], d, config)); err != nil {
		return fmt.Errorf("Error reading EnvReferences: %s", err)
	}

	return nil
}

func resourceApigeeEnvReferencesDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}{{env_id}}/references/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting EnvReferences %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "EnvReferences")
	}

	log.Printf("[DEBUG] Finished deleting EnvReferences %q: %#v", d.Id(), res)
	return nil
}

func resourceApigeeEnvReferencesImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats cannot import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{
		"(?P<env_id>.+)/references/(?P<name>.+)",
		"(?P<env_id>.+)/(?P<name>.+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{env_id}}/references/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenApigeeEnvReferencesName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeEnvReferencesDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeEnvReferencesResourceType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeEnvReferencesRefers(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandApigeeEnvReferencesName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvReferencesDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvReferencesResourceType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvReferencesRefers(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
