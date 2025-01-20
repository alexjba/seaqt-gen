#pragma once
#ifndef MIQT_QT6_GEN_QTEXTOBJECT_H
#define MIQT_QT6_GEN_QTEXTOBJECT_H

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
class QGlyphRun;
class QMetaMethod;
class QMetaObject;
class QObject;
class QTextBlock;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QTextBlock__iterator)
typedef QTextBlock::iterator QTextBlock__iterator;
#else
class QTextBlock__iterator;
#endif
class QTextBlockFormat;
class QTextBlockGroup;
class QTextBlockUserData;
class QTextCharFormat;
class QTextCursor;
class QTextDocument;
class QTextFormat;
class QTextFragment;
class QTextFrame;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QTextFrame__iterator)
typedef QTextFrame::iterator QTextFrame__iterator;
#else
class QTextFrame__iterator;
#endif
class QTextFrameFormat;
class QTextFrameLayoutData;
class QTextLayout;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QTextLayout__FormatRange)
typedef QTextLayout::FormatRange QTextLayout__FormatRange;
#else
class QTextLayout__FormatRange;
#endif
class QTextList;
class QTextObject;
class QTimerEvent;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QGlyphRun QGlyphRun;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTextBlock QTextBlock;
typedef struct QTextBlock__iterator QTextBlock__iterator;
typedef struct QTextBlockFormat QTextBlockFormat;
typedef struct QTextBlockGroup QTextBlockGroup;
typedef struct QTextBlockUserData QTextBlockUserData;
typedef struct QTextCharFormat QTextCharFormat;
typedef struct QTextCursor QTextCursor;
typedef struct QTextDocument QTextDocument;
typedef struct QTextFormat QTextFormat;
typedef struct QTextFragment QTextFragment;
typedef struct QTextFrame QTextFrame;
typedef struct QTextFrame__iterator QTextFrame__iterator;
typedef struct QTextFrameFormat QTextFrameFormat;
typedef struct QTextFrameLayoutData QTextFrameLayoutData;
typedef struct QTextLayout QTextLayout;
typedef struct QTextLayout__FormatRange QTextLayout__FormatRange;
typedef struct QTextList QTextList;
typedef struct QTextObject QTextObject;
typedef struct QTimerEvent QTimerEvent;
#endif

void QTextObject_virtbase(QTextObject* src, QObject** outptr_QObject);
QMetaObject* QTextObject_MetaObject(const QTextObject* self);
void* QTextObject_Metacast(QTextObject* self, const char* param1);
struct miqt_string QTextObject_Tr(const char* s);
QTextFormat* QTextObject_Format(const QTextObject* self);
int QTextObject_FormatIndex(const QTextObject* self);
QTextDocument* QTextObject_Document(const QTextObject* self);
int QTextObject_ObjectIndex(const QTextObject* self);
struct miqt_string QTextObject_Tr2(const char* s, const char* c);
struct miqt_string QTextObject_Tr3(const char* s, const char* c, int n);

void QTextBlockGroup_virtbase(QTextBlockGroup* src, QTextObject** outptr_QTextObject);
QMetaObject* QTextBlockGroup_MetaObject(const QTextBlockGroup* self);
void* QTextBlockGroup_Metacast(QTextBlockGroup* self, const char* param1);
struct miqt_string QTextBlockGroup_Tr(const char* s);
void QTextBlockGroup_BlockInserted(QTextBlockGroup* self, QTextBlock* block);
void QTextBlockGroup_BlockRemoved(QTextBlockGroup* self, QTextBlock* block);
void QTextBlockGroup_BlockFormatChanged(QTextBlockGroup* self, QTextBlock* block);
struct miqt_string QTextBlockGroup_Tr2(const char* s, const char* c);
struct miqt_string QTextBlockGroup_Tr3(const char* s, const char* c, int n);

void QTextFrameLayoutData_OperatorAssign(QTextFrameLayoutData* self, QTextFrameLayoutData* param1);
void QTextFrameLayoutData_Delete(QTextFrameLayoutData* self);

