// +build !ignore_autogenerated

/*
Copyright 2019 The Kubernetes Authors.

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
// Code generated by main. DO NOT EDIT.

package v1alpha3

import (
	"os"

	common "github.com/kubeflow/katib/pkg/apis/controller/common/v1alpha3"
)

func (e *Experiment) SetDefault() {
	e.setDefaultParallelTrialCount()
	e.setDefaultTrialTemplate()
	e.setDefaultMetricsCollector()
}

func (e *Experiment) setDefaultParallelTrialCount() {
	if e.Spec.ParallelTrialCount == nil {
		e.Spec.ParallelTrialCount = new(int32)
		*e.Spec.ParallelTrialCount = DefaultTrialParallelCount
	}
}

func (e *Experiment) setDefaultTrialTemplate() {
	t := e.Spec.TrialTemplate
	if t == nil {
		t = &TrialTemplate{
			Retain: true,
		}
	}
	if t.GoTemplate == nil {
		t.GoTemplate = &GoTemplate{}
	}
	if t.GoTemplate.RawTemplate == "" && t.GoTemplate.TemplateSpec == nil {
		t.GoTemplate.TemplateSpec = &TemplateSpec{
			ConfigMapNamespace: os.Getenv(DefaultKatibNamespaceEnvName),
			ConfigMapName:      DefaultTrialConfigMapName,
			TemplatePath:       DefaultTrialTemplatePath,
		}
	}
	e.Spec.TrialTemplate = t
}

func (e *Experiment) setDefaultMetricsCollector() {
	if e.Spec.MetricsCollectorSpec == nil {
		e.Spec.MetricsCollectorSpec = &MetricsCollectorSpec{}
	}
	if e.Spec.MetricsCollectorSpec.Collector == nil {
		e.Spec.MetricsCollectorSpec.Collector = &common.CollectorSpec{
			Kind: common.StdOutCollector,
		}
	}
}