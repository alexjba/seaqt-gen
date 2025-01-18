#pragma once
#ifndef MIQT_QT_GEN_QLOCALE_H
#define MIQT_QT_GEN_QLOCALE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QCalendar;
class QChar;
class QDate;
class QDateTime;
class QLocale;
class QTime;
#else
typedef struct QCalendar QCalendar;
typedef struct QChar QChar;
typedef struct QDate QDate;
typedef struct QDateTime QDateTime;
typedef struct QLocale QLocale;
typedef struct QTime QTime;
#endif

QLocale* QLocale_new();
QLocale* QLocale_new2(struct miqt_string name);
QLocale* QLocale_new3(int language);
QLocale* QLocale_new4(int language, int script, int country);
QLocale* QLocale_new5(QLocale* other);
QLocale* QLocale_new6(int language, int country);
void QLocale_OperatorAssign(QLocale* self, QLocale* other);
void QLocale_Swap(QLocale* self, QLocale* other);
int QLocale_Language(const QLocale* self);
int QLocale_Script(const QLocale* self);
int QLocale_Country(const QLocale* self);
struct miqt_string QLocale_Name(const QLocale* self);
struct miqt_string QLocale_Bcp47Name(const QLocale* self);
struct miqt_string QLocale_NativeLanguageName(const QLocale* self);
struct miqt_string QLocale_NativeCountryName(const QLocale* self);
int16_t QLocale_ToShort(const QLocale* self, struct miqt_string s);
uint16_t QLocale_ToUShort(const QLocale* self, struct miqt_string s);
int QLocale_ToInt(const QLocale* self, struct miqt_string s);
unsigned int QLocale_ToUInt(const QLocale* self, struct miqt_string s);
long QLocale_ToLong(const QLocale* self, struct miqt_string s);
unsigned long QLocale_ToULong(const QLocale* self, struct miqt_string s);
long long QLocale_ToLongLong(const QLocale* self, struct miqt_string s);
unsigned long long QLocale_ToULongLong(const QLocale* self, struct miqt_string s);
float QLocale_ToFloat(const QLocale* self, struct miqt_string s);
double QLocale_ToDouble(const QLocale* self, struct miqt_string s);
struct miqt_string QLocale_ToString(const QLocale* self, long long i);
struct miqt_string QLocale_ToStringWithQulonglong(const QLocale* self, unsigned long long i);
struct miqt_string QLocale_ToStringWithLong(const QLocale* self, long i);
struct miqt_string QLocale_ToStringWithUlong(const QLocale* self, unsigned long i);
struct miqt_string QLocale_ToStringWithShort(const QLocale* self, int16_t i);
struct miqt_string QLocale_ToStringWithUshort(const QLocale* self, uint16_t i);
struct miqt_string QLocale_ToStringWithInt(const QLocale* self, int i);
struct miqt_string QLocale_ToStringWithUint(const QLocale* self, unsigned int i);
struct miqt_string QLocale_ToStringWithDouble(const QLocale* self, double i);
struct miqt_string QLocale_ToStringWithFloat(const QLocale* self, float i);
struct miqt_string QLocale_ToString2(const QLocale* self, QDate* date, struct miqt_string formatStr);
struct miqt_string QLocale_ToString3(const QLocale* self, QTime* time, struct miqt_string formatStr);
struct miqt_string QLocale_ToString4(const QLocale* self, QDateTime* dateTime, struct miqt_string format);
struct miqt_string QLocale_ToStringWithDate(const QLocale* self, QDate* date);
struct miqt_string QLocale_ToStringWithTime(const QLocale* self, QTime* time);
struct miqt_string QLocale_ToStringWithDateTime(const QLocale* self, QDateTime* dateTime);
struct miqt_string QLocale_ToString9(const QLocale* self, QDate* date, int format, QCalendar* cal);
struct miqt_string QLocale_ToString10(const QLocale* self, QDateTime* dateTime, int format, QCalendar* cal);
struct miqt_string QLocale_DateFormat(const QLocale* self);
struct miqt_string QLocale_TimeFormat(const QLocale* self);
struct miqt_string QLocale_DateTimeFormat(const QLocale* self);
QDate* QLocale_ToDate(const QLocale* self, struct miqt_string stringVal);
QTime* QLocale_ToTime(const QLocale* self, struct miqt_string stringVal);
QDateTime* QLocale_ToDateTime(const QLocale* self, struct miqt_string stringVal);
QDate* QLocale_ToDate2(const QLocale* self, struct miqt_string stringVal, struct miqt_string format);
QTime* QLocale_ToTime2(const QLocale* self, struct miqt_string stringVal, struct miqt_string format);
QDateTime* QLocale_ToDateTime2(const QLocale* self, struct miqt_string stringVal, struct miqt_string format);
QDate* QLocale_ToDate3(const QLocale* self, struct miqt_string stringVal, int format, QCalendar* cal);
QDateTime* QLocale_ToDateTime3(const QLocale* self, struct miqt_string stringVal, int format, QCalendar* cal);
QDate* QLocale_ToDate4(const QLocale* self, struct miqt_string stringVal, struct miqt_string format, QCalendar* cal);
QDateTime* QLocale_ToDateTime4(const QLocale* self, struct miqt_string stringVal, struct miqt_string format, QCalendar* cal);
QTime* QLocale_ToTime3(const QLocale* self, struct miqt_string stringVal, int format, QCalendar* cal);
QTime* QLocale_ToTime4(const QLocale* self, struct miqt_string stringVal, struct miqt_string format, QCalendar* cal);
QChar* QLocale_DecimalPoint(const QLocale* self);
QChar* QLocale_GroupSeparator(const QLocale* self);
QChar* QLocale_Percent(const QLocale* self);
QChar* QLocale_ZeroDigit(const QLocale* self);
QChar* QLocale_NegativeSign(const QLocale* self);
QChar* QLocale_PositiveSign(const QLocale* self);
QChar* QLocale_Exponential(const QLocale* self);
struct miqt_string QLocale_MonthName(const QLocale* self, int param1);
struct miqt_string QLocale_StandaloneMonthName(const QLocale* self, int param1);
struct miqt_string QLocale_DayName(const QLocale* self, int param1);
struct miqt_string QLocale_StandaloneDayName(const QLocale* self, int param1);
int QLocale_FirstDayOfWeek(const QLocale* self);
struct miqt_array /* of int */  QLocale_Weekdays(const QLocale* self);
struct miqt_string QLocale_AmText(const QLocale* self);
struct miqt_string QLocale_PmText(const QLocale* self);
int QLocale_MeasurementSystem(const QLocale* self);
QLocale* QLocale_Collation(const QLocale* self);
int QLocale_TextDirection(const QLocale* self);
struct miqt_string QLocale_ToUpper(const QLocale* self, struct miqt_string str);
struct miqt_string QLocale_ToLower(const QLocale* self, struct miqt_string str);
struct miqt_string QLocale_CurrencySymbol(const QLocale* self);
struct miqt_string QLocale_ToCurrencyString(const QLocale* self, long long param1);
struct miqt_string QLocale_ToCurrencyStringWithQulonglong(const QLocale* self, unsigned long long param1);
struct miqt_string QLocale_ToCurrencyStringWithShort(const QLocale* self, int16_t param1);
struct miqt_string QLocale_ToCurrencyStringWithUshort(const QLocale* self, uint16_t param1);
struct miqt_string QLocale_ToCurrencyStringWithInt(const QLocale* self, int param1);
struct miqt_string QLocale_ToCurrencyStringWithUint(const QLocale* self, unsigned int param1);
struct miqt_string QLocale_ToCurrencyStringWithDouble(const QLocale* self, double param1);
struct miqt_string QLocale_ToCurrencyString2(const QLocale* self, double param1, struct miqt_string symbol, int precision);
struct miqt_string QLocale_ToCurrencyStringWithFloat(const QLocale* self, float i);
struct miqt_string QLocale_ToCurrencyString3(const QLocale* self, float i, struct miqt_string symbol, int precision);
struct miqt_string QLocale_FormattedDataSize(QLocale* self, long long bytes);
struct miqt_string QLocale_FormattedDataSizeWithBytes(const QLocale* self, long long bytes);
struct miqt_array /* of struct miqt_string */  QLocale_UiLanguages(const QLocale* self);
bool QLocale_OperatorEqual(const QLocale* self, QLocale* other);
bool QLocale_OperatorNotEqual(const QLocale* self, QLocale* other);
struct miqt_string QLocale_LanguageToString(int language);
struct miqt_string QLocale_CountryToString(int country);
struct miqt_string QLocale_ScriptToString(int script);
void QLocale_SetDefault(QLocale* locale);
QLocale* QLocale_C();
QLocale* QLocale_System();
struct miqt_array /* of QLocale* */  QLocale_MatchingLocales(int language, int script, int country);
struct miqt_array /* of int */  QLocale_CountriesForLanguage(int lang);
void QLocale_SetNumberOptions(QLocale* self, int options);
int QLocale_NumberOptions(const QLocale* self);
struct miqt_string QLocale_QuoteString(const QLocale* self, struct miqt_string str);
struct miqt_string QLocale_CreateSeparatedList(const QLocale* self, struct miqt_array /* of struct miqt_string */  strl);
int16_t QLocale_ToShort2(const QLocale* self, struct miqt_string s, bool* ok);
uint16_t QLocale_ToUShort2(const QLocale* self, struct miqt_string s, bool* ok);
int QLocale_ToInt2(const QLocale* self, struct miqt_string s, bool* ok);
unsigned int QLocale_ToUInt2(const QLocale* self, struct miqt_string s, bool* ok);
long QLocale_ToLong2(const QLocale* self, struct miqt_string s, bool* ok);
unsigned long QLocale_ToULong2(const QLocale* self, struct miqt_string s, bool* ok);
long long QLocale_ToLongLong2(const QLocale* self, struct miqt_string s, bool* ok);
unsigned long long QLocale_ToULongLong2(const QLocale* self, struct miqt_string s, bool* ok);
float QLocale_ToFloat2(const QLocale* self, struct miqt_string s, bool* ok);
double QLocale_ToDouble2(const QLocale* self, struct miqt_string s, bool* ok);
struct miqt_string QLocale_ToString22(const QLocale* self, double i, char f);
struct miqt_string QLocale_ToString32(const QLocale* self, double i, char f, int prec);
struct miqt_string QLocale_ToString23(const QLocale* self, float i, char f);
struct miqt_string QLocale_ToString33(const QLocale* self, float i, char f, int prec);
struct miqt_string QLocale_ToString24(const QLocale* self, QDate* date, int format);
struct miqt_string QLocale_ToString25(const QLocale* self, QTime* time, int format);
struct miqt_string QLocale_ToString26(const QLocale* self, QDateTime* dateTime, int format);
struct miqt_string QLocale_DateFormat1(const QLocale* self, int format);
struct miqt_string QLocale_TimeFormat1(const QLocale* self, int format);
struct miqt_string QLocale_DateTimeFormat1(const QLocale* self, int format);
QDate* QLocale_ToDate22(const QLocale* self, struct miqt_string stringVal, int param2);
QTime* QLocale_ToTime22(const QLocale* self, struct miqt_string stringVal, int param2);
QDateTime* QLocale_ToDateTime22(const QLocale* self, struct miqt_string stringVal, int format);
struct miqt_string QLocale_MonthName2(const QLocale* self, int param1, int format);
struct miqt_string QLocale_StandaloneMonthName2(const QLocale* self, int param1, int format);
struct miqt_string QLocale_DayName2(const QLocale* self, int param1, int format);
struct miqt_string QLocale_StandaloneDayName2(const QLocale* self, int param1, int format);
struct miqt_string QLocale_CurrencySymbol1(const QLocale* self, int param1);
struct miqt_string QLocale_ToCurrencyString22(const QLocale* self, long long param1, struct miqt_string symbol);
struct miqt_string QLocale_ToCurrencyString23(const QLocale* self, unsigned long long param1, struct miqt_string symbol);
struct miqt_string QLocale_ToCurrencyString24(const QLocale* self, int16_t param1, struct miqt_string symbol);
struct miqt_string QLocale_ToCurrencyString25(const QLocale* self, uint16_t param1, struct miqt_string symbol);
struct miqt_string QLocale_ToCurrencyString26(const QLocale* self, int param1, struct miqt_string symbol);
struct miqt_string QLocale_ToCurrencyString27(const QLocale* self, unsigned int param1, struct miqt_string symbol);
struct miqt_string QLocale_ToCurrencyString28(const QLocale* self, double param1, struct miqt_string symbol);
struct miqt_string QLocale_ToCurrencyString29(const QLocale* self, float i, struct miqt_string symbol);
struct miqt_string QLocale_FormattedDataSize2(QLocale* self, long long bytes, int precision);
struct miqt_string QLocale_FormattedDataSize3(QLocale* self, long long bytes, int precision, int format);
struct miqt_string QLocale_FormattedDataSize22(const QLocale* self, long long bytes, int precision);
struct miqt_string QLocale_FormattedDataSize32(const QLocale* self, long long bytes, int precision, int format);
struct miqt_string QLocale_QuoteString2(const QLocale* self, struct miqt_string str, int style);
void QLocale_Delete(QLocale* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
