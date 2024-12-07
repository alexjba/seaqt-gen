package qt

/*

#include "gen_qobjectdefs.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QMetaObject__Call int

const (
	QMetaObject__InvokeMetaMethod               QMetaObject__Call = 0
	QMetaObject__ReadProperty                   QMetaObject__Call = 1
	QMetaObject__WriteProperty                  QMetaObject__Call = 2
	QMetaObject__ResetProperty                  QMetaObject__Call = 3
	QMetaObject__QueryPropertyDesignable        QMetaObject__Call = 4
	QMetaObject__QueryPropertyScriptable        QMetaObject__Call = 5
	QMetaObject__QueryPropertyStored            QMetaObject__Call = 6
	QMetaObject__QueryPropertyEditable          QMetaObject__Call = 7
	QMetaObject__QueryPropertyUser              QMetaObject__Call = 8
	QMetaObject__CreateInstance                 QMetaObject__Call = 9
	QMetaObject__IndexOfMethod                  QMetaObject__Call = 10
	QMetaObject__RegisterPropertyMetaType       QMetaObject__Call = 11
	QMetaObject__RegisterMethodArgumentMetaType QMetaObject__Call = 12
)

type QGenericArgument struct {
	h          *C.QGenericArgument
	isSubclass bool
}

func (this *QGenericArgument) cPointer() *C.QGenericArgument {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QGenericArgument) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQGenericArgument constructs the type using only CGO pointers.
func newQGenericArgument(h *C.QGenericArgument) *QGenericArgument {
	if h == nil {
		return nil
	}
	return &QGenericArgument{h: h}
}

// UnsafeNewQGenericArgument constructs the type using only unsafe pointers.
func UnsafeNewQGenericArgument(h unsafe.Pointer) *QGenericArgument {
	if h == nil {
		return nil
	}

	return &QGenericArgument{h: (*C.QGenericArgument)(h)}
}

// NewQGenericArgument constructs a new QGenericArgument object.
func NewQGenericArgument() *QGenericArgument {
	var outptr_QGenericArgument *C.QGenericArgument = nil

	C.QGenericArgument_new(&outptr_QGenericArgument)
	ret := newQGenericArgument(outptr_QGenericArgument)
	ret.isSubclass = true
	return ret
}

// NewQGenericArgument2 constructs a new QGenericArgument object.
func NewQGenericArgument2(param1 *QGenericArgument) *QGenericArgument {
	var outptr_QGenericArgument *C.QGenericArgument = nil

	C.QGenericArgument_new2(param1.cPointer(), &outptr_QGenericArgument)
	ret := newQGenericArgument(outptr_QGenericArgument)
	ret.isSubclass = true
	return ret
}

// NewQGenericArgument3 constructs a new QGenericArgument object.
func NewQGenericArgument3(aName string) *QGenericArgument {
	aName_Cstring := C.CString(aName)
	defer C.free(unsafe.Pointer(aName_Cstring))
	var outptr_QGenericArgument *C.QGenericArgument = nil

	C.QGenericArgument_new3(aName_Cstring, &outptr_QGenericArgument)
	ret := newQGenericArgument(outptr_QGenericArgument)
	ret.isSubclass = true
	return ret
}

// NewQGenericArgument4 constructs a new QGenericArgument object.
func NewQGenericArgument4(aName string, aData unsafe.Pointer) *QGenericArgument {
	aName_Cstring := C.CString(aName)
	defer C.free(unsafe.Pointer(aName_Cstring))
	var outptr_QGenericArgument *C.QGenericArgument = nil

	C.QGenericArgument_new4(aName_Cstring, aData, &outptr_QGenericArgument)
	ret := newQGenericArgument(outptr_QGenericArgument)
	ret.isSubclass = true
	return ret
}

func (this *QGenericArgument) Data() unsafe.Pointer {
	return (unsafe.Pointer)(C.QGenericArgument_Data(this.h))
}

func (this *QGenericArgument) Name() string {
	_ret := C.QGenericArgument_Name(this.h)
	return C.GoString(_ret)
}

// Delete this object from C++ memory.
func (this *QGenericArgument) Delete() {
	C.QGenericArgument_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QGenericArgument) GoGC() {
	runtime.SetFinalizer(this, func(this *QGenericArgument) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QGenericReturnArgument struct {
	h          *C.QGenericReturnArgument
	isSubclass bool
	*QGenericArgument
}

func (this *QGenericReturnArgument) cPointer() *C.QGenericReturnArgument {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QGenericReturnArgument) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQGenericReturnArgument constructs the type using only CGO pointers.
func newQGenericReturnArgument(h *C.QGenericReturnArgument, h_QGenericArgument *C.QGenericArgument) *QGenericReturnArgument {
	if h == nil {
		return nil
	}
	return &QGenericReturnArgument{h: h,
		QGenericArgument: newQGenericArgument(h_QGenericArgument)}
}

// UnsafeNewQGenericReturnArgument constructs the type using only unsafe pointers.
func UnsafeNewQGenericReturnArgument(h unsafe.Pointer, h_QGenericArgument unsafe.Pointer) *QGenericReturnArgument {
	if h == nil {
		return nil
	}

	return &QGenericReturnArgument{h: (*C.QGenericReturnArgument)(h),
		QGenericArgument: UnsafeNewQGenericArgument(h_QGenericArgument)}
}

// NewQGenericReturnArgument constructs a new QGenericReturnArgument object.
func NewQGenericReturnArgument() *QGenericReturnArgument {
	var outptr_QGenericReturnArgument *C.QGenericReturnArgument = nil
	var outptr_QGenericArgument *C.QGenericArgument = nil

	C.QGenericReturnArgument_new(&outptr_QGenericReturnArgument, &outptr_QGenericArgument)
	ret := newQGenericReturnArgument(outptr_QGenericReturnArgument, outptr_QGenericArgument)
	ret.isSubclass = true
	return ret
}

// NewQGenericReturnArgument2 constructs a new QGenericReturnArgument object.
func NewQGenericReturnArgument2(param1 *QGenericReturnArgument) *QGenericReturnArgument {
	var outptr_QGenericReturnArgument *C.QGenericReturnArgument = nil
	var outptr_QGenericArgument *C.QGenericArgument = nil

	C.QGenericReturnArgument_new2(param1.cPointer(), &outptr_QGenericReturnArgument, &outptr_QGenericArgument)
	ret := newQGenericReturnArgument(outptr_QGenericReturnArgument, outptr_QGenericArgument)
	ret.isSubclass = true
	return ret
}

// NewQGenericReturnArgument3 constructs a new QGenericReturnArgument object.
func NewQGenericReturnArgument3(aName string) *QGenericReturnArgument {
	aName_Cstring := C.CString(aName)
	defer C.free(unsafe.Pointer(aName_Cstring))
	var outptr_QGenericReturnArgument *C.QGenericReturnArgument = nil
	var outptr_QGenericArgument *C.QGenericArgument = nil

	C.QGenericReturnArgument_new3(aName_Cstring, &outptr_QGenericReturnArgument, &outptr_QGenericArgument)
	ret := newQGenericReturnArgument(outptr_QGenericReturnArgument, outptr_QGenericArgument)
	ret.isSubclass = true
	return ret
}

// NewQGenericReturnArgument4 constructs a new QGenericReturnArgument object.
func NewQGenericReturnArgument4(aName string, aData unsafe.Pointer) *QGenericReturnArgument {
	aName_Cstring := C.CString(aName)
	defer C.free(unsafe.Pointer(aName_Cstring))
	var outptr_QGenericReturnArgument *C.QGenericReturnArgument = nil
	var outptr_QGenericArgument *C.QGenericArgument = nil

	C.QGenericReturnArgument_new4(aName_Cstring, aData, &outptr_QGenericReturnArgument, &outptr_QGenericArgument)
	ret := newQGenericReturnArgument(outptr_QGenericReturnArgument, outptr_QGenericArgument)
	ret.isSubclass = true
	return ret
}

// Delete this object from C++ memory.
func (this *QGenericReturnArgument) Delete() {
	C.QGenericReturnArgument_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QGenericReturnArgument) GoGC() {
	runtime.SetFinalizer(this, func(this *QGenericReturnArgument) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QMetaObject struct {
	h          *C.QMetaObject
	isSubclass bool
}

func (this *QMetaObject) cPointer() *C.QMetaObject {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QMetaObject) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQMetaObject constructs the type using only CGO pointers.
func newQMetaObject(h *C.QMetaObject) *QMetaObject {
	if h == nil {
		return nil
	}
	return &QMetaObject{h: h}
}

// UnsafeNewQMetaObject constructs the type using only unsafe pointers.
func UnsafeNewQMetaObject(h unsafe.Pointer) *QMetaObject {
	if h == nil {
		return nil
	}

	return &QMetaObject{h: (*C.QMetaObject)(h)}
}

// NewQMetaObject constructs a new QMetaObject object.
func NewQMetaObject() *QMetaObject {
	var outptr_QMetaObject *C.QMetaObject = nil

	C.QMetaObject_new(&outptr_QMetaObject)
	ret := newQMetaObject(outptr_QMetaObject)
	ret.isSubclass = true
	return ret
}

// NewQMetaObject2 constructs a new QMetaObject object.
func NewQMetaObject2(param1 *QMetaObject) *QMetaObject {
	var outptr_QMetaObject *C.QMetaObject = nil

	C.QMetaObject_new2(param1.cPointer(), &outptr_QMetaObject)
	ret := newQMetaObject(outptr_QMetaObject)
	ret.isSubclass = true
	return ret
}

func (this *QMetaObject) ClassName() string {
	_ret := C.QMetaObject_ClassName(this.h)
	return C.GoString(_ret)
}

func (this *QMetaObject) SuperClass() *QMetaObject {
	return newQMetaObject(C.QMetaObject_SuperClass(this.h))
}

func (this *QMetaObject) Inherits(metaObject *QMetaObject) bool {
	return (bool)(C.QMetaObject_Inherits(this.h, metaObject.cPointer()))
}

func (this *QMetaObject) Cast(obj *QObject) *QObject {
	return newQObject(C.QMetaObject_Cast(this.h, obj.cPointer()))
}

func (this *QMetaObject) CastWithObj(obj *QObject) *QObject {
	return newQObject(C.QMetaObject_CastWithObj(this.h, obj.cPointer()))
}

func (this *QMetaObject) Tr(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMetaObject_Tr(this.h, s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMetaObject) MethodOffset() int {
	return (int)(C.QMetaObject_MethodOffset(this.h))
}

func (this *QMetaObject) EnumeratorOffset() int {
	return (int)(C.QMetaObject_EnumeratorOffset(this.h))
}

func (this *QMetaObject) PropertyOffset() int {
	return (int)(C.QMetaObject_PropertyOffset(this.h))
}

func (this *QMetaObject) ClassInfoOffset() int {
	return (int)(C.QMetaObject_ClassInfoOffset(this.h))
}

func (this *QMetaObject) ConstructorCount() int {
	return (int)(C.QMetaObject_ConstructorCount(this.h))
}

func (this *QMetaObject) MethodCount() int {
	return (int)(C.QMetaObject_MethodCount(this.h))
}

func (this *QMetaObject) EnumeratorCount() int {
	return (int)(C.QMetaObject_EnumeratorCount(this.h))
}

func (this *QMetaObject) PropertyCount() int {
	return (int)(C.QMetaObject_PropertyCount(this.h))
}

func (this *QMetaObject) ClassInfoCount() int {
	return (int)(C.QMetaObject_ClassInfoCount(this.h))
}

func (this *QMetaObject) IndexOfConstructor(constructor string) int {
	constructor_Cstring := C.CString(constructor)
	defer C.free(unsafe.Pointer(constructor_Cstring))
	return (int)(C.QMetaObject_IndexOfConstructor(this.h, constructor_Cstring))
}

func (this *QMetaObject) IndexOfMethod(method string) int {
	method_Cstring := C.CString(method)
	defer C.free(unsafe.Pointer(method_Cstring))
	return (int)(C.QMetaObject_IndexOfMethod(this.h, method_Cstring))
}

func (this *QMetaObject) IndexOfSignal(signal string) int {
	signal_Cstring := C.CString(signal)
	defer C.free(unsafe.Pointer(signal_Cstring))
	return (int)(C.QMetaObject_IndexOfSignal(this.h, signal_Cstring))
}

func (this *QMetaObject) IndexOfSlot(slot string) int {
	slot_Cstring := C.CString(slot)
	defer C.free(unsafe.Pointer(slot_Cstring))
	return (int)(C.QMetaObject_IndexOfSlot(this.h, slot_Cstring))
}

func (this *QMetaObject) IndexOfEnumerator(name string) int {
	name_Cstring := C.CString(name)
	defer C.free(unsafe.Pointer(name_Cstring))
	return (int)(C.QMetaObject_IndexOfEnumerator(this.h, name_Cstring))
}

func (this *QMetaObject) IndexOfProperty(name string) int {
	name_Cstring := C.CString(name)
	defer C.free(unsafe.Pointer(name_Cstring))
	return (int)(C.QMetaObject_IndexOfProperty(this.h, name_Cstring))
}

func (this *QMetaObject) IndexOfClassInfo(name string) int {
	name_Cstring := C.CString(name)
	defer C.free(unsafe.Pointer(name_Cstring))
	return (int)(C.QMetaObject_IndexOfClassInfo(this.h, name_Cstring))
}

func (this *QMetaObject) Constructor(index int) *QMetaMethod {
	_goptr := newQMetaMethod(C.QMetaObject_Constructor(this.h, (C.int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaObject) Method(index int) *QMetaMethod {
	_goptr := newQMetaMethod(C.QMetaObject_Method(this.h, (C.int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaObject) Enumerator(index int) *QMetaEnum {
	_goptr := newQMetaEnum(C.QMetaObject_Enumerator(this.h, (C.int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaObject) Property(index int) *QMetaProperty {
	_goptr := newQMetaProperty(C.QMetaObject_Property(this.h, (C.int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaObject) ClassInfo(index int) *QMetaClassInfo {
	_goptr := newQMetaClassInfo(C.QMetaObject_ClassInfo(this.h, (C.int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaObject) UserProperty() *QMetaProperty {
	_goptr := newQMetaProperty(C.QMetaObject_UserProperty(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QMetaObject_CheckConnectArgs(signal string, method string) bool {
	signal_Cstring := C.CString(signal)
	defer C.free(unsafe.Pointer(signal_Cstring))
	method_Cstring := C.CString(method)
	defer C.free(unsafe.Pointer(method_Cstring))
	return (bool)(C.QMetaObject_CheckConnectArgs(signal_Cstring, method_Cstring))
}

func QMetaObject_CheckConnectArgs2(signal *QMetaMethod, method *QMetaMethod) bool {
	return (bool)(C.QMetaObject_CheckConnectArgs2(signal.cPointer(), method.cPointer()))
}

func QMetaObject_NormalizedSignature(method string) []byte {
	method_Cstring := C.CString(method)
	defer C.free(unsafe.Pointer(method_Cstring))
	var _bytearray C.struct_miqt_string = C.QMetaObject_NormalizedSignature(method_Cstring)
	_ret := C.GoBytes(unsafe.Pointer(_bytearray.data), C.int(int64(_bytearray.len)))
	C.free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QMetaObject_NormalizedType(typeVal string) []byte {
	typeVal_Cstring := C.CString(typeVal)
	defer C.free(unsafe.Pointer(typeVal_Cstring))
	var _bytearray C.struct_miqt_string = C.QMetaObject_NormalizedType(typeVal_Cstring)
	_ret := C.GoBytes(unsafe.Pointer(_bytearray.data), C.int(int64(_bytearray.len)))
	C.free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QMetaObject_Connect(sender *QObject, signal_index int, receiver *QObject, method_index int) *QMetaObject__Connection {
	_goptr := newQMetaObject__Connection(C.QMetaObject_Connect(sender.cPointer(), (C.int)(signal_index), receiver.cPointer(), (C.int)(method_index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QMetaObject_Disconnect(sender *QObject, signal_index int, receiver *QObject, method_index int) bool {
	return (bool)(C.QMetaObject_Disconnect(sender.cPointer(), (C.int)(signal_index), receiver.cPointer(), (C.int)(method_index)))
}

func QMetaObject_DisconnectOne(sender *QObject, signal_index int, receiver *QObject, method_index int) bool {
	return (bool)(C.QMetaObject_DisconnectOne(sender.cPointer(), (C.int)(signal_index), receiver.cPointer(), (C.int)(method_index)))
}

func QMetaObject_ConnectSlotsByName(o *QObject) {
	C.QMetaObject_ConnectSlotsByName(o.cPointer())
}

func QMetaObject_InvokeMethod(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod(obj.cPointer(), member_Cstring, (C.int)(param3), retVal.cPointer()))
}

func QMetaObject_InvokeMethod2(obj *QObject, member string, retVal QGenericReturnArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod2(obj.cPointer(), member_Cstring, retVal.cPointer()))
}

func QMetaObject_InvokeMethod3(obj *QObject, member string, typeVal ConnectionType) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod3(obj.cPointer(), member_Cstring, (C.int)(typeVal)))
}

func QMetaObject_InvokeMethod4(obj *QObject, member string) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod4(obj.cPointer(), member_Cstring))
}

func (this *QMetaObject) NewInstance() *QObject {
	return newQObject(C.QMetaObject_NewInstance(this.h))
}

func (this *QMetaObject) Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMetaObject_Tr3(this.h, s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMetaObject_Connect5(sender *QObject, signal_index int, receiver *QObject, method_index int, typeVal int) *QMetaObject__Connection {
	_goptr := newQMetaObject__Connection(C.QMetaObject_Connect5(sender.cPointer(), (C.int)(signal_index), receiver.cPointer(), (C.int)(method_index), (C.int)(typeVal)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QMetaObject_Connect6(sender *QObject, signal_index int, receiver *QObject, method_index int, typeVal int, types *int) *QMetaObject__Connection {
	_goptr := newQMetaObject__Connection(C.QMetaObject_Connect6(sender.cPointer(), (C.int)(signal_index), receiver.cPointer(), (C.int)(method_index), (C.int)(typeVal), (*C.int)(unsafe.Pointer(types))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QMetaObject_InvokeMethod5(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument, val0 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod5(obj.cPointer(), member_Cstring, (C.int)(param3), retVal.cPointer(), val0.cPointer()))
}

func QMetaObject_InvokeMethod6(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod6(obj.cPointer(), member_Cstring, (C.int)(param3), retVal.cPointer(), val0.cPointer(), val1.cPointer()))
}

func QMetaObject_InvokeMethod7(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod7(obj.cPointer(), member_Cstring, (C.int)(param3), retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer()))
}

func QMetaObject_InvokeMethod8(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod8(obj.cPointer(), member_Cstring, (C.int)(param3), retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer()))
}

func QMetaObject_InvokeMethod9(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod9(obj.cPointer(), member_Cstring, (C.int)(param3), retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer()))
}

func QMetaObject_InvokeMethod10(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod10(obj.cPointer(), member_Cstring, (C.int)(param3), retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer()))
}

func QMetaObject_InvokeMethod11(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod11(obj.cPointer(), member_Cstring, (C.int)(param3), retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer()))
}

func QMetaObject_InvokeMethod12(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod12(obj.cPointer(), member_Cstring, (C.int)(param3), retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer()))
}

func QMetaObject_InvokeMethod13(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod13(obj.cPointer(), member_Cstring, (C.int)(param3), retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer()))
}

func QMetaObject_InvokeMethod14(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument, val9 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod14(obj.cPointer(), member_Cstring, (C.int)(param3), retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer(), val9.cPointer()))
}

func QMetaObject_InvokeMethod42(obj *QObject, member string, retVal QGenericReturnArgument, val0 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod42(obj.cPointer(), member_Cstring, retVal.cPointer(), val0.cPointer()))
}

func QMetaObject_InvokeMethod52(obj *QObject, member string, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod52(obj.cPointer(), member_Cstring, retVal.cPointer(), val0.cPointer(), val1.cPointer()))
}

func QMetaObject_InvokeMethod62(obj *QObject, member string, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod62(obj.cPointer(), member_Cstring, retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer()))
}

func QMetaObject_InvokeMethod72(obj *QObject, member string, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod72(obj.cPointer(), member_Cstring, retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer()))
}

func QMetaObject_InvokeMethod82(obj *QObject, member string, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod82(obj.cPointer(), member_Cstring, retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer()))
}

func QMetaObject_InvokeMethod92(obj *QObject, member string, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod92(obj.cPointer(), member_Cstring, retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer()))
}

func QMetaObject_InvokeMethod102(obj *QObject, member string, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod102(obj.cPointer(), member_Cstring, retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer()))
}

func QMetaObject_InvokeMethod112(obj *QObject, member string, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod112(obj.cPointer(), member_Cstring, retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer()))
}

func QMetaObject_InvokeMethod122(obj *QObject, member string, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod122(obj.cPointer(), member_Cstring, retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer()))
}

func QMetaObject_InvokeMethod132(obj *QObject, member string, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument, val9 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod132(obj.cPointer(), member_Cstring, retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer(), val9.cPointer()))
}

func QMetaObject_InvokeMethod43(obj *QObject, member string, typeVal ConnectionType, val0 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod43(obj.cPointer(), member_Cstring, (C.int)(typeVal), val0.cPointer()))
}

func QMetaObject_InvokeMethod53(obj *QObject, member string, typeVal ConnectionType, val0 QGenericArgument, val1 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod53(obj.cPointer(), member_Cstring, (C.int)(typeVal), val0.cPointer(), val1.cPointer()))
}

func QMetaObject_InvokeMethod63(obj *QObject, member string, typeVal ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod63(obj.cPointer(), member_Cstring, (C.int)(typeVal), val0.cPointer(), val1.cPointer(), val2.cPointer()))
}

func QMetaObject_InvokeMethod73(obj *QObject, member string, typeVal ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod73(obj.cPointer(), member_Cstring, (C.int)(typeVal), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer()))
}

func QMetaObject_InvokeMethod83(obj *QObject, member string, typeVal ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod83(obj.cPointer(), member_Cstring, (C.int)(typeVal), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer()))
}

func QMetaObject_InvokeMethod93(obj *QObject, member string, typeVal ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod93(obj.cPointer(), member_Cstring, (C.int)(typeVal), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer()))
}

func QMetaObject_InvokeMethod103(obj *QObject, member string, typeVal ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod103(obj.cPointer(), member_Cstring, (C.int)(typeVal), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer()))
}

func QMetaObject_InvokeMethod113(obj *QObject, member string, typeVal ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod113(obj.cPointer(), member_Cstring, (C.int)(typeVal), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer()))
}

func QMetaObject_InvokeMethod123(obj *QObject, member string, typeVal ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod123(obj.cPointer(), member_Cstring, (C.int)(typeVal), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer()))
}

func QMetaObject_InvokeMethod133(obj *QObject, member string, typeVal ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument, val9 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod133(obj.cPointer(), member_Cstring, (C.int)(typeVal), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer(), val9.cPointer()))
}

func QMetaObject_InvokeMethod32(obj *QObject, member string, val0 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod32(obj.cPointer(), member_Cstring, val0.cPointer()))
}

func QMetaObject_InvokeMethod44(obj *QObject, member string, val0 QGenericArgument, val1 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod44(obj.cPointer(), member_Cstring, val0.cPointer(), val1.cPointer()))
}

func QMetaObject_InvokeMethod54(obj *QObject, member string, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod54(obj.cPointer(), member_Cstring, val0.cPointer(), val1.cPointer(), val2.cPointer()))
}

func QMetaObject_InvokeMethod64(obj *QObject, member string, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod64(obj.cPointer(), member_Cstring, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer()))
}

func QMetaObject_InvokeMethod74(obj *QObject, member string, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod74(obj.cPointer(), member_Cstring, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer()))
}

func QMetaObject_InvokeMethod84(obj *QObject, member string, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod84(obj.cPointer(), member_Cstring, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer()))
}

func QMetaObject_InvokeMethod94(obj *QObject, member string, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod94(obj.cPointer(), member_Cstring, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer()))
}

func QMetaObject_InvokeMethod104(obj *QObject, member string, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod104(obj.cPointer(), member_Cstring, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer()))
}

func QMetaObject_InvokeMethod114(obj *QObject, member string, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod114(obj.cPointer(), member_Cstring, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer()))
}

func QMetaObject_InvokeMethod124(obj *QObject, member string, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument, val9 QGenericArgument) bool {
	member_Cstring := C.CString(member)
	defer C.free(unsafe.Pointer(member_Cstring))
	return (bool)(C.QMetaObject_InvokeMethod124(obj.cPointer(), member_Cstring, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer(), val9.cPointer()))
}

func (this *QMetaObject) NewInstance1(val0 QGenericArgument) *QObject {
	return newQObject(C.QMetaObject_NewInstance1(this.h, val0.cPointer()))
}

func (this *QMetaObject) NewInstance2(val0 QGenericArgument, val1 QGenericArgument) *QObject {
	return newQObject(C.QMetaObject_NewInstance2(this.h, val0.cPointer(), val1.cPointer()))
}

func (this *QMetaObject) NewInstance3(val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument) *QObject {
	return newQObject(C.QMetaObject_NewInstance3(this.h, val0.cPointer(), val1.cPointer(), val2.cPointer()))
}

func (this *QMetaObject) NewInstance4(val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument) *QObject {
	return newQObject(C.QMetaObject_NewInstance4(this.h, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer()))
}

func (this *QMetaObject) NewInstance5(val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument) *QObject {
	return newQObject(C.QMetaObject_NewInstance5(this.h, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer()))
}

func (this *QMetaObject) NewInstance6(val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument) *QObject {
	return newQObject(C.QMetaObject_NewInstance6(this.h, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer()))
}

func (this *QMetaObject) NewInstance7(val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument) *QObject {
	return newQObject(C.QMetaObject_NewInstance7(this.h, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer()))
}

func (this *QMetaObject) NewInstance8(val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument) *QObject {
	return newQObject(C.QMetaObject_NewInstance8(this.h, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer()))
}

func (this *QMetaObject) NewInstance9(val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument) *QObject {
	return newQObject(C.QMetaObject_NewInstance9(this.h, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer()))
}

func (this *QMetaObject) NewInstance10(val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument, val9 QGenericArgument) *QObject {
	return newQObject(C.QMetaObject_NewInstance10(this.h, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer(), val9.cPointer()))
}

// Delete this object from C++ memory.
func (this *QMetaObject) Delete() {
	C.QMetaObject_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QMetaObject) GoGC() {
	runtime.SetFinalizer(this, func(this *QMetaObject) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QMetaObject__Connection struct {
	h          *C.QMetaObject__Connection
	isSubclass bool
}

func (this *QMetaObject__Connection) cPointer() *C.QMetaObject__Connection {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QMetaObject__Connection) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQMetaObject__Connection constructs the type using only CGO pointers.
func newQMetaObject__Connection(h *C.QMetaObject__Connection) *QMetaObject__Connection {
	if h == nil {
		return nil
	}
	return &QMetaObject__Connection{h: h}
}

// UnsafeNewQMetaObject__Connection constructs the type using only unsafe pointers.
func UnsafeNewQMetaObject__Connection(h unsafe.Pointer) *QMetaObject__Connection {
	if h == nil {
		return nil
	}

	return &QMetaObject__Connection{h: (*C.QMetaObject__Connection)(h)}
}

// NewQMetaObject__Connection constructs a new QMetaObject::Connection object.
func NewQMetaObject__Connection() *QMetaObject__Connection {
	var outptr_QMetaObject__Connection *C.QMetaObject__Connection = nil

	C.QMetaObject__Connection_new(&outptr_QMetaObject__Connection)
	ret := newQMetaObject__Connection(outptr_QMetaObject__Connection)
	ret.isSubclass = true
	return ret
}

// NewQMetaObject__Connection2 constructs a new QMetaObject::Connection object.
func NewQMetaObject__Connection2(other *QMetaObject__Connection) *QMetaObject__Connection {
	var outptr_QMetaObject__Connection *C.QMetaObject__Connection = nil

	C.QMetaObject__Connection_new2(other.cPointer(), &outptr_QMetaObject__Connection)
	ret := newQMetaObject__Connection(outptr_QMetaObject__Connection)
	ret.isSubclass = true
	return ret
}

func (this *QMetaObject__Connection) OperatorAssign(other *QMetaObject__Connection) {
	C.QMetaObject__Connection_OperatorAssign(this.h, other.cPointer())
}

// Delete this object from C++ memory.
func (this *QMetaObject__Connection) Delete() {
	C.QMetaObject__Connection_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QMetaObject__Connection) GoGC() {
	runtime.SetFinalizer(this, func(this *QMetaObject__Connection) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QMetaObject__SuperData struct {
	h          *C.QMetaObject__SuperData
	isSubclass bool
}

func (this *QMetaObject__SuperData) cPointer() *C.QMetaObject__SuperData {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QMetaObject__SuperData) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQMetaObject__SuperData constructs the type using only CGO pointers.
func newQMetaObject__SuperData(h *C.QMetaObject__SuperData) *QMetaObject__SuperData {
	if h == nil {
		return nil
	}
	return &QMetaObject__SuperData{h: h}
}

// UnsafeNewQMetaObject__SuperData constructs the type using only unsafe pointers.
func UnsafeNewQMetaObject__SuperData(h unsafe.Pointer) *QMetaObject__SuperData {
	if h == nil {
		return nil
	}

	return &QMetaObject__SuperData{h: (*C.QMetaObject__SuperData)(h)}
}

// NewQMetaObject__SuperData constructs a new QMetaObject::SuperData object.
func NewQMetaObject__SuperData() *QMetaObject__SuperData {
	var outptr_QMetaObject__SuperData *C.QMetaObject__SuperData = nil

	C.QMetaObject__SuperData_new(&outptr_QMetaObject__SuperData)
	ret := newQMetaObject__SuperData(outptr_QMetaObject__SuperData)
	ret.isSubclass = true
	return ret
}

// NewQMetaObject__SuperData2 constructs a new QMetaObject::SuperData object.
func NewQMetaObject__SuperData2(mo *QMetaObject) *QMetaObject__SuperData {
	var outptr_QMetaObject__SuperData *C.QMetaObject__SuperData = nil

	C.QMetaObject__SuperData_new2(mo.cPointer(), &outptr_QMetaObject__SuperData)
	ret := newQMetaObject__SuperData(outptr_QMetaObject__SuperData)
	ret.isSubclass = true
	return ret
}

// NewQMetaObject__SuperData3 constructs a new QMetaObject::SuperData object.
func NewQMetaObject__SuperData3(param1 *QMetaObject__SuperData) *QMetaObject__SuperData {
	var outptr_QMetaObject__SuperData *C.QMetaObject__SuperData = nil

	C.QMetaObject__SuperData_new3(param1.cPointer(), &outptr_QMetaObject__SuperData)
	ret := newQMetaObject__SuperData(outptr_QMetaObject__SuperData)
	ret.isSubclass = true
	return ret
}

func (this *QMetaObject__SuperData) OperatorMinusGreater() *QMetaObject {
	return newQMetaObject(C.QMetaObject__SuperData_OperatorMinusGreater(this.h))
}

func (this *QMetaObject__SuperData) OperatorAssign(param1 *QMetaObject__SuperData) {
	C.QMetaObject__SuperData_OperatorAssign(this.h, param1.cPointer())
}

// Delete this object from C++ memory.
func (this *QMetaObject__SuperData) Delete() {
	C.QMetaObject__SuperData_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QMetaObject__SuperData) GoGC() {
	runtime.SetFinalizer(this, func(this *QMetaObject__SuperData) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
