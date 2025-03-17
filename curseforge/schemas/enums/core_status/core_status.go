package core_status

/*
CoreStatus
Possible enum values:

1=Draft

2=Test

3=PendingReview

4=Rejected

5=Approved

6=Live
*/

type CoreStatus int

const (
	Draft         CoreStatus = 1
	Test          CoreStatus = 2
	PendingReview CoreStatus = 3
	Rejected      CoreStatus = 4
	Approved      CoreStatus = 5
	Live          CoreStatus = 6
)
