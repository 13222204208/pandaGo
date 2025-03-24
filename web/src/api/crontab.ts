import { http } from "@/utils/http";
import { baseUrlApi } from "./utils";
import type { ApiResponse, ListResponse } from "./utils";

// 定时任务参数类型
export interface CrontabCreateParams {
  title: string;
  name: string;
  params?: string;
  pattern: string;
  policy: number;
  sort?: number;
  status: number;
  remark?: string;
}

// 定时任务信息类型
export interface CrontabInfo {
  id: number;
  title: string;
  name: string;
  params: string;
  pattern: string;
  policy: number;
  sort: number;
  status: number;
  remark: string;
  createdAt: string;
  updatedAt: string;
}

// 创建定时任务
export const createCrontab = (data: CrontabCreateParams) => {
  return http.request<ApiResponse>("post", baseUrlApi("crontab"), { data });
};

// 获取定时任务列表
export const getCrontabList = (params?: object) => {
  return http.request<ListResponse<CrontabInfo>>("get", baseUrlApi("crontab"), {
    params
  });
};

// 更新定时任务
export const updateCrontab = (
  id: number,
  data: Partial<CrontabCreateParams>
) => {
  return http.request<ApiResponse>("put", baseUrlApi(`crontab/${id}`), {
    data
  });
};

// 删除定时任务
export const deleteCrontab = (id: number) => {
  return http.request<ApiResponse>("delete", baseUrlApi(`crontab/${id}`));
};

// 修改定时任务状态
export const updateCrontabStatus = (id: number, status: number) => {
  return http.request<ApiResponse>("put", baseUrlApi(`crontab/${id}/status`), {
    data: { status }
  });
};
