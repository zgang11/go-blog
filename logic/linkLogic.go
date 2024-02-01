package logic

import (
	"blog/dao/link"
	"blog/dao/topic"
	"blog/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

type LinkLogic struct {
}

func (this *LinkLogic) CreateLinkLogic(c *gin.Context) error {
	linkName := c.PostForm("linkName")
	url := c.PostForm("url")
	types, errType := strconv.Atoi(c.PostForm("type"))
	if errType != nil {
		types = 0
	}
	companyId, err_ := strconv.Atoi(c.PostForm("companyId"))
	if err_ != nil {
		companyId = 0
	}

	topicIdsStr := c.PostForm("topicIds")
	ids := []int{}
	topicIds := strings.Split(topicIdsStr, ",")
	for _, topicId := range topicIds {
		id, err := strconv.Atoi(topicId)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}

	topics := topic.GetTopicByIds(ids)

	params := &models.Link{LinkName: linkName, Url: url, CompanyId: companyId, Topics: topics, CreateTime: time.Now().UnixNano() / 1e6,
		UpdateTime: time.Now().UnixNano() / 1e6, Type: types}
	err := link.CreateLink(params)
	return err
}
