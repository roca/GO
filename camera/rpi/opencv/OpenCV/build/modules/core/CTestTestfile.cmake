# CMake generated Testfile for 
# Source directory: /go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/modules/core
# Build directory: /go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/modules/core
# 
# This file includes the relevant testing commands required for 
# testing this directory and lists subdirectories to be tested as well.
ADD_TEST(opencv_test_core "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/bin/opencv_test_core" "--gtest_output=xml:opencv_test_core.xml")
SET_TESTS_PROPERTIES(opencv_test_core PROPERTIES  LABELS "Main;opencv_core;Accuracy" WORKING_DIRECTORY "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/test-reports/accuracy")
ADD_TEST(opencv_perf_core "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/bin/opencv_perf_core" "--gtest_output=xml:opencv_perf_core.xml")
SET_TESTS_PROPERTIES(opencv_perf_core PROPERTIES  LABELS "Main;opencv_core;Performance" WORKING_DIRECTORY "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/test-reports/performance")
ADD_TEST(opencv_sanity_core "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/bin/opencv_perf_core" "--gtest_output=xml:opencv_perf_core.xml" "--perf_min_samples=1" "--perf_force_samples=1" "--perf_verify_sanity")
SET_TESTS_PROPERTIES(opencv_sanity_core PROPERTIES  LABELS "Main;opencv_core;Sanity" WORKING_DIRECTORY "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/test-reports/sanity")
