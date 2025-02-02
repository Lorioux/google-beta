// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package vmwareengine

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/lorioux/google-beta/google-beta/tpgresource"
	transport_tpg "github.com/lorioux/google-beta/google-beta/transport"
)

func DataSourceVmwareengineCluster() *schema.Resource {

	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceVmwareengineCluster().Schema)
	tpgresource.AddRequiredFieldsToSchema(dsSchema, "parent", "name")
	return &schema.Resource{
		Read:   dataSourceVmwareengineClusterRead,
		Schema: dsSchema,
	}
}

func dataSourceVmwareengineClusterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/clusters/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	return resourceVmwareengineClusterRead(d, meta)
}
