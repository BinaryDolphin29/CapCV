package CapCV

// Type https://docs.opencv.org/4.5.2/d1/d1b/group__core__hal__interface.html#ga32b18d904ee2b1731a9416a8eef67d06
type Type uint8

const (
	CV_8U Type = iota
	CV_8S
	CV_16U
	CV_16S
	CV_32S
	CV_32F
	CV_64F
	CV_16F
)

const (
	CV_8UC1 Type = iota * 8
	CV_8UC2
	CV_8UC3
	CV_8UC4
)

const (
	CV_8SC1 Type = (iota * 8) + 1
	CV_8SC2
	CV_8SC3
	CV_8SC4
)

const (
	CV_16UC1 Type = (iota * 8) + 2
	CV_16UC2
	CV_16UC3
	CV_16UC4
)

const (
	CV_16SC1 Type = (iota * 8) + 3
	CV_16SC2
	CV_16SC3
	CV_16SC4
)

const (
	CV_32SC1 Type = (iota * 8) + 4
	CV_32SC2
	CV_32SC3
	CV_32SC4
)

const (
	CV_32FC1 Type = (iota * 8) + 5
	CV_32FC2
	CV_32FC3
	CV_32FC4
)

const (
	CV_64FC1 Type = (iota * 8) + 6
	CV_64FC2
	CV_64FC3
	CV_64FC4
)

const (
	CV_16FC1 Type = (iota * 8) + 7
	CV_16FC2
	CV_16FC3
	CV_16FC4
)
