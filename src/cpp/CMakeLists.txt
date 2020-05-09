cmake_minimum_required(VERSION 3.13)
project(cpp)

set(CMAKE_CXX_STANDARD 17)

aux_source_directory(. DIR_SRCS)
aux_source_directory(list DIR_SRCS)
aux_source_directory(number DIR_SRCS)
add_executable(main ${DIR_SRCS})