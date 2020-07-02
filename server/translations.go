package server

import (
	"context"
	"fmt"
	"time"

	"github.com/mattermost/mattermost-mattermod/model"
	"github.com/mattermost/mattermost-server/v5/mlog"
)

func (s *Server) handleTranslationPR(ctx context.Context, pr *model.PullRequest) {
	if pr.Username != s.Config.TranslationsBot {
		return
	}

	dataMsg := fmt.Sprintf("####[%v translations PR %v](%v)\n", pr.RepoName, time.Now().UTC().Format(time.RFC3339), pr.URL)
	msg := dataMsg + s.Config.TranslationsMattermostMessage
	mlog.Debug("Sending Mattermost message", mlog.String("message", msg))

	webhookRequest := &Payload{Username: "Weblate", Text: msg}
	r, err := s.sendToWebhook(ctx, s.Config.TranslationsMattermostWebhookURL, webhookRequest)
	if err != nil {
		mlog.Error("Unable to post to Mattermost webhook", mlog.Err(err))
		return
	}

	closeBody(r)
}
