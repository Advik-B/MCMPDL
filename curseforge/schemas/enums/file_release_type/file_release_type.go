package file_release_type

/*
FileReleaseType

Possible enum values:

1=Release

2=Beta

3=Alpha
*/

type FileReleaseType int

// FileReleaseType values as const
const (
	Release FileReleaseType = 1
	Beta    FileReleaseType = 2
	Alpha   FileReleaseType = 3
)
