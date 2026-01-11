package runetui

// FlexChild represents a child component with flex properties.
type FlexChild struct {
	Component  Component
	Size       Size
	FlexGrow   float64
	FlexShrink float64
}

// calculateFlexGrow distributes extra space proportionally based on flex-grow values.
func calculateFlexGrow(children []FlexChild, extraSpace int) []int {
	result := make([]int, len(children))

	if extraSpace <= 0 {
		return result
	}

	var totalGrow float64
	for _, child := range children {
		totalGrow += child.FlexGrow
	}

	if totalGrow == 0 {
		return result
	}

	for i, child := range children {
		result[i] = int(float64(extraSpace) * child.FlexGrow / totalGrow)
	}

	return result
}

// calculateFlexShrink reduces size proportionally based on flex-shrink values when constrained.
func calculateFlexShrink(children []FlexChild, deficit int) []int {
	result := make([]int, len(children))

	if deficit <= 0 {
		return result
	}

	var totalShrink float64
	for _, child := range children {
		totalShrink += child.FlexShrink
	}

	if totalShrink == 0 {
		return result
	}

	for i, child := range children {
		result[i] = int(float64(deficit) * child.FlexShrink / totalShrink)
	}

	return result
}

// alignItems aligns children on the cross-axis based on AlignItems value.
func alignItems(children []*LayoutTree, props BoxProps, crossSize int) {
	for _, child := range children {
		if props.Direction == Column {
			switch props.AlignItems {
			case AlignStart:
			case AlignCenter:
				child.Layout.X = (crossSize - child.Layout.Width) / 2
			case AlignEnd:
				child.Layout.X = crossSize - child.Layout.Width
			case AlignStretch:
				child.Layout.Width = crossSize
			}
		} else {
			switch props.AlignItems {
			case AlignStart:
			case AlignCenter:
				child.Layout.Y = (crossSize - child.Layout.Height) / 2
			case AlignEnd:
				child.Layout.Y = crossSize - child.Layout.Height
			case AlignStretch:
				child.Layout.Height = crossSize
			}
		}
	}
}

// justifyContent distributes children on the main-axis based on JustifyContent value.
func justifyContent(children []*LayoutTree, props BoxProps, mainSize int) {
	if len(children) == 0 {
		return
	}

	if props.Direction == Column {
		justifyColumn(children, props, mainSize)
	} else {
		justifyRow(children, props, mainSize)
	}
}

func justifyColumn(children []*LayoutTree, props BoxProps, mainSize int) {
	switch props.JustifyContent {
	case JustifyStart:
	case JustifyCenter:
		totalHeight := getTotalHeight(children)
		offset := (mainSize - totalHeight) / 2
		for _, child := range children {
			child.Layout.Y += offset
		}
	case JustifyEnd:
		totalHeight := getTotalHeight(children)
		offset := mainSize - totalHeight
		for _, child := range children {
			child.Layout.Y += offset
		}
	case JustifySpaceBetween:
		if len(children) <= 1 {
			return
		}
		totalHeight := getTotalHeight(children)
		space := (mainSize - totalHeight) / (len(children) - 1)
		for i := 1; i < len(children); i++ {
			children[i].Layout.Y = children[i-1].Layout.Y + children[i-1].Layout.Height + space
		}
	case JustifySpaceAround:
		totalHeight := getTotalHeight(children)
		space := (mainSize - totalHeight) / len(children)
		halfSpace := space / 2
		for i, child := range children {
			child.Layout.Y = halfSpace + i*(child.Layout.Height+space)
		}
	}
}

func justifyRow(children []*LayoutTree, props BoxProps, mainSize int) {
	switch props.JustifyContent {
	case JustifyStart:
	case JustifyCenter:
		totalWidth := getTotalWidth(children)
		offset := (mainSize - totalWidth) / 2
		for _, child := range children {
			child.Layout.X += offset
		}
	case JustifyEnd:
		totalWidth := getTotalWidth(children)
		offset := mainSize - totalWidth
		for _, child := range children {
			child.Layout.X += offset
		}
	case JustifySpaceBetween:
		if len(children) <= 1 {
			return
		}
		totalWidth := getTotalWidth(children)
		space := (mainSize - totalWidth) / (len(children) - 1)
		for i := 1; i < len(children); i++ {
			children[i].Layout.X = children[i-1].Layout.X + children[i-1].Layout.Width + space
		}
	case JustifySpaceAround:
		totalWidth := getTotalWidth(children)
		space := (mainSize - totalWidth) / len(children)
		halfSpace := space / 2
		for i, child := range children {
			child.Layout.X = halfSpace + i*(child.Layout.Width+space)
		}
	}
}

func getTotalHeight(children []*LayoutTree) int {
	if len(children) == 0 {
		return 0
	}
	first := children[0].Layout.Y
	last := children[len(children)-1]
	return last.Layout.Y + last.Layout.Height - first
}

func getTotalWidth(children []*LayoutTree) int {
	if len(children) == 0 {
		return 0
	}
	first := children[0].Layout.X
	last := children[len(children)-1]
	return last.Layout.X + last.Layout.Width - first
}
