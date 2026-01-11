package runetui

import "testing"

func TestDirection_Column_IsZero(t *testing.T) {
	if Column != 0 {
		t.Errorf("Column should be 0, got %d", Column)
	}
}

func TestDirection_Row_IsOne(t *testing.T) {
	if Row != 1 {
		t.Errorf("Row should be 1, got %d", Row)
	}
}

func TestDimensionAuto_CanBeCreated(t *testing.T) {
	dim := DimensionAuto()
	if dim == nil {
		t.Error("DimensionAuto should not be nil")
	}
}

func TestDimensionFixed_StoresValue(t *testing.T) {
	dim := DimensionFixed(100)
	fixed, ok := dim.(interface{ Value() int })
	if !ok {
		t.Fatal("DimensionFixed should expose Value() method")
	}
	if got := fixed.Value(); got != 100 {
		t.Errorf("expected 100, got %d", got)
	}
}

func TestDimensionPercent_StoresValue(t *testing.T) {
	dim := DimensionPercent(50)
	percent, ok := dim.(interface{ Value() int })
	if !ok {
		t.Fatal("DimensionPercent should expose Value() method")
	}
	if got := percent.Value(); got != 50 {
		t.Errorf("expected 50, got %d", got)
	}
}

func TestSpacing_WithValues_CreatesCorrectly(t *testing.T) {
	spacing := Spacing{Top: 1, Right: 2, Bottom: 3, Left: 4}
	if spacing.Top != 1 || spacing.Right != 2 || spacing.Bottom != 3 || spacing.Left != 4 {
		t.Error("Spacing values should match constructor")
	}
}

func TestSpacingAll_SetsAllSides(t *testing.T) {
	spacing := SpacingAll(5)
	if spacing.Top != 5 || spacing.Right != 5 || spacing.Bottom != 5 || spacing.Left != 5 {
		t.Error("SpacingAll should set all sides to the same value")
	}
}

func TestSpacingVertical_SetsTopAndBottom(t *testing.T) {
	spacing := SpacingVertical(3)
	if spacing.Top != 3 || spacing.Bottom != 3 || spacing.Left != 0 || spacing.Right != 0 {
		t.Error("SpacingVertical should set top and bottom, leave sides at 0")
	}
}

func TestSpacingHorizontal_SetsLeftAndRight(t *testing.T) {
	spacing := SpacingHorizontal(2)
	if spacing.Left != 2 || spacing.Right != 2 || spacing.Top != 0 || spacing.Bottom != 0 {
		t.Error("SpacingHorizontal should set left and right, leave top/bottom at 0")
	}
}

func TestSpacing_ZeroValue_CreatesZeroSpacing(t *testing.T) {
	spacing := Spacing{}
	if spacing.Top != 0 || spacing.Right != 0 || spacing.Bottom != 0 || spacing.Left != 0 {
		t.Error("Zero value Spacing should have all fields as 0")
	}
}

func TestSpacingAll_WithZero_CreatesZeroSpacing(t *testing.T) {
	spacing := SpacingAll(0)
	if spacing.Top != 0 || spacing.Right != 0 || spacing.Bottom != 0 || spacing.Left != 0 {
		t.Error("SpacingAll(0) should create zero spacing")
	}
}

func TestSpacing_WithNegativeValues_AllowsNegative(t *testing.T) {
	spacing := Spacing{Top: -1, Right: -2, Bottom: -3, Left: -4}
	if spacing.Top != -1 || spacing.Right != -2 || spacing.Bottom != -3 || spacing.Left != -4 {
		t.Error("Spacing should allow negative values")
	}
}

func TestBorderStyle_BorderNone_IsZero(t *testing.T) {
	if BorderNone != 0 {
		t.Errorf("BorderNone should be 0, got %d", BorderNone)
	}
}

func TestBorderStyle_BorderSingle_IsOne(t *testing.T) {
	if BorderSingle != 1 {
		t.Errorf("BorderSingle should be 1, got %d", BorderSingle)
	}
}

func TestBorderStyle_BorderDouble_IsTwo(t *testing.T) {
	if BorderDouble != 2 {
		t.Errorf("BorderDouble should be 2, got %d", BorderDouble)
	}
}

func TestBorderStyle_BorderRounded_IsThree(t *testing.T) {
	if BorderRounded != 3 {
		t.Errorf("BorderRounded should be 3, got %d", BorderRounded)
	}
}

func TestAlign_AlignStart_IsZero(t *testing.T) {
	if AlignStart != 0 {
		t.Errorf("AlignStart should be 0, got %d", AlignStart)
	}
}

func TestAlign_AlignCenter_IsOne(t *testing.T) {
	if AlignCenter != 1 {
		t.Errorf("AlignCenter should be 1, got %d", AlignCenter)
	}
}

func TestAlign_AlignEnd_IsTwo(t *testing.T) {
	if AlignEnd != 2 {
		t.Errorf("AlignEnd should be 2, got %d", AlignEnd)
	}
}

func TestAlign_AlignStretch_IsThree(t *testing.T) {
	if AlignStretch != 3 {
		t.Errorf("AlignStretch should be 3, got %d", AlignStretch)
	}
}

func TestJustify_JustifyStart_IsZero(t *testing.T) {
	if JustifyStart != 0 {
		t.Errorf("JustifyStart should be 0, got %d", JustifyStart)
	}
}

func TestJustify_JustifyCenter_IsOne(t *testing.T) {
	if JustifyCenter != 1 {
		t.Errorf("JustifyCenter should be 1, got %d", JustifyCenter)
	}
}

func TestJustify_JustifyEnd_IsTwo(t *testing.T) {
	if JustifyEnd != 2 {
		t.Errorf("JustifyEnd should be 2, got %d", JustifyEnd)
	}
}

func TestJustify_JustifySpaceBetween_IsThree(t *testing.T) {
	if JustifySpaceBetween != 3 {
		t.Errorf("JustifySpaceBetween should be 3, got %d", JustifySpaceBetween)
	}
}

func TestJustify_JustifySpaceAround_IsFour(t *testing.T) {
	if JustifySpaceAround != 4 {
		t.Errorf("JustifySpaceAround should be 4, got %d", JustifySpaceAround)
	}
}

func TestWrapMode_WrapNone_IsZero(t *testing.T) {
	if WrapNone != 0 {
		t.Errorf("WrapNone should be 0, got %d", WrapNone)
	}
}

func TestWrapMode_WrapWord_IsOne(t *testing.T) {
	if WrapWord != 1 {
		t.Errorf("WrapWord should be 1, got %d", WrapWord)
	}
}

func TestWrapMode_WrapChar_IsTwo(t *testing.T) {
	if WrapChar != 2 {
		t.Errorf("WrapChar should be 2, got %d", WrapChar)
	}
}

func TestWrapMode_WrapTruncate_IsThree(t *testing.T) {
	if WrapTruncate != 3 {
		t.Errorf("WrapTruncate should be 3, got %d", WrapTruncate)
	}
}

func TestTextAlign_TextAlignLeft_IsZero(t *testing.T) {
	if TextAlignLeft != 0 {
		t.Errorf("TextAlignLeft should be 0, got %d", TextAlignLeft)
	}
}

func TestTextAlign_TextAlignCenter_IsOne(t *testing.T) {
	if TextAlignCenter != 1 {
		t.Errorf("TextAlignCenter should be 1, got %d", TextAlignCenter)
	}
}

func TestTextAlign_TextAlignRight_IsTwo(t *testing.T) {
	if TextAlignRight != 2 {
		t.Errorf("TextAlignRight should be 2, got %d", TextAlignRight)
	}
}
