<!-- ユーザ登録フォーム -->

<template>
  <section>
    <b-button
      label="新規登録"
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
          @close="props.close"
          @displayMessage="isMessageModalActive = true"
        />
      </template>
    </b-modal>
    <b-modal v-model="isMessageModalActive">
      <b-message type="is-success" has-icon>
        ユーザ登録が完了しました
        <br />
        マイページに移動します
      </b-message>
    </b-modal>
  </section>
</template>

<script>
import * as api from "API";

const ModalForm = {
  data() {
    return {
      user: {
        username: "",
        email: "",
        password: "",
        confirm_password: ""
      },
      invalid: false,
      errorMessage: ""
    };
  },
  watch: {
    user: {
      handler() {
        // メールアドレスとパスワードが正しく入力されていればアラートを消す
        if (this.isAllEntered() && this.isPasswordsCorrect()) {
          this.invalid = false;
        }
      },
      deep: true
    }
  },
  methods: {
    // 全項目入力されているかのチェック
    isAllEntered() {
      return (
        this.user.username.length *
          this.user.email.length *
          this.user.password.length *
          this.user.confirm_password.length >
        0
      );
    },
    // パスワードが条件を満たしているかチェック
    isPasswordSecure() {
      return this.user.password.match(
        /^(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9]).{8,}$/
      );
    },
    // パスワードが一致しているかチェック
    isPasswordsCorrect() {
      return this.user.password === this.user.confirm_password;
    },
    // ユーザ登録処理
    register() {
      // すべての情報が正しく入力されていれば
      if (
        this.isAllEntered() &&
        this.isPasswordSecure() &&
        this.isPasswordsCorrect()
      ) {
        api.register(this, this.user);
      } else {
        this.invalid = true;
        if (!this.isAllEntered()) {
          this.errorMessage = "すべての項目を入力してください";
        } else if (!this.isPasswordSecure()) {
          this.errorMessage =
            "パスワードは少なくとも1文字以上の大文字を含んだ、半角英数字8文字以上で入力してください";
        } else if (!this.isPasswordsCorrect()) {
          this.errorMessage = "パスワードが違います";
        }
      }
    }
  },
  /* html */
  template: `
    <form action="">
      <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
          <p class="modal-card-title has-text-centered">新規登録</p>
          <button type="button" class="delete" @click="$emit('close')" />
        </header>
        <section class="modal-card-body">
          <b-message v-show="invalid" type="is-danger">
            {{ errorMessage }}
          </b-message>
          <b-field label="ユーザ名">
            <div class="control has-icons-left">
              <b-icon icon="account" size="is-small"></b-icon>
              <b-input
                type="text"
                v-model="user.username"
                placeholder="username"
                required
              >
              </b-input>
            </div>
          </b-field>
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
                placeholder="パスワード(大文字を含む半角英数字で8文字以上)"
                pattern="^(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9]).{8,}$"
                required
              >
              </b-input>
            </div>
          </b-field>
          <b-field label="パスワード(再入力)">
            <div class="control has-icons-left">
              <b-icon icon="key-outline" size="is-small"></b-icon>
              <b-input
                type="password"
                v-model="user.confirm_password"
                password-reveal
                placeholder="パスワードを再入力してください"
                required
              >
              </b-input>
            </div>
          </b-field>

          <!--b-field>
            <div class="control">
              <label class="checkbox">
                <input type="checkbox">
                <a href="#">利用規約</a>に同意します。
              </label>
            </div>
          </b-field-->
        </section>
        <footer class="modal-card-foot is-flex is-justify-content-center">
          <b-button label="新規登録" type="is-primary" @click="register" />
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
      isComponentModalActive: false,
      isMessageModalActive: false
    };
  },
  watch: {
    // ユーザ登録成功メッセージを閉じたらページをリロードする
    isMessageModalActive(newVal, oldVal) {
      if (newVal === false && oldVal === true) {
        const user_id = localStorage.getItem("user_id");
        if (this.$route.path !== "/user") {
          this.$router.push({ name: "MyPage", params: { user_id } });
        } else {
          this.$router.go({ name: "MyPage", params: { user_id } });
        }
      }
    }
  }
};
</script>
