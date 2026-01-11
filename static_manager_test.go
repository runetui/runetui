package runetui

import "testing"

func TestNewStaticManager_ReturnsNonNil(t *testing.T) {
	sm := NewStaticManager()
	if sm == nil {
		t.Error("NewStaticManager should return non-nil")
	}
}

func TestAppendStatic_WithContent_ReturnsCount(t *testing.T) {
	sm := NewStaticManager()
	content := []string{"line1", "line2", "line3"}
	count := sm.AppendStatic("key1", content)
	if count != 3 {
		t.Errorf("expected count 3, got %d", count)
	}
}

func TestRenderStatic_WithContent_ReturnsAccumulatedLines(t *testing.T) {
	sm := NewStaticManager()
	sm.AppendStatic("key1", []string{"line1", "line2"})
	result := sm.RenderStatic()
	expected := "line1\nline2"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestAppendStatic_WithSameKey_ReturnsZeroOnSecondCall(t *testing.T) {
	sm := NewStaticManager()
	count1 := sm.AppendStatic("key1", []string{"line1", "line2"})
	count2 := sm.AppendStatic("key1", []string{"line3", "line4"})
	if count1 != 2 {
		t.Errorf("first call: expected count 2, got %d", count1)
	}
	if count2 != 0 {
		t.Errorf("second call: expected count 0, got %d", count2)
	}
}

func TestAppendStatic_WithSameKey_DoesNotDuplicateContent(t *testing.T) {
	sm := NewStaticManager()
	sm.AppendStatic("key1", []string{"line1", "line2"})
	sm.AppendStatic("key1", []string{"line3", "line4"})
	result := sm.RenderStatic()
	expected := "line1\nline2"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestClear_AfterAppending_ClearsBuffer(t *testing.T) {
	sm := NewStaticManager()
	sm.AppendStatic("key1", []string{"line1", "line2"})
	sm.Clear()
	result := sm.RenderStatic()
	if result != "" {
		t.Errorf("expected empty string, got %q", result)
	}
}

func TestClear_AfterAppending_AllowsReuse(t *testing.T) {
	sm := NewStaticManager()
	sm.AppendStatic("key1", []string{"line1", "line2"})
	sm.Clear()
	count := sm.AppendStatic("key1", []string{"line3", "line4"})
	if count != 2 {
		t.Errorf("expected count 2, got %d", count)
	}
	result := sm.RenderStatic()
	expected := "line3\nline4"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestAppendStatic_WithEmptyContent_ReturnsZero(t *testing.T) {
	sm := NewStaticManager()
	count := sm.AppendStatic("key1", []string{})
	if count != 0 {
		t.Errorf("expected count 0, got %d", count)
	}
}

func TestRenderStatic_WithEmptyBuffer_ReturnsEmpty(t *testing.T) {
	sm := NewStaticManager()
	result := sm.RenderStatic()
	if result != "" {
		t.Errorf("expected empty string, got %q", result)
	}
}

func TestAppendStatic_WithMultipleKeys_AccumulatesAll(t *testing.T) {
	sm := NewStaticManager()
	sm.AppendStatic("key1", []string{"line1", "line2"})
	sm.AppendStatic("key2", []string{"line3", "line4"})
	result := sm.RenderStatic()
	expected := "line1\nline2\nline3\nline4"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestRenderStatic_WithSingleLine_ReturnsWithoutNewline(t *testing.T) {
	sm := NewStaticManager()
	sm.AppendStatic("key1", []string{"line1"})
	result := sm.RenderStatic()
	expected := "line1"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
