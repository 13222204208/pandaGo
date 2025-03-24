<script setup lang="ts">
import { ref } from "vue";
import ReCol from "@/components/ReCol";
import { formRules } from "../utils/rule";
import { FormProps } from "../utils/types";
import { usePublicHooks } from "../../hooks";

const props = withDefaults(defineProps<FormProps>(), {
  formInline: () => ({
    title: "新增",
    higherDeptOptions: [],
    label: "",
    value: "",
    valueType: "string",
    type: "",
    listClass: "primary",
    isDefault: 0,
    sort: 0,
    remark: "",
    status: 1
  })
});

const valueOptions = [
  {
    value: "string",
    label: "字符串"
  },
  {
    value: "int",
    label: "整数"
  },
  {
    value: "float64",
    label: "小数"
  },
  {
    value: "bool",
    label: "布尔值"
  }
];

const labelCssOptions: {
  value: "primary" | "success" | "warning" | "info" | "danger";
  label: string;
}[] = [
  {
    value: "primary",
    label: "主要"
  },
  {
    value: "success",
    label: "成功"
  },
  {
    value: "warning",
    label: "警告"
  },
  {
    value: "info",
    label: "信息"
  },
  {
    value: "danger",
    label: "危险"
  }
];

const ruleFormRef = ref();
const { switchStyle } = usePublicHooks();
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
    label-width="82px"
  >
    <el-row :gutter="30">
      <re-col :value="12" :xs="24" :sm="24">
        <el-form-item label="字典类型">
          <el-cascader
            v-model="newFormInline.type"
            class="w-full"
            :options="newFormInline.higherDeptOptions"
            :props="{
              value: 'type',
              label: 'name',
              emitPath: false
            }"
            clearable
            filterable
            placeholder="请选择字典类型"
          >
            <template #default="{ node, data }">
              <span>{{ data.name }}</span>
              <span v-if="!node.isLeaf"> ({{ data.children.length }}) </span>
            </template>
          </el-cascader>
        </el-form-item>
      </re-col>

      <re-col :value="12" :xs="24" :sm="24">
        <el-form-item label="字典标签" prop="label">
          <el-input
            v-model="newFormInline.label"
            clearable
            placeholder="请输入字典标签"
          />
        </el-form-item>
      </re-col>

      <re-col :value="12" :xs="24" :sm="24">
        <el-form-item label="标签样式">
          <el-select
            v-model="newFormInline.listClass"
            placeholder="请选择样式"
            class="w-full"
            clearable
          >
            <el-option
              v-for="(item, index) in labelCssOptions"
              :key="index"
              :label="item.label"
              :value="item.value"
            >
              <el-tag :type="item.value">
                {{ item.label }}
              </el-tag>
            </el-option>
          </el-select>
        </el-form-item>
      </re-col>

      <re-col :value="12" :xs="24" :sm="24">
        <el-form-item label="字典键值" prop="value">
          <el-input
            v-model="newFormInline.value"
            clearable
            placeholder="请输入字典键值"
          />
        </el-form-item>
      </re-col>

      <re-col :value="12" :xs="24" :sm="24">
        <el-form-item label="键值类型">
          <el-select
            v-model="newFormInline.valueType"
            placeholder="请选择类型"
            class="w-full"
            clearable
          >
            <el-option
              v-for="(item, index) in valueOptions"
              :key="index"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </re-col>
      <re-col :value="12" :xs="24" :sm="24">
        <el-form-item label="排序" prop="sort">
          <el-input
            v-model="newFormInline.sort"
            type="number"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
      </re-col>

      <re-col
        v-if="newFormInline.title === '新增'"
        :value="12"
        :xs="24"
        :sm="24"
      >
        <el-form-item label="状态">
          <el-switch
            v-model="newFormInline.status"
            inline-prompt
            :active-value="1"
            :inactive-value="0"
            active-text="启用"
            inactive-text="停用"
            :style="switchStyle"
          />
        </el-form-item>
      </re-col>

      <re-col>
        <el-form-item label="备注">
          <el-input
            v-model="newFormInline.remark"
            placeholder="请输入备注信息"
            type="textarea"
          />
        </el-form-item>
      </re-col>
    </el-row>
  </el-form>
</template>