QTextFrame* QTextFrame_new(QTextDocument* doc);
void QTextFrame_virtbase(QTextFrame* src, QTextObject** outptr_QTextObject);
QMetaObject* QTextFrame_MetaObject(const QTextFrame* self);
void* QTextFrame_Metacast(QTextFrame* self, const char* param1);
struct miqt_string QTextFrame_Tr(const char* s);
void QTextFrame_SetFrameFormat(QTextFrame* self, QTextFrameFormat* format);
QTextFrameFormat* QTextFrame_FrameFormat(const QTextFrame* self);
QTextCursor* QTextFrame_FirstCursorPosition(const QTextFrame* self);
QTextCursor* QTextFrame_LastCursorPosition(const QTextFrame* self);
int QTextFrame_FirstPosition(const QTextFrame* self);
int QTextFrame_LastPosition(const QTextFrame* self);
QTextFrameLayoutData* QTextFrame_LayoutData(const QTextFrame* self);
void QTextFrame_SetLayoutData(QTextFrame* self, QTextFrameLayoutData* data);
struct miqt_array /* of QTextFrame* */  QTextFrame_ChildFrames(const QTextFrame* self);
QTextFrame* QTextFrame_ParentFrame(const QTextFrame* self);
QTextFrame__iterator* QTextFrame_Begin(const QTextFrame* self);
QTextFrame__iterator* QTextFrame_End(const QTextFrame* self);
struct miqt_string QTextFrame_Tr2(const char* s, const char* c);
struct miqt_string QTextFrame_Tr3(const char* s, const char* c, int n);
bool QTextFrame_override_virtual_Event(void* self, intptr_t slot);
bool QTextFrame_virtualbase_Event(void* self, QEvent* event);
bool QTextFrame_override_virtual_EventFilter(void* self, intptr_t slot);
bool QTextFrame_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
bool QTextFrame_override_virtual_TimerEvent(void* self, intptr_t slot);
void QTextFrame_virtualbase_TimerEvent(void* self, QTimerEvent* event);
bool QTextFrame_override_virtual_ChildEvent(void* self, intptr_t slot);
void QTextFrame_virtualbase_ChildEvent(void* self, QChildEvent* event);
bool QTextFrame_override_virtual_CustomEvent(void* self, intptr_t slot);
void QTextFrame_virtualbase_CustomEvent(void* self, QEvent* event);
bool QTextFrame_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QTextFrame_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
bool QTextFrame_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QTextFrame_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
void QTextFrame_Delete(QTextFrame* self);

void QTextBlockUserData_OperatorAssign(QTextBlockUserData* self, QTextBlockUserData* param1);
void QTextBlockUserData_Delete(QTextBlockUserData* self);

QTextBlock* QTextBlock_new();
QTextBlock* QTextBlock_new2(QTextBlock* o);
void QTextBlock_OperatorAssign(QTextBlock* self, QTextBlock* o);
bool QTextBlock_IsValid(const QTextBlock* self);
bool QTextBlock_OperatorEqual(const QTextBlock* self, QTextBlock* o);
bool QTextBlock_OperatorNotEqual(const QTextBlock* self, QTextBlock* o);
bool QTextBlock_OperatorLesser(const QTextBlock* self, QTextBlock* o);
int QTextBlock_Position(const QTextBlock* self);
int QTextBlock_Length(const QTextBlock* self);
bool QTextBlock_Contains(const QTextBlock* self, int position);
QTextLayout* QTextBlock_Layout(const QTextBlock* self);
void QTextBlock_ClearLayout(QTextBlock* self);
QTextBlockFormat* QTextBlock_BlockFormat(const QTextBlock* self);
int QTextBlock_BlockFormatIndex(const QTextBlock* self);
QTextCharFormat* QTextBlock_CharFormat(const QTextBlock* self);
int QTextBlock_CharFormatIndex(const QTextBlock* self);
int QTextBlock_TextDirection(const QTextBlock* self);
struct miqt_string QTextBlock_Text(const QTextBlock* self);
struct miqt_array /* of QTextLayout__FormatRange* */  QTextBlock_TextFormats(const QTextBlock* self);
QTextDocument* QTextBlock_Document(const QTextBlock* self);
QTextList* QTextBlock_TextList(const QTextBlock* self);
QTextBlockUserData* QTextBlock_UserData(const QTextBlock* self);
void QTextBlock_SetUserData(QTextBlock* self, QTextBlockUserData* data);
int QTextBlock_UserState(const QTextBlock* self);
void QTextBlock_SetUserState(QTextBlock* self, int state);
int QTextBlock_Revision(const QTextBlock* self);
void QTextBlock_SetRevision(QTextBlock* self, int rev);
bool QTextBlock_IsVisible(const QTextBlock* self);
void QTextBlock_SetVisible(QTextBlock* self, bool visible);
int QTextBlock_BlockNumber(const QTextBlock* self);
int QTextBlock_FirstLineNumber(const QTextBlock* self);
void QTextBlock_SetLineCount(QTextBlock* self, int count);
int QTextBlock_LineCount(const QTextBlock* self);
QTextBlock__iterator* QTextBlock_Begin(const QTextBlock* self);
QTextBlock__iterator* QTextBlock_End(const QTextBlock* self);
QTextBlock* QTextBlock_Next(const QTextBlock* self);
QTextBlock* QTextBlock_Previous(const QTextBlock* self);
int QTextBlock_FragmentIndex(const QTextBlock* self);
void QTextBlock_Delete(QTextBlock* self);

