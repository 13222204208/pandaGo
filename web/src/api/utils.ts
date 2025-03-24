export const baseUrlApi = (url: string) => `/api/admin/${url}`;

// 基础响应接口
export interface BaseResponse {
  code: number;
  message: string;
}

// 通用API响应（data可能为空）
export type ApiResponse<T = any> = BaseResponse & {
  data: T;
};

// 列表数据响应
export interface ListData<T = any> {
  list: T[];
  total: number;
  pageSize: number;
  currentPage: number;
}

// 列表响应类型
export type ListResponse<T = any> = BaseResponse & {
  data: ListData<T>;
};
