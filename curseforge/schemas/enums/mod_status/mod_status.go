package mod_status

/*
ModStatus

Possible enum values:

1=New

2=ChangesRequired

3=UnderSoftReview

4=Approved

5=Rejected

6=ChangesMade

7=Inactive

8=Abandoned

9=Deleted

10=UnderReview
*/

type ModStatus int

// ModStatus values as const
const (
	New             ModStatus = 1
	ChangesRequired ModStatus = 2
	UnderSoftReview ModStatus = 3
	Approved        ModStatus = 4
	Rejected        ModStatus = 5
	ChangesMade     ModStatus = 6
	Inactive        ModStatus = 7
	Abandoned       ModStatus = 8
	Deleted         ModStatus = 9
	UnderReview     ModStatus = 10
)
