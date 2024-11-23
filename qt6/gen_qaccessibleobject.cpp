#define WORKAROUND_INNER_CLASS_DEFINITION_QAccessible__State
#include <QAccessibleApplication>
#include <QAccessibleInterface>
#include <QAccessibleObject>
#include <QObject>
#include <QRect>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWindow>
#include <qaccessibleobject.h>
#include "gen_qaccessibleobject.h"
#include "_cgo_export.h"

bool QAccessibleObject_IsValid(const QAccessibleObject* self) {
	return self->isValid();
}

QObject* QAccessibleObject_Object(const QAccessibleObject* self) {
	return self->object();
}

QRect* QAccessibleObject_Rect(const QAccessibleObject* self) {
	return new QRect(self->rect());
}

void QAccessibleObject_SetText(QAccessibleObject* self, int t, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->setText(static_cast<QAccessible::Text>(t), text_QString);
}

QAccessibleInterface* QAccessibleObject_ChildAt(const QAccessibleObject* self, int x, int y) {
	return self->childAt(static_cast<int>(x), static_cast<int>(y));
}

void QAccessibleApplication_new(QAccessibleApplication** outptr_QAccessibleApplication, QAccessibleObject** outptr_QAccessibleObject, QAccessibleInterface** outptr_QAccessibleInterface) {
	QAccessibleApplication* ret = new QAccessibleApplication();
	*outptr_QAccessibleApplication = ret;
	*outptr_QAccessibleObject = static_cast<QAccessibleObject*>(ret);
	*outptr_QAccessibleInterface = static_cast<QAccessibleInterface*>(ret);
}

QWindow* QAccessibleApplication_Window(const QAccessibleApplication* self) {
	return self->window();
}

int QAccessibleApplication_ChildCount(const QAccessibleApplication* self) {
	return self->childCount();
}

int QAccessibleApplication_IndexOfChild(const QAccessibleApplication* self, QAccessibleInterface* param1) {
	return self->indexOfChild(param1);
}

QAccessibleInterface* QAccessibleApplication_FocusChild(const QAccessibleApplication* self) {
	return self->focusChild();
}

QAccessibleInterface* QAccessibleApplication_Parent(const QAccessibleApplication* self) {
	return self->parent();
}

QAccessibleInterface* QAccessibleApplication_Child(const QAccessibleApplication* self, int index) {
	return self->child(static_cast<int>(index));
}

struct miqt_string QAccessibleApplication_Text(const QAccessibleApplication* self, int t) {
	QString _ret = self->text(static_cast<QAccessible::Text>(t));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QAccessibleApplication_Role(const QAccessibleApplication* self) {
	QAccessible::Role _ret = self->role();
	return static_cast<int>(_ret);
}

QAccessible__State* QAccessibleApplication_State(const QAccessibleApplication* self) {
	return new QAccessible::State(self->state());
}

void QAccessibleApplication_Delete(QAccessibleApplication* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QAccessibleApplication*>( self );
	} else {
		delete self;
	}
}

