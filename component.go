package runetui

// Layout defines the position and size of a component after layout calculation.
type Layout struct {
	X      int
	Y      int
	Width  int
	Height int
}

// Size defines size constraints returned by component's Measure method.
type Size struct {
	Width     int
	Height    int
	MinWidth  int
	MinHeight int
	MaxWidth  int
	MaxHeight int
}

// Props is a marker interface for component properties.
type Props interface {
	isProps()
}

// Component represents a UI component that can be rendered.
type Component interface {
	// Render generates the string representation of the component within the given layout.
	Render(layout Layout) string

	// Children returns the child components.
	Children() []Component

	// Key returns a unique identifier for this component instance.
	Key() string

	// Measure calculates the size requirements for this component.
	Measure(availableWidth, availableHeight int) Size
}

// ComponentFunc is a function that returns a Component, allowing functional component definitions.
type ComponentFunc func() Component

// Render delegates to the component returned by the function.
func (f ComponentFunc) Render(layout Layout) string {
	return f().Render(layout)
}

// Children delegates to the component returned by the function.
func (f ComponentFunc) Children() []Component {
	return f().Children()
}

// Key delegates to the component returned by the function.
func (f ComponentFunc) Key() string {
	return f().Key()
}

// Measure delegates to the component returned by the function.
func (f ComponentFunc) Measure(availableWidth, availableHeight int) Size {
	return f().Measure(availableWidth, availableHeight)
}
