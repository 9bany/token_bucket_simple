package rule

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetBucket(t *testing.T) {
	userType := "gen-user"
	indentifier := "123.34556.4545.4.45"
	bucket := GetBucket(indentifier, userType)

	require.Equal(t, int64(20), bucket.MaxTokens)
	require.Equal(t, int64(20), bucket.Rate)

}
