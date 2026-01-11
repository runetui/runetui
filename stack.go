package runetui

// StackProps defines simplified properties for stack components.
type StackProps struct {
	Gap            int
	Padding        Spacing
	AlignItems     Align
	JustifyContent Justify
	Width          Dimension
	Height         Dimension
	Key            string
}

// VStack creates a vertical stack with default properties.
func VStack(children ...Component) Component {
	props := BoxProps{
		Direction: Column,
	}
	return Box(props, children...)
}

// VStackWithProps creates a vertical stack with custom properties.
func VStackWithProps(props StackProps, children ...Component) Component {
	boxProps := BoxProps{
		Direction:      Column,
		Gap:            props.Gap,
		Padding:        props.Padding,
		AlignItems:     props.AlignItems,
		JustifyContent: props.JustifyContent,
		Width:          props.Width,
		Height:         props.Height,
		Key:            props.Key,
	}
	return Box(boxProps, children...)
}

// HStack creates a horizontal stack with default properties.
func HStack(children ...Component) Component {
	props := BoxProps{
		Direction: Row,
	}
	return Box(props, children...)
}

// HStackWithProps creates a horizontal stack with custom properties.
func HStackWithProps(props StackProps, children ...Component) Component {
	boxProps := BoxProps{
		Direction:      Row,
		Gap:            props.Gap,
		Padding:        props.Padding,
		AlignItems:     props.AlignItems,
		JustifyContent: props.JustifyContent,
		Width:          props.Width,
		Height:         props.Height,
		Key:            props.Key,
	}
	return Box(boxProps, children...)
}
