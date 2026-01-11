package runetui

import "strings"

var currentStaticManager *StaticManager

// SetStaticManager sets the current static manager for rendering.
func SetStaticManager(sm *StaticManager) {
	currentStaticManager = sm
}

// StaticProps defines properties for Static component.
type StaticProps struct {
	Key string
}

func (StaticProps) isProps() {}

// Static creates a static component that accumulates content across renders.
func Static(props StaticProps, itemsFunc func() []Component) Component {
	return &static{
		props:     props,
		itemsFunc: itemsFunc,
	}
}

type static struct {
	props     StaticProps
	itemsFunc func() []Component
}

func (s *static) Render(layout Layout) string {
	items := s.itemsFunc()
	lines := []string{}
	for _, item := range items {
		rendered := item.Render(layout)
		lines = append(lines, rendered)
	}

	if currentStaticManager != nil {
		count := currentStaticManager.AppendStatic(s.props.Key, lines)
		if count == 0 {
			return ""
		}
		if count < len(lines) {
			return strings.Join(lines[len(lines)-count:], "\n")
		}
	}

	return strings.Join(lines, "\n")
}

func (s *static) Children() []Component {
	return []Component{}
}

func (s *static) Key() string {
	return s.props.Key
}

func (s *static) Measure(availableWidth, availableHeight int) Size {
	items := s.itemsFunc()
	totalHeight := 0
	maxWidth := 0

	for _, item := range items {
		size := item.Measure(availableWidth, availableHeight)
		totalHeight += size.Height
		if size.Width > maxWidth {
			maxWidth = size.Width
		}
	}

	return Size{
		Width:  maxWidth,
		Height: totalHeight,
	}
}
