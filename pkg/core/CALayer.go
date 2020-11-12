package core

import "github.com/progrium/macdriver/pkg/objc"

type CALayer struct {
	objc.Object
}

func (l CALayer) CornerRadius() float64 {
	return l.Get("cornerRadius").Float()
}

func (l CALayer) SetCornerRadius(r float64) {
	l.Set("cornerRadius:", r)
}
