// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/omar-orrantia/Strezless-Musick-Productionz-Founder/internal/mocktest"
	"github.com/omar-orrantia/Strezless-Musick-Productionz-Founder/internal/requestflag"
)

func TestPetCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pet", "create",
			"--name", "doggie",
			"--photo-url", "string",
			"--id", "10",
			"--category", "{id: 1, name: Dogs}",
			"--status", "available",
			"--tag", "{id: 0, name: name}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(petCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: doggie\n" +
			"photoUrls:\n" +
			"  - string\n" +
			"id: 10\n" +
			"category:\n" +
			"  id: 1\n" +
			"  name: Dogs\n" +
			"status: available\n" +
			"tags:\n" +
			"  - id: 0\n" +
			"    name: name\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"pet", "create",
		)
	})
}

func TestPetRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pet", "retrieve",
			"--pet-id", "0",
		)
	})
}

func TestPetUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pet", "update",
			"--name", "doggie",
			"--photo-url", "string",
			"--id", "10",
			"--category", "{id: 1, name: Dogs}",
			"--status", "available",
			"--tag", "{id: 0, name: name}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(petUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: doggie\n" +
			"photoUrls:\n" +
			"  - string\n" +
			"id: 10\n" +
			"category:\n" +
			"  id: 1\n" +
			"  name: Dogs\n" +
			"status: available\n" +
			"tags:\n" +
			"  - id: 0\n" +
			"    name: name\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"pet", "update",
		)
	})
}

func TestPetDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pet", "delete",
			"--pet-id", "0",
		)
	})
}

func TestPetFindByStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pet", "find-by-status",
			"--status", "available",
		)
	})
}

func TestPetFindByTags(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pet", "find-by-tags",
			"--tag", "string",
		)
	})
}

func TestPetUpdateWithForm(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pet", "update-with-form",
			"--pet-id", "0",
			"--name", "name",
			"--status", "status",
		)
	})
}

func TestPetUploadImage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pet", "upload-image",
			"--pet-id", "0",
			"--body", mocktest.TestFile(t, "Example data"),
			"--additional-metadata", "additionalMetadata",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("Example data")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"pet", "upload-image",
			"--pet-id", "0",
			"--additional-metadata", "additionalMetadata",
		)
	})
}
