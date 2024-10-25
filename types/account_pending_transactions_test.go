package types

import (
	"encoding/json"
	"testing"
)

func TestMarshalTransactionStatus(t *testing.T) {
	测试用例 := []struct {
		名称     string
		状态     TransactionStatus
		期望JSON string
	}{
		{
			名称:     "已打包状态",
			状态:     TransactionStatus{Status: "packed"},
			期望JSON: `"packed"`,
		},
		{
			名称:     "就绪状态",
			状态:     TransactionStatus{Status: "ready"},
			期望JSON: `"ready"`,
		},
		{
			名称:     "待处理状态-未来nonce",
			状态:     TransactionStatus{Status: "pending", PendingReason: "futureNonce"},
			期望JSON: `{"pending":"futureNonce"}`,
		},
		{
			名称:     "待处理状态-余额不足",
			状态:     TransactionStatus{Status: "pending", PendingReason: "notEnoughCash"},
			期望JSON: `{"pending":"notEnoughCash"}`,
		},
	}

	for _, 案例 := range 测试用例 {
		t.Run(案例.名称, func(t *testing.T) {
			结果, 错误 := json.Marshal(案例.状态)
			if 错误 != nil {
				t.Fatalf("序列化失败: %v", 错误)
			}
			if string(结果) != 案例.期望JSON {
				t.Errorf("期望 %s, 得到 %s", 案例.期望JSON, string(结果))
			}
		})
	}
}

func TestUnmarshalTransactionStatus(t *testing.T) {
	测试用例 := []struct {
		名称     string
		JSON数据 string
		期望状态   TransactionStatus
	}{
		{
			名称:     "已打包状态",
			JSON数据: `"packed"`,
			期望状态:   TransactionStatus{Status: "packed"},
		},
		{
			名称:     "就绪状态",
			JSON数据: `"ready"`,
			期望状态:   TransactionStatus{Status: "ready"},
		},
		{
			名称:     "待处理状态-未来nonce",
			JSON数据: `{"pending":"futureNonce"}`,
			期望状态:   TransactionStatus{Status: "pending", PendingReason: "futureNonce"},
		},
		{
			名称:     "待处理状态-余额不足",
			JSON数据: `{"pending":"notEnoughCash"}`,
			期望状态:   TransactionStatus{Status: "pending", PendingReason: "notEnoughCash"},
		},
	}

	for _, 案例 := range 测试用例 {
		t.Run(案例.名称, func(t *testing.T) {
			var 结果 TransactionStatus
			错误 := json.Unmarshal([]byte(案例.JSON数据), &结果)
			if 错误 != nil {
				t.Fatalf("反序列化失败: %v", 错误)
			}
			if 结果 != 案例.期望状态 {
				t.Errorf("期望 %+v, 得到 %+v", 案例.期望状态, 结果)
			}
		})
	}
}
