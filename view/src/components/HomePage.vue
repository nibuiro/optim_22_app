<template>
  <div class="container">
    <section class="hero is-small is-primary mb-3">
      <div class="hero-body is-flex is-justify-content-space-between">
        <p class="title mb-0">
          リクエスト一覧
        </p>
        <request-form class="is-light" />
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
        :data="data"
        :accepting="onlyAccepting"
        :paginated="isPaginated"
        :per-page="perPage"
        :current-page.sync="currentPage"
        :sort-icon="sortIcon"
        :sort-icon-size="sortIconSize"
        :row-class="
          row => onlyAccepting && row.accepting === false && 'is-hidden'
        "
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
            :type="props.row.accepting === true ? 'is-success' : 'is-danger'"
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
          sortable
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
          <router-link
            :to="{
              name: 'RequestPage',
              params: { request_id: props.row.requestid }
            }"
          >
            {{ props.row.request }}
          </router-link>
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
          label="参加者"
          width="20%"
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
    </section>
  </div>
</template>

<script>
import RequestForm from "@/components/RequestForm";

const data = require("../../src/assets/sampleRequests.json");

export default {
  data() {
    return {
      data,
      onlyAccepting: false,
      isPaginated: true,
      defaultSort: ["date", "desc"],
      sortIcon: "chevron-up",
      sortIconSize: "",
      currentPage: 1,
      perPage: 10
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
    "request-form": RequestForm
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
