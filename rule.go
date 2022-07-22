package main

var clientBucketMap = make(map[string]*TokenBucket)

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
	"gen-user": {MaxTokens: 2, Rate: 5},
}

func GetBucket(indentifier string, userType string) *TokenBucket {
	if clientBucketMap[indentifier] == nil {
		clientBucketMap[indentifier] = NewTokenBucket(rulesMap[userType].MaxTokens, rulesMap[userType].Rate)
	}
	return clientBucketMap[indentifier]
}
