#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
    typedef cv::Mat *Mat;
    typedef cv::VideoCapture *Dev;
    typedef cv::Vec3b *V3b;
#else
    typedef void *Mat;
    typedef void *Dev;
    typedef void *V3b;
#endif

    Mat NewMat();
    int isEmpty(Mat);
    void cvtColor(Mat, Mat, int, int);
    V3b Ptr(Mat);
    V3b GetPix(Mat, int, int);
    Dev OpenCaptureDev(int);
    int isOpenedCapDev(Dev);
    int DevSet(Dev, int, double);
    double DevGet(Dev, int);
    int Read(Dev, Mat);

#ifdef __cplusplus
}
#endif