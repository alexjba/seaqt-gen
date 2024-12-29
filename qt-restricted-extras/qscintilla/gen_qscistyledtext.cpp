#include <QString>
#include <QByteArray>
#include <cstring>
#include <qscistyledtext.h>
#include "gen_qscistyledtext.h"

#ifndef _Bool
#define _Bool bool
#endif
#include "_cgo_export.h"

QsciStyledText* QsciStyledText_new(struct miqt_string text, int style) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return new QsciStyledText(text_QString, static_cast<int>(style));
}

QsciStyledText* QsciStyledText_new2(struct miqt_string text, QsciStyle* style) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return new QsciStyledText(text_QString, *style);
}

QsciStyledText* QsciStyledText_new3(QsciStyledText* param1) {
	return new QsciStyledText(*param1);
}

void QsciStyledText_Apply(const QsciStyledText* self, QsciScintillaBase* sci) {
	self->apply(sci);
}

struct miqt_string QsciStyledText_Text(const QsciStyledText* self) {
	const QString _ret = self->text();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QsciStyledText_Style(const QsciStyledText* self) {
	return self->style();
}

void QsciStyledText_Delete(QsciStyledText* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QsciStyledText*>( self );
	} else {
		delete self;
	}
}

