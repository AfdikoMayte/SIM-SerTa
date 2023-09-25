package helper

import (
	"github.com/afdikomayte/sim-layanan-sertifikat-tanah/model/domain"
	"github.com/afdikomayte/sim-layanan-sertifikat-tanah/model/web"
)

func ToUserResponse(user *domain.User) web.UserResponse {
	return web.UserResponse{
		Id:   user.Id,
		Nik:  user.Nik,
		Nama: user.Nama,
	}
}

func ToUserResponses(users []*domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}
