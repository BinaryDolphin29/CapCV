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
func NewMat() *Mat {
	return &Mat{C.NewMat()}
}

// Release Close the Mat.
func (m *Mat) Release() {
	C.MatRelease(m.mat)
}

// IsEmpty Whether the Mat is empty.
func (m *Mat) IsEmpty() bool {
	ptr := C.isEmpty(m.mat)
	return *(*bool)(unsafe.Pointer(&ptr))
}

func (m *Mat) Type() Type {
	ptr := C.type_(m.mat)
	return *(*Type)(unsafe.Pointer(&ptr))
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
func CvtColor(dst, src *Mat, code Type, channel int) {
	C.cvtColor(src.mat, dst.mat, C.int(code), C.int(channel))
}

// OpenCaptureDevice Open the Capture Device. You will need to check whether the camera is opened.
func OpenCaptureDevice(index int) *VideoCapture {
	return &VideoCapture{C.OpenCaptureDev(C.int(index))}
}

// Release Close the CaptureDevice.
func (dev *VideoCapture) Release() {
	C.DevRelease(dev.vcd)
}

// IsOpenedCaptureDevice Whether the is opened Capture Device.
func (dev *VideoCapture) IsOpenedCaptureDevice() bool {
	ptr := C.isOpenedCapDev(dev.vcd)
	return *(*bool)(unsafe.Pointer(&ptr))
}

// Read Read frame to the Mat.
func (dev *VideoCapture) Read(m *Mat) bool {
	ptr := C.Read(dev.vcd, m.mat)
	return *(*bool)(unsafe.Pointer(&ptr))
}

// Set Set property to the Capture Device. You will need warning to incompatible device properties.
func (dev *VideoCapture) Set(p VidCapProp, v float64) bool {
	ptr := C.DevSet(dev.vcd, C.int(p), C.double(v))
	return *(*bool)(unsafe.Pointer(&ptr))
}

// Get Get the Capture Device's value.
func (dev *VideoCapture) Get(p VidCapProp) float64 {
	return float64(C.DevGet(dev.vcd, C.int(p)))
}
