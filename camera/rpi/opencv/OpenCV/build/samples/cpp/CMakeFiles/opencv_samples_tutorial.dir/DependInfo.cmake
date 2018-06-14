# The set of languages for which implicit dependencies are needed:
SET(CMAKE_DEPENDS_LANGUAGES
  )
# The set of files for implicit dependencies of each language:

# Preprocessor definitions for this target.
SET(CMAKE_TARGET_DEFINITIONS
  "vtkRenderingCore_AUTOINIT=3(vtkInteractionStyle,vtkRenderingFreeType,vtkRenderingOpenGL)"
  )

# Targets to which this target links.
SET(CMAKE_TARGET_LINKED_INFO_FILES
  )

# The include file search paths:
SET(CMAKE_C_TARGET_INCLUDE_PATH
  "../modules/core/include"
  "../modules/calib3d/include"
  "../modules/features2d/include"
  "../modules/highgui/include"
  "../modules/videoio/include"
  "../modules/imgcodecs/include"
  "../modules/video/include"
  "../modules/photo/include"
  "../modules/imgproc/include"
  "../modules/flann/include"
  "../modules/viz/include"
  "../modules/videostab/include"
  "../modules/superres/include"
  "../modules/stitching/include"
  "../modules/shape/include"
  "../modules/objdetect/include"
  "../modules/ml/include"
  "../modules/dnn/include"
  "3rdparty/ippicv/ippicv_lnx/include"
  "3rdparty/ippicv/ippiw_lnx/include"
  "."
  "/usr/include/gdal"
  "/usr/include/eigen3"
  "/usr/include/vtk-6.0"
  )
SET(CMAKE_CXX_TARGET_INCLUDE_PATH ${CMAKE_C_TARGET_INCLUDE_PATH})
SET(CMAKE_Fortran_TARGET_INCLUDE_PATH ${CMAKE_C_TARGET_INCLUDE_PATH})
SET(CMAKE_ASM_TARGET_INCLUDE_PATH ${CMAKE_C_TARGET_INCLUDE_PATH})
