#include <QFile>
#include <QFileDevice>
#include <QIODevice>
#include <QIODeviceBase>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTemporaryFile>
#include <qtemporaryfile.h>
#include "gen_qtemporaryfile.h"

#ifdef __cplusplus
extern "C" {
#endif

struct miqt_string miqt_exec_callback_QTemporaryFile_FileName(void*, intptr_t);
bool miqt_exec_callback_QTemporaryFile_OpenWithFlags(void*, intptr_t, int);
long long miqt_exec_callback_QTemporaryFile_Size(void*, intptr_t);
bool miqt_exec_callback_QTemporaryFile_Resize(void*, intptr_t, long long);
int miqt_exec_callback_QTemporaryFile_Permissions(void*, intptr_t);
bool miqt_exec_callback_QTemporaryFile_SetPermissions(void*, intptr_t, int);
#ifdef __cplusplus
} /* extern C */
#endif

class MiqtVirtualQTemporaryFile final : public QTemporaryFile {
public:

	MiqtVirtualQTemporaryFile(): QTemporaryFile() {};
	MiqtVirtualQTemporaryFile(const QString& templateName): QTemporaryFile(templateName) {};
	MiqtVirtualQTemporaryFile(QObject* parent): QTemporaryFile(parent) {};
	MiqtVirtualQTemporaryFile(const QString& templateName, QObject* parent): QTemporaryFile(templateName, parent) {};

	virtual ~MiqtVirtualQTemporaryFile() override = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__FileName = 0;

	// Subclass to allow providing a Go implementation
	virtual QString fileName() const override {
		if (handle__FileName == 0) {
			return QTemporaryFile::fileName();
		}
		

		struct miqt_string callback_return_value = miqt_exec_callback_QTemporaryFile_FileName(const_cast<MiqtVirtualQTemporaryFile*>(this), handle__FileName);
		QString callback_return_value_QString = QString::fromUtf8(callback_return_value.data, callback_return_value.len);

		return callback_return_value_QString;
	}

	// Wrapper to allow calling protected method
	struct miqt_string virtualbase_FileName() const {

		QString _ret = QTemporaryFile::fileName();
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _b = _ret.toUtf8();
		struct miqt_string _ms;
		_ms.len = _b.length();
		_ms.data = static_cast<char*>(malloc(_ms.len));
		memcpy(_ms.data, _b.data(), _ms.len);
		return _ms;

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__OpenWithFlags = 0;

