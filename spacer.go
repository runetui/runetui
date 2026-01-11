package runetui

// Spacer creates a fixed-size spacer component.
// Returns an empty Box with both width and height set to the specified size.
// The layout engine will use the appropriate dimension based on parent direction.
func Spacer(size int) Component {
	return Box(BoxProps{
		Width:  DimensionFixed(size),
		Height: DimensionFixed(size),
	})
}

// FlexSpacer creates a flexible spacer that fills available space.
// Returns an empty Box with FlexGrow set to 1.0.
func FlexSpacer() Component {
	return Box(BoxProps{FlexGrow: 1.0})
}
