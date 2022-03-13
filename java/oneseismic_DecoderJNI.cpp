#include "oneseismic_DecoderJNI.h"
#include <oneseismic/decoder.hpp>
//#include "/home/thannant/development/equinor/oneseismic/core/include/oneseismic/decoder.hpp"
//#include <oneseismic/decoder.hpp>
#include <iostream>

JNIEXPORT void JNICALL Java_oneseismic_DecoderJNI_decode
(JNIEnv *, jobject) {
  std::cout << "Hello from C++ !!" << std::endl;
}


JNIEXPORT jlong JNICALL Java_oneseismic_DecoderJNI_create_1decoder
(JNIEnv *, jobject) {
  one::decoder* instance_ptr = new one::decoder();
  return (jlong) instance_ptr;
}

JNIEXPORT void JNICALL Java_oneseismic_DecoderJNI_status
(JNIEnv *, jobject, jlong ptr) {
  auto dec = (one::decoder*) ptr;
  auto status = dec->process();
  if (status == one::decoder::status::done)
    {
      std::cout << "status done"<< std::endl;
    }
  else {
    std::cout << "status: paused" << std::endl;
  }
}



JNIEXPORT void JNICALL Java_oneseismic_DecoderJNI_free_1decoder
(JNIEnv *, jobject, jlong ptr) {
  //  delete (one::decoder) ptr;
}
