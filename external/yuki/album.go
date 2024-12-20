package yuki

import (
	"STUOJ/internal/model"
	"encoding/json"
	"errors"

	"github.com/mitchellh/mapstructure"
)

func GetAlbumList() ([]model.YukiAlbum, error) {
	bodystr, err := httpInteraction("/album", "GET", nil)
	if err != nil {
		return nil, err
	}

	type tmpResponses struct {
		Code    int                      `json:"code"`
		Message string                   `json:"message"`
		Data    []map[string]interface{} `json:"data"`
	}
	var responses tmpResponses

	err = json.Unmarshal([]byte(bodystr), &responses)
	if err != nil {
		return nil, err
	}
	if responses.Code == 0 {
		return nil, errors.New(responses.Message)
	}
	var albumList []model.YukiAlbum
	err = mapstructure.Decode(responses.Data, &albumList)
	if err != nil {
		return nil, err
	}
	return albumList, nil
}

func GetAlbum(albumId uint64) (model.YukiAlbum, error) {
	bodystr, err := httpInteraction("/album/"+string(albumId), "GET", nil)
	if err != nil {
		return model.YukiAlbum{}, err
	}
	var responses model.YukiResponses
	err = json.Unmarshal([]byte(bodystr), &responses)
	if err != nil {
		return model.YukiAlbum{}, err
	}
	if responses.Code == 0 {
		return model.YukiAlbum{}, errors.New(responses.Message)
	}
	var album model.YukiAlbum
	err = mapstructure.Decode(responses.Data, &album)
	if err != nil {
		return model.YukiAlbum{}, err
	}
	return album, nil
}
