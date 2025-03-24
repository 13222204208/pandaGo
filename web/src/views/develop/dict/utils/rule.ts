import { reactive } from "vue";
import type { FormRules } from "element-plus";

/** 自定义表单规则校验 */
export const formRules = reactive(<FormRules>{
  label: [{ required: true, message: "字典标签为必填项", trigger: "blur" }],
  value: [{ required: true, message: "字典键值为必填项", trigger: "blur" }],
  type: [{ required: true, message: "字典类型为必填项", trigger: "blur" }]
});
