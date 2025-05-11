package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework CoreGraphics -framework CoreFoundation -framework IOKit

#include <CoreGraphics/CoreGraphics.h>
#include <CoreGraphics/CoreGraphics.h>
#include <CoreFoundation/CoreFoundation.h>

void testCG() {
    CGPoint point = CGPointMake(10, 20);
    printf("Point: %.1f, %.1f\n", point.x, point.y);
}

CGEventMask EventMaskBit(CGEventType eventType) {
    return (CGEventMask)(1 << eventType);
}
CGEventRef eventCallback(CGEventTapProxy proxy, CGEventType type, CGEventRef event, void *userInfo) {
    // Just pass through the event for now
    return event;
}
*/
import "C"

func Init() {
	println("init running")
	C.testCG()
	// C.eventCallback()

	eventMask := C.EventMaskBit(C.kCGEventKeyDown) | C.EventMaskBit(C.kCGEventKeyUp)
	print(eventMask)
	// C.Cr
}
