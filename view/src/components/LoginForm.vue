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
        <modal-form v-bind="formProps" @close="props.close"></modal-form>
      </template>
    </b-modal>
  </section>
</template>

<script>
const ModalForm = {
  props: ["email", "password", "canCancel"],
  /* html */
  template: `
    <form action="">
      <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
          <p class="modal-card-title has-text-centered">ログイン</p>
          <button type="button" class="delete" @click="$emit('close')" />
        </header>
        <section class="modal-card-body">
          <b-field label="メールアドレス">
            <div class="control has-icons-left">
              <b-icon icon="email" size="is-small"></b-icon>
              <b-input
                type="email"
                :value="email"
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
                :value="password"
                password-reveal
                placeholder="Enter password"
                required
              >
              </b-input>
            </div>
          </b-field>

          <!--b-checkbox>Remember me</!--b-checkbox-->
        </section>
        <footer class="modal-card-foot is-flex is-justify-content-center">
          <b-button label="ログイン" type="is-primary" />
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
      formProps: {
        email: "",
        password: ""
      }
    };
  }
};
</script>
