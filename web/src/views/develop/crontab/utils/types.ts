// 定时任务相关类型定义

interface FormItemProps {
  /** ID */
  id?: number;
  /** 任务标题 */
  title: string;
  /** 任务名称 */
  name: string;
  /** 任务参数 */
  params?: string;
  /** Cron表达式 */
  pattern: string;
  /** 执行策略 */
  policy: number;
  /** 排序 */
  sort?: number;
  /** 状态 */
  status: number;
  /** 备注 */
  remark?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
}

interface FormProps {
  formInline: FormItemProps;
}

export type { FormItemProps, FormProps };
