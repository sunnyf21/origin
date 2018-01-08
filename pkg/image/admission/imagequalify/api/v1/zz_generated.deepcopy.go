// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageQualifyConfig) DeepCopyInto(out *ImageQualifyConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]ImageQualifyRule, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageQualifyConfig.
func (in *ImageQualifyConfig) DeepCopy() *ImageQualifyConfig {
	if in == nil {
		return nil
	}
	out := new(ImageQualifyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageQualifyConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageQualifyRule) DeepCopyInto(out *ImageQualifyRule) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageQualifyRule.
func (in *ImageQualifyRule) DeepCopy() *ImageQualifyRule {
	if in == nil {
		return nil
	}
	out := new(ImageQualifyRule)
	in.DeepCopyInto(out)
	return out
}
