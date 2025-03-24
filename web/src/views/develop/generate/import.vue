<script setup lang="ts">
import { ref } from "vue";
import { ElMessage } from "element-plus";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";

interface Props {
  visible: boolean;
}

const props = defineProps<Props>();
const emit = defineEmits(["update:visible"]);

const formRef = ref();
const loading = ref(false);
const tableData = ref([]);
const selectedTables = ref([]);
const searchForm = ref({
  name: "",
  description: ""
});

const columns = [
  {
    type: "selection",
    width: 55,
    align: "center",
    fixed: "left"
  },
  {
    label: "表名称",
    prop: "name",
    minWidth: 120
  },
  {
    label: "表描述",
    prop: "description",
    minWidth: 120
  },
  {
    label: "创建时间",
    prop: "created_at",
    minWidth: 180
  },
  {
    label: "更新时间",
    prop: "updated_at",
    minWidth: 180
  }
];

const pagination = ref({
  total: 0,
  pageSize: 10,
  currentPage: 1,
  background: true
});

// 搜索
const onSearch = async () => {
  loading.value = true;
  try {
    // TODO: 调用获取数据库表列表接口
    tableData.value = [
      {
        name: "wx_material",
        description: "微信永久素材库",
        created_at: "2024-12-23 08:52:43",
        updated_at: "2024-12-23 08:52:43"
      }
      // ... 其他数据
    ];
  } catch (error) {
    ElMessage.error("获取数据库表列表失败");
  } finally {
    loading.value = false;
  }
};

// 重置
const onReset = () => {
  searchForm.value = {
    name: "",
    description: ""
  };
  onSearch();
};

// 选择变化
const handleSelectionChange = (selection) => {
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
  searchForm.value = {
    name: "",
    description: ""
  };
  selectedTables.value = [];
  emit("update:visible", false);
};

onSearch();
</script>

<template>
  <el-dialog
    title="导入表"
    v-model="props.visible"
    width="900px"
    :close-on-click-modal="false"
    @close="handleClose"
  >
    <div class="search-form bg-bg_color w-[99/100] pl-8 pt-[12px] overflow-auto">
      <el-form ref="formRef" :inline="true" :model="searchForm">
        <el-form-item label="表名称" prop="name">
          <el-input
            v-model="searchForm.name"
            placeholder="请输入表名称"
            clearable
            class="!w-[150px]"
          />
        </el-form-item>
        <el-form-item label="表描述" prop="description">
          <el-input
            v-model="searchForm.description"
            placeholder="请输入表描述"
            clearable
            class="!w-[150px]"
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
          <el-button
            :icon="useRenderIcon('ep:refresh')"
            @click="onReset"
          >
            重置
          </el-button>
        </el-form-item>
      </el-form>
    </div>

    <div v-if="selectedTables.length > 0" class="mb-2 pl-4 flex items-center bg-[var(--el-fill-color-light)] h-[46px]">
      <div class="flex-auto">
        <span class="text-[rgba(42,46,54,0.5)] dark:text-[rgba(220,220,242,0.5)]">
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

<style lang="scss" scoped>
.search-form {
  :deep(.el-form-item) {
    margin-bottom: 12px;
  }
}
</style>
