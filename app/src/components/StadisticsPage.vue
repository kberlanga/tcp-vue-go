<template>
  <div class="content-table">
    <v-data-table
      :headers="headers"
      :items="stadistics"
      :items-per-page="10"
      :height="500"
      :fixed-header="true"
      :loading="false"
      group-by="toChannel"
      class="elevation-1"
      show-group-by
    >
      <template v-slot:[`item.status`]="{ item }">
        <v-chip :color="getColor(item.status)">
          {{ item.status }}
        </v-chip>
      </template>
    </v-data-table>
    <div class="content-table__filters">
      <h3>Filters</h3>
      <section class="content-table__filters__selects">
        <v-select
          v-model="filters.status"
          :items="status"
          item-text="label"
          label="Status"
          @change="filter()"
          return-object
          outlined
        ></v-select>
        <v-select
          v-model="filters.client"
          :items="clients"
          item-text="label"
          label="Client"
          @change="filter()"
          return-object
          outlined
        ></v-select>
        <v-select
          v-model="filters.channel"
          :items="channels"
          item-text="label"
          label="Channel"
          @change="filter()"
          return-object
          outlined
        ></v-select>
        <section class="content-table__filters__buttons">
          <v-btn v-on:click="cleanFilters">
            <v-icon left> mdi-eraser </v-icon>
            Clear</v-btn
          >
        </section>
      </section>
    </div>
  </div>
</template>

<script>
export default {
  name: "ConnectionsPage",
  props: ["stds"],

  data: function () {
    return {
      headers: [
        { text: "Status", value: "status", groupable: false },
        { text: "Client", value: "toClient" },
        { text: "Channel", value: "toChannel" },
        { text: "Shipping Time", value: "shippingTime", groupable: false },
        { text: "Receipt Time", value: "receiptTime", groupable: false },
      ],
      statusColors: {
        PENDING: "#a1a1a1",
        SUCCESS: "#9cc295",
      },
      stadistics: [],
      filters: {
        client: { value: "", label: "None" },
        channel: { value: "", label: "None" },
        status: { value: "", label: "None" },
      },
    };
  },
  created() {
    this.stadistics = this.stds.map((stadistic) => {
      return {
        ...stadistic,
        shippingTime: this.formatDate(stadistic.shippingTime),
        receiptTime: this.formatDate(stadistic.receiptTime),
      };
    });
  },
  computed: {
    data() {
      return this.stds.map((stadistic) => {
        return {
          ...stadistic,
          shippingTime: this.formatDate(stadistic.shippingTime),
          receiptTime: this.formatDate(stadistic.receiptTime),
        };
      });
    },
    clients: {
      get: function () {
        return [
          { value: "", label: "None" },
          ...this.stds.map((stadistic) => {
            return { value: stadistic.toClient, label: stadistic.toClient };
          }),
        ];
      },
    },
    channels: {
      get: function () {
        return [
          { value: "", label: "None" },
          ...this.stds.map((stadistic) => {
            return { value: stadistic.toChannel, label: stadistic.toChannel };
          }),
        ];
      },
    },
    status() {
      return [
        { value: "", label: "None" },
        ...Object.keys(this.statusColors).map((status) => {
          return { value: status, label: status };
        }),
      ];
    },
  },
  watch: {
    data: {
      deep: true,
      handler: function (newVal) {
        this.stadistics = newVal;
      },
    },
  },
  methods: {
    formatDate: function (date) {
      const newDate = new Date(date).toLocaleString();
      if (newDate === "Invalid Date") {
        return "---";
      }

      return newDate;
    },
    getColor: function (status) {
      return this.statusColors[status] ? this.statusColors[status] : "white";
    },
    filterByStatus: function (stadistics) {
      return stadistics.filter((stadistic) => {
        return stadistic.status === this.filters.status.value;
      });
    },
    filterByClient: function (stadistics) {
      return stadistics.filter((stadistic) => {
        return stadistic.toClient === this.filters.client.value;
      });
    },
    filterByChannel: function (stadistics) {
      return stadistics.filter((stadistic) => {
        return stadistic.toChannel === this.filters.channel.value;
      });
    },
    filter: function () {
      let filtered = this.data;
      if (this.filters.status.value !== "") {
        filtered = this.filterByStatus(filtered);
      }

      if (this.filters.channel.value !== "") {
        filtered = this.filterByChannel(filtered);
      }

      if (this.filters.client.value !== "") {
        filtered = this.filterByClient(filtered);
      }

      this.stadistics = filtered;
    },
    cleanFilters: function () {
      this.stadistics = this.data;
      this.filters = {
        client: { value: "", label: "None" },
        channel: { value: "", label: "None" },
        status: { value: "", label: "None" },
      };
    },
  },
};
</script>
