#pragma once
#ifndef MIQT_QT6_GEN_QFILEDEVICE_H
#define MIQT_QT6_GEN_QFILEDEVICE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QDateTime;
class QFileDevice;
class QIODevice;
class QIODeviceBase;
class QMetaObject;
class QObject;
#else
typedef struct QDateTime QDateTime;
typedef struct QFileDevice QFileDevice;
typedef struct QIODevice QIODevice;
typedef struct QIODeviceBase QIODeviceBase;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
#endif

void QFileDevice_virtbase(QFileDevice* src, QIODevice** outptr_QIODevice);
QMetaObject* QFileDevice_MetaObject(const QFileDevice* self);
void* QFileDevice_Metacast(QFileDevice* self, const char* param1);
struct miqt_string QFileDevice_Tr(const char* s);
int QFileDevice_Error(const QFileDevice* self);
void QFileDevice_UnsetError(QFileDevice* self);
void QFileDevice_Close(QFileDevice* self);
bool QFileDevice_IsSequential(const QFileDevice* self);
int QFileDevice_Handle(const QFileDevice* self);
struct miqt_string QFileDevice_FileName(const QFileDevice* self);
long long QFileDevice_Pos(const QFileDevice* self);
bool QFileDevice_Seek(QFileDevice* self, long long offset);
bool QFileDevice_AtEnd(const QFileDevice* self);
bool QFileDevice_Flush(QFileDevice* self);
long long QFileDevice_Size(const QFileDevice* self);
bool QFileDevice_Resize(QFileDevice* self, long long sz);
int QFileDevice_Permissions(const QFileDevice* self);
bool QFileDevice_SetPermissions(QFileDevice* self, int permissionSpec);
unsigned char* QFileDevice_Map(QFileDevice* self, long long offset, long long size);
bool QFileDevice_Unmap(QFileDevice* self, unsigned char* address);
QDateTime* QFileDevice_FileTime(const QFileDevice* self, int time);
bool QFileDevice_SetFileTime(QFileDevice* self, QDateTime* newDate, int fileTime);
long long QFileDevice_ReadData(QFileDevice* self, char* data, long long maxlen);
long long QFileDevice_WriteData(QFileDevice* self, const char* data, long long lenVal);
long long QFileDevice_ReadLineData(QFileDevice* self, char* data, long long maxlen);
struct miqt_string QFileDevice_Tr2(const char* s, const char* c);
struct miqt_string QFileDevice_Tr3(const char* s, const char* c, int n);
unsigned char* QFileDevice_Map3(QFileDevice* self, long long offset, long long size, int flags);
void QFileDevice_Delete(QFileDevice* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
