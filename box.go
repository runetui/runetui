package runetui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// BoxProps defines the properties for a Box component.
type BoxProps struct {
	Direction      Direction
	Width          Dimension
	Height         Dimension
	MinWidth       int
	MinHeight      int
	MaxWidth       int
	MaxHeight      int
	FlexGrow       float64
	FlexShrink     float64
	AlignItems     Align
	JustifyContent Justify
	Padding        Spacing
	Margin         Spacing
	Gap            int
	Border         BorderStyle
	BorderColor    string
	Background     string
	IsStatic       bool
	Key            string
}

func (BoxProps) isProps() {}

// box is the private implementation of the Box component.
type box struct {
	props    BoxProps
	children []Component
}

// Box creates a new Box component with the given properties and children.
func Box(props BoxProps, children ...Component) Component {
	if children == nil {
		children = []Component{}
	}
	return &box{
		props:    props,
		children: children,
	}
}

// Render generates the string representation of the box.
func (b *box) Render(layout Layout) string {
	if len(b.children) == 0 {
		return ""
	}

	var parts []string
	for _, child := range b.children {
		childLayout := Layout{
			X:      layout.X,
			Y:      layout.Y,
			Width:  layout.Width,
			Height: layout.Height,
		}
		parts = append(parts, child.Render(childLayout))
	}

	var content string
	if b.props.Direction == Row {
		content = strings.Join(parts, "")
	} else {
		content = strings.Join(parts, "\n")
	}

	style := lipgloss.NewStyle()

	if b.props.Border != BorderNone {
		style = b.applyBorder(style)
	}

	if b.props.Background != "" {
		style = style.Background(lipgloss.Color(b.props.Background))
	}

	return style.Render(content)
}

func (b *box) applyBorder(style lipgloss.Style) lipgloss.Style {
	switch b.props.Border {
	case BorderSingle:
		style = style.Border(lipgloss.NormalBorder())
	case BorderDouble:
		style = style.Border(lipgloss.DoubleBorder())
	case BorderRounded:
		style = style.Border(lipgloss.RoundedBorder())
	}

	if b.props.BorderColor != "" {
		style = style.BorderForeground(lipgloss.Color(b.props.BorderColor))
	}

	return style
}

// Children returns the child components.
func (b *box) Children() []Component {
	return b.children
}

// Key returns the unique identifier for this component.
func (b *box) Key() string {
	return b.props.Key
}

// Measure calculates the size requirements for this component.
func (b *box) Measure(availableWidth, availableHeight int) Size {
	return measureBox(b.props, b.children, availableWidth, availableHeight)
}
