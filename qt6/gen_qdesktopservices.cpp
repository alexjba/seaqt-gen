#include <QDesktopServices>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QUrl>
#include <qdesktopservices.h>
#include "gen_qdesktopservices.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
} /* extern C */
#endif

bool QDesktopServices_OpenUrl(QUrl* url) {
	return QDesktopServices::openUrl(*url);
}

void QDesktopServices_SetUrlHandler(struct miqt_string scheme, QObject* receiver, const char* method) {
	QString scheme_QString = QString::fromUtf8(scheme.data, scheme.len);
	QDesktopServices::setUrlHandler(scheme_QString, receiver, method);
}

void QDesktopServices_UnsetUrlHandler(struct miqt_string scheme) {
	QString scheme_QString = QString::fromUtf8(scheme.data, scheme.len);
	QDesktopServices::unsetUrlHandler(scheme_QString);
}

void QDesktopServices_Delete(QDesktopServices* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QDesktopServices*>( self );
	} else {
		delete self;
	}
}

