package eip

import (
	"github.com/hdksky/ksyungo/common"
	"time"
)

var (
	LINE_EXPIRED_SEC int64 = 3600
)

type LineType struct {
	LineName string
	LineId   string
	LineType string
}

type GetLinesResponse struct {
	common.Response
	LineSet struct {
		Item []LineType `xml:"item"`
	}
}

// GetLines 获取用户可选链路信息
// You can read doc at https://docs.ksyun.com/read/latest/57/_book/Action/GetLines.html
func (c *Client) GetLines() ([]LineType, error) {
	response := &GetLinesResponse{}
	err := c.Invoke("GetLines", nil, response)
	if err == nil {
		return response.LineSet.Item, nil
	}
	return nil, err
}

var cachedLines struct {
	lineM      map[string][]LineType
	lastSecond int64
}

func GetAllLines(accessId, accessKey string) (map[string][]LineType, error) {
	if cacheExpired() {
		l, err := getAllLines(accessId, accessKey)
		if err != nil {
			return nil, err
		}

		if len(l) == len(regions) {
			cachedLines.lineM = l
			cachedLines.lastSecond = time.Now().Unix()

			return l, nil
		}
	}

	return cachedLines.lineM, nil
}

func cacheExpired() bool {
	// cachedLines.lastSecond == 0 means first opt
	return time.Now().Unix()-cachedLines.lastSecond > LINE_EXPIRED_SEC
}

func getAllLines(accessId, accessKey string) (map[string][]LineType, error) {
	ret := make(map[string][]LineType, len(regions))

	for _, region := range regions {
		rLines, err := NewClient(accessId, accessKey, region).GetLines()
		if err != nil {
			return nil, err
		}

		ret[region] = rLines
	}

	return ret, nil
}
