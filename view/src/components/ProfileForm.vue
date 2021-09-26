<template>
  <section>
    <b-button
      label="編集"
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
        <modal-form :profile="formProps" @close="props.close"></modal-form>
      </template>
    </b-modal>
  </section>
</template>

<script>
const ModalForm = {
  props: ["profile"],
  data() {
    return {
      file: null,
      password: "",
      confirm_password: ""
    };
  },
  /* html */
  template: `
    <form action="">
      <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
          <p class="modal-card-title has-text-centered">プロフィールの編集</p>
          <button type="button" class="delete" @click="$emit('close')" />
        </header>
        <section class="modal-card-body">
          <b-field label="ユーザ名">
            <div class="control has-icons-left">
              <b-icon icon="account" size="is-small"></b-icon>
              <b-input
                type="text"
                :value="profile.username"
                placeholder="username"
                required
              />
            </div>
          </b-field>
          <b-field label="アイコン画像">
            <p class="control">
              <div class="control has-icons-left">
                <b-icon icon="image" size="is-small" />
                <b-input :value="!!file?file.name:''" />
              </div>
            </p>
            <b-field class="file is-primary">
              <b-upload v-model="file" class="file-label">
                <span class="file-cta">
                  <b-icon class="file-icon" icon="upload" />
                  <span class="file-label">アップロード</span>
                </span>
              </b-upload>
            </b-field>
          </b-field>
          <b-field label="メールアドレス">
            <div class="control has-icons-left">
              <b-icon icon="email" size="is-small"></b-icon>
              <b-input
                type="email"
                :value="profile.email"
                placeholder="email@example.com"
                required
              />
            </div>
          </b-field>
          <b-field label="自己紹介">
            <div class="control has-icons-left">
              <b-icon icon="comment" size="is-small"></b-icon>
              <b-input
                :value="profile.comment"
                placeholder="email@example.com"
                required
              />
            </div>
          </b-field>
          <b-field label="SNSアカウント" grouped group-multiline>
            <p class="control">
              <b-field>
                <p class="control">
                  <span class="button is-static is-flex is-align-items-center" style="background-color: #171516;">
                    <b-icon class="is-inline-flex" icon="github" type="is-light" />
                  </span>
                </p>
                <b-input :value="profile.SNS.Github" placeholder="Github" />
              </b-field>
            </p>
            <p class="control">
              <b-field>
                <p class="control">
                  <span class="button is-static is-flex is-align-items-center" style="background-color: #1D9BF0;">
                    <b-icon class="is-inline-flex" icon="twitter" type="is-light" />
                  </span>
                </p>
                <b-input :value="profile.SNS.Twitter" placeholder="Twitter" />
              </b-field>
            </p>
            <p class="control">
              <b-field>
                <p class="control">
                  <span class="button is-static is-flex is-align-items-center" style="background-color: #1877F2;">
                    <b-icon class="is-inline-flex" icon="facebook" type="is-light" />
                  </span>
                </p>
                <b-input :value="profile.SNS.Facebook" placeholder="Facebook" />
              </b-field>
            </p>
          </b-field>
          <b-field label="パスワード">
            <div class="control has-icons-left">
              <b-icon icon="key" size="is-small"></b-icon>
              <b-input
                type="password"
                :value="password"
                password-reveal
                placeholder="Enter password"
                required
              />
            </div>
          </b-field>
          <b-field label="パスワード(再入力)">
            <div class="control has-icons-left">
              <b-icon icon="key-outline" size="is-small"></b-icon>
              <b-input
                type="password"
                :value="confirm_password"
                password-reveal
                placeholder="Confirm password"
                required
              >
              </b-input>
            </div>
          </b-field>
        </section>
        <footer class="modal-card-foot is-flex is-justify-content-center">
          <b-button label="編集する" type="is-primary" />
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
      formProps: {
        username: this.profile.username,
        email: this.profile["e-mail"],
        comment: this.profile.comment,
        SNS: {
          Github: this.profile.SNS.Github,
          Facebook: this.profile.SNS.Facebook,
          Twitter: this.profile.SNS.Twitter
        }
      }
    };
  },
  props: ["profile"],
  components: {
    ModalForm
  }
};
</script>
