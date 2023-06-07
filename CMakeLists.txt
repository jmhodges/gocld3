# This cmake scripts only builds a static cld3 lib and the unittests.

project(cld3)

# Old versions of cmake dont search/find protobuf lite
cmake_minimum_required(VERSION 3.9)

find_package(Protobuf REQUIRED)
message(STATUS "Protobuf_FOUND= ${Protobuf_FOUND}")
message(STATUS "Protobuf_VERSION= ${Protobuf_VERSION}")
message(STATUS "You need generate Protobuf files yourself.")
message(STATUS "Protobuf_LIBRARIES= ${Protobuf_LIBRARIES}")
message(STATUS "Protobuf_LITE_LIBRARIES= ${Protobuf_LITE_LIBRARIES}") # Usually /usr/lib64/libprotobuf-lite.so

message(STATUS "PROTO_HDRS= ${PROTO_HDRS}")

add_definitions(-fPIC) # Position Independant Code
add_definitions(-std=c++11) # Needed for std::to_string(), ...

include_directories(${CMAKE_CURRENT_BINARY_DIR} ${Protobuf_INCLUDE_DIRS}) # needed to include generated pb headers

add_library(${PROJECT_NAME}
        cld3/cld3.h
        cld3/cld3.cc

        cld3/feature_extractor.pb.h
        cld3/sentence.pb.h
        cld3/task_spec.pb.h
        cld3/feature_extractor.pb.cc
        cld3/sentence.pb.cc
        cld3/task_spec.pb.cc

        cld3/base.h
        cld3/casts.h
        cld3/embedding_feature_extractor.h
        cld3/embedding_network.h
        cld3/embedding_network_params.h
        cld3/feature_extractor.h
        cld3/feature_types.h
        cld3/fixunicodevalue.h
        cld3/float16.h
        cld3/fml_parser.h
        cld3/generated_ulscript.h
        cld3/getonescriptspan.h
        cld3/integral_types.h
        cld3/lang_id_nn_params.h
        cld3/language_identifier_features.h
        cld3/nnet_lang_id_test_data.h
        cld3/nnet_language_identifier.h
        cld3/offsetmap.h
        cld3/port.h
        cld3/registry.h
        cld3/relevant_script_feature.h
        cld3/script_detector.h
        cld3/sentence_features.h
        cld3/simple_adder.h
        cld3/stringpiece.h
        cld3/task_context.h
        cld3/task_context_params.h
        cld3/text_processing.h
        cld3/unicodetext.h
        cld3/utf8acceptinterchange.h
        cld3/utf8prop_lettermarkscriptnum.h
        cld3/utf8repl_lettermarklower.h
        cld3/utf8scannot_lettermarkspecial.h
        cld3/utf8statetable.h
        cld3/utils.h
        cld3/workspace.h

        cld3/base.cc
        cld3/embedding_feature_extractor.cc
        cld3/embedding_network.cc
        cld3/feature_extractor.cc
        cld3/feature_types.cc
        cld3/fixunicodevalue.cc
        cld3/fml_parser.cc
        cld3/generated_entities.cc
        cld3/generated_ulscript.cc
        cld3/getonescriptspan.cc
        cld3/lang_id_nn_params.cc
        cld3/language_identifier_features.cc
        cld3/nnet_lang_id_test_data.cc
        cld3/nnet_language_identifier.cc
        cld3/offsetmap.cc
        cld3/registry.cc
        cld3/relevant_script_feature.cc
        cld3/sentence_features.cc
        cld3/task_context.cc
        cld3/task_context_params.cc
        cld3/text_processing.cc
        cld3/unicodetext.cc
        cld3/utf8statetable.cc
        cld3/utils.cc
        cld3/workspace.cc
        )
