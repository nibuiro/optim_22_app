<template>
  <section>
    <b-button
      label="参加する"
      type="is-primary"
      @click="isComponentModalActive = true"
    />
    <b-modal
      v-model="isComponentModalActive"
      has-modal-card
      trap-focus
      :destroy-on-hide="false"
      aria-role="dialog"
      aria-label="Example Modal"
      aria-modal
    >
      <template #default="props">
        <modal-form
          :requestProps="formProps"
          @close="props.close"
          @displayMessage="isMessageModalActive = true"
        />
      </template>
    </b-modal>
    <b-modal v-model="isMessageModalActive">
      <b-message type="is-success" has-icon>
        このリクエストに参加しました
        <br />
        ページを更新します
      </b-message>
    </b-modal>
  </section>
</template>

<script>
import * as api from "@/modules/API";

const ModalForm = {
  props: ["requestProps"],
  data() {
    return {
      request: this.requestProps
    };
  },
  methods: {
    // リクエストへの参加
    async joinRequest() {
      console.log(this.request);
      const access_token = localStorage.getItem("access_token");
      api.joinRequest(this, this.request, access_token);
    },
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
  /* html */
  template: `
    <form action="">
      <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
          <p class="modal-card-title has-text-centered">依頼への参加</p>
          <button type="button" class="delete" @click="$emit('close')" />
        </header>
        <section class="modal-card-body">
          <p>このリクエストに参加しますか？</p>
        </section>
        <footer class="modal-card-foot is-flex is-justify-content-center">
          <b-button label="参加する" type="is-primary" @click="joinRequest" />
          <b-button label="キャンセル" @click="$emit('close')" />
        </footer>
      </div>
    </form>
  `
};

export default {
  data() {
    return {
      isComponentModalActive: false,
      isMessageModalActive: false,
      formProps: {
        request_id: this.$route.params.request_id,
        client_id: this.client_id
      }
    };
  },
  watch: {
    // ユーザが参加成功メッセージを閉じたらページをリロードする
    isMessageModalActive(newVal, oldVal) {
      if (newVal === false && oldVal === true) {
        const request_id = this.formProps.request_id;
        this.$router.go({ name: "RequestPage", params: { request_id } });
      }
    }
  },
  props: ["client_id"],
  components: {
    ModalForm
  }
};
</script>
