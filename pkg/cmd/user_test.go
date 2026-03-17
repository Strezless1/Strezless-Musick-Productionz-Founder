// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/omar-orrantia/Strezless-Musick-Productionz-Founder/internal/mocktest"
	"github.com/omar-orrantia/Strezless-Musick-Productionz-Founder/internal/requestflag"
)

func TestUserCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"user", "create",
			"--id", "10",
			"--email", "john@email.com",
			"--first-name", "John",
			"--last-name", "James",
			"--password", "12345",
			"--phone", "12345",
			"--username", "theUser",
			"--user-status", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: 10\n" +
			"email: john@email.com\n" +
			"firstName: John\n" +
			"lastName: James\n" +
			"password: '12345'\n" +
			"phone: '12345'\n" +
			"username: theUser\n" +
			"userStatus: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"user", "create",
		)
	})
}

func TestUserRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"user", "retrieve",
			"--username", "username",
		)
	})
}

func TestUserUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"user", "update",
			"--existing-username", "username",
			"--id", "10",
			"--email", "john@email.com",
			"--first-name", "John",
			"--last-name", "James",
			"--password", "12345",
			"--phone", "12345",
			"--username", "theUser",
			"--user-status", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: 10\n" +
			"email: john@email.com\n" +
			"firstName: John\n" +
			"lastName: James\n" +
			"password: '12345'\n" +
			"phone: '12345'\n" +
			"username: theUser\n" +
			"userStatus: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"user", "update",
			"--existing-username", "username",
		)
	})
}

func TestUserDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"user", "delete",
			"--username", "username",
		)
	})
}

func TestUserCreateWithList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"user", "create-with-list",
			"--body", "{id: 10, email: john@email.com, firstName: John, lastName: James, password: '12345', phone: '12345', username: theUser, userStatus: 1}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(userCreateWithList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"user", "create-with-list",
			"--body.id", "10",
			"--body.email", "john@email.com",
			"--body.first-name", "John",
			"--body.last-name", "James",
			"--body.password", "12345",
			"--body.phone", "12345",
			"--body.username", "theUser",
			"--body.user-status", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"- id: 10\n" +
			"  email: john@email.com\n" +
			"  firstName: John\n" +
			"  lastName: James\n" +
			"  password: '12345'\n" +
			"  phone: '12345'\n" +
			"  username: theUser\n" +
			"  userStatus: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"user", "create-with-list",
		)
	})
}

func TestUserLogin(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"user", "login",
			"--password", "password",
			"--username", "username",
		)
	})
}

func TestUserLogout(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"user", "logout",
		)
	})
}
