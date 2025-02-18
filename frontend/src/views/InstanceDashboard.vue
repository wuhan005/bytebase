<template>
  <div class="flex flex-col">
    <div class="px-5 py-2 flex justify-between items-center">
      <!-- eslint-disable vue/attribute-hyphenation -->
      <EnvironmentTabFilter
        :selectedId="state.selectedEnvironment?.id"
        @select-environment="selectEnvironment"
      />
      <BBTableSearch
        ref="searchField"
        :placeholder="$t('instance.search-instance-name')"
        @change-text="(text) => changeSearchText(text)"
      />
    </div>
    <InstanceTable :instance-list="filteredList(instanceList)" />
  </div>

  <BBAlert
    v-if="state.showGuide"
    :style="'INFO'"
    :ok-text="$t('common.do-not-show-again')"
    :cancel-text="$t('common.dismiss')"
    :title="$t('instance.how-to-setup-instance')"
    :description="$t('instance.how-to-setup-instance-description')"
    @ok="
      () => {
        doDismissGuide();
      }
    "
    @cancel="state.showGuide = false"
  >
  </BBAlert>
</template>

<script lang="ts">
import { computed, watchEffect, onMounted, reactive, ref } from "vue";
import { useRouter } from "vue-router";
import EnvironmentTabFilter from "../components/EnvironmentTabFilter.vue";
import InstanceTable from "../components/InstanceTable.vue";
import { useStore } from "vuex";
import { Environment, Instance } from "../types";
import { cloneDeep } from "lodash-es";
import { sortInstanceList } from "../utils";

interface LocalState {
  searchText: string;
  selectedEnvironment?: Environment;
  showGuide: boolean;
}

export default {
  name: "InstanceDashboard",
  components: {
    EnvironmentTabFilter,
    InstanceTable,
  },
  setup() {
    const searchField = ref();

    const store = useStore();
    const router = useRouter();

    const environmentList = computed(() => {
      return store.getters["environment/environmentList"](["NORMAL"]);
    });

    const state = reactive<LocalState>({
      searchText: "",
      selectedEnvironment: router.currentRoute.value.query.environment
        ? store.getters["environment/environmentById"](
            router.currentRoute.value.query.environment
          )
        : undefined,
      showGuide: false,
    });

    onMounted(() => {
      // Focus on the internal search field when mounted
      searchField.value.$el.querySelector("#search").focus();

      if (!store.getters["uistate/introStateByKey"]("guide.instance")) {
        setTimeout(() => {
          state.showGuide = true;
          store.dispatch("uistate/saveIntroStateByKey", {
            key: "instance.visit",
            newState: true,
          });
        }, 1000);
      }
    });

    const prepareInstanceList = () => {
      store.dispatch("instance/fetchInstanceList");
    };

    watchEffect(prepareInstanceList);

    const selectEnvironment = (environment: Environment) => {
      state.selectedEnvironment = environment;
      if (environment) {
        router.replace({
          name: "workspace.instance",
          query: { environment: environment.id },
        });
      } else {
        router.replace({ name: "workspace.instance" });
      }
    };

    const changeSearchText = (searchText: string) => {
      state.searchText = searchText;
    };

    const doDismissGuide = () => {
      store.dispatch("uistate/saveIntroStateByKey", {
        key: "guide.instance",
        newState: true,
      });
      state.showGuide = false;
    };

    const instanceList = computed(() => {
      const list = store.getters["instance/instanceList"]();
      return sortInstanceList(cloneDeep(list), environmentList.value);
    });

    const filteredList = (list: Instance[]) => {
      if (!state.selectedEnvironment && !state.searchText) {
        // Select "All"
        return list;
      }
      return list.filter((instance) => {
        return (
          (!state.selectedEnvironment ||
            instance.environment.id == state.selectedEnvironment.id) &&
          (!state.searchText ||
            instance.name
              .toLowerCase()
              .includes(state.searchText.toLowerCase()))
        );
      });
    };

    return {
      searchField,
      state,
      instanceList,
      filteredList,
      selectEnvironment,
      changeSearchText,
      doDismissGuide,
    };
  },
};
</script>
