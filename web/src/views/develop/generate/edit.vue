<script setup lang="ts">
import { ref } from "vue";
import GenerateInfo from "./gen.vue";
import FieldInfo from "./field.vue";
const active = ref(0);
const loading = ref(false);

const basicForm = ref({
  tableName: "demo_gen",
  tableDesc: "代码生成测试表",
  className: "DemoGen",
  author: "panda",
  sortField: "ID",
  sortType: "顺序",
  remark: "",
  hasFile: false,
  hasExport: true,
  hasImport: false,
  hasDetail: true,
  hasAdd: true,
  hasEdit: true,
  hasDelete: true,
  hasQuery: true
});

const moduleList = ref([
  { label: "覆盖api", value: "api" },
  { label: "覆盖controller", value: "controller" },
  { label: "覆盖dao", value: "dao" },
  { label: "覆盖dao_internal", value: "dao_internal" },
  { label: "覆盖logic", value: "logic" },
  { label: "覆盖model", value: "model" },
  { label: "覆盖model_do", value: "model_do" },
  { label: "覆盖model_entity", value: "model_entity" },
  { label: "覆盖router", value: "router" },
  { label: "覆盖router_func", value: "router_func" },
  { label: "覆盖service", value: "service" },
  { label: "覆盖类型sql", value: "sql" },
  { label: "覆盖tsApi", value: "tsApi" },
  { label: "覆盖tsModel", value: "tsModel" },
  { label: "覆盖vue-list", value: "vue-list" },
  { label: "覆盖vue-detail", value: "vue-detail" },
  { label: "覆盖vue-edit", value: "vue-edit" }
]);

const selectedModules = ref<string[]>([]);

const rules = {
  tableName: [{ required: true, message: "请输入表名称", trigger: "blur" }],
  tableDesc: [{ required: true, message: "请输入表描述", trigger: "blur" }],
  className: [{ required: true, message: "请输入实体类名称", trigger: "blur" }],
  author: [{ required: true, message: "请输入作者", trigger: "blur" }]
};

const generateInfoRef = ref();

const next = async () => {
  if (active.value === 0) {
    // 验证基本信息
    // TODO: 表单验证
  } else if (active.value === 1) {
    // 验证生成信息
    // TODO: 表单验证
  }
  if (active.value < 2) {
    active.value++;
  }
};

const prev = () => {
  if (active.value > 0) {
    active.value--;
  }
};

const close = () => {
  // TODO: 关闭页面
};

const fieldInfoRef = ref();
</script>

<template>
  <div class="main p-6">
    <el-steps :active="active + 1" finish-status="success" class="mb-4">
      <el-step title="1.基本信息" />
      <el-step title="2.生成信息" />
      <el-step title="3.字段信息" />
    </el-steps>

    <div v-show="active === 0" class="bg-white p-6 rounded-lg shadow-sm">
      <div class="text-lg font-bold mb-4">基本信息配置</div>
      <el-form
        ref="basicFormRef"
        :model="basicForm"
        :rules="rules"
        label-width="120px"
      >
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="表名称" prop="tableName">
              <el-input v-model="basicForm.tableName" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="表描述" prop="tableDesc">
              <el-input v-model="basicForm.tableDesc" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="实体类名称" prop="className">
              <el-input v-model="basicForm.className" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="作者" prop="author">
              <el-input v-model="basicForm.author" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="排序字段" prop="sortField">
              <el-select v-model="basicForm.sortField" class="w-full">
                <el-option label="ID" value="ID" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="排序方式" prop="sortType">
              <el-select v-model="basicForm.sortType" class="w-full">
                <el-option label="顺序" value="顺序" />
                <el-option label="倒序" value="倒序" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="备注" prop="remark">
              <el-input v-model="basicForm.remark" type="textarea" :rows="3" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item>
          <el-checkbox v-model="basicForm.hasFile">
            是否覆盖现有文件
          </el-checkbox>
          <el-checkbox v-model="basicForm.hasAdd" class="ml-4">
            新增功能
          </el-checkbox>
          <el-checkbox v-model="basicForm.hasEdit" class="ml-4">
            修改功能
          </el-checkbox>
          <el-checkbox v-model="basicForm.hasDelete" class="ml-4">
            删除功能
          </el-checkbox>
          <el-checkbox v-model="basicForm.hasQuery" class="ml-4">
            查询功能
          </el-checkbox>
          <el-checkbox v-model="basicForm.hasExport" class="ml-4">
            导出excel
          </el-checkbox>
          <el-checkbox v-model="basicForm.hasImport" class="ml-4">
            导入excel
          </el-checkbox>
          <el-checkbox v-model="basicForm.hasDetail" class="ml-4">
            是否显示详情功能
          </el-checkbox>
        </el-form-item>
      </el-form>

      <div class="mt-4">
        <el-divider content-position="left">模块选择</el-divider>
        <el-checkbox-group v-model="selectedModules" class="w-full">
          <el-row :gutter="20">
            <el-col
              v-for="item in moduleList"
              :key="item.value"
              :span="6"
              class="mb-4"
            >
              <el-checkbox :value="item.value">
                {{ item.label }}
              </el-checkbox>
            </el-col>
          </el-row>
        </el-checkbox-group>
      </div>
    </div>

    <div v-show="active === 1">
      <GenerateInfo ref="generateInfoRef" />
    </div>

    <div v-show="active === 2">
      <FieldInfo ref="fieldInfoRef" />
    </div>

    <div class="flex justify-center mt-8">
      <el-button @click="close">关闭</el-button>
      <el-button v-if="active > 0" @click="prev">上一步</el-button>
      <el-button type="primary" @click="next">
        {{ active < 2 ? "下一步" : "保存" }}
      </el-button>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.main {
  background-color: var(--el-bg-color);
  border-radius: 8px;
  min-height: calc(100vh - 100px);

  :deep(.el-step__title) {
    font-size: 16px;
  }

  :deep(.el-form-item__label) {
    font-weight: 500;
  }

  :deep(.el-checkbox.is-bordered) {
    padding: 8px 16px;
    border-radius: 4px;
    margin-right: 0;
    height: auto;

    &:hover {
      border-color: var(--el-color-primary);
    }
  }

  :deep(.el-divider__text) {
    font-size: 15px;
    font-weight: 500;
    color: var(--el-text-color-secondary);
  }
}
</style>
