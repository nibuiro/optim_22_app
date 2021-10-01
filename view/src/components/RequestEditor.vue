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
        <modal-form :request="formProps" @close="props.close"></modal-form>
      </template>
    </b-modal>
  </section>
</template>

<script>
const ModalForm = {
  props: ["request"],
  /* html */
  template: `
    <form action="">
      <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
          <p class="modal-card-title has-text-centered">リクエスト内容の編集</p>
          <button type="button" class="delete" @click="$emit('close')" />
        </header>
        <section class="modal-card-body">
          <b-field label="依頼名">
            <div class="control">
              <b-input
                type="text"
                :value="request.title"
                placeholder="依頼内容を分かりやすく一言で！"
                required
              />
            </div>
          </b-field>
          <b-field label="依頼内容の詳細">
            <div class="control">
              <b-input
                type="textarea"
                :value="request.detail"
                placeholder="依頼内容について具体的に説明してください。(500字以内)"
                maxlength="500"
                required
              />
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
        title: this.request.request,
        detail: this.request.detail
      }
    };
  },
  props: ["request"],
  components: {
    ModalForm
  }
};
</script>
