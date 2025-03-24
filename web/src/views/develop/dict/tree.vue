<script setup lang="ts">
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import { ref, computed, watch, getCurrentInstance } from "vue";
import { dictdata } from "./utils/hook";
import More2Fill from "@iconify-icons/ri/more-2-fill";
import ExpandIcon from "./svg/expand.svg?component";
import UnExpandIcon from "./svg/unexpand.svg?component";
import { ElMessageBox } from "element-plus";
import { message } from "@/utils/message";
import { deleteDictType, getDictTypeList } from "@/api/dict";

// 修改 Tree 接口定义
interface Tree {
  id: number;
  pid: number;
  name: string;
  type: string;
  sort: number;
  remark: string;
  status: number;
  createdAt: string;
  updatedAt: string;
  highlight?: boolean;
  children?: Tree[];
}

defineProps({
  treeLoading: Boolean,
  treeData: {
    type: Array as PropType<Tree[]>,
    default: () => []
  }
});

// 修改 emit 定义
const emit = defineEmits(["tree-select", "check"]);

// 替换 nodeClick 函数为 handleCheck
// 添加选中节点数量的响应式变量
const checkedCount = ref(0);

// 修改 handleCheck 函数
function handleCheck(data, { checkedNodes, checkedKeys }) {
  checkedCount.value = checkedNodes.length;
  emit("check", { checkedNodes, checkedKeys });
}

const treeRef = ref();
const tableRef = ref();
const isExpand = ref(true);
const searchValue = ref("");
const highlightMap = ref({});
const { proxy } = getCurrentInstance();
const defaultProps = {
  children: "children",
  label: "name"
};
const buttonClass = computed(() => {
  return [
    "!h-[20px]",
    "!text-sm",
    "reset-margin",
    "!text-[var(--el-text-color-regular)]",
    "dark:!text-white",
    "dark:hover:!text-primary"
  ];
});

const filterNode = (value: string, data: Tree) => {
  if (!value) return true;
  return data.name.includes(value);
};

function nodeClick(value) {
  const nodeId = value.$treeNodeId;
  highlightMap.value[nodeId] = highlightMap.value[nodeId]?.highlight
    ? Object.assign({ id: nodeId }, highlightMap.value[nodeId], {
        highlight: false
      })
    : Object.assign({ id: nodeId }, highlightMap.value[nodeId], {
        highlight: true
      });
  Object.values(highlightMap.value).forEach((v: Tree) => {
    if (v.id !== nodeId) {
      v.highlight = false;
    }
  });
  emit(
    "tree-select",
    highlightMap.value[nodeId]?.highlight
      ? Object.assign({ ...value, selected: true })
      : Object.assign({ ...value, selected: false })
  );
}
const currentNode = ref(null);
function toggleRowExpansionAll(status) {
  isExpand.value = status;
  const nodes = (proxy.$refs["treeRef"] as any).store._getAllNodes();
  for (let i = 0; i < nodes.length; i++) {
    nodes[i].expanded = status;
  }
}

function handleEdit() {
  const checkedNodes = treeRef.value.getCheckedNodes();
  if (!checkedNodes || checkedNodes.length === 0) {
    message("请选择要修改的字典类型", { type: "warning" });
    return;
  }
  // 获取第一个选中的节点
  const firstNode = checkedNodes[0];
  openTypeDialog("修改", firstNode);
}

function handleDelete() {
  // 获取选中的节点
  const checkedNodes = treeRef.value.getCheckedNodes();
  if (!checkedNodes || checkedNodes.length === 0) {
    message("请选择要删除的字典类型", { type: "warning" });
    return;
  }

  ElMessageBox.confirm(
    `确认要删除选中的 ${checkedNodes.length} 个字典类型吗？`,
    "系统提示",
    {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
      type: "warning",
      draggable: true
    }
  ).then(async () => {
    try {
      // 批量删除
      const promises = checkedNodes.map(node => deleteDictType(node.id));
      const results = await Promise.all(promises);
      // 修改解构，添加 searchTreeData
      // 在 handleDelete 成功后调用
      if (results.every(res => res.code === 0)) {
        message("删除成功", { type: "success" });
        onTreeReset();
        await searchTreeData(); // 添加这一行
        checkedCount.value = 0;
        currentNode.value = null;
      } else {
        message("部分删除失败，请重试", { type: "error" });
      }
    } catch (error) {
      message("删除失败", { type: "error" });
    }
  });
}

