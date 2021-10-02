<template>
  <div class="container" :profile="profile">
    <section class="hero is-primary is-small mb-3">
      <b-tooltip
        style="position: absolute;"
        :label="profile.comment"
        type="is-light"
        position="is-right"
        always
      >
        <div class="ml-3 mt-3 mb-6" :style="iconStyle(64, profile.icon)" />
      </b-tooltip>
      <div class="hero-body is-flex pt-0 pb-5">
        <p class="title mb-0 pt-2" style="margin-left: 64px;">
          {{ profile.username }}
        </p>
        <profile-editor class="is-light ml-auto mt-5" :profile="profile" />
      </div>
    </section>
    <section class="mb-3 is-flex is-justify-content-space-evenly">
      <b-taglist class="m-0" attached>
        <b-tag style="background-color: #171516;">
          <b-icon icon="github" type="is-light" />
        </b-tag>
        <b-tag type="is-light">
          <a href="">@{{ profile.SNS.Github }}</a>
        </b-tag>
      </b-taglist>
      <b-taglist class="m-0" attached>
        <b-tag style="background-color: #1D9BF0;">
          <b-icon icon="twitter" type="is-light" />
        </b-tag>
        <b-tag type="is-light">
          <a href="">@{{ profile.SNS.Twitter }}</a>
        </b-tag>
      </b-taglist>
      <b-taglist class="m-0" attached>
        <b-tag style="background-color: #1877F2;">
          <b-icon icon="facebook" type="is-light" />
        </b-tag>
        <b-tag type="is-light">
          <a href="">@{{ profile.SNS.Facebook }}</a>
        </b-tag>
      </b-taglist>
    </section>
    <section class="mb-3">
      <b-tabs type="is-boxed">
        <b-tab-item>
          <template #header>
            <b-icon icon="account-question-outline"></b-icon>
            <span>
              過去にした依頼
              <b-tag rounded> {{ profile.requests.length }} </b-tag>
            </span>
          </template>
          <b-table :data="profile.requests" :default-sort="['date', 'desc']">
            <b-table-column
              cell-class="is-vcentered"
              field="state"
              label="状態"
              width="10%"
              centered
              v-slot="props"
            >
              <b-tag
                :type="props.row.finish === false ? 'is-success' : 'is-danger'"
              >
                {{ props.row.finish === false ? "受付中" : "終了" }}
              </b-tag>
            </b-table-column>
            <b-table-column
              cell-class="is-vcentered"
              field="date"
              label="依頼日時"
              width="10%"
              sortable
              centered
              v-slot="props"
            >
              {{ new Date(props.row.createdat).toLocaleDateString() }}
              <br />
              {{ new Date(props.row.createdat).toLocaleTimeString() }}
            </b-table-column>
            <b-table-column
              cell-class="is-vcentered"
              field="request"
              label="依頼名"
              width="20%"
              sortable
              v-slot="props"
            >
              <router-link
                :to="{
                  name: 'RequestPage',
                  params: { request_id: props.row.request_id }
                }"
              >
                {{ props.row.requestname }}
              </router-link>
            </b-table-column>
            <b-table-column
              cell-class="is-vcentered"
              field="detail"
              label="詳細"
              width="35%"
              v-slot="props"
            >
              {{ props.row.content }}
            </b-table-column>
            <b-table-column
              cell-class="is-vcentered"
              field="engineer"
              label="参加者"
              width="25%"
              v-slot="props"
            >
              <router-link
                v-for="engineer in props.row.engineers"
                :key="engineer.user_id"
                :to="{
                  name: 'MyPage',
                  params: { user_id: engineer.user_id }
                }"
              >
                <b-tooltip :label="engineer.username">
                  <div :style="iconStyle(48, engineer.icon)" />
                </b-tooltip>
              </router-link>
            </b-table-column>
          </b-table>
        </b-tab-item>
        <b-tab-item>
          <template #header>
            <b-icon icon="account-edit-outline"></b-icon>
            <span>
              過去に受けた依頼
              <b-tag rounded> {{ profile.submissions.length }} </b-tag>
            </span>
          </template>
          <b-table :data="profile.submissions" :default-sort="['date', 'desc']">
            <b-table-column
              cell-class="is-vcentered"
              field="state"
              label="状態"
              width="10%"
              centered
              v-slot="props"
            >
              <b-tag
                :type="props.row.finish === false ? 'is-success' : 'is-danger'"
              >
                {{ props.row.finish === false ? "受付中" : "終了" }}
              </b-tag>
            </b-table-column>
            <b-table-column
              cell-class="is-vcentered"
              field="date"
              label="依頼日時"
              width="10%"
              sortable
              centered
              v-slot="props"
            >
              {{ new Date(props.row.createdat).toLocaleDateString() }}
              <br />
              {{ new Date(props.row.createdat).toLocaleTimeString() }}
            </b-table-column>
            <b-table-column
              cell-class="is-vcentered"
              field="client"
              label="依頼者"
              width="10%"
              centered
              v-slot="props"
            >
              <router-link
                v-if="!!props.row.request"
                :to="{
                  name: 'MyPage',
                  params: { user_id: props.row.request.client.user_id }
                }"
              >
                <b-tooltip :label="props.row.request.client.username">
                  <div :style="iconStyle(64, props.row.request.client.icon)" />
                </b-tooltip>
              </router-link>
            </b-table-column>
            <b-table-column
              cell-class="is-vcentered"
              field="request"
              label="依頼名"
              width="20%"
              sortable
              v-slot="props"
            >
              <router-link
                v-if="!!props.row.request"
                :to="{
                  name: 'RequestPage',
                  params: { request_id: props.row.request_id }
                }"
              >
                {{ props.row.request.requestname }}
              </router-link>
            </b-table-column>
            <b-table-column
              cell-class="is-vcentered"
              field="detail"
              label="詳細"
              width="30%"
              v-slot="props"
            >
              {{ props.row.content }}
            </b-table-column>
            <b-table-column
              cell-class="is-vcentered"
              field="engineer"
              label="提出物"
              width="20%"
              v-slot="props"
            >
              <router-link
                :to="{
                  name: 'SubmissionPage',
                  query: { id: props.submission_id }
                }"
              >
                提出物
              </router-link>
            </b-table-column>
          </b-table>
        </b-tab-item>
      </b-tabs>
    </section>
  </div>
