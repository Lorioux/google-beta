// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package composer

import (
	"fmt"
	"time"

	"github.com/lorioux/google-beta/google-beta/tpgresource"
	transport_tpg "github.com/lorioux/google-beta/google-beta/transport"

	composer "google.golang.org/api/composer/v1beta1"
)

type ComposerOperationWaiter struct {
	Service *composer.ProjectsLocationsService
	tpgresource.CommonOperationWaiter
}

func (w *ComposerOperationWaiter) QueryOp() (interface{}, error) {
	if w == nil {
		return nil, fmt.Errorf("Cannot query operation, it's unset or nil.")
	}
	return w.Service.Operations.Get(w.Op.Name).Do()
}

func ComposerOperationWaitTime(config *transport_tpg.Config, op *composer.Operation, project, activity, userAgent string, timeout time.Duration) error {
	w := &ComposerOperationWaiter{
		Service: config.NewComposerClient(userAgent).Projects.Locations,
	}
	if err := w.SetOp(op); err != nil {
		return err
	}
	return tpgresource.OperationWait(w, activity, timeout, config.PollInterval)
}
