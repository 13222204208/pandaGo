<script setup lang="ts">
import { ref, watch } from "vue";
import { ElMessage } from "element-plus";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import { getTableList, type TableItem } from "@/api/generate";

interface Props {
  visible: boolean;
}

const props = defineProps<Props>();
const emit = defineEmits(["update:visible"]);

const formRef = ref();
const loading = ref(false);
const tableData = ref<TableItem[]>([]);
const selectedTables = ref([]);

const columns = [
  {
    type: "selection",
    width: 55,
    align: "center",
    fixed: "left"
  },
  {
    label: "表名称",
    prop: "tableName",
    minWidth: 120
  },
  {
    label: "表描述",
    prop: "tableComment",
    minWidth: 120
  },
  {
    label: "创建时间",
    prop: "createTime",
    minWidth: 180
  },
  {
    label: "更新时间",
    prop: "updateTime",
    minWidth: 180
  }
];

const pagination = ref({
  total: 0,
  pageSize: 10,
  currentPage: 1,
  background: true
});

// 获取数据
const fetchData = async (
  currentPage = pagination.value.currentPage,
  pageSize = pagination.value.pageSize
) => {
  loading.value = true;
  try {
    const { data } = await getTableList({
      currentPage,
      pageSize
    });
    tableData.value = data.list;
    pagination.value.total = data.total;
    pagination.value.currentPage = data.currentPage;
    pagination.value.pageSize = data.pageSize;
  } catch (error) {
    ElMessage.error("获取数据库表列表失败");
  } finally {
    loading.value = false;
  }
};

// 添加分页变化处理
const handleCurrentChange = (currentPage: number) => {
  fetchData(currentPage, pagination.value.pageSize);
};

const handleSizeChange = (pageSize: number) => {
  fetchData(1, pageSize); // 切换每页条数时，重置为第一页
};

// 选择变化
const handleSelectionChange = selection => {
  selectedTables.value = selection;
};

// 确定导入
const onSubmit = async () => {
  if (selectedTables.value.length === 0) {
    ElMessage.warning("请选择要导入的表");
    return;
  }
  loading.value = true;
  try {
    // TODO: 调用导入表接口
    setTimeout(() => {
      ElMessage.success("导入成功");
      handleClose();
    }, 1000);
  } catch (error) {
    ElMessage.error("导入失败");
  } finally {
    loading.value = false;
  }
};

const handleClose = () => {
  selectedTables.value = [];
  emit("update:visible", false);
};

// 监听对话框可见性变化，当显示时才加载数据
watch(
  () => props.visible,
  newVal => {
    if (newVal) {
      fetchData();
    }
  }
);
</script>

<template>
  <el-dialog
    :model-value="props.visible"
    title="导入表"
    width="900px"
    :close-on-click-modal="false"
    @update:model-value="emit('update:visible', $event)"
    @close="handleClose"
  >
    <div
      v-if="selectedTables.length > 0"
      class="mb-2 pl-4 flex items-center bg-[var(--el-fill-color-light)] h-[46px]"
    >
      <div class="flex-auto">
        <span
          class="text-[rgba(42,46,54,0.5)] dark:text-[rgba(220,220,242,0.5)]"
        >
          已选 {{ selectedTables.length }} 项
        </span>
        <el-button type="primary" text @click="selectedTables = []">
          取消选择
        </el-button>
      </div>
    </div>

    <pure-table
      row-key="name"
      align-whole="center"
      table-layout="auto"
      :loading="loading"
      :data="tableData"
      :columns="columns"
      :pagination="pagination"
      :header-cell-style="{
        background: 'var(--el-fill-color-light)',
        color: 'var(--el-text-color-primary)'
      }"
      @selection-change="handleSelectionChange"
      @page-size-change="handleSizeChange"
      @page-current-change="handleCurrentChange"
    />

    <template #footer>
      <div class="flex items-center justify-between">
        <span class="text-[var(--el-text-color-secondary)]">
          共 {{ pagination.total }} 条数据
        </span>
        <div>
          <el-button @click="handleClose">取 消</el-button>
          <el-button type="primary" :loading="loading" @click="onSubmit">
            确 定
          </el-button>
        </div>
      </div>
    </template>
  </el-dialog>
</template>

<style lang="scss" scoped></style>
