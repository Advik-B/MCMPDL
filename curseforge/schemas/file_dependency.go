package schemas

import "MCMPDL/curseforge/schemas/enums/file_relation_type"

type FileDependency struct {
	ModId        int                                 `json:"modId"`
	RelationType file_relation_type.FileRelationType `json:"relationType"`
}
