# CMake generated Testfile for 
# Source directory: /go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/modules/shape
# Build directory: /go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/modules/shape
# 
# This file includes the relevant testing commands required for 
# testing this directory and lists subdirectories to be tested as well.
ADD_TEST(opencv_test_shape "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/bin/opencv_test_shape" "--gtest_output=xml:opencv_test_shape.xml")
SET_TESTS_PROPERTIES(opencv_test_shape PROPERTIES  LABELS "Main;opencv_shape;Accuracy" WORKING_DIRECTORY "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/test-reports/accuracy")
