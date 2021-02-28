package commands

type CategoryType string

const (
	InfoCategory  CategoryType = "info"
	AdminCategory CategoryType = "admin"
	AudioCategory CategoryType = "audio"
)

func emptyCategories() (categories []*Category) {
	categories = append(categories, &Category{
		Name:     string(InfoCategory),
		Type:     InfoCategory,
		Commands: []*Command{},
	})
	categories = append(categories, &Category{
		Name:     string(AdminCategory),
		Type:     AdminCategory,
		Commands: []*Command{},
	})
	categories = append(categories, &Category{
		Name:     string(AudioCategory),
		Type:     AudioCategory,
		Commands: []*Command{},
	})

	return
}
