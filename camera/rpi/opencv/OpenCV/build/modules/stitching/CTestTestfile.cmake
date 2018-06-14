# CMake generated Testfile for 
# Source directory: /go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/modules/stitching
# Build directory: /go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/modules/stitching
# 
# This file includes the relevant testing commands required for 
# testing this directory and lists subdirectories to be tested as well.
ADD_TEST(opencv_test_stitching "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/bin/opencv_test_stitching" "--gtest_output=xml:opencv_test_stitching.xml")
SET_TESTS_PROPERTIES(opencv_test_stitching PROPERTIES  LABELS "Main;opencv_stitching;Accuracy" WORKING_DIRECTORY "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/test-reports/accuracy")
ADD_TEST(opencv_perf_stitching "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/bin/opencv_perf_stitching" "--gtest_output=xml:opencv_perf_stitching.xml")
SET_TESTS_PROPERTIES(opencv_perf_stitching PROPERTIES  LABELS "Main;opencv_stitching;Performance" WORKING_DIRECTORY "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/test-reports/performance")
ADD_TEST(opencv_sanity_stitching "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/bin/opencv_perf_stitching" "--gtest_output=xml:opencv_perf_stitching.xml" "--perf_min_samples=1" "--perf_force_samples=1" "--perf_verify_sanity")
SET_TESTS_PROPERTIES(opencv_sanity_stitching PROPERTIES  LABELS "Main;opencv_stitching;Sanity" WORKING_DIRECTORY "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/test-reports/sanity")
