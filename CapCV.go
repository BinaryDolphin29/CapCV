package CapCV

// If you use windows, You need to changing ↓this path(CPPFLAGS, LDFLAGS).

// #cgo !linux CPPFLAGS: -IC:/opencv/build/install/include
// #cgo !linux LDFLAGS: -LC:/opencv/build/install/x64/vc16/lib -lopencv_core -lopencv_videoio -lopencv_imgproc
// #cgo CPPFLAGS: -I/usr/local/include/opencv4/
// #cgo LDFLAGS: -L/usr/local/lib/ -lopencv_core -lopencv_videoio -lopencv_imgproc
// #include "cv.hpp"
import "C"
import (
	"unsafe"
)

// Mat Minimum necessary implementation for cv::Mat.
type Mat struct {
	mat C.Mat
}

// VideoCapture Minimum necessary implementation for cv::VideoCapture.
type VideoCapture struct {
	vcd C.Dev
}

// NewMat Create cv::Mat.
func (m *Mat) NewMat() *Mat {
	m.mat = C.NewMat()
	return m
}

// IsEmpty Whether the Mat is empty.
func (m *Mat) IsEmpty() bool {
	return *(*bool)(unsafe.Pointer(uintptr(C.isEmpty(m.mat))))
}

// GetPix Get one Pixel values from Mat.
func (m *Mat) GetPix(x, y int) [3]uint8 {
	return *(*[3]uint8)(unsafe.Pointer(C.GetPix(m.mat, C.int(x), C.int(y))))
}

// Ptr Return is Ptr(uintptr).
func (m *Mat) Ptr() uintptr {
	return uintptr(C.Ptr(m.mat))
}

// RawPtr Retrun is cv::Vec3b type.
func (m *Mat) RawPtr() C.V3b {
	return C.Ptr(m.mat)
}

// CvtColor Convert the Mat color.
func (m *Mat) CvtColor(dst *Mat, code, channel int) {
	C.cvtColor(m.mat, dst.mat, C.int(code), C.int(channel))
}

// OpenCaptureDevice Open the Capture Device. You will need to check whether the camera is opened.
func (dev *VideoCapture) OpenCaptureDevice(index int) *VideoCapture {
	dev.vcd = C.OpenCaptureDev(C.int(index))
	return dev
}

// IsOpenedCaptureDevice Whether the is opened Capture Device.
func (dev *VideoCapture) IsOpenedCaptureDevice() bool {
	return *(*bool)(unsafe.Pointer(uintptr(C.isOpenedCapDev(dev.vcd))))
}

// Read Read frame to the Mat.
func (dev *VideoCapture) Read(m *Mat) bool {
	return *(*bool)(unsafe.Pointer(uintptr(C.Read(dev.vcd, m.mat))))
}

// Set Set property to the Capture Device. You will need warning to incompatible device properties.
func (dev *VideoCapture) Set(p VidCapProp, v float64) bool {
	return *(*bool)(unsafe.Pointer(uintptr(C.DevSet(dev.vcd, C.int(p), C.double(v)))))
}

// Get Get the Capture Device's value.
func (dev *VideoCapture) Get(p VidCapProp) float64 {
	return float64(C.DevGet(dev.vcd, C.int(p)))
}
