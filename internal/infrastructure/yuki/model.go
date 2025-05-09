package yuki

const (
	JPEG uint64 = 1
	PNG  uint64 = 2
	GIF  uint64 = 3
)

const (
	YukiAvatarAlbum  uint8 = 1
	YukiProblemAlbum uint8 = 2
	YukiBlogAlbum    uint8 = 3
)

var wantAlbum = []YukiAlbum{
	{
		Name:      GetAlbumName(YukiAvatarAlbum),
		MaxHeight: 512,
		MaxWidth:  512,
		FormatSupport: []YukiFormat{
			{
				Id: JPEG,
			},
			{
				Id: PNG,
			},
		},
	},
	{
		Name:      GetAlbumName(YukiProblemAlbum),
		MaxHeight: 1080,
		MaxWidth:  1920,
		FormatSupport: []YukiFormat{
			{
				Id: JPEG,
			},
			{
				Id: PNG,
			},
			{
				Id: GIF,
			},
		},
	},
	{
		Name:      GetAlbumName(YukiBlogAlbum),
		MaxHeight: 1080,
		MaxWidth:  1920,
		FormatSupport: []YukiFormat{
			{
				Id: JPEG,
			},
			{
				Id: PNG,
			},
			{
				Id: GIF,
			},
		},
	},
}

func GetFormatName(id uint64) string {
	switch id {
	case JPEG:
		return "jpeg"
	case PNG:
		return "png"
	case GIF:
		return "gif"
	default:
		return "unknown"
	}
}

// GetAlbumRoleByName 根据相册名称查找其对应的 uint8 角色常量。
func GetAlbumRoleByName(name string) uint8 {
	switch name {
	case "avatar":
		return YukiAvatarAlbum
	case "problem":
		return YukiProblemAlbum
	case "blog":
		return YukiBlogAlbum
	default:
		return 0 // 或者一个表示未知的常量
	}
}

// GetAlbumName 根据角色获取相册名称。
func GetAlbumName(role uint8) string {
	switch role {
	case YukiAvatarAlbum:
		return "avatar"
	case YukiProblemAlbum:
		return "problem"
	case YukiBlogAlbum:
		return "blog"
	default:
		return "unknown"
	}
}

type YukiResponses struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type YukiAlbum struct {
	Id            uint64        `json:"id,omitempty"`
	Name          string        `json:"name,omitempty"`
	MaxHeight     uint64        `json:"max_height,omitempty"`
	MaxWidth      uint64        `json:"max_width,omitempty"`
	FormatSupport []YukiFormat  `json:"format_support,omitempty"`
	UpdateTime    string        `json:"update_time,omitempty"`
	CreateTime    string        `json:"create_time,omitempty"`
	Image         YukiImageList `json:"image,omitempty"`
}

type YukiImageList struct {
	Image []YukiImage `json:"image,omitempty"`
	Page
}

type YukiFormat struct {
	Id   uint64 `json:"id"`
	Name string `json:"name,omitempty"`
}

type YukiFormatSupport struct {
	FormatId uint64 `json:"format_id"`
	AlbumId  uint64 `json:"album_id"`
}

type YukiImage struct {
	Key        string `json:"key"`
	Name       string `json:"name"`
	Url        string `json:"url,omitempty"`
	AlbumId    uint64 `json:"album_id,omitempty"`
	Pathname   string `json:"pathname"`
	OriginName string `json:"origin_name"`
	Size       uint64 `json:"size"`
	Mimetype   string `json:"mimetype"`
	Time       string `json:"time,omitempty"`
}

type YukiTmpInfo struct {
	Size  uint64 `json:"size"`
	Count uint64 `json:"count"`
}

type Page struct {
	Page  uint64 `json:"page,omitempty"`
	Size  uint64 `json:"size,omitempty"`
	Total uint64 `json:"total"`
}
