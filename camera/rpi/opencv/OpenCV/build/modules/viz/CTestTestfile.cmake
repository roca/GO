# CMake generated Testfile for 
# Source directory: /go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/modules/viz
# Build directory: /go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/modules/viz
# 
# This file includes the relevant testing commands required for 
# testing this directory and lists subdirectories to be tested as well.
ADD_TEST(opencv_test_viz "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/bin/opencv_test_viz" "--gtest_output=xml:opencv_test_viz.xml")
SET_TESTS_PROPERTIES(opencv_test_viz PROPERTIES  LABELS "Main;opencv_viz;Accuracy" WORKING_DIRECTORY "/go/src/github.com/GOCODE/camera/rpi/opencv/OpenCV/build/test-reports/accuracy")
