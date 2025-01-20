#pragma once
#ifndef MIQT_QT_RESTRICTED_EXTRAS_QSCINTILLA_GEN_QSCILEXERCSS_H
#define MIQT_QT_RESTRICTED_EXTRAS_QSCINTILLA_GEN_QSCILEXERCSS_H

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
class QsciLexerCSS;
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
typedef struct QsciLexerCSS QsciLexerCSS;
typedef struct QsciScintilla QsciScintilla;
#endif

QsciLexerCSS* QsciLexerCSS_new();
QsciLexerCSS* QsciLexerCSS_new2(QObject* parent);
void QsciLexerCSS_virtbase(QsciLexerCSS* src, QsciLexer** outptr_QsciLexer);
QMetaObject* QsciLexerCSS_MetaObject(const QsciLexerCSS* self);
void* QsciLexerCSS_Metacast(QsciLexerCSS* self, const char* param1);
struct miqt_string QsciLexerCSS_Tr(const char* s);
struct miqt_string QsciLexerCSS_TrUtf8(const char* s);
const char* QsciLexerCSS_Language(const QsciLexerCSS* self);
const char* QsciLexerCSS_Lexer(const QsciLexerCSS* self);
const char* QsciLexerCSS_BlockEnd(const QsciLexerCSS* self);
const char* QsciLexerCSS_BlockStart(const QsciLexerCSS* self);
const char* QsciLexerCSS_WordCharacters(const QsciLexerCSS* self);
QColor* QsciLexerCSS_DefaultColor(const QsciLexerCSS* self, int style);
QFont* QsciLexerCSS_DefaultFont(const QsciLexerCSS* self, int style);
const char* QsciLexerCSS_Keywords(const QsciLexerCSS* self, int set);
struct miqt_string QsciLexerCSS_Description(const QsciLexerCSS* self, int style);
void QsciLexerCSS_RefreshProperties(QsciLexerCSS* self);
bool QsciLexerCSS_FoldComments(const QsciLexerCSS* self);
bool QsciLexerCSS_FoldCompact(const QsciLexerCSS* self);
void QsciLexerCSS_SetHSSLanguage(QsciLexerCSS* self, bool enabled);
bool QsciLexerCSS_HSSLanguage(const QsciLexerCSS* self);
void QsciLexerCSS_SetLessLanguage(QsciLexerCSS* self, bool enabled);
bool QsciLexerCSS_LessLanguage(const QsciLexerCSS* self);
void QsciLexerCSS_SetSCSSLanguage(QsciLexerCSS* self, bool enabled);
bool QsciLexerCSS_SCSSLanguage(const QsciLexerCSS* self);
void QsciLexerCSS_SetFoldComments(QsciLexerCSS* self, bool fold);
void QsciLexerCSS_SetFoldCompact(QsciLexerCSS* self, bool fold);
struct miqt_string QsciLexerCSS_Tr2(const char* s, const char* c);
struct miqt_string QsciLexerCSS_Tr3(const char* s, const char* c, int n);
struct miqt_string QsciLexerCSS_TrUtf82(const char* s, const char* c);
struct miqt_string QsciLexerCSS_TrUtf83(const char* s, const char* c, int n);
const char* QsciLexerCSS_BlockEnd1(const QsciLexerCSS* self, int* style);
const char* QsciLexerCSS_BlockStart1(const QsciLexerCSS* self, int* style);
bool QsciLexerCSS_override_virtual_SetFoldComments(void* self, intptr_t slot);
void QsciLexerCSS_virtualbase_SetFoldComments(void* self, bool fold);
bool QsciLexerCSS_override_virtual_SetFoldCompact(void* self, intptr_t slot);
void QsciLexerCSS_virtualbase_SetFoldCompact(void* self, bool fold);
bool QsciLexerCSS_override_virtual_Language(void* self, intptr_t slot);
const char* QsciLexerCSS_virtualbase_Language(const void* self);
bool QsciLexerCSS_override_virtual_Lexer(void* self, intptr_t slot);
const char* QsciLexerCSS_virtualbase_Lexer(const void* self);
bool QsciLexerCSS_override_virtual_LexerId(void* self, intptr_t slot);
int QsciLexerCSS_virtualbase_LexerId(const void* self);
bool QsciLexerCSS_override_virtual_AutoCompletionFillups(void* self, intptr_t slot);
const char* QsciLexerCSS_virtualbase_AutoCompletionFillups(const void* self);
bool QsciLexerCSS_override_virtual_AutoCompletionWordSeparators(void* self, intptr_t slot);
struct miqt_array /* of struct miqt_string */  QsciLexerCSS_virtualbase_AutoCompletionWordSeparators(const void* self);
bool QsciLexerCSS_override_virtual_BlockEnd(void* self, intptr_t slot);
const char* QsciLexerCSS_virtualbase_BlockEnd(const void* self, int* style);
bool QsciLexerCSS_override_virtual_BlockLookback(void* self, intptr_t slot);
int QsciLexerCSS_virtualbase_BlockLookback(const void* self);
bool QsciLexerCSS_override_virtual_BlockStart(void* self, intptr_t slot);
const char* QsciLexerCSS_virtualbase_BlockStart(const void* self, int* style);
bool QsciLexerCSS_override_virtual_BlockStartKeyword(void* self, intptr_t slot);
const char* QsciLexerCSS_virtualbase_BlockStartKeyword(const void* self, int* style);
bool QsciLexerCSS_override_virtual_BraceStyle(void* self, intptr_t slot);
int QsciLexerCSS_virtualbase_BraceStyle(const void* self);
bool QsciLexerCSS_override_virtual_CaseSensitive(void* self, intptr_t slot);
bool QsciLexerCSS_virtualbase_CaseSensitive(const void* self);
bool QsciLexerCSS_override_virtual_Color(void* self, intptr_t slot);
QColor* QsciLexerCSS_virtualbase_Color(const void* self, int style);
bool QsciLexerCSS_override_virtual_EolFill(void* self, intptr_t slot);
bool QsciLexerCSS_virtualbase_EolFill(const void* self, int style);
bool QsciLexerCSS_override_virtual_Font(void* self, intptr_t slot);
QFont* QsciLexerCSS_virtualbase_Font(const void* self, int style);
bool QsciLexerCSS_override_virtual_IndentationGuideView(void* self, intptr_t slot);
int QsciLexerCSS_virtualbase_IndentationGuideView(const void* self);
bool QsciLexerCSS_override_virtual_Keywords(void* self, intptr_t slot);
const char* QsciLexerCSS_virtualbase_Keywords(const void* self, int set);
bool QsciLexerCSS_override_virtual_DefaultStyle(void* self, intptr_t slot);
int QsciLexerCSS_virtualbase_DefaultStyle(const void* self);
bool QsciLexerCSS_override_virtual_Description(void* self, intptr_t slot);
struct miqt_string QsciLexerCSS_virtualbase_Description(const void* self, int style);
bool QsciLexerCSS_override_virtual_Paper(void* self, intptr_t slot);
QColor* QsciLexerCSS_virtualbase_Paper(const void* self, int style);
bool QsciLexerCSS_override_virtual_DefaultColorWithStyle(void* self, intptr_t slot);
QColor* QsciLexerCSS_virtualbase_DefaultColorWithStyle(const void* self, int style);
bool QsciLexerCSS_override_virtual_DefaultEolFill(void* self, intptr_t slot);
bool QsciLexerCSS_virtualbase_DefaultEolFill(const void* self, int style);
bool QsciLexerCSS_override_virtual_DefaultFontWithStyle(void* self, intptr_t slot);
QFont* QsciLexerCSS_virtualbase_DefaultFontWithStyle(const void* self, int style);
bool QsciLexerCSS_override_virtual_DefaultPaperWithStyle(void* self, intptr_t slot);
QColor* QsciLexerCSS_virtualbase_DefaultPaperWithStyle(const void* self, int style);
bool QsciLexerCSS_override_virtual_SetEditor(void* self, intptr_t slot);
void QsciLexerCSS_virtualbase_SetEditor(void* self, QsciScintilla* editor);
bool QsciLexerCSS_override_virtual_RefreshProperties(void* self, intptr_t slot);
void QsciLexerCSS_virtualbase_RefreshProperties(void* self);
bool QsciLexerCSS_override_virtual_StyleBitsNeeded(void* self, intptr_t slot);
int QsciLexerCSS_virtualbase_StyleBitsNeeded(const void* self);
bool QsciLexerCSS_override_virtual_WordCharacters(void* self, intptr_t slot);
const char* QsciLexerCSS_virtualbase_WordCharacters(const void* self);
bool QsciLexerCSS_override_virtual_SetAutoIndentStyle(void* self, intptr_t slot);
void QsciLexerCSS_virtualbase_SetAutoIndentStyle(void* self, int autoindentstyle);
bool QsciLexerCSS_override_virtual_SetColor(void* self, intptr_t slot);
void QsciLexerCSS_virtualbase_SetColor(void* self, QColor* c, int style);
bool QsciLexerCSS_override_virtual_SetEolFill(void* self, intptr_t slot);
void QsciLexerCSS_virtualbase_SetEolFill(void* self, bool eoffill, int style);
bool QsciLexerCSS_override_virtual_SetFont(void* self, intptr_t slot);
void QsciLexerCSS_virtualbase_SetFont(void* self, QFont* f, int style);
bool QsciLexerCSS_override_virtual_SetPaper(void* self, intptr_t slot);
void QsciLexerCSS_virtualbase_SetPaper(void* self, QColor* c, int style);
bool QsciLexerCSS_override_virtual_ReadProperties(void* self, intptr_t slot);
bool QsciLexerCSS_virtualbase_ReadProperties(void* self, QSettings* qs, struct miqt_string prefix);
bool QsciLexerCSS_override_virtual_WriteProperties(void* self, intptr_t slot);
bool QsciLexerCSS_virtualbase_WriteProperties(const void* self, QSettings* qs, struct miqt_string prefix);
bool QsciLexerCSS_override_virtual_Event(void* self, intptr_t slot);
bool QsciLexerCSS_virtualbase_Event(void* self, QEvent* event);
bool QsciLexerCSS_override_virtual_EventFilter(void* self, intptr_t slot);
bool QsciLexerCSS_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
bool QsciLexerCSS_override_virtual_TimerEvent(void* self, intptr_t slot);
void QsciLexerCSS_virtualbase_TimerEvent(void* self, QTimerEvent* event);
bool QsciLexerCSS_override_virtual_ChildEvent(void* self, intptr_t slot);
void QsciLexerCSS_virtualbase_ChildEvent(void* self, QChildEvent* event);
bool QsciLexerCSS_override_virtual_CustomEvent(void* self, intptr_t slot);
void QsciLexerCSS_virtualbase_CustomEvent(void* self, QEvent* event);
bool QsciLexerCSS_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QsciLexerCSS_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
bool QsciLexerCSS_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QsciLexerCSS_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
void QsciLexerCSS_Delete(QsciLexerCSS* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
