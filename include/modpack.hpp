#pragma once
#include <string>
#include <nlohmann/json.hpp>
#include <vector>
#include <CurseMod.hpp>

namespace mcmpdl {
    using std::string;
    using json = nlohmann::json;
    class ModPackExtractor {
        public:
            ModPackExtractor(string);
            virtual void init();
            virtual void extract(string, bool);

    };

    class CurseForgeModpack : public ModPackExtractor {
        public:
            CurseForgeModpack(string);
            void init() override;
            void extract(string, bool) override;
        private:
            string path_or_url;
            json manifest;
            std::vector<cf::CurseMod> mods;            
    };
}
