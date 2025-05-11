
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework CoreGraphics -framework CoreFoundation
#include <CoreGraphics/CoreGraphics.h>

// We need to declare the callback in C
CGEventRef eventCallbackC(CGEventTapProxy proxy, CGEventType type, CGEventRef event, void *userInfo) {
    return event;
}
