package CapCV

// #cgo !linux CPPFLAGS: -IC:/opencv/build/install/include
// #cgo !linux LDFLAGS: -LC:/opencv/build/install/x64/vc16/lib
// #cgo CPPFLAGS: -I/usr/local/include/opencv4/
// #cgo LDFLAGS: -L/usr/local/lib/ -lopencv_imgcodecs -lopencv_videoio -lopencv_core -lopencv_imgproc
// #include "cv.hpp"
import (
	"C"
)
import (
	"unsafe"
)

type Mat struct {
	mat C.Mat
}

type VideoCapture struct {
	vcd C.Dev
}

func (m *Mat) NewMat() *Mat {
	m.mat = C.NewMat()
	return m
}

func (m *Mat) IsEmpty() bool {
	return *(*bool)(unsafe.Pointer(uintptr(C.isEmpty(m.mat))))
}

func (m *Mat) GetPix(x, y int) [3]uint8 {
	return *(*[3]uint8)(unsafe.Pointer(C.GetPix(m.mat, C.int(x), C.int(y))))
}

func (m *Mat) Ptr() uintptr {
	return uintptr(C.Ptr(m.mat))
}

func (m *Mat) RawPtr() C.V3b {
	return C.Ptr(m.mat)
}

func (m *Mat) CvtColor(dst *Mat, code, channel int) {
	C.cvtColor(m.mat, dst.mat, C.int(code), C.int(channel))
}

func (dev *VideoCapture) OpenCaptureDevice(index int) *VideoCapture {
	dev.vcd = C.OpenCaptureDev(C.int(index))
	return dev
}

func (dev *VideoCapture) IsOpenedCaptureDevice() bool {
	return *(*bool)(unsafe.Pointer(uintptr(C.isOpenedCapDev(dev.vcd))))
}

func (dev *VideoCapture) Read(m *Mat) bool {
	return *(*bool)(unsafe.Pointer(uintptr(C.Read(dev.vcd, m.mat))))
}

func (dev *VideoCapture) Set(p VidCapProp, v float64) bool {
	return *(*bool)(unsafe.Pointer(uintptr(C.DevSet(dev.vcd, C.int(p), C.double(v)))))
}

func (dev *VideoCapture) Get(p VidCapProp) float64 {
	return float64(C.DevGet(dev.vcd, C.int(p)))
}
