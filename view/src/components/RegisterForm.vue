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
        <modal-form v-bind="formProps" @close="props.close"></modal-form>
      </template>
    </b-modal>
  </section>
</template>

<script>
const ModalForm = {
  props: ["email", "password", "canCancel"],
  template: `
            <form action="">
                <div class="modal-card" style="width: auto">
                    <header class="modal-card-head">
                        <p class="modal-card-title has-text-centered">ログイン</p>
                        <button
                            type="button"
                            class="delete"
                            @click="$emit('close')"/>
                    </header>
                    <section class="modal-card-body">
                        <b-field label="ユーザ名">
                        <div class="control has-icons-left">
                            <b-icon icon="account" size="is-small"></b-icon>
                            <b-input
                              type="text"
                              :value="username"
                              placeholder="username"
                              required>
                            </div>
                        </b-field>

                        <b-field label="メールアドレス">
                          <div class="control has-icons-left">
                            <b-icon icon="email" size="is-small"></b-icon>
                            <b-input
                                type="email"
                                :value="email"
                                placeholder="email@example.com"
                                required>
                            </b-input>
                          </div>
                        </b-field>

                        <b-field label="パスワード">
                          <div class="control has-icons-left">
                            <b-icon icon="key" size="is-small"></b-icon>
                            <b-input
                                type="password"
                                :value="password"
                                password-reveal
                                placeholder="Enter password"
                                required>
                            </b-input>
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
                                required>
                            </b-input>
                          </div>
                        </b-field>

                    </section>
                    <footer class="modal-card-foot is-flex is-justify-content-center">
                        <b-button
                            label="新規登録"
                            type="is-primary" />
                        <b-button
                            label="キャンセル"
                            @click="$emit('close')" />
                        
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
      formProps: {
        username: "",
        email: "",
        password: "",
        confirm_password: ""
      }
    };
  }
};
</script>
