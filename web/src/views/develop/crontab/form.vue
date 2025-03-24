<script setup lang="ts">
import { ref } from "vue";
import { formRules } from "./utils/rule";
import { FormProps } from "./utils/types";

const props = withDefaults(defineProps<FormProps>(), {
  formInline: () => ({
    id: undefined,
    title: "",
    name: "",
    params: "",
    pattern: "",
    policy: 1,
    sort: 0,
    status: 1,
    remark: "",
    createdAt: "",
    updatedAt: ""
  })
});

const ruleFormRef = ref();
const newFormInline = ref(props.formInline);

function getRef() {
  return ruleFormRef.value;
}

defineExpose({ getRef });
</script>

<template>
  <el-form
    ref="ruleFormRef"
    :model="newFormInline"
    :rules="formRules"
    label-width="100px"
  >
    <el-form-item label="任务标题" prop="title">
      <el-input
        v-model="newFormInline.title"
        clearable
        placeholder="请输入任务标题"
      />
    </el-form-item>

    <el-form-item label="任务名称" prop="name">
      <el-input
        v-model="newFormInline.name"
        clearable
        placeholder="请输入任务名称"
      />
    </el-form-item>

    <el-form-item label="任务参数" prop="params">
      <el-input
        v-model="newFormInline.params"
        clearable
        placeholder="请输入任务参数"
      />
    </el-form-item>

    <el-form-item label="Cron表达式" prop="pattern">
      <el-input
        v-model="newFormInline.pattern"
        clearable
        placeholder="请输入Cron表达式"
      />
    </el-form-item>

    <el-form-item label="执行策略" prop="policy">
      <el-select v-model="newFormInline.policy" placeholder="请选择执行策略">
        <el-option label="立即执行" :value="1" />
        <el-option label="执行一次" :value="2" />
        <el-option label="放弃执行" :value="3" />
      </el-select>
    </el-form-item>

    <el-form-item label="排序" prop="sort">
      <el-input-number v-model="newFormInline.sort" :min="0" />
    </el-form-item>

    <el-form-item label="状态" prop="status">
      <el-radio-group v-model="newFormInline.status">
        <el-radio :label="1">启用</el-radio>
        <el-radio :label="0">停用</el-radio>
      </el-radio-group>
    </el-form-item>

    <el-form-item label="备注">
      <el-input
        v-model="newFormInline.remark"
        placeholder="请输入备注信息"
        type="textarea"
      />
    </el-form-item>
  </el-form>
</template>
