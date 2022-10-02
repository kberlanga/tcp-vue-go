<template>
  <div class="content-table">
    <v-data-table
      :headers="headers"
      :items="connections"
      :items-per-page="10"
      :height="500"
      :fixed-header="true"
      :loading="false"
      class="elevation-1"
    ></v-data-table>
    <div class="content-table__filters">
      <h3>Filters</h3>
      <section class="content-table__filters__selects">
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
          outlined
          return-object
          @change="filter()"
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
  props: ["conns"],

  data: function () {
    return {
      headers: [
        { text: "Client", value: "client" },
        { text: "Channel", value: "channel" },
        { text: "Subscription Time", value: "subscriptionTime" },
      ],
      connections: [],
      filters: {
        client: { value: "", label: "None" },
        channel: { value: "", label: "None" },
      },
    };
  },
  created() {
    this.connections = this.conns.map((connection) => {
      return {
        ...connection,
        subscriptionTime: new Date(
          connection.subscriptionTime
        ).toLocaleString(),
      };
    });
  },
  computed: {
    data() {
      return this.conns.map((connection) => {
        return {
          ...connection,
          subscriptionTime: new Date(
            connection.subscriptionTime
          ).toLocaleString(),
        };
      });
    },
    clients: {
      get: function () {
        return [
          { value: "", label: "None" },
          ...this.conns.map((connection) => {
            return { value: connection.client, label: connection.client };
          }),
        ];
      },
    },
    channels: {
      get: function () {
        return [
          { value: "", label: "None" },
          ...this.conns.map((connection) => {
            return { value: connection.channel, label: connection.channel };
          }),
        ];
      },
    },
  },
  watch: {
    data: {
      deep: true,
      handler: function (newVal) {
        this.connections = newVal;
      },
    },
  },
  methods: {
    filterByClient: function (connections) {
      return connections.filter((connection) => {
        return connection.client === this.filters.client.value;
      });
    },
    filterByChannel: function (connections) {
      return connections.filter((connection) => {
        return connection.channel === this.filters.channel.value;
      });
    },
    filter: function () {
      let filtered = this.data;

      if (this.filters.channel.value !== "") {
        filtered = this.filterByChannel(filtered);
      }

      if (this.filters.client.value !== "") {
        filtered = this.filterByClient(filtered);
      }

      this.connections = filtered;
    },
    cleanFilters: function () {
      this.connections = this.data;
      this.filters = {
        client: { value: "", text: "None" },
        channel: { value: "", text: "None" },
      };
    },
  },
};
</script>

<style>
@import "../assets/styles/table-filters.css";
</style>
