<template>
  <h1>{{ title }}</h1>
</template>

<script lang="ts">
import {
  Component,
  Prop,
  Vue
} from "nuxt-property-decorator"

@Component({})
export default class Header extends Vue {
  @Prop() title!: string
}
</script>