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
*/

// This file was automatically generated by informer-gen

package internalversion

import (
	"fmt"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
	api "k8s.io/kubernetes/pkg/api"
	admissionregistration "k8s.io/kubernetes/pkg/apis/admissionregistration"
	apps "k8s.io/kubernetes/pkg/apis/apps"
	autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
	batch "k8s.io/kubernetes/pkg/apis/batch"
	certificates "k8s.io/kubernetes/pkg/apis/certificates"
	extensions "k8s.io/kubernetes/pkg/apis/extensions"
	policy "k8s.io/kubernetes/pkg/apis/policy"
	rbac "k8s.io/kubernetes/pkg/apis/rbac"
	settings "k8s.io/kubernetes/pkg/apis/settings"
	storage "k8s.io/kubernetes/pkg/apis/storage"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=Admissionregistration, Version=InternalVersion
	case admissionregistration.SchemeGroupVersion.WithResource("externaladmissionhookconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Admissionregistration().InternalVersion().ExternalAdmissionHookConfigurations().Informer()}, nil
	case admissionregistration.SchemeGroupVersion.WithResource("initializerconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Admissionregistration().InternalVersion().InitializerConfigurations().Informer()}, nil

		// Group=Apps, Version=InternalVersion
	case apps.SchemeGroupVersion.WithResource("controllerrevisions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().InternalVersion().ControllerRevisions().Informer()}, nil
	case apps.SchemeGroupVersion.WithResource("statefulsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().InternalVersion().StatefulSets().Informer()}, nil

		// Group=Autoscaling, Version=InternalVersion
	case autoscaling.SchemeGroupVersion.WithResource("horizontalpodautoscalers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Autoscaling().InternalVersion().HorizontalPodAutoscalers().Informer()}, nil

		// Group=Batch, Version=InternalVersion
	case batch.SchemeGroupVersion.WithResource("cronjobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Batch().InternalVersion().CronJobs().Informer()}, nil
	case batch.SchemeGroupVersion.WithResource("jobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Batch().InternalVersion().Jobs().Informer()}, nil

		// Group=Certificates, Version=InternalVersion
	case certificates.SchemeGroupVersion.WithResource("certificatesigningrequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Certificates().InternalVersion().CertificateSigningRequests().Informer()}, nil

		// Group=Core, Version=InternalVersion
	case api.SchemeGroupVersion.WithResource("componentstatuses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().ComponentStatuses().Informer()}, nil
	case api.SchemeGroupVersion.WithResource("configmaps"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().ConfigMaps().Informer()}, nil
	case api.SchemeGroupVersion.WithResource("endpoints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().Endpoints().Informer()}, nil
	case api.SchemeGroupVersion.WithResource("events"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().Events().Informer()}, nil
	case api.SchemeGroupVersion.WithResource("limitranges"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().LimitRanges().Informer()}, nil
	case api.SchemeGroupVersion.WithResource("namespaces"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().Namespaces().Informer()}, nil
	case api.SchemeGroupVersion.WithResource("nodes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().Nodes().Informer()}, nil
	case api.SchemeGroupVersion.WithResource("persistentvolumes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().PersistentVolumes().Informer()}, nil
	case api.SchemeGroupVersion.WithResource("persistentvolumeclaims"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().PersistentVolumeClaims().Informer()}, nil
	case api.SchemeGroupVersion.WithResource("pods"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().Pods().Informer()}, nil
	case api.SchemeGroupVersion.WithResource("podtemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().PodTemplates().Informer()}, nil
	case api.SchemeGroupVersion.WithResource("replicationcontrollers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().ReplicationControllers().Informer()}, nil
	case api.SchemeGroupVersion.WithResource("resourcequotas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().ResourceQuotas().Informer()}, nil
	case api.SchemeGroupVersion.WithResource("secrets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().Secrets().Informer()}, nil
	case api.SchemeGroupVersion.WithResource("services"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().Services().Informer()}, nil
	case api.SchemeGroupVersion.WithResource("serviceaccounts"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().ServiceAccounts().Informer()}, nil

		// Group=Extensions, Version=InternalVersion
	case extensions.SchemeGroupVersion.WithResource("daemonsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().InternalVersion().DaemonSets().Informer()}, nil
	case extensions.SchemeGroupVersion.WithResource("deployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().InternalVersion().Deployments().Informer()}, nil
	case extensions.SchemeGroupVersion.WithResource("ingresses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().InternalVersion().Ingresses().Informer()}, nil
	case extensions.SchemeGroupVersion.WithResource("networkpolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().InternalVersion().NetworkPolicies().Informer()}, nil
	case extensions.SchemeGroupVersion.WithResource("podsecuritypolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().InternalVersion().PodSecurityPolicies().Informer()}, nil
	case extensions.SchemeGroupVersion.WithResource("replicasets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().InternalVersion().ReplicaSets().Informer()}, nil
	case extensions.SchemeGroupVersion.WithResource("thirdpartyresources"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().InternalVersion().ThirdPartyResources().Informer()}, nil

		// Group=Policy, Version=InternalVersion
	case policy.SchemeGroupVersion.WithResource("poddisruptionbudgets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Policy().InternalVersion().PodDisruptionBudgets().Informer()}, nil

		// Group=Rbac, Version=InternalVersion
	case rbac.SchemeGroupVersion.WithResource("clusterroles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().InternalVersion().ClusterRoles().Informer()}, nil
	case rbac.SchemeGroupVersion.WithResource("clusterrolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().InternalVersion().ClusterRoleBindings().Informer()}, nil
	case rbac.SchemeGroupVersion.WithResource("roles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().InternalVersion().Roles().Informer()}, nil
	case rbac.SchemeGroupVersion.WithResource("rolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().InternalVersion().RoleBindings().Informer()}, nil

		// Group=Settings, Version=InternalVersion
	case settings.SchemeGroupVersion.WithResource("podpresets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Settings().InternalVersion().PodPresets().Informer()}, nil

		// Group=Storage, Version=InternalVersion
	case storage.SchemeGroupVersion.WithResource("storageclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().InternalVersion().StorageClasses().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
