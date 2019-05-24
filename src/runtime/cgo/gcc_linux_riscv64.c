// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build cgo
// +build linux
// +build riscv64

#include <pthread.h>
#include <errno.h>
#include <string.h> // strerror
#include <signal.h>
#include <stdlib.h>
#include "libcgo.h"
#include "libcgo_unix.h"

static void* threadentry(void*);
static void (*setg_gcc)(void*);

// This will be set in gcc_android.c for android-specific customization.
void (*x_cgo_inittls)(void **tlsg, void **tlsbase);

void
x_cgo_init(G *g, void (*setg)(void*), void **tlsg, void **tlsbase)
{
	pthread_attr_t *attr;
	size_t size;

	setg_gcc = setg;
	attr = (pthread_attr_t*)malloc(sizeof *attr);
	if (attr == NULL) {
		fatalf("malloc failed: %s", strerror(errno));
	}
	pthread_attr_init(attr);
	pthread_attr_getstacksize(attr, &size);
	g->stacklo = (uintptr)&size - size + 4096;
	pthread_attr_destroy(attr);
	free(attr);

	if (x_cgo_inittls) {
		x_cgo_inittls(tlsg, tlsbase);
	}
}

void
_cgo_sys_thread_start(ThreadStart *ts)
{
	pthread_attr_t attr;
	sigset_t ign, oset;
	pthread_t p;
	size_t size;
	int err;

	sigfillset(&ign);
	pthread_sigmask(SIG_SETMASK, &ign, &oset);

	// Not sure why the memset is necessary here,
	// but without it, we get a bogus stack size
	// out of pthread_attr_getstacksize. C'est la Linux.
	memset(&attr, 0, sizeof attr);
	pthread_attr_init(&attr);
    size = 0;
	pthread_attr_getstacksize(&attr, &size);
	// Leave stacklo=0 and set stackhi=size; mstart will do the rest.
	ts->g->stackhi = size;
	err = _cgo_try_pthread_create(&p, &attr, threadentry, ts);

	pthread_sigmask(SIG_SETMASK, &oset, nil);

	if (err != 0) {
		fatalf("pthread_create failed: %s", strerror(err));
	}
}

extern void crosscall1(void (*fn)(void), void (*setg_gcc)(void*), void *g);
static void*
threadentry(void *v)
{
	ThreadStart ts;

	ts = *(ThreadStart*)v;
	free(v);

	crosscall1(ts.fn, setg_gcc, (void*)ts.g);
	return nil;
}