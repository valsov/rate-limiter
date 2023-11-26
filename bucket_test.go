package rlimit

import (
	"testing"
	"time"
)

func TestBucketTryAllow(t *testing.T) {
	baseTime := time.Now().UTC()
	config := BucketConfig{
		RefillRate:  time.Hour,
		RefillCount: 10,
		BucketSize:  15,
	}
	userValue := BucketValue{
		LastRefillUtc:   baseTime,
		RemainingTokens: 15,
	}

	testCases := []struct {
		count                   int
		nowUtc                  time.Time
		expectedAllow           bool
		expectedLastRefillUtc   time.Time
		expectedRemainingTokens int
	}{
		{
			count:                   1,
			nowUtc:                  baseTime.Add(10 * time.Minute),
			expectedAllow:           true,
			expectedLastRefillUtc:   baseTime,
			expectedRemainingTokens: 14,
		},
		{
			count:                   14,
			nowUtc:                  baseTime.Add(20 * time.Minute),
			expectedAllow:           true,
			expectedLastRefillUtc:   baseTime,
			expectedRemainingTokens: 0,
		},
		{
			count:                   1,
			nowUtc:                  baseTime.Add(30 * time.Minute),
			expectedAllow:           false,
			expectedLastRefillUtc:   baseTime,
			expectedRemainingTokens: 0,
		},
		{
			count:                   1,
			nowUtc:                  baseTime.Add(time.Hour),
			expectedAllow:           true,
			expectedLastRefillUtc:   baseTime.Add(time.Hour),
			expectedRemainingTokens: 9, // RefillCount is only 10 while BucketSize is 15
		},
		{
			count:                   1,
			nowUtc:                  baseTime.Add(2 * time.Hour),
			expectedAllow:           true,
			expectedLastRefillUtc:   baseTime.Add(2 * time.Hour),
			expectedRemainingTokens: 14, // Fully refilled but allow 1
		},
	}
	for _, tc := range testCases {
		limiter := BucketLimiter{}
		result, newUserValue := limiter.TryAllow(tc.count, config, userValue, tc.nowUtc)

		if result != tc.expectedAllow {
			t.Fatalf("wrong result. expected=%v got=%v", tc.expectedAllow, result)
		}
		if newUserValue.LastRefillUtc != tc.expectedLastRefillUtc {
			t.Fatalf("wrong userValue.LastRefillUtc. expected=%v got=%v", tc.expectedLastRefillUtc, userValue.LastRefillUtc)
		}
		if newUserValue.RemainingTokens != tc.expectedRemainingTokens {
			t.Fatalf("wrong userValue.RemainingTokens. expected=%d got=%d", tc.expectedRemainingTokens, userValue.RemainingTokens)
		}

		// Update user value for next test cases
		userValue = newUserValue
	}
}
