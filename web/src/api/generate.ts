import { http } from "@/utils/http";
import { baseUrlApi } from "./utils";
import type { ListResponse, ApiResponse } from "./utils";

// 数据表接口
export interface TableItem {
  tableName: string;
  tableComment: string;
  createTime: string;
  updateTime: string;
}

// sql 生成接口
export interface SqlGenerateItem {
  sql: string;
}

// 获取数据表列表
export const getTableList = (params?: object) => {
  return http.request<ListResponse<TableItem>>("get", baseUrlApi("table"), {
    params
  });
};

// 根据描述生成 SQL
export const generateSQL = (data?: object) => {
  return http.request<ApiResponse<SqlGenerateItem>>(
    "post",
    baseUrlApi("generate/sql"),
    {
      data
    }
  );
};

// 执行 SQL 语句
export const executeSQL = (sql: string) => {
  return http.request<ApiResponse>("post", baseUrlApi("generate/execute"), {
    data: { sql }
  });
};
