// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/omar-orrantia/Strezless-Musick-Productionz-Founder/internal/mocktest"
)

func TestStoreOrderCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"store:order", "create",
		"--api-key", "string",
		"--id", "10",
		"--complete=true",
		"--pet-id", "198772",
		"--quantity", "7",
		"--ship-date", "'2019-12-27T18:11:19.117Z'",
		"--status", "approved",
	)
}

func TestStoreOrderRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"store:order", "retrieve",
		"--api-key", "string",
		"--order-id", "0",
	)
}

func TestStoreOrderDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"store:order", "delete",
		"--api-key", "string",
		"--order-id", "0",
	)
}
