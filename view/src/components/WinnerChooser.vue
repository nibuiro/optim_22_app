<!-- 勝者決定フォーム -->

<template>
  <section>
    <b-button
      label="勝者の決定"
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
          :requestProps="request"
          @close="props.close"
          @displayMessage="isMessageModalActive = true"
        />
      </template>
    </b-modal>
    <b-modal v-model="isMessageModalActive">
      <b-message type="is-success" has-icon>
        勝者を決定しました
        <br />
        リクエストを終了します
      </b-message>
    </b-modal>
  </section>
</template>

<script>
import * as api from "API";

const ModalForm = {
  props: ["requestProps"],
  data() {
    return {
      request: {
        request_id: this.requestProps.request_id,
        client_id: this.requestProps.client.user_id,
        engineer_id: ""
      },
      submissions: this.requestProps.submissions,
      invalid: false,
      errorMessage: ""
    };
  },
  methods: {
    // 勝者を決定する
    async chooseWinner() {
      // 勝者が選択されていれば
      if (this.request.engineer_id !== "") {
        const access_token = localStorage.getItem("access_token");
        api.chooseWinner(this, this.request, access_token);
      } else {
        this.errorMessage = "すべての項目を入力してください";
        this.invalid = true;
      }
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
          <p class="modal-card-title has-text-centered">勝者の決定</p>
          <button type="button" class="delete" @click="$emit('close')" />
        </header>
        <section class="modal-card-body">
          <b-message v-show="invalid" type="is-danger">
            {{ errorMessage }}
          </b-message>
          <b-field label="提出したエンジニア一覧">
            <section>
              <b-field
                v-for="submission in submissions"
                :key="submission.engineer.userid"
              >
                <b-radio v-model="request.engineer_id" :native-value="submission.engineer.user_id">
                  <div class="is-flex is-align-items-center">
                    <div :style="iconStyle(32, submission.engineer.icon)" />
                    <span>{{ submission.engineer.username }}</span>
                  </div>
                </b-radio>
              </b-field>
            </section>
          </b-field>
        </section>
        <footer class="modal-card-foot is-flex is-justify-content-center">
          <b-button label="決定する" type="is-primary" @click="chooseWinner" />
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
      isMessageModalActive: false
    };
  },
  watch: {
    // ユーザが勝者決定成功メッセージを閉じたらページをリロードする
    isMessageModalActive(newVal, oldVal) {
      if (newVal === false && oldVal === true) {
        const request_id = this.$route.params.request_id;
        this.$router.go({ name: "RequestPage", params: { request_id } });
      }
    }
  },
  props: ["request"],
  components: {
    ModalForm
  }
};
</script>
