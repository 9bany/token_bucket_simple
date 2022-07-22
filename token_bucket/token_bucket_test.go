package token_bucket

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func createBucket(t *testing.T, maxTokens int64, rate int64) *TokenBucket {
	bucket := NewTokenBucket(rate, maxTokens)

	require.Equal(t, int64(rate), bucket.Rate)
	require.Equal(t, int64(maxTokens), bucket.MaxTokens)
	return bucket
}

func TestCreateBucket(t *testing.T) {
	createBucket(t, int64(2), int64(20))
}

func TestIsRequestAllowedTrue(t *testing.T) {
	bucket := createBucket(t, int64(2), int64(20))
	maxTokens := bucket.MaxTokens
	resultAllowed := bucket.IsRequestAllowed(1)

	require.Equal(t, resultAllowed, true)
	require.Equal(t, maxTokens-1, bucket.currentTokens)
}

func TestIsRequestAlloweFalse(t *testing.T) {
	bucket := createBucket(t, int64(2), int64(20))
	maxTokens := bucket.MaxTokens
	resultAllowed := bucket.IsRequestAllowed(4)

	require.Equal(t, resultAllowed, false)
	require.Equal(t, maxTokens, bucket.currentTokens)
}
