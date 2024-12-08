#include <QByteArray>
#include <QChar>
#include <QFont>
#include <QImage>
#include <QList>
#include <QPainterPath>
#include <QPointF>
#include <QRawFont>
#include <QRectF>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTransform>
#include <qrawfont.h>
#include "gen_qrawfont.h"
#include "_cgo_export.h"

QRawFont* QRawFont_new() {
	return new QRawFont();
}

QRawFont* QRawFont_new2(struct miqt_string fileName, double pixelSize) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	return new QRawFont(fileName_QString, static_cast<qreal>(pixelSize));
}

QRawFont* QRawFont_new3(struct miqt_string fontData, double pixelSize) {
	QByteArray fontData_QByteArray(fontData.data, fontData.len);
	return new QRawFont(fontData_QByteArray, static_cast<qreal>(pixelSize));
}

QRawFont* QRawFont_new4(QRawFont* other) {
	return new QRawFont(*other);
}

QRawFont* QRawFont_new5(struct miqt_string fileName, double pixelSize, int hintingPreference) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	return new QRawFont(fileName_QString, static_cast<qreal>(pixelSize), static_cast<QFont::HintingPreference>(hintingPreference));
}

QRawFont* QRawFont_new6(struct miqt_string fontData, double pixelSize, int hintingPreference) {
	QByteArray fontData_QByteArray(fontData.data, fontData.len);
	return new QRawFont(fontData_QByteArray, static_cast<qreal>(pixelSize), static_cast<QFont::HintingPreference>(hintingPreference));
}

void QRawFont_OperatorAssign(QRawFont* self, QRawFont* other) {
	self->operator=(*other);
}

void QRawFont_Swap(QRawFont* self, QRawFont* other) {
	self->swap(*other);
}

bool QRawFont_IsValid(const QRawFont* self) {
	return self->isValid();
}

bool QRawFont_OperatorEqual(const QRawFont* self, QRawFont* other) {
	return (*self == *other);
}

bool QRawFont_OperatorNotEqual(const QRawFont* self, QRawFont* other) {
	return (*self != *other);
}

