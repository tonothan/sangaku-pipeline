package models

import (
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sangaku struct {
	Id primitive.ObjectID `json:"id,omitempty"`

	// Time-Space metadata
	Temple_en   string `json:"temple_en,omitempty"`
	Temple_jp   string `json:"temple_jp,omitempty"`
	Location_en string `json:"location_en,omitempty"`
	Location_jp string `json:"location_jp,omitempty"`
	Year_from   string `json:"year_from,omitempty"`
	Year_to     string `json:"year_to,omitempty"`

	// Content data
	Transcription_en string `json:"transcription_en,omitempty"`
	Transcription_jp string `json:"transcription_jp,omitempty"`
	Problem_en       string `json:"problem_en,omitempty"`
	Problem_jp       string `json:"problem_jp,omitempty"`
	Formula_en       string `json:"formula_en,omitempty"`
	Formula_jp       string `json:"formula_jp,omitempty"`
	Solution_en      string `json:"solution_en,omitempty"`
	Solution_jp      string `json:"solution_jp,omitempty"`
	School_en        string `json:"school_en,omitempty"`
	School_jp        string `json:"school_jp,omitempty"`
	Author_en        string `json:"author_en,omitempty"`
	Author_jp        string `json:"author_jp,omitempty"`

	// Medium metadata
	Dimension   string `json:"dimension,omitempty"`
	Material_en string `json:"material_en,omitempty"`
	Material_jp string `json:"material_jp,omitempty"`

	// Images metadata
	Images Sangaku_Images `json:"images,omitempty"`
}

type Sangaku_Images struct {
	Image []Sangaku_Image
}

type Sangaku_Image struct {
	Name     string                `form:"name"`
	FileData *multipart.FileHeader `form:"file"`
}
