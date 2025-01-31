#include <QChildEvent>
#include <QEvent>
#include <QLocalServer>
#include <QLocalSocket>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTimerEvent>
#include <qlocalserver.h>
#include "gen_qlocalserver.h"

#ifdef __cplusplus
extern "C" {
#endif

void miqt_exec_callback_QLocalServer_NewConnection(intptr_t);
bool miqt_exec_callback_QLocalServer_HasPendingConnections(const QLocalServer*, intptr_t);
QLocalSocket* miqt_exec_callback_QLocalServer_NextPendingConnection(QLocalServer*, intptr_t);
void miqt_exec_callback_QLocalServer_IncomingConnection(QLocalServer*, intptr_t, uintptr_t);
bool miqt_exec_callback_QLocalServer_Event(QLocalServer*, intptr_t, QEvent*);
bool miqt_exec_callback_QLocalServer_EventFilter(QLocalServer*, intptr_t, QObject*, QEvent*);
void miqt_exec_callback_QLocalServer_TimerEvent(QLocalServer*, intptr_t, QTimerEvent*);
void miqt_exec_callback_QLocalServer_ChildEvent(QLocalServer*, intptr_t, QChildEvent*);
void miqt_exec_callback_QLocalServer_CustomEvent(QLocalServer*, intptr_t, QEvent*);
void miqt_exec_callback_QLocalServer_ConnectNotify(QLocalServer*, intptr_t, QMetaMethod*);
void miqt_exec_callback_QLocalServer_DisconnectNotify(QLocalServer*, intptr_t, QMetaMethod*);
#ifdef __cplusplus
} /* extern C */
#endif

class MiqtVirtualQLocalServer final : public QLocalServer {
public:

	MiqtVirtualQLocalServer(): QLocalServer() {};
	MiqtVirtualQLocalServer(QObject* parent): QLocalServer(parent) {};

	virtual ~MiqtVirtualQLocalServer() override = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HasPendingConnections = 0;

