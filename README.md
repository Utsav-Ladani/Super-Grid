## Super Grid

This is a widget for fyne developer that allows to create a grid of widgets with a lot of options. You can feel like CSS flexbox. 

### Installation

- This project is in development, so you can't install it yet using `go get .` command.
- But you can clone this repository and use it in your project.


### Usage

- Wrapper you widgets or canvas object in `NewSuperGridElement` function or directly define `SuperGridElement` struct.
- Set necessary options for this SuperGridElements.
- Also create a variable for `SuperGridOptions` with required values.
- Create a `SuperGrid` object with `NewSuperGrid` function.
- Pass your `SuperGridOptions` and `SuperGridElements` to `SuperGrid` object.
- Now `SuperGrid` object is ready to use. You can use it as a widget or as a canvas object.
- To see the example of usage, please, look at `main.go` file.

### Options

#### SuperGridOptions

- `Direction` - direction of the grid. Can be `horizontal` or `vertical`.
    - Value: `DirectionHorizontal` or `DirectionVertical`
- `Spacing` - gap between elements.
    - Value: `float32`

#### SuperGridElementOptions

- `IsBlock` - if true, element will take all available space in the direction of layout. After placing non-block elements, block elements will take all available space equally.
    - Value: `bool`
- `Alignment` - alignment of the element in the grid. Can be `start`, `center`, `end`.
    - Value: `AlignmentStart`, `AlignmentCenter`, `AlignmentEnd`
- `Fill` - if true, element will take all available space in the perpendicular direction of layout.
    - Value: `bool`
- `Margin` - margin of the element.
    - Value: `[4]float32` - top, right, bottom, left

### Note

- You can also add other properties to `SuperGridElementOptions` and `SuperGridOptions` structs to make this widget more flexible.
- This is a project in development, so it can have some bugs.
- If you find a bug or have a suggestion, please, create an issue.
- If you want to contribute, please, feel free to create a pull request.
