package output

type Destination interface {
	Close() error
	Print(state [][]int) error
}