QTextFragment* QTextFragment_new();
QTextFragment* QTextFragment_new2(QTextFragment* o);
void QTextFragment_OperatorAssign(QTextFragment* self, QTextFragment* o);
bool QTextFragment_IsValid(const QTextFragment* self);
bool QTextFragment_OperatorEqual(const QTextFragment* self, QTextFragment* o);
bool QTextFragment_OperatorNotEqual(const QTextFragment* self, QTextFragment* o);
bool QTextFragment_OperatorLesser(const QTextFragment* self, QTextFragment* o);
int QTextFragment_Position(const QTextFragment* self);
int QTextFragment_Length(const QTextFragment* self);
bool QTextFragment_Contains(const QTextFragment* self, int position);
QTextCharFormat* QTextFragment_CharFormat(const QTextFragment* self);
int QTextFragment_CharFormatIndex(const QTextFragment* self);
struct miqt_string QTextFragment_Text(const QTextFragment* self);
struct miqt_array /* of QGlyphRun* */  QTextFragment_GlyphRuns(const QTextFragment* self);
struct miqt_array /* of QGlyphRun* */  QTextFragment_GlyphRuns1(const QTextFragment* self, int from);
struct miqt_array /* of QGlyphRun* */  QTextFragment_GlyphRuns2(const QTextFragment* self, int from, int length);
void QTextFragment_Delete(QTextFragment* self);

QTextFrame__iterator* QTextFrame__iterator_new();
QTextFrame__iterator* QTextFrame__iterator_new2(QTextFrame__iterator* param1);
QTextFrame* QTextFrame__iterator_ParentFrame(const QTextFrame__iterator* self);
QTextFrame* QTextFrame__iterator_CurrentFrame(const QTextFrame__iterator* self);
QTextBlock* QTextFrame__iterator_CurrentBlock(const QTextFrame__iterator* self);
bool QTextFrame__iterator_AtEnd(const QTextFrame__iterator* self);
bool QTextFrame__iterator_OperatorEqual(const QTextFrame__iterator* self, QTextFrame__iterator* o);
bool QTextFrame__iterator_OperatorNotEqual(const QTextFrame__iterator* self, QTextFrame__iterator* o);
QTextFrame__iterator* QTextFrame__iterator_OperatorPlusPlus(QTextFrame__iterator* self);
QTextFrame__iterator* QTextFrame__iterator_OperatorPlusPlusWithInt(QTextFrame__iterator* self, int param1);
QTextFrame__iterator* QTextFrame__iterator_OperatorMinusMinus(QTextFrame__iterator* self);
QTextFrame__iterator* QTextFrame__iterator_OperatorMinusMinusWithInt(QTextFrame__iterator* self, int param1);
void QTextFrame__iterator_Delete(QTextFrame__iterator* self);

QTextBlock__iterator* QTextBlock__iterator_new();
QTextBlock__iterator* QTextBlock__iterator_new2(QTextBlock__iterator* param1);
QTextFragment* QTextBlock__iterator_Fragment(const QTextBlock__iterator* self);
bool QTextBlock__iterator_AtEnd(const QTextBlock__iterator* self);
bool QTextBlock__iterator_OperatorEqual(const QTextBlock__iterator* self, QTextBlock__iterator* o);
bool QTextBlock__iterator_OperatorNotEqual(const QTextBlock__iterator* self, QTextBlock__iterator* o);
QTextBlock__iterator* QTextBlock__iterator_OperatorPlusPlus(QTextBlock__iterator* self);
QTextBlock__iterator* QTextBlock__iterator_OperatorPlusPlusWithInt(QTextBlock__iterator* self, int param1);
QTextBlock__iterator* QTextBlock__iterator_OperatorMinusMinus(QTextBlock__iterator* self);
QTextBlock__iterator* QTextBlock__iterator_OperatorMinusMinusWithInt(QTextBlock__iterator* self, int param1);
void QTextBlock__iterator_Delete(QTextBlock__iterator* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
