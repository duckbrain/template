// +build !release

package assets

func NewHelper(c Config) (Helper, error) {
	return NewDebugHelper(c)
}
