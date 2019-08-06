// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package network

import (
	calicov1alpha1 "github.com/gardener/gardener-extensions/controllers/networking-calico/pkg/apis/calico/v1alpha1"
	"github.com/gardener/gardener-extensions/controllers/networking-calico/pkg/controller"
	extensionswebhook "github.com/gardener/gardener-extensions/pkg/webhook"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
)

func mutateNetworkConfig(network *extensionsv1alpha1.Network) error {
	extensionswebhook.LogMutation(logger, "Network", network.Namespace, network.Name)
	networkConfig, err := controller.CalicoNetworkConfigFromNetworkResource(network)
	if err != nil {
		return err
	}
	networkConfig.Backend = calicov1alpha1.None
	network.Spec.ProviderConfig = &runtime.RawExtension{
		Object: networkConfig,
	}

	return nil
}
