package hash_algo

type HashAlgo int

/*
Possible enum values:

1=Sha1

2=Md5
*/

const (
	Sha1 HashAlgo = 1
	Md5  HashAlgo = 2
)
