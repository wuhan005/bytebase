<template>
  <BBTable
    :column-list="columnList"
    :data-source="task.taskRunList"
    :show-header="true"
    :left-bordered="true"
    :right-bordered="true"
    :row-clickable="false"
  >
    <template #body="{ rowData: taskRun }">
      <BBTableCell :left-padding="4" class="table-cell w-4">
        <div class="flex flex-row space-x-2">
          <div
            class="relative w-5 h-5 flex flex-shrink-0 items-center justify-center rounded-full select-none"
            :class="statusIconClass(taskRun.status)"
          >
            <template v-if="taskRun.status == 'RUNNING'">
              <span
                class="h-2.5 w-2.5 bg-info rounded-full"
                style="
                  animation: pulse 2.5s cubic-bezier(0.4, 0, 0.6, 1) infinite;
                "
                aria-hidden="true"
              />
            </template>
            <template v-else-if="taskRun.status == 'DONE'">
              <heroicons-outline:check class="w-5 h-5" />
            </template>
            <template v-else-if="taskRun.status == 'FAILED'">
              <span class="text-white font-medium text-base" aria-hidden="true">
                !
              </span>
            </template>
            <template v-else-if="taskRun.status == 'CANCELED'">
              <heroicons-outline:minus-sm class="w-5 h-5" />
            </template>
          </div>
        </div>
      </BBTableCell>
      <BBTableCell class="table-cell w-36">
        {{ comment(taskRun) }}
        <template v-if="commentLink(taskRun).link">
          <router-link
            class="ml-1 normal-link"
            :to="commentLink(taskRun).link"
            >{{ commentLink(taskRun).title }}</router-link
          >
        </template>
      </BBTableCell>
      <BBTableCell class="table-cell w-12">
        <div class="flex flex-row items-center space-x-2">
          <PrincipalAvatar :principal="taskRun.creator" :size="'SMALL'" />
          <div class="flex flex-col">
            <div class="flex flex-row items-center space-x-2">
              <router-link :to="`/u/${taskRun.creator.id}`"
                >{{ taskRun.creator.name }}
              </router-link>
            </div>
          </div>
        </div>
      </BBTableCell>
      <BBTableCell class="table-cell w-12">
        {{ humanizeTs(taskRun.createdTs) }}
      </BBTableCell>
      <BBTableCell class="table-cell w-12">
        {{ humanizeTs(taskRun.updatedTs) }}
      </BBTableCell>
    </template>
  </BBTable>
</template>

<script lang="ts">
import { computed, defineComponent, PropType } from "vue";
import PrincipalAvatar from "./PrincipalAvatar.vue";
import { BBTableColumn } from "../bbkit/types";
import { MigrationErrorCode, Task, TaskRun, TaskRunStatus } from "../types";
import { databaseSlug, instanceSlug, migrationHistorySlug } from "../utils";
import { useI18n } from "vue-i18n";

type CommentLink = {
  title: string;
  link: string;
};

export default defineComponent({
  name: "TaskRunTable",
  components: { PrincipalAvatar },
  props: {
    task: {
      required: true,
      type: Object as PropType<Task>,
    },
  },
  setup(props) {
    const { t } = useI18n();

    const columnList = computed((): BBTableColumn[] => [
      {
        title: "",
      },
      {
        title: t("task.comment"),
      },
      {
        title: t("task.invoker"),
      },
      {
        title: t("task.started"),
      },
      {
        title: t("task.ended"),
      },
    ]);

    const statusIconClass = (status: TaskRunStatus) => {
      switch (status) {
        case "RUNNING":
          return "bg-white border-2 border-info text-info";
        case "DONE":
          return "bg-success text-white";
        case "FAILED":
          return "bg-error text-white";
        case "CANCELED":
          return "bg-white border-2 border-gray-400 text-gray-400";
      }
    };

    const comment = (taskRun: TaskRun): string => {
      if (taskRun.status == "FAILED") {
        return taskRun.result.detail;
      }
      // Returns result detail if we get the result, otherwise, returns the comment.
      return taskRun.result.detail || taskRun.comment;
    };

    const commentLink = (taskRun: TaskRun): CommentLink => {
      if (taskRun.status == "DONE") {
        switch (taskRun.type) {
          case "bb.task.database.schema.update": {
            return {
              title: t("task.view-migration"),
              link: `/db/${databaseSlug(
                props.task.database!
              )}/history/${migrationHistorySlug(
                taskRun.result.migrationId!,
                taskRun.result.version!
              )}`,
            };
          }
        }
      } else if (taskRun.status == "FAILED") {
        if (taskRun.code == MigrationErrorCode.MIGRATION_SCHEMA_MISSING) {
          return {
            title: "Check instance",
            link: `/instance/${instanceSlug(props.task.instance)}`,
          };
        } else if (
          taskRun.code == MigrationErrorCode.MIGRAITON_ALREADY_APPLIED ||
          taskRun.code == MigrationErrorCode.MGIRATION_OUT_OF_ORDER ||
          taskRun.code == MigrationErrorCode.MIGRATION_BASELINE_MISSING
        ) {
          return {
            title: t("task.view-migration-history"),
            link: `/db/${databaseSlug(props.task.database!)}#migration-history`,
          };
        }
      }
      return {
        title: "",
        link: "",
      };
    };

    return {
      columnList,
      statusIconClass,
      comment,
      commentLink,
    };
  },
});
</script>
