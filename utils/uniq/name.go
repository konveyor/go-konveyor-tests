package uniq

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/konveyor/tackle2-hub/api"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var charset = "QWERTZUIOPASDFGHJKLYXCVBNMqwertzuiopasdfghjklyxcvbnm0123456789"

//
// Random string for given lenght with [A-Za-z0-9] chars.
func RandString(lenght int) string {
	result := make([]byte, lenght)
	for i := range result {
		result[i] = charset[r.Intn(len(charset))]
	}
	return string(result)
}

//
// Append random 5 chars to Application Name.
func ApplicationName(r *api.Application) {
	r.Name = fmt.Sprintf("%s %s", r.Name, RandString(5))
}