struct miqt_string QRawFont_FamilyName(const QRawFont* self) {
	QString _ret = self->familyName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QRawFont_StyleName(const QRawFont* self) {
	QString _ret = self->styleName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QRawFont_Style(const QRawFont* self) {
	QFont::Style _ret = self->style();
	return static_cast<int>(_ret);
}

int QRawFont_Weight(const QRawFont* self) {
	return self->weight();
}

struct miqt_array /* of unsigned int */  QRawFont_GlyphIndexesForString(const QRawFont* self, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QVector<quint32> _ret = self->glyphIndexesForString(text_QString);
	// Convert QList<> from C++ memory to manually-managed C memory
	unsigned int* _arr = static_cast<unsigned int*>(malloc(sizeof(unsigned int) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of QPointF* */  QRawFont_AdvancesForGlyphIndexes(const QRawFont* self, struct miqt_array /* of unsigned int */  glyphIndexes) {
	QVector<quint32> glyphIndexes_QList;
	glyphIndexes_QList.reserve(glyphIndexes.len);
	unsigned int* glyphIndexes_arr = static_cast<unsigned int*>(glyphIndexes.data);
	for(size_t i = 0; i < glyphIndexes.len; ++i) {
		glyphIndexes_QList.push_back(static_cast<unsigned int>(glyphIndexes_arr[i]));
	}
	QVector<QPointF> _ret = self->advancesForGlyphIndexes(glyphIndexes_QList);
	// Convert QList<> from C++ memory to manually-managed C memory
	QPointF** _arr = static_cast<QPointF**>(malloc(sizeof(QPointF*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QPointF(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of QPointF* */  QRawFont_AdvancesForGlyphIndexes2(const QRawFont* self, struct miqt_array /* of unsigned int */  glyphIndexes, int layoutFlags) {
	QVector<quint32> glyphIndexes_QList;
	glyphIndexes_QList.reserve(glyphIndexes.len);
	unsigned int* glyphIndexes_arr = static_cast<unsigned int*>(glyphIndexes.data);
	for(size_t i = 0; i < glyphIndexes.len; ++i) {
		glyphIndexes_QList.push_back(static_cast<unsigned int>(glyphIndexes_arr[i]));
	}
	QVector<QPointF> _ret = self->advancesForGlyphIndexes(glyphIndexes_QList, static_cast<QRawFont::LayoutFlags>(layoutFlags));
	// Convert QList<> from C++ memory to manually-managed C memory
	QPointF** _arr = static_cast<QPointF**>(malloc(sizeof(QPointF*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QPointF(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

bool QRawFont_GlyphIndexesForChars(const QRawFont* self, QChar* chars, int numChars, unsigned int* glyphIndexes, int* numGlyphs) {
	return self->glyphIndexesForChars(chars, static_cast<int>(numChars), static_cast<quint32*>(glyphIndexes), static_cast<int*>(numGlyphs));
}

bool QRawFont_AdvancesForGlyphIndexes3(const QRawFont* self, const unsigned int* glyphIndexes, QPointF* advances, int numGlyphs) {
	return self->advancesForGlyphIndexes(static_cast<const quint32*>(glyphIndexes), advances, static_cast<int>(numGlyphs));
}

bool QRawFont_AdvancesForGlyphIndexes4(const QRawFont* self, const unsigned int* glyphIndexes, QPointF* advances, int numGlyphs, int layoutFlags) {
	return self->advancesForGlyphIndexes(static_cast<const quint32*>(glyphIndexes), advances, static_cast<int>(numGlyphs), static_cast<QRawFont::LayoutFlags>(layoutFlags));
}

QImage* QRawFont_AlphaMapForGlyph(const QRawFont* self, unsigned int glyphIndex) {
	return new QImage(self->alphaMapForGlyph(static_cast<quint32>(glyphIndex)));
}

QPainterPath* QRawFont_PathForGlyph(const QRawFont* self, unsigned int glyphIndex) {
	return new QPainterPath(self->pathForGlyph(static_cast<quint32>(glyphIndex)));
}

QRectF* QRawFont_BoundingRect(const QRawFont* self, unsigned int glyphIndex) {
	return new QRectF(self->boundingRect(static_cast<quint32>(glyphIndex)));
}

void QRawFont_SetPixelSize(QRawFont* self, double pixelSize) {
	self->setPixelSize(static_cast<qreal>(pixelSize));
}

double QRawFont_PixelSize(const QRawFont* self) {
	qreal _ret = self->pixelSize();
	return static_cast<double>(_ret);
}

int QRawFont_HintingPreference(const QRawFont* self) {
	QFont::HintingPreference _ret = self->hintingPreference();
	return static_cast<int>(_ret);
}

double QRawFont_Ascent(const QRawFont* self) {
	qreal _ret = self->ascent();
	return static_cast<double>(_ret);
}

double QRawFont_CapHeight(const QRawFont* self) {
	qreal _ret = self->capHeight();
	return static_cast<double>(_ret);
}

double QRawFont_Descent(const QRawFont* self) {
	qreal _ret = self->descent();
	return static_cast<double>(_ret);
}

double QRawFont_Leading(const QRawFont* self) {
	qreal _ret = self->leading();
	return static_cast<double>(_ret);
}

double QRawFont_XHeight(const QRawFont* self) {
	qreal _ret = self->xHeight();
	return static_cast<double>(_ret);
}

double QRawFont_AverageCharWidth(const QRawFont* self) {
	qreal _ret = self->averageCharWidth();
	return static_cast<double>(_ret);
}

double QRawFont_MaxCharWidth(const QRawFont* self) {
	qreal _ret = self->maxCharWidth();
	return static_cast<double>(_ret);
}

double QRawFont_LineThickness(const QRawFont* self) {
	qreal _ret = self->lineThickness();
	return static_cast<double>(_ret);
}

double QRawFont_UnderlinePosition(const QRawFont* self) {
	qreal _ret = self->underlinePosition();
	return static_cast<double>(_ret);
}

double QRawFont_UnitsPerEm(const QRawFont* self) {
	qreal _ret = self->unitsPerEm();
	return static_cast<double>(_ret);
}

void QRawFont_LoadFromFile(QRawFont* self, struct miqt_string fileName, double pixelSize, int hintingPreference) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	self->loadFromFile(fileName_QString, static_cast<qreal>(pixelSize), static_cast<QFont::HintingPreference>(hintingPreference));
}

void QRawFont_LoadFromData(QRawFont* self, struct miqt_string fontData, double pixelSize, int hintingPreference) {
	QByteArray fontData_QByteArray(fontData.data, fontData.len);
	self->loadFromData(fontData_QByteArray, static_cast<qreal>(pixelSize), static_cast<QFont::HintingPreference>(hintingPreference));
}

bool QRawFont_SupportsCharacter(const QRawFont* self, unsigned int ucs4) {
	return self->supportsCharacter(static_cast<uint>(ucs4));
}

bool QRawFont_SupportsCharacterWithCharacter(const QRawFont* self, QChar* character) {
	return self->supportsCharacter(*character);
}

struct miqt_array /* of int */  QRawFont_SupportedWritingSystems(const QRawFont* self) {
	QList<QFontDatabase::WritingSystem> _ret = self->supportedWritingSystems();
	// Convert QList<> from C++ memory to manually-managed C memory
	int* _arr = static_cast<int*>(malloc(sizeof(int) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QFontDatabase::WritingSystem _lv_ret = _ret[i];
		_arr[i] = static_cast<int>(_lv_ret);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_string QRawFont_FontTable(const QRawFont* self, const char* tagName) {
	QByteArray _qb = self->fontTable(tagName);
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

QRawFont* QRawFont_FromFont(QFont* font) {
	return new QRawFont(QRawFont::fromFont(*font));
}

QImage* QRawFont_AlphaMapForGlyph2(const QRawFont* self, unsigned int glyphIndex, int antialiasingType) {
	return new QImage(self->alphaMapForGlyph(static_cast<quint32>(glyphIndex), static_cast<QRawFont::AntialiasingType>(antialiasingType)));
}

QImage* QRawFont_AlphaMapForGlyph3(const QRawFont* self, unsigned int glyphIndex, int antialiasingType, QTransform* transform) {
	return new QImage(self->alphaMapForGlyph(static_cast<quint32>(glyphIndex), static_cast<QRawFont::AntialiasingType>(antialiasingType), *transform));
}

QRawFont* QRawFont_FromFont2(QFont* font, int writingSystem) {
	return new QRawFont(QRawFont::fromFont(*font, static_cast<QFontDatabase::WritingSystem>(writingSystem)));
}

void QRawFont_Delete(QRawFont* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QRawFont*>( self );
	} else {
		delete self;
	}
}

