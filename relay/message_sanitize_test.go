package relay

import (
	"testing"

	"github.com/QuantumNous/new-api/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSanitizeOpenAIMessagesTrimsAssistantTrailingWhitespace(t *testing.T) {
	req := &dto.GeneralOpenAIRequest{
		Messages: []dto.Message{
			{Role: "user", Content: "keep me  \n"},
			{Role: "assistant", Content: "trim me  \t\n"},
		},
	}

	sanitizeOpenAIMessages(req, "gpt-5.5")

	assert.Equal(t, "keep me  \n", req.Messages[0].StringContent())
	assert.Equal(t, "trim me", req.Messages[1].StringContent())
}

func TestSanitizeOpenAIMessagesRemovesThinkingModelAssistantPrefill(t *testing.T) {
	req := &dto.GeneralOpenAIRequest{
		Messages: []dto.Message{
			{Role: "user", Content: "question"},
			{Role: "assistant", Content: "prefill"},
		},
	}

	sanitizeOpenAIMessages(req, "claude-sonnet-4-5-thinking")

	require.Len(t, req.Messages, 1)
	assert.Equal(t, "user", req.Messages[0].Role)
}

func TestSanitizeOpenAIMessagesKeepsNonThinkingAssistantTail(t *testing.T) {
	req := &dto.GeneralOpenAIRequest{
		Messages: []dto.Message{
			{Role: "user", Content: "question"},
			{Role: "assistant", Content: "allowed"},
		},
	}

	sanitizeOpenAIMessages(req, "gpt-5.5")

	require.Len(t, req.Messages, 2)
}
