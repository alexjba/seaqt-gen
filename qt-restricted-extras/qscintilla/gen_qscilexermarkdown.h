#pragma once
#ifndef MIQT_QT_RESTRICTED_EXTRAS_QSCINTILLA_GEN_QSCILEXERMARKDOWN_H
#define MIQT_QT_RESTRICTED_EXTRAS_QSCINTILLA_GEN_QSCILEXERMARKDOWN_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QChildEvent;
class QColor;
class QEvent;
class QFont;
class QMetaMethod;
class QMetaObject;
class QObject;
class QSettings;
class QTimerEvent;
class QsciLexer;
class QsciLexerMarkdown;
class QsciScintilla;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QColor QColor;
typedef struct QEvent QEvent;
typedef struct QFont QFont;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QSettings QSettings;
typedef struct QTimerEvent QTimerEvent;
typedef struct QsciLexer QsciLexer;
typedef struct QsciLexerMarkdown QsciLexerMarkdown;
typedef struct QsciScintilla QsciScintilla;
#endif

QsciLexerMarkdown* QsciLexerMarkdown_new();
QsciLexerMarkdown* QsciLexerMarkdown_new2(QObject* parent);
void QsciLexerMarkdown_virtbase(QsciLexerMarkdown* src, QsciLexer** outptr_QsciLexer);
QMetaObject* QsciLexerMarkdown_MetaObject(const QsciLexerMarkdown* self);
void* QsciLexerMarkdown_Metacast(QsciLexerMarkdown* self, const char* param1);
struct miqt_string QsciLexerMarkdown_Tr(const char* s);
struct miqt_string QsciLexerMarkdown_TrUtf8(const char* s);
const char* QsciLexerMarkdown_Language(const QsciLexerMarkdown* self);
const char* QsciLexerMarkdown_Lexer(const QsciLexerMarkdown* self);
QColor* QsciLexerMarkdown_DefaultColor(const QsciLexerMarkdown* self, int style);
QFont* QsciLexerMarkdown_DefaultFont(const QsciLexerMarkdown* self, int style);
QColor* QsciLexerMarkdown_DefaultPaper(const QsciLexerMarkdown* self, int style);
struct miqt_string QsciLexerMarkdown_Description(const QsciLexerMarkdown* self, int style);
struct miqt_string QsciLexerMarkdown_Tr2(const char* s, const char* c);
struct miqt_string QsciLexerMarkdown_Tr3(const char* s, const char* c, int n);
struct miqt_string QsciLexerMarkdown_TrUtf82(const char* s, const char* c);
struct miqt_string QsciLexerMarkdown_TrUtf83(const char* s, const char* c, int n);
bool QsciLexerMarkdown_override_virtual_Language(void* self, intptr_t slot);
const char* QsciLexerMarkdown_virtualbase_Language(const void* self);
bool QsciLexerMarkdown_override_virtual_Lexer(void* self, intptr_t slot);
const char* QsciLexerMarkdown_virtualbase_Lexer(const void* self);
bool QsciLexerMarkdown_override_virtual_LexerId(void* self, intptr_t slot);
int QsciLexerMarkdown_virtualbase_LexerId(const void* self);
bool QsciLexerMarkdown_override_virtual_AutoCompletionFillups(void* self, intptr_t slot);
const char* QsciLexerMarkdown_virtualbase_AutoCompletionFillups(const void* self);
bool QsciLexerMarkdown_override_virtual_AutoCompletionWordSeparators(void* self, intptr_t slot);
struct miqt_array /* of struct miqt_string */  QsciLexerMarkdown_virtualbase_AutoCompletionWordSeparators(const void* self);
bool QsciLexerMarkdown_override_virtual_BlockEnd(void* self, intptr_t slot);
const char* QsciLexerMarkdown_virtualbase_BlockEnd(const void* self, int* style);
bool QsciLexerMarkdown_override_virtual_BlockLookback(void* self, intptr_t slot);
int QsciLexerMarkdown_virtualbase_BlockLookback(const void* self);
bool QsciLexerMarkdown_override_virtual_BlockStart(void* self, intptr_t slot);
const char* QsciLexerMarkdown_virtualbase_BlockStart(const void* self, int* style);
bool QsciLexerMarkdown_override_virtual_BlockStartKeyword(void* self, intptr_t slot);
const char* QsciLexerMarkdown_virtualbase_BlockStartKeyword(const void* self, int* style);
bool QsciLexerMarkdown_override_virtual_BraceStyle(void* self, intptr_t slot);
int QsciLexerMarkdown_virtualbase_BraceStyle(const void* self);
bool QsciLexerMarkdown_override_virtual_CaseSensitive(void* self, intptr_t slot);
bool QsciLexerMarkdown_virtualbase_CaseSensitive(const void* self);
bool QsciLexerMarkdown_override_virtual_Color(void* self, intptr_t slot);
QColor* QsciLexerMarkdown_virtualbase_Color(const void* self, int style);
bool QsciLexerMarkdown_override_virtual_EolFill(void* self, intptr_t slot);
bool QsciLexerMarkdown_virtualbase_EolFill(const void* self, int style);
bool QsciLexerMarkdown_override_virtual_Font(void* self, intptr_t slot);
QFont* QsciLexerMarkdown_virtualbase_Font(const void* self, int style);
bool QsciLexerMarkdown_override_virtual_IndentationGuideView(void* self, intptr_t slot);
int QsciLexerMarkdown_virtualbase_IndentationGuideView(const void* self);
bool QsciLexerMarkdown_override_virtual_Keywords(void* self, intptr_t slot);
const char* QsciLexerMarkdown_virtualbase_Keywords(const void* self, int set);
bool QsciLexerMarkdown_override_virtual_DefaultStyle(void* self, intptr_t slot);
int QsciLexerMarkdown_virtualbase_DefaultStyle(const void* self);
bool QsciLexerMarkdown_override_virtual_Description(void* self, intptr_t slot);
struct miqt_string QsciLexerMarkdown_virtualbase_Description(const void* self, int style);
bool QsciLexerMarkdown_override_virtual_Paper(void* self, intptr_t slot);
QColor* QsciLexerMarkdown_virtualbase_Paper(const void* self, int style);
bool QsciLexerMarkdown_override_virtual_DefaultColorWithStyle(void* self, intptr_t slot);
QColor* QsciLexerMarkdown_virtualbase_DefaultColorWithStyle(const void* self, int style);
bool QsciLexerMarkdown_override_virtual_DefaultEolFill(void* self, intptr_t slot);
bool QsciLexerMarkdown_virtualbase_DefaultEolFill(const void* self, int style);
bool QsciLexerMarkdown_override_virtual_DefaultFontWithStyle(void* self, intptr_t slot);
QFont* QsciLexerMarkdown_virtualbase_DefaultFontWithStyle(const void* self, int style);
bool QsciLexerMarkdown_override_virtual_DefaultPaperWithStyle(void* self, intptr_t slot);
QColor* QsciLexerMarkdown_virtualbase_DefaultPaperWithStyle(const void* self, int style);
bool QsciLexerMarkdown_override_virtual_SetEditor(void* self, intptr_t slot);
void QsciLexerMarkdown_virtualbase_SetEditor(void* self, QsciScintilla* editor);
bool QsciLexerMarkdown_override_virtual_RefreshProperties(void* self, intptr_t slot);
void QsciLexerMarkdown_virtualbase_RefreshProperties(void* self);
bool QsciLexerMarkdown_override_virtual_StyleBitsNeeded(void* self, intptr_t slot);
int QsciLexerMarkdown_virtualbase_StyleBitsNeeded(const void* self);
bool QsciLexerMarkdown_override_virtual_WordCharacters(void* self, intptr_t slot);
const char* QsciLexerMarkdown_virtualbase_WordCharacters(const void* self);
bool QsciLexerMarkdown_override_virtual_SetAutoIndentStyle(void* self, intptr_t slot);
void QsciLexerMarkdown_virtualbase_SetAutoIndentStyle(void* self, int autoindentstyle);
bool QsciLexerMarkdown_override_virtual_SetColor(void* self, intptr_t slot);
void QsciLexerMarkdown_virtualbase_SetColor(void* self, QColor* c, int style);
bool QsciLexerMarkdown_override_virtual_SetEolFill(void* self, intptr_t slot);
void QsciLexerMarkdown_virtualbase_SetEolFill(void* self, bool eoffill, int style);
bool QsciLexerMarkdown_override_virtual_SetFont(void* self, intptr_t slot);
void QsciLexerMarkdown_virtualbase_SetFont(void* self, QFont* f, int style);
bool QsciLexerMarkdown_override_virtual_SetPaper(void* self, intptr_t slot);
void QsciLexerMarkdown_virtualbase_SetPaper(void* self, QColor* c, int style);
bool QsciLexerMarkdown_override_virtual_ReadProperties(void* self, intptr_t slot);
bool QsciLexerMarkdown_virtualbase_ReadProperties(void* self, QSettings* qs, struct miqt_string prefix);
bool QsciLexerMarkdown_override_virtual_WriteProperties(void* self, intptr_t slot);
bool QsciLexerMarkdown_virtualbase_WriteProperties(const void* self, QSettings* qs, struct miqt_string prefix);
bool QsciLexerMarkdown_override_virtual_Event(void* self, intptr_t slot);
bool QsciLexerMarkdown_virtualbase_Event(void* self, QEvent* event);
bool QsciLexerMarkdown_override_virtual_EventFilter(void* self, intptr_t slot);
bool QsciLexerMarkdown_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
bool QsciLexerMarkdown_override_virtual_TimerEvent(void* self, intptr_t slot);
void QsciLexerMarkdown_virtualbase_TimerEvent(void* self, QTimerEvent* event);
bool QsciLexerMarkdown_override_virtual_ChildEvent(void* self, intptr_t slot);
void QsciLexerMarkdown_virtualbase_ChildEvent(void* self, QChildEvent* event);
bool QsciLexerMarkdown_override_virtual_CustomEvent(void* self, intptr_t slot);
void QsciLexerMarkdown_virtualbase_CustomEvent(void* self, QEvent* event);
bool QsciLexerMarkdown_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QsciLexerMarkdown_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
bool QsciLexerMarkdown_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QsciLexerMarkdown_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
void QsciLexerMarkdown_Delete(QsciLexerMarkdown* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