	// Subclass to allow providing a Go implementation
	virtual bool hasPendingConnections() const override {
		if (handle__HasPendingConnections == 0) {
			return QLocalServer::hasPendingConnections();
		}
		

		bool callback_return_value = miqt_exec_callback_QLocalServer_HasPendingConnections(this, handle__HasPendingConnections);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_HasPendingConnections() const {

		return QLocalServer::hasPendingConnections();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__NextPendingConnection = 0;

	// Subclass to allow providing a Go implementation
	virtual QLocalSocket* nextPendingConnection() override {
		if (handle__NextPendingConnection == 0) {
			return QLocalServer::nextPendingConnection();
		}
		

		QLocalSocket* callback_return_value = miqt_exec_callback_QLocalServer_NextPendingConnection(this, handle__NextPendingConnection);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QLocalSocket* virtualbase_NextPendingConnection() {

		return QLocalServer::nextPendingConnection();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__IncomingConnection = 0;

	// Subclass to allow providing a Go implementation
	virtual void incomingConnection(quintptr socketDescriptor) override {
		if (handle__IncomingConnection == 0) {
			QLocalServer::incomingConnection(socketDescriptor);
			return;
		}
		
		quintptr socketDescriptor_ret = socketDescriptor;
		uintptr_t sigval1 = static_cast<uintptr_t>(socketDescriptor_ret);

		miqt_exec_callback_QLocalServer_IncomingConnection(this, handle__IncomingConnection, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_IncomingConnection(uintptr_t socketDescriptor) {

		QLocalServer::incomingConnection(static_cast<quintptr>(socketDescriptor));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Event = 0;

	// Subclass to allow providing a Go implementation
	virtual bool event(QEvent* event) override {
		if (handle__Event == 0) {
			return QLocalServer::event(event);
		}
		
		QEvent* sigval1 = event;

		bool callback_return_value = miqt_exec_callback_QLocalServer_Event(this, handle__Event, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_Event(QEvent* event) {

		return QLocalServer::event(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__EventFilter = 0;

	// Subclass to allow providing a Go implementation
	virtual bool eventFilter(QObject* watched, QEvent* event) override {
		if (handle__EventFilter == 0) {
			return QLocalServer::eventFilter(watched, event);
		}
		
		QObject* sigval1 = watched;
		QEvent* sigval2 = event;

		bool callback_return_value = miqt_exec_callback_QLocalServer_EventFilter(this, handle__EventFilter, sigval1, sigval2);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_EventFilter(QObject* watched, QEvent* event) {

		return QLocalServer::eventFilter(watched, event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__TimerEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void timerEvent(QTimerEvent* event) override {
		if (handle__TimerEvent == 0) {
			QLocalServer::timerEvent(event);
			return;
		}
		
		QTimerEvent* sigval1 = event;

		miqt_exec_callback_QLocalServer_TimerEvent(this, handle__TimerEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_TimerEvent(QTimerEvent* event) {

		QLocalServer::timerEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ChildEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void childEvent(QChildEvent* event) override {
		if (handle__ChildEvent == 0) {
			QLocalServer::childEvent(event);
			return;
		}
		
		QChildEvent* sigval1 = event;

		miqt_exec_callback_QLocalServer_ChildEvent(this, handle__ChildEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ChildEvent(QChildEvent* event) {

		QLocalServer::childEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__CustomEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void customEvent(QEvent* event) override {
		if (handle__CustomEvent == 0) {
			QLocalServer::customEvent(event);
			return;
		}
		
		QEvent* sigval1 = event;

		miqt_exec_callback_QLocalServer_CustomEvent(this, handle__CustomEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_CustomEvent(QEvent* event) {

		QLocalServer::customEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ConnectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void connectNotify(const QMetaMethod& signal) override {
		if (handle__ConnectNotify == 0) {
			QLocalServer::connectNotify(signal);
			return;
		}
		
		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);

		miqt_exec_callback_QLocalServer_ConnectNotify(this, handle__ConnectNotify, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ConnectNotify(QMetaMethod* signal) {

		QLocalServer::connectNotify(*signal);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DisconnectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void disconnectNotify(const QMetaMethod& signal) override {
		if (handle__DisconnectNotify == 0) {
			QLocalServer::disconnectNotify(signal);
			return;
		}
		
		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);

		miqt_exec_callback_QLocalServer_DisconnectNotify(this, handle__DisconnectNotify, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DisconnectNotify(QMetaMethod* signal) {

		QLocalServer::disconnectNotify(*signal);

	}

};

QLocalServer* QLocalServer_new() {
	return new MiqtVirtualQLocalServer();
}

QLocalServer* QLocalServer_new2(QObject* parent) {
	return new MiqtVirtualQLocalServer(parent);
}

void QLocalServer_virtbase(QLocalServer* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QLocalServer_MetaObject(const QLocalServer* self) {
	return (QMetaObject*) self->metaObject();
}

void* QLocalServer_Metacast(QLocalServer* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QLocalServer_Tr(const char* s) {
	QString _ret = QLocalServer::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocalServer_TrUtf8(const char* s) {
	QString _ret = QLocalServer::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QLocalServer_NewConnection(QLocalServer* self) {
	self->newConnection();
}

void QLocalServer_connect_NewConnection(QLocalServer* self, intptr_t slot) {
	MiqtVirtualQLocalServer::connect(self, static_cast<void (QLocalServer::*)()>(&QLocalServer::newConnection), self, [=]() {
		miqt_exec_callback_QLocalServer_NewConnection(slot);
	});
}

void QLocalServer_Close(QLocalServer* self) {
	self->close();
}

struct miqt_string QLocalServer_ErrorString(const QLocalServer* self) {
	QString _ret = self->errorString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QLocalServer_HasPendingConnections(const QLocalServer* self) {
	return self->hasPendingConnections();
}

bool QLocalServer_IsListening(const QLocalServer* self) {
	return self->isListening();
}

bool QLocalServer_Listen(QLocalServer* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return self->listen(name_QString);
}

bool QLocalServer_ListenWithSocketDescriptor(QLocalServer* self, intptr_t socketDescriptor) {
	return self->listen((qintptr)(socketDescriptor));
}

int QLocalServer_MaxPendingConnections(const QLocalServer* self) {
	return self->maxPendingConnections();
}

QLocalSocket* QLocalServer_NextPendingConnection(QLocalServer* self) {
	return self->nextPendingConnection();
}

struct miqt_string QLocalServer_ServerName(const QLocalServer* self) {
	QString _ret = self->serverName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocalServer_FullServerName(const QLocalServer* self) {
	QString _ret = self->fullServerName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QLocalServer_RemoveServer(struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return QLocalServer::removeServer(name_QString);
}

int QLocalServer_ServerError(const QLocalServer* self) {
	QAbstractSocket::SocketError _ret = self->serverError();
	return static_cast<int>(_ret);
}

void QLocalServer_SetMaxPendingConnections(QLocalServer* self, int numConnections) {
	self->setMaxPendingConnections(static_cast<int>(numConnections));
}

bool QLocalServer_WaitForNewConnection(QLocalServer* self) {
	return self->waitForNewConnection();
}

void QLocalServer_SetSocketOptions(QLocalServer* self, int options) {
	self->setSocketOptions(static_cast<QLocalServer::SocketOptions>(options));
}

int QLocalServer_SocketOptions(const QLocalServer* self) {
	QLocalServer::SocketOptions _ret = self->socketOptions();
	return static_cast<int>(_ret);
}

intptr_t QLocalServer_SocketDescriptor(const QLocalServer* self) {
	qintptr _ret = self->socketDescriptor();
	return (intptr_t)(_ret);
}

struct miqt_string QLocalServer_Tr2(const char* s, const char* c) {
	QString _ret = QLocalServer::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocalServer_Tr3(const char* s, const char* c, int n) {
	QString _ret = QLocalServer::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocalServer_TrUtf82(const char* s, const char* c) {
	QString _ret = QLocalServer::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocalServer_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QLocalServer::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QLocalServer_WaitForNewConnection1(QLocalServer* self, int msec) {
	return self->waitForNewConnection(static_cast<int>(msec));
}

bool QLocalServer_WaitForNewConnection2(QLocalServer* self, int msec, bool* timedOut) {
	return self->waitForNewConnection(static_cast<int>(msec), timedOut);
}

bool QLocalServer_override_virtual_HasPendingConnections(void* self, intptr_t slot) {
	MiqtVirtualQLocalServer* self_cast = dynamic_cast<MiqtVirtualQLocalServer*>( (QLocalServer*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__HasPendingConnections = slot;
	return true;
}

bool QLocalServer_virtualbase_HasPendingConnections(const void* self) {
	return ( (const MiqtVirtualQLocalServer*)(self) )->virtualbase_HasPendingConnections();
}

bool QLocalServer_override_virtual_NextPendingConnection(void* self, intptr_t slot) {
	MiqtVirtualQLocalServer* self_cast = dynamic_cast<MiqtVirtualQLocalServer*>( (QLocalServer*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__NextPendingConnection = slot;
	return true;
}

QLocalSocket* QLocalServer_virtualbase_NextPendingConnection(void* self) {
	return ( (MiqtVirtualQLocalServer*)(self) )->virtualbase_NextPendingConnection();
}

bool QLocalServer_override_virtual_IncomingConnection(void* self, intptr_t slot) {
	MiqtVirtualQLocalServer* self_cast = dynamic_cast<MiqtVirtualQLocalServer*>( (QLocalServer*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__IncomingConnection = slot;
	return true;
}

void QLocalServer_virtualbase_IncomingConnection(void* self, uintptr_t socketDescriptor) {
	( (MiqtVirtualQLocalServer*)(self) )->virtualbase_IncomingConnection(socketDescriptor);
}

bool QLocalServer_override_virtual_Event(void* self, intptr_t slot) {
	MiqtVirtualQLocalServer* self_cast = dynamic_cast<MiqtVirtualQLocalServer*>( (QLocalServer*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__Event = slot;
	return true;
}

bool QLocalServer_virtualbase_Event(void* self, QEvent* event) {
	return ( (MiqtVirtualQLocalServer*)(self) )->virtualbase_Event(event);
}

bool QLocalServer_override_virtual_EventFilter(void* self, intptr_t slot) {
	MiqtVirtualQLocalServer* self_cast = dynamic_cast<MiqtVirtualQLocalServer*>( (QLocalServer*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__EventFilter = slot;
	return true;
}

bool QLocalServer_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event) {
	return ( (MiqtVirtualQLocalServer*)(self) )->virtualbase_EventFilter(watched, event);
}

bool QLocalServer_override_virtual_TimerEvent(void* self, intptr_t slot) {
	MiqtVirtualQLocalServer* self_cast = dynamic_cast<MiqtVirtualQLocalServer*>( (QLocalServer*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__TimerEvent = slot;
	return true;
}

void QLocalServer_virtualbase_TimerEvent(void* self, QTimerEvent* event) {
	( (MiqtVirtualQLocalServer*)(self) )->virtualbase_TimerEvent(event);
}

bool QLocalServer_override_virtual_ChildEvent(void* self, intptr_t slot) {
	MiqtVirtualQLocalServer* self_cast = dynamic_cast<MiqtVirtualQLocalServer*>( (QLocalServer*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__ChildEvent = slot;
	return true;
}

void QLocalServer_virtualbase_ChildEvent(void* self, QChildEvent* event) {
	( (MiqtVirtualQLocalServer*)(self) )->virtualbase_ChildEvent(event);
}

bool QLocalServer_override_virtual_CustomEvent(void* self, intptr_t slot) {
	MiqtVirtualQLocalServer* self_cast = dynamic_cast<MiqtVirtualQLocalServer*>( (QLocalServer*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__CustomEvent = slot;
	return true;
}

void QLocalServer_virtualbase_CustomEvent(void* self, QEvent* event) {
	( (MiqtVirtualQLocalServer*)(self) )->virtualbase_CustomEvent(event);
}

bool QLocalServer_override_virtual_ConnectNotify(void* self, intptr_t slot) {
	MiqtVirtualQLocalServer* self_cast = dynamic_cast<MiqtVirtualQLocalServer*>( (QLocalServer*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__ConnectNotify = slot;
	return true;
}

void QLocalServer_virtualbase_ConnectNotify(void* self, QMetaMethod* signal) {
	( (MiqtVirtualQLocalServer*)(self) )->virtualbase_ConnectNotify(signal);
}

bool QLocalServer_override_virtual_DisconnectNotify(void* self, intptr_t slot) {
	MiqtVirtualQLocalServer* self_cast = dynamic_cast<MiqtVirtualQLocalServer*>( (QLocalServer*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__DisconnectNotify = slot;
	return true;
}

void QLocalServer_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal) {
	( (MiqtVirtualQLocalServer*)(self) )->virtualbase_DisconnectNotify(signal);
}

void QLocalServer_Delete(QLocalServer* self) {
	delete self;
}

