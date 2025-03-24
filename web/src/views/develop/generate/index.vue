<script setup lang="ts">
import { ref } from "vue";
import { useRole } from "./hook";
import { getPickerShortcuts } from "../utils";
import { PureTableBar } from "@/components/RePureTableBar";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import { ElMessage } from "element-plus";

import Delete from "@iconify-icons/ep/delete";
import Refresh from "@iconify-icons/ep/refresh";
import ImportTable from "./import.vue";
// 添加图标导入
import EditPen from "@iconify-icons/ep/edit-pen";
import Preview from "@iconify-icons/ep/view";
import Sync from "@iconify-icons/ep/refresh";
import Code from "@iconify-icons/ep/document";

defineOptions({
  name: "Generate"
});

const formRef = ref();
const tableRef = ref();
const importVisible = ref(false); // 添加导入弹窗的显示控制

const {
  form,
  loading,
  columns,
  dataList,
  pagination,
  selectedNum,
  onSearch,
  clearAll,
  resetForm,
  onbatchDel,
  // handleEdit,  // 如果hook中没有这个函数，注释掉
  // handleDelete, // 如果hook中没有这个函数，注释掉
  handleSizeChange,
  onSelectionCancel,
  handleCurrentChange,
  handleSelectionChange
} = useRole(tableRef);

// 自定义编辑处理函数
const handleEdit = (row: TableRow) => {
  ElMessage.info(`编辑表：${row.name}`);
};

// 自定义删除处理函数
const handleDelete = (row: TableRow) => {
  ElMessage.info(`删除表：${row.name}`);
  onSearch(); // 刷新表格数据
};

// 修改配置
const onEdit = () => {
  if (selectedNum.value !== 1) {
    ElMessage.warning("请选择一条记录");
    return;
  }
  const selected = tableRef.value.getTableRef().getSelectionRows()[0];
  handleEdit(selected);
};

// 生成代码
const onGenerate = () => {
  ElMessage.info("代码生成功能开发中");
};

// 新增配置
const onCreate = () => {
  ElMessage.info("新增功能开发中");
};

// 导入配置
const onImport = () => {
  importVisible.value = true;
};

// 添加缺失的处理函数
// 添加一个接口定义表结构
interface TableRow {
  id: number;
  name: string;
  description?: string;
  [key: string]: any;
}

// 修改缺失的处理函数，添加类型
const handlePreview = (row: TableRow) => {
  ElMessage.info(`预览表：${row.name}`);
};

const handleSync = (row: TableRow) => {
  ElMessage.info(`同步表：${row.name}`);
};

const handleGenerate = (row: TableRow) => {
  ElMessage.info(`生成代码：${row.name}`);
};
</script>

