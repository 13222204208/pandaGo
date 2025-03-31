<script setup lang="ts">
// 添加导入
import { ref } from "vue";
import { ElMessage } from "element-plus";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import { generateSQL, executeSQL } from "@/api/generate";

interface Props {
  visible: boolean;
}

const props = defineProps<Props>();
const emit = defineEmits(["update:visible"]);

const loading = ref(false);
const generating = ref(false); // 添加生成状态
const prompt = ref("");
const sqlResult = ref("");

// 生成 SQL
const onGenerate = async () => {
  if (!prompt.value) {
    ElMessage.warning("请输入要求描述");
    return;
  }
  loading.value = true;
  generating.value = true;
  try {
    sqlResult.value = "";
    const dots = [".", "..", "..."];
    let i = 0;
    const interval = setInterval(() => {
      sqlResult.value = `正在生成${dots[i]}`;
      i = (i + 1) % dots.length;
    }, 1000);

    try {
      // 调用生成 SQL 接口
      const response = await generateSQL({ prompt: prompt.value });
      clearInterval(interval);
      if (response && response.data && response.data.sql) {
        sqlResult.value = response.data.sql;
        ElMessage.success("生成成功");
      } else {
        sqlResult.value = "";
        ElMessage.warning("生成SQL失败，请检查描述或稍后重试");
      }
    } catch (error: any) {
      clearInterval(interval);

      // 处理超时错误
      if (error.code === "ECONNABORTED" || error.message?.includes("timeout")) {
        sqlResult.value = "";
        ElMessage.error("请求超时，AI生成服务可能繁忙，请稍后重试");
      } else {
        sqlResult.value = "";
        ElMessage.error("生成失败，请稍后重试");
      }
      console.error("生成SQL错误:", error);
    }
  } catch (error) {
    ElMessage.error("生成失败");
  } finally {
    loading.value = false;
    generating.value = false;
  }
};

// 添加执行 SQL 函数
const onExecute = async () => {
  if (!sqlResult.value) {
    ElMessage.warning("请先生成 SQL");
    return;
  }
  loading.value = true;
  try {
    // 调用执行 SQL 接口
    await executeSQL(sqlResult.value);
    ElMessage.success("执行成功");
    handleClose();
  } catch (error) {
    ElMessage.error("执行失败");
  } finally {
    loading.value = false;
  }
};

const handleClose = () => {
  prompt.value = "";
  sqlResult.value = "";
  emit("update:visible", false);
};
</script>

<template>
  <el-dialog
    :model-value="props.visible"
    title="SQL 预览"
    width="800px"
    :close-on-click-modal="false"
    @update:model-value="emit('update:visible', $event)"
    @close="handleClose"
  >
    <el-form>
      <el-form-item label="需求描述">
        <el-input
          v-model="prompt"
          type="textarea"
          :rows="3"
          placeholder="请输入创建表的需求描述，例如：创建一个用户表，包含用户名、邮箱等字段"
        />
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          :loading="loading"
          :icon="useRenderIcon('ep:magic-stick')"
          @click="onGenerate"
        >
          {{ loading ? "AI 生成中..." : "生成 SQL" }}
        </el-button>
      </el-form-item>
    </el-form>

    <el-form-item v-if="sqlResult || generating" label="SQL 语句">
      <div
        class="sql-editor-container"
        :class="{ 'is-generating': generating }"
      >
        <el-input
          v-model="sqlResult"
          type="textarea"
          :rows="10"
          class="sql-editor"
          resize="none"
          spellcheck="false"
          :readonly="generating"
        />
      </div>
    </el-form-item>

    <template #footer>
      <div class="flex items-center justify-between">
        <span class="text-[var(--el-text-color-secondary)]">
          提示：SQL 语句可以手动修改
        </span>
        <div class="flex gap-2">
          <el-button @click="handleClose">取 消</el-button>
          <el-button
            type="primary"
            :loading="loading"
            :icon="useRenderIcon('ep:position')"
            @click="onExecute"
          >
            执行 SQL
          </el-button>
        </div>
      </div>
    </template>
  </el-dialog>
</template>

<style lang="scss" scoped>
:deep(.el-form-item__content) {
  flex: 1;
}

.sql-editor-container {
  width: 100%;
  height: 300px;
  border-radius: 4px;
  overflow: hidden;
  border: 1px solid var(--el-border-color);

  .sql-editor {
    height: 100%;

    :deep(.el-textarea__inner) {
      height: 100%;
      font-family: "Menlo", "Monaco", "Consolas", "Courier New", monospace;
      font-size: 14px;
      line-height: 1.6;
      padding: 12px;
      background-color: var(--el-fill-color-light);
      color: var(--el-text-color-primary);
      border: none;
      border-radius: 0;
      resize: none;
      box-shadow: none;

      &:focus {
        outline: none;
        box-shadow: none;
      }
    }
  }
}
</style>
