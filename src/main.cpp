#include <iostream>
#include <string>
#include <sstream>
#include <elzip.hpp>
#include <filesystem>

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
    // elz::extractFile("test.zip", "test.txt");
    try {
        // Get the path to the temporary directory
        fs::path tempPath = fs::temp_directory_path();
        std::cout << tempPath / "mcmpdl";
        
        std::cout << "Temporary directory: " << tempPath << std::endl;
    } catch (const fs::filesystem_error& e) {
        std::cerr << "Error getting temp directory: " << e.what() << std::endl;
    }
    return 0;
}

