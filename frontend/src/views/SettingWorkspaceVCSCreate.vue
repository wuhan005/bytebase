<template>
  <div class="mt-4 space-y-4 divide-y divide-block-border">
    <div class="textinfolabel">
      <i18n-t keypath="version-control.setting.add-git-provider.description">
        <template #guide>
          <a
            href="https://docs.bytebase.com/use-bytebase/vcs-integration?ref=console"
            target="_blank"
            class="normal-link"
          >
            {{ $t('common.detailed-guide') }}</a
          >
      </template>
      </i18n-t>
    </div>
    <div class="pt-4">
      <VCSSetupWizard />
    </div>
  </div>
</template>

<script lang="ts">
import { reactive } from "vue";
import VCSSetupWizard from "../components/VCSSetupWizard.vue";

// eslint-disable-next-line @typescript-eslint/no-empty-interface
interface LocalState {}

export default {
  name: "SettingWorkspaceVCSCreate",
  components: {
    VCSSetupWizard,
  },
  setup() {
    const state = reactive<LocalState>({});

    return {
      state,
    };
  },
};
</script>
