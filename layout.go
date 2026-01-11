package runetui

// LayoutEngine calculates positions for components based on terminal dimensions.
type LayoutEngine struct {
	terminalWidth  int
	terminalHeight int
}

// NewLayoutEngine creates a new layout engine with the given terminal dimensions.
func NewLayoutEngine(width, height int) *LayoutEngine {
	return &LayoutEngine{
		terminalWidth:  width,
		terminalHeight: height,
	}
}

// LayoutTree represents a component and its calculated layout along with its children.
type LayoutTree struct {
	Component Component
	Layout    Layout
	Children  []*LayoutTree
}

// CalculateLayout is the main entry point for layout calculation.
func (e *LayoutEngine) CalculateLayout(root Component) *LayoutTree {
	return e.measureAndLayout(root, e.terminalWidth, e.terminalHeight, 0, 0)
}

// measureAndLayout recursively measures and positions components.
func (e *LayoutEngine) measureAndLayout(component Component, availableWidth, availableHeight, x, y int) *LayoutTree {
	marginLeft := 0
	marginTop := 0

	if b, ok := component.(*box); ok {
		marginLeft = b.props.Margin.Left
		marginTop = b.props.Margin.Top
	}

	adjustedX := x + marginLeft
	adjustedY := y + marginTop

	size := component.Measure(availableWidth, availableHeight)

	layout := Layout{
		X:      adjustedX,
		Y:      adjustedY,
		Width:  size.Width,
		Height: size.Height,
	}

	children := component.Children()
	childTrees := make([]*LayoutTree, 0, len(children))

	if len(children) > 0 {
		if b, ok := component.(*box); ok {
			paddingLeft := b.props.Padding.Left
			paddingTop := b.props.Padding.Top

			borderWidth, borderHeight := borderSize(b.props.Border)
			borderLeft := borderWidth / 2
			borderTop := borderHeight / 2

			switch b.props.Direction {
			case Column:
				currentY := adjustedY + paddingTop + borderTop
				for i, child := range children {
					childTree := e.measureAndLayout(child, availableWidth, availableHeight, adjustedX+paddingLeft+borderLeft, currentY)
					childTrees = append(childTrees, childTree)
					currentY += childTree.Layout.Height
					if i < len(children)-1 && b.props.Gap > 0 {
						currentY += b.props.Gap
					}
				}
			case Row:
				currentX := adjustedX + paddingLeft + borderLeft
				for i, child := range children {
					childTree := e.measureAndLayout(child, availableWidth, availableHeight, currentX, adjustedY+paddingTop+borderTop)
					childTrees = append(childTrees, childTree)
					currentX += childTree.Layout.Width
					if i < len(children)-1 && b.props.Gap > 0 {
						currentX += b.props.Gap
					}
				}
			}
		}
	}

	return &LayoutTree{
		Component: component,
		Layout:    layout,
		Children:  childTrees,
	}
}
