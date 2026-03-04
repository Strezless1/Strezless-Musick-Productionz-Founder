// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/omar-orrantia/Strezless-Musick-Productionz-Founder/internal/mocktest"
	"github.com/omar-orrantia/Strezless-Musick-Productionz-Founder/internal/requestflag"
)

func TestPetCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"pet", "create",
		"--api-key", "string",
		"--name", "doggie",
		"--photo-url", "string",
		"--id", "10",
		"--category", "{id: 1, name: Dogs}",
		"--status", "available",
		"--tag", "{id: 0, name: name}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(petCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"pet", "create",
		"--name", "doggie",
		"--photo-url", "string",
		"--id", "10",
		"--category.id", "1",
		"--category.name", "Dogs",
		"--status", "available",
		"--tag.id", "0",
		"--tag.name", "name",
	)
}

func TestPetRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"pet", "retrieve",
		"--api-key", "string",
		"--pet-id", "0",
	)
}

func TestPetUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"pet", "update",
		"--api-key", "string",
		"--name", "doggie",
		"--photo-url", "string",
		"--id", "10",
		"--category", "{id: 1, name: Dogs}",
		"--status", "available",
		"--tag", "{id: 0, name: name}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(petUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"pet", "update",
		"--name", "doggie",
		"--photo-url", "string",
		"--id", "10",
		"--category.id", "1",
		"--category.name", "Dogs",
		"--status", "available",
		"--tag.id", "0",
		"--tag.name", "name",
	)
}

func TestPetDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"pet", "delete",
		"--api-key", "string",
		"--pet-id", "0",
	)
}

func TestPetFindByStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"pet", "find-by-status",
		"--api-key", "string",
		"--status", "available",
	)
}

func TestPetFindByTags(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"pet", "find-by-tags",
		"--api-key", "string",
		"--tag", "string",
	)
}

func TestPetUpdateWithForm(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"pet", "update-with-form",
		"--api-key", "string",
		"--pet-id", "0",
		"--name", "name",
		"--status", "status",
	)
}

func TestPetUploadImage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"pet", "upload-image",
		"--api-key", "string",
		"--pet-id", "0",
		"--body", mocktest.TestFile(t, "..."),
		"--additional-metadata", "additionalMetadata",
	)
}
