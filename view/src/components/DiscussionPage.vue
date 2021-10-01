<!-- ディスカッションページ -->

<template>
  <div class="container" :request="request">
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
              params: { user_id: comment.user.userid }
            }"
          >
            <b-tooltip
              type="is-light"
              :label="comment.text"
              position="is-right"
              always
            >
              <div class="mx-auto" :style="iconStyle(32, comment.user.icon)" />
            </b-tooltip>
            <span>{{ comment.user.username }}</span>
          </router-link>
        </div>
        <span
          v-if="!!comment.attachment"
          class="is-flex is-justify-content-end"
        >
          <a class="is-flex is-align-items-center" :href="comment.attachment">
            <b-icon icon="attachment" />
            {{ comment.attachment }}
          </a>
        </span>
      </div>
    </div>
  </div>
</template>

<script>
const comments = require("../../src/assets/sampleComments.json");

export default {
  data() {
    return {
      comments: comments.comments
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
  }
};
</script>

<!-- RequestApplier- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
