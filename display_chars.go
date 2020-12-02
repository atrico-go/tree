package tree

type BoxType int

const (
	// Type of line
	BoxNone   BoxType = 0
	BoxSingle BoxType = 1
)

type direction uint32

func (d direction) getUp() BoxType {
	return d.getDir(0xff000000, 24)
}
func (d *direction) setUp(boxType BoxType) *direction {
	return d.setDir(0x00ffffff, 24, boxType)
}
func (d direction) getDown() BoxType {
	return d.getDir(0x00ff0000, 16)
}
func (d *direction) setDown(boxType BoxType) *direction {
	return d.setDir(0xff00ffff, 16, boxType)
}
func (d direction) getLeft() BoxType {
	return d.getDir(0x0000ff00, 8)
}
func (d *direction) setLeft(boxType BoxType) *direction {
	return d.setDir(0xffff00ff, 8, boxType)
}
func (d direction) getRight() BoxType {
	return d.getDir(0x000000ff, 0)
}
func (d *direction) setRight(boxType BoxType) *direction {
	return d.setDir(0xffffff00, 0, boxType)
}
func (d direction) getDir(mask uint32, shift int) BoxType {
	return BoxType((uint32(d) & mask) >> shift)
}
func (d *direction) setDir(mask uint32, shift int, boxType BoxType) *direction {
	*d = *d & direction(mask)
	*d = *d | direction(boxType<<shift)
	return d
}
func createDir(up BoxType, down BoxType, left BoxType, right BoxType) direction {
	d := direction(0)
	return *d.setUp(up).setDown(down).setLeft(left).setRight(right)
}

var boxParts = map[direction]string{
	    createDir(BoxNone, BoxNone, BoxNone, BoxNone): " ",
		createDir(BoxNone, BoxNone, BoxNone, BoxSingle): "╶",
		createDir(BoxNone, BoxNone, BoxSingle, BoxNone): "╴",
		createDir(BoxNone, BoxNone, BoxSingle, BoxSingle): "─",
		createDir(BoxNone, BoxSingle, BoxNone, BoxNone): "╷",
		createDir(BoxNone, BoxSingle, BoxNone, BoxSingle): "┌",
		createDir(BoxNone, BoxSingle, BoxSingle, BoxNone): "┐",
		createDir(BoxNone, BoxSingle, BoxSingle, BoxSingle): "┬",
		createDir(BoxSingle, BoxNone, BoxNone, BoxNone): "╵",
		createDir(BoxSingle, BoxNone, BoxNone, BoxSingle): "└",
		createDir(BoxSingle, BoxNone, BoxSingle, BoxNone): "┘",
		createDir(BoxSingle, BoxNone, BoxSingle, BoxSingle): "┴",
		createDir(BoxSingle, BoxSingle, BoxNone, BoxNone): "│",
		createDir(BoxSingle, BoxSingle, BoxNone, BoxSingle): "├",
		createDir(BoxSingle, BoxSingle, BoxSingle, BoxNone): "┤",
		createDir(BoxSingle, BoxSingle, BoxSingle, BoxSingle): "┼",
}


func getBoxChar(up BoxType,down BoxType,left BoxType,right BoxType) string {
	return boxParts[createDir(up,down,left,right)]
}