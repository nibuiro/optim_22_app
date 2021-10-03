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
    // パスワードが一致しているかチェック
    isPasswordsCorrect() {
      return this.user.password === this.user.confirm_password;
    },
    // ユーザ登録処理
    async register() {
      // すべての情報が正しく入力されていれば
      if (this.isAllEntered() && this.isPasswordsCorrect()) {
        const msgUint8 = new TextEncoder().encode(this.user.password); // パスワードをUint8Array(utf-8)としてエンコード
        const hashBuffer = await crypto.subtle.digest("SHA-256", msgUint8); // エンコードされたパスワードをハッシュ化
        const hashArray = Array.from(new Uint8Array(hashBuffer)); // バッファをbyte配列に変換
        const hashHex = hashArray
          .map(b => b.toString(16).padStart(2, "0"))
          .join(""); // byte配列を16進文字列に変換
        // ログイン情報をサーバに送信し，レスポンスを得る
        fetch(`${process.env.API}/register`, {
          method: "POST",
          body: JSON.stringify({
            username: this.user.username,
            email: this.user.email,
            password: hashHex
          })
        }).then(response => {
          // 登録成功時
          if (response.status === 200) {
            const access_token = response.headers.get("Authorization");
            const refresh_token = response.headers.get("Refresh-Token");
            if (process.env.NODE_ENV === "development") {
              console.log("access_token:");
              console.log(access_token);
              console.log("refresh_token:");
              console.log(refresh_token);
            }
            // レスポンスのbodyをjsonに変換
            response.json().then(data => {
              const user_id = data.user_id;
              if (process.env.NODE_ENV === "development") {
                console.log(`user_id: ${user_id}`);
              }
              // localStorageにユーザIDを保存
              localStorage.setItem("user_id", user_id);
              // localStorageにアクセストークンを保存
              localStorage.setItem("access_token", access_token);
              // cookieにリフレッシュトークンを保存（有効期限: 1ヶ月）
              this.$cookies.set("refresh_token", refresh_token, "1m");
              // 新規登録フォームを閉じる
              this.$emit("close");
              // ユーザ登録成功メッセージを表示する
              this.$emit("displayMessage");
            });
          }
        });
      } else {
        this.invalid = true;
        if (!this.isAllEntered()) {
          this.errorMessage = "すべての項目を入力してください";
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
                placeholder="Enter password"
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
                placeholder="Confirm password"
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
        this.$router.go({ name: "MyPage", params: { user_id } });
      }
    }
  }
};
</script>
