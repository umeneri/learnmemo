<template>
  <v-app-bar app>
    <v-toolbar-title>LearnMemo</v-toolbar-title>
    <v-spacer></v-spacer>
    <v-btn
      class="my-10"
      outlined
      color="blue"
      depressed
      theme="light"
      primary="text"
      @click="logout"
      >ログアウト</v-btn
    >
  </v-app-bar>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'

@Component({})
export default class Header extends Vue {
  logout() {
    window.location.href = '/api/user/logout'
  }
}
</script>
