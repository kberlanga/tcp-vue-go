<template>
  <div class="content-table">
    <v-data-table
      :headers="headers"
      :items="stadistics"
      :items-per-page="10"
      :height="630"
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
        RECEIVING: "#64b2d9",
        SUCCESS: "#9cc295",
      },
    };
  },
  computed: {
    stadistics: {
      get: function () {
        return this.stds.map((stadistic) => {
          return {
            ...stadistic,
            shippingTime: this.formatDate(stadistic.shippingTime),
            receiptTime: this.formatDate(stadistic.receiptTime),
          };
        });
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
  },
};
</script>
