<template>
  <div class="hidden md:flex justify-center items-center border m-4">
    <img class="m-4 h-72" src="../assets/illustration/guide.webp" alt="" />
    <div class="m-4">
      <h2 class="text-xl text-main">
          <i18n-t keypath="intro.content">
            <template v-slot:quickstart>
              <span class="text-accent">{{ $t("common.quickstart") }}</span>
            </template>
            <template v-slot:doc>
              <a class="normal-link" href="https://docs.bytebase.com" target="__blank">
                {{ $t("intro.doc") }}
              </a>
            </template>
            <template v-slot:issue>
              <a
                class="normal-link"
                href="https://github.com/bytebase/bytebase/issues"
                target="__blank"
              >
                {{ $t("intro.issue") }}
              </a>
            </template>
          </i18n-t>
      </h2>
      <div class="mt-4">
        <button class="btn-normal" @click.prevent="dismiss">
          {{ $t("common.dismiss") }}
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { useStore } from "vuex";

export default {
  name: "IntroBanner",
  components: {},
  setup() {
    const store = useStore();
    const dismiss = () => {
      store.dispatch("uistate/saveIntroStateByKey", {
        key: "general.overview",
        newState: true,
      });
    };
    return {
      dismiss,
    };
  },
};
</script>
