package rlimit

import (
	"testing"
	"time"
)

func TestLeakyBucketTryAllow(t *testing.T) {
	baseTime := time.Now().UTC()
	config := LeakyBucketConfig{
		NewTokensRate:  time.Minute,
		NewTokensCount: 1,
		BucketSize:     10,
	}
	userValue := LeakyBucketValue{
		LastCheckedUtc:  baseTime,
		RemainingTokens: 10,
	}

	testCases := []struct {
		count                   int
		nowUtc                  time.Time
		expectedAllow           bool
		expectedRemainingTokens int
	}{
		{
			count:                   6,
			nowUtc:                  baseTime.Add(time.Minute),
			expectedAllow:           true,
			expectedRemainingTokens: 4,
		},
		{
			count:                   7,
			nowUtc:                  baseTime.Add(4 * time.Minute), // 3 minutes after last request
			expectedAllow:           true,
			expectedRemainingTokens: 0,
		},
		{
			count:                   1,
			nowUtc:                  baseTime.Add(4 * time.Minute), // Same time as previous
			expectedAllow:           false,
			expectedRemainingTokens: 0,
		},
		{
			count:                   1,
			nowUtc:                  baseTime.Add(6 * time.Minute),
			expectedAllow:           true,
			expectedRemainingTokens: 1,
		},
	}
	for _, tc := range testCases {
		limiter := LeakyBucketLimiter{}
		result, newUserValue := limiter.TryAllow(tc.count, config, userValue, tc.nowUtc)

		if result != tc.expectedAllow {
			t.Fatalf("wrong result. expected=%v got=%v", tc.expectedAllow, result)
		}
		if newUserValue.LastCheckedUtc != tc.nowUtc {
			t.Fatalf("wrong userValue.LastCheckedUtc. expected=%v got=%v", tc.nowUtc, userValue.LastCheckedUtc)
		}
		if newUserValue.RemainingTokens != tc.expectedRemainingTokens {
			t.Fatalf("wrong userValue.RemainingTokens. expected=%d got=%d", tc.expectedRemainingTokens, userValue.RemainingTokens)
		}

		// Update user value for next test cases
		userValue = newUserValue
	}
}
