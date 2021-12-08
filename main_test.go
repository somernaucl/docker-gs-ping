package main

import (
	//"fmt"
	
	//"net/http"
	//"net/http/httptest"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
)

// The tests in this suite check if the business requirements
// have been satisfied. Each test checks a single requirement.

// Every test in this suite is isolated from all other tests,
// there is no data or container instances shared at all.

// Beware, that by default HTTP client does not set a timeout.
// In tests, however, this isn't an issue because the tests *do*
// have a timeout, which defaults to 10 minutes, but can be reduced
// via a command-line parameter.

// After a container has started, there is some time required before
// the application inside it is ready to service the incoming requests.
// This creates a race condition with the test. To avoid this, the tests
// retry with exponential back-off.

// Requirement 1: The application response must contain a heart ("<3")


// Requirement 2: The application must include a health-check end-point at /ping,
// responding with JSON document { "Status": "OK" } if everything is well.
func TestHealthCheck(t *testing.T) {

	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "could not connect to Docker")

	resource, err := pool.Run("docker-gs-ping", "latest", []string{})
	require.NoError(t, err, "could not start container")

	t.Cleanup(func() {
		require.NoError(t, pool.Purge(resource), "failed to remove container")
	})
}


