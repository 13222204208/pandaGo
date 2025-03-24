import dayjs from "dayjs";
import editForm from "../form.vue";
import { message } from "@/utils/message";
import { ElMessageBox } from "element-plus";
import { usePublicHooks } from "../../hooks";
import { transformI18n } from "@/plugins/i18n";
import { addDialog } from "@/components/ReDialog";
import type { FormItemProps } from "./types";
import type { PaginationProps } from "@pureadmin/table";
import { deviceDetection } from "@pureadmin/utils";
import { type Ref, reactive, ref, onMounted, h, toRaw } from "vue";
import {
  getCrontabList,
  createCrontab,
  updateCrontab,
  deleteCrontab,
  updateCrontabStatus
} from "@/api/crontab";

export function useColumns(tableRef: Ref) {
  const form = reactive({
    title: "",
    name: "",
    params: "",
    pattern: "",
    policy: null,
    status: null,
    remark: ""
  });
  const curRow = ref();
  const formRef = ref();
  const dataList = ref([]);
  const isShow = ref(false);
  const loading = ref(true);
  const switchLoadMap = ref({});
  const { switchStyle } = usePublicHooks();
  const pagination = reactive<PaginationProps>({
    total: 0,
    pageSize: 10,
    currentPage: 1,
    background: true
  });
  const columns: TableColumnList = [
    {
      label: "任务标题",
      prop: "title"
    },
    {
      label: "执行方法",
      prop: "name"
    },
    {
      label: "任务参数",
      prop: "params",
      minWidth: 120
    },
    {
      label: "Cron表达式",
      prop: "pattern",
      minWidth: 120
    },
    {
      label: "执行策略",
      prop: "policy",
      formatter: ({ policy }) => {
        const policyMap = {
          1: "并行策略",
          2: "单例策略",
          3: "单次策略",
          4: "多次策略"
        };
        return policyMap[policy] || "未知策略";
      }
    },
    {
      label: "状态",
      cellRenderer: scope => (
        <el-switch
          size={scope.props.size === "small" ? "small" : "default"}
          loading={switchLoadMap.value[scope.index]?.loading}
          v-model={scope.row.status}
          active-value={1}
          inactive-value={0}
          active-text="已启用"
          inactive-text="已停用"
          inline-prompt
          style={switchStyle.value}
          onChange={() => onChange(scope as any)}
        />
      ),
      minWidth: 80
    },
    {
      label: "备注",
      prop: "remark",
      minWidth: 160
    },
    {
      label: "创建时间",
      prop: "createdAt",
      minWidth: 100,
      formatter: ({ createdAt }) =>
        dayjs(createdAt).format("YYYY-MM-DD HH:mm:ss")
    },
    {
      label: "操作",
      fixed: "right",
      width: 270,
      slot: "operation"
    }
  ];

  function onChange({ row, index }) {
    ElMessageBox.confirm(
      `确认要<strong>${
        row.status === 0 ? "停用" : "启用"
      }</strong><strong style='color:var(--el-color-primary)'>${
        row.name
      }</strong>吗?`,
      "系统提示",
      {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
        dangerouslyUseHTMLString: true,
        draggable: true
      }
    )
      .then(() => {
        switchLoadMap.value[index] = Object.assign(
          {},
          switchLoadMap.value[index],
          {
            loading: true
          }
        );
        // 调用更新状态API
        updateCrontabStatus(row.id, row.status)
          .then(res => {
            switchLoadMap.value[index].loading = false;
            if (res.code === 0) {
              message(`已${row.status === 0 ? "停用" : "启用"}${row.name}`, {
                type: "success"
              });
            } else {
              message(res.message, { type: "error" });
              row.status = row.status === 0 ? 1 : 0;
            }
          })
          .catch(error => {
            switchLoadMap.value[index].loading = false;
            message(error.message || "操作失败", { type: "error" });
            row.status = row.status === 0 ? 1 : 0;
          });
      })
      .catch(() => {
        row.status = row.status === 0 ? 1 : 0;
      });
  }

  function handleDelete(row) {
    // 调用删除API
    deleteCrontab(row.id)
      .then(res => {
        if (res.code === 0) {
          message(`删除成功`, { type: "success" });
          onSearch();
        } else {
          message(res.message, { type: "error" });
        }
      })
      .catch(error => {
        message(error.message || "删除失败", { type: "error" });
      });
  }

  function handleSizeChange(val: number) {
    console.log(`${val} items per page`);
  }

  function handleCurrentChange(val: number) {
    console.log(`current page: ${val}`);
  }

  function handleSelectionChange(val) {
    console.log("handleSelectionChange", val);
  }

  async function onSearch() {
    loading.value = true;
    try {
      const { data } = await getCrontabList(toRaw(form));
      dataList.value = data.list;
      pagination.total = data.total;
      pagination.pageSize = data.pageSize;
      pagination.currentPage = data.currentPage;
    } catch (error) {
      message(error.message || "获取数据失败", { type: "error" });
      dataList.value = [];
      pagination.total = 0;
    } finally {
      loading.value = false;
    }
  }

  const resetForm = formEl => {
    if (!formEl) return;
    formEl.resetFields();
    onSearch();
  };

  function openDialog(title = "新增", row?: FormItemProps) {
    addDialog({
      title: `${title}任务`,
      props: {
        formInline: {
          id: row?.id,
          title: row?.title ?? "",
          name: row?.name ?? "",
          params: row?.params ?? "",
          pattern: row?.pattern ?? "",
          policy: row?.policy ?? 1,
          sort: row?.sort ?? 0,
          status: row?.status ?? 1,
          remark: row?.remark ?? "",
          createdAt: row?.createdAt,
          updatedAt: row?.updatedAt
        }
      },
      width: "40%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(editForm, { ref: formRef, formInline: null }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = options.props.formInline as FormItemProps;

        FormRef.validate(valid => {
          if (valid) {
            if (title === "新增") {
              // 调用新增接口
              createCrontab(curData)
                .then(res => {
                  if (res.code === 0) {
                    message(`新增任务成功`, { type: "success" });
                    done(); // 关闭弹框
                    onSearch(); // 刷新表格数据
                  } else {
                    message(res.message, { type: "error" });
                  }
                })
                .catch(error => {
                  message(error.message || "新增失败", { type: "error" });
                });
            } else {
              // 调用修改接口
              updateCrontab(curData.id, curData)
                .then(res => {
                  if (res.code === 0) {
                    message(`修改任务成功`, { type: "success" });
                    done(); // 关闭弹框
                    onSearch(); // 刷新表格数据
                  } else {
                    message(res.message, { type: "error" });
                  }
                })
                .catch(error => {
                  message(error.message || "修改失败", { type: "error" });
                });
            }
          }
        });
      }
    });
  }

  /** 数据权限 可自行开发 */
  // function handleDatabase() {}

  const onQueryChanged = (query: string) => {
    tableRef.value!.filter(query);
  };

  const filterMethod = (query: string, node) => {
    return transformI18n(node.title)!.includes(query);
  };

  onMounted(async () => {
    onSearch();
  });

  return {
    form,
    isShow,
    curRow,
    loading,
    columns,
    dataList,
    pagination,
    // buttonClass,
    onSearch,
    resetForm,
    openDialog,
    handleDelete,
    filterMethod,
    transformI18n,
    onQueryChanged,
    // handleDatabase,
    handleSizeChange,
    handleCurrentChange,
    handleSelectionChange
  };
}
