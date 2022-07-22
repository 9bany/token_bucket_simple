package token_bucket

import (
	"math"
	"sync"
	"time"
)

type TokenBucket struct {
	rate                int64
	maxTokens           int64
	currentTokens       int64
	lastRefillTimestamp time.Time
	mutex               sync.Mutex
}

/*
* Returns a new bucket for a client if the client doesn’t exist in the system
* and sets his current tokens to the max tokens.
 */
func NewTokenBucket(Rate int64, MaxTokens int64) *TokenBucket {
	return &TokenBucket{
		rate:                Rate,
		maxTokens:           MaxTokens,
		lastRefillTimestamp: time.Now(),
		currentTokens:       MaxTokens,
	}
}

/*
* Refills the number of tokens needed to be added since the time elapsed after last refill.
* Also we will reset the lastRefillTimeStamp in this for the next request.
 */
func (tokenBucket *TokenBucket) refill() {
	now := time.Now()
	end := time.Since(tokenBucket.lastRefillTimestamp)

	tokensTobeAdded := (end.Nanoseconds() * tokenBucket.rate) / 1000000000
	tokenBucket.currentTokens = int64(math.Min(float64(tokenBucket.currentTokens+tokensTobeAdded), float64(tokenBucket.maxTokens)))

	tokenBucket.lastRefillTimestamp = now
}

/*
* The one which takes the decision wether the request should be discarded or not.
* It first refills the clients bucket and after that checks if the client can perform
* the required request or not.
 */
func (tokenBucket *TokenBucket) IsRequestAllowed(token int64) bool {
	tokenBucket.mutex.Lock()
	defer tokenBucket.mutex.Unlock()
	tokenBucket.refill()
	if tokenBucket.currentTokens >= token {
		tokenBucket.currentTokens = tokenBucket.currentTokens - token
		return true
	}
	return false
}
