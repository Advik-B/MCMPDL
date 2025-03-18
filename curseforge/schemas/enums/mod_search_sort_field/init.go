package mod_search_sort_field

/*
ModsSearchSortField

Possible enum values:

1=Featured

2=Popularity

3=LastUpdated

4=Name

5=Author

6=TotalDownloads

7=Category

8=GameVersion

9=EarlyAccess

10=FeaturedReleased

11=ReleasedDate

12=Rating
*/

type ModsSearchSortField int

const (
	Featured         ModsSearchSortField = 1
	Popularity       ModsSearchSortField = 2
	LastUpdated      ModsSearchSortField = 3
	Name             ModsSearchSortField = 4
	Author           ModsSearchSortField = 5
	TotalDownloads   ModsSearchSortField = 6
	Category         ModsSearchSortField = 7
	GameVersion      ModsSearchSortField = 8
	EarlyAccess      ModsSearchSortField = 9
	FeaturedReleased ModsSearchSortField = 10
	ReleasedDate     ModsSearchSortField = 11
	Rating           ModsSearchSortField = 12
)

func (e ModsSearchSortField) Ptr() *ModsSearchSortField { return &e }
