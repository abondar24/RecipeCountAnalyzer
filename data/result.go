package data

type CountResult struct {
	UniqueCount     int32        `json:"unique_recipe_count"`
	RecipeCount     []Recipe     `json:"count_per_recipe"`
	BusyPostcode    PostcodeBusy `json:"busiest_postcode"`
	PostcodePerTime PostcodeTime `json:"count_per_postcode_and_time"`
	Matches         []string     `json:"match_by_name"`
}
