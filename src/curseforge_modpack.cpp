#include <string>
#include <modpack.hpp>

mcmpdl::CurseForgeModpack::CurseForgeModpack(std::string path_or_url) : ModPackExtractor(path_or_url) {
    this->path_or_url = path_or_url;
}

