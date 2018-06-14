# CMake generated Testfile for 
# Source directory: /go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/modules/features2d
# Build directory: /go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/modules/features2d
# 
# This file includes the relevant testing commands required for 
# testing this directory and lists subdirectories to be tested as well.
ADD_TEST(opencv_test_features2d "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/bin/opencv_test_features2d" "--gtest_output=xml:opencv_test_features2d.xml")
SET_TESTS_PROPERTIES(opencv_test_features2d PROPERTIES  LABELS "Main;opencv_features2d;Accuracy" WORKING_DIRECTORY "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/test-reports/accuracy")
ADD_TEST(opencv_perf_features2d "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/bin/opencv_perf_features2d" "--gtest_output=xml:opencv_perf_features2d.xml")
SET_TESTS_PROPERTIES(opencv_perf_features2d PROPERTIES  LABELS "Main;opencv_features2d;Performance" WORKING_DIRECTORY "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/test-reports/performance")
ADD_TEST(opencv_sanity_features2d "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/bin/opencv_perf_features2d" "--gtest_output=xml:opencv_perf_features2d.xml" "--perf_min_samples=1" "--perf_force_samples=1" "--perf_verify_sanity")
SET_TESTS_PROPERTIES(opencv_sanity_features2d PROPERTIES  LABELS "Main;opencv_features2d;Sanity" WORKING_DIRECTORY "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/test-reports/sanity")
