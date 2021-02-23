%module bishop
#define __STDC__
#define __STDC_VERSION__ 199901L
#define __i386__
#define SAMPGDK_EMBEDDED
#define LINUX
#define __attribute__(x)

%{
#include "sampgdk.h"
%}

%import "plugincommon.h"
%import "amx/amx.h"
%include "sampgdk.h"