</template>

<script>
import ProfileEditor from "@/components/ProfileEditor";

export default {
  data() {
    return {
      profile: {
        user_id: null,
        username: "",
        email: "",
        icon: "",
        comment: "",
        SNS: {
          Github: "",
          Twitter: "",
          Facebook: ""
        },
        requests: [],
        submissions: []
      }
    };
  },
  watch: {
    $route(to, from) {
      this.getProfile(this.$route.params.user_id);
    }
  },
  methods: {
    // ユーザプロフィールの取得
    getProfile(user_id) {
      fetch(`${process.env.API}/user/${user_id}`)
        .then(data => data.json())
        .then(profile => {
          if (process.env.NODE_ENV === "development") {
            console.log(`Profile:`);
            console.log(profile);
          }
          this.profile = profile;
          this.getRequest();
        });
    },
    // リクエストの取得
    getRequest() {
      this.profile.submissions.forEach((submission, i) => {
        fetch(`${process.env.API}/request/${submission.request_id}`)
          .then(data => data.json())
          .then(request => {
            if (process.env.NODE_ENV === "development") {
              console.log(`Request of Submission #${i}:`);
              console.log(request);
            }
            submission.request = request;
            this.$forceUpdate();
          });
      });
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
  components: {
    "profile-editor": ProfileEditor
  },
  created() {
    this.getProfile(this.$route.params.user_id);
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
