// go-libtor - Self-contained Tor from Go
// Copyright (c) 2018 Péter Szilágyi. All rights reserved.
// +build darwin,amd64

package libtor

/*
#define BUILDDIR ""

#include <../src/ext/ed25519/ref10/ge_double_scalarmult.c>
*/
import "C"
