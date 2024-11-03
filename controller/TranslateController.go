package controllers

import (
	"ApiController/untils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	url2 "net/url"
)

type TranslateController struct {
}
type UserRequest struct {
	Be      string `json:"be"`
	To      string `json:"to"`
	Message string `json:"message"`
	Type    int    `json:"type"`
}

func (Tr TranslateController) Translation(c *gin.Context) {
	var userrequest UserRequest
	if err := c.BindJSON(&userrequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tk := untils.JavaScript("tools.js", "VL", userrequest.Message)
	url := untils.ConstructUrl("https://translate.s9y.in", tk, url2.QueryEscape(userrequest.Message), userrequest.Be, userrequest.To)
	body := untils.GetUrl(url)
	if userrequest.Type == 1 {
		var parsedData []interface{}
		json.Unmarshal(body, &parsedData)
		// 提取原文和翻译文本
		translationInfo := parsedData[0].([]interface{})[0].([]interface{})
		originalText := translationInfo[1].(string)
		translatedText := translationInfo[0].(string)
		// 提取汉语拼音
		pinyin := ""
		if userrequest.Be == "en" {
			pinyinInfo := parsedData[0].([]interface{})[1].([]interface{})
			pinyin = pinyinInfo[2].(string)
		}

		// 构造 JSON 数据结构
		type Phrase struct {
			Text        string `json:"text"`
			Translation string `json:"translation"`
		}

		type ExtractedData struct {
			OriginalText   string `json:"original_text"`
			TranslatedText string `json:"translated_text"`
			Pinyin         string `json:"pinyin"`
		}

		// 构造提取的数据结构
		extractedData := ExtractedData{
			OriginalText:   originalText,
			TranslatedText: translatedText,
			Pinyin:         pinyin,
		}
		c.JSON(200, extractedData)
	} else {
		c.String(200, string(body))
	}

}

func (Tr TranslateController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "请使用Post请求",
	})
}
