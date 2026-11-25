//go:build !amd64 && !arm64
// +build !amd64,!arm64

package gomonkey

func buildCorePatchDirective(_, double, _ uintptr) ([]byte, uintptr) {
	return buildJmpDirective(double), 0
}

func releaseCorePatchTrampoline(trampoline uintptr) {
	releasePrivateMethodTrampoline(trampoline)
}
