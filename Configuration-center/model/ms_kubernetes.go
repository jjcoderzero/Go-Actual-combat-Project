package model

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type ObjectMeta struct {
	Name              string    `json:",omitempty"`
	NameSpace         string    `json:"namespace,omitempty"`
	CreationTimestamp v1.Time   `json:"creationTimestamp,omitempty"`
	UID               types.UID `json:"uid,omitempty"`
}

type TypeMeta struct {
	TotalItems int `json:"totalItems"`
}

type ResourceKind string

const (
	ResourceKindConfigMap                = "configmap"
	ResourceKindDaemonSet                = "daemonset"
	ResourceKindDeployment               = "deployment"
	ResourceKindEvent                    = "event"
	ResourceKindHorizontalPodAutoscaler  = "horizontalpodautoscaler"
	ResourceKindIngress                  = "ingress"
	ResourceKindServiceAccount           = "serviceaccount"
	ResourceKindJob                      = "job"
	ResourceKindCronJob                  = "cronjob"
	ResourceKindLimitRange               = "limitrange"
	ResourceKindNamespace                = "namespace"
	ResourceKindNode                     = "node"
	ResourceKindPersistentVolumeClaim    = "persistentvolumeclaim"
	ResourceKindPersistentVolume         = "persistentvolume"
	ResourceKindCustomResourceDefinition = "customresourcedefinition"
	ResourceKindPod                      = "pod"
	ResourceKindReplicaSet               = "replicaset"
	ResourceKindReplicationController    = "replicationcontroller"
	ResourceKindResourceQuota            = "resourcequota"
	ResourceKindSecret                   = "secret"
	ResourceKindService                  = "service"
	ResourceKindStatefulSet              = "statefulset"
	ResourceKindStorageClass             = "storageclass"
	ResourceKindClusterRole              = "clusterrole"
	ResourceKindClusterRoleBinding       = "clusterrolebinding"
	ResourceKindRole                     = "role"
	ResourceKindRoleBinding              = "rolebinding"
	ResourceKindPlugin                   = "plugin"
	ResourceKindEndpoint                 = "endpoint"
	ResourceKindNetworkPolicy            = "networkpolicy"
)
