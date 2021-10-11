<!-- リクエスト一覧ページ -->

<template>
  <div class="container">
    <section class="hero is-small is-primary mb-3">
      <div class="hero-body is-flex is-justify-content-space-between">
        <p class="title mb-0">リクエスト一覧</p>
        <request-form v-if="loggedin" class="is-light" />
      </div>
    </section>
    <section class="mb-3">
      <div class="is-flex is-justify-content-space-between">
        <div class="control is-flex">
          <b-switch v-model="onlyAccepting">受付中のみ</b-switch>
        </div>
        <b-field grouped group-multiline>
          <b-select v-model="perPage" :disabled="!isPaginated">
            <option value="5">5 per page</option>
            <option value="10">10 per page</option>
            <option value="15">15 per page</option>
            <option value="20">20 per page</option>
          </b-select>
        </b-field>
      </div>

      <b-table
        :loading="loading"
        :data="requests"
        :accepting="onlyAccepting"
        :paginated="isPaginated"
        :per-page="perPage"
        :current-page.sync="currentPage"
        :sort-icon="sortIcon"
        :sort-icon-size="sortIconSize"
        :row-class="row => onlyAccepting && row.finish === true && 'is-hidden'"
        :default-sort="defaultSort"
      >
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
          sortable
          centered
          v-slot="props"
        >
          <router-link
            :to="{
              name: 'MyPage',
              params: { user_id: props.row.client.user_id }
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
          width="30%"
          v-slot="props"
        >
          {{ props.row.content }}
        </b-table-column>
        <b-table-column
          cell-class="is-vcentered"
          field="engineer"
          label="参加者"
          width="20%"
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
    </section>
  </div>
</template>

<script>
import RequestForm from "@/components/RequestForm";
import * as api from "API";
import { iconStyle } from "iconStyle";

export default {
  data() {
    return {
      loggedin: this.$cookies.get("refresh_token") !== null,
      requests: [],
      onlyAccepting: false,
      isPaginated: true,
      defaultSort: ["date", "desc"],
      sortIcon: "chevron-up",
      sortIconSize: "",
      currentPage: 1,
      perPage: 10,
      loading: false,
      iconStyle
    };
  },
  components: {
    "request-form": RequestForm
  },
  async created() {
    this.loading = true;
    // リクエスト一覧の取得
    this.requests = await api.getRequests();
    this.loading = false;
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
