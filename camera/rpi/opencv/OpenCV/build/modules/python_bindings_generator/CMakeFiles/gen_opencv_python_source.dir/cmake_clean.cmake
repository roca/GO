FILE(REMOVE_RECURSE
  "CMakeFiles/gen_opencv_python_source"
  "pyopencv_generated_include.h"
  "pyopencv_generated_funcs.h"
  "pyopencv_generated_types.h"
  "pyopencv_generated_type_reg.h"
  "pyopencv_generated_ns_reg.h"
  "pyopencv_signatures.json"
)

# Per-language clean rules from dependency scanning.
FOREACH(lang)
  INCLUDE(CMakeFiles/gen_opencv_python_source.dir/cmake_clean_${lang}.cmake OPTIONAL)
ENDFOREACH(lang)
