<template>
  <section>
    <b-button
      label="勝者の決定"
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
        <modal-form :request="formProps" @close="props.close"></modal-form>
      </template>
    </b-modal>
  </section>
</template>

<script>
const ModalForm = {
  props: ["request"],
  data() {
    return {
      winner_id: null
    };
  },
  methods: {
    iconStyle(size, image) {
      return {
        width: `${size}px`,
        height: `${size}px`,
        backgroundImage: `url("${image}")`,
        backgroundSize: "contain",
        backgroundRepeat: "no-repeat",
        backgroundPosition: "center",
        borderRadius: "100%"
      };
    }
  },
  /* html */
  template: `
    <form action="">
      <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
          <p class="modal-card-title has-text-centered">勝者の決定</p>
          <button type="button" class="delete" @click="$emit('close')" />
        </header>
        <section class="modal-card-body">
          <b-field label="提出したエンジニア">
            <section>
              <b-field 
                v-for="engineer in request.engineers"
                :key="engineer.userid"
              >
                <b-radio v-model="winner_id" :native-value="engineer.userid">
                  <div class="is-flex is-align-items-center">
                    <div :style="iconStyle(32, engineer.icon)" />
                    <span>{{ engineer.username }}</span>
                  </div>
                </b-radio>
              </b-field>
            </section>
          </b-field>
        </section>
        <footer class="modal-card-foot is-flex is-justify-content-center">
          <b-button label="決定する" type="is-primary" />
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
        engineers: this.request.engineers
      }
    };
  },
  props: ["request"],
  components: {
    ModalForm
  }
};
</script>
