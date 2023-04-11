package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Manifest struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Image    string             `json:"image,omitempty" validate:"required"`
	Site_en     string             `json:"site_en,omitempty" validate:"required"`
	Site_jp     string             `json:"site_jp,omitempty" validate:"required"`
	Location_en string             `json:"location,omitempty" validate:"required"`
	Location_jp string             `json:"location,omitempty" validate:"required"`
	Year_from     string             `json:"year_from,omitempty" validate:"required"`
	Year_until     string             `json:"year_until,omitempty" validate:"required"`
	Problem_en string             `json:"problem_en,omitempty" validate:"required"`
	Problem_jp string             `json:"problem_jp,omitempty" validate:"required"`
	Dimensions string             `json:"dimensions,omitempty" validate:"required"`
	Material string	 	      `json:"material_en,omitempty" validate:"required"`
	Material string	 	      `json:"material_jp,omitempty" validate:"required"`
	Formula_en string	 	      `json:"formula_en,omitempty" validate:"required"`
	Formula_jp string	 	      `json:"formula_jp,omitempty" validate:"required"`
	Solution_en string	 	      `json:"solution_en,omitempty" validate:"required"`
	Solution_jp string	 	      `json:"solution_jp,omitempty" validate:"required"`
	School_en string	 	      `json:"school_en,omitempty" validate:"required"`
	School_jp string	 	      `json:"school_jp,omitempty" validate:"required"`
	Author_en string	 	      `json:"author_en,omitempty" validate:"required"`
	Author_jp string	 	      `json:"author_jp,omitempty" validate:"required"`
	Accessible boolean 	 	      `json:"accessible,omitempty" validate:"required"`
	Lit_sources string	 	      `json:"author_jp,omitempty" validate:"required"`
}
