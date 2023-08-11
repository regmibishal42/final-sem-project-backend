package model

import "backend/pkg/util"

func (p *OffsetPaginationFilter) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *OffsetPaginationFilter) GetLimit() int {
	if p.Limit == nil {
		p.Limit = util.Ref(10)
	}
	return *p.Limit
}

func (p *OffsetPaginationFilter) GetPage() int {
	if p.Page == nil {
		p.Page = util.Ref(1)
	}
	return *p.Page
}

func (p *OffsetPaginationFilter) GetSort() string {
	if p.Sort == nil {
		p.Sort = util.Ref(SortTypeDesc)
	}
	if p.Column == nil {
		p.Column = util.Ref("created_at")
	}
	return *p.Column + " " + p.Sort.String()
}
