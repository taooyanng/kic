//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 Kong, Inc.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	"github.com/kong/go-kong/kong"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigSource) DeepCopyInto(out *ConfigSource) {
	*out = *in
	out.SecretValue = in.SecretValue
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigSource.
func (in *ConfigSource) DeepCopy() *ConfigSource {
	if in == nil {
		return nil
	}
	out := new(ConfigSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongClusterPlugin) DeepCopyInto(out *KongClusterPlugin) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Config.DeepCopyInto(&out.Config)
	if in.ConfigFrom != nil {
		in, out := &in.ConfigFrom, &out.ConfigFrom
		*out = new(NamespacedConfigSource)
		**out = **in
	}
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]KongProtocol, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongClusterPlugin.
func (in *KongClusterPlugin) DeepCopy() *KongClusterPlugin {
	if in == nil {
		return nil
	}
	out := new(KongClusterPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongClusterPlugin) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongClusterPluginList) DeepCopyInto(out *KongClusterPluginList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KongClusterPlugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongClusterPluginList.
func (in *KongClusterPluginList) DeepCopy() *KongClusterPluginList {
	if in == nil {
		return nil
	}
	out := new(KongClusterPluginList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongClusterPluginList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongConsumer) DeepCopyInto(out *KongConsumer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongConsumer.
func (in *KongConsumer) DeepCopy() *KongConsumer {
	if in == nil {
		return nil
	}
	out := new(KongConsumer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongConsumer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongConsumerList) DeepCopyInto(out *KongConsumerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KongConsumer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongConsumerList.
func (in *KongConsumerList) DeepCopy() *KongConsumerList {
	if in == nil {
		return nil
	}
	out := new(KongConsumerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongConsumerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongIngress) DeepCopyInto(out *KongIngress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Upstream != nil {
		in, out := &in.Upstream, &out.Upstream
		*out = new(KongIngressUpstream)
		(*in).DeepCopyInto(*out)
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(KongIngressService)
		(*in).DeepCopyInto(*out)
	}
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = new(KongIngressRoute)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongIngress.
func (in *KongIngress) DeepCopy() *KongIngress {
	if in == nil {
		return nil
	}
	out := new(KongIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongIngress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongIngressList) DeepCopyInto(out *KongIngressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KongIngress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongIngressList.
func (in *KongIngressList) DeepCopy() *KongIngressList {
	if in == nil {
		return nil
	}
	out := new(KongIngressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongIngressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongIngressRoute) DeepCopyInto(out *KongIngressRoute) {
	*out = *in
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]*KongProtocol, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(KongProtocol)
				**out = **in
			}
		}
	}
	if in.RegexPriority != nil {
		in, out := &in.RegexPriority, &out.RegexPriority
		*out = new(int)
		**out = **in
	}
	if in.StripPath != nil {
		in, out := &in.StripPath, &out.StripPath
		*out = new(bool)
		**out = **in
	}
	if in.PreserveHost != nil {
		in, out := &in.PreserveHost, &out.PreserveHost
		*out = new(bool)
		**out = **in
	}
	if in.HTTPSRedirectStatusCode != nil {
		in, out := &in.HTTPSRedirectStatusCode, &out.HTTPSRedirectStatusCode
		*out = new(int)
		**out = **in
	}
	if in.PathHandling != nil {
		in, out := &in.PathHandling, &out.PathHandling
		*out = new(string)
		**out = **in
	}
	if in.SNIs != nil {
		in, out := &in.SNIs, &out.SNIs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RequestBuffering != nil {
		in, out := &in.RequestBuffering, &out.RequestBuffering
		*out = new(bool)
		**out = **in
	}
	if in.ResponseBuffering != nil {
		in, out := &in.ResponseBuffering, &out.ResponseBuffering
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongIngressRoute.
func (in *KongIngressRoute) DeepCopy() *KongIngressRoute {
	if in == nil {
		return nil
	}
	out := new(KongIngressRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongIngressService) DeepCopyInto(out *KongIngressService) {
	*out = *in
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Retries != nil {
		in, out := &in.Retries, &out.Retries
		*out = new(int)
		**out = **in
	}
	if in.ConnectTimeout != nil {
		in, out := &in.ConnectTimeout, &out.ConnectTimeout
		*out = new(int)
		**out = **in
	}
	if in.ReadTimeout != nil {
		in, out := &in.ReadTimeout, &out.ReadTimeout
		*out = new(int)
		**out = **in
	}
	if in.WriteTimeout != nil {
		in, out := &in.WriteTimeout, &out.WriteTimeout
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongIngressService.
func (in *KongIngressService) DeepCopy() *KongIngressService {
	if in == nil {
		return nil
	}
	out := new(KongIngressService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongIngressUpstream) DeepCopyInto(out *KongIngressUpstream) {
	*out = *in
	if in.HostHeader != nil {
		in, out := &in.HostHeader, &out.HostHeader
		*out = new(string)
		**out = **in
	}
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
		*out = new(string)
		**out = **in
	}
	if in.Slots != nil {
		in, out := &in.Slots, &out.Slots
		*out = new(int)
		**out = **in
	}
	if in.Healthchecks != nil {
		in, out := &in.Healthchecks, &out.Healthchecks
		*out = new(kong.Healthcheck)
		(*in).DeepCopyInto(*out)
	}
	if in.HashOn != nil {
		in, out := &in.HashOn, &out.HashOn
		*out = new(string)
		**out = **in
	}
	if in.HashFallback != nil {
		in, out := &in.HashFallback, &out.HashFallback
		*out = new(string)
		**out = **in
	}
	if in.HashOnHeader != nil {
		in, out := &in.HashOnHeader, &out.HashOnHeader
		*out = new(string)
		**out = **in
	}
	if in.HashFallbackHeader != nil {
		in, out := &in.HashFallbackHeader, &out.HashFallbackHeader
		*out = new(string)
		**out = **in
	}
	if in.HashOnCookie != nil {
		in, out := &in.HashOnCookie, &out.HashOnCookie
		*out = new(string)
		**out = **in
	}
	if in.HashOnCookiePath != nil {
		in, out := &in.HashOnCookiePath, &out.HashOnCookiePath
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongIngressUpstream.
func (in *KongIngressUpstream) DeepCopy() *KongIngressUpstream {
	if in == nil {
		return nil
	}
	out := new(KongIngressUpstream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongPlugin) DeepCopyInto(out *KongPlugin) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Config.DeepCopyInto(&out.Config)
	if in.ConfigFrom != nil {
		in, out := &in.ConfigFrom, &out.ConfigFrom
		*out = new(ConfigSource)
		**out = **in
	}
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]KongProtocol, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongPlugin.
func (in *KongPlugin) DeepCopy() *KongPlugin {
	if in == nil {
		return nil
	}
	out := new(KongPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongPlugin) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongPluginList) DeepCopyInto(out *KongPluginList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KongPlugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongPluginList.
func (in *KongPluginList) DeepCopy() *KongPluginList {
	if in == nil {
		return nil
	}
	out := new(KongPluginList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongPluginList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedConfigSource) DeepCopyInto(out *NamespacedConfigSource) {
	*out = *in
	out.SecretValue = in.SecretValue
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedConfigSource.
func (in *NamespacedConfigSource) DeepCopy() *NamespacedConfigSource {
	if in == nil {
		return nil
	}
	out := new(NamespacedConfigSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedSecretValueFromSource) DeepCopyInto(out *NamespacedSecretValueFromSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedSecretValueFromSource.
func (in *NamespacedSecretValueFromSource) DeepCopy() *NamespacedSecretValueFromSource {
	if in == nil {
		return nil
	}
	out := new(NamespacedSecretValueFromSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretValueFromSource) DeepCopyInto(out *SecretValueFromSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretValueFromSource.
func (in *SecretValueFromSource) DeepCopy() *SecretValueFromSource {
	if in == nil {
		return nil
	}
	out := new(SecretValueFromSource)
	in.DeepCopyInto(out)
	return out
}