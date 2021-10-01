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
        <modal-form :submission="formProps" @close="props.close"></modal-form>
      </template>
    </b-modal>
  </section>
</template>

<script>
const ModalForm = {
  props: ["submission"],
  /* html */
  template: `
    <form action="">
      <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
          <p class="modal-card-title has-text-centered">提出物の編集</p>
          <button type="button" class="delete" @click="$emit('close')" />
        </header>
        <section class="modal-card-body">
          <b-field label="提出物がダウンロード可能なURL">
            <div class="control">
              <b-input
                type="text"
                :value="submission.url"
                placeholder="アップロード先URL"
                required
              />
            </div>
          </b-field>
          <b-field label="提出物の詳細">
            <div class="control">
              <b-input
                type="textarea"
                :value="submission.comment"
                placeholder="提出物について具体的に説明してください。(500字以内)"
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
        url: this.submission.url,
        comment: this.submission.comment
      }
    };
  },
  props: ["submission"],
  components: {
    ModalForm
  }
};
</script>
