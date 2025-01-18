#pragma once
#ifndef MIQT_QT6_GEN_QUNDOGROUP_H
#define MIQT_QT6_GEN_QUNDOGROUP_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAction;
class QChildEvent;
class QEvent;
class QMetaMethod;
class QMetaObject;
class QObject;
class QTimerEvent;
class QUndoGroup;
class QUndoStack;
#else
typedef struct QAction QAction;
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;
typedef struct QUndoGroup QUndoGroup;
typedef struct QUndoStack QUndoStack;
#endif

QUndoGroup* QUndoGroup_new();
QUndoGroup* QUndoGroup_new2(QObject* parent);
void QUndoGroup_virtbase(QUndoGroup* src, QObject** outptr_QObject);
QMetaObject* QUndoGroup_MetaObject(const QUndoGroup* self);
void* QUndoGroup_Metacast(QUndoGroup* self, const char* param1);
struct miqt_string QUndoGroup_Tr(const char* s);
void QUndoGroup_AddStack(QUndoGroup* self, QUndoStack* stack);
void QUndoGroup_RemoveStack(QUndoGroup* self, QUndoStack* stack);
struct miqt_array /* of QUndoStack* */  QUndoGroup_Stacks(const QUndoGroup* self);
QUndoStack* QUndoGroup_ActiveStack(const QUndoGroup* self);
QAction* QUndoGroup_CreateUndoAction(const QUndoGroup* self, QObject* parent);
QAction* QUndoGroup_CreateRedoAction(const QUndoGroup* self, QObject* parent);
bool QUndoGroup_CanUndo(const QUndoGroup* self);
bool QUndoGroup_CanRedo(const QUndoGroup* self);
struct miqt_string QUndoGroup_UndoText(const QUndoGroup* self);
struct miqt_string QUndoGroup_RedoText(const QUndoGroup* self);
bool QUndoGroup_IsClean(const QUndoGroup* self);
void QUndoGroup_Undo(QUndoGroup* self);
void QUndoGroup_Redo(QUndoGroup* self);
void QUndoGroup_SetActiveStack(QUndoGroup* self, QUndoStack* stack);
void QUndoGroup_ActiveStackChanged(QUndoGroup* self, QUndoStack* stack);
void QUndoGroup_connect_ActiveStackChanged(QUndoGroup* self, intptr_t slot);
void QUndoGroup_IndexChanged(QUndoGroup* self, int idx);
void QUndoGroup_connect_IndexChanged(QUndoGroup* self, intptr_t slot);
void QUndoGroup_CleanChanged(QUndoGroup* self, bool clean);
void QUndoGroup_connect_CleanChanged(QUndoGroup* self, intptr_t slot);
void QUndoGroup_CanUndoChanged(QUndoGroup* self, bool canUndo);
void QUndoGroup_connect_CanUndoChanged(QUndoGroup* self, intptr_t slot);
void QUndoGroup_CanRedoChanged(QUndoGroup* self, bool canRedo);
void QUndoGroup_connect_CanRedoChanged(QUndoGroup* self, intptr_t slot);
void QUndoGroup_UndoTextChanged(QUndoGroup* self, struct miqt_string undoText);
void QUndoGroup_connect_UndoTextChanged(QUndoGroup* self, intptr_t slot);
void QUndoGroup_RedoTextChanged(QUndoGroup* self, struct miqt_string redoText);
void QUndoGroup_connect_RedoTextChanged(QUndoGroup* self, intptr_t slot);
struct miqt_string QUndoGroup_Tr2(const char* s, const char* c);
struct miqt_string QUndoGroup_Tr3(const char* s, const char* c, int n);
QAction* QUndoGroup_CreateUndoAction2(const QUndoGroup* self, QObject* parent, struct miqt_string prefix);
QAction* QUndoGroup_CreateRedoAction2(const QUndoGroup* self, QObject* parent, struct miqt_string prefix);
bool QUndoGroup_override_virtual_Event(void* self, intptr_t slot);
bool QUndoGroup_virtualbase_Event(void* self, QEvent* event);
bool QUndoGroup_override_virtual_EventFilter(void* self, intptr_t slot);
bool QUndoGroup_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
bool QUndoGroup_override_virtual_TimerEvent(void* self, intptr_t slot);
void QUndoGroup_virtualbase_TimerEvent(void* self, QTimerEvent* event);
bool QUndoGroup_override_virtual_ChildEvent(void* self, intptr_t slot);
void QUndoGroup_virtualbase_ChildEvent(void* self, QChildEvent* event);
bool QUndoGroup_override_virtual_CustomEvent(void* self, intptr_t slot);
void QUndoGroup_virtualbase_CustomEvent(void* self, QEvent* event);
bool QUndoGroup_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QUndoGroup_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
bool QUndoGroup_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QUndoGroup_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
void QUndoGroup_Delete(QUndoGroup* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
