<!-- ディスカッションページ -->

<template>
  <div class="container">
    <section class="mb-3">
      <div v-for="comment in comments" :key="comment.comment_id" class="card">
        <div class="card-content">
          <div class="media">
            <div class="media-content is-flex is-justify-content-space-between">
              <span>{{ `${comment.comment_id}: ${comment.title}` }}</span>
              <span>
                {{
                  `${new Date(comment.createdat).toLocaleDateString()}
                ${new Date(comment.createdat).toLocaleTimeString()}`
                }}
              </span>
            </div>
          </div>
          <div class="is-flex">
            <router-link
              class="is-flex is-flex-direction-column"
              :to="{
                name: 'MyPage',
                params: { user_id: comment.user.user_id }
              }"
            >
              <div class="mx-auto" :style="iconStyle(32, comment.user.icon)" />
              <span>{{ comment.user.username }}</span>
            </router-link>
            <div class="content pl-3">
              {{ comment.text }}
            </div>
          </div>
          <span
            v-if="!!comment.attachment"
            class="is-flex is-justify-content-end"
          >
            <a class="is-flex is-align-items-center" :href="comment.attachment">
              <b-icon icon="attachment" />
              添付ファイル
            </a>
          </span>
        </div>
      </div>
    </section>
    <!-- ログイン済みなら -->
    <section class="mt-5">
      <article class="media">
        <figure class="media-left">
          <p class="image is-64x64">
            <img src="https://bulma.io/images/placeholders/128x128.png" />
          </p>
        </figure>
        <div class="media-content">
          <b-field>
            <div class="control">
              <b-input
                type="textarea"
                :value="comment"
                placeholder="コメントを入力する(500字以内)"
                maxlength="500"
              />
              <div class="has-text-centered">
                <b-button type="is-primary" label="送信" />
              </div>
            </div>
          </b-field>
        </div>
      </article>
    </section>
  </div>
</template>

<script>
import * as api from "@/modules/API";

export default {
  data() {
    return {
      comments: [],
      comment: ""
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
  async created() {
    const request_id = this.$route.params.request_id;
    this.comments = await api.getComments(request_id);
  }
};
</script>

<!-- RequestApplier- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
