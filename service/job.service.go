package service

import (
	"dansmultipro/recruitment/dto"
	"dansmultipro/recruitment/utils"
	"net/url"
)

type JobService interface {
	GetList(query url.Values) []dto.JobResponse
	GetById(id string) dto.JobResponse
}

type jobService struct {
}

func NewJobService() JobService {
	return &jobService{}
}


func (s *jobService) GetList(query url.Values) []dto.JobResponse {

	url := "http://dev3.dansmultipro.co.id/api/recruitment/positions.json?"

	if val, ok := query["description"]; ok {
		url = url + "description=" + val[0]
	}

	if val, ok := query["location"]; ok {
		url = url + "location=" + val[0]
	}

	if val, ok := query["full_time"]; ok {
		url = url + "full_time=" + val[0]
	}

	if val, ok := query["page"]; ok {
		url = url + "page=" + val[0]
	}

	res := utils.GetResponses(url)
	return res

}


func (s *jobService) GetById(id string) dto.JobResponse  {
	url := "http://dev3.dansmultipro.co.id/api/recruitment/positions/"
	res := utils.GetResponse(url, id)
	return res
}