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
        <profile-form class="is-light ml-auto mt-5" :profile="profile" />
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
                :type="
                  props.row.accepting === true ? 'is-success' : 'is-danger'
                "
              >
                {{ props.row.accepting === true ? "受付中" : "終了" }}
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
              {{ new Date(props.row.date).toLocaleDateString() }}
              <br />
              {{ new Date(props.row.date).toLocaleTimeString() }}
            </b-table-column>
            <b-table-column
              cell-class="is-vcentered"
              field="request"
              label="依頼名"
              width="20%"
              sortable
              v-slot="props"
            >
              <a href="">
                {{ props.row.request }}
              </a>
            </b-table-column>
            <b-table-column
              cell-class="is-vcentered"
              field="detail"
              label="詳細"
              width="35%"
              v-slot="props"
            >
              {{ props.row.detail }}
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
                :key="engineer.userid"
                :to="{
                  name: 'MyPage',
                  params: { user_id: engineer.userid }
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
                :type="
                  props.row.accepting === true ? 'is-success' : 'is-danger'
                "
              >
                {{ props.row.accepting === true ? "受付中" : "終了" }}
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
              {{ new Date(props.row.date).toLocaleDateString() }}
              <br />
              {{ new Date(props.row.date).toLocaleTimeString() }}
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
                :to="{
                  name: 'MyPage',
                  params: { user_id: props.row.client.userid }
                }"
              >
                <b-tooltip :label="props.row.client.username">
                  <div :style="iconStyle(64, props.row.client.icon)" />
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
              <a href="">
                {{ props.row.request }}
              </a>
            </b-table-column>
            <b-table-column
              cell-class="is-vcentered"
              field="detail"
              label="詳細"
              width="30%"
              v-slot="props"
            >
              {{ props.row.detail }}
            </b-table-column>
            <b-table-column
              cell-class="is-vcentered"
              field="engineer"
              label="提出物"
              width="20%"
              v-slot="props"
            >
              <a href="">提出物</a>
            </b-table-column>
          </b-table>
        </b-tab-item>
      </b-tabs>
    </section>
  </div>
</template>

<script>
import ProfileForm from "@/components/ProfileForm";

const profile = require("../../src/assets/sampleProfile.json");

export default {
  data() {
    return {
      profile
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
  components: {
    "profile-form": ProfileForm
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
