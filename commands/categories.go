package commands

type CategoryType string

const (
	AdminCategory CategoryType = "admin"
	AudioCategory CategoryType = "audio"
	FunCategory   CategoryType = "fun"
	InfoCategory  CategoryType = "info"
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
	categories = append(categories, &Category{
		Name:     string(FunCategory),
		Type:     FunCategory,
		Commands: []*Command{},
	})

	return
}
