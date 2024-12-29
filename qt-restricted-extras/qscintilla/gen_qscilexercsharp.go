package qscintilla

/*

#include "gen_qscilexercsharp.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QsciLexerCSharp struct {
	h          *C.QsciLexerCSharp
	isSubclass bool
	*QsciLexerCPP
}

func (this *QsciLexerCSharp) cPointer() *C.QsciLexerCSharp {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QsciLexerCSharp) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQsciLexerCSharp constructs the type using only CGO pointers.
func newQsciLexerCSharp(h *C.QsciLexerCSharp) *QsciLexerCSharp {
	if h == nil {
		return nil
	}
	var outptr_QsciLexerCPP *C.QsciLexerCPP = nil
	C.QsciLexerCSharp_virtbase(h, &outptr_QsciLexerCPP)

	return &QsciLexerCSharp{h: h,
		QsciLexerCPP: newQsciLexerCPP(outptr_QsciLexerCPP)}
}

// UnsafeNewQsciLexerCSharp constructs the type using only unsafe pointers.
func UnsafeNewQsciLexerCSharp(h unsafe.Pointer) *QsciLexerCSharp {
	return newQsciLexerCSharp((*C.QsciLexerCSharp)(h))
}

// NewQsciLexerCSharp constructs a new QsciLexerCSharp object.
func NewQsciLexerCSharp() *QsciLexerCSharp {

	ret := newQsciLexerCSharp(C.QsciLexerCSharp_new())
	ret.isSubclass = true
	return ret
}

// NewQsciLexerCSharp2 constructs a new QsciLexerCSharp object.
func NewQsciLexerCSharp2(parent *qt.QObject) *QsciLexerCSharp {

	ret := newQsciLexerCSharp(C.QsciLexerCSharp_new2((*C.QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QsciLexerCSharp) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(C.QsciLexerCSharp_MetaObject(this.h)))
}

func (this *QsciLexerCSharp) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QsciLexerCSharp_Metacast(this.h, param1_Cstring))
}

func QsciLexerCSharp_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QsciLexerCSharp_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QsciLexerCSharp_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QsciLexerCSharp_TrUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QsciLexerCSharp) Language() string {
	_ret := C.QsciLexerCSharp_Language(this.h)
	return C.GoString(_ret)
}

func (this *QsciLexerCSharp) DefaultColor(style int) *qt.QColor {
	_goptr := qt.UnsafeNewQColor(unsafe.Pointer(C.QsciLexerCSharp_DefaultColor(this.h, (C.int)(style))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QsciLexerCSharp) DefaultEolFill(style int) bool {
	return (bool)(C.QsciLexerCSharp_DefaultEolFill(this.h, (C.int)(style)))
}

func (this *QsciLexerCSharp) DefaultFont(style int) *qt.QFont {
	_goptr := qt.UnsafeNewQFont(unsafe.Pointer(C.QsciLexerCSharp_DefaultFont(this.h, (C.int)(style))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QsciLexerCSharp) DefaultPaper(style int) *qt.QColor {
	_goptr := qt.UnsafeNewQColor(unsafe.Pointer(C.QsciLexerCSharp_DefaultPaper(this.h, (C.int)(style))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QsciLexerCSharp) Keywords(set int) string {
	_ret := C.QsciLexerCSharp_Keywords(this.h, (C.int)(set))
	return C.GoString(_ret)
}

func (this *QsciLexerCSharp) Description(style int) string {
	var _ms C.struct_miqt_string = C.QsciLexerCSharp_Description(this.h, (C.int)(style))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QsciLexerCSharp_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QsciLexerCSharp_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QsciLexerCSharp_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QsciLexerCSharp_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QsciLexerCSharp_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QsciLexerCSharp_TrUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QsciLexerCSharp_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QsciLexerCSharp_TrUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QsciLexerCSharp) callVirtualBase_SetFoldAtElse(fold bool) {

	C.QsciLexerCSharp_virtualbase_SetFoldAtElse(unsafe.Pointer(this.h), (C.bool)(fold))

}
func (this *QsciLexerCSharp) OnSetFoldAtElse(slot func(super func(fold bool), fold bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QsciLexerCSharp_override_virtual_SetFoldAtElse(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciLexerCSharp_SetFoldAtElse
func miqt_exec_callback_QsciLexerCSharp_SetFoldAtElse(self *C.QsciLexerCSharp, cb C.intptr_t, fold C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(fold bool), fold bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(fold)

	gofunc((&QsciLexerCSharp{h: self}).callVirtualBase_SetFoldAtElse, slotval1)

}

func (this *QsciLexerCSharp) callVirtualBase_SetFoldComments(fold bool) {

	C.QsciLexerCSharp_virtualbase_SetFoldComments(unsafe.Pointer(this.h), (C.bool)(fold))

}
func (this *QsciLexerCSharp) OnSetFoldComments(slot func(super func(fold bool), fold bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QsciLexerCSharp_override_virtual_SetFoldComments(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciLexerCSharp_SetFoldComments
func miqt_exec_callback_QsciLexerCSharp_SetFoldComments(self *C.QsciLexerCSharp, cb C.intptr_t, fold C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(fold bool), fold bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(fold)

	gofunc((&QsciLexerCSharp{h: self}).callVirtualBase_SetFoldComments, slotval1)

}

func (this *QsciLexerCSharp) callVirtualBase_SetFoldCompact(fold bool) {

	C.QsciLexerCSharp_virtualbase_SetFoldCompact(unsafe.Pointer(this.h), (C.bool)(fold))

}
func (this *QsciLexerCSharp) OnSetFoldCompact(slot func(super func(fold bool), fold bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QsciLexerCSharp_override_virtual_SetFoldCompact(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciLexerCSharp_SetFoldCompact
func miqt_exec_callback_QsciLexerCSharp_SetFoldCompact(self *C.QsciLexerCSharp, cb C.intptr_t, fold C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(fold bool), fold bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(fold)

	gofunc((&QsciLexerCSharp{h: self}).callVirtualBase_SetFoldCompact, slotval1)

}

func (this *QsciLexerCSharp) callVirtualBase_SetFoldPreprocessor(fold bool) {

	C.QsciLexerCSharp_virtualbase_SetFoldPreprocessor(unsafe.Pointer(this.h), (C.bool)(fold))

}
func (this *QsciLexerCSharp) OnSetFoldPreprocessor(slot func(super func(fold bool), fold bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QsciLexerCSharp_override_virtual_SetFoldPreprocessor(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciLexerCSharp_SetFoldPreprocessor
func miqt_exec_callback_QsciLexerCSharp_SetFoldPreprocessor(self *C.QsciLexerCSharp, cb C.intptr_t, fold C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(fold bool), fold bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(fold)

	gofunc((&QsciLexerCSharp{h: self}).callVirtualBase_SetFoldPreprocessor, slotval1)

}

func (this *QsciLexerCSharp) callVirtualBase_SetStylePreprocessor(style bool) {

	C.QsciLexerCSharp_virtualbase_SetStylePreprocessor(unsafe.Pointer(this.h), (C.bool)(style))

}
func (this *QsciLexerCSharp) OnSetStylePreprocessor(slot func(super func(style bool), style bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QsciLexerCSharp_override_virtual_SetStylePreprocessor(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QsciLexerCSharp_SetStylePreprocessor
func miqt_exec_callback_QsciLexerCSharp_SetStylePreprocessor(self *C.QsciLexerCSharp, cb C.intptr_t, style C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(style bool), style bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(style)

	gofunc((&QsciLexerCSharp{h: self}).callVirtualBase_SetStylePreprocessor, slotval1)

}

// Delete this object from C++ memory.
func (this *QsciLexerCSharp) Delete() {
	C.QsciLexerCSharp_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QsciLexerCSharp) GoGC() {
	runtime.SetFinalizer(this, func(this *QsciLexerCSharp) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
