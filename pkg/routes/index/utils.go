package index

type sectionsStruct struct {
	Text   string
	Link   string
	Active bool
}

func getSections(active string) []sectionsStruct {
	return []sectionsStruct{
		{
			Text:   "Tracks",
			Link:   "/dashboard/tracks",
			Active: active == "tracks",
		},
		{
			Text:   "Folders",
			Link:   "/dashboard/folders",
			Active: active == "folders",
		},
		{
			Text:   "ID3 Tags",
			Link:   "/dashboard/id3Tags",
			Active: active == "id3Tags",
		},
		{
			Text:   "Settings",
			Link:   "/dashboard/settings",
			Active: active == "settings",
		},
	}
}
