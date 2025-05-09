package yuki

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/mitchellh/mapstructure"
)

func GetAlbumList() ([]YukiAlbum, error) {
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
	var albumList []YukiAlbum
	err = mapstructure.Decode(responses.Data, &albumList)
	if err != nil {
		return nil, err
	}
	return albumList, nil
}

func GetAlbum(albumId uint64) (YukiAlbum, error) {
	bodystr, err := httpInteraction("/album/"+string(albumId), "GET", nil)
	if err != nil {
		return YukiAlbum{}, err
	}
	var responses YukiResponses
	err = json.Unmarshal([]byte(bodystr), &responses)
	if err != nil {
		return YukiAlbum{}, err
	}
	if responses.Code == 0 {
		return YukiAlbum{}, errors.New(responses.Message)
	}
	var album YukiAlbum
	err = mapstructure.Decode(responses.Data, &album)
	if err != nil {
		return YukiAlbum{}, err
	}
	return album, nil
}

func CreateAlbum(name string, height, width int64) (YukiAlbum, error) {
	data := map[string]interface{}{
		"name":   name,
		"height": height,
		"width":  width,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return YukiAlbum{}, err
	}
	bodystr, err := httpInteraction("/album", "POST", bytes.NewReader(jsonData))
	if err != nil {
		return YukiAlbum{}, err
	}
	var responses YukiResponses
	err = json.Unmarshal([]byte(bodystr), &responses)
	if err != nil {
		return YukiAlbum{}, err
	}
	if responses.Code == 0 {
		return YukiAlbum{}, errors.New(responses.Message)
	}
	var album YukiAlbum
	err = mapstructure.Decode(responses.Data, &album)
	if err != nil {
		return YukiAlbum{}, err
	}
	return album, nil
}

func AlbumAddFormat(album uint8, format uint64) error {
	data := map[string]interface{}{
		"format_name": GetFormatName(format),
		"album_name":  GetAlbumName(album),
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	bodystr, err := httpInteraction("/album/format", "POST", bytes.NewReader(jsonData))
	if err != nil {
		return err
	}
	var responses YukiResponses
	err = json.Unmarshal([]byte(bodystr), &responses)
	if err != nil {
		return err
	}
	if responses.Code == 0 {
		return errors.New(responses.Message)
	}
	return nil
}
