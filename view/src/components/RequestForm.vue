<template>
  <section>
    <b-button
      label="新規リクエスト"
      type="is-light"
      outlined
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
          @close="props.close"
          @displayMessage="isMessageModalActive = true"
        />
      </template>
    </b-modal>
    <b-modal v-model="isMessageModalActive">
      <b-message type="is-success" has-icon>
        新規リクエストが完了しました
        <br />
        リクエスト一覧を更新します
      </b-message>
    </b-modal>
  </section>
</template>

<script>
const ModalForm = {
  data() {
    return {
      access_token: localStorage.getItem("access_token"),
      user_id: localStorage.getItem("user_id"),
      request: {
        title: "",
        detail: ""
      },
      invalid: false,
      errorMessage: ""
    };
  },
  watch: {
    request: {
      handler() {
        // 依頼名と内容が正しく入力されていればアラートを消す
        if (this.isAllEntered()) {
          this.invalid = false;
        }
      },
      deep: true
    }
  },
  methods: {
    // 全項目入力されているかのチェック
    isAllEntered() {
      return this.request.title.length * this.request.detail.length > 0;
    },
    // 新規リクエスト
    makeRequest() {
      // すべての情報が正しく入力されていれば
      if (this.isAllEntered()) {
        fetch(`${process.env.API}/request`, {
          method: "POST",
          headers: {
            Authorization: this.access_token
          },
          body: JSON.stringify({
            requestname: this.request.title,
            client_id: this.user_id,
            content: this.request.detail
          })
        }).then(response => {
          // 登録成功時
          if (response.status === 200) {
            const access_token = response.headers.get("Authorization");
            if (process.env.NODE_ENV === "development") {
              console.log("access_token:");
              console.log(access_token);
            }
            // localStorageにアクセストークンを保存
            localStorage.setItem("access_token", access_token);
            // 新規リクエストフォームを閉じる
            this.$emit("close");
            // ユーザ登録成功メッセージを表示する
            this.$emit("displayMessage");
          }
        });
      } else {
        this.invalid = true;
        this.errorMessage = "すべての項目を入力してください";
      }
    }
  },
  /* html */
  template: `
    <form action="">
      <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
          <p class="modal-card-title has-text-centered">新規リクエスト</p>
          <button type="button" class="delete" @click="$emit('close')" />
        </header>
        <section class="modal-card-body">
          <b-message v-show="invalid" type="is-danger">
            {{ errorMessage }}
          </b-message>
          <b-field label="依頼名">
            <div class="control">
              <b-input
                type="text"
                v-model="request.title"
                placeholder="依頼内容を分かりやすく一言で！"
                required
              />
            </div>
          </b-field>
          <b-field label="依頼内容の詳細">
            <div class="control">
              <b-input
                type="textarea"
                v-model="request.detail"
                placeholder="依頼内容について具体的に説明してください。(500字以内)"
                maxlength="500"
                required
              />
            </div>
          </b-field>
        </section>
        <footer class="modal-card-foot is-flex is-justify-content-center">
          <b-button label="投稿する" type="is-primary" @click="makeRequest" />
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
  components: {
    ModalForm
  },
  watch: {
    // 新規リクエスト成功メッセージを閉じたら依頼一覧ページをリロードする
    isMessageModalActive(newVal, oldVal) {
      if (newVal === false && oldVal === true) {
        this.$router.go("/");
      }
    }
  }
};
</script>
