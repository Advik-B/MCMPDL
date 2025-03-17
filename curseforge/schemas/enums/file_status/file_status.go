package file_status

/*
FileStatus

1

Possible enum values:

1=Processing

2=ChangesRequired

3=UnderReview

4=Approved

5=Rejected

6=MalwareDetected

7=Deleted

8=Archived

9=Testing

10=Released

11=ReadyForReview

12=Deprecated

13=Baking

14=AwaitingPublishing

15=FailedPublishing

16=Cooking

17=Cooked

18=UnderManualReview

19=ScanningForMalware

20=ProcessingFile

21=PendingRelease

22=ReadyForCooking

23=PostProcessing
*/

type FileStatus int

// FileStatus values as const
const (
	Processing         FileStatus = 1
	ChangesRequired    FileStatus = 2
	UnderReview        FileStatus = 3
	Approved           FileStatus = 4
	Rejected           FileStatus = 5
	MalwareDetected    FileStatus = 6
	Deleted            FileStatus = 7
	Archived           FileStatus = 8
	Testing            FileStatus = 9
	Released           FileStatus = 10
	ReadyForReview     FileStatus = 11
	Deprecated         FileStatus = 12
	Baking             FileStatus = 13
	AwaitingPublishing FileStatus = 14
	FailedPublishing   FileStatus = 15
	Cooking            FileStatus = 16
	Cooked             FileStatus = 17
	UnderManualReview  FileStatus = 18
	ScanningForMalware FileStatus = 19
	ProcessingFile     FileStatus = 20
	PendingRelease     FileStatus = 21
	ReadyForCooking    FileStatus = 22
	PostProcessing     FileStatus = 23
)
