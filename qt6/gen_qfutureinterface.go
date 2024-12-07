package qt6

/*

#include "gen_qfutureinterface.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QFutureInterfaceBase__State int

const (
	QFutureInterfaceBase__NoState    QFutureInterfaceBase__State = 0
	QFutureInterfaceBase__Running    QFutureInterfaceBase__State = 1
	QFutureInterfaceBase__Started    QFutureInterfaceBase__State = 2
	QFutureInterfaceBase__Finished   QFutureInterfaceBase__State = 4
	QFutureInterfaceBase__Canceled   QFutureInterfaceBase__State = 8
	QFutureInterfaceBase__Suspending QFutureInterfaceBase__State = 16
	QFutureInterfaceBase__Suspended  QFutureInterfaceBase__State = 32
	QFutureInterfaceBase__Throttled  QFutureInterfaceBase__State = 64
	QFutureInterfaceBase__Pending    QFutureInterfaceBase__State = 128
)

type QFutureInterfaceBase__CancelMode int

const (
	QFutureInterfaceBase__CancelOnly      QFutureInterfaceBase__CancelMode = 0
	QFutureInterfaceBase__CancelAndFinish QFutureInterfaceBase__CancelMode = 1
)

type QFutureInterfaceBase struct {
	h          *C.QFutureInterfaceBase
	isSubclass bool
}

func (this *QFutureInterfaceBase) cPointer() *C.QFutureInterfaceBase {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QFutureInterfaceBase) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQFutureInterfaceBase constructs the type using only CGO pointers.
func newQFutureInterfaceBase(h *C.QFutureInterfaceBase) *QFutureInterfaceBase {
	if h == nil {
		return nil
	}

	return &QFutureInterfaceBase{h: h}
}

// UnsafeNewQFutureInterfaceBase constructs the type using only unsafe pointers.
func UnsafeNewQFutureInterfaceBase(h unsafe.Pointer) *QFutureInterfaceBase {
	return newQFutureInterfaceBase((*C.QFutureInterfaceBase)(h))
}

// NewQFutureInterfaceBase constructs a new QFutureInterfaceBase object.
func NewQFutureInterfaceBase() *QFutureInterfaceBase {

	ret := newQFutureInterfaceBase(C.QFutureInterfaceBase_new())
	ret.isSubclass = true
	return ret
}

// NewQFutureInterfaceBase2 constructs a new QFutureInterfaceBase object.
func NewQFutureInterfaceBase2(other *QFutureInterfaceBase) *QFutureInterfaceBase {

	ret := newQFutureInterfaceBase(C.QFutureInterfaceBase_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFutureInterfaceBase3 constructs a new QFutureInterfaceBase object.
func NewQFutureInterfaceBase3(initialState QFutureInterfaceBase__State) *QFutureInterfaceBase {

	ret := newQFutureInterfaceBase(C.QFutureInterfaceBase_new3((C.int)(initialState)))
	ret.isSubclass = true
	return ret
}

func (this *QFutureInterfaceBase) OperatorAssign(other *QFutureInterfaceBase) {
	C.QFutureInterfaceBase_OperatorAssign(this.h, other.cPointer())
}

func (this *QFutureInterfaceBase) ReportStarted() {
	C.QFutureInterfaceBase_ReportStarted(this.h)
}

func (this *QFutureInterfaceBase) ReportFinished() {
	C.QFutureInterfaceBase_ReportFinished(this.h)
}

func (this *QFutureInterfaceBase) ReportCanceled() {
	C.QFutureInterfaceBase_ReportCanceled(this.h)
}

func (this *QFutureInterfaceBase) ReportResultsReady(beginIndex int, endIndex int) {
	C.QFutureInterfaceBase_ReportResultsReady(this.h, (C.int)(beginIndex), (C.int)(endIndex))
}

func (this *QFutureInterfaceBase) SetRunnable(runnable *QRunnable) {
	C.QFutureInterfaceBase_SetRunnable(this.h, runnable.cPointer())
}

func (this *QFutureInterfaceBase) SetThreadPool(pool *QThreadPool) {
	C.QFutureInterfaceBase_SetThreadPool(this.h, pool.cPointer())
}

func (this *QFutureInterfaceBase) ThreadPool() *QThreadPool {
	return newQThreadPool(C.QFutureInterfaceBase_ThreadPool(this.h))
}

func (this *QFutureInterfaceBase) SetFilterMode(enable bool) {
	C.QFutureInterfaceBase_SetFilterMode(this.h, (C.bool)(enable))
}

func (this *QFutureInterfaceBase) SetProgressRange(minimum int, maximum int) {
	C.QFutureInterfaceBase_SetProgressRange(this.h, (C.int)(minimum), (C.int)(maximum))
}

func (this *QFutureInterfaceBase) ProgressMinimum() int {
	return (int)(C.QFutureInterfaceBase_ProgressMinimum(this.h))
}

func (this *QFutureInterfaceBase) ProgressMaximum() int {
	return (int)(C.QFutureInterfaceBase_ProgressMaximum(this.h))
}

func (this *QFutureInterfaceBase) IsProgressUpdateNeeded() bool {
	return (bool)(C.QFutureInterfaceBase_IsProgressUpdateNeeded(this.h))
}

func (this *QFutureInterfaceBase) SetProgressValue(progressValue int) {
	C.QFutureInterfaceBase_SetProgressValue(this.h, (C.int)(progressValue))
}

func (this *QFutureInterfaceBase) ProgressValue() int {
	return (int)(C.QFutureInterfaceBase_ProgressValue(this.h))
}

func (this *QFutureInterfaceBase) SetProgressValueAndText(progressValue int, progressText string) {
	progressText_ms := C.struct_miqt_string{}
	progressText_ms.data = C.CString(progressText)
	progressText_ms.len = C.size_t(len(progressText))
	defer C.free(unsafe.Pointer(progressText_ms.data))
	C.QFutureInterfaceBase_SetProgressValueAndText(this.h, (C.int)(progressValue), progressText_ms)
}

func (this *QFutureInterfaceBase) ProgressText() string {
	var _ms C.struct_miqt_string = C.QFutureInterfaceBase_ProgressText(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFutureInterfaceBase) SetExpectedResultCount(resultCount int) {
	C.QFutureInterfaceBase_SetExpectedResultCount(this.h, (C.int)(resultCount))
}

func (this *QFutureInterfaceBase) ExpectedResultCount() int {
	return (int)(C.QFutureInterfaceBase_ExpectedResultCount(this.h))
}

func (this *QFutureInterfaceBase) ResultCount() int {
	return (int)(C.QFutureInterfaceBase_ResultCount(this.h))
}

func (this *QFutureInterfaceBase) QueryState(state QFutureInterfaceBase__State) bool {
	return (bool)(C.QFutureInterfaceBase_QueryState(this.h, (C.int)(state)))
}

func (this *QFutureInterfaceBase) IsRunning() bool {
	return (bool)(C.QFutureInterfaceBase_IsRunning(this.h))
}

func (this *QFutureInterfaceBase) IsStarted() bool {
	return (bool)(C.QFutureInterfaceBase_IsStarted(this.h))
}

func (this *QFutureInterfaceBase) IsCanceled() bool {
	return (bool)(C.QFutureInterfaceBase_IsCanceled(this.h))
}

func (this *QFutureInterfaceBase) IsFinished() bool {
	return (bool)(C.QFutureInterfaceBase_IsFinished(this.h))
}

func (this *QFutureInterfaceBase) IsPaused() bool {
	return (bool)(C.QFutureInterfaceBase_IsPaused(this.h))
}

func (this *QFutureInterfaceBase) SetPaused(paused bool) {
	C.QFutureInterfaceBase_SetPaused(this.h, (C.bool)(paused))
}

func (this *QFutureInterfaceBase) TogglePaused() {
	C.QFutureInterfaceBase_TogglePaused(this.h)
}

func (this *QFutureInterfaceBase) IsSuspending() bool {
	return (bool)(C.QFutureInterfaceBase_IsSuspending(this.h))
}

func (this *QFutureInterfaceBase) IsSuspended() bool {
	return (bool)(C.QFutureInterfaceBase_IsSuspended(this.h))
}

func (this *QFutureInterfaceBase) IsThrottled() bool {
	return (bool)(C.QFutureInterfaceBase_IsThrottled(this.h))
}

func (this *QFutureInterfaceBase) IsResultReadyAt(index int) bool {
	return (bool)(C.QFutureInterfaceBase_IsResultReadyAt(this.h, (C.int)(index)))
}

func (this *QFutureInterfaceBase) IsValid() bool {
	return (bool)(C.QFutureInterfaceBase_IsValid(this.h))
}

func (this *QFutureInterfaceBase) LoadState() int {
	return (int)(C.QFutureInterfaceBase_LoadState(this.h))
}

func (this *QFutureInterfaceBase) Cancel() {
	C.QFutureInterfaceBase_Cancel(this.h)
}

func (this *QFutureInterfaceBase) CancelAndFinish() {
	C.QFutureInterfaceBase_CancelAndFinish(this.h)
}

func (this *QFutureInterfaceBase) SetSuspended(suspend bool) {
	C.QFutureInterfaceBase_SetSuspended(this.h, (C.bool)(suspend))
}

func (this *QFutureInterfaceBase) ToggleSuspended() {
	C.QFutureInterfaceBase_ToggleSuspended(this.h)
}

func (this *QFutureInterfaceBase) ReportSuspended() {
	C.QFutureInterfaceBase_ReportSuspended(this.h)
}

func (this *QFutureInterfaceBase) SetThrottled(enable bool) {
	C.QFutureInterfaceBase_SetThrottled(this.h, (C.bool)(enable))
}

func (this *QFutureInterfaceBase) WaitForFinished() {
	C.QFutureInterfaceBase_WaitForFinished(this.h)
}

func (this *QFutureInterfaceBase) WaitForNextResult() bool {
	return (bool)(C.QFutureInterfaceBase_WaitForNextResult(this.h))
}

func (this *QFutureInterfaceBase) WaitForResult(resultIndex int) {
	C.QFutureInterfaceBase_WaitForResult(this.h, (C.int)(resultIndex))
}

func (this *QFutureInterfaceBase) WaitForResume() {
	C.QFutureInterfaceBase_WaitForResume(this.h)
}

func (this *QFutureInterfaceBase) SuspendIfRequested() {
	C.QFutureInterfaceBase_SuspendIfRequested(this.h)
}

func (this *QFutureInterfaceBase) Mutex() *QMutex {
	return newQMutex(C.QFutureInterfaceBase_Mutex(this.h))
}

func (this *QFutureInterfaceBase) HasException() bool {
	return (bool)(C.QFutureInterfaceBase_HasException(this.h))
}

func (this *QFutureInterfaceBase) OperatorEqual(other *QFutureInterfaceBase) bool {
	return (bool)(C.QFutureInterfaceBase_OperatorEqual(this.h, other.cPointer()))
}

func (this *QFutureInterfaceBase) OperatorNotEqual(other *QFutureInterfaceBase) bool {
	return (bool)(C.QFutureInterfaceBase_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QFutureInterfaceBase) Swap(other *QFutureInterfaceBase) {
	C.QFutureInterfaceBase_Swap(this.h, other.cPointer())
}

func (this *QFutureInterfaceBase) IsChainCanceled() bool {
	return (bool)(C.QFutureInterfaceBase_IsChainCanceled(this.h))
}

// Delete this object from C++ memory.
func (this *QFutureInterfaceBase) Delete() {
	C.QFutureInterfaceBase_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QFutureInterfaceBase) GoGC() {
	runtime.SetFinalizer(this, func(this *QFutureInterfaceBase) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
