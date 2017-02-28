package kec

import (
	"github.com/hdksky/ksyungo/common"
	"github.com/hdksky/ksyungo/util"
)

// ImageType https://docs.ksyun.com/read/latest/52/_book/oaImage.html
type ImageType struct {
	Name         string
	ImageId      string
	CreationDate util.ISO6801Time
	IsPublic     string
	Platform     string
	ImageState   string
	InstanceId   string
}

type DescribeImagesResponse struct {
	common.Response
	ImagesSet struct {
		Item []ImageType `json:"ImagesSet"`
	}
}

// DescribeImages 获取所有镜像
// You can read doc at https://docs.ksyun.com/read/latest/52/_book/oaDescribeImages.html
func (c *Client) DescribeImages(imageid string) ([]ImageType, error) {
	response := DescribeImagesResponse{}
	var args interface{}
	if imageid != "" {
		args = imageid
	}
	err := c.Invoke("DescribeImages", args, &response)
	if err == nil {
		return response.ImagesSet.Item, nil
	}
	return nil, err
}

type CreateImageArgs struct {
	InstanceId string
	Name       string
}

type CreateImageResponse struct {
	common.Response
	ImageId string
}

// CreateImage create image
// You can read doc at https://docs.ksyun.com/read/latest/52/_book/oaCreateImage.html
func (c *Client) CreateImage(args *CreateImageArgs) (string, error) {
	response := CreateImageResponse{}
	err := c.Invoke("CreateImage", args, &response)
	if err == nil {
		return response.ImageId, nil
	}
	return "", err
}

type ImageDeletion struct {
	ImageId string
	Return  string
}

type RemoveImageResponse struct {
	common.Response
	ReturnSet struct {
		Item []ImageDeletion
	}
}

// RemoveImages remove image
// You can read doc at https://docs.ksyun.com/read/latest/52/_book/oaRemoveImages.html
func (c *Client) RemoveImages(imageIds []string) ([]ImageDeletion, error) {
	response := RemoveImageResponse{}
	err := c.Invoke("RemoveImages", imageIds, &response)
	if err == nil {
		return response.ReturnSet.Item, nil
	}
	return nil, err
}

type ModifyImageAttributeArgs struct {
	ImageId string
	Name    string
}

type ModifyImageAttributeResponse struct {
	common.Response
	Return string
}

// ModifyImageAttribute modify image name attribute
// You can read doc at https://docs.ksyun.com/read/latest/52/_book/oaModifyImageAttribute.html
func (c *Client) ModifyImageAttribute(args *ModifyImageAttributeArgs) (string, error) {
	response := ModifyImageAttributeResponse{}
	err := c.Invoke("ModifyImageAttribute", args, &response)
	if err == nil {
		return response.Return, nil
	}
	return "", err
}
