package singleton

import (
	"testing"
)

func TestGetSingletonInstance(t *testing.T) {
	if GetSingletonInstance() != GetSingletonInstance() {
		t.Errorf("Test Failed: Instances returned are not equal")
	}
}

func BenchmarkGetSingletonInstance(b *testing.B) {
	instance := GetSingletonInstance()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if instance != GetSingletonInstance() {
				b.Errorf("Test Failed: Instances returned are not equal")
			}
		}
	})
}
