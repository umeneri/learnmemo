<template>
  <v-app id="inspire">
    <v-app-bar app>
      <v-toolbar-title>LearnMemo</v-toolbar-title>
    </v-app-bar>

    <v-main>
      <v-container class="fill-height content" fluid>
        <v-row align="center" justify="center">
          <v-col class="text-center" cols="12" sm="8" md="4">
            <h4 class="text-h4 my-10">登録まであと一歩。</h4>

            <v-form ref="form" v-model="valid">
              <v-text-field
                v-model="user.name"
                label="名前"
                placeholder="nickname"
                :rules="nameRules"
                required
                maxlength="100"
              ></v-text-field>
            </v-form>
            <v-btn
              large
              depressed
              rounded
              theme="light"
              primary="text"
              color="primary"
              :disabled="!valid"
              @click="save"
              >登録する
            </v-btn>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
    <v-footer app>
      <v-col class="text-center">
        <span>&copy; {{ new Date().getFullYear() }} Masato Koishi</span>
      </v-col>
    </v-footer>
  </v-app>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'nuxt-property-decorator'

interface User {
  name: string
}

@Component({
  components: {},
})
export default class UserForm extends Vue {
  nameRules: Array<any> = [
    (v: string) => !!v || '名前が未入力です',
    (v: string) => v.length <= 100 || '名前は100文字以内です',
  ]
  valid = true
  user: User = {
    name: '',
  }

  async created() {
    // this.user = await this.fetchUser()
    this.user = { name: '' }
  }

  async save() {
    const user = await this.updateUser(this.user)
    if (user !== null) {
      this.$router.push('/')
    }
  }

  // async fetchUser(): Promise<User> {
  //   try {
  //     const response = await this.$axios.$get(`${window.location.origin}/api/user/v1/me`)
  //     console.log(response.data)
  //     return response.data
  //   } catch (error) {
  //     console.log(error)
  //     return { name: '' }
  //   }
  // }

  async updateUser(user: User): Promise<User> {
    try {
      const response = await this.$axios.$put(`${window.location.origin}/api/user/v1/update`, user)
      console.log(response.data)
      return response.data
    } catch (error) {
      console.log(error)
      return { name: '' }
    }
  }
}
</script>
<style scoped lang="scss"></style>
