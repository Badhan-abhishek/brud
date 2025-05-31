package pagination

type PaginationInput struct {
	Skip  *int `json:"skip"`
	Limit *int `json:"limit"`
}

func CreatePagination(skip, limit int) *PaginationInput {
	if skip < 0 {
		skip = 0
	}
	if limit <= 0 {
		limit = 10
	}
	return &PaginationInput{
		Skip:  &skip,
		Limit: &limit,
	}
}

func CalculateOffset(skip, limit int) int {
	if skip < 0 {
		skip = 0
	}
	if limit <= 0 {
		limit = 10
	}
	return skip * limit
}

type PaginationMeta struct {
	TotalCount int `json:"total_count"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
}

type PaginationOutput struct {
	Items any            `json:"items"`
	Meta  PaginationMeta `json:"meta"`
}

func CreatePaginationOutput(totalCount, skip, limit int, data any) *PaginationOutput {
	if totalCount < 0 {
		totalCount = 0
	}
	if skip < 0 {
		skip = 0
	}
	if limit <= 0 {
		limit = 10
	}

	page := skip / limit
	if skip%limit > 0 {
		page++
	}

	return &PaginationOutput{
		Items: data,
		Meta: PaginationMeta{
			TotalCount: totalCount,
			Page:       page,
			Limit:      limit,
		},
	}
}
