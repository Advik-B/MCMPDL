#include <iostream>
#include <string>
#include <sstream>

std::string encodeKey(const std::string& key) {
    std::string encodedKey;
    for (char c : key) {
        encodedKey += static_cast<char>(c ^ 0x55); // XOR with 0x55
    }
    return encodedKey;
}

std::string toHexString(const std::string& str) {
    std::stringstream ss;
    for (unsigned char c : str) {
        ss << "\\x" << std::hex << static_cast<int>(c);
    }
    return ss.str();
}

std::string decodeKey(const std::string& encodedKey) {
    std::string key;
    for (char c : encodedKey) {
        key += c ^ 0x55; // Simple XOR with 0x55
    }
    return key;
}

int main() {
    std::string apiKey = "my_secret_key"; // Your original API key
    getline(std::cin, apiKey);
    std::string encodedKey = encodeKey(apiKey);
    std::string hexEncodedKey = toHexString(encodedKey);

    std::cout << "Encoded API Key: " << hexEncodedKey << std::endl;
    std::cout << "Decoded API Key: " << decodeKey(encodedKey) << std::endl;
    return 0;
}
