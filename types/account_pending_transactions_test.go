package types

import (
	"encoding/json"
	"testing"
)

func TestMarshalTransactionStatus(t *testing.T) {
	testCases := []struct {
		name         string
		status       TransactionStatus
		expectedJSON string
	}{
		{
			name:         "Packed status",
			status:       TransactionStatus{Status: "packed"},
			expectedJSON: `"packed"`,
		},
		{
			name:         "Ready status",
			status:       TransactionStatus{Status: "ready"},
			expectedJSON: `"ready"`,
		},
		{
			name:         "Pending status - future nonce",
			status:       TransactionStatus{Status: "pending", PendingReason: "futureNonce"},
			expectedJSON: `{"pending":"futureNonce"}`,
		},
		{
			name:         "Pending status - not enough cash",
			status:       TransactionStatus{Status: "pending", PendingReason: "notEnoughCash"},
			expectedJSON: `{"pending":"notEnoughCash"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := json.Marshal(tc.status)
			if err != nil {
				t.Fatalf("Failed to marshal: %v", err)
			}
			if string(result) != tc.expectedJSON {
				t.Errorf("Expected %s, got %s", tc.expectedJSON, string(result))
			}
		})
	}
}

func TestUnmarshalTransactionStatus(t *testing.T) {
	testCases := []struct {
		name           string
		jsonData       string
		expectedStatus TransactionStatus
	}{
		{
			name:           "Packed status",
			jsonData:       `"packed"`,
			expectedStatus: TransactionStatus{Status: "packed"},
		},
		{
			name:           "Ready status",
			jsonData:       `"ready"`,
			expectedStatus: TransactionStatus{Status: "ready"},
		},
		{
			name:           "Pending status - future nonce",
			jsonData:       `{"pending":"futureNonce"}`,
			expectedStatus: TransactionStatus{Status: "pending", PendingReason: "futureNonce"},
		},
		{
			name:           "Pending status - not enough cash",
			jsonData:       `{"pending":"notEnoughCash"}`,
			expectedStatus: TransactionStatus{Status: "pending", PendingReason: "notEnoughCash"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result TransactionStatus
			err := json.Unmarshal([]byte(tc.jsonData), &result)
			if err != nil {
				t.Fatalf("Failed to unmarshal: %v", err)
			}
			if result != tc.expectedStatus {
				t.Errorf("Expected %+v, got %+v", tc.expectedStatus, result)
			}
		})
	}
}
