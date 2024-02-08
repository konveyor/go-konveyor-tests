package utils

import (
	"time"

	"github.com/konveyor/tackle2-hub/api"
	"github.com/onsi/gomega"
)

func CheckReferenceNotEmpty(ticket *api.Ticket) string {
	// Wait for reference field to be populated
	var got *api.Ticket
	var err error
	for i := 0; i < RETRY_NUM; i++ {
		got, err = Ticket.Get(ticket.ID)
		if err != nil || got.Reference != "" {
			break
		}
		time.Sleep(5 * time.Second)
	}
	gomega.Expect(got.Reference).NotTo(gomega.BeEmpty())
	return got.Reference
}
