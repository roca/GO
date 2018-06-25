FILE(REMOVE_RECURSE
  "CMakeFiles/gen-pkgconfig"
  "unix-install/opencv.pc"
)

# Per-language clean rules from dependency scanning.
FOREACH(lang)
  INCLUDE(CMakeFiles/gen-pkgconfig.dir/cmake_clean_${lang}.cmake OPTIONAL)
ENDFOREACH(lang)
