import { http } from "@/utils/http";
import { baseUrlApi, type ApiResponse, type ListResponse } from "./utils";

export interface AttachmentItem {
  id: number;
  type: string;
  name: string;
  path: string;
  size: string;
  ext: string;
  drive: number;
  status: number;
  createdAt: string;
  updatedAt: string;
}

// 获取附件列表
export const getAttachmentList = (params?: object) => {
  return http.request<ListResponse<AttachmentItem>>(
    "get",
    baseUrlApi("attachment"),
    {
      params
    }
  );
};

// 上传附件
export const uploadAttachment = (data: FormData) => {
  return http.request<ApiResponse<AttachmentItem>>(
    "post",
    baseUrlApi("attachment/upload"),
    {
      data,
      headers: {
        "Content-Type": "multipart/form-data"
      }
    }
  );
};

// 删除附件
export const deleteAttachment = (id: number) => {
  return http.request<ApiResponse>("delete", baseUrlApi(`attachment/${id}`));
};

// 修改附件信息
export const updateAttachment = (id: number, data: Partial<AttachmentItem>) => {
  return http.request<ApiResponse<AttachmentItem>>(
    "put",
    baseUrlApi(`attachment/${id}`),
    {
      data
    }
  );
};

// 获取附件详情
export const getAttachmentDetail = (id: number) => {
  return http.request<ApiResponse<AttachmentItem>>(
    "get",
    baseUrlApi(`attachment/${id}`)
  );
};
