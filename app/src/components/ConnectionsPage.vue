<template>
  <div class="content-table">
    <v-data-table
      :headers="headers"
      :items="connections"
      :items-per-page="10"
      :height="550"
      :fixed-header="true"
      :loading="false"
      class="elevation-1"
    ></v-data-table>
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
    };
  },
  computed: {
    connections: {
      get: function () {
        return this.conns.map((connection) => {
          return {
            ...connection,
            subscriptionTime: new Date(
              connection.subscriptionTime
            ).toLocaleString(),
          };
        });
      },
    },
  },
};
</script>
