package contorolers

import (
	"github.com/MrMohebi/golang-gin-boilerplate.git/common"
	"github.com/MrMohebi/golang-gin-boilerplate.git/myTel"
	"github.com/MrMohebi/golang-gin-boilerplate.git/postReqs"
	"github.com/anaskhan96/soup"
	"github.com/gin-gonic/gin"
	tele "gopkg.in/telebot.v3"
)

type SendTelegramRequest struct {
	Message *tele.Message `json:"message"`
	User    *tele.User    `json:"user"`
}

func Telegram() gin.HandlerFunc {
	return func(context *gin.Context) {

		var requestBody SendTelegramRequest
		err := context.BindJSON(&requestBody)
		common.IsErr(err)

		postReqBody, err := postReqs.PostSendRequest(requestBody.Message.Text)
		common.IsErr(err)
		doc := soup.HTMLParse(postReqBody)
		trackers := doc.Find("div", "id", "pnlResult").FindAll("div")[0].FindAll("div")

		result := ""
		for _, tracker := range trackers {
			text := ""
			allDivs := tracker.FindAll("div")

			if len(allDivs) == 1 {
				text += allDivs[0].Text() + "\n"
			}

			if len(allDivs) == 3 {
				text += "-------------------------------------------------" + "\n"
				text += "```" + allDivs[0].Text() + "```\n"
			}

			if len(allDivs) == 4 {
				text += allDivs[3].Text() + " ==> " + allDivs[1].Text() + "\n\n"
			}

			if len(text) > 0 {
				result += text
			}

		}
		myTel.SendText(result, requestBody.User.ID)

		context.JSON(200, gin.H{
			"message": "ok",
		})
	}
}
