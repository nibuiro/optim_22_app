<!-- ディスカッションページ -->

<template>
  <div class="container">
    <section class="mb-3">
      <div v-for="comment in comments" :key="comment.comment_id" class="card">
        <div class="card-content">
          <div class="media">
            <div class="media-content is-flex is-justify-content-space-between">
              <span>
                {{ `${comment.comment_id}: ${comment.title}` }}
                <b-tooltip
                  v-show="comment.reply_id !== null"
                  :label="comments[Number(comment.reply_id)].text"
                  position="is-right"
                  size="is-large"
                  type="is-light"
                  multilined
                >
                  <b-icon icon="chevron-double-right" />
                  <span class="has-text-weight-bold">{{
                    `#${comment.reply_id}`
                  }}</span>
                </b-tooltip>
              </span>
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
                params: { user_id: comment.user_id }
              }"
            >
              <div class="mx-auto" :style="iconStyle(32, comment.icon)" />
              <span>{{ comment.username }}</span>
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
    <section v-show="refresh_token !== null" class="mt-5">
      <b-message v-show="invalid" type="is-danger">
        {{ errorMessage }}
      </b-message>
      <article class="media">
        <div class="mr-3">
          <b-field label="返信先No.">
            <b-select v-model="comment.reply_id">
              <option :value="null" selected>なし</option>
              <option
                v-for="comment in comments"
                :value="comment.comment_id"
                :key="comment.comment_id"
              >
                {{ comment.comment_id }}
              </option>
            </b-select>
          </b-field>
          <div class="mx-auto" :style="iconStyle(64, icon)" />
        </div>
        <div class="media-content">
          <b-field label="コメント">
            <b-input
              v-model="comment.title"
              placeholder="タイトルを入力してください"
              required
            />
          </b-field>
          <div>
            <b-input
              type="textarea"
              v-model="comment.text"
              placeholder="コメントを入力してください(500字以内)"
              maxlength="500"
              required
            />
            <b-input
              type="text"
              v-model="comment.attachment"
              placeholder="添付ファイルのURL"
            />
            <div class="has-text-centered mt-4">
              <b-button type="is-primary" label="送信" @click="addComment" />
            </div>
          </div>
        </div>
      </article>
    </section>
    <b-modal v-model="isMessageModalActive">
      <b-message type="is-success" has-icon>
        コメントを投稿しました
        <br />
        ディスカッションを更新します
      </b-message>
    </b-modal>
  </div>
</template>

<script>
import * as api from "@/modules/API";

export default {
  data() {
    return {
      refresh_token: this.$cookies.get("refresh_token"),
      icon: "",
      comments: [],
      comment: {
        request_id: null,
        user_id: null,
        reply_id: null,
        title: "",
        text: "",
        attachment: ""
      },
      invalid: false,
      errorMessage: "",
      isMessageModalActive: false
    };
  },
  watch: {
    comment: {
      handler() {
        // タイトルととコメントが入力されていればアラートを消す
        if (this.comment.title.length * this.comment.text.length > 0) {
          this.invalid = false;
        }
      },
      deep: true
    },
    // コメント投稿成功メッセージを閉じたらディスカッションをリロードする
    async isMessageModalActive(newVal, oldVal) {
      if (newVal === false && oldVal === true) {
        this.comment.reply_id = null;
        this.comment.title = "";
        this.comment.text = "";
        this.comments = await api.getComments(this.comment.request_id);
      }
    }
  },
  methods: {
    // コメントを投稿する
    async addComment() {
      // タイトルとコメントが入力されていれば
      if (this.comment.title.length * this.comment.text.length > 0) {
        const access_token = localStorage.getItem("access_token");
        this.isMessageModalActive = await api.addComment(
          this,
          this.comment,
          access_token
        );
      } else {
        this.errorMessage = "タイトルとコメントを入力してください";
        this.invalid = true;
      }
    },
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
    this.comment.request_id = request_id;
    const user_id = localStorage.getItem("user_id");
    this.comment.user_id = user_id;
    this.comments = await api.getComments(request_id);
    const access_token = localStorage.getItem("access_token");
    const profile = await api.getProfile(user_id, access_token);
    this.icon = profile.icon;
  }
};
</script>

<!-- RequestApplier- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
