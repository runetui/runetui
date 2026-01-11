package runetui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

func init() {
	lipgloss.SetColorProfile(termenv.TrueColor)
}

// TextProps defines properties for the Text component.
type TextProps struct {
	Content       string
	Color         string
	Background    string
	Bold          bool
	Italic        bool
	Underline     bool
	Strikethrough bool
	Wrap          WrapMode
	Align         TextAlign
	Key           string
}

func (TextProps) isProps() {}

type text struct {
	content string
	props   TextProps
}

// Text creates a new text component with the given content and optional properties.
func Text(content string, props ...TextProps) Component {
	p := TextProps{Content: content}
	if len(props) > 0 {
		p = props[0]
		p.Content = content
	}
	return &text{
		content: content,
		props:   p,
	}
}

func (t *text) Render(layout Layout) string {
	style := lipgloss.NewStyle()

	if t.props.Color != "" {
		style = style.Foreground(lipgloss.Color(t.props.Color))
	}

	if t.props.Background != "" {
		style = style.Background(lipgloss.Color(t.props.Background))
	}

	if t.props.Bold {
		style = style.Bold(true)
	}

	if t.props.Italic {
		style = style.Italic(true)
	}

	if t.props.Underline {
		style = style.Underline(true)
	}

	if t.props.Strikethrough {
		style = style.Strikethrough(true)
	}

	style = style.Width(layout.Width)

	switch t.props.Wrap {
	case WrapWord:
		style = style.MaxWidth(layout.Width)
	case WrapTruncate:
		style = style.MaxWidth(layout.Width).Inline(true)
	}

	switch t.props.Align {
	case TextAlignLeft:
		style = style.Align(lipgloss.Left)
	case TextAlignCenter:
		style = style.Align(lipgloss.Center)
	case TextAlignRight:
		style = style.Align(lipgloss.Right)
	}

	return style.Render(t.content)
}

func (t *text) Children() []Component {
	return []Component{}
}

func (t *text) Key() string {
	return t.props.Key
}

func (t *text) Measure(availableWidth, availableHeight int) Size {
	lines := 1
	width := len(t.content)

	if t.props.Wrap == WrapWord && width > availableWidth {
		width = availableWidth
		lines = (len(t.content) + availableWidth - 1) / availableWidth
	}

	if t.props.Wrap == WrapTruncate && width > availableWidth {
		width = availableWidth
		lines = 1
	}

	return Size{
		Width:  width,
		Height: lines,
	}
}
