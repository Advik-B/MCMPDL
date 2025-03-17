package schemas

import "MCMPDL/curseforge/schemas/enums/hash_algo"

type FileHash struct {
	Value string             `json:"value"`
	Algo  hash_algo.HashAlgo `json:"algo"`
}
