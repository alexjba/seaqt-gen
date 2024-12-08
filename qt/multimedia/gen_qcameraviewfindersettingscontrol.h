#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QCAMERAVIEWFINDERSETTINGSCONTROL_H
#define MIQT_QT_MULTIMEDIA_GEN_QCAMERAVIEWFINDERSETTINGSCONTROL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QCameraViewfinderSettings;
class QCameraViewfinderSettingsControl;
class QCameraViewfinderSettingsControl2;
class QMediaControl;
class QMetaObject;
class QObject;
class QVariant;
#else
typedef struct QCameraViewfinderSettings QCameraViewfinderSettings;
typedef struct QCameraViewfinderSettingsControl QCameraViewfinderSettingsControl;
typedef struct QCameraViewfinderSettingsControl2 QCameraViewfinderSettingsControl2;
typedef struct QMediaControl QMediaControl;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QVariant QVariant;
#endif

void QCameraViewfinderSettingsControl_virtbase(QCameraViewfinderSettingsControl* src, QMediaControl** outptr_QMediaControl);
QMetaObject* QCameraViewfinderSettingsControl_MetaObject(const QCameraViewfinderSettingsControl* self);
void* QCameraViewfinderSettingsControl_Metacast(QCameraViewfinderSettingsControl* self, const char* param1);
struct miqt_string QCameraViewfinderSettingsControl_Tr(const char* s);
struct miqt_string QCameraViewfinderSettingsControl_TrUtf8(const char* s);
bool QCameraViewfinderSettingsControl_IsViewfinderParameterSupported(const QCameraViewfinderSettingsControl* self, int parameter);
QVariant* QCameraViewfinderSettingsControl_ViewfinderParameter(const QCameraViewfinderSettingsControl* self, int parameter);
void QCameraViewfinderSettingsControl_SetViewfinderParameter(QCameraViewfinderSettingsControl* self, int parameter, QVariant* value);
struct miqt_string QCameraViewfinderSettingsControl_Tr2(const char* s, const char* c);
struct miqt_string QCameraViewfinderSettingsControl_Tr3(const char* s, const char* c, int n);
struct miqt_string QCameraViewfinderSettingsControl_TrUtf82(const char* s, const char* c);
struct miqt_string QCameraViewfinderSettingsControl_TrUtf83(const char* s, const char* c, int n);
void QCameraViewfinderSettingsControl_Delete(QCameraViewfinderSettingsControl* self, bool isSubclass);

void QCameraViewfinderSettingsControl2_virtbase(QCameraViewfinderSettingsControl2* src, QMediaControl** outptr_QMediaControl);
QMetaObject* QCameraViewfinderSettingsControl2_MetaObject(const QCameraViewfinderSettingsControl2* self);
void* QCameraViewfinderSettingsControl2_Metacast(QCameraViewfinderSettingsControl2* self, const char* param1);
struct miqt_string QCameraViewfinderSettingsControl2_Tr(const char* s);
struct miqt_string QCameraViewfinderSettingsControl2_TrUtf8(const char* s);
struct miqt_array /* of QCameraViewfinderSettings* */  QCameraViewfinderSettingsControl2_SupportedViewfinderSettings(const QCameraViewfinderSettingsControl2* self);
QCameraViewfinderSettings* QCameraViewfinderSettingsControl2_ViewfinderSettings(const QCameraViewfinderSettingsControl2* self);
void QCameraViewfinderSettingsControl2_SetViewfinderSettings(QCameraViewfinderSettingsControl2* self, QCameraViewfinderSettings* settings);
struct miqt_string QCameraViewfinderSettingsControl2_Tr2(const char* s, const char* c);
struct miqt_string QCameraViewfinderSettingsControl2_Tr3(const char* s, const char* c, int n);
struct miqt_string QCameraViewfinderSettingsControl2_TrUtf82(const char* s, const char* c);
struct miqt_string QCameraViewfinderSettingsControl2_TrUtf83(const char* s, const char* c, int n);
void QCameraViewfinderSettingsControl2_Delete(QCameraViewfinderSettingsControl2* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
