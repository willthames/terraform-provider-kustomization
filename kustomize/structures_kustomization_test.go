package kustomize

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/kustomize/api/filesys"
	"sigs.k8s.io/kustomize/api/krusty"
)

func TestFlattenKustomizationIDs(t *testing.T) {
	fSys := filesys.MakeFsOnDisk()
	opts := krusty.MakeDefaultOptions()
	k := krusty.MakeKustomizer(fSys, opts)

	rm, err := k.Run("test_kustomizations/basic/initial")
	assert.Equal(t, err, nil, nil)

	ids, idsPrio := flattenKustomizationIDs(rm)

	expMerged := append(idsPrio[0], idsPrio[1]...)
	expMerged = append(expMerged, idsPrio[2]...)
	assert.ElementsMatch(t, expMerged, ids, nil)

	expIds := []string{"~G_~V_Namespace|~X|test-basic", "apps_~V_Deployment|test-basic|test", "networking.k8s.io_~V_Ingress|test-basic|test", "~G_~V_Service|test-basic|test"}
	assert.ElementsMatch(t, expIds, ids, nil)

	expP1 := []string{"~G_~V_Namespace|~X|test-basic"}
	assert.ElementsMatch(t, expP1, idsPrio[0], nil)

	expP2 := []string{"apps_~V_Deployment|test-basic|test", "networking.k8s.io_~V_Ingress|test-basic|test", "~G_~V_Service|test-basic|test"}
	assert.ElementsMatch(t, expP2, idsPrio[1], nil)

	expP3 := []string{}
	assert.ElementsMatch(t, expP3, idsPrio[2], nil)
}
