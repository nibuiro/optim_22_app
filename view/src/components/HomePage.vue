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
          <b-switch v-model="isAccepting" :disabled="!isAccepting"
            >受付中のみ</b-switch
          >
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
        :accepting="isAccepting"
        :paginated="isPaginated"
        :per-page="perPage"
        :current-page.sync="currentPage"
        :sort-icon="sortIcon"
        :sort-icon-size="sortIconSize"
        default-sort="user.first_name"
        aria-next-label="Next page"
        aria-previous-label="Previous page"
        aria-page-label="Page"
        aria-current-label="Current page"
      >
        <b-table-column
          field="id"
          label="ID"
          width="40"
          sortable
          numeric
          v-slot="props"
        >
          {{ props.row.id }}
        </b-table-column>

        <b-table-column
          field="user.first_name"
          label="First Name"
          sortable
          v-slot="props"
        >
          {{ props.row.user.first_name }}
        </b-table-column>

        <b-table-column
          field="user.last_name"
          label="Last Name"
          sortable
          v-slot="props"
        >
          {{ props.row.user.last_name }}
        </b-table-column>

        <b-table-column
          field="date"
          label="Date"
          sortable
          centered
          v-slot="props"
        >
          <span class="tag is-success">
            {{ new Date(props.row.date).toLocaleDateString() }}
          </span>
        </b-table-column>

        <b-table-column label="Gender" v-slot="props">
          <span>
            <b-icon
              pack="fas"
              :icon="props.row.gender === 'Male' ? 'mars' : 'venus'"
            >
            </b-icon>
            {{ props.row.gender }}
          </span>
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
      isAccepting: false,
      isPaginated: true,
      sortIcon: "chevron-up",
      sortIconSize: "",
      currentPage: 1,
      perPage: 5
    };
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
