#pragma once
#ifndef MIQT_QT_GEN_QSOCKETNOTIFIER_H
#define MIQT_QT_GEN_QSOCKETNOTIFIER_H

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
class QEvent;
class QMetaMethod;
class QMetaObject;
class QObject;
class QSocketDescriptor;
class QSocketNotifier;
class QTimerEvent;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QSocketDescriptor QSocketDescriptor;
typedef struct QSocketNotifier QSocketNotifier;
typedef struct QTimerEvent QTimerEvent;
#endif

QSocketNotifier* QSocketNotifier_new(intptr_t socket, int param2);
QSocketNotifier* QSocketNotifier_new2(intptr_t socket, int param2, QObject* parent);
void QSocketNotifier_virtbase(QSocketNotifier* src, QObject** outptr_QObject);
QMetaObject* QSocketNotifier_MetaObject(const QSocketNotifier* self);
void* QSocketNotifier_Metacast(QSocketNotifier* self, const char* param1);
struct miqt_string QSocketNotifier_Tr(const char* s);
struct miqt_string QSocketNotifier_TrUtf8(const char* s);
intptr_t QSocketNotifier_Socket(const QSocketNotifier* self);
int QSocketNotifier_Type(const QSocketNotifier* self);
bool QSocketNotifier_IsEnabled(const QSocketNotifier* self);
void QSocketNotifier_SetEnabled(QSocketNotifier* self, bool enabled);
bool QSocketNotifier_Event(QSocketNotifier* self, QEvent* param1);
struct miqt_string QSocketNotifier_Tr2(const char* s, const char* c);
struct miqt_string QSocketNotifier_Tr3(const char* s, const char* c, int n);
struct miqt_string QSocketNotifier_TrUtf82(const char* s, const char* c);
struct miqt_string QSocketNotifier_TrUtf83(const char* s, const char* c, int n);
bool QSocketNotifier_override_virtual_Event(void* self, intptr_t slot);
bool QSocketNotifier_virtualbase_Event(void* self, QEvent* param1);
bool QSocketNotifier_override_virtual_EventFilter(void* self, intptr_t slot);
bool QSocketNotifier_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
bool QSocketNotifier_override_virtual_TimerEvent(void* self, intptr_t slot);
void QSocketNotifier_virtualbase_TimerEvent(void* self, QTimerEvent* event);
bool QSocketNotifier_override_virtual_ChildEvent(void* self, intptr_t slot);
void QSocketNotifier_virtualbase_ChildEvent(void* self, QChildEvent* event);
bool QSocketNotifier_override_virtual_CustomEvent(void* self, intptr_t slot);
void QSocketNotifier_virtualbase_CustomEvent(void* self, QEvent* event);
bool QSocketNotifier_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QSocketNotifier_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
bool QSocketNotifier_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QSocketNotifier_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
void QSocketNotifier_Delete(QSocketNotifier* self);

QSocketDescriptor* QSocketDescriptor_new();
QSocketDescriptor* QSocketDescriptor_new2(QSocketDescriptor* param1);
QSocketDescriptor* QSocketDescriptor_new3(int descriptor);
int QSocketDescriptor_ToInt(const QSocketDescriptor* self);
bool QSocketDescriptor_IsValid(const QSocketDescriptor* self);
void QSocketDescriptor_Delete(QSocketDescriptor* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
