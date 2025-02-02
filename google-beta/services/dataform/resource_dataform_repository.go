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

package dataform

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/lorioux/google-beta/google-beta/tpgresource"
	transport_tpg "github.com/lorioux/google-beta/google-beta/transport"
)

func ResourceDataformRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataformRepositoryCreate,
		Read:   resourceDataformRepositoryRead,
		Update: resourceDataformRepositoryUpdate,
		Delete: resourceDataformRepositoryDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDataformRepositoryImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The repository's name.`,
			},
			"git_remote_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Optional. If set, configures this repository to be linked to a Git remote.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication_token_secret_version": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The name of the Secret Manager secret version to use as an authentication token for Git operations. Must be in the format projects/*/secrets/*/versions/*.`,
						},
						"default_branch": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The Git remote's default branch name.`,
						},
						"url": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The Git remote's URL.`,
						},
						"token_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the status of the Git access token. https://cloud.google.com/dataform/reference/rest/v1beta1/projects.locations.repositories#TokenStatus`,
						},
					},
				},
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `A reference to the region`,
			},
			"workspace_compilation_overrides": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Optional. If set, fields of workspaceCompilationOverrides override the default compilation settings that are specified in dataform.json when creating workspace-scoped compilation results.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_database": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Optional. The default database (Google Cloud project ID).`,
						},
						"schema_suffix": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Optional. The suffix that should be appended to all schema (BigQuery dataset ID) names.`,
						},
						"table_prefix": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Optional. The prefix that should be prepended to all table names.`,
						},
					},
				},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceDataformRepositoryCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandDataformRepositoryName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	gitRemoteSettingsProp, err := expandDataformRepositoryGitRemoteSettings(d.Get("git_remote_settings"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("git_remote_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(gitRemoteSettingsProp)) && (ok || !reflect.DeepEqual(v, gitRemoteSettingsProp)) {
		obj["gitRemoteSettings"] = gitRemoteSettingsProp
	}
	workspaceCompilationOverridesProp, err := expandDataformRepositoryWorkspaceCompilationOverrides(d.Get("workspace_compilation_overrides"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("workspace_compilation_overrides"); !tpgresource.IsEmptyValue(reflect.ValueOf(workspaceCompilationOverridesProp)) && (ok || !reflect.DeepEqual(v, workspaceCompilationOverridesProp)) {
		obj["workspaceCompilationOverrides"] = workspaceCompilationOverridesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataformBasePath}}projects/{{project}}/locations/{{region}}/repositories?repositoryId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Repository: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Repository: %s", err)
	}
	billingProject = project

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
		return fmt.Errorf("Error creating Repository: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/repositories/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Repository %q: %#v", d.Id(), res)

	return resourceDataformRepositoryRead(d, meta)
}

func resourceDataformRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataformBasePath}}projects/{{project}}/locations/{{region}}/repositories/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Repository: %s", err)
	}
	billingProject = project

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DataformRepository %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}

	if err := d.Set("name", flattenDataformRepositoryName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("git_remote_settings", flattenDataformRepositoryGitRemoteSettings(res["gitRemoteSettings"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("workspace_compilation_overrides", flattenDataformRepositoryWorkspaceCompilationOverrides(res["workspaceCompilationOverrides"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}

	return nil
}

func resourceDataformRepositoryUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Repository: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	gitRemoteSettingsProp, err := expandDataformRepositoryGitRemoteSettings(d.Get("git_remote_settings"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("git_remote_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, gitRemoteSettingsProp)) {
		obj["gitRemoteSettings"] = gitRemoteSettingsProp
	}
	workspaceCompilationOverridesProp, err := expandDataformRepositoryWorkspaceCompilationOverrides(d.Get("workspace_compilation_overrides"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("workspace_compilation_overrides"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, workspaceCompilationOverridesProp)) {
		obj["workspaceCompilationOverrides"] = workspaceCompilationOverridesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataformBasePath}}projects/{{project}}/locations/{{region}}/repositories/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Repository %q: %#v", d.Id(), obj)

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
		return fmt.Errorf("Error updating Repository %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Repository %q: %#v", d.Id(), res)
	}

	return resourceDataformRepositoryRead(d, meta)
}

func resourceDataformRepositoryDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Repository: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DataformBasePath}}projects/{{project}}/locations/{{region}}/repositories/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Repository %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "Repository")
	}

	log.Printf("[DEBUG] Finished deleting Repository %q: %#v", d.Id(), res)
	return nil
}

func resourceDataformRepositoryImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/repositories/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/repositories/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDataformRepositoryName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenDataformRepositoryGitRemoteSettings(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["url"] =
		flattenDataformRepositoryGitRemoteSettingsUrl(original["url"], d, config)
	transformed["default_branch"] =
		flattenDataformRepositoryGitRemoteSettingsDefaultBranch(original["defaultBranch"], d, config)
	transformed["authentication_token_secret_version"] =
		flattenDataformRepositoryGitRemoteSettingsAuthenticationTokenSecretVersion(original["authenticationTokenSecretVersion"], d, config)
	transformed["token_status"] =
		flattenDataformRepositoryGitRemoteSettingsTokenStatus(original["tokenStatus"], d, config)
	return []interface{}{transformed}
}
func flattenDataformRepositoryGitRemoteSettingsUrl(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataformRepositoryGitRemoteSettingsDefaultBranch(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataformRepositoryGitRemoteSettingsAuthenticationTokenSecretVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataformRepositoryGitRemoteSettingsTokenStatus(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataformRepositoryWorkspaceCompilationOverrides(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["default_database"] =
		flattenDataformRepositoryWorkspaceCompilationOverridesDefaultDatabase(original["defaultDatabase"], d, config)
	transformed["schema_suffix"] =
		flattenDataformRepositoryWorkspaceCompilationOverridesSchemaSuffix(original["schemaSuffix"], d, config)
	transformed["table_prefix"] =
		flattenDataformRepositoryWorkspaceCompilationOverridesTablePrefix(original["tablePrefix"], d, config)
	return []interface{}{transformed}
}
func flattenDataformRepositoryWorkspaceCompilationOverridesDefaultDatabase(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataformRepositoryWorkspaceCompilationOverridesSchemaSuffix(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataformRepositoryWorkspaceCompilationOverridesTablePrefix(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDataformRepositoryName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryGitRemoteSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUrl, err := expandDataformRepositoryGitRemoteSettingsUrl(original["url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUrl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["url"] = transformedUrl
	}

	transformedDefaultBranch, err := expandDataformRepositoryGitRemoteSettingsDefaultBranch(original["default_branch"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDefaultBranch); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["defaultBranch"] = transformedDefaultBranch
	}

	transformedAuthenticationTokenSecretVersion, err := expandDataformRepositoryGitRemoteSettingsAuthenticationTokenSecretVersion(original["authentication_token_secret_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAuthenticationTokenSecretVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["authenticationTokenSecretVersion"] = transformedAuthenticationTokenSecretVersion
	}

	transformedTokenStatus, err := expandDataformRepositoryGitRemoteSettingsTokenStatus(original["token_status"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTokenStatus); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tokenStatus"] = transformedTokenStatus
	}

	return transformed, nil
}

func expandDataformRepositoryGitRemoteSettingsUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryGitRemoteSettingsDefaultBranch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryGitRemoteSettingsAuthenticationTokenSecretVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryGitRemoteSettingsTokenStatus(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkspaceCompilationOverrides(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDefaultDatabase, err := expandDataformRepositoryWorkspaceCompilationOverridesDefaultDatabase(original["default_database"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDefaultDatabase); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["defaultDatabase"] = transformedDefaultDatabase
	}

	transformedSchemaSuffix, err := expandDataformRepositoryWorkspaceCompilationOverridesSchemaSuffix(original["schema_suffix"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSchemaSuffix); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["schemaSuffix"] = transformedSchemaSuffix
	}

	transformedTablePrefix, err := expandDataformRepositoryWorkspaceCompilationOverridesTablePrefix(original["table_prefix"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTablePrefix); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tablePrefix"] = transformedTablePrefix
	}

	return transformed, nil
}

func expandDataformRepositoryWorkspaceCompilationOverridesDefaultDatabase(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkspaceCompilationOverridesSchemaSuffix(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkspaceCompilationOverridesTablePrefix(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
