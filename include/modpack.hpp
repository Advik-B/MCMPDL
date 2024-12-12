#pragma once
#include <string>
#include <nlohmann/json.hpp>
#include <vector>
#include <CurseMod.hpp>
#include <spdlog/spdlog.h>

namespace mcmpdl {
    using std::string;
    using json = nlohmann::json;

    class CurseForgeModpack {
        public:
            CurseForgeModpack(json, std::shared_ptr<spdlog::logger>);
            void init();
            void extract(string);
        private:
            json manifest;
            std::vector<cf::CurseMod> mods;
            std::shared_ptr<spdlog::logger> logger;
    };
}
