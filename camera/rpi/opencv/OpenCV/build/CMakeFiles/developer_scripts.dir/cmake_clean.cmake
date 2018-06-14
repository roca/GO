FILE(REMOVE_RECURSE
  "CMakeFiles/developer_scripts"
)

# Per-language clean rules from dependency scanning.
FOREACH(lang)
  INCLUDE(CMakeFiles/developer_scripts.dir/cmake_clean_${lang}.cmake OPTIONAL)
ENDFOREACH(lang)
