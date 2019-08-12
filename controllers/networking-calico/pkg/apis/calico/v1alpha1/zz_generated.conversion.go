// +build !ignore_autogenerated

/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	calico "github.com/gardener/gardener-extensions/controllers/networking-calico/pkg/apis/calico"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*IPAM)(nil), (*calico.IPAM)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IPAM_To_calico_IPAM(a.(*IPAM), b.(*calico.IPAM), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*calico.IPAM)(nil), (*IPAM)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_calico_IPAM_To_v1alpha1_IPAM(a.(*calico.IPAM), b.(*IPAM), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NetworkConfig)(nil), (*calico.NetworkConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_NetworkConfig_To_calico_NetworkConfig(a.(*NetworkConfig), b.(*calico.NetworkConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*calico.NetworkConfig)(nil), (*NetworkConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_calico_NetworkConfig_To_v1alpha1_NetworkConfig(a.(*calico.NetworkConfig), b.(*NetworkConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NetworkStatus)(nil), (*calico.NetworkStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_NetworkStatus_To_calico_NetworkStatus(a.(*NetworkStatus), b.(*calico.NetworkStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*calico.NetworkStatus)(nil), (*NetworkStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_calico_NetworkStatus_To_v1alpha1_NetworkStatus(a.(*calico.NetworkStatus), b.(*NetworkStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_IPAM_To_calico_IPAM(in *IPAM, out *calico.IPAM, s conversion.Scope) error {
	out.Type = in.Type
	out.CIDR = (*calico.CIDR)(unsafe.Pointer(in.CIDR))
	return nil
}

// Convert_v1alpha1_IPAM_To_calico_IPAM is an autogenerated conversion function.
func Convert_v1alpha1_IPAM_To_calico_IPAM(in *IPAM, out *calico.IPAM, s conversion.Scope) error {
	return autoConvert_v1alpha1_IPAM_To_calico_IPAM(in, out, s)
}

func autoConvert_calico_IPAM_To_v1alpha1_IPAM(in *calico.IPAM, out *IPAM, s conversion.Scope) error {
	out.Type = in.Type
	out.CIDR = (*CIDR)(unsafe.Pointer(in.CIDR))
	return nil
}

// Convert_calico_IPAM_To_v1alpha1_IPAM is an autogenerated conversion function.
func Convert_calico_IPAM_To_v1alpha1_IPAM(in *calico.IPAM, out *IPAM, s conversion.Scope) error {
	return autoConvert_calico_IPAM_To_v1alpha1_IPAM(in, out, s)
}

func autoConvert_v1alpha1_NetworkConfig_To_calico_NetworkConfig(in *NetworkConfig, out *calico.NetworkConfig, s conversion.Scope) error {
	out.Backend = calico.Backend(in.Backend)
	out.IPAM = (*calico.IPAM)(unsafe.Pointer(in.IPAM))
	out.IPAutoDetectionMethod = (*string)(unsafe.Pointer(in.IPAutoDetectionMethod))
	return nil
}

// Convert_v1alpha1_NetworkConfig_To_calico_NetworkConfig is an autogenerated conversion function.
func Convert_v1alpha1_NetworkConfig_To_calico_NetworkConfig(in *NetworkConfig, out *calico.NetworkConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_NetworkConfig_To_calico_NetworkConfig(in, out, s)
}

func autoConvert_calico_NetworkConfig_To_v1alpha1_NetworkConfig(in *calico.NetworkConfig, out *NetworkConfig, s conversion.Scope) error {
	out.Backend = Backend(in.Backend)
	out.IPAM = (*IPAM)(unsafe.Pointer(in.IPAM))
	out.IPAutoDetectionMethod = (*string)(unsafe.Pointer(in.IPAutoDetectionMethod))
	return nil
}

// Convert_calico_NetworkConfig_To_v1alpha1_NetworkConfig is an autogenerated conversion function.
func Convert_calico_NetworkConfig_To_v1alpha1_NetworkConfig(in *calico.NetworkConfig, out *NetworkConfig, s conversion.Scope) error {
	return autoConvert_calico_NetworkConfig_To_v1alpha1_NetworkConfig(in, out, s)
}

func autoConvert_v1alpha1_NetworkStatus_To_calico_NetworkStatus(in *NetworkStatus, out *calico.NetworkStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_NetworkStatus_To_calico_NetworkStatus is an autogenerated conversion function.
func Convert_v1alpha1_NetworkStatus_To_calico_NetworkStatus(in *NetworkStatus, out *calico.NetworkStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_NetworkStatus_To_calico_NetworkStatus(in, out, s)
}

func autoConvert_calico_NetworkStatus_To_v1alpha1_NetworkStatus(in *calico.NetworkStatus, out *NetworkStatus, s conversion.Scope) error {
	return nil
}

// Convert_calico_NetworkStatus_To_v1alpha1_NetworkStatus is an autogenerated conversion function.
func Convert_calico_NetworkStatus_To_v1alpha1_NetworkStatus(in *calico.NetworkStatus, out *NetworkStatus, s conversion.Scope) error {
	return autoConvert_calico_NetworkStatus_To_v1alpha1_NetworkStatus(in, out, s)
}
