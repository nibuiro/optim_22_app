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
        <modal-form @close="props.close"></modal-form>
      </template>
    </b-modal>
  </section>
</template>

<script>
const ModalForm = {
  data() {
    return {
      user: {
        email: "",
        password: ""
      },
      invalid: false
    };
  },
  watch: {
    user: {
      handler() {
        // メールアドレスとパスワードが入力されていればアラートを消す
        if (this.user.email.length * this.user.password.length > 0) {
          this.invalid = false;
        }
      },
      deep: true
    }
  },
  methods: {
    // ログイン処理
    async login() {
      // メールアドレスとパスワードの両方が入力されていれば
      if (this.user.email.length * this.user.password.length > 0) {
        const msgUint8 = new TextEncoder().encode(this.user.password); // パスワードをUint8Array(utf-8)としてエンコード
        const hashBuffer = await crypto.subtle.digest("SHA-256", msgUint8); // エンコードされたパスワードをハッシュ化
        const hashArray = Array.from(new Uint8Array(hashBuffer)); // バッファをbyte配列に変換
        const hashHex = hashArray
          .map(b => b.toString(16).padStart(2, "0"))
          .join(""); // byte配列を16進文字列に変換
        // ログイン情報をサーバに送信し，レスポンスを得る
        fetch(`${process.env.API}/auth`, {
          method: "POST",
          body: JSON.stringify({
            email: this.user.email,
            password: hashHex
          })
        }).then(response => {
          // ログイン成功時
          if (response.status === 200) {
            const access_token = response.headers.get("Authorization");
            if (process.env.NODE_ENV === "development") {
              console.log("access_token:");
              console.log(access_token);
            }
            // レスポンスのbodyをjsonに変換
            response.json().then(data => {
              if (process.env.NODE_ENV === "development") {
                console.log("user_id:");
                console.log(data);
              }
              // localStorageにユーザIDを保存
              localStorage.setItem("user_id", data.user_id);
              // localStorageにアクセストークンを保存
              localStorage.setItem("access_token", access_token);
            });
          }
        });
      } else {
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
            正しく入力してください
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
