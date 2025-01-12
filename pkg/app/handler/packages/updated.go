// Used to show recently updated versions

package packages

import (
	"net/http"
	"soko/pkg/app/handler/feeds"
)

// Updated renders a template containing
// a list of 50 recently updated versions
func Updated(w http.ResponseWriter, r *http.Request) {
	updatedVersions := GetUpdatedVersions(50)
	renderPackageTemplates("changedVersions", "changedVersions", "changedVersionRow", GetFuncMap(), CreateFeedData("Updated", updatedVersions), w)
}

func UpdatedFeed(w http.ResponseWriter, r *http.Request) {
	updatedVersions := GetUpdatedVersions(250)
	feeds.Changes("Added packages in Gentoo.", "Added packages in Gentoo.", updatedVersions, w)
}
