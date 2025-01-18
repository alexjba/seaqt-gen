#pragma once
#ifndef MIQT_QT6_GEN_QCOREEVENT_H
#define MIQT_QT6_GEN_QCOREEVENT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QChildEvent;
class QDynamicPropertyChangeEvent;
class QEvent;
class QObject;
class QTimerEvent;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QDynamicPropertyChangeEvent QDynamicPropertyChangeEvent;
typedef struct QEvent QEvent;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;
#endif

QEvent* QEvent_new(int typeVal);
int QEvent_Type(const QEvent* self);
bool QEvent_Spontaneous(const QEvent* self);
void QEvent_SetAccepted(QEvent* self, bool accepted);
bool QEvent_IsAccepted(const QEvent* self);
void QEvent_Accept(QEvent* self);
void QEvent_Ignore(QEvent* self);
bool QEvent_IsInputEvent(const QEvent* self);
bool QEvent_IsPointerEvent(const QEvent* self);
bool QEvent_IsSinglePointEvent(const QEvent* self);
int QEvent_RegisterEventType();
QEvent* QEvent_Clone(const QEvent* self);
int QEvent_RegisterEventType1(int hint);
bool QEvent_override_virtual_SetAccepted(void* self, intptr_t slot);
void QEvent_virtualbase_SetAccepted(void* self, bool accepted);
bool QEvent_override_virtual_Clone(void* self, intptr_t slot);
QEvent* QEvent_virtualbase_Clone(const void* self);
void QEvent_Delete(QEvent* self);

QTimerEvent* QTimerEvent_new(int timerId);
void QTimerEvent_virtbase(QTimerEvent* src, QEvent** outptr_QEvent);
QTimerEvent* QTimerEvent_Clone(const QTimerEvent* self);
int QTimerEvent_TimerId(const QTimerEvent* self);
bool QTimerEvent_override_virtual_Clone(void* self, intptr_t slot);
QTimerEvent* QTimerEvent_virtualbase_Clone(const void* self);
bool QTimerEvent_override_virtual_SetAccepted(void* self, intptr_t slot);
void QTimerEvent_virtualbase_SetAccepted(void* self, bool accepted);
void QTimerEvent_Delete(QTimerEvent* self);

QChildEvent* QChildEvent_new(int typeVal, QObject* child);
void QChildEvent_virtbase(QChildEvent* src, QEvent** outptr_QEvent);
QChildEvent* QChildEvent_Clone(const QChildEvent* self);
QObject* QChildEvent_Child(const QChildEvent* self);
bool QChildEvent_Added(const QChildEvent* self);
bool QChildEvent_Polished(const QChildEvent* self);
bool QChildEvent_Removed(const QChildEvent* self);
bool QChildEvent_override_virtual_Clone(void* self, intptr_t slot);
QChildEvent* QChildEvent_virtualbase_Clone(const void* self);
bool QChildEvent_override_virtual_SetAccepted(void* self, intptr_t slot);
void QChildEvent_virtualbase_SetAccepted(void* self, bool accepted);
void QChildEvent_Delete(QChildEvent* self);

QDynamicPropertyChangeEvent* QDynamicPropertyChangeEvent_new(struct miqt_string name);
void QDynamicPropertyChangeEvent_virtbase(QDynamicPropertyChangeEvent* src, QEvent** outptr_QEvent);
QDynamicPropertyChangeEvent* QDynamicPropertyChangeEvent_Clone(const QDynamicPropertyChangeEvent* self);
struct miqt_string QDynamicPropertyChangeEvent_PropertyName(const QDynamicPropertyChangeEvent* self);
bool QDynamicPropertyChangeEvent_override_virtual_Clone(void* self, intptr_t slot);
QDynamicPropertyChangeEvent* QDynamicPropertyChangeEvent_virtualbase_Clone(const void* self);
bool QDynamicPropertyChangeEvent_override_virtual_SetAccepted(void* self, intptr_t slot);
void QDynamicPropertyChangeEvent_virtualbase_SetAccepted(void* self, bool accepted);
void QDynamicPropertyChangeEvent_Delete(QDynamicPropertyChangeEvent* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
