# Specify the minimum version of CMake required
cmake_minimum_required(VERSION 3.10)

# Define the project name and the programming languages used
project(FadakaProject VERSION 1.0 LANGUAGES CXX)

# Specify the C++ standard
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED True)

# Add include directories
include_directories(include)

# Add the executable target
add_executable(FadakaApp src/main.cpp src/other_source.cpp)

# Link libraries
target_link_libraries(FadakaApp PRIVATE some_library)

# Define installation rules
install(TARGETS FadakaApp DESTINATION bin)
install(DIRECTORY include/ DESTINATION include)
cmake_minimum_required(VERSION 3.1)
project(oracle-client)

set(CMAKE_CXX_STANDARD 11)

find_package(oatpp REQUIRED)
find_package(oatpp-curl REQUIRED)

include_directories(${OATPP_INCLUDE_DIRS})

add_executable(oracle-client main.cpp)
target_link_libraries(oracle-client
  oatpp::oatpp
  oatpp::oatpp-curl
)
