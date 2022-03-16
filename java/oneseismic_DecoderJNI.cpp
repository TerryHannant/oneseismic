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

JNIEXPORT void JNICALL Java_oneseismic_DecoderJNI_buffer_1and_1process
(JNIEnv * env, jobject, jlong ptr, jbyteArray array, jint len) {
  auto dec = (one::decoder*) ptr;
  //int len = env->GetArrayLength (array);
  char* buf = new  char[len];
  env->GetByteArrayRegion (array, 0, len, reinterpret_cast<jbyte*>(buf)); 
  dec->buffer_and_process(buf,len);
    
  
}

JNIEXPORT void JNICALL Java_oneseismic_DecoderJNI_header
(JNIEnv *, jobject, jlong ptr) {
  auto dec = (one::decoder*) ptr;
  auto header = dec->header();
}

JNIEXPORT jobjectArray JNICALL Java_oneseismic_DecoderJNI_header_1attributes
(JNIEnv *env, jobject, jlong ptr) {
  auto dec = (one::decoder*) ptr;
  auto header = dec->header();
  if (header == nullptr) {
    return nullptr;
  }
  auto attrs = header->attributes;
  auto attrs_size = attrs.size();
  auto result = env->NewObjectArray(attrs_size, env->FindClass("java/lang/String"), nullptr);
  
  for (int i = 0; i < attrs_size; ++i)
    {
      env->SetObjectArrayElement(result, i, env->NewStringUTF(attrs[i].c_str()));
    }
  
  return result;
}

JNIEXPORT jintArray JNICALL Java_oneseismic_DecoderJNI_header_1shape
(JNIEnv *env, jobject, jlong ptr) {
  auto dec = (one::decoder*) ptr;
  auto header = dec->header();
  if (header == nullptr) {
    return nullptr;
  }
  auto shape = header->shapes;
  auto shape_size = shape.size();
  auto result = env->NewIntArray(shape_size);

  jint fill[shape_size];
  for (int i = 0; i < shape_size; i++) {
    fill[i] = shape[i];
  }
 // move from the temp structure to the java structure
 env->SetIntArrayRegion(result, 0, shape_size, fill);
 return result;

}

JNIEXPORT void JNICALL Java_oneseismic_DecoderJNI_register_1writer
(JNIEnv *env, jobject, jlong ptr, jstring name, jobject byte_buffer) {
  auto dec = (one::decoder*) ptr;
  char *buf = (char*)env->GetDirectBufferAddress(byte_buffer);
  

  const char *str = env->GetStringUTFChars(name, 0);
  std::string string = std::string(str);
  env->ReleaseStringUTFChars(name, str);
    
  dec->register_writer(string,buf);
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
    //std::cout << "status: paused" << std::endl;
  }
}



JNIEXPORT void JNICALL Java_oneseismic_DecoderJNI_free_1decoder
(JNIEnv *, jobject, jlong ptr) {
  //  delete (one::decoder) ptr;
}
