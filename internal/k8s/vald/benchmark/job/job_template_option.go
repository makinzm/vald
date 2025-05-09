//
// Copyright (C) 2019-2025 vdaas.org vald team <vald@vdaas.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package job

import (
	"github.com/vdaas/vald/internal/k8s"
	corev1 "k8s.io/api/core/v1"
)

type BenchmarkJobTplOption func(b *benchmarkJobTpl) error

var defaultBenchmarkJobTplOpts = []BenchmarkJobTplOption{
	WithContainerName("vald-benchmark-job"),
	WithContainerImage("vdaas/vald-benchmark-job"),
	WithImagePullPolicy(PullAlways),
}

// WithContainerName sets the docker container name of benchmark job.
func WithContainerName(name string) BenchmarkJobTplOption {
	return func(b *benchmarkJobTpl) error {
		if len(name) > 0 {
			b.containerName = name
		}
		return nil
	}
}

// WithContainerImage sets the docker image path for benchmark job.
func WithContainerImage(name string) BenchmarkJobTplOption {
	return func(b *benchmarkJobTpl) error {
		if len(name) > 0 {
			b.containerImageName = name
		}
		return nil
	}
}

// WithImagePullPolicy sets the docker image pull policy for benchmark job.
func WithImagePullPolicy(p ImagePullPolicy) BenchmarkJobTplOption {
	return func(b *benchmarkJobTpl) error {
		if len(p) > 0 {
			b.imagePullPolicy = p
		}
		return nil
	}
}

// WithOperatorConfigMap sets the configMapName for mounting Job Pod.
func WithOperatorConfigMap(cm string) BenchmarkJobTplOption {
	return func(b *benchmarkJobTpl) error {
		if len(cm) > 0 {
			b.configMapName = cm
		}
		return nil
	}
}

// BenchmarkJobOption represents the option for create benchmark job template.
type BenchmarkJobOption func(b *k8s.Job) error

const (
	// defaultTTLSeconds represents the default TTLSecondsAfterFinished for benchmark job template.
	defaultTTLSeconds int32 = 600
	// defaultCompletions represents the default completions for benchmark job template.
	defaultCompletions int32 = 1
	// defaultParallelism represents the default parallelism for benchmark job template.
	defaultParallelism int32 = 1
)

var defaultBenchmarkJobOpts = []BenchmarkJobOption{
	WithSvcAccountName(svcAccount),
	WithRestartPolicy(RestartPolicyNever),
	WithTTLSecondsAfterFinished(defaultTTLSeconds),
	WithCompletions(defaultCompletions),
	WithParallelism(defaultParallelism),
}

// WithSvcAccountName sets the service account name for benchmark job.
func WithSvcAccountName(name string) BenchmarkJobOption {
	return func(b *k8s.Job) error {
		if len(name) > 0 {
			b.Spec.Template.Spec.ServiceAccountName = name
		}
		return nil
	}
}

// WithRestartPolicy sets the job restart policy for benchmark job.
func WithRestartPolicy(rp RestartPolicy) BenchmarkJobOption {
	return func(b *k8s.Job) error {
		if len(rp) > 0 {
			b.Spec.Template.Spec.RestartPolicy = corev1.RestartPolicy(rp)
		}
		return nil
	}
}

// WithBackoffLimit sets the job backoff limit for benchmark job.
func WithBackoffLimit(bo int32) BenchmarkJobOption {
	return func(b *k8s.Job) error {
		b.Spec.BackoffLimit = &bo
		return nil
	}
}

// WithName sets the job name of benchmark job.
func WithName(name string) BenchmarkJobOption {
	return func(b *k8s.Job) error {
		b.Name = name
		return nil
	}
}

// WithNamespace specify namespace where job will execute.
func WithNamespace(ns string) BenchmarkJobOption {
	return func(b *k8s.Job) error {
		b.Namespace = ns
		return nil
	}
}

// WithOwnerRef sets the OwnerReference to the job resource.
func WithOwnerRef(refs []k8s.OwnerReference) BenchmarkJobOption {
	return func(b *k8s.Job) error {
		if len(refs) > 0 {
			b.OwnerReferences = refs
		}
		return nil
	}
}

// WithCompletions sets the job completion.
func WithCompletions(com int32) BenchmarkJobOption {
	return func(b *k8s.Job) error {
		if com > 1 {
			b.Spec.Completions = &com
		}
		return nil
	}
}

// WithParallelism sets the job parallelism.
func WithParallelism(parallelism int32) BenchmarkJobOption {
	return func(b *k8s.Job) error {
		if parallelism > 1 {
			b.Spec.Parallelism = &parallelism
		}
		return nil
	}
}

// WithLabel sets the label to the job resource.
func WithLabel(label map[string]string) BenchmarkJobOption {
	return func(b *k8s.Job) error {
		if len(label) > 0 {
			b.Labels = label
		}
		return nil
	}
}

// WithTTLSecondsAfterFinished sets the TTLSecondsAfterFinished to the job template.
func WithTTLSecondsAfterFinished(ttl int32) BenchmarkJobOption {
	return func(b *k8s.Job) error {
		if ttl > 0 {
			b.Spec.TTLSecondsAfterFinished = &ttl
		}
		return nil
	}
}
