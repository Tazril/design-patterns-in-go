package chain

import (
	"fmt"
	"testing"
)

func TestErrorLogger_Log(t *testing.T) {
	// given
	il := InfoLogger{}
	wl := WarnLogger{&il}
	el := ErrorLogger{&wl}
	message := "error message"
	// when
	result := el.Log("error", message)
	// then
	expected := fmt.Sprintf("Error: %v\n", message)
	if result != expected {
		t.Errorf("Got result= %v Expected %v", result, expected)
	}
}

func TestWarnLogger_Log(t *testing.T) {
	// given
	il := InfoLogger{}
	wl := WarnLogger{&il}
	el := ErrorLogger{&wl}
	message := "warning message"
	// when
	result := el.Log("warn", message)
	// then
	expected := fmt.Sprintf("Warn: %v\n", message)
	if result != expected {
		t.Errorf("Got result= %v Expected %v", result, expected)
	}
}

func TestInfoLogger_Log(t *testing.T) {
	// given
	il := InfoLogger{}
	wl := WarnLogger{&il}
	el := ErrorLogger{&wl}
	message := "info message"
	// when
	result := el.Log("info", message)
	// then
	expected := fmt.Sprintf("Info: %v\n", message)
	if result != expected {
		t.Errorf("Got result= %v Expected %v", result, expected)
	}
}
