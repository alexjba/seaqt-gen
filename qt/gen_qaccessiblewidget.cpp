#define WORKAROUND_INNER_CLASS_DEFINITION_QAccessible__State
#include <QAccessibleActionInterface>
#include <QAccessibleInterface>
#include <QAccessibleObject>
#include <QAccessibleWidget>
#include <QColor>
#include <QList>
#include <QPair>
#include <QRect>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <QWindow>
#include <qaccessiblewidget.h>
#include "gen_qaccessiblewidget.h"
#include "_cgo_export.h"

void QAccessibleWidget_new(QWidget* o, QAccessibleWidget** outptr_QAccessibleWidget, QAccessibleObject** outptr_QAccessibleObject, QAccessibleInterface** outptr_QAccessibleInterface, QAccessibleActionInterface** outptr_QAccessibleActionInterface) {
	QAccessibleWidget* ret = new QAccessibleWidget(o);
	*outptr_QAccessibleWidget = ret;
	*outptr_QAccessibleObject = static_cast<QAccessibleObject*>(ret);
	*outptr_QAccessibleInterface = static_cast<QAccessibleInterface*>(ret);
	*outptr_QAccessibleActionInterface = static_cast<QAccessibleActionInterface*>(ret);
}

void QAccessibleWidget_new2(QWidget* o, int r, QAccessibleWidget** outptr_QAccessibleWidget, QAccessibleObject** outptr_QAccessibleObject, QAccessibleInterface** outptr_QAccessibleInterface, QAccessibleActionInterface** outptr_QAccessibleActionInterface) {
	QAccessibleWidget* ret = new QAccessibleWidget(o, static_cast<QAccessible::Role>(r));
	*outptr_QAccessibleWidget = ret;
	*outptr_QAccessibleObject = static_cast<QAccessibleObject*>(ret);
	*outptr_QAccessibleInterface = static_cast<QAccessibleInterface*>(ret);
	*outptr_QAccessibleActionInterface = static_cast<QAccessibleActionInterface*>(ret);
}

void QAccessibleWidget_new3(QWidget* o, int r, struct miqt_string name, QAccessibleWidget** outptr_QAccessibleWidget, QAccessibleObject** outptr_QAccessibleObject, QAccessibleInterface** outptr_QAccessibleInterface, QAccessibleActionInterface** outptr_QAccessibleActionInterface) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	QAccessibleWidget* ret = new QAccessibleWidget(o, static_cast<QAccessible::Role>(r), name_QString);
	*outptr_QAccessibleWidget = ret;
	*outptr_QAccessibleObject = static_cast<QAccessibleObject*>(ret);
	*outptr_QAccessibleInterface = static_cast<QAccessibleInterface*>(ret);
	*outptr_QAccessibleActionInterface = static_cast<QAccessibleActionInterface*>(ret);
}

bool QAccessibleWidget_IsValid(const QAccessibleWidget* self) {
	return self->isValid();
}

QWindow* QAccessibleWidget_Window(const QAccessibleWidget* self) {
	return self->window();
}

int QAccessibleWidget_ChildCount(const QAccessibleWidget* self) {
	return self->childCount();
}

int QAccessibleWidget_IndexOfChild(const QAccessibleWidget* self, QAccessibleInterface* child) {
	return self->indexOfChild(child);
}

struct miqt_array /* of struct miqt_map  tuple of QAccessibleInterface* and int   */  QAccessibleWidget_Relations(const QAccessibleWidget* self, int match) {
	QVector<QPair<QAccessibleInterface *, QAccessible::Relation>> _ret = self->relations(static_cast<QAccessible::Relation>(match));
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_map /* tuple of QAccessibleInterface* and int */ * _arr = static_cast<struct miqt_map /* tuple of QAccessibleInterface* and int */ *>(malloc(sizeof(struct miqt_map /* tuple of QAccessibleInterface* and int */ ) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QPair<QAccessibleInterface*, QFlags<QAccessible::RelationFlag>> _vv_ret = _ret[i];
		// Convert QPair<> from C++ memory to manually-managed C memory
		QAccessibleInterface** _vv_first_arr = static_cast<QAccessibleInterface**>(malloc(sizeof(QAccessibleInterface*)));
		int* _vv_second_arr = static_cast<int*>(malloc(sizeof(int)));
		_vv_first_arr[0] = _vv_ret.first;
		QFlags<QAccessible::RelationFlag> _vv_second_ret = _vv_ret.second;
		_vv_second_arr[0] = static_cast<int>(_vv_second_ret);
		struct miqt_map _vv_out;
		_vv_out.len = 1;
		_vv_out.keys = static_cast<void*>(_vv_first_arr);
		_vv_out.values = static_cast<void*>(_vv_second_arr);
		_arr[i] = _vv_out;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

QAccessibleInterface* QAccessibleWidget_FocusChild(const QAccessibleWidget* self) {
	return self->focusChild();
}

QRect* QAccessibleWidget_Rect(const QAccessibleWidget* self) {
	return new QRect(self->rect());
}

QAccessibleInterface* QAccessibleWidget_Parent(const QAccessibleWidget* self) {
	return self->parent();
}

QAccessibleInterface* QAccessibleWidget_Child(const QAccessibleWidget* self, int index) {
	return self->child(static_cast<int>(index));
}

struct miqt_string QAccessibleWidget_Text(const QAccessibleWidget* self, int t) {
	QString _ret = self->text(static_cast<QAccessible::Text>(t));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QAccessibleWidget_Role(const QAccessibleWidget* self) {
	QAccessible::Role _ret = self->role();
	return static_cast<int>(_ret);
}

QAccessible__State* QAccessibleWidget_State(const QAccessibleWidget* self) {
	return new QAccessible::State(self->state());
}

QColor* QAccessibleWidget_ForegroundColor(const QAccessibleWidget* self) {
	return new QColor(self->foregroundColor());
}

QColor* QAccessibleWidget_BackgroundColor(const QAccessibleWidget* self) {
	return new QColor(self->backgroundColor());
}

void* QAccessibleWidget_InterfaceCast(QAccessibleWidget* self, int t) {
	return self->interface_cast(static_cast<QAccessible::InterfaceType>(t));
}

struct miqt_array /* of struct miqt_string */  QAccessibleWidget_ActionNames(const QAccessibleWidget* self) {
	QStringList _ret = self->actionNames();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

void QAccessibleWidget_DoAction(QAccessibleWidget* self, struct miqt_string actionName) {
	QString actionName_QString = QString::fromUtf8(actionName.data, actionName.len);
	self->doAction(actionName_QString);
}

struct miqt_array /* of struct miqt_string */  QAccessibleWidget_KeyBindingsForAction(const QAccessibleWidget* self, struct miqt_string actionName) {
	QString actionName_QString = QString::fromUtf8(actionName.data, actionName.len);
	QStringList _ret = self->keyBindingsForAction(actionName_QString);
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

