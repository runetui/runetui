package runetui

// Direction defines the layout direction for flex containers.
type Direction int

const (
	// Column arranges children vertically (top to bottom).
	Column Direction = iota
	// Row arranges children horizontally (left to right).
	Row
)

// Dimension represents a sizing constraint (auto, fixed, or percentage).
type Dimension interface {
	isDimension()
}

// dimensionAuto represents automatic sizing based on content.
type dimensionAuto struct{}

func (dimensionAuto) isDimension() {}

// DimensionAuto creates an automatic dimension.
func DimensionAuto() Dimension {
	return dimensionAuto{}
}

// dimensionFixed represents a fixed size in cells.
type dimensionFixed struct {
	value int
}

func (dimensionFixed) isDimension() {}

// Value returns the fixed size value.
func (d dimensionFixed) Value() int {
	return d.value
}

// DimensionFixed creates a fixed dimension with the given size.
func DimensionFixed(value int) Dimension {
	return dimensionFixed{value: value}
}

// dimensionPercent represents a percentage of available space.
type dimensionPercent struct {
	value int
}

func (dimensionPercent) isDimension() {}

// Value returns the percentage value.
func (d dimensionPercent) Value() int {
	return d.value
}

// DimensionPercent creates a percentage dimension (0-100).
func DimensionPercent(value int) Dimension {
	return dimensionPercent{value: value}
}

// Spacing defines space around an element (like CSS padding/margin).
type Spacing struct {
	Top    int
	Right  int
	Bottom int
	Left   int
}

// SpacingAll creates spacing with the same value on all sides.
func SpacingAll(value int) Spacing {
	return Spacing{Top: value, Right: value, Bottom: value, Left: value}
}

// SpacingVertical creates spacing with values on top and bottom only.
func SpacingVertical(value int) Spacing {
	return Spacing{Top: value, Bottom: value}
}

// SpacingHorizontal creates spacing with values on left and right only.
func SpacingHorizontal(value int) Spacing {
	return Spacing{Left: value, Right: value}
}

// BorderStyle defines the border rendering style.
type BorderStyle int

const (
	// BorderNone renders no border.
	BorderNone BorderStyle = iota
	// BorderSingle renders a single-line border.
	BorderSingle
	// BorderDouble renders a double-line border.
	BorderDouble
	// BorderRounded renders a rounded border.
	BorderRounded
)

// Align defines cross-axis alignment in flex containers.
type Align int

const (
	// AlignStart aligns items to the start of the cross axis.
	AlignStart Align = iota
	// AlignCenter centers items on the cross axis.
	AlignCenter
	// AlignEnd aligns items to the end of the cross axis.
	AlignEnd
	// AlignStretch stretches items to fill the cross axis.
	AlignStretch
)

// Justify defines main-axis alignment in flex containers.
type Justify int

const (
	// JustifyStart packs items at the start.
	JustifyStart Justify = iota
	// JustifyCenter centers items.
	JustifyCenter
	// JustifyEnd packs items at the end.
	JustifyEnd
	// JustifySpaceBetween distributes items with space between them.
	JustifySpaceBetween
	// JustifySpaceAround distributes items with space around them.
	JustifySpaceAround
)

// WrapMode defines how text wraps or truncates.
type WrapMode int

const (
	// WrapNone disables text wrapping (overflow hidden).
	WrapNone WrapMode = iota
	// WrapWord wraps text at word boundaries.
	WrapWord
	// WrapChar wraps text at any character.
	WrapChar
	// WrapTruncate truncates text with ellipsis.
	WrapTruncate
)

// TextAlign defines horizontal text alignment.
type TextAlign int

const (
	// TextAlignLeft aligns text to the left.
	TextAlignLeft TextAlign = iota
	// TextAlignCenter centers text horizontally.
	TextAlignCenter
	// TextAlignRight aligns text to the right.
	TextAlignRight
)
