package view

// BaseViewModel struct
type BaseViewModel struct {
	Title       string
	CurrentUser string
}

func (v *BaseViewModel) setTitle(title string) {
	v.Title = title
}

func (v *BaseViewModel) setCurrentUser(username string) {
	v.CurrentUser = username
}

// PageViewModel struct
type PageViewModel struct {
	PrevPage    int
	NextPage    int
	Total       int
	CurrentPage int
	Limit       int
}

func (v *PageViewModel) setPageViewModel(total, page, limit int) {
	v.Total = total
	v.CurrentPage = page
	v.Limit = limit

	v.setPrevAndNextPage()
}

func (v *PageViewModel) setPrevAndNextPage() {
	if v.CurrentPage > 1 {
		v.PrevPage = v.CurrentPage - 1
	}

	if (v.Total-1)/v.Limit >= v.CurrentPage {
		v.NextPage = v.CurrentPage + 1
	}
}
