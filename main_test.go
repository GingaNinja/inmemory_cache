package main_test

import (
	main "nandos"
	"testing"
	"time"
)

func TestNew_cache(t *testing.T) {
	tests := []struct {
		desc string
		key  string // description of this test case
		val  string
	}{
		{desc: "chicken with valid ttl is returned", key: "item1", val: "chicken"},
		{desc: "sweetcorn with valid ttl is returned", key: "item2", val: "sweetcorn"},
	}
	cache := main.New_cache()
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			cache.Set(tt.key, tt.val, 60)
			actual := cache.Get(tt.key)
			if actual != tt.val {
				t.Errorf("didn't get expected value: actual = %v, want %v", actual, tt.val)
			}
		})
	}
}

func Test_missing_key(t *testing.T) {
	cache := main.New_cache()

	val := cache.Get("coleslaw")

	if val != nil {
		t.Errorf("coleslaw doesn't exist in the store. Should be nil, but got %v", val)
	}
}

func Test_expired_key(t *testing.T) {
	customTime := func() time.Time {
		return time.Now().Add(time.Duration(20) * time.Second)
	}

	cache := main.New_cache_with_time(customTime)

	cache.Set("expired_key", "chicken", 10)
	val := cache.Get("expired_key")

	if val != nil {
		t.Errorf("expired_key should be expired, but it is %v", val)
	}
}

func Test_delete_key(t *testing.T) {
	cache := main.New_cache()

	cache.Set("key_to_delete", "chicken", 10)
	cache.Delete("key_to_delete")
	val := cache.Get("key_to_delete")

	if val != nil {
		t.Errorf("key_to_delete should be deleted, but it is %v", val)
	}
}
