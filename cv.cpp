#include "opencv2/core.hpp"
#include "opencv2/videoio.hpp"
#include "opencv2/imgproc.hpp"
#include "cv.hpp"

Mat NewMat() {
    return (new cv::Mat());
}

int isEmpty(const Mat m) {
    return m->empty();
}

int type_(Mat m) {
    return m->type();
}

void cvtColor(Mat src, Mat dst, int code, int chan) {
    cv::cvtColor(*src, *dst, code, chan);
}

V3b Ptr(Mat m) {
    // cvtColor(m, m, cv::COLOR_BGR2RGBA, 4);
    return m->ptr<cv::Vec3b>();
}

V3b GetPix(Mat m, int x, int y) {
    return &m->at<cv::Vec3b>(x, y);
}

Dev OpenCaptureDev(int idx) {
    return (new cv::VideoCapture(idx));
}

int isOpenedCapDev(Dev camera) {
    return camera->isOpened();
}

int DevSet(Dev d, int propID, double value) {
    // d->set(cv::CAP_PROP_FRAME_WIDTH, 96);
    // d->set(cv::CAP_PROP_FRAME_HEIGHT, 64);
    return d->set(propID, value);
}

double DevGet(Dev d, int propID) {
    return d->get(propID);
}

int Read(Dev c, Mat m) {
    return c->read(*m);
}

void MatRelease(Mat m) {
    m->release();
}

void DevRelease(Dev d) {
    d->release();
}