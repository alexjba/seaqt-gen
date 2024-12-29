#pragma once
#ifndef MIQT_QT_RESTRICTED_EXTRAS_QSCINTILLA6_GEN_QSCILEXERLUA_H
#define MIQT_QT_RESTRICTED_EXTRAS_QSCINTILLA6_GEN_QSCILEXERLUA_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QColor;
class QFont;
class QMetaObject;
class QObject;
class QSettings;
class QsciLexer;
class QsciLexerLua;
class QsciScintilla;
#else
typedef struct QColor QColor;
typedef struct QFont QFont;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QSettings QSettings;
typedef struct QsciLexer QsciLexer;
typedef struct QsciLexerLua QsciLexerLua;
typedef struct QsciScintilla QsciScintilla;
#endif

QsciLexerLua* QsciLexerLua_new();
QsciLexerLua* QsciLexerLua_new2(QObject* parent);
void QsciLexerLua_virtbase(QsciLexerLua* src, QsciLexer** outptr_QsciLexer);
QMetaObject* QsciLexerLua_MetaObject(const QsciLexerLua* self);
void* QsciLexerLua_Metacast(QsciLexerLua* self, const char* param1);
struct miqt_string QsciLexerLua_Tr(const char* s);
const char* QsciLexerLua_Language(const QsciLexerLua* self);
const char* QsciLexerLua_Lexer(const QsciLexerLua* self);
struct miqt_array /* of struct miqt_string */  QsciLexerLua_AutoCompletionWordSeparators(const QsciLexerLua* self);
const char* QsciLexerLua_BlockStart(const QsciLexerLua* self);
int QsciLexerLua_BraceStyle(const QsciLexerLua* self);
QColor* QsciLexerLua_DefaultColor(const QsciLexerLua* self, int style);
bool QsciLexerLua_DefaultEolFill(const QsciLexerLua* self, int style);
QFont* QsciLexerLua_DefaultFont(const QsciLexerLua* self, int style);
QColor* QsciLexerLua_DefaultPaper(const QsciLexerLua* self, int style);
const char* QsciLexerLua_Keywords(const QsciLexerLua* self, int set);
struct miqt_string QsciLexerLua_Description(const QsciLexerLua* self, int style);
void QsciLexerLua_RefreshProperties(QsciLexerLua* self);
bool QsciLexerLua_FoldCompact(const QsciLexerLua* self);
void QsciLexerLua_SetFoldCompact(QsciLexerLua* self, bool fold);
struct miqt_string QsciLexerLua_Tr2(const char* s, const char* c);
struct miqt_string QsciLexerLua_Tr3(const char* s, const char* c, int n);
const char* QsciLexerLua_BlockStart1(const QsciLexerLua* self, int* style);
void QsciLexerLua_override_virtual_SetFoldCompact(void* self, intptr_t slot);
void QsciLexerLua_virtualbase_SetFoldCompact(void* self, bool fold);
void QsciLexerLua_override_virtual_Language(void* self, intptr_t slot);
const char* QsciLexerLua_virtualbase_Language(const void* self);
void QsciLexerLua_override_virtual_Lexer(void* self, intptr_t slot);
const char* QsciLexerLua_virtualbase_Lexer(const void* self);
void QsciLexerLua_override_virtual_LexerId(void* self, intptr_t slot);
int QsciLexerLua_virtualbase_LexerId(const void* self);
void QsciLexerLua_override_virtual_AutoCompletionFillups(void* self, intptr_t slot);
const char* QsciLexerLua_virtualbase_AutoCompletionFillups(const void* self);
void QsciLexerLua_override_virtual_AutoCompletionWordSeparators(void* self, intptr_t slot);
struct miqt_array /* of struct miqt_string */  QsciLexerLua_virtualbase_AutoCompletionWordSeparators(const void* self);
void QsciLexerLua_override_virtual_BlockEnd(void* self, intptr_t slot);
const char* QsciLexerLua_virtualbase_BlockEnd(const void* self, int* style);
void QsciLexerLua_override_virtual_BlockLookback(void* self, intptr_t slot);
int QsciLexerLua_virtualbase_BlockLookback(const void* self);
void QsciLexerLua_override_virtual_BlockStart(void* self, intptr_t slot);
const char* QsciLexerLua_virtualbase_BlockStart(const void* self, int* style);
void QsciLexerLua_override_virtual_BlockStartKeyword(void* self, intptr_t slot);
const char* QsciLexerLua_virtualbase_BlockStartKeyword(const void* self, int* style);
void QsciLexerLua_override_virtual_BraceStyle(void* self, intptr_t slot);
int QsciLexerLua_virtualbase_BraceStyle(const void* self);
void QsciLexerLua_override_virtual_CaseSensitive(void* self, intptr_t slot);
bool QsciLexerLua_virtualbase_CaseSensitive(const void* self);
void QsciLexerLua_override_virtual_Color(void* self, intptr_t slot);
QColor* QsciLexerLua_virtualbase_Color(const void* self, int style);
void QsciLexerLua_override_virtual_EolFill(void* self, intptr_t slot);
bool QsciLexerLua_virtualbase_EolFill(const void* self, int style);
void QsciLexerLua_override_virtual_Font(void* self, intptr_t slot);
QFont* QsciLexerLua_virtualbase_Font(const void* self, int style);
void QsciLexerLua_override_virtual_IndentationGuideView(void* self, intptr_t slot);
int QsciLexerLua_virtualbase_IndentationGuideView(const void* self);
void QsciLexerLua_override_virtual_Keywords(void* self, intptr_t slot);
const char* QsciLexerLua_virtualbase_Keywords(const void* self, int set);
void QsciLexerLua_override_virtual_DefaultStyle(void* self, intptr_t slot);
int QsciLexerLua_virtualbase_DefaultStyle(const void* self);
void QsciLexerLua_override_virtual_Description(void* self, intptr_t slot);
struct miqt_string QsciLexerLua_virtualbase_Description(const void* self, int style);
void QsciLexerLua_override_virtual_Paper(void* self, intptr_t slot);
QColor* QsciLexerLua_virtualbase_Paper(const void* self, int style);
void QsciLexerLua_override_virtual_DefaultColorWithStyle(void* self, intptr_t slot);
QColor* QsciLexerLua_virtualbase_DefaultColorWithStyle(const void* self, int style);
void QsciLexerLua_override_virtual_DefaultEolFill(void* self, intptr_t slot);
bool QsciLexerLua_virtualbase_DefaultEolFill(const void* self, int style);
void QsciLexerLua_override_virtual_DefaultFontWithStyle(void* self, intptr_t slot);
QFont* QsciLexerLua_virtualbase_DefaultFontWithStyle(const void* self, int style);
void QsciLexerLua_override_virtual_DefaultPaperWithStyle(void* self, intptr_t slot);
QColor* QsciLexerLua_virtualbase_DefaultPaperWithStyle(const void* self, int style);
void QsciLexerLua_override_virtual_SetEditor(void* self, intptr_t slot);
void QsciLexerLua_virtualbase_SetEditor(void* self, QsciScintilla* editor);
void QsciLexerLua_override_virtual_RefreshProperties(void* self, intptr_t slot);
void QsciLexerLua_virtualbase_RefreshProperties(void* self);
void QsciLexerLua_override_virtual_StyleBitsNeeded(void* self, intptr_t slot);
int QsciLexerLua_virtualbase_StyleBitsNeeded(const void* self);
void QsciLexerLua_override_virtual_WordCharacters(void* self, intptr_t slot);
const char* QsciLexerLua_virtualbase_WordCharacters(const void* self);
void QsciLexerLua_override_virtual_SetAutoIndentStyle(void* self, intptr_t slot);
void QsciLexerLua_virtualbase_SetAutoIndentStyle(void* self, int autoindentstyle);
void QsciLexerLua_override_virtual_SetColor(void* self, intptr_t slot);
void QsciLexerLua_virtualbase_SetColor(void* self, QColor* c, int style);
void QsciLexerLua_override_virtual_SetEolFill(void* self, intptr_t slot);
void QsciLexerLua_virtualbase_SetEolFill(void* self, bool eoffill, int style);
void QsciLexerLua_override_virtual_SetFont(void* self, intptr_t slot);
void QsciLexerLua_virtualbase_SetFont(void* self, QFont* f, int style);
void QsciLexerLua_override_virtual_SetPaper(void* self, intptr_t slot);
void QsciLexerLua_virtualbase_SetPaper(void* self, QColor* c, int style);
void QsciLexerLua_override_virtual_ReadProperties(void* self, intptr_t slot);
bool QsciLexerLua_virtualbase_ReadProperties(void* self, QSettings* qs, struct miqt_string prefix);
void QsciLexerLua_override_virtual_WriteProperties(void* self, intptr_t slot);
bool QsciLexerLua_virtualbase_WriteProperties(const void* self, QSettings* qs, struct miqt_string prefix);
void QsciLexerLua_Delete(QsciLexerLua* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
