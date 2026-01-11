package runetui

import (
	"strings"
	"unicode/utf8"
)

// resolveDimension resolves a Dimension to a concrete integer value.
// For Fixed dimensions, returns the fixed value.
// For Auto dimensions, returns 0 (caller must provide intrinsic size).
// For Percent dimensions, calculates percentage of available space.
func resolveDimension(dim Dimension, available int) int {
	switch d := dim.(type) {
	case dimensionFixed:
		return d.Value()
	case dimensionPercent:
		return (available * d.Value()) / 100
	case dimensionAuto:
		return 0
	default:
		return 0
	}
}

// measureText calculates the size of text based on content and wrap mode.
func measureText(content string, wrap WrapMode, availableWidth int) Size {
	if content == "" {
		return Size{Width: 0, Height: 0}
	}

	lines := strings.Split(content, "\n")
	height := len(lines)
	width := 0

	for _, line := range lines {
		lineWidth := utf8.RuneCountInString(line)
		if lineWidth > width {
			width = lineWidth
		}
	}

	if wrap == WrapNone {
		return Size{Width: width, Height: height}
	}

	if wrap == WrapTruncate && width > availableWidth {
		return Size{Width: availableWidth, Height: height}
	}

	if wrap == WrapWord || wrap == WrapChar {
		if width > availableWidth && availableWidth > 0 {
			totalRunes := 0
			for _, line := range lines {
				totalRunes += utf8.RuneCountInString(line)
			}
			wrappedHeight := (totalRunes + availableWidth - 1) / availableWidth
			return Size{Width: availableWidth, Height: wrappedHeight}
		}
	}

	return Size{Width: width, Height: height}
}

// spacingWidth returns the total horizontal spacing (left + right).
func spacingWidth(s Spacing) int {
	return s.Left + s.Right
}

// spacingHeight returns the total vertical spacing (top + bottom).
func spacingHeight(s Spacing) int {
	return s.Top + s.Bottom
}

// borderSize returns the width and height added by a border.
// Any border style except BorderNone adds 2 to width and 2 to height.
func borderSize(style BorderStyle) (width, height int) {
	if style == BorderNone {
		return 0, 0
	}
	return 2, 2
}

// applyConstraints applies min/max constraints to a size.
func applyConstraints(size Size, minWidth, minHeight, maxWidth, maxHeight int) Size {
	if minWidth > 0 && size.Width < minWidth {
		size.Width = minWidth
	}
	if maxWidth > 0 && size.Width > maxWidth {
		size.Width = maxWidth
	}
	if minHeight > 0 && size.Height < minHeight {
		size.Height = minHeight
	}
	if maxHeight > 0 && size.Height > maxHeight {
		size.Height = maxHeight
	}
	return size
}

// measureBox calculates the size of a box including its children.
func measureBox(props BoxProps, children []Component, availableWidth, availableHeight int) Size {
	if len(children) == 0 {
		return Size{Width: 0, Height: 0}
	}

	var totalWidth, totalHeight int
	var maxWidth, maxHeight int

	for i, child := range children {
		childSize := child.Measure(availableWidth, availableHeight)

		if props.Direction == Row {
			totalWidth += childSize.Width
			if i > 0 && props.Gap > 0 {
				totalWidth += props.Gap
			}
			if childSize.Height > maxHeight {
				maxHeight = childSize.Height
			}
		} else {
			totalHeight += childSize.Height
			if i > 0 && props.Gap > 0 {
				totalHeight += props.Gap
			}
			if childSize.Width > maxWidth {
				maxWidth = childSize.Width
			}
		}
	}

	var width, height int
	if props.Direction == Row {
		width = totalWidth
		height = maxHeight
	} else {
		width = maxWidth
		height = totalHeight
	}

	width += spacingWidth(props.Padding)
	height += spacingHeight(props.Padding)

	width += spacingWidth(props.Margin)
	height += spacingHeight(props.Margin)

	borderWidth, borderHeight := borderSize(props.Border)
	width += borderWidth
	height += borderHeight

	resolvedWidth := resolveDimension(props.Width, availableWidth)
	if resolvedWidth > 0 {
		width = resolvedWidth
	}

	resolvedHeight := resolveDimension(props.Height, availableHeight)
	if resolvedHeight > 0 {
		height = resolvedHeight
	}

	size := Size{Width: width, Height: height}
	size = applyConstraints(size, props.MinWidth, props.MinHeight, props.MaxWidth, props.MaxHeight)

	return size
}
