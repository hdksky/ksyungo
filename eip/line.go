package eip

import "github.com/hdksky/ksyungo/common"

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
