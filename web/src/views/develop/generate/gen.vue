<script setup lang="ts">
import { ref } from "vue";

const form = ref({
  packageName: "internal/app/demo", // 包名
  moduleName: "demo",              // 模块名
  businessName: "demo",           // 业务名
  frontendPath: "src/views/demo",  // 前端代码路径
  backendPath: "internal/app/demo", // 后端代码路径
  apiPrefix: "/api/v1/demo",      // 接口前缀
});

const rules = {
  packageName: [{ required: true, message: "请输入包名", trigger: "blur" }],
  moduleName: [{ required: true, message: "请输入模块名", trigger: "blur" }],
  businessName: [{ required: true, message: "请输入业务名", trigger: "blur" }],
  frontendPath: [{ required: true, message: "请输入前端代码路径", trigger: "blur" }],
  backendPath: [{ required: true, message: "请输入后端代码路径", trigger: "blur" }],
  apiPrefix: [{ required: true, message: "请输入接口前缀", trigger: "blur" }]
};

defineExpose({
  form,
  rules
});
</script>

<template>
  <div class="bg-white p-6 rounded-lg shadow-sm">
    <div class="text-lg font-bold mb-4">生成信息配置</div>
    <el-form ref="formRef" :model="form" :rules="rules" label-width="120px">
      <el-row :gutter="20">
        <el-col :span="12">
          <el-form-item label="包名" prop="packageName">
            <el-input v-model="form.packageName" placeholder="例如: internal/app/demo" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="模块名" prop="moduleName">
            <el-input v-model="form.moduleName" placeholder="例如: demo" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="业务名" prop="businessName">
            <el-input v-model="form.businessName" placeholder="例如: demo" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="前端代码路径" prop="frontendPath">
            <el-input v-model="form.frontendPath" placeholder="例如: src/views/demo" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="后端代码路径" prop="backendPath">
            <el-input v-model="form.backendPath" placeholder="例如: internal/app/demo" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="接口前缀" prop="apiPrefix">
            <el-input v-model="form.apiPrefix" placeholder="例如: /api/v1/demo" />
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
  </div>
</template>

<style scoped>
.el-form :deep(.el-form-item__label) {
  font-weight: 500;
}
</style>