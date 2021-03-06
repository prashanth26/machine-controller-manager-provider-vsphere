/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

This file was copied and modified from the kubernetes-csi/drivers project
https://github.com/kubernetes-csi/drivers/blob/release-1.0/app/nfsplugin/main.go

Modifications Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved.
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gardener/machine-controller-manager-provider-vsphere/pkg/vsphere"
	"github.com/spf13/cobra"
)

var (
	version  string
	endpoint string
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Set("v", "3")
}

func main() {

	err := flag.CommandLine.Parse([]string{})
	if err != nil {
		panic(err)
	}

	cmd := &cobra.Command{
		Use:   "cmi-plugin",
		Short: "vSphere gRPC CMI Plugin for machine-controller-manager",
		Run: func(cmd *cobra.Command, args []string) {
			handle()
		},
	}

	cmd.PersistentFlags().StringVar(&endpoint, "endpoint", "", "Endpoint to be used for plugin")
	err = cmd.MarkPersistentFlagRequired("endpoint")
	if err != nil {
		panic(err)
	}

	err = cmd.ParseFlags(os.Args[1:])
	if err != nil {
		panic(err)
	}
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

func handle() {
	d := vsphere.NewPlugin(endpoint, version)
	d.Run()
}
