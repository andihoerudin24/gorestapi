package utils

import (
	"fmt"
	"math"
)

type PageLink struct {
	Page          int64
	Url           string
	IsCurrentPage bool
}

type PaginationLink struct {
	CurrentPage string
	NextPage    string
	PrevPage    string
	TotalRows   int64
	TotalPages  int64
	Links       []PageLink
}

type PaginationParams struct {
	Path        string
	TotalRows   int64
	PerPage     int64
	CurrentPage int64
}

func GetPaginationLinks(params PaginationParams) (PaginationLink, error) {
	var links []PageLink

	totalPages := int64(math.Ceil(float64(params.TotalRows) / float64(params.PerPage)))
	for i := 1; int64(i) <= totalPages; i++ {
		links = append(links, PageLink{
			Page:          int64(i),
			Url:           fmt.Sprintf("%s?page=%s", params.Path, fmt.Sprint(i)),
			IsCurrentPage: int64(i) == params.CurrentPage,
		})
	}
	var nextPage int64
	var prevPage int64
	prevPage = 1
	nextPage = totalPages

	if params.CurrentPage > 2 {
		prevPage = params.CurrentPage - 1
	}

	if params.CurrentPage < totalPages {
		nextPage = params.CurrentPage + 1
	}

	return PaginationLink{
		CurrentPage: fmt.Sprintf("%s?page=%s", params.Path, fmt.Sprint(params.CurrentPage)),
		NextPage:    fmt.Sprintf("%s?page=%s", params.Path, fmt.Sprint(nextPage)),
		PrevPage:    fmt.Sprintf("%s?page=%s", params.Path, fmt.Sprint(prevPage)),
		TotalRows:   params.TotalRows,
		Links:       links,
		TotalPages:  int64(len(links)),
	}, nil

}
