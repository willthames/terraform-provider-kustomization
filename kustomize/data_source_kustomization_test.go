package kustomize

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/kustomize/api/resid"
)

func TestDeterminePrefix(t *testing.T) {
	for _, id := range residListFirst {
		gvk := resid.GvkFromString(id)
		p := determinePrefix(gvk)
		e := uint32(1)
		assert.Equal(t, e, p, nil)
	}

	for _, id := range residListDefault {
		gvk := resid.GvkFromString(id)
		p := determinePrefix(gvk)
		e := uint32(5)
		assert.Equal(t, e, p, nil)
	}

	for _, id := range residListLast {
		gvk := resid.GvkFromString(id)
		p := determinePrefix(gvk)
		e := uint32(9)
		assert.Equal(t, e, p, nil)
	}
}

func TestPrefixHash(t *testing.T) {
	ti := uint32(math.MaxInt32 / 1000)
	i := prefixHash(uint32(1), ti)
	assert.Equal(t, 100021474, i, nil)

	i = prefixHash(uint32(5), ti)
	assert.Equal(t, 500021474, i, nil)

	i = prefixHash(uint32(9), ti)
	assert.Equal(t, 900021474, i, nil)
}

func TestPrefixHashMaxInt(t *testing.T) {
	ti := uint32(math.MaxInt32)
	i := prefixHash(uint32(1), ti)
	assert.Equal(t, 121474836, i, nil)

	i = prefixHash(uint32(5), ti)
	assert.Equal(t, 521474836, i, nil)

	i = prefixHash(uint32(9), ti)
	assert.Equal(t, 921474836, i, nil)
}

func TestIdSetHash(t *testing.T) {
	idList := []string{}
	idList = append(idList, residListFirst...)
	idList = append(idList, residListDefault...)
	idList = append(idList, residListLast...)
	setIDs := []int{}

	for _, s := range idList {
		rid := resid.FromString(s)

		f := idSetHash(rid.String())

		assert.NotContains(t, setIDs, f, nil)
		setIDs = append(setIDs, f)
	}
}

var residListFirst = []string{
	"~G_~V_Namespace|~X|test",
	"apiextensions.k8s.io_~V_CustomResourceDefinition|~X|test",
	"apiextensions.k8s.io_~V_CustomResourceDefinition|test-ns|test",
}

