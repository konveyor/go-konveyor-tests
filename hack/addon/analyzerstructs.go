package addon

import (
	"github.com/konveyor/tackle2-addon/repository"
	"github.com/konveyor/tackle2-hub/api"
)

//
// Originally from https://github.com/konveyor/tackle2-addon-windup/blob/main/cmd as of 2023-06-26
// Would love replace this when https://github.com/konveyor/tackle2-addon-windup/issues/98 gets resolved.

//
// Data Addon data passed in the secret.
type Data struct {
	// Output directory within application bucket.
	Output string `json:"output" binding:"required"`
	// Mode options.
	Mode Mode `json:"mode"`
	// Sources list.
	Sources Sources `json:"sources"`
	// Targets list.
	Targets Targets `json:"targets"`
	// Scope options.
	Scope Scope `json:"scope"`
	// Rules options.
	Rules Rules `json:"rules"`
	// Tagger options.
	Tagger Tagger `json:"tagger"`
}

//
// Mode settings.
type Mode struct {
	Binary     bool   `json:"binary"`
	Artifact   string `json:"artifact"`
	WithDeps   bool   `json:"withDeps"`
	Repository repository.SCM
	//
	//path struct {
	//	appDir string
	//	binary string
	//	maven  struct {
	//		settings string
	//	}
	//}
}

//
// Sources list of sources.
type Sources []string

//
// Targets list of target.
type Targets []string

//
// Scope settings.
type Scope struct {
	WithKnown bool `json:"withKnown"`
	Packages  struct {
		Included []string `json:"included,omitempty"`
		Excluded []string `json:"excluded,omitempty"`
	} `json:"packages"`
}

//
// Rules settings.
type Rules struct {
	Path       string          `json:"path"`
	Repository *api.Repository `json:"repository"`
	Identity   *api.Ref        `json:"identity"`
	Labels     Labels          `json:"labels"`
}

//
// Labels collection.
type Labels struct {
	Included []string `json:"included,omitempty"`
	Excluded []string `json:"excluded,omitempty"`
}

//
// Tagger tags an application.
type Tagger struct {
	Enabled bool `json:"enabled"`
}

////
//// RuleSet an XML document.
//type RuleSet struct {
//	Metadata struct {
//		Target struct {
//			ID string `xml:"id,attr"`
//		} `xml:"targetTechnology"`
//	} `xml:"metadata"`
//}
//
////
//// Windup application analyzer.
//type Windup struct {
//	application *api.Application
//	*Data
//}
