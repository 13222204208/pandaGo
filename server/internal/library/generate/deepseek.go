package generate

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/guonaihong/gout"
)

// Message DeepSeek消息结构
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// DeepSeekRequest 请求结构
type DeepSeekRequest struct {
	Model            string    `json:"model"`
	Messages         []Message `json:"messages"`
	MaxTokens        int       `json:"max_tokens"`
	Temperature      float64   `json:"temperature"`
	TopP             float64   `json:"top_p"`
	FrequencyPenalty float64   `json:"frequency_penalty,omitempty"`
	PresencePenalty  float64   `json:"presence_penalty,omitempty"`
	Stream           bool      `json:"stream,omitempty"`
	ResponseFormat   struct {
		Type string `json:"type"`
	} `json:"response_format,omitempty"`
	ToolChoice string `json:"tool_choice,omitempty"`
}

// DeepSeekResponse 响应结构
type DeepSeekResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
}

// 提示词常量
const (
	UserRole     = "user"
	SystemRole   = "system"
	SQLAssistant = `请帮我创建一个MySQL 8.0及以上版本兼容的数据表，要求如下：
					1. 只回答关于创建数据表的SQL语句，不要回答其他内容。
					2. 存储引擎：InnoDB
					3. 字符集：utf8mb4（支持完整Unicode，包括emoji）
					4. 排序规则：utf8mb4_0900_ai_ci（MySQL 8.0默认）
					5. 需要包含的字段：
					- [字段1名称] [数据类型] [约束] COMMENT '[字段注释]'
					- [字段2名称] [数据类型] [约束] COMMENT '[字段注释]'
					- ...
					6. 主键：[指定主键字段]
					7. 需要创建的索引（如果有）：
					- [索引1名称] ([字段名])
					8. 表注释：'[表用途描述]'
					9. 使用MySQL 8.0特性（如自增持久化、窗口函数等，如有需要）
					10. 包含创建时间和更新时间字段（自动维护）
					11. 输出结果为完整的SQL语句,不要包含任何其他内容
					12. 去除可能存在的Markdown标记,生成的sql语句必须是直接可以执行的sql语句

					请使用规范的SQL语法，包含完整的字段注释和表注释。`
)

// ChatCompletion 发送聊天请求到DeepSeek
func chatCompletion(messages []Message) (*DeepSeekResponse, error) {
	ctx := gctx.New()
	var (
		apiURL   = g.Cfg().MustGet(ctx, "deepseek.url").String()
		apiModel = g.Cfg().MustGet(ctx, "deepseek.model").String()
		apiKey   = g.Cfg().MustGet(ctx, "deepseek.apiKey").String()
	)

	// 构建请求体
	reqData := DeepSeekRequest{
		Model:       apiModel,
		Messages:    messages,
		MaxTokens:   2048,
		Temperature: 0.7,
		TopP:        0.9,
		ResponseFormat: struct {
			Type string `json:"type"`
		}{
			Type: "text",
		},
	}

	// 使用gout发送请求
	var response DeepSeekResponse
	var code int
	err := gout.POST(apiURL).
		SetJSON(reqData).
		SetHeader(gout.H{
			"Content-Type":  "application/json",
			"Accept":        "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", apiKey),
		}).
		BindJSON(&response).
		Code(&code).
		Do()

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("DeepSeek API请求失败，状态码: %d", code)
	}

	return &response, nil
}

// CreateTableSql 根据描述创建表SQL语句
func CreateTableSql(tableDesc string) (string, error) {
	messages := []Message{
		{
			Role:    SystemRole,
			Content: SQLAssistant,
		},
		{
			Role:    UserRole,
			Content: tableDesc,
		},
	}

	response, err := chatCompletion(messages)
	if err != nil {
		return "", err
	}

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}

	return "", nil
}
