#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QTCPSOCKET_H
#define MIQT_QT_NETWORK_GEN_QTCPSOCKET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractSocket;
class QIODevice;
class QMetaObject;
class QObject;
class QTcpSocket;
class QVariant;
#else
typedef struct QAbstractSocket QAbstractSocket;
typedef struct QIODevice QIODevice;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTcpSocket QTcpSocket;
typedef struct QVariant QVariant;
#endif

QTcpSocket* QTcpSocket_new();
QTcpSocket* QTcpSocket_new2(QObject* parent);
void QTcpSocket_virtbase(QTcpSocket* src, QAbstractSocket** outptr_QAbstractSocket);
QMetaObject* QTcpSocket_MetaObject(const QTcpSocket* self);
void* QTcpSocket_Metacast(QTcpSocket* self, const char* param1);
struct miqt_string QTcpSocket_Tr(const char* s);
struct miqt_string QTcpSocket_TrUtf8(const char* s);
struct miqt_string QTcpSocket_Tr2(const char* s, const char* c);
struct miqt_string QTcpSocket_Tr3(const char* s, const char* c, int n);
struct miqt_string QTcpSocket_TrUtf82(const char* s, const char* c);
struct miqt_string QTcpSocket_TrUtf83(const char* s, const char* c, int n);
void QTcpSocket_override_virtual_Resume(void* self, intptr_t slot);
void QTcpSocket_virtualbase_Resume(void* self);
void QTcpSocket_override_virtual_ConnectToHost(void* self, intptr_t slot);
void QTcpSocket_virtualbase_ConnectToHost(void* self, struct miqt_string hostName, uint16_t port, int mode, int protocol);
void QTcpSocket_override_virtual_DisconnectFromHost(void* self, intptr_t slot);
void QTcpSocket_virtualbase_DisconnectFromHost(void* self);
void QTcpSocket_override_virtual_BytesAvailable(void* self, intptr_t slot);
long long QTcpSocket_virtualbase_BytesAvailable(const void* self);
void QTcpSocket_override_virtual_BytesToWrite(void* self, intptr_t slot);
long long QTcpSocket_virtualbase_BytesToWrite(const void* self);
void QTcpSocket_override_virtual_CanReadLine(void* self, intptr_t slot);
bool QTcpSocket_virtualbase_CanReadLine(const void* self);
void QTcpSocket_override_virtual_SetReadBufferSize(void* self, intptr_t slot);
void QTcpSocket_virtualbase_SetReadBufferSize(void* self, long long size);
void QTcpSocket_override_virtual_SocketDescriptor(void* self, intptr_t slot);
intptr_t QTcpSocket_virtualbase_SocketDescriptor(const void* self);
void QTcpSocket_override_virtual_SetSocketDescriptor(void* self, intptr_t slot);
bool QTcpSocket_virtualbase_SetSocketDescriptor(void* self, intptr_t socketDescriptor, int state, int openMode);
void QTcpSocket_override_virtual_SetSocketOption(void* self, intptr_t slot);
void QTcpSocket_virtualbase_SetSocketOption(void* self, int option, QVariant* value);
void QTcpSocket_override_virtual_SocketOption(void* self, intptr_t slot);
QVariant* QTcpSocket_virtualbase_SocketOption(void* self, int option);
void QTcpSocket_override_virtual_Close(void* self, intptr_t slot);
void QTcpSocket_virtualbase_Close(void* self);
void QTcpSocket_override_virtual_IsSequential(void* self, intptr_t slot);
bool QTcpSocket_virtualbase_IsSequential(const void* self);
void QTcpSocket_override_virtual_AtEnd(void* self, intptr_t slot);
bool QTcpSocket_virtualbase_AtEnd(const void* self);
void QTcpSocket_override_virtual_WaitForConnected(void* self, intptr_t slot);
bool QTcpSocket_virtualbase_WaitForConnected(void* self, int msecs);
void QTcpSocket_override_virtual_WaitForReadyRead(void* self, intptr_t slot);
bool QTcpSocket_virtualbase_WaitForReadyRead(void* self, int msecs);
void QTcpSocket_override_virtual_WaitForBytesWritten(void* self, intptr_t slot);
bool QTcpSocket_virtualbase_WaitForBytesWritten(void* self, int msecs);
void QTcpSocket_override_virtual_WaitForDisconnected(void* self, intptr_t slot);
bool QTcpSocket_virtualbase_WaitForDisconnected(void* self, int msecs);
void QTcpSocket_override_virtual_ReadData(void* self, intptr_t slot);
long long QTcpSocket_virtualbase_ReadData(void* self, char* data, long long maxlen);
void QTcpSocket_override_virtual_ReadLineData(void* self, intptr_t slot);
long long QTcpSocket_virtualbase_ReadLineData(void* self, char* data, long long maxlen);
void QTcpSocket_override_virtual_WriteData(void* self, intptr_t slot);
long long QTcpSocket_virtualbase_WriteData(void* self, const char* data, long long lenVal);
void QTcpSocket_Delete(QTcpSocket* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
