#include <iostream>
#include <string>
#include <sstream>

std::string decodeKey(const std::string& encodedKey) {
    std::string key;
    for (char c : encodedKey) {
        key += c ^ 0x55; // Simple XOR with 0x55
    }
    return key;
}

int main() {
    std::cout <<  decodeKey(CURSEFORGE_API_KEY) << std::endl;   
    return 0;
}
