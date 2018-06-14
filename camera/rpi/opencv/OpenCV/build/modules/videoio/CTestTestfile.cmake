# CMake generated Testfile for 
# Source directory: /go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/modules/videoio
# Build directory: /go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/modules/videoio
# 
# This file includes the relevant testing commands required for 
# testing this directory and lists subdirectories to be tested as well.
ADD_TEST(opencv_test_videoio "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/bin/opencv_test_videoio" "--gtest_output=xml:opencv_test_videoio.xml")
SET_TESTS_PROPERTIES(opencv_test_videoio PROPERTIES  LABELS "Main;opencv_videoio;Accuracy" WORKING_DIRECTORY "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/test-reports/accuracy")
ADD_TEST(opencv_perf_videoio "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/bin/opencv_perf_videoio" "--gtest_output=xml:opencv_perf_videoio.xml")
SET_TESTS_PROPERTIES(opencv_perf_videoio PROPERTIES  LABELS "Main;opencv_videoio;Performance" WORKING_DIRECTORY "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/test-reports/performance")
ADD_TEST(opencv_sanity_videoio "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/bin/opencv_perf_videoio" "--gtest_output=xml:opencv_perf_videoio.xml" "--perf_min_samples=1" "--perf_force_samples=1" "--perf_verify_sanity")
SET_TESTS_PROPERTIES(opencv_sanity_videoio PROPERTIES  LABELS "Main;opencv_videoio;Sanity" WORKING_DIRECTORY "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/test-reports/sanity")
