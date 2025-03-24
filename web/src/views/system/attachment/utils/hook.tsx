import editForm from "../form.vue";
import { message } from "@/utils/message";
import { ElMessageBox } from "element-plus";
import { usePublicHooks } from "../../hooks";
import { transformI18n } from "@/plugins/i18n";
import { addDialog } from "@/components/ReDialog";
import type { PaginationProps } from "@pureadmin/table";
import { deviceDetection } from "@pureadmin/utils";
import {
  getAttachmentList,
  deleteAttachment,
  updateAttachment
} from "@/api/attachment";
import { reactive, ref, onMounted, h, toRaw, computed } from "vue";

import EditPen from "@iconify-icons/ep/edit-pen";
import Check from "@iconify-icons/ep/check";

export function useColumns() {
  const form = reactive({
    name: "",
    type: "",
    drive: "",
    createdAt: [],
    startTime: undefined, // 实际传给后端的开始时间
    endTime: undefined, // 实际传给后端的结束时间
    status: ""
  });
  const curRow = ref();
  const formRef = ref();
  const dataList = ref([]);
  const isShow = ref(false);
  const loading = ref(true);
  const isLinkage = ref(false);
  const switchLoadMap = ref({});
  const { switchStyle } = usePublicHooks();

  const editMap = ref({});
  const activeIndex = ref(-1);

  const editing = computed(() => {
    return index => {
      return editMap.value[index]?.editing;
    };
  });

  const handleTimeChange = (val: [Date, Date] | null) => {
    if (val) {
      form.startTime = val[0];
      form.endTime = val[1];
    } else {
      form.startTime = "";
      form.endTime = "";
    }
  };
  const iconClass = computed(() => {
    return (index, other = false) => {
      return [
        "cursor-pointer",
        "ml-2",
        "transition",
        "delay-100",
        other
          ? ["hover:scale-110", "hover:text-red-500"]
          : editing.value(index) && ["scale-150", "text-red-500"]
      ];
    };
  });

  const pagination = reactive<PaginationProps>({
    total: 0,
    pageSize: 10,
    currentPage: 1,
    background: true
  });
  const columns: TableColumnList = [
    {
      label: "ID",
      prop: "id",
      width: 80
    },
    {
      label: "文件原始名(可修改)",
      prop: "name",
      width: "auto",
      cellRenderer: ({ row, index }) => (
        <div
          class="flex-bc w-full h-[32px]"
          onMouseenter={() => (activeIndex.value = index)}
          onMouseleave={() => onMouseleave(index)}
        >
          {!editing.value(index) ? (
            <p>{row.name}</p>
          ) : (
            <>
              <el-input v-model={row.name} />
              <iconify-icon-offline
                icon={Check}
                class={iconClass.value(index)}
                onClick={() => onSave(index)}
              />
            </>
          )}
          <iconify-icon-offline
            v-show={activeIndex.value === index && !editing.value(index)}
            icon={EditPen}
            class={iconClass.value(index, true)}
            onClick={() => onEdit(row, index)}
          />
        </div>
      )
    },
    {
      label: "文件类型",
      prop: "type",
      width: 100
    },
    {
      label: "文件路径",
      prop: "path"
    },
    {
      label: "文件大小",
      prop: "size",
      width: 100
    },
    {
      label: "后缀名",
      prop: "ext",
      width: 80
    },
    {
      label: "上传驱动",
      prop: "drive",
      width: 90,
      formatter: ({ drive }) => (drive === 1 ? "本地上传" : "未知")
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
      width: 90
    },
    {
      label: "上传时间",
      prop: "createdAt",
      width: 160
    },
    {
      label: "操作",
      fixed: "right",
      width: 140,
      slot: "operation"
    }
  ];
  function onMouseleave(index) {
    editing.value[index]
      ? (activeIndex.value = index)
      : (activeIndex.value = -1);
  }

  function onEdit(row, index) {
    console.log("编辑", row, index);
    editMap.value[index] = Object.assign({ ...row, editing: true });
  }

  async function onSave(index) {
    const currentRow = dataList.value[index];
    const editedRow = editMap.value[index];
    try {
      const res = await updateAttachment(editedRow.id, {
        name: currentRow.name // 使用当前行的最新名称
      });
      if (res.code === 0) {
        message("文件名修改成功", { type: "success" });
        editMap.value[index].editing = false;
        onSearch(); // 刷新表格数据
      } else {
        message(res.message, { type: "error" });
      }
    } catch (error) {
      message(error.message, { type: "error" });
    }
  }

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
      .then(async () => {
        switchLoadMap.value[index] = Object.assign(
          {},
          switchLoadMap.value[index],
          {
            loading: true
          }
        );
        console.log(row);
        try {
          const res = await updateAttachment(row.id, {
            status: row.status
          });

          if (res.code === 0) {
            message(`已${row.status === 0 ? "停用" : "启用"}${row.name}`, {
              type: "success"
            });
            // onSearch(); // 刷新表格数据
          } else {
            message(res.message, { type: "error" });
            // 状态更新失败，恢复原状态
            row.status = row.status === 0 ? 1 : 0;
          }
        } catch (error) {
          message(error, { type: "error" });
          // 发生错误时恢复原状态
          row.status = row.status === 0 ? 1 : 0;
        } finally {
          switchLoadMap.value[index] = Object.assign(
            {},
            switchLoadMap.value[index],
            {
              loading: false
            }
          );
        }
      })
      .catch(() => {
        row.status = row.status === 0 ? 1 : 0;
      });
  }

  async function handleDelete(row) {
    const res = await deleteAttachment(row.id);
    if (res.code === 0) {
      message(`已删除文件：${row.name}`, { type: "success" });
      onSearch(); // 刷新表格数据
    } else {
      message(res.message, { type: "error" });
    }
    onSearch();
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
    const { data } = await getAttachmentList(toRaw(form));
    dataList.value = data.list;
    pagination.total = data.total;
    pagination.pageSize = data.pageSize;
    pagination.currentPage = data.currentPage;

    setTimeout(() => {
      loading.value = false;
    }, 500);
  }

  const resetForm = formEl => {
    if (!formEl) return;
    formEl.resetFields();
    onSearch();
  };

  function openDialog(title = "新增") {
    addDialog({
      title: `${title}`,
      width: "40%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      hideFooter: true, // 使用 hideFooter 来隐藏底部按钮区域
      contentRenderer: () =>
        h(editForm, {
          ref: formRef,
          onSuccess: () => {
            onSearch(); // 刷新列表数据
          },
          onClose: () => {
            // 关闭弹窗的逻辑会由 ReDialog 组件自动处理
          }
        })
    });
  }

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
    isLinkage,
    pagination,
    handleTimeChange,
    onSearch,
    resetForm,
    openDialog,
    handleDelete,
    transformI18n,
    // handleDatabase,
    handleSizeChange,
    handleCurrentChange,
    handleSelectionChange
  };
}
