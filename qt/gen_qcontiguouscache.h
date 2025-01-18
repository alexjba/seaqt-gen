#pragma once
#ifndef MIQT_QT_GEN_QCONTIGUOUSCACHE_H
#define MIQT_QT_GEN_QCONTIGUOUSCACHE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QContiguousCacheData;
#else
typedef struct QContiguousCacheData QContiguousCacheData;
#endif

QContiguousCacheData* QContiguousCacheData_AllocateData(int size, int alignment);
void QContiguousCacheData_FreeData(QContiguousCacheData* data);
void QContiguousCacheData_Delete(QContiguousCacheData* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
