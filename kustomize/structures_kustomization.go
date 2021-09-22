package kustomize

import (
	"regexp"

	"sigs.k8s.io/kustomize/api/resmap"
)

var reVersionMatch = regexp.MustCompile("_.*_")

func removeVersion(id string) string {
	return reVersionMatch.ReplaceAllString(id, "_~V_")
}

func flattenKustomizationIDs(rm resmap.ResMap) (ids []string, idsPrio [][]string) {
	p0 := []string{}
	p1 := []string{}
	p2 := []string{}
	for _, id := range rm.AllIds() {
		versionLessId := removeVersion(id.String())
		ids = append(ids, versionLessId)

		p := determinePrefix(id.Gvk)
		if p < 5 {
			p0 = append(p0, versionLessId)
		} else if p == 9 {
			p2 = append(p2, versionLessId)
		} else {
			p1 = append(p1, versionLessId)
		}
	}

	idsPrio = append(idsPrio, p0)
	idsPrio = append(idsPrio, p1)
	idsPrio = append(idsPrio, p2)

	return ids, idsPrio
}

func flattenKustomizationResources(rm resmap.ResMap) (res map[string]string, err error) {
	res = make(map[string]string)
	for _, r := range rm.Resources() {
		id := removeVersion(r.CurId().String())
		json, err := r.MarshalJSON()
		if err != nil {
			return nil, err
		}
		res[id] = string(json)
	}

	return res, nil
}
