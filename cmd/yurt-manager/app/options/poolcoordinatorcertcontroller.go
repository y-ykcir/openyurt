/*
Copyright 2023 The OpenYurt Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package options

import (
	"github.com/spf13/pflag"

	"github.com/openyurtio/openyurt/pkg/controller/poolcoordinator/config"
)

type PoolCoordinatorCertControllerOptions struct {
	*config.PoolCoordinatorCertControllerConfiguration
}

func NewPoolCoordinatorCertControllerOptions() *PoolCoordinatorCertControllerOptions {
	return &PoolCoordinatorCertControllerOptions{
		&config.PoolCoordinatorCertControllerConfiguration{
			PoolCoordinatorNameSpace: "kube-system",
		},
	}
}

// AddFlags adds flags related to nodepool for yurt-manager to the specified FlagSet.
func (p *PoolCoordinatorCertControllerOptions) AddFlags(fs *pflag.FlagSet) {
	if p == nil {
		return
	}

	fs.StringVar(&p.PoolCoordinatorNameSpace, "pool-coordinator-namespace", p.PoolCoordinatorNameSpace, "Specify the namespace of pool coordinator.")
}

// ApplyTo fills up nodepool config with options.
func (p *PoolCoordinatorCertControllerOptions) ApplyTo(cfg *config.PoolCoordinatorCertControllerConfiguration) error {
	if p == nil {
		return nil
	}
	cfg.PoolCoordinatorNameSpace = p.PoolCoordinatorNameSpace

	return nil
}

// Validate checks validation of PoolCoordinatorCertControllerOptions.
func (p *PoolCoordinatorCertControllerOptions) Validate() []error {
	if p == nil {
		return nil
	}
	errs := []error{}
	return errs
}
