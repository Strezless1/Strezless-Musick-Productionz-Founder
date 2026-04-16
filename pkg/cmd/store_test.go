// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/omar-orrantia/Strezless-Musick-Productionz-Founder/internal/mocktest"
)

func TestStoreListInventory(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"store", "list-inventory",
		)
	})
}
