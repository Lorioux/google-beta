// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server"
	"github.com/hashicorp/terraform-plugin-mux/tf5muxserver"

	ver "github.com/lorioux/google-beta/version"
	provider "github.com/lorioux/google-beta/google-beta"
	"github.com/lorioux/google-beta/google-beta/fwprovider"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = ver.ProviderVersion

	// goreleaser can also pass the specific commit if you want
	// commit  string = ""
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	// concat with sdkv2 provider
	providers := []func() tfprotov5.ProviderServer{
		providerserver.NewProtocol5(fwprovider.New(version)), // framework provider
		provider.Provider().GRPCProvider,                     // sdk provider
	}

	// use the muxer
	muxServer, err := tf5muxserver.NewMuxServer(context.Background(), providers...)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var serveOpts []tf5server.ServeOpt

	if debug {
		serveOpts = append(serveOpts, tf5server.WithManagedDebug())
	}

	err = tf5server.Serve(
		"registry.terraform.io/hashicorp/google-beta",
		muxServer.ProviderServer,
		serveOpts...,
	)

	if err != nil {
		log.Fatal(err)
	}
}