var residListDefault = []string{
	"autoscaling_~V_HorizontalPodAutoscaler|~X|test",
	"autoscaling_~V_HorizontalPodAutoscaler|test-ns|test",
	"coordination.k8s.io_~V_Lease|~X|test",
	"coordination.k8s.io_~V_Lease|test-ns|test",
	"discovery.k8s.io_~V_EndpointSlice|~X|test",
	"discovery.k8s.io_~V_EndpointSlice|test-ns|test",
	"extensions_~V_Ingress|~X|test",
	"extensions_~V_Ingress|test-ns|test",
	"rbac.authorization.k8s.io_~V_ClusterRoleBinding|~X|test",
	"rbac.authorization.k8s.io_~V_ClusterRoleBinding|test-ns|test",
	"rbac.authorization.k8s.io_~V_ClusterRole|~X|test",
	"rbac.authorization.k8s.io_~V_ClusterRole|test-ns|test",
	"rbac.authorization.k8s.io_~V_RoleBinding|~X|test",
	"rbac.authorization.k8s.io_~V_RoleBinding|test-ns|test",
	"rbac.authorization.k8s.io_~V_Role|~X|test",
	"rbac.authorization.k8s.io_~V_Role|test-ns|test",
	"apiregistration.k8s.io_~V_APIService|~X|test",
	"apiregistration.k8s.io_~V_APIService|test-ns|test",
	"batch_~V_CronJob|~X|test",
	"batch_~V_CronJob|test-ns|test",
	"batch_~V_Job|~X|test",
	"batch_~V_Job|test-ns|test",
	"authorization.k8s.io_~V_LocalSubjectAccessReview|~X|test",
	"authorization.k8s.io_~V_LocalSubjectAccessReview|test-ns|test",
	"authorization.k8s.io_~V_SelfSubjectAccessReview|~X|test",
	"authorization.k8s.io_~V_SelfSubjectAccessReview|test-ns|test",
	"authorization.k8s.io_~V_SelfSubjectRulesReview|~X|test",
	"authorization.k8s.io_~V_SelfSubjectRulesReview|test-ns|test",
	"authorization.k8s.io_~V_SubjectAccessReview|~X|test",
	"authorization.k8s.io_~V_SubjectAccessReview|test-ns|test",
	"certificates.k8s.io_~V_CertificateSigningRequest|~X|test",
	"certificates.k8s.io_~V_CertificateSigningRequest|test-ns|test",
	"events.k8s.io_~V_Event|~X|test",
	"events.k8s.io_~V_Event|test-ns|test",
	"node.k8s.io_~V_RuntimeClass|~X|test",
	"node.k8s.io_~V_RuntimeClass|test-ns|test",
	"policy_~V_PodDisruptionBudget|~X|test",
	"policy_~V_PodDisruptionBudget|test-ns|test",
	"policy_~V_PodSecurityPolicy|~X|test",
	"policy_~V_PodSecurityPolicy|test-ns|test",
	"apps_~V_ControllerRevision|~X|test",
	"apps_~V_ControllerRevision|test-ns|test",
	"apps_~V_DaemonSet|~X|test",
	"apps_~V_DaemonSet|test-ns|test",
	"apps_~V_Deployment|~X|test",
	"apps_~V_Deployment|test-ns|test",
	"apps_~V_ReplicaSet|~X|test",
	"apps_~V_ReplicaSet|test-ns|test",
	"apps_~V_StatefulSet|~X|test",
	"apps_~V_StatefulSet|test-ns|test",
	"authentication.k8s.io_~V_TokenReview|~X|test",
	"authentication.k8s.io_~V_TokenReview|test-ns|test",
	"networking.k8s.io_~V_IngressClass|~X|test",
	"networking.k8s.io_~V_IngressClass|test-ns|test",
	"networking.k8s.io_~V_Ingress|~X|test",
	"networking.k8s.io_~V_Ingress|test-ns|test",
	"networking.k8s.io_~V_NetworkPolicy|~X|test",
	"networking.k8s.io_~V_NetworkPolicy|test-ns|test",
	"scheduling.k8s.io_~V_PriorityClass|~X|test",
	"scheduling.k8s.io_~V_PriorityClass|test-ns|test",
	"storage.k8s.io_~V_CSIDriver|~X|test",
	"storage.k8s.io_~V_CSIDriver|test-ns|test",
	"storage.k8s.io_~V_CSINode|~X|test",
	"storage.k8s.io_~V_CSINode|test-ns|test",
	"storage.k8s.io_~V_StorageClass|~X|test",
	"storage.k8s.io_~V_StorageClass|test-ns|test",
	"storage.k8s.io_~V_VolumeAttachment|~X|test",
	"storage.k8s.io_~V_VolumeAttachment|test-ns|test",
	"~G_~V_Binding|~X|test",
	"~G_~V_Binding|test-ns|test",
	"~G_~V_ComponentStatus|~X|test",
	"~G_~V_ComponentStatus|test-ns|test",
	"~G_~V_ConfigMap|~X|test",
	"~G_~V_ConfigMap|test-ns|test",
	"~G_~V_Endpoints|~X|test",
	"~G_~V_Endpoints|test-ns|test",
	"~G_~V_Event|~X|test",
	"~G_~V_Event|test-ns|test",
	"~G_~V_LimitRange|~X|test",
	"~G_~V_LimitRange|test-ns|test",
	"~G_~V_Node|~X|test",
	"~G_~V_PersistentVolumeClaim|~X|test",
	"~G_~V_PersistentVolumeClaim|test-ns|test",
	"~G_~V_PersistentVolume|~X|test",
	"~G_~V_PersistentVolume|test-ns|test",
	"~G_~V_Pod|~X|test",
	"~G_~V_Pod|test-ns|test",
	"~G_~V_PodTemplate|~X|test",
	"~G_~V_PodTemplate|test-ns|test",
	"~G_~V_ReplicationController|~X|test",
	"~G_~V_ReplicationController|test-ns|test",
	"~G_~V_ResourceQuota|~X|test",
	"~G_~V_ResourceQuota|test-ns|test",
	"~G_~V_Secret|~X|test",
	"~G_~V_Secret|test-ns|test",
	"~G_~V_ServiceAccount|~X|test",
	"~G_~V_ServiceAccount|test-ns|test",
	"~G_~V_Service|~X|test",
	"~G_~V_Service|test-ns|test",
}

var residListLast = []string{
	"admissionregistration.k8s.io_~V_MutatingWebhookConfiguration|~X|test",
	"admissionregistration.k8s.io_~V_MutatingWebhookConfiguration|test-ns|test",
	"admissionregistration.k8s.io_~V_ValidatingWebhookConfiguration|~X|test",
	"admissionregistration.k8s.io_~V_ValidatingWebhookConfiguration|test-ns|test",
}
