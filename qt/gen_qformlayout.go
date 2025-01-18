package qt

/*

#include "gen_qformlayout.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QFormLayout__FieldGrowthPolicy int

const (
	QFormLayout__FieldsStayAtSizeHint  QFormLayout__FieldGrowthPolicy = 0
	QFormLayout__ExpandingFieldsGrow   QFormLayout__FieldGrowthPolicy = 1
	QFormLayout__AllNonFixedFieldsGrow QFormLayout__FieldGrowthPolicy = 2
)

type QFormLayout__RowWrapPolicy int

const (
	QFormLayout__DontWrapRows QFormLayout__RowWrapPolicy = 0
	QFormLayout__WrapLongRows QFormLayout__RowWrapPolicy = 1
	QFormLayout__WrapAllRows  QFormLayout__RowWrapPolicy = 2
)

type QFormLayout__ItemRole int

const (
	QFormLayout__LabelRole    QFormLayout__ItemRole = 0
	QFormLayout__FieldRole    QFormLayout__ItemRole = 1
	QFormLayout__SpanningRole QFormLayout__ItemRole = 2
)

type QFormLayout struct {
	h          *C.QFormLayout
	isSubclass bool
	*QLayout
}

func (this *QFormLayout) cPointer() *C.QFormLayout {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QFormLayout) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQFormLayout constructs the type using only CGO pointers.
func newQFormLayout(h *C.QFormLayout) *QFormLayout {
	if h == nil {
		return nil
	}
	var outptr_QLayout *C.QLayout = nil
	C.QFormLayout_virtbase(h, &outptr_QLayout)

	return &QFormLayout{h: h,
		QLayout: newQLayout(outptr_QLayout)}
}

// UnsafeNewQFormLayout constructs the type using only unsafe pointers.
func UnsafeNewQFormLayout(h unsafe.Pointer) *QFormLayout {
	return newQFormLayout((*C.QFormLayout)(h))
}

// NewQFormLayout constructs a new QFormLayout object.
func NewQFormLayout(parent *QWidget) *QFormLayout {

	ret := newQFormLayout(C.QFormLayout_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFormLayout2 constructs a new QFormLayout object.
func NewQFormLayout2() *QFormLayout {

	ret := newQFormLayout(C.QFormLayout_new2())
	ret.isSubclass = true
	return ret
}

func (this *QFormLayout) MetaObject() *QMetaObject {
	return newQMetaObject(C.QFormLayout_MetaObject(this.h))
}

func (this *QFormLayout) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QFormLayout_Metacast(this.h, param1_Cstring))
}

func QFormLayout_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QFormLayout_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFormLayout_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QFormLayout_TrUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFormLayout) SetFieldGrowthPolicy(policy QFormLayout__FieldGrowthPolicy) {
	C.QFormLayout_SetFieldGrowthPolicy(this.h, (C.int)(policy))
}

func (this *QFormLayout) FieldGrowthPolicy() QFormLayout__FieldGrowthPolicy {
	return (QFormLayout__FieldGrowthPolicy)(C.QFormLayout_FieldGrowthPolicy(this.h))
}

func (this *QFormLayout) SetRowWrapPolicy(policy QFormLayout__RowWrapPolicy) {
	C.QFormLayout_SetRowWrapPolicy(this.h, (C.int)(policy))
}

func (this *QFormLayout) RowWrapPolicy() QFormLayout__RowWrapPolicy {
	return (QFormLayout__RowWrapPolicy)(C.QFormLayout_RowWrapPolicy(this.h))
}

func (this *QFormLayout) SetLabelAlignment(alignment AlignmentFlag) {
	C.QFormLayout_SetLabelAlignment(this.h, (C.int)(alignment))
}

func (this *QFormLayout) LabelAlignment() AlignmentFlag {
	return (AlignmentFlag)(C.QFormLayout_LabelAlignment(this.h))
}

func (this *QFormLayout) SetFormAlignment(alignment AlignmentFlag) {
	C.QFormLayout_SetFormAlignment(this.h, (C.int)(alignment))
}

func (this *QFormLayout) FormAlignment() AlignmentFlag {
	return (AlignmentFlag)(C.QFormLayout_FormAlignment(this.h))
}

func (this *QFormLayout) SetHorizontalSpacing(spacing int) {
	C.QFormLayout_SetHorizontalSpacing(this.h, (C.int)(spacing))
}

func (this *QFormLayout) HorizontalSpacing() int {
	return (int)(C.QFormLayout_HorizontalSpacing(this.h))
}

func (this *QFormLayout) SetVerticalSpacing(spacing int) {
	C.QFormLayout_SetVerticalSpacing(this.h, (C.int)(spacing))
}

func (this *QFormLayout) VerticalSpacing() int {
	return (int)(C.QFormLayout_VerticalSpacing(this.h))
}

func (this *QFormLayout) Spacing() int {
	return (int)(C.QFormLayout_Spacing(this.h))
}

func (this *QFormLayout) SetSpacing(spacing int) {
	C.QFormLayout_SetSpacing(this.h, (C.int)(spacing))
}

func (this *QFormLayout) AddRow(label *QWidget, field *QWidget) {
	C.QFormLayout_AddRow(this.h, label.cPointer(), field.cPointer())
}

func (this *QFormLayout) AddRow2(label *QWidget, field *QLayout) {
	C.QFormLayout_AddRow2(this.h, label.cPointer(), field.cPointer())
}

func (this *QFormLayout) AddRow3(labelText string, field *QWidget) {
	labelText_ms := C.struct_miqt_string{}
	labelText_ms.data = C.CString(labelText)
	labelText_ms.len = C.size_t(len(labelText))
	defer C.free(unsafe.Pointer(labelText_ms.data))
	C.QFormLayout_AddRow3(this.h, labelText_ms, field.cPointer())
}

func (this *QFormLayout) AddRow4(labelText string, field *QLayout) {
	labelText_ms := C.struct_miqt_string{}
	labelText_ms.data = C.CString(labelText)
	labelText_ms.len = C.size_t(len(labelText))
	defer C.free(unsafe.Pointer(labelText_ms.data))
	C.QFormLayout_AddRow4(this.h, labelText_ms, field.cPointer())
}

func (this *QFormLayout) AddRowWithWidget(widget *QWidget) {
	C.QFormLayout_AddRowWithWidget(this.h, widget.cPointer())
}

func (this *QFormLayout) AddRowWithLayout(layout *QLayout) {
	C.QFormLayout_AddRowWithLayout(this.h, layout.cPointer())
}

func (this *QFormLayout) InsertRow(row int, label *QWidget, field *QWidget) {
	C.QFormLayout_InsertRow(this.h, (C.int)(row), label.cPointer(), field.cPointer())
}

func (this *QFormLayout) InsertRow2(row int, label *QWidget, field *QLayout) {
	C.QFormLayout_InsertRow2(this.h, (C.int)(row), label.cPointer(), field.cPointer())
}

func (this *QFormLayout) InsertRow3(row int, labelText string, field *QWidget) {
	labelText_ms := C.struct_miqt_string{}
	labelText_ms.data = C.CString(labelText)
	labelText_ms.len = C.size_t(len(labelText))
	defer C.free(unsafe.Pointer(labelText_ms.data))
	C.QFormLayout_InsertRow3(this.h, (C.int)(row), labelText_ms, field.cPointer())
}

func (this *QFormLayout) InsertRow4(row int, labelText string, field *QLayout) {
	labelText_ms := C.struct_miqt_string{}
	labelText_ms.data = C.CString(labelText)
	labelText_ms.len = C.size_t(len(labelText))
	defer C.free(unsafe.Pointer(labelText_ms.data))
	C.QFormLayout_InsertRow4(this.h, (C.int)(row), labelText_ms, field.cPointer())
}

func (this *QFormLayout) InsertRow5(row int, widget *QWidget) {
	C.QFormLayout_InsertRow5(this.h, (C.int)(row), widget.cPointer())
}

func (this *QFormLayout) InsertRow6(row int, layout *QLayout) {
	C.QFormLayout_InsertRow6(this.h, (C.int)(row), layout.cPointer())
}

func (this *QFormLayout) RemoveRow(row int) {
	C.QFormLayout_RemoveRow(this.h, (C.int)(row))
}

func (this *QFormLayout) RemoveRowWithWidget(widget *QWidget) {
	C.QFormLayout_RemoveRowWithWidget(this.h, widget.cPointer())
}

func (this *QFormLayout) RemoveRowWithLayout(layout *QLayout) {
	C.QFormLayout_RemoveRowWithLayout(this.h, layout.cPointer())
}

func (this *QFormLayout) TakeRow(row int) *QFormLayout__TakeRowResult {
	_goptr := newQFormLayout__TakeRowResult(C.QFormLayout_TakeRow(this.h, (C.int)(row)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFormLayout) TakeRowWithWidget(widget *QWidget) *QFormLayout__TakeRowResult {
	_goptr := newQFormLayout__TakeRowResult(C.QFormLayout_TakeRowWithWidget(this.h, widget.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFormLayout) TakeRowWithLayout(layout *QLayout) *QFormLayout__TakeRowResult {
	_goptr := newQFormLayout__TakeRowResult(C.QFormLayout_TakeRowWithLayout(this.h, layout.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFormLayout) SetItem(row int, role QFormLayout__ItemRole, item *QLayoutItem) {
	C.QFormLayout_SetItem(this.h, (C.int)(row), (C.int)(role), item.cPointer())
}

func (this *QFormLayout) SetWidget(row int, role QFormLayout__ItemRole, widget *QWidget) {
	C.QFormLayout_SetWidget(this.h, (C.int)(row), (C.int)(role), widget.cPointer())
}

func (this *QFormLayout) SetLayout(row int, role QFormLayout__ItemRole, layout *QLayout) {
	C.QFormLayout_SetLayout(this.h, (C.int)(row), (C.int)(role), layout.cPointer())
}

func (this *QFormLayout) ItemAt(row int, role QFormLayout__ItemRole) *QLayoutItem {
	return newQLayoutItem(C.QFormLayout_ItemAt(this.h, (C.int)(row), (C.int)(role)))
}

func (this *QFormLayout) LabelForField(field *QWidget) *QWidget {
	return newQWidget(C.QFormLayout_LabelForField(this.h, field.cPointer()))
}

func (this *QFormLayout) LabelForFieldWithField(field *QLayout) *QWidget {
	return newQWidget(C.QFormLayout_LabelForFieldWithField(this.h, field.cPointer()))
}

func (this *QFormLayout) AddItem(item *QLayoutItem) {
	C.QFormLayout_AddItem(this.h, item.cPointer())
}

func (this *QFormLayout) ItemAtWithIndex(index int) *QLayoutItem {
	return newQLayoutItem(C.QFormLayout_ItemAtWithIndex(this.h, (C.int)(index)))
}

func (this *QFormLayout) TakeAt(index int) *QLayoutItem {
	return newQLayoutItem(C.QFormLayout_TakeAt(this.h, (C.int)(index)))
}

func (this *QFormLayout) SetGeometry(rect *QRect) {
	C.QFormLayout_SetGeometry(this.h, rect.cPointer())
}

func (this *QFormLayout) MinimumSize() *QSize {
	_goptr := newQSize(C.QFormLayout_MinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFormLayout) SizeHint() *QSize {
	_goptr := newQSize(C.QFormLayout_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFormLayout) Invalidate() {
	C.QFormLayout_Invalidate(this.h)
}

func (this *QFormLayout) HasHeightForWidth() bool {
	return (bool)(C.QFormLayout_HasHeightForWidth(this.h))
}

func (this *QFormLayout) HeightForWidth(width int) int {
	return (int)(C.QFormLayout_HeightForWidth(this.h, (C.int)(width)))
}

func (this *QFormLayout) ExpandingDirections() Orientation {
	return (Orientation)(C.QFormLayout_ExpandingDirections(this.h))
}

func (this *QFormLayout) Count() int {
	return (int)(C.QFormLayout_Count(this.h))
}

func (this *QFormLayout) RowCount() int {
	return (int)(C.QFormLayout_RowCount(this.h))
}

func QFormLayout_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QFormLayout_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFormLayout_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QFormLayout_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFormLayout_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QFormLayout_TrUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFormLayout_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QFormLayout_TrUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFormLayout) callVirtualBase_AddItem(item *QLayoutItem) {

	C.QFormLayout_virtualbase_AddItem(unsafe.Pointer(this.h), item.cPointer())

}
func (this *QFormLayout) OnAddItem(slot func(super func(item *QLayoutItem), item *QLayoutItem)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QFormLayout_override_virtual_AddItem(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_AddItem
func miqt_exec_callback_QFormLayout_AddItem(self *C.QFormLayout, cb C.intptr_t, item *C.QLayoutItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(item *QLayoutItem), item *QLayoutItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQLayoutItem(item)

	gofunc((&QFormLayout{h: self}).callVirtualBase_AddItem, slotval1)

}

func (this *QFormLayout) callVirtualBase_ItemAtWithIndex(index int) *QLayoutItem {

	return newQLayoutItem(C.QFormLayout_virtualbase_ItemAtWithIndex(unsafe.Pointer(this.h), (C.int)(index)))

}
func (this *QFormLayout) OnItemAtWithIndex(slot func(super func(index int) *QLayoutItem, index int) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QFormLayout_override_virtual_ItemAtWithIndex(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_ItemAtWithIndex
func miqt_exec_callback_QFormLayout_ItemAtWithIndex(self *C.QFormLayout, cb C.intptr_t, index C.int) *C.QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int) *QLayoutItem, index int) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_ItemAtWithIndex, slotval1)

	return virtualReturn.cPointer()

}

func (this *QFormLayout) callVirtualBase_TakeAt(index int) *QLayoutItem {

	return newQLayoutItem(C.QFormLayout_virtualbase_TakeAt(unsafe.Pointer(this.h), (C.int)(index)))

}
func (this *QFormLayout) OnTakeAt(slot func(super func(index int) *QLayoutItem, index int) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QFormLayout_override_virtual_TakeAt(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_TakeAt
func miqt_exec_callback_QFormLayout_TakeAt(self *C.QFormLayout, cb C.intptr_t, index C.int) *C.QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int) *QLayoutItem, index int) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_TakeAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QFormLayout) callVirtualBase_SetGeometry(rect *QRect) {

	C.QFormLayout_virtualbase_SetGeometry(unsafe.Pointer(this.h), rect.cPointer())

}
func (this *QFormLayout) OnSetGeometry(slot func(super func(rect *QRect), rect *QRect)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QFormLayout_override_virtual_SetGeometry(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_SetGeometry
func miqt_exec_callback_QFormLayout_SetGeometry(self *C.QFormLayout, cb C.intptr_t, rect *C.QRect) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rect *QRect), rect *QRect))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(rect)

	gofunc((&QFormLayout{h: self}).callVirtualBase_SetGeometry, slotval1)

}

func (this *QFormLayout) callVirtualBase_MinimumSize() *QSize {

	_goptr := newQSize(C.QFormLayout_virtualbase_MinimumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFormLayout) OnMinimumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QFormLayout_override_virtual_MinimumSize(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_MinimumSize
func miqt_exec_callback_QFormLayout_MinimumSize(self *C.QFormLayout, cb C.intptr_t) *C.QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_MinimumSize)

	return virtualReturn.cPointer()

}

func (this *QFormLayout) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(C.QFormLayout_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFormLayout) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QFormLayout_override_virtual_SizeHint(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_SizeHint
func miqt_exec_callback_QFormLayout_SizeHint(self *C.QFormLayout, cb C.intptr_t) *C.QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QFormLayout) callVirtualBase_Invalidate() {

	C.QFormLayout_virtualbase_Invalidate(unsafe.Pointer(this.h))

}
func (this *QFormLayout) OnInvalidate(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QFormLayout_override_virtual_Invalidate(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_Invalidate
func miqt_exec_callback_QFormLayout_Invalidate(self *C.QFormLayout, cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QFormLayout{h: self}).callVirtualBase_Invalidate)

}

func (this *QFormLayout) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(C.QFormLayout_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QFormLayout) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QFormLayout_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_HasHeightForWidth
func miqt_exec_callback_QFormLayout_HasHeightForWidth(self *C.QFormLayout, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_HasHeightForWidth)

	return (C.bool)(virtualReturn)

}

func (this *QFormLayout) callVirtualBase_HeightForWidth(width int) int {

	return (int)(C.QFormLayout_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (C.int)(width)))

}
func (this *QFormLayout) OnHeightForWidth(slot func(super func(width int) int, width int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QFormLayout_override_virtual_HeightForWidth(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_HeightForWidth
func miqt_exec_callback_QFormLayout_HeightForWidth(self *C.QFormLayout, cb C.intptr_t, width C.int) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(width int) int, width int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(width)

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (C.int)(virtualReturn)

}

func (this *QFormLayout) callVirtualBase_ExpandingDirections() Orientation {

	return (Orientation)(C.QFormLayout_virtualbase_ExpandingDirections(unsafe.Pointer(this.h)))

}
func (this *QFormLayout) OnExpandingDirections(slot func(super func() Orientation) Orientation) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QFormLayout_override_virtual_ExpandingDirections(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_ExpandingDirections
func miqt_exec_callback_QFormLayout_ExpandingDirections(self *C.QFormLayout, cb C.intptr_t) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() Orientation) Orientation)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_ExpandingDirections)

	return (C.int)(virtualReturn)

}

func (this *QFormLayout) callVirtualBase_Count() int {

	return (int)(C.QFormLayout_virtualbase_Count(unsafe.Pointer(this.h)))

}
func (this *QFormLayout) OnCount(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QFormLayout_override_virtual_Count(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_Count
func miqt_exec_callback_QFormLayout_Count(self *C.QFormLayout, cb C.intptr_t) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_Count)

	return (C.int)(virtualReturn)

}

func (this *QFormLayout) callVirtualBase_Geometry() *QRect {

	_goptr := newQRect(C.QFormLayout_virtualbase_Geometry(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFormLayout) OnGeometry(slot func(super func() *QRect) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QFormLayout_override_virtual_Geometry(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_Geometry
func miqt_exec_callback_QFormLayout_Geometry(self *C.QFormLayout, cb C.intptr_t) *C.QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QRect) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_Geometry)

	return virtualReturn.cPointer()

}

func (this *QFormLayout) callVirtualBase_MaximumSize() *QSize {

	_goptr := newQSize(C.QFormLayout_virtualbase_MaximumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFormLayout) OnMaximumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QFormLayout_override_virtual_MaximumSize(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_MaximumSize
func miqt_exec_callback_QFormLayout_MaximumSize(self *C.QFormLayout, cb C.intptr_t) *C.QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_MaximumSize)

	return virtualReturn.cPointer()

}

func (this *QFormLayout) callVirtualBase_IndexOf(param1 *QWidget) int {

	return (int)(C.QFormLayout_virtualbase_IndexOf(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QFormLayout) OnIndexOf(slot func(super func(param1 *QWidget) int, param1 *QWidget) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QFormLayout_override_virtual_IndexOf(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_IndexOf
func miqt_exec_callback_QFormLayout_IndexOf(self *C.QFormLayout, cb C.intptr_t, param1 *C.QWidget) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QWidget) int, param1 *QWidget) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(param1)

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_IndexOf, slotval1)

	return (C.int)(virtualReturn)

}

func (this *QFormLayout) callVirtualBase_IsEmpty() bool {

	return (bool)(C.QFormLayout_virtualbase_IsEmpty(unsafe.Pointer(this.h)))

}
func (this *QFormLayout) OnIsEmpty(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QFormLayout_override_virtual_IsEmpty(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_IsEmpty
func miqt_exec_callback_QFormLayout_IsEmpty(self *C.QFormLayout, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_IsEmpty)

	return (C.bool)(virtualReturn)

}

func (this *QFormLayout) callVirtualBase_ControlTypes() QSizePolicy__ControlType {

	return (QSizePolicy__ControlType)(C.QFormLayout_virtualbase_ControlTypes(unsafe.Pointer(this.h)))

}
func (this *QFormLayout) OnControlTypes(slot func(super func() QSizePolicy__ControlType) QSizePolicy__ControlType) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QFormLayout_override_virtual_ControlTypes(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_ControlTypes
func miqt_exec_callback_QFormLayout_ControlTypes(self *C.QFormLayout, cb C.intptr_t) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() QSizePolicy__ControlType) QSizePolicy__ControlType)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_ControlTypes)

	return (C.int)(virtualReturn)

}

func (this *QFormLayout) callVirtualBase_Layout() *QLayout {

	return newQLayout(C.QFormLayout_virtualbase_Layout(unsafe.Pointer(this.h)))

}
func (this *QFormLayout) OnLayout(slot func(super func() *QLayout) *QLayout) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QFormLayout_override_virtual_Layout(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_Layout
func miqt_exec_callback_QFormLayout_Layout(self *C.QFormLayout, cb C.intptr_t) *C.QLayout {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QLayout) *QLayout)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_Layout)

	return virtualReturn.cPointer()

}

func (this *QFormLayout) callVirtualBase_ChildEvent(e *QChildEvent) {

	C.QFormLayout_virtualbase_ChildEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QFormLayout) OnChildEvent(slot func(super func(e *QChildEvent), e *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QFormLayout_override_virtual_ChildEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_ChildEvent
func miqt_exec_callback_QFormLayout_ChildEvent(self *C.QFormLayout, cb C.intptr_t, e *C.QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QChildEvent), e *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(e)

	gofunc((&QFormLayout{h: self}).callVirtualBase_ChildEvent, slotval1)

}

// Delete this object from C++ memory.
func (this *QFormLayout) Delete() {
	C.QFormLayout_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QFormLayout) GoGC() {
	runtime.SetFinalizer(this, func(this *QFormLayout) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QFormLayout__TakeRowResult struct {
	h          *C.QFormLayout__TakeRowResult
	isSubclass bool
}

func (this *QFormLayout__TakeRowResult) cPointer() *C.QFormLayout__TakeRowResult {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QFormLayout__TakeRowResult) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQFormLayout__TakeRowResult constructs the type using only CGO pointers.
func newQFormLayout__TakeRowResult(h *C.QFormLayout__TakeRowResult) *QFormLayout__TakeRowResult {
	if h == nil {
		return nil
	}

	return &QFormLayout__TakeRowResult{h: h}
}

// UnsafeNewQFormLayout__TakeRowResult constructs the type using only unsafe pointers.
func UnsafeNewQFormLayout__TakeRowResult(h unsafe.Pointer) *QFormLayout__TakeRowResult {
	return newQFormLayout__TakeRowResult((*C.QFormLayout__TakeRowResult)(h))
}

// Delete this object from C++ memory.
func (this *QFormLayout__TakeRowResult) Delete() {
	C.QFormLayout__TakeRowResult_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QFormLayout__TakeRowResult) GoGC() {
	runtime.SetFinalizer(this, func(this *QFormLayout__TakeRowResult) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
