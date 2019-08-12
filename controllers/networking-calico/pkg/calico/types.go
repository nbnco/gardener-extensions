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

package calico

import "path/filepath"

const (
	Name = "networking-calico"

	// ImageNames
	CNIImageName             = "calico-cni"
	NodeImageName            = "calico-node"
	TyphaImageName           = "calico-typha"
	KubeControllersImageName = "calico-kube-controllers"

	// CalicoRelease is the name of the Calico Release
	ReleaseName = "calico"
)

var (
	// ChartsPath is the path to the charts
	ChartsPath = filepath.Join("controllers", Name, "charts")
	// InternalChartsPath is the path to the internal charts
	InternalChartsPath = filepath.Join(ChartsPath, "internal")

	// CalicoChartPath path for internal Calico Chart
	ChartPath = filepath.Join(InternalChartsPath, "calico")
)
