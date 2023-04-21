package catalog

type catalog struct {
	ID   string `json:"id"`
	DATA string `json:"data"`
}

var catalogs = []catalog{
	{ID: "1", DATA: "DATA"},
	{ID: "3", DATA: "DATA 3"},
	{ID: "14", DATA: "DATA 14"},
}

func GetCatalog(catalod_id string) (catalog, error) {

	var foundCatalog catalog

	for key, cat_item := range catalogs {
		if catalogs[key].ID == catalod_id {
			foundCatalog = cat_item
		}
	}

	return foundCatalog, nil
}
