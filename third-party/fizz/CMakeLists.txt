include(ExternalProject)
include(HPHPFunctions)

get_target_property(BOOST_INCLUDE_DIR boost INTERFACE_INCLUDE_DIRECTORIES)
get_target_property(BOOST_LIBRARIES boost INTERFACE_LINK_LIBRARIES)
list(GET BOOST_LIBRARIES 0 FIRST_LIB)
if("${FIRST_LIB}" MATCHES ".+/${CMAKE_STATIC_LIBRARY_PREFIX}boost_.+${CMAKE_STATIC_LIBRARY_SUFFIX}$")
  set(Boost_USE_STATIC_LIBS ON)
else()
  set(Boost_USE_STATIC_LIBS OFF)
endif()
get_filename_component(BOOST_LIBRARY_DIR "${FIRST_LIB}" DIRECTORY)

get_target_property(DOUBLE_CONVERSION_INCLUDE_DIR double-conversion INTERFACE_INCLUDE_DIRECTORIES)
get_target_property(DOUBLE_CONVERSION_LIBRARY double-conversion INTERFACE_LINK_LIBRARIES)
get_target_property(LIBSODIUM_INCLUDE_DIR libsodium INTERFACE_INCLUDE_DIRECTORIES)
get_target_property(LIBSODIUM_LIBRARY libsodium INTERFACE_LINK_LIBRARIES)
get_target_property(ZSTD_INCLUDE_DIR zstd INTERFACE_INCLUDE_DIRECTORIES)
get_target_property(ZSTD_LIBRARY zstd INTERFACE_LINK_LIBRARIES)
get_target_property(JEMALLOC_INCLUDE_DIR jemalloc INTERFACE_INCLUDE_DIRECTORIES)
get_target_property(GLOG_INCLUDE_DIR glog INTERFACE_INCLUDE_DIRECTORIES)
get_target_property(GLOG_LIBRARY glog INTERFACE_LINK_LIBRARIES)

set(
  FIZZ_DEPS
  boost
  folly
  jemalloc
  libsodium
  zstd
)
add_library(fizz_deps INTERFACE)
target_link_libraries(fizz_deps INTERFACE ${FIZZ_DEPS} ${OPENSSL_LIBRARIES})

ExternalProject_Add(
  bundled_fizz
  SOURCE_DIR ${CMAKE_CURRENT_SOURCE_DIR}/src/
  BUILD_ALWAYS ON
  PATCH_COMMAND
    "${CMAKE_COMMAND}" -E copy
    "${CMAKE_SOURCE_DIR}/CMake/FindLibSodium.cmake"
    <SOURCE_DIR>/fizz/cmake/FindLibSodium.cmake
    &&
    "${CMAKE_CURRENT_SOURCE_DIR}/../apply-quilt-patches.sh"
    "${CMAKE_CURRENT_SOURCE_DIR}/patches"
  SOURCE_SUBDIR fizz
  CMAKE_ARGS
    "-DCMAKE_PREFIX_PATH=${CMAKE_PREFIX_PATH}"
    -DCMAKE_MODULE_PATH=${CMAKE_SOURCE_DIR}/CMake
    -DCMAKE_BUILD_TYPE=${CMAKE_BUILD_TYPE}
    -DCMAKE_INSTALL_PREFIX=<INSTALL_DIR>
    -DCMAKE_INSTALL_INCLUDEDIR=include

    -DCMAKE_C_COMPILER=${CMAKE_C_COMPILER}
    -DCMAKE_CXX_COMPILER=${CMAKE_CXX_COMPILER}
    -DCMAKE_OSX_SYSROOT=${CMAKE_OSX_SYSROOT}

    -DBUILD_TESTS=OFF
    -DBUILD_EXAMPLES=OFF

    -DOPENSSL_INCLUDE_DIR=${OPENSSL_INCLUDE_DIR}
    -DOPENSSL_CRYPTO_LIBRARY=${OPENSSL_CRYPTO_LIBRARY}
    -DBoost_NO_SYSTEM_PATHS=ON
    -DBoost_NO_BOOST_CMAKE=ON
    -DBoost_USE_STATIC_LIBS=${Boost_USE_STATIC_LIBS}
    -DBoost_INCLUDE_DIR=${BOOST_INCLUDE_DIR}
    "-DBOOST_LIBRARYDIR=${BOOST_LIBRARY_DIR}"
    -DDOUBLE_CONVERSION_INCLUDE_DIR=${DOUBLE_CONVERSION_INCLUDE_DIR}
    -DDOUBLE_CONVERSION_LIBRARY=${DOUBLE_CONVERSION_LIBRARY}
    -DGLOG_INCLUDE_DIR=${GLOG_INCLUDE_DIR}
    -DGLOG_LIBRARY=${GLOG_LIBRARY}
    -DLIBSODIUM_INCLUDEDIR=${LIBSODIUM_INCLUDE_DIR}
    -DLIBSODIUM_LIBRARIES=${LIBSODIUM_LIBRARY}
    -DZSTD_INCLUDE_DIR=${ZSTD_INCLUDE_DIR}
    -DZSTD_LIBRARY_RELEASE=${ZSTD_LIBRARY}

    "-Dfmt_DIR=${fmt_DIR}"
    "-Dfolly_DIR=${FOLLY_INSTALL_DIR}/lib/cmake/folly"

    "-DCMAKE_CXX_FLAGS=-I${JEMALLOC_INCLUDE_DIR} ${CMAKE_CXX_FLAGS}"
    "-DCMAKE_C_FLAGS=${CMAKE_C_FLAGS}"
  DEPENDS ${FIZZ_DEPS}
)
ExternalProject_Get_property(bundled_fizz INSTALL_DIR)

add_library(fizz INTERFACE)
add_dependencies(fizz bundled_fizz)
target_include_directories(fizz INTERFACE "${INSTALL_DIR}/include")
target_link_libraries(
  fizz
  INTERFACE
  "${INSTALL_DIR}/lib/${CMAKE_STATIC_LIBRARY_PREFIX}fizz${CMAKE_STATIC_LIBRARY_SUFFIX}"
  fizz_deps
)

set(FIZZ_INSTALL_DIR "${INSTALL_DIR}" PARENT_SCOPE)
