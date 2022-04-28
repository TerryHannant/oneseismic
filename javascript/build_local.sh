#!/usr/bin/env bash

set -xe

if [ "$#" -ne 1 ]; then
    echo "build dir missing"
fi

#git clone https://github.com/msgpack/msgpack-c.git
#cd msgpack-c
#git checkout cpp_master

BUILD_DIR=$(realpath $1)

echo $BUILD_DIR
shift

emcmake cmake \
    -S core/ -B $BUILD_DIR/core/ \
    -DBUILD_CORE=OFF \
    -DCMAKE_BUILD_TYPE=Release \
    -DCMAKE_CXX_FLAGS=-I/home/thannant/development/equinor/oneseismic/javascript/deps/boost_1_78_0/ \
    -DCMAKE_CXX_FLAGS=-I/home/thannant/development/equinor/oneseismic/javascript/deps/msgpack-c/include \
    -DCMAKE_TOOLCHAIN_FILE=/home/thannant/.asdf/installs/emsdk/3.1.8/upstream/emscripten/cmake/Modules/Platform/Emscripten.cmake \

pushd $BUILD_DIR/core
emmake make
popd


emcmake cmake \
    -S javascript/ -B $BUILD_DIR/javascript/ \
    -Doneseismic_DIR=$BUILD_DIR/core/ \
    -DCMAKE_BUILD_TYPE=Release \
    -DCMAKE_CXX_FLAGS=-I/home/thannant/development/equinor/oneseismic/javascript/deps/boost_1_78_0/ \
    -DCMAKE_CXX_FLAGS=-I/home/thannant/development/equinor/oneseismic/javascript/deps/msgpack-c/include  \
    -DCMAKE_INTERPROCEDURAL_OPTIMIZATION=1 \
    -DCMAKE_TOOLCHAIN_FILE=/home/thannant/.asdf/installs/emsdk/3.1.8/upstream/emscripten/cmake/Modules/Platform/Emscripten.cmake \
    "$@" \

pushd $BUILD_DIR/javascript
emmake make
popd
