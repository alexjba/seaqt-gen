package qt6

/*

#include "gen_qpen.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QPen struct {
	h *C.QPen
}

func (this *QPen) cPointer() *C.QPen {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QPen) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQPen constructs the type using only CGO pointers.
func newQPen(h *C.QPen) *QPen {
	if h == nil {
		return nil
	}

	return &QPen{h: h}
}

// UnsafeNewQPen constructs the type using only unsafe pointers.
func UnsafeNewQPen(h unsafe.Pointer) *QPen {
	return newQPen((*C.QPen)(h))
}

// NewQPen constructs a new QPen object.
func NewQPen() *QPen {

	return newQPen(C.QPen_new())
}

// NewQPen2 constructs a new QPen object.
func NewQPen2(param1 PenStyle) *QPen {

	return newQPen(C.QPen_new2((C.int)(param1)))
}

// NewQPen3 constructs a new QPen object.
func NewQPen3(color *QColor) *QPen {

	return newQPen(C.QPen_new3(color.cPointer()))
}

// NewQPen4 constructs a new QPen object.
func NewQPen4(brush *QBrush, width float64) *QPen {

	return newQPen(C.QPen_new4(brush.cPointer(), (C.double)(width)))
}

// NewQPen5 constructs a new QPen object.
func NewQPen5(pen *QPen) *QPen {

	return newQPen(C.QPen_new5(pen.cPointer()))
}

// NewQPen6 constructs a new QPen object.
func NewQPen6(brush *QBrush, width float64, s PenStyle) *QPen {

	return newQPen(C.QPen_new6(brush.cPointer(), (C.double)(width), (C.int)(s)))
}

// NewQPen7 constructs a new QPen object.
func NewQPen7(brush *QBrush, width float64, s PenStyle, c PenCapStyle) *QPen {

	return newQPen(C.QPen_new7(brush.cPointer(), (C.double)(width), (C.int)(s), (C.int)(c)))
}

// NewQPen8 constructs a new QPen object.
func NewQPen8(brush *QBrush, width float64, s PenStyle, c PenCapStyle, j PenJoinStyle) *QPen {

	return newQPen(C.QPen_new8(brush.cPointer(), (C.double)(width), (C.int)(s), (C.int)(c), (C.int)(j)))
}

func (this *QPen) OperatorAssign(pen *QPen) {
	C.QPen_OperatorAssign(this.h, pen.cPointer())
}

func (this *QPen) Swap(other *QPen) {
	C.QPen_Swap(this.h, other.cPointer())
}

func (this *QPen) Style() PenStyle {
	return (PenStyle)(C.QPen_Style(this.h))
}

func (this *QPen) SetStyle(style PenStyle) {
	C.QPen_SetStyle(this.h, (C.int)(style))
}

func (this *QPen) DashPattern() []float64 {
	var _ma C.struct_miqt_array = C.QPen_DashPattern(this.h)
	_ret := make([]float64, int(_ma.len))
	_outCast := (*[0xffff]C.double)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (float64)(_outCast[i])
	}
	return _ret
}

func (this *QPen) SetDashPattern(pattern []float64) {
	pattern_CArray := (*[0xffff]C.double)(C.malloc(C.size_t(8 * len(pattern))))
	defer C.free(unsafe.Pointer(pattern_CArray))
	for i := range pattern {
		pattern_CArray[i] = (C.double)(pattern[i])
	}
	pattern_ma := C.struct_miqt_array{len: C.size_t(len(pattern)), data: unsafe.Pointer(pattern_CArray)}
	C.QPen_SetDashPattern(this.h, pattern_ma)
}

func (this *QPen) DashOffset() float64 {
	return (float64)(C.QPen_DashOffset(this.h))
}

func (this *QPen) SetDashOffset(doffset float64) {
	C.QPen_SetDashOffset(this.h, (C.double)(doffset))
}

func (this *QPen) MiterLimit() float64 {
	return (float64)(C.QPen_MiterLimit(this.h))
}

func (this *QPen) SetMiterLimit(limit float64) {
	C.QPen_SetMiterLimit(this.h, (C.double)(limit))
}

func (this *QPen) WidthF() float64 {
	return (float64)(C.QPen_WidthF(this.h))
}

func (this *QPen) SetWidthF(width float64) {
	C.QPen_SetWidthF(this.h, (C.double)(width))
}

func (this *QPen) Width() int {
	return (int)(C.QPen_Width(this.h))
}

func (this *QPen) SetWidth(width int) {
	C.QPen_SetWidth(this.h, (C.int)(width))
}

func (this *QPen) Color() *QColor {
	_goptr := newQColor(C.QPen_Color(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPen) SetColor(color *QColor) {
	C.QPen_SetColor(this.h, color.cPointer())
}

func (this *QPen) Brush() *QBrush {
	_goptr := newQBrush(C.QPen_Brush(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPen) SetBrush(brush *QBrush) {
	C.QPen_SetBrush(this.h, brush.cPointer())
}

func (this *QPen) IsSolid() bool {
	return (bool)(C.QPen_IsSolid(this.h))
}

func (this *QPen) CapStyle() PenCapStyle {
	return (PenCapStyle)(C.QPen_CapStyle(this.h))
}

func (this *QPen) SetCapStyle(pcs PenCapStyle) {
	C.QPen_SetCapStyle(this.h, (C.int)(pcs))
}

func (this *QPen) JoinStyle() PenJoinStyle {
	return (PenJoinStyle)(C.QPen_JoinStyle(this.h))
}

func (this *QPen) SetJoinStyle(pcs PenJoinStyle) {
	C.QPen_SetJoinStyle(this.h, (C.int)(pcs))
}

func (this *QPen) IsCosmetic() bool {
	return (bool)(C.QPen_IsCosmetic(this.h))
}

func (this *QPen) SetCosmetic(cosmetic bool) {
	C.QPen_SetCosmetic(this.h, (C.bool)(cosmetic))
}

func (this *QPen) OperatorEqual(p *QPen) bool {
	return (bool)(C.QPen_OperatorEqual(this.h, p.cPointer()))
}

func (this *QPen) OperatorNotEqual(p *QPen) bool {
	return (bool)(C.QPen_OperatorNotEqual(this.h, p.cPointer()))
}

func (this *QPen) ToQVariant() *QVariant {
	_goptr := newQVariant(C.QPen_ToQVariant(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPen) IsDetached() bool {
	return (bool)(C.QPen_IsDetached(this.h))
}

// Delete this object from C++ memory.
func (this *QPen) Delete() {
	C.QPen_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QPen) GoGC() {
	runtime.SetFinalizer(this, func(this *QPen) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
