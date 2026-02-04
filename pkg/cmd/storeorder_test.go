// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/strezless-musick-nexus-metadata-cli/internal/mocktest"
)

func TestStoreOrderCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"store:order", "create",
		"--id", "10",
		"--complete=true",
		"--pet-id", "198772",
		"--quantity", "7",
		"--ship-date", "2019-12-27T18:11:19.117Z",
		"--status", "approved",
	)
}

func TestStoreOrderRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"store:order", "retrieve",
		"--order-id", "0",
	)
}

func TestStoreOrderDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"store:order", "delete",
		"--order-id", "0",
	)
}
