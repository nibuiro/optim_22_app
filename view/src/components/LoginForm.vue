<!-- ログインフォーム -->

<template>
  <section>
    <b-button
      label="ログイン"
      type="is-primary"
      @click="isComponentModalActive = true"
      outlined
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
        <modal-form @close="props.close" />
      </template>
    </b-modal>
  </section>
</template>

<script>
import * as api from "@/modules/API.js";

const ModalForm = {
  data() {
    return {
      user: {
        email: "",
        password: ""
      },
      invalid: false,
      errorMessage: ""
    };
  },
  watch: {
    user: {
      handler() {
        // メールアドレスとパスワードが入力されていればアラートを消す
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
      return this.user.email.length * this.user.password.length > 0;
    },
    // ログイン処理
    login() {
      // メールアドレスとパスワードの両方が入力されていれば
      if (this.isAllEntered()) {
        api.login(this, this.user);
      } else {
        this.errorMessage = "すべての項目を入力してください";
        this.invalid = true;
      }
    }
  },
  /* html */
  template: `
    <form action="">
      <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
          <p class="modal-card-title has-text-centered">ログイン</p>
          <button type="button" class="delete" @click="$emit('close')" />
        </header>
        <section class="modal-card-body">
          <b-message v-show="invalid" type="is-danger">
            {{ errorMessage }}
          </b-message>
          <b-field label="メールアドレス">
            <div class="control has-icons-left">
              <b-icon icon="email" size="is-small"></b-icon>
              <b-input
                type="email"
                v-model="user.email"
                placeholder="email@example.com"
                required
              >
              </b-input>
            </div>
          </b-field>
          <b-field label="パスワード">
            <div class="control has-icons-left">
              <b-icon icon="key" size="is-small"></b-icon>
              <b-input
                type="password"
                v-model="user.password"
                password-reveal
                placeholder="Enter password"
                required
              >
              </b-input>
            </div>
          </b-field>
        </section>
        <footer class="modal-card-foot is-flex is-justify-content-center">
          <b-button label="ログイン" type="is-primary" @click="login" />
          <b-button label="キャンセル" @click="$emit('close')" />
        </footer>
      </div>
    </form>
  `
};

export default {
  components: {
    ModalForm
  },
  data() {
    return {
      isComponentModalActive: false
    };
  }
};
</script>
