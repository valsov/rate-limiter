package rlimit

import (
	"testing"
	"time"
)

func TestFixedWindowTryAllow(t *testing.T) {
	baseTime := time.Now().UTC()
	config := FixedWindowConfig{
		WindowLength: time.Hour,
		MaxTokens:    10,
	}
	userValue := FixedWindowValue{
		WindowStartUtc:  baseTime,
		RemainingTokens: 10,
	}

	testCases := []struct {
		count                   int
		nowUtc                  time.Time
		expectedAllow           bool
		expectedWindowStartUtc  time.Time
		expectedRemainingTokens int
	}{
		{
			count:                   1,
			nowUtc:                  baseTime.Add(10 * time.Minute),
			expectedAllow:           true,
			expectedWindowStartUtc:  baseTime,
			expectedRemainingTokens: 9,
		},
		{
			count:                   9,
			nowUtc:                  baseTime.Add(20 * time.Minute),
			expectedAllow:           true,
			expectedWindowStartUtc:  baseTime,
			expectedRemainingTokens: 0,
		},
		{
			count:                   1,
			nowUtc:                  baseTime.Add(30 * time.Minute),
			expectedAllow:           false,
			expectedWindowStartUtc:  baseTime,
			expectedRemainingTokens: 0,
		},
		{
			count:                   1,
			nowUtc:                  baseTime.Add(time.Hour),
			expectedAllow:           true,
			expectedWindowStartUtc:  baseTime.Add(time.Hour),
			expectedRemainingTokens: 9,
		},
	}
	for _, tc := range testCases {
		limiter := FixedWindowLimiter{}
		result, newUserValue := limiter.TryAllow(tc.count, config, userValue, tc.nowUtc)

		if result != tc.expectedAllow {
			t.Fatalf("wrong result. expected=%v got=%v", tc.expectedAllow, result)
		}
		if newUserValue.WindowStartUtc != tc.expectedWindowStartUtc {
			t.Fatalf("wrong userValue.WindowStartUtc. expected=%v got=%v", tc.expectedWindowStartUtc, userValue.WindowStartUtc)
		}
		if newUserValue.RemainingTokens != tc.expectedRemainingTokens {
			t.Fatalf("wrong userValue.RemainingTokens. expected=%d got=%d", tc.expectedRemainingTokens, userValue.RemainingTokens)
		}

		// Update user value for next test cases
		userValue = newUserValue
	}
}
