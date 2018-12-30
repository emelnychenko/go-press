package models

type (
	PaginationQuery struct {
		Limit int `query:"limit" validate:"gte=1,lte=100"`
		Start int `query:"start" validate:"gte=0"`
		Page  int `query:"page" validate:"gte=1"`
	}

	PaginationResult struct {
		Total int         `json:"total" xml:"total"`
		Data  interface{} `json:"data" xml:"data"`
	}
)

func (q *PaginationQuery) Offset() int {
	return (q.Page-1)*q.Limit + q.Start
}
