# CMake generated Testfile for 
# Source directory: /go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/modules/photo
# Build directory: /go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/modules/photo
# 
# This file includes the relevant testing commands required for 
# testing this directory and lists subdirectories to be tested as well.
ADD_TEST(opencv_test_photo "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/bin/opencv_test_photo" "--gtest_output=xml:opencv_test_photo.xml")
SET_TESTS_PROPERTIES(opencv_test_photo PROPERTIES  LABELS "Main;opencv_photo;Accuracy" WORKING_DIRECTORY "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/test-reports/accuracy")
ADD_TEST(opencv_perf_photo "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/bin/opencv_perf_photo" "--gtest_output=xml:opencv_perf_photo.xml")
SET_TESTS_PROPERTIES(opencv_perf_photo PROPERTIES  LABELS "Main;opencv_photo;Performance" WORKING_DIRECTORY "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/test-reports/performance")
ADD_TEST(opencv_sanity_photo "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/bin/opencv_perf_photo" "--gtest_output=xml:opencv_perf_photo.xml" "--perf_min_samples=1" "--perf_force_samples=1" "--perf_verify_sanity")
SET_TESTS_PROPERTIES(opencv_sanity_photo PROPERTIES  LABELS "Main;opencv_photo;Sanity" WORKING_DIRECTORY "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/test-reports/sanity")