/** 重置状态（选中状态、搜索框值、树初始化） */
function onTreeReset() {
  highlightMap.value = {};
  searchValue.value = "";
  toggleRowExpansionAll(true);
}

watch(searchValue, val => {
  treeRef.value!.filter(val);
});

defineExpose({ onTreeReset });

const {
  openTypeDialog,
  treeData: dictTreeData,
  searchTreeData
} = dictdata(tableRef, treeRef);
</script>

<template>
  <div
    v-loading="treeLoading"
    class="h-full bg-bg_color overflow-hidden relative"
    :style="{ minHeight: `calc(100vh - 141px)` }"
  >
    <div class="flex flex-col gap-2">
      <div class="flex items-center justify-between px-4 pt-2">
        <el-button
          type="primary"
          size="small"
          :icon="useRenderIcon('ep:plus')"
          @click="openTypeDialog()"
        >
          新增
        </el-button>
        <el-button
          type="primary"
          size="small"
          :icon="useRenderIcon('ep:edit')"
          :disabled="checkedCount === 0"
          @click="handleEdit"
        >
          编辑
        </el-button>
        <el-button
          type="danger"
          size="small"
          :icon="useRenderIcon('ep:delete')"
          :disabled="checkedCount === 0"
          @click="handleDelete"
        >
          删除
        </el-button>
      </div>
      <div class="flex items-center h-[34px]">
        <el-input
          v-model="searchValue"
          class="ml-2"
          size="small"
          placeholder="请输入字典名称"
          clearable
        >
          <template #suffix>
            <el-icon class="el-input__icon">
              <IconifyIconOffline
                v-show="searchValue.length === 0"
                icon="ri:search-line"
              />
            </el-icon>
          </template>
        </el-input>
        <el-dropdown :hide-on-click="false">
          <IconifyIconOffline
            class="w-[28px] cursor-pointer"
            width="18px"
            :icon="More2Fill"
          />
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item>
                <el-button
                  :class="buttonClass"
                  link
                  type="primary"
                  :icon="useRenderIcon(isExpand ? ExpandIcon : UnExpandIcon)"
                  @click="toggleRowExpansionAll(isExpand ? false : true)"
                >
                  {{ isExpand ? "折叠全部" : "展开全部" }}
                </el-button>
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </div>
    <el-divider />
    <el-scrollbar height="calc(90vh - 88px)">
      <el-tree
        ref="treeRef"
        :data="dictTreeData"
        node-key="id"
        size="small"
        :props="defaultProps"
        default-expand-all
        show-checkbox
        :expand-on-click-node="false"
        :filter-node-method="filterNode"
        @check="handleCheck"
        @node-click="nodeClick"
      >
        <template #default="{ node }">
          <div
            :class="[
              'rounded',
              'flex',
              'items-center',
              'select-none',
              'hover:text-primary',
              searchValue.trim().length > 0 &&
                node.label.includes(searchValue) &&
                'text-red-500',
              highlightMap[node.id]?.highlight ? 'dark:text-primary' : ''
            ]"
            :style="{
              color: highlightMap[node.id]?.highlight
                ? 'var(--el-color-primary)'
                : '',
              background: highlightMap[node.id]?.highlight
                ? 'var(--el-color-primary-light-7)'
                : 'transparent'
            }"
          >
            <span class="!w-[120px] !truncate" :title="node.label">
              {{ node.label }}
            </span>
          </div>
        </template>
      </el-tree>
    </el-scrollbar>
  </div>
</template>

<style lang="scss" scoped>
:deep(.el-divider) {
  margin: 0;
}

:deep(.el-tree) {
  --el-tree-node-hover-bg-color: transparent;
}
</style>
