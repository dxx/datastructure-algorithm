package greedy

import (
    "fmt"
    "testing"
)

func TestGetBroadcast(t *testing.T) {
    broadcastMap := map[string][]string {
        "B1": {"北京", "上海", "天津"},
        "B2": {"广州", "北京", "深圳"},
        "B3": {"成都", "上海", "杭州"},
        "B4": {"上海", "天津"},
        "B5": {"杭州", "大连"},
    }
    broadcasts := getBroadcast(broadcastMap)
    fmt.Printf("最少选择的广播电台: %v\n", broadcasts)
}
