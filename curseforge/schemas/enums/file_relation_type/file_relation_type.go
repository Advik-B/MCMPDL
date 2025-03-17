package file_relation_type

/*
FileRelationType


Possible enum values:

1=EmbeddedLibrary

2=OptionalDependency

3=RequiredDependency

4=Tool

5=Incompatible

6=Include
*/

type FileRelationType int

const (
	EmbeddedLibrary    FileRelationType = 1
	OptionalDependency FileRelationType = 2
	RequiredDependency FileRelationType = 3
	Tool               FileRelationType = 4
	Incompatible       FileRelationType = 5
	Include            FileRelationType = 6
)
