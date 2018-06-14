# The set of languages for which implicit dependencies are needed:
SET(CMAKE_DEPENDS_LANGUAGES
  )
# The set of files for implicit dependencies of each language:

# Targets to which this target links.
SET(CMAKE_TARGET_LINKED_INFO_FILES
  )

# The include file search paths:
SET(CMAKE_C_TARGET_INCLUDE_PATH
  "../modules/videoio/include"
  "../modules/imgcodecs/include"
  "../modules/imgproc/include"
  "../modules/core/include"
  "../modules/highgui/include"
  "3rdparty/ippicv/ippicv_lnx/include"
  "3rdparty/ippicv/ippiw_lnx/include"
  "."
  "/usr/include/gdal"
  "/usr/include/eigen3"
  )
SET(CMAKE_CXX_TARGET_INCLUDE_PATH ${CMAKE_C_TARGET_INCLUDE_PATH})
SET(CMAKE_Fortran_TARGET_INCLUDE_PATH ${CMAKE_C_TARGET_INCLUDE_PATH})
SET(CMAKE_ASM_TARGET_INCLUDE_PATH ${CMAKE_C_TARGET_INCLUDE_PATH})
