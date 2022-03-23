#!/usr/bin/env bash

set -xe

if [ "$#" -ne 1 ]; then
    echo "build dir missing"
fi



BUILD_DIR=$(realpath $1)

echo $BUILD_DIR
shift

emcmake cmake \
    -S core/ -B $BUILD_DIR/core/ \
    -DBUILD_CORE=OFF \
    -DCMAKE_BUILD_TYPE=Release \
    -DCMAKE_CXX_FLAGS=-I/home/thannant/development/equinor/oneseismic/javascript/deps/msgpack-c/include \
    -DCMAKE_TOOLCHAIN_FILE=/usr/lib/emscripten/cmake/Modules/Platform/Emscripten.cmake \

pushd $BUILD_DIR/core
emmake make
popd


emcmake cmake \
    -S javascript/ -B $BUILD_DIR/javascript/ \
    -Doneseismic_DIR=$BUILD_DIR/core/ \
    -DCMAKE_BUILD_TYPE=Release \
    -DCMAKE_CXX_FLAGS=-I/home/thannant/development/equinor/oneseismic/javascript/deps/msgpack-c/include  \
    -DCMAKE_INTERPROCEDURAL_OPTIMIZATION=1 \
    -DCMAKE_TOOLCHAIN_FILE=/usr/lib/emscripten/cmake/Modules/Platform/Emscripten.cmake \
    "$@" \

pushd $BUILD_DIR/javascript
emmake make
popd
