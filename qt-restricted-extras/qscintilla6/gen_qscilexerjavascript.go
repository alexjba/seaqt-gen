package qscintilla6

/*

#include "gen_qscilexerjavascript.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt6"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QsciLexerJavaScript struct {
	h *C.QsciLexerJavaScript
	*QsciLexerCPP
}

func (this *QsciLexerJavaScript) cPointer() *C.QsciLexerJavaScript {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QsciLexerJavaScript) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQsciLexerJavaScript constructs the type using only CGO pointers.
func newQsciLexerJavaScript(h *C.QsciLexerJavaScript) *QsciLexerJavaScript {
	if h == nil {
		return nil
	}
	var outptr_QsciLexerCPP *C.QsciLexerCPP = nil
	C.QsciLexerJavaScript_virtbase(h, &outptr_QsciLexerCPP)

	return &QsciLexerJavaScript{h: h,
		QsciLexerCPP: newQsciLexerCPP(outptr_QsciLexerCPP)}
}

// UnsafeNewQsciLexerJavaScript constructs the type using only unsafe pointers.
func UnsafeNewQsciLexerJavaScript(h unsafe.Pointer) *QsciLexerJavaScript {
	return newQsciLexerJavaScript((*C.QsciLexerJavaScript)(h))
}

// NewQsciLexerJavaScript constructs a new QsciLexerJavaScript object.
func NewQsciLexerJavaScript() *QsciLexerJavaScript {

	return newQsciLexerJavaScript(C.QsciLexerJavaScript_new())
}

// NewQsciLexerJavaScript2 constructs a new QsciLexerJavaScript object.
func NewQsciLexerJavaScript2(parent *qt6.QObject) *QsciLexerJavaScript {

	return newQsciLexerJavaScript(C.QsciLexerJavaScript_new2((*C.QObject)(parent.UnsafePointer())))
}

func (this *QsciLexerJavaScript) MetaObject() *qt6.QMetaObject {
	return qt6.UnsafeNewQMetaObject(unsafe.Pointer(C.QsciLexerJavaScript_MetaObject(this.h)))
}

func (this *QsciLexerJavaScript) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QsciLexerJavaScript_Metacast(this.h, param1_Cstring))
}

func QsciLexerJavaScript_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QsciLexerJavaScript_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QsciLexerJavaScript) Language() string {
	_ret := C.QsciLexerJavaScript_Language(this.h)
	return C.GoString(_ret)
}

func (this *QsciLexerJavaScript) DefaultColor(style int) *qt6.QColor {
	_goptr := qt6.UnsafeNewQColor(unsafe.Pointer(C.QsciLexerJavaScript_DefaultColor(this.h, (C.int)(style))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QsciLexerJavaScript) DefaultEolFill(style int) bool {
	return (bool)(C.QsciLexerJavaScript_DefaultEolFill(this.h, (C.int)(style)))
}

func (this *QsciLexerJavaScript) DefaultFont(style int) *qt6.QFont {
	_goptr := qt6.UnsafeNewQFont(unsafe.Pointer(C.QsciLexerJavaScript_DefaultFont(this.h, (C.int)(style))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QsciLexerJavaScript) DefaultPaper(style int) *qt6.QColor {
	_goptr := qt6.UnsafeNewQColor(unsafe.Pointer(C.QsciLexerJavaScript_DefaultPaper(this.h, (C.int)(style))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QsciLexerJavaScript) Keywords(set int) string {
	_ret := C.QsciLexerJavaScript_Keywords(this.h, (C.int)(set))
	return C.GoString(_ret)
}

func (this *QsciLexerJavaScript) Description(style int) string {
	var _ms C.struct_miqt_string = C.QsciLexerJavaScript_Description(this.h, (C.int)(style))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QsciLexerJavaScript_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QsciLexerJavaScript_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QsciLexerJavaScript_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QsciLexerJavaScript_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QsciLexerJavaScript) callVirtualBase_SetFoldAtElse(fold bool) {

	C.QsciLexerJavaScript_virtualbase_SetFoldAtElse(unsafe.Pointer(this.h), (C.bool)(fold))

}
func (this *QsciLexerJavaScript) OnSetFoldAtElse(slot func(super func(fold bool), fold bool)) {
	ok := C.QsciLexerJavaScript_override_virtual_SetFoldAtElse(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerJavaScript_SetFoldAtElse
func miqt_exec_callback_QsciLexerJavaScript_SetFoldAtElse(self *C.QsciLexerJavaScript, cb C.intptr_t, fold C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(fold bool), fold bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(fold)

	gofunc((&QsciLexerJavaScript{h: self}).callVirtualBase_SetFoldAtElse, slotval1)

}

func (this *QsciLexerJavaScript) callVirtualBase_SetFoldComments(fold bool) {

	C.QsciLexerJavaScript_virtualbase_SetFoldComments(unsafe.Pointer(this.h), (C.bool)(fold))

}
func (this *QsciLexerJavaScript) OnSetFoldComments(slot func(super func(fold bool), fold bool)) {
	ok := C.QsciLexerJavaScript_override_virtual_SetFoldComments(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerJavaScript_SetFoldComments
func miqt_exec_callback_QsciLexerJavaScript_SetFoldComments(self *C.QsciLexerJavaScript, cb C.intptr_t, fold C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(fold bool), fold bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(fold)

	gofunc((&QsciLexerJavaScript{h: self}).callVirtualBase_SetFoldComments, slotval1)

}

func (this *QsciLexerJavaScript) callVirtualBase_SetFoldCompact(fold bool) {

	C.QsciLexerJavaScript_virtualbase_SetFoldCompact(unsafe.Pointer(this.h), (C.bool)(fold))

}
func (this *QsciLexerJavaScript) OnSetFoldCompact(slot func(super func(fold bool), fold bool)) {
	ok := C.QsciLexerJavaScript_override_virtual_SetFoldCompact(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerJavaScript_SetFoldCompact
func miqt_exec_callback_QsciLexerJavaScript_SetFoldCompact(self *C.QsciLexerJavaScript, cb C.intptr_t, fold C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(fold bool), fold bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(fold)

	gofunc((&QsciLexerJavaScript{h: self}).callVirtualBase_SetFoldCompact, slotval1)

}

func (this *QsciLexerJavaScript) callVirtualBase_SetFoldPreprocessor(fold bool) {

	C.QsciLexerJavaScript_virtualbase_SetFoldPreprocessor(unsafe.Pointer(this.h), (C.bool)(fold))

}
func (this *QsciLexerJavaScript) OnSetFoldPreprocessor(slot func(super func(fold bool), fold bool)) {
	ok := C.QsciLexerJavaScript_override_virtual_SetFoldPreprocessor(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerJavaScript_SetFoldPreprocessor
func miqt_exec_callback_QsciLexerJavaScript_SetFoldPreprocessor(self *C.QsciLexerJavaScript, cb C.intptr_t, fold C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(fold bool), fold bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(fold)

	gofunc((&QsciLexerJavaScript{h: self}).callVirtualBase_SetFoldPreprocessor, slotval1)

}

func (this *QsciLexerJavaScript) callVirtualBase_SetStylePreprocessor(style bool) {

	C.QsciLexerJavaScript_virtualbase_SetStylePreprocessor(unsafe.Pointer(this.h), (C.bool)(style))

}
func (this *QsciLexerJavaScript) OnSetStylePreprocessor(slot func(super func(style bool), style bool)) {
	ok := C.QsciLexerJavaScript_override_virtual_SetStylePreprocessor(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerJavaScript_SetStylePreprocessor
func miqt_exec_callback_QsciLexerJavaScript_SetStylePreprocessor(self *C.QsciLexerJavaScript, cb C.intptr_t, style C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(style bool), style bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(style)

	gofunc((&QsciLexerJavaScript{h: self}).callVirtualBase_SetStylePreprocessor, slotval1)

}

// Delete this object from C++ memory.
func (this *QsciLexerJavaScript) Delete() {
	C.QsciLexerJavaScript_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QsciLexerJavaScript) GoGC() {
	runtime.SetFinalizer(this, func(this *QsciLexerJavaScript) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
