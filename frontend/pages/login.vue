<template>
  <v-app id="inspire">
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
        @click="moveToOauth"
        >Googleアカウントでログイン</v-btn
      >
    </v-app-bar>

    <v-main class="relative">
      <v-responsive class="movie">
        <video
          autoplay
          loop
          muted
          src="https://res.cloudinary.com/dmqxwobws/video/upload/v1596726217/study_zshdsf.mp4"
        ></video>
      </v-responsive>

      <v-container class="fill-height content" fluid>
        <v-row align="center" justify="center">
          <v-col class="text-center" cols="12">
            <h1 class="text-h1 py-4">学び続けよう。</h1>
            <h2 class="py-3">
              LearnMemoは学習を記録して成長を可視化するサービスです。
            </h2>
            <v-btn
              class="my-10"
              large
              depressed
              rounded
              theme="light"
              primary="text"
              color="primary"
              @click="moveToSimpleLogin"
              >かんたんログイン(機能テスト用)</v-btn
            >
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
import { Component, Vue } from 'nuxt-property-decorator'

@Component({
  components: {},
})
export default class SignUpForm extends Vue {
  moveToOauth() {
    window.location.href = '/api/user/auth/google'
  }
 moveToSimpleLogin () {
    window.location.href = '/api/user/simple-login'
  }
}
</script>

<style scoped lang="scss">
.relative {
  position: relative;
  overflow: hidden;
  z-index: 0;
}

.movie {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 1;
  contain: strict;
}

.content {
  color: #fff;
  height: 100%;
  z-index: 2;
  position: relative;
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding: 0 1rem;
}

$white: #fff;
$google-blue: #4285f4;
$button-active-blue: #1669f2;

.google-btn {
  width: 184px;
  height: 42px;
  background-color: $google-blue;
  border-radius: 2px;
  box-shadow: 0 3px 4px 0 rgba(0, 0, 0, 0.25);
  .google-icon-wrapper {
    position: absolute;
    margin-top: 1px;
    margin-left: 1px;
    width: 40px;
    height: 40px;
    border-radius: 2px;
    background-color: $white;
  }
  .google-icon {
    position: absolute;
    margin-top: 11px;
    margin-left: 11px;
    width: 18px;
    height: 18px;
  }
  .btn-text {
    float: right;
    margin: 11px 11px 0 0;
    color: $white;
    font-size: 14px;
    letter-spacing: 0.2px;
    font-family: 'Roboto';
  }
  &:hover {
    box-shadow: 0 0 6px $google-blue;
  }
  &:active {
    background: $button-active-blue;
  }
}

@import url(https://fonts.googleapis.com/css?family=Roboto:500);
</style>