<template>
  <div class="main">
    <el-form
      ref="formRef"
      :inline="true"
      :model="form"
      class="search-form bg-bg_color w-[99/100] pl-8 pt-[12px] overflow-auto"
    >
      <el-form-item label="表名称" prop="name">
        <el-input
          v-model="form.name"
          placeholder="请输入表名称"
          clearable
          class="!w-[150px]"
        />
      </el-form-item>
      <el-form-item label="表描述" prop="description">
        <el-input
          v-model="form.description"
          placeholder="请输入表描述"
          clearable
          class="!w-[150px]"
        />
      </el-form-item>
      <el-form-item label="创建时间" prop="loginTime">
        <el-date-picker
          v-model="form.created_at"
          :shortcuts="getPickerShortcuts()"
          type="datetimerange"
          range-separator="至"
          start-placeholder="开始日期时间"
          end-placeholder="结束日期时间"
        />
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          :icon="useRenderIcon('ri:search-line')"
          :loading="loading"
          @click="onSearch"
        >
          搜索
        </el-button>
        <el-button :icon="useRenderIcon(Refresh)" @click="resetForm(formRef)">
          重置
        </el-button>
      </el-form-item>
    </el-form>

    <PureTableBar title="代码生成配置" :columns="columns" @refresh="onSearch">
      <template #buttons>
        <el-button
          type="success"
          :icon="useRenderIcon('ep:upload')"
          @click="onImport"
        >
          导入
        </el-button>
        <el-button
          type="info"
          :icon="useRenderIcon('ep:document')"
          :disabled="!selectedNum"
          @click="onGenerate"
        >
          生成代码
        </el-button>
        <el-popconfirm title="是否确认删除选中数据?" @confirm="onbatchDel">
          <template #reference>
            <el-button
              type="danger"
              :icon="useRenderIcon('ep:delete')"
              :disabled="!selectedNum"
            >
              删除
            </el-button>
          </template>
        </el-popconfirm>
      </template>
      <template v-slot="{ size, dynamicColumns }">
        <div
          v-if="selectedNum > 0"
          v-motion-fade
          class="bg-[var(--el-fill-color-light)] w-full h-[46px] mb-2 pl-4 flex items-center"
        >
          <div class="flex-auto">
            <span
              style="font-size: var(--el-font-size-base)"
              class="text-[rgba(42,46,54,0.5)] dark:text-[rgba(220,220,242,0.5)]"
            >
              已选 {{ selectedNum }} 项
            </span>
            <el-button type="primary" text @click="onSelectionCancel">
              取消选择
            </el-button>
          </div>
          <el-popconfirm title="是否确认删除?" @confirm="onbatchDel">
            <template #reference>
              <el-button type="danger" text class="mr-1"> 批量删除 </el-button>
            </template>
          </el-popconfirm>
        </div>
        <pure-table
          ref="tableRef"
          row-key="id"
          align-whole="center"
          table-layout="auto"
          :loading="loading"
          :size="size"
          adaptive
          :adaptiveConfig="{ offsetBottom: 108 }"
          :data="dataList"
          :columns="dynamicColumns"
          :pagination="{ ...pagination, size }"
          :header-cell-style="{
            background: 'var(--el-fill-color-light)',
            color: 'var(--el-text-color-primary)'
          }"
          @selection-change="handleSelectionChange"
          @page-size-change="handleSizeChange"
          @page-current-change="handleCurrentChange"
        >
          <template #operation="{ row, size }">
            <el-button
              class="reset-margin"
              link
              type="primary"
              :size="size"
              :icon="useRenderIcon(Preview)"
              @click="handlePreview(row)"
            >
              预览
            </el-button>
            <el-button
              class="reset-margin"
              link
              type="primary"
              :size="size"
              :icon="useRenderIcon(EditPen)"
              @click="handleEdit(row)"
            >
              编辑
            </el-button>
            <el-button
              class="reset-margin"
              link
              type="primary"
              :size="size"
              :icon="useRenderIcon(Sync)"
              @click="handleSync(row)"
            >
              同步
            </el-button>
            <el-button
              class="reset-margin"
              link
              type="primary"
              :size="size"
              :icon="useRenderIcon(Code)"
              @click="handleGenerate(row)"
            >
              生成代码
            </el-button>
            <el-popconfirm
              :title="`是否确认删除表${row.name}?`"
              @confirm="handleDelete(row)"
            >
              <template #reference>
                <el-button
                  class="reset-margin"
                  link
                  type="danger"
                  :size="size"
                  :icon="useRenderIcon(Delete)"
                >
                  删除
                </el-button>
              </template>
            </el-popconfirm>
          </template>
        </pure-table>
      </template>
    </PureTableBar>

    <ImportTable v-model:visible="importVisible" />
  </div>
</template>

<style lang="scss" scoped>
:deep(.el-dropdown-menu__item i) {
  margin: 0;
}

.main-content {
  margin: 24px 24px 0 !important;
}

.search-form {
  :deep(.el-form-item) {
    margin-bottom: 12px;
  }
}
</style>
