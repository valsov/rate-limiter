package rlimit

import (
	"testing"
	"time"
)

func TestSlidingWindowTryAllow(t *testing.T) {
	baseTime := time.Now().UTC()
	config := SlidingWindowConfig{
		WindowLength: time.Hour,
		MaxTokens:    10,
	}
	userValue := SlidingWindowValue{
		PreviousConsumption: 0,
		CurrentConsumption:  0,
		WindowStartUtc:      baseTime,
	}

	testCases := []struct {
		count                       int
		nowUtc                      time.Time
		expectedAllow               bool
		expectedPreviousConsumption int
		expectedCurrentConsumption  int
		expectedWindowStartUtc      time.Time
	}{
		{
			count:                       5,
			nowUtc:                      baseTime.Add(time.Minute),
			expectedAllow:               true,
			expectedPreviousConsumption: 0,
			expectedCurrentConsumption:  5,
			expectedWindowStartUtc:      baseTime,
		},
		{
			count:                       6,
			nowUtc:                      baseTime.Add(time.Minute),
			expectedAllow:               false,
			expectedPreviousConsumption: 0,
			expectedCurrentConsumption:  5,
			expectedWindowStartUtc:      baseTime,
		},
		{
			count:                       5,
			nowUtc:                      baseTime.Add(time.Hour),
			expectedAllow:               true,
			expectedPreviousConsumption: 5,
			expectedCurrentConsumption:  5,
			expectedWindowStartUtc:      baseTime.Add(time.Hour),
		},
		{
			count:                       4,
			nowUtc:                      baseTime.Add(time.Hour + 15*time.Minute),
			expectedAllow:               false,
			expectedPreviousConsumption: 5,
			expectedCurrentConsumption:  5,
			expectedWindowStartUtc:      baseTime.Add(time.Hour),
		},
		{
			count:                       2,
			nowUtc:                      baseTime.Add(time.Hour + 15*time.Minute),
			expectedAllow:               true,
			expectedPreviousConsumption: 5,
			expectedCurrentConsumption:  7,
			expectedWindowStartUtc:      baseTime.Add(time.Hour),
		},
		{
			count:                       5,
			nowUtc:                      baseTime.Add(3 * time.Hour),
			expectedAllow:               true,
			expectedPreviousConsumption: 0, // Previous window is out of bounds: reset previous counter
			expectedCurrentConsumption:  5,
			expectedWindowStartUtc:      baseTime.Add(3 * time.Hour),
		},
	}
	for _, tc := range testCases {
		limiter := SlidingWindowLimiter{}
		result, newUserValue := limiter.TryAllow(tc.count, config, userValue, tc.nowUtc)

		if result != tc.expectedAllow {
			t.Fatalf("wrong result. expected=%v got=%v", tc.expectedAllow, result)
		}
		if newUserValue.PreviousConsumption != tc.expectedPreviousConsumption {
			t.Fatalf("wrong userValue.PreviousConsumption. expected=%d got=%d", tc.expectedPreviousConsumption, userValue.PreviousConsumption)
		}
		if newUserValue.CurrentConsumption != tc.expectedCurrentConsumption {
			t.Fatalf("wrong userValue.CurrentConsumption. expected=%d got=%d", tc.expectedCurrentConsumption, userValue.CurrentConsumption)
		}
		if newUserValue.WindowStartUtc != tc.expectedWindowStartUtc {
			t.Fatalf("wrong userValue.WindowStartUtc. expected=%v got=%v", tc.expectedWindowStartUtc, userValue.WindowStartUtc)
		}

		// Update user value for next test cases
		userValue = newUserValue
	}
}
