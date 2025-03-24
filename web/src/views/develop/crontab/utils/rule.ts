import { reactive } from "vue";
import type { FormRules } from "element-plus";

/** 自定义表单规则校验 */
export const formRules = reactive(<FormRules>{
  title: [{ required: true, message: "任务标题为必填项", trigger: "blur" }],
  name: [{ required: true, message: "任务名称为必填项", trigger: "blur" }],
  pattern: [{ required: true, message: "Cron表达式为必填项", trigger: "blur" }],
  policy: [{ required: true, message: "执行策略为必填项", trigger: "change" }],
  status: [{ required: true, message: "状态为必填项", trigger: "change" }]
});
