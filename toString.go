package CapCV

func (t Type) String() string {
	switch t {
	case CV_8U:
		return "CV_8U,CV_8UC1"
	case CV_8S:
		return "CV_8S,CV_8SC1"
	case CV_16U:
		return "CV_16U,CV_16UC1"
	case CV_16S:
		return "CV_16S,CV_16SC1"
	case CV_32S:
		return "CV_32S,CV_32SC1"
	case CV_32F:
		return "CV_32F,CV_32FC1"
	case CV_64F:
		return "CV_64F,CV_64FC1"
	case CV_16F:
		return "CV_16F,CV_16FC1"
	case CV_8UC2:
		return "CV_8UC2"
	case CV_8UC3:
		return "CV_8UC3"
	case CV_8UC4:
		return "CV_8UC4"
	case CV_8SC2:
		return "CV_8SC2"
	case CV_8SC3:
		return "CV_8SC3"
	case CV_8SC4:
		return "CV_8SC4"
	case CV_16UC2:
		return "CV_16UC2"
	case CV_16UC3:
		return "CV_16UC3"
	case CV_16UC4:
		return "CV_16UC4"
	case CV_16SC2:
		return "CV_16SC2"
	case CV_16SC3:
		return "CV_16SC3"
	case CV_16SC4:
		return "CV_16SC4"
	case CV_32SC2:
		return "CV_32SC2"
	case CV_32SC3:
		return "CV_32SC3"
	case CV_32SC4:
		return "CV_32SC4"
	case CV_32FC2:
		return "CV_32FC2"
	case CV_32FC3:
		return "CV_32FC3"
	case CV_32FC4:
		return "CV_32FC4"
	case CV_64FC2:
		return "CV_64FC2"
	case CV_64FC3:
		return "CV_64FC3"
	case CV_64FC4:
		return "CV_64FC4"
	case CV_16FC2:
		return "CV_16FC2"
	case CV_16FC3:
		return "CV_16FC3"
	case CV_16FC4:
		return "CV_16FC4"

	default:
		return ""
	}
}
