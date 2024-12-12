#include <string>
#include <modpack.hpp>

mcmpdl::CurseForgeModpack::CurseForgeModpack(json manifest, std::shared_ptr<spdlog::logger> logger) {
    this->manifest = manifest;
    this->logger = logger->clone("curseforge");
}

void mcmpdl::CurseForgeModpack::init() {
    logger->info("Initializing CurseForge modpack");
    
}

void mcmpdl::CurseForgeModpack::extract(std::string output_dir) {
    logger->info("Extracting CurseForge modpack");
    // Extract the mods to the output_dir
    // If extract_mods is true, extract the mods
    // If extract_mods is false, just download the mods
}

