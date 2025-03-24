<script setup lang="ts">
import { ref } from "vue";

const columns = [
  { label: "序号", type: "index", width: 60 },
  { label: "字段列名", prop: "columnName", width: 120 },
  { label: "字段描述", prop: "columnComment", width: 120 },
  { label: "物理类型", prop: "columnType", width: 120 },
  { label: "必填", prop: "required", width: 60 },
  { label: "查询方式", prop: "queryType", width: 120 },
  { label: "显示类型", prop: "showType", width: 120 },
  { label: "列表排序", prop: "listSort", width: 80 },
  { label: "编辑排序", prop: "editSort", width: 80 },
  { label: "查询排序", prop: "querySort", width: 80 },
  { label: "导出排序", prop: "exportSort", width: 80 },
  { label: "列表宽度", prop: "listWidth", width: 80 },
  { label: "字典类型", prop: "dictType", width: 120 },
  { label: "关联表", prop: "relationTable", width: 120 }
];

const tableData = ref([
  {
    columnName: "created_at",
    columnComment: "创建日期",
    columnType: "DATETIME",
    required: true,
    queryType: "BETWEEN",
    showType: "日期时间选择器",
    listSort: 7,
    editSort: 7,
    querySort: 7,
    exportSort: 7,
    listWidth: 100,
    dictType: "",
    relationTable: ""
  }
  // ... 其他字段数据
]);

defineExpose({
  tableData
});
</script>

<template>
  <div class="bg-white p-6 rounded-lg shadow-sm">
    <div class="text-lg font-bold mb-4">字段信息配置</div>
    <el-table
      :data="tableData"
      border
      style="width: 100%"
      height="calc(100vh - 300px)"
    >
      <el-table-column
        v-for="col in columns"
        :key="col.prop"
        :label="col.label"
        :prop="col.prop"
        :width="col.width"
        align="center"
      >
        <template v-if="col.prop === 'required'" #default="{ row }">
          <el-checkbox v-model="row.required" />
        </template>
        <template v-else-if="col.prop === 'queryType'" #default="{ row }">
          <el-select v-model="row.queryType" class="w-full">
            <el-option label="=" value="=" />
            <el-option label="BETWEEN" value="BETWEEN" />
            <el-option label="单图上传" value="单图上传" />
            <el-option label="多图上传" value="多图上传" />
            <el-option label="富文本" value="富文本" />
          </el-select>
        </template>
        <template v-else-if="col.prop === 'showType'" #default="{ row }">
          <el-select v-model="row.showType" class="w-full">
            <el-option label="文本框" value="文本框" />
            <el-option label="文本域" value="文本域" />
            <el-option label="下拉框" value="下拉框" />
            <el-option label="单选框" value="单选框" />
            <el-option label="复选框" value="复选框" />
            <el-option label="日期选择器" value="日期选择器" />
            <el-option label="日期时间选择器" value="日期时间选择器" />
          </el-select>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<style scoped>
.el-table {
  --el-table-border-color: var(--el-border-color-lighter);
}
</style>
