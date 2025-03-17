package mod_loader_type

/*
Possible enum values:

0=Any

1=Forge

2=Cauldron

3=LiteLoader

4=Fabric

5=Quilt

6=NeoForge
*/

type ModLoaderType int

const (
	Any        ModLoaderType = 0
	Forge      ModLoaderType = 1
	Cauldron   ModLoaderType = 2
	LiteLoader ModLoaderType = 3
	Fabric     ModLoaderType = 4
	Quilt      ModLoaderType = 5
	NeoForge   ModLoaderType = 6
)
