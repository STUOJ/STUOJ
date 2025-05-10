package yuki

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

// InitYukiImage 初始化 Yuki Image 服务客户端，并确保预期的相册和格式存在。
// 它会连接到 Yuki Image 服务，获取现有相册列表，
// 然后遍历 `wantAlbum` 中定义的期望相册配置。
// 如果相册不存在，则创建它；接着，检查并补充该相册所缺少的 `FormatSupport`。
func InitYukiImage(host, port, token string) error {
	config = YukiConf{
		Host:  host,
		Port:  port,
		Token: token,
	}
	preUrl = config.Host + ":" + config.Port + "/api/v1"
	log.Println("Connecting to yuki-image service: " + preUrl)
	albums, err := GetAlbumList()
	if err != nil {
		return err
	}
	albumsMap := make(map[string]YukiAlbum)
	for _, album := range albums {
		albumsMap[album.Name] = album
	}

	for _, wantedAlbum := range wantAlbum {
		currentAlbum, albumExists := albumsMap[wantedAlbum.Name]
		if !albumExists {
			log.Printf("Album '%s' not found, creating...", wantedAlbum.Name)
			createdAlbum, err := CreateAlbum(wantedAlbum.Name, int64(wantedAlbum.MaxHeight), int64(wantedAlbum.MaxWidth))
			if err != nil {
				log.Printf("Error creating album '%s': %v", wantedAlbum.Name, err)
				return err
			}
			albumsMap[createdAlbum.Name] = createdAlbum
			currentAlbum = createdAlbum
			log.Printf("Album '%s' created successfully with ID: %d", currentAlbum.Name, currentAlbum.Id)
		} else {
			log.Printf("Album '%s' already exists with ID: %d", currentAlbum.Name, currentAlbum.Id)
		}

		// 确保获取最新的相册信息，特别是 FormatSupport
		fullCurrentAlbum, err := GetAlbum(currentAlbum.Id)
		if err != nil {
			log.Printf("Error fetching full details for album '%s' (ID: %d): %v", currentAlbum.Name, currentAlbum.Id, err)
			return err
		}
		currentAlbum = fullCurrentAlbum

		// 检查并补充 FormatSupport
		currentAlbum.FormatSupport, err = AlbumGetFormat(GetAlbumRoleByName(wantedAlbum.Name))
		if err != nil {
			log.Printf("Error fetching formats for album '%s' (ID: %d): %v", currentAlbum.Name, currentAlbum.Id, err)
			return err
		}

		existingFormats := make(map[uint64]bool)
		for _, format := range currentAlbum.FormatSupport {
			existingFormats[format.Id] = true
		}

		albumRole := GetAlbumRoleByName(currentAlbum.Name)
		if albumRole == 0 {
			log.Printf("Unknown album role for name: %s", currentAlbum.Name)
			// 根据需要处理未知角色，例如跳过或返回错误
			continue
		}

		for _, wantedFormat := range wantedAlbum.FormatSupport {
			if !existingFormats[wantedFormat.Id] {
				log.Printf("Format '%s' not found in album '%s', adding...", GetFormatName(wantedFormat.Id), currentAlbum.Name)
				err := AlbumAddFormat(albumRole, wantedFormat.Id)
				if err != nil {
					log.Printf("Error adding format '%s' to album '%s': %v", GetFormatName(wantedFormat.Id), currentAlbum.Name, err)
					return err
				}
				log.Printf("Format '%s' added successfully to album '%s'", GetFormatName(wantedFormat.Id), currentAlbum.Name)
			} else {
				log.Printf("Format '%s' already exists in album '%s'", GetFormatName(wantedFormat.Id), currentAlbum.Name)
			}
		}
	}

	log.Println("Yuki-image service initialization and album/format check completed.")
	return nil
}

func httpInteraction(route string, httpMethod string, reader *bytes.Reader) (string, error) {
	url := preUrl + route
	var req *http.Request
	var err error
	if reader == nil {
		req, err = http.NewRequest(httpMethod, url, nil)
	} else {
		req, err = http.NewRequest(httpMethod, url, reader)
	}
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+config.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	bodystr := string(body)
	return bodystr, nil
}
