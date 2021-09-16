<template>
  <div class="container">
    <section class="hero is-small is-primary mb-3">
      <div class="hero-body is-flex is-justify-content-space-between">
        <p class="title mb-0">
          リクエスト一覧
        </p>
        <b-button class="is-light" label="新規リクエスト" outlined />
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
          field="client"
          label="依頼者"
          width="10%"
          sortable
          centered
          v-slot="props"
        >
          <a href="">
            <b-tooltip :label="props.row.client.username">
              <b-image
                class="image is-64x64 is-inline-block is-centered"
                :src="props.row.client.icon"
                rounded
              />
            </b-tooltip>
          </a>
        </b-table-column>

        <b-table-column
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

        <b-table-column field="detail" label="詳細" width="30%" v-slot="props">
          {{ props.row.detail }}
        </b-table-column>

        <b-table-column
          field="engineer"
          label="参加者"
          width="20%"
          v-slot="props"
        >
          <a
            href=""
            v-for="engineer in props.row.engineers"
            :key="engineer.username"
          >
            <b-tooltip :label="engineer.username">
              <b-image
                class="image is-48x48 is-inline-block is-centered has-border"
                :src="engineer.icon"
                rounded
              />
            </b-tooltip>
          </a>
        </b-table-column>
      </b-table>
    </section>
  </div>
</template>

<script>
const data = require("../../src/assets/sample.json");

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
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
