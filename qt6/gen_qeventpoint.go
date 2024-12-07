package qt6

/*

#include "gen_qeventpoint.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QEventPoint__State byte

const (
	QEventPoint__Unknown    QEventPoint__State = 0
	QEventPoint__Stationary QEventPoint__State = 4
	QEventPoint__Pressed    QEventPoint__State = 1
	QEventPoint__Updated    QEventPoint__State = 2
	QEventPoint__Released   QEventPoint__State = 8
)

type QEventPoint struct {
	h          *C.QEventPoint
	isSubclass bool
}

func (this *QEventPoint) cPointer() *C.QEventPoint {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QEventPoint) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQEventPoint constructs the type using only CGO pointers.
func newQEventPoint(h *C.QEventPoint) *QEventPoint {
	if h == nil {
		return nil
	}
	return &QEventPoint{h: h}
}

// UnsafeNewQEventPoint constructs the type using only unsafe pointers.
func UnsafeNewQEventPoint(h unsafe.Pointer) *QEventPoint {
	if h == nil {
		return nil
	}

	return &QEventPoint{h: (*C.QEventPoint)(h)}
}

// NewQEventPoint constructs a new QEventPoint object.
func NewQEventPoint() *QEventPoint {
	var outptr_QEventPoint *C.QEventPoint = nil

	C.QEventPoint_new(&outptr_QEventPoint)
	ret := newQEventPoint(outptr_QEventPoint)
	ret.isSubclass = true
	return ret
}

// NewQEventPoint2 constructs a new QEventPoint object.
func NewQEventPoint2(pointId int, state QEventPoint__State, scenePosition *QPointF, globalPosition *QPointF) *QEventPoint {
	var outptr_QEventPoint *C.QEventPoint = nil

	C.QEventPoint_new2((C.int)(pointId), (C.uint8_t)(state), scenePosition.cPointer(), globalPosition.cPointer(), &outptr_QEventPoint)
	ret := newQEventPoint(outptr_QEventPoint)
	ret.isSubclass = true
	return ret
}

// NewQEventPoint3 constructs a new QEventPoint object.
func NewQEventPoint3(other *QEventPoint) *QEventPoint {
	var outptr_QEventPoint *C.QEventPoint = nil

	C.QEventPoint_new3(other.cPointer(), &outptr_QEventPoint)
	ret := newQEventPoint(outptr_QEventPoint)
	ret.isSubclass = true
	return ret
}

// NewQEventPoint4 constructs a new QEventPoint object.
func NewQEventPoint4(id int) *QEventPoint {
	var outptr_QEventPoint *C.QEventPoint = nil

	C.QEventPoint_new4((C.int)(id), &outptr_QEventPoint)
	ret := newQEventPoint(outptr_QEventPoint)
	ret.isSubclass = true
	return ret
}

// NewQEventPoint5 constructs a new QEventPoint object.
func NewQEventPoint5(id int, device *QPointingDevice) *QEventPoint {
	var outptr_QEventPoint *C.QEventPoint = nil

	C.QEventPoint_new5((C.int)(id), device.cPointer(), &outptr_QEventPoint)
	ret := newQEventPoint(outptr_QEventPoint)
	ret.isSubclass = true
	return ret
}

func (this *QEventPoint) OperatorAssign(other *QEventPoint) {
	C.QEventPoint_OperatorAssign(this.h, other.cPointer())
}

func (this *QEventPoint) OperatorEqual(other *QEventPoint) bool {
	return (bool)(C.QEventPoint_OperatorEqual(this.h, other.cPointer()))
}

func (this *QEventPoint) OperatorNotEqual(other *QEventPoint) bool {
	return (bool)(C.QEventPoint_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QEventPoint) Swap(other *QEventPoint) {
	C.QEventPoint_Swap(this.h, other.cPointer())
}

func (this *QEventPoint) Position() *QPointF {
	_goptr := newQPointF(C.QEventPoint_Position(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) PressPosition() *QPointF {
	_goptr := newQPointF(C.QEventPoint_PressPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) GrabPosition() *QPointF {
	_goptr := newQPointF(C.QEventPoint_GrabPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) LastPosition() *QPointF {
	_goptr := newQPointF(C.QEventPoint_LastPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) ScenePosition() *QPointF {
	_goptr := newQPointF(C.QEventPoint_ScenePosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) ScenePressPosition() *QPointF {
	_goptr := newQPointF(C.QEventPoint_ScenePressPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) SceneGrabPosition() *QPointF {
	_goptr := newQPointF(C.QEventPoint_SceneGrabPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) SceneLastPosition() *QPointF {
	_goptr := newQPointF(C.QEventPoint_SceneLastPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) GlobalPosition() *QPointF {
	_goptr := newQPointF(C.QEventPoint_GlobalPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) GlobalPressPosition() *QPointF {
	_goptr := newQPointF(C.QEventPoint_GlobalPressPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) GlobalGrabPosition() *QPointF {
	_goptr := newQPointF(C.QEventPoint_GlobalGrabPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) GlobalLastPosition() *QPointF {
	_goptr := newQPointF(C.QEventPoint_GlobalLastPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) NormalizedPosition() *QPointF {
	_goptr := newQPointF(C.QEventPoint_NormalizedPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) Pos() *QPointF {
	_goptr := newQPointF(C.QEventPoint_Pos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) StartPos() *QPointF {
	_goptr := newQPointF(C.QEventPoint_StartPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) ScenePos() *QPointF {
	_goptr := newQPointF(C.QEventPoint_ScenePos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) StartScenePos() *QPointF {
	_goptr := newQPointF(C.QEventPoint_StartScenePos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) ScreenPos() *QPointF {
	_goptr := newQPointF(C.QEventPoint_ScreenPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) StartScreenPos() *QPointF {
	_goptr := newQPointF(C.QEventPoint_StartScreenPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) StartNormalizedPos() *QPointF {
	_goptr := newQPointF(C.QEventPoint_StartNormalizedPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) NormalizedPos() *QPointF {
	_goptr := newQPointF(C.QEventPoint_NormalizedPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) LastPos() *QPointF {
	_goptr := newQPointF(C.QEventPoint_LastPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) LastScenePos() *QPointF {
	_goptr := newQPointF(C.QEventPoint_LastScenePos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) LastScreenPos() *QPointF {
	_goptr := newQPointF(C.QEventPoint_LastScreenPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) LastNormalizedPos() *QPointF {
	_goptr := newQPointF(C.QEventPoint_LastNormalizedPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) Velocity() *QVector2D {
	_goptr := newQVector2D(C.QEventPoint_Velocity(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) State() QEventPoint__State {
	return (QEventPoint__State)(C.QEventPoint_State(this.h))
}

func (this *QEventPoint) Device() *QPointingDevice {
	return newQPointingDevice(C.QEventPoint_Device(this.h), nil, nil)
}

func (this *QEventPoint) Id() int {
	return (int)(C.QEventPoint_Id(this.h))
}

func (this *QEventPoint) UniqueId() *QPointingDeviceUniqueId {
	_goptr := newQPointingDeviceUniqueId(C.QEventPoint_UniqueId(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) Timestamp() uint64 {
	return (uint64)(C.QEventPoint_Timestamp(this.h))
}

func (this *QEventPoint) LastTimestamp() uint64 {
	return (uint64)(C.QEventPoint_LastTimestamp(this.h))
}

func (this *QEventPoint) PressTimestamp() uint64 {
	return (uint64)(C.QEventPoint_PressTimestamp(this.h))
}

func (this *QEventPoint) TimeHeld() float64 {
	return (float64)(C.QEventPoint_TimeHeld(this.h))
}

func (this *QEventPoint) Pressure() float64 {
	return (float64)(C.QEventPoint_Pressure(this.h))
}

func (this *QEventPoint) Rotation() float64 {
	return (float64)(C.QEventPoint_Rotation(this.h))
}

func (this *QEventPoint) EllipseDiameters() *QSizeF {
	_goptr := newQSizeF(C.QEventPoint_EllipseDiameters(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) IsAccepted() bool {
	return (bool)(C.QEventPoint_IsAccepted(this.h))
}

func (this *QEventPoint) SetAccepted() {
	C.QEventPoint_SetAccepted(this.h)
}

func (this *QEventPoint) SetAccepted1(accepted bool) {
	C.QEventPoint_SetAccepted1(this.h, (C.bool)(accepted))
}

// Delete this object from C++ memory.
func (this *QEventPoint) Delete() {
	C.QEventPoint_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QEventPoint) GoGC() {
	runtime.SetFinalizer(this, func(this *QEventPoint) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
