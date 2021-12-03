package author

import (
	"github.com/jakecoffman/crud"
)

// this must match the Query validation below
type authorQuery struct {
	Limit int64  `form:"limit" bson:"-"`
	Skip  int64  `form:"skip" bson:"-"`
	Sort  string `form:"sort" bson:"-"`
	Order int64  `form:"order" bson:"-"`

	Name *string `form:"name" bson:"name,omitempty"`
}

var tags = []string{"Authors"}

var Routes = []crud.Spec{
	{
		Method:  "GET",
		Path:    "/authors",
		Summary: "List authors",
		Tags:    tags,
		Handler: ListAuthors,
		Validate: crud.Validate{
			Query: crud.Object(map[string]crud.Field{
				"limit": crud.Integer().Min(0).Max(100).Default(100).Description("Amount of records to return (for pagination)"),
				"skip":  crud.Integer().Default(0).Description("Amount of records to skip (for pagination)"),
				"sort":  crud.String().Default("created").Description("Provide the field to sort by"),
				"order": crud.Integer().Enum(1, -1).Default(-1).Description("Sort the results ascending (1) or descending (-1)"),
				"name":  crud.String().Description("Filter by exact author name"),
			}).Unknown(false),
		},
	},
	{
		Method:  "POST",
		Path:    "/authors",
		Summary: "Create an author",
		Tags:    tags,
		Handler: CreateAuthor,
		Validate: crud.Validate{
			Body: crud.Object(map[string]crud.Field{
				"name": crud.String().Required().Description("Name of the author"),
				"born": crud.Date().Description("Author's date of birth"),
				"books": crud.Array().Items(crud.Object(map[string]crud.Field{
					"title": crud.String().Required().Description("Title of book"),
					"genre": crud.String().Enum("", "fantasy", "non-fiction", "political").Description("Book genre"),
				})),
			}),
		},
	},
	{
		Method:  "PATCH",
		Path:    "/authors/{id}",
		Summary: "Patch an author",
		Tags:    tags,
		Handler: PatchAuthor,
		Validate: crud.Validate{
			Path: crud.Object(map[string]crud.Field{
				"id": crud.String().Required(),
			}),
			Body: crud.Object(map[string]crud.Field{
				"name": crud.String().Description("Name of the author"),
				"born": crud.Date().Description("Author's date of birth"),
				"books": crud.Array().Items(crud.Object(map[string]crud.Field{
					"title": crud.String().Description("Title of book"),
					"genre": crud.String().Enum("", "fantasy", "non-fiction", "political").Description("Book genre"),
				})),
			}),
		},
	},
	{
		Method:  "DELETE",
		Path:    "/authors/{id}",
		Summary: "Delete an author",
		Tags:    tags,
		Handler: DeleteAuthor,
		Validate: crud.Validate{
			Path: crud.Object(map[string]crud.Field{
				"id": crud.String().Required(),
			}),
		},
	},
}
