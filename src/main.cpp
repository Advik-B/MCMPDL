#include <iostream>
#include <string>
#include <sstream>
#include <elzip.hpp>
#include <filesystem>
#include <spdlog/spdlog.h>
#include <spdlog/sinks/stdout_color_sinks.h>
#include <modpack.hpp>

namespace fs = std::filesystem;

std::string decodeKey(const std::string& encodedKey) {
    std::string key;
    key.reserve(60);
    for (char c : encodedKey) {
        key += c ^ 0x55; // Simple XOR with 0x55
    }
    return key;
}

int main() {
    std::cout <<  decodeKey(CURSEFORGE_API_KEY) << "\n";
    auto logger = spdlog::stdout_color_mt("mcmpdl");
    mcmpdl::CurseForgeModpack modpack("asdas", logger);
    modpack.init();
    return 0;
}

