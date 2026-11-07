package llm

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"

	"github.com/openai/openai-go/v3"
)

const captionPrompt = `Describe this image in 2-3 sentences for a search index. 
	Focus on the visible subject - what it is, key details, style. 
	Do not interpret or speculate; describe only what is shown.`

func (c *Client) HasVision() bool {
	return c.cfg.VisionModel != ""
}

func (c *Client) DescribeImage(ctx context.Context, mime string, image []byte) (string, error) {
	if !c.HasVision() {
		return "", errors.New("vision model not configured")
	}

	if len(image) == 0 {
		return "", errors.New("image is empty")
	}

	if mime == "" {
		mime = http.DetectContentType(image)
	}

	dataURL := "data:" + mime + ";base64," + base64.StdEncoding.EncodeToString(image)

	resp, err := c.sdk.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model: c.cfg.VisionModel,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage([]openai.ChatCompletionContentPartUnionParam{
				openai.TextContentPart(captionPrompt),
				openai.ImageContentPart(openai.ChatCompletionContentPartImageImageURLParam{
					URL: dataURL,
				}),
			}),
		},
	})
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from model")
	}

	return resp.Choices[0].Message.Content, nil

}