	// Subclass to allow providing a Go implementation
	virtual bool open(QIODeviceBase::OpenMode flags) override {
		if (handle__OpenWithFlags == 0) {
			return QTemporaryFile::open(flags);
		}
		
		QIODeviceBase::OpenMode flags_ret = flags;
		int sigval1 = static_cast<int>(flags_ret);

		bool callback_return_value = miqt_exec_callback_QTemporaryFile_OpenWithFlags(this, handle__OpenWithFlags, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_OpenWithFlags(int flags) {

		return QTemporaryFile::open(static_cast<QIODeviceBase::OpenMode>(flags));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Size = 0;

	// Subclass to allow providing a Go implementation
	virtual qint64 size() const override {
		if (handle__Size == 0) {
			return QTemporaryFile::size();
		}
		

		long long callback_return_value = miqt_exec_callback_QTemporaryFile_Size(const_cast<MiqtVirtualQTemporaryFile*>(this), handle__Size);

		return static_cast<qint64>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	long long virtualbase_Size() const {

		qint64 _ret = QTemporaryFile::size();
		return static_cast<long long>(_ret);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Resize = 0;

	// Subclass to allow providing a Go implementation
	virtual bool resize(qint64 sz) override {
		if (handle__Resize == 0) {
			return QTemporaryFile::resize(sz);
		}
		
		qint64 sz_ret = sz;
		long long sigval1 = static_cast<long long>(sz_ret);

		bool callback_return_value = miqt_exec_callback_QTemporaryFile_Resize(this, handle__Resize, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_Resize(long long sz) {

		return QTemporaryFile::resize(static_cast<qint64>(sz));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Permissions = 0;

	// Subclass to allow providing a Go implementation
	virtual QFileDevice::Permissions permissions() const override {
		if (handle__Permissions == 0) {
			return QTemporaryFile::permissions();
		}
		

		int callback_return_value = miqt_exec_callback_QTemporaryFile_Permissions(const_cast<MiqtVirtualQTemporaryFile*>(this), handle__Permissions);

		return static_cast<QFileDevice::Permissions>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_Permissions() const {

		QFileDevice::Permissions _ret = QTemporaryFile::permissions();
		return static_cast<int>(_ret);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetPermissions = 0;

	// Subclass to allow providing a Go implementation
	virtual bool setPermissions(QFileDevice::Permissions permissionSpec) override {
		if (handle__SetPermissions == 0) {
			return QTemporaryFile::setPermissions(permissionSpec);
		}
		
		QFileDevice::Permissions permissionSpec_ret = permissionSpec;
		int sigval1 = static_cast<int>(permissionSpec_ret);

		bool callback_return_value = miqt_exec_callback_QTemporaryFile_SetPermissions(this, handle__SetPermissions, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_SetPermissions(int permissionSpec) {

		return QTemporaryFile::setPermissions(static_cast<QFileDevice::Permissions>(permissionSpec));

	}

};

QTemporaryFile* QTemporaryFile_new() {
	return new MiqtVirtualQTemporaryFile();
}

QTemporaryFile* QTemporaryFile_new2(struct miqt_string templateName) {
	QString templateName_QString = QString::fromUtf8(templateName.data, templateName.len);
	return new MiqtVirtualQTemporaryFile(templateName_QString);
}

QTemporaryFile* QTemporaryFile_new3(QObject* parent) {
	return new MiqtVirtualQTemporaryFile(parent);
}

QTemporaryFile* QTemporaryFile_new4(struct miqt_string templateName, QObject* parent) {
	QString templateName_QString = QString::fromUtf8(templateName.data, templateName.len);
	return new MiqtVirtualQTemporaryFile(templateName_QString, parent);
}

void QTemporaryFile_virtbase(QTemporaryFile* src, QFile** outptr_QFile) {
	*outptr_QFile = static_cast<QFile*>(src);
}

QMetaObject* QTemporaryFile_MetaObject(const QTemporaryFile* self) {
	return (QMetaObject*) self->metaObject();
}

void* QTemporaryFile_Metacast(QTemporaryFile* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QTemporaryFile_Tr(const char* s) {
	QString _ret = QTemporaryFile::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QTemporaryFile_AutoRemove(const QTemporaryFile* self) {
	return self->autoRemove();
}

void QTemporaryFile_SetAutoRemove(QTemporaryFile* self, bool b) {
	self->setAutoRemove(b);
}

bool QTemporaryFile_Open(QTemporaryFile* self) {
	return self->open();
}

struct miqt_string QTemporaryFile_FileName(const QTemporaryFile* self) {
	QString _ret = self->fileName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QTemporaryFile_FileTemplate(const QTemporaryFile* self) {
	QString _ret = self->fileTemplate();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QTemporaryFile_SetFileTemplate(QTemporaryFile* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	self->setFileTemplate(name_QString);
}

bool QTemporaryFile_Rename(QTemporaryFile* self, struct miqt_string newName) {
	QString newName_QString = QString::fromUtf8(newName.data, newName.len);
	return self->rename(newName_QString);
}

QTemporaryFile* QTemporaryFile_CreateNativeFile(struct miqt_string fileName) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	return QTemporaryFile::createNativeFile(fileName_QString);
}

QTemporaryFile* QTemporaryFile_CreateNativeFileWithFile(QFile* file) {
	return QTemporaryFile::createNativeFile(*file);
}

struct miqt_string QTemporaryFile_Tr2(const char* s, const char* c) {
	QString _ret = QTemporaryFile::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QTemporaryFile_Tr3(const char* s, const char* c, int n) {
	QString _ret = QTemporaryFile::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QTemporaryFile_override_virtual_FileName(void* self, intptr_t slot) {
	MiqtVirtualQTemporaryFile* self_cast = dynamic_cast<MiqtVirtualQTemporaryFile*>( (QTemporaryFile*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__FileName = slot;
	return true;
}

struct miqt_string QTemporaryFile_virtualbase_FileName(const void* self) {
	return ( (const MiqtVirtualQTemporaryFile*)(self) )->virtualbase_FileName();
}

bool QTemporaryFile_override_virtual_OpenWithFlags(void* self, intptr_t slot) {
	MiqtVirtualQTemporaryFile* self_cast = dynamic_cast<MiqtVirtualQTemporaryFile*>( (QTemporaryFile*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__OpenWithFlags = slot;
	return true;
}

bool QTemporaryFile_virtualbase_OpenWithFlags(void* self, int flags) {
	return ( (MiqtVirtualQTemporaryFile*)(self) )->virtualbase_OpenWithFlags(flags);
}

bool QTemporaryFile_override_virtual_Size(void* self, intptr_t slot) {
	MiqtVirtualQTemporaryFile* self_cast = dynamic_cast<MiqtVirtualQTemporaryFile*>( (QTemporaryFile*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__Size = slot;
	return true;
}

long long QTemporaryFile_virtualbase_Size(const void* self) {
	return ( (const MiqtVirtualQTemporaryFile*)(self) )->virtualbase_Size();
}

bool QTemporaryFile_override_virtual_Resize(void* self, intptr_t slot) {
	MiqtVirtualQTemporaryFile* self_cast = dynamic_cast<MiqtVirtualQTemporaryFile*>( (QTemporaryFile*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__Resize = slot;
	return true;
}

bool QTemporaryFile_virtualbase_Resize(void* self, long long sz) {
	return ( (MiqtVirtualQTemporaryFile*)(self) )->virtualbase_Resize(sz);
}

bool QTemporaryFile_override_virtual_Permissions(void* self, intptr_t slot) {
	MiqtVirtualQTemporaryFile* self_cast = dynamic_cast<MiqtVirtualQTemporaryFile*>( (QTemporaryFile*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__Permissions = slot;
	return true;
}

int QTemporaryFile_virtualbase_Permissions(const void* self) {
	return ( (const MiqtVirtualQTemporaryFile*)(self) )->virtualbase_Permissions();
}

bool QTemporaryFile_override_virtual_SetPermissions(void* self, intptr_t slot) {
	MiqtVirtualQTemporaryFile* self_cast = dynamic_cast<MiqtVirtualQTemporaryFile*>( (QTemporaryFile*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__SetPermissions = slot;
	return true;
}

bool QTemporaryFile_virtualbase_SetPermissions(void* self, int permissionSpec) {
	return ( (MiqtVirtualQTemporaryFile*)(self) )->virtualbase_SetPermissions(permissionSpec);
}

void QTemporaryFile_Delete(QTemporaryFile* self) {
	delete self;
}

