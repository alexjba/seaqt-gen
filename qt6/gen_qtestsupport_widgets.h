#pragma once
#ifndef MIQT_QT6_GEN_QTESTSUPPORT_WIDGETS_H
#define MIQT_QT6_GEN_QTESTSUPPORT_WIDGETS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QPoint;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QTest__QTouchEventSequence)
typedef QTest::QTouchEventSequence QTest__QTouchEventSequence;
#else
class QTest__QTouchEventSequence;
#endif
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QTest__QTouchEventWidgetSequence)
typedef QTest::QTouchEventWidgetSequence QTest__QTouchEventWidgetSequence;
#else
class QTest__QTouchEventWidgetSequence;
#endif
class QWidget;
#else
typedef struct QPoint QPoint;
typedef struct QTest__QTouchEventSequence QTest__QTouchEventSequence;
typedef struct QTest__QTouchEventWidgetSequence QTest__QTouchEventWidgetSequence;
typedef struct QWidget QWidget;
#endif

QTest__QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_new(QTest__QTouchEventWidgetSequence* param1);
void QTest__QTouchEventWidgetSequence_virtbase(QTest__QTouchEventWidgetSequence* src, QTest__QTouchEventSequence** outptr_QTest__QTouchEventSequence);
QTest__QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Press(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt);
QTest__QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Move(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt);
QTest__QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Release(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt);
QTest__QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Stationary(QTest__QTouchEventWidgetSequence* self, int touchId);
bool QTest__QTouchEventWidgetSequence_Commit(QTest__QTouchEventWidgetSequence* self, bool processEvents);
QTest__QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Press3(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt, QWidget* widget);
QTest__QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Move3(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt, QWidget* widget);
QTest__QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Release3(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt, QWidget* widget);
bool QTest__QTouchEventWidgetSequence_override_virtual_Stationary(void* self, intptr_t slot);
QTest__QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_virtualbase_Stationary(void* self, int touchId);
bool QTest__QTouchEventWidgetSequence_override_virtual_Commit(void* self, intptr_t slot);
bool QTest__QTouchEventWidgetSequence_virtualbase_Commit(void* self, bool processEvents);
void QTest__QTouchEventWidgetSequence_Delete(QTest__QTouchEventWidgetSequence* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
