package yuki

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

func UploadImage(reader io.Reader, filename string, role uint8) (YukiImage, error) {
	url := preUrl + "/image"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return YukiImage{}, err
	}
	req.Header.Set("Authorization", "Bearer "+config.Token)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	err = writer.WriteField("album_name", GetAlbumName(role))
	if err != nil {
		return YukiImage{}, err
	}
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return YukiImage{}, err
	}

	_, err = io.Copy(part, reader)
	if err != nil {
		return YukiImage{}, err
	}
	err = writer.Close()
	if err != nil {
		return YukiImage{}, err
	}
	req.Body = io.NopCloser(body)
	req.ContentLength = int64(body.Len())
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// 如果发送请求失败，返回错误信息
		return YukiImage{}, err
	}
	defer resp.Body.Close()
	log.Println("resp.StatusCode: ", resp.StatusCode)
	bodys, err := io.ReadAll(resp.Body)
	if err != nil {
		return YukiImage{}, err
	}
	bodystr := string(bodys)
	var responses YukiResponses
	err = json.Unmarshal([]byte(bodystr), &responses)
	if err != nil {
		return YukiImage{}, err
	}
	if resp.StatusCode != http.StatusCreated {
		return YukiImage{}, errors.New(responses.Message)
	}
	var image YukiImage
	err = mapstructure.Decode(responses.Data, &image)
	if err != nil {
		return YukiImage{}, err
	}
	return image, nil
}

func GetImageList(page uint64, role uint8) (YukiImageList, error) {
	bodystr, err := httpInteraction("/album/image/"+GetAlbumName(role)+"/?page="+strconv.FormatUint(page, 10), "GET", nil)
	if err != nil {
		return YukiImageList{}, err
	}
	var responses YukiResponses
	err = json.Unmarshal([]byte(bodystr), &responses)
	if err != nil {
		return YukiImageList{}, err
	}
	if responses.Code == 0 {
		return YukiImageList{}, errors.New(responses.Message)
	}
	var imageList YukiImageList
	err = mapstructure.Decode(responses.Data, &imageList)
	if err != nil {
		return YukiImageList{}, err
	}
	return imageList, nil
}

func GetImage(key string) (YukiImage, error) {
	bodystr, err := httpInteraction("/image/"+key, "GET", nil)
	if err != nil {
		return YukiImage{}, err
	}
	var responses YukiResponses
	err = json.Unmarshal([]byte(bodystr), &responses)
	if err != nil {
		return YukiImage{}, err
	}
	if responses.Code == 0 {
		return YukiImage{}, errors.New(responses.Message)
	}
	var image YukiImage
	err = mapstructure.Decode(responses.Data, &image)
	if err != nil {
		return YukiImage{}, err
	}
	return image, nil
}

func GetImageFromUrl(url string) (YukiImage, error) {
	bodystr, err := httpInteraction("/image/?url="+url, "GET", nil)
	if err != nil {
		return YukiImage{}, err
	}
	var responses YukiResponses
	err = json.Unmarshal([]byte(bodystr), &responses)
	if err != nil {
		return YukiImage{}, err
	}
	if responses.Code == 0 {
		return YukiImage{}, errors.New(responses.Message)
	}
	var image YukiImage
	err = mapstructure.Decode(responses.Data, &image)
	if err != nil {
		return YukiImage{}, err
	}
	return image, nil
}

func DeleteImageByKey(key string) error {
	bodeystr, err := httpInteraction("/image"+"/"+key, "DELETE", nil)
	if err != nil {
		return err
	}
	var responses YukiResponses
	err = json.Unmarshal([]byte(bodeystr), &responses)
	if err != nil {
		return err
	}
	if responses.Code == 0 {
		return errors.New(responses.Message)
	}
	return nil
}

func DeleteImageByUrl(url string) error {
	image, err := GetImageFromUrl(url)
	if err != nil {
		return err
	}
	return DeleteImageByKey(image.Key)
}
