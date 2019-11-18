package jsonrpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/torusresearch/torus-common/bijson"
)

func TestRequestID(t *testing.T) {

	c := context.Background()
	id := bijson.RawMessage("1")
	c = WithRequestID(c, &id)
	var pick *bijson.RawMessage
	require.NotPanics(t, func() {
		pick = RequestID(c)
	})
	require.Equal(t, &id, pick)
}
