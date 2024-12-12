cmake_minimum_required(VERSION 3.25)

set(PROJECT_NAME "MCMPDL")
set(PROJECT_VERSION 1.0.0)

project(${PROJECT_NAME} VERSION ${PROJECT_VERSION} LANGUAGES CXX)

find_program(CCACHE_FOUND ccache)
if (CCACHE_FOUND)
    set_property(GLOBAL PROPERTY RULE_LAUNCH_COMPILE ccache)
    set_property(GLOBAL PROPERTY RULE_LAUNCH_LINK ccache)
endif()

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin)
set(CMAKE_POSITION_INDEPENDENT_CODE ON)
option(BUILD_SHARED_LIBS "Build shared libraries" OFF)

include(FetchContent)
FetchContent_Declare(
    cfapi
    GIT_REPOSITORY https://github.com/Advik-B/CurseForge-API-CPP.git
    GIT_TAG        v1.0.0
)
FetchContent_MakeAvailable(cfapi)

# Get environment variable "CURSEFORGE_API_KEY" and set it as a preprocessor definition
if (DEFINED ENV{CURSEFORGE_API_KEY})
    set(CURSEFORGE_API_KEY $ENV{CURSEFORGE_API_KEY})
    add_compile_definitions(CURSEFORGE_API_KEY="${CURSEFORGE_API_KEY}")
else()
    message(FATAL_ERROR "CURSEFORGE_API_KEY environment variable not set")
endif()

add_subdirectory(extern/11zip)

file(GLOB_RECURSE SOURCES "src/*.cpp")
file(GLOB_RECURSE HEADERS "include/*.hpp")

add_executable(${PROJECT_NAME} ${SOURCES} ${HEADERS})
target_include_directories(${PROJECT_NAME} PRIVATE "include")
target_link_libraries(${PROJECT_NAME} PRIVATE curseforge::static elzip)