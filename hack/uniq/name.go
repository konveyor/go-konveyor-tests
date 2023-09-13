package uniq

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/konveyor/tackle2-hub/api"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var charset = "QWERTZUIOPASDFGHJKLYXCVBNMqwertzuiopasdfghjklyxcvbnm0123456789"

// Random string for given lenght with [A-Za-z0-9] chars.
func RandString(lenght int) string {
	result := make([]byte, lenght)
	for i := range result {
		result[i] = charset[r.Intn(len(charset))]
	}
	return string(result)
}

// Append random 5 chars to Application Name.
func ApplicationName(r *api.Application) {
	r.Name = fmt.Sprintf("%s %s", r.Name, RandString(5))
}

// Append random 3 chars to RuleSet Name.
func RuleSetName(r *api.RuleSet) {
	r.Name = fmt.Sprintf("%s %s", r.Name, RandString(3))
}

// IdentityName returns random Identity Name
func IdentityName(r *api.Identity) {
	r.Name = fmt.Sprint(r.Kind, "-", RandString(5))
}

// TrackerName returns random Tracker Name
func TrackerName(r *api.Tracker) {
	r.Name = fmt.Sprint(r.Kind, "-", RandString(5))
}
