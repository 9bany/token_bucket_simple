package rule

import bucket "9bany/rate-limiter-token-bucket/token_bucket"

var clientBucketMap = make(map[string]*bucket.TokenBucket)

type Rule struct {
	MaxTokens int64
	Rate      int64
}

/*
* These rules map can be fetched form a database rather than
* herdcoding at application layer and also can be periodically
* update in a background job in case some dynamic changed are needed.
 */
var rulesMap = map[string]Rule{
	"gen-user": {MaxTokens: 20, Rate: 20},
}

func GetBucket(indentifier string, userType string) *bucket.TokenBucket {
	if clientBucketMap[indentifier] == nil {
		clientBucketMap[indentifier] = bucket.NewTokenBucket(rulesMap[userType].Rate, rulesMap[userType].MaxTokens)
	}
	return clientBucketMap[indentifier]
}
