# CMake generated Testfile for 
# Source directory: /go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/modules/videostab
# Build directory: /go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/modules/videostab
# 
# This file includes the relevant testing commands required for 
# testing this directory and lists subdirectories to be tested as well.
ADD_TEST(opencv_test_videostab "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/bin/opencv_test_videostab" "--gtest_output=xml:opencv_test_videostab.xml")
SET_TESTS_PROPERTIES(opencv_test_videostab PROPERTIES  LABELS "Main;opencv_videostab;Accuracy" WORKING_DIRECTORY "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/test-reports/accuracy")
