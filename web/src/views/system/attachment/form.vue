<script setup lang="ts">
import { reactive, ref } from "vue";
import { message } from "@/utils/message";
import { createFormData } from "@pureadmin/utils";
import { uploadAttachment } from "@/api/attachment";

import UploadIcon from "@iconify-icons/ri/upload-2-line";

// 定义事件
const emit = defineEmits(["success", "close"]);

const formRef = ref();
const uploadRef = ref();
const validateForm = reactive({
  fileList: [],
  date: ""
});

const submitForm = formEl => {
  if (!formEl) return;
  formEl.validate(async valid => {
    if (valid) {
      const formData = createFormData({
        files: validateForm.fileList.map(file => ({ raw: file.raw })),
        date: validateForm.date
      });
      const res = await uploadAttachment(formData);
      if (res.code === 0) {
        message("文件上传成功", { type: "success" });
        // 触发成功事件，通知父组件关闭弹窗和刷新数据
        emit("success");
        emit("close");
      } else {
        message(res.message, { type: "error" });
      }
    } else {
      return false;
    }
  });
};

const resetForm = formEl => {
  if (!formEl) return;
  formEl.resetFields();
};
</script>

<template>
  <el-form ref="formRef" :model="validateForm" label-width="82px">
    <el-form-item
      label="附件"
      prop="fileList"
      :rules="[{ required: true, message: '附件不能为空' }]"
    >
      <el-upload
        ref="uploadRef"
        v-model:file-list="validateForm.fileList"
        drag
        multiple
        action="#"
        class="!w-[200px]"
        :auto-upload="false"
      >
        <div class="el-upload__text">
          <IconifyIconOffline
            :icon="UploadIcon"
            width="26"
            class="m-auto mb-2"
          />
          可点击或拖拽上传
        </div>
      </el-upload>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" text bg @click="submitForm(formRef)">
        提交
      </el-button>
      <el-button text bg @click="resetForm(formRef)">重置</el-button>
    </el-form-item>
  </el-form>
</template>
