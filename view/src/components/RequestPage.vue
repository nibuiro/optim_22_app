<!-- リクエスト詳細ページ -->

<template>
  <div class="container" :request="request">
    <section class="hero is-primary is-small mb-3">
      <b-tooltip
        style="position: absolute;"
        :label="request.detail"
        type="is-light"
        position="is-right"
        always
      >
        <div
          class="ml-3 mt-3 mb-6"
          :style="iconStyle(64, request.client.icon)"
        />
      </b-tooltip>
      <div class="hero-body is-flex pt-0 pb-5">
        <p class="title mb-0 pt-2" style="margin-left: 64px;">
          {{ request.request }}
        </p>
        <request-form class="is-light ml-auto mt-5" :profile="request" />
      </div>
    </section>
    <section class="mb-3">
      a
    </section>
  </div>
</template>

<script>
import RequestForm from "@/components/RequestForm";

const request = require("../../src/assets/sampleRequest.json");

export default {
  data() {
    return {
      request
    };
  },
  methods: {
    iconStyle(size, image) {
      return {
        width: `${size}px`,
        height: `${size}px`,
        backgroundImage: `url("${image}")`,
        backgroundSize: "contain",
        backgroundRepeat: "no-repeat",
        backgroundPosition: "center",
        borderRadius: "100%"
      };
    }
  },
  components: {
    "request-form": RequestForm
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
