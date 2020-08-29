package test

import (
	"github.com/patrickmn/go-cache"
	"testing"
	"time"
	"ukiyo/internal/caching/cacheupdate"
)

const (
	_TimeZone = "Asia/Kolkata"
)

func TestHistoryFetch(t *testing.T) {
	c := cache.New(1*time.Minute, 1*time.Minute)
	var history cacheupdate.History
	var historyArr []cacheupdate.History
	loc, _ := time.LoadLocation(_TimeZone)
	history.Action = ""
	history.Comment = ""
	history.Status = 0
	history.Time = time.Now().In(loc).Format("2006-01-02 15:04:05")
	historyArr = append(historyArr, history)
	c.Set("history", historyArr, 2*time.Minute)
	val := cacheupdate.HistoryFetch(c)
	if len(val) > 0 {
		t.Logf("passed expected  %v and got value %v", len(val), len(val))
	} else {
		t.Errorf("failed expected %v but got value %v", 0, len(val))
	}
}