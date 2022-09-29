<template>
  <div class="content">
    <div>
      <v-btn-toggle mandatory tile group class="content__menu__options">
        <v-btn
          v-for="option in options"
          :key="option.label"
          v-on:click="changeComponent(option)"
          large
          plain
          :value="option"
        >
          <span>{{ option.label }}</span>
        </v-btn>
      </v-btn-toggle>
      <v-divider></v-divider>
    </div>
    <section>
      <div class="content__page-title">
        <h2>{{ componentSelected.label }}</h2>
      </div>
      <v-divider></v-divider>
      <div class="content__page-component">
        <component
          :is="componentSelected.component"
          :conns="connections"
          :stds="stadistics"
        />
      </div>
    </section>
  </div>
</template>

<script>
import ConnectionsPage from "./ConnectionsPage.vue";
import StadisticsPage from "./StadisticsPage.vue";

export default {
  name: "MenuOptions",
  props: ["connections", "stadistics"],

  data: () => ({
    tab: null,
    options: [
      {
        label: "Connections",
        component: ConnectionsPage,
      },
      {
        label: "Stadistics",
        component: StadisticsPage,
      },
    ],
    componentSelected: {
      label: "Connections",
      component: ConnectionsPage,
    },
  }),
  methods: {
    changeComponent: function (option) {
      this.componentSelected = option;
    },
  },
};
</script>

<style>
.content {
  width: 100%;
}

.content__menu__options {
  padding: 0px 8px;
}

.content__page-title {
  margin: 34px 0px;
  letter-spacing: 1px;
  padding: 0px 26px;
}

.content__menu__options__btn--active {
  border-bottom: 1px solid royalblue;
  background: red;
  color: white;
}

.content__page-component {
  padding: 0px 26px;
  margin-top: 22px;
}

@media screen and (min-width: 768px) {
  .content__menu__options {
    padding: 0px 88px;
  }

  .content__page-component,
  .content__page-title {
    padding: 0px 100px;
  }
}

@media screen and (min-width: 1024px) {
  .content__menu__options {
    padding: 0px 170px;
  }

  .content__page-title,
  .content__page-component {
    padding: 0px 192px;
  }
}

@media screen and (min-width: 1440px) {
  .content__menu__options {
    padding: 0px 270px;
  }

  .content__page-title,
  .content__page-component {
    padding: 0px 292px;
  }
}
</style>
