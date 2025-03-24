package consts

// 定时任务

const (
	CronSplitStr     = "," // 变量分割符
	CronPolicySame   = 1   // 并行策略
	CronPolicySingle = 2   // 单例策略
	CronPolicyOnce   = 3   // 单次策略
	CronPolicyTimes  = 4   // 多次策略
)

// 任务状态码
const (
	StatusALL     int = -1 // 全部状态
	StatusEnabled int = 1  // 启用
	StatusDisable int = 2  // 禁用
	StatusDelete  int = 3  // 已删除
)

type CtxKey string

const (
	ContextHTTPKey     CtxKey = "httpContext" // http上下文变量名称
	ContextKeyCronArgs CtxKey = "cronArgs"    // 定时任务参数
	ContextKeyCronSn   CtxKey = "cronSn"      // 定时任务序列号
)
