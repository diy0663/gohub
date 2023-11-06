package link_dto

import "github.com/diy0663/gohub/app/models/link"

//保存之后自动import

func ConvertLinkToDTO(link *link.Link) *LinkDTO {
	if link == nil {
		return nil
	}
	return &LinkDTO{

		LinkId:      link.GetStringID(),
		NameWithUrl: link.Name + " : " + link.Url,
		CreatedAt:   link.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ConvertLinksToDTO(links []*link.Link) []*LinkDTO {
	if links == nil {
		return nil
	}
	var dtos []*LinkDTO
	for _, link := range links {
		dtos = append(dtos, ConvertLinkToDTO(link))
	}
	return dtos
}
