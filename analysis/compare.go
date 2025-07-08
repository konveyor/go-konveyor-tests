package analysis

import (
	"fmt"

	"github.com/konveyor/tackle2-hub/api"
)

// Compare diffs for resources collections relevant for Analysis
func getTagsDiff(want, got []api.Tag) (notFound []api.Tag, extras []api.Tag) {
	tagKey := func(t api.Tag) string { return fmt.Sprintf("%s-%s", t.Category.Name, t.Name) }
	wantTags := map[string]api.Tag{}
	gotTags := map[string]api.Tag{}
	for _, gotTag := range got {
		gotTags[tagKey(gotTag)] = gotTag
	}
	for _, wantTag := range want {
		wantTags[tagKey(wantTag)] = wantTag
		// find tags we expect, but not present in output
		if _, found := gotTags[tagKey(wantTag)]; !found {
			notFound = append(notFound, wantTag)
		}
	}
	for _, gotTag := range got {
		// find tags we got, but we didn't expect
		if _, found := wantTags[tagKey(gotTag)]; !found {
			extras = append(extras, gotTag)
		}
	}
	return
}

func getInsightsDiff(wantList, gotList []api.Insight) (notFound []api.Insight, extras []api.Insight) {
	key := func(r api.Insight) string { return fmt.Sprintf("%s-%s-%s-%d", r.Category, r.RuleSet, r.Rule, r.Effort) }
	wantMap := map[string]api.Insight{}
	gotMap := map[string]api.Insight{}
	for _, got := range gotList {
		gotMap[key(got)] = got
	}
	for _, want := range wantList {
		wantMap[key(want)] = want
		// find what we expect, but not present in output
		if _, found := gotMap[key(want)]; !found {
			notFound = append(notFound, want)
		}
	}
	for _, got := range gotList {
		// find what we got, but we didn't expect
		if _, found := wantMap[key(got)]; !found {
			extras = append(extras, got)
		}
	}
	return
}

func getIncidentsDiff(wantList, gotList []api.Incident) (notFound []api.Incident, extras []api.Incident) {
	key := func(r api.Incident) string { return fmt.Sprintf("%s-%d", r.File, r.Line) }
	wantMap := map[string]api.Incident{}
	gotMap := map[string]api.Incident{}
	for _, got := range gotList {
		gotMap[key(got)] = got
	}
	for _, want := range wantList {
		wantMap[key(want)] = want
		// find tags we expect, but not present in output
		if _, found := gotMap[key(want)]; !found {
			notFound = append(notFound, want)
		}
	}
	for _, got := range gotList {
		// find tags we got, but we didn't expect
		if _, found := wantMap[key(got)]; !found {
			extras = append(extras, got)
		}
	}
	return
}
