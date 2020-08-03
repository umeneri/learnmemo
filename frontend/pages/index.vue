<template>
  <v-app>
    <v-navigation-drawer app permanent>
      <v-toolbar flat>
        <v-list>
          <v-list-tile>
            <v-list-tile-title class="title">
              Sample App
            </v-list-tile-title>
          </v-list-tile>
        </v-list>
      </v-toolbar>

      <v-divider></v-divider>

      <v-list dense class="pt-0">
        <v-list-tile v-for="item in items" :key="item.title">
          <v-list-tile-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-tile-action>

          <v-list-tile-content>
            <v-list-tile-title>{{ item.title }}</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>

    <v-main class="bar-container">
      <p class="title">
        This year progress
      </p>
      <v-progress-linear v-model="yearProgress"></v-progress-linear>
      <p class="percentage">
        {{ `${this.yearProgress.toFixed(6)}%` }}
      </p>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'

@Component({
  components: {},
})
export default class extends Vue {
  items: Object[] = []
  yearProgress: number = 20

  mounted() {
    this.items = [
      {
        title: 'Progress',
        icon: 'mdi-timer',
      },
    ]
  }
}
</script>

<style scoped>
.title {
  font-family: 'Roboto', sans-serif;
  font-size: 20px;
  text-align: center;
}

.percentage {
  font-family: 'Roboto', sans-serif;
  font-size: 18px;
  text-align: center;
}

.bar-container {
  margin: 40px;
}
</style>
