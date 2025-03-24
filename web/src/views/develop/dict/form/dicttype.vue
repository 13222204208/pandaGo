<script setup lang="ts">
import { ref } from "vue";
import ReCol from "@/components/ReCol";
import { formRules } from "../utils/rule";
import { DictTypeDataProps } from "../utils/types";

const props = withDefaults(defineProps<DictTypeDataProps>(), {
  dictTypeInline: () => ({
    title: "新增",
    higherDeptOptions: [],
    pid: 0,
    name: "",
    type: ""
  })
});

const ruleFormRef = ref();
const newFormInline = ref(props.dictTypeInline);

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
        <el-form-item label="上级类型">
          <el-cascader
            v-model="newFormInline.pid"
            class="w-full"
            :options="newFormInline.higherDeptOptions"
            :props="{
              value: 'id',
              label: 'name',
              emitPath: false,
              checkStrictly: true
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

      <re-col :value="24">
        <el-form-item label="类型名称" prop="name">
          <el-input
            v-model="newFormInline.name"
            clearable
            placeholder="请输入类型名称"
          />
        </el-form-item>
      </re-col>

      <re-col :value="24">
        <el-form-item label="类型标识" prop="type">
          <el-input
            v-model="newFormInline.type"
            clearable
            placeholder="请输入类型标识"
          />
        </el-form-item>
      </re-col>
    </el-row>
  </el-form>
</template>
