<template>
  <v-app>
    <Header />
    <v-main>
      <v-data-table
        :headers="headers"
        :items="tasks"
        sort-by="createdAt"
        sort-desc
        class="elevation-1 pt-6"
      >
        <template v-slot:top>
          <v-toolbar flat color="white">
            <v-toolbar-title>学習記録</v-toolbar-title>
            <v-divider class="mx-4" inset vertical></v-divider>
            <v-spacer></v-spacer>
            <v-dialog v-model="dialog" max-width="500px">
              <template v-slot:activator="{ on, attrs }">
                <v-btn
                  color="primary"
                  dark
                  class="mb-2"
                  v-bind="attrs"
                  v-on="on"
                  >新しい記録</v-btn
                >
              </template>
              <v-card>
                <v-form ref="form" v-model="valid">
                  <v-card-title>
                    <span class="headline">{{ formTitle }}</span>
                  </v-card-title>

                  <v-card-text>
                    <v-container>
                      <v-row>
                        <v-col cols="12" sm="6" md="4">
                          <v-text-field
                            v-model="editedTask.title"
                            label="タイトル"
                            placeholder="英単語"
                            :rules="titleRules"
                            required
                            maxlength="100"
                          ></v-text-field>
                        </v-col>
                        <v-col cols="12" sm="6" md="4">
                          <v-text-field
                            v-model.number="editedTask.elapsedTimeHours"
                            label="学習時間(時)"
                            placeholder="1"
                            type="number"
                            required
                          ></v-text-field>
                        </v-col>
                        <v-col cols="12" sm="6" md="4">
                          <v-text-field
                            v-model.number="editedTask.elapsedTimeMinutes"
                            label="学習時間(分)"
                            placeholder="30"
                            type="number"
                            required
                          ></v-text-field>
                        </v-col>
                      </v-row>
                    </v-container>
                  </v-card-text>

                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" text @click="close"
                      >キャンセル</v-btn
                    >
                    <v-btn
                      color="blue darken-1"
                      text
                      @click="save"
                      :disabled="!valid"
                      >保存</v-btn
                    >
                  </v-card-actions>
                </v-form>
              </v-card>
            </v-dialog>
          </v-toolbar>
        </template>
        <template v-slot:item.actions="{ item }">
          <v-icon small class="mr-2" @click="editTask(item)">
            mdi-pencil
          </v-icon>
          <v-icon small @click="deleteTask(item)">
            mdi-delete
          </v-icon>
        </template>
        <template v-slot:item.elapsedTime="{ item }">
          <span>{{ getFormatedElapsedTime(item.elapsedTime) }}</span>
        </template>
        <template v-slot:item.createdAt="{ item }">
          <span>{{ getFormatedDate(item.createdAt) }}</span>
        </template>
        <template v-slot:no-data>
          <v-btn color="primary" @click="fetchTaskList">リセット</v-btn>
        </template>
      </v-data-table>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'nuxt-property-decorator'
import Header from '~/components/Header.vue'

interface Task {
  id: number
  title: string
  elapsedTime: number
  status: number
  createdAt: string
  updatedAt: string
}

interface EditedTask {
  id: number
  title: string
  elapsedTimeHours: number
  elapsedTimeMinutes: number
  status: number
  createdAt: string
  updatedAt: string
}

@Component({
  components: {
    Header
  },
})
export default class DataTable extends Vue {
  headers = [
    {
      text: 'タイトル',
      align: 'start',
      sortable: false,
      value: 'title',
    },
    { text: '学習時間', value: 'elapsedTime' },
    { text: '作成日', value: 'createdAt' },
    { text: 'アクション', value: 'actions', sortable: false },
  ]
  titleRules: Array<any> = [
    (v: string) => !!v || 'タイトルが未入力です',
    (v: string) => v.length <= 100 || 'タイトルは100文字以内です',
  ]
  defaultTask: EditedTask = {
    id: 0,
    title: '',
    elapsedTimeHours: 0,
    elapsedTimeMinutes: 0,
    status: 0,
    createdAt: new Date().toLocaleString(),
    updatedAt: new Date().toLocaleString(),
  }

  dialog = false
  valid = true
  tasks: Array<Task> = []
  editedIndex = -1
  editedTask: EditedTask = {
    id: 0,
    title: '',
    elapsedTimeHours: 0,
    elapsedTimeMinutes: 0,
    status: 0,
    createdAt: new Date().toLocaleString(),
    updatedAt: new Date().toLocaleString(),
  }

  get formTitle() {
    return this.editedIndex === -1 ? '新しい記録' : '編集'
  }

  getFormatedDate(d: string) {
    const date = new Date(d)
    return `${date.getFullYear()}/${date.getMonth() + 1}/${date.getDate()}`
  }

  getFormatedElapsedTime(elapsedTime: number) {
    const elapsedTimeHours = Math.floor(elapsedTime / 60)
    const elapsedTimeMinutes = Math.floor(elapsedTime % 60)
    if (elapsedTimeHours === 0) {
      return `${elapsedTimeMinutes}分`
    }
    return `${elapsedTimeHours}時間${elapsedTimeMinutes}分`
  }

  async created() {
    // this.tasks = await this.fetchTaskList()
    this.fetchTaskList()
  }

  taskToEditedTask(task: Task): EditedTask {
    const task2 = Object.assign({}, task)
    return {
      ...task2,
      elapsedTimeHours: Math.floor(task2.elapsedTime / 60),
      elapsedTimeMinutes: Math.floor(task2.elapsedTime % 60),
    }
  }

  editedTaskToTask(editedTask: EditedTask): Task {
    const task = Object.assign({}, editedTask)
    return {
      ...task,
      elapsedTime: task.elapsedTimeHours * 60 + task.elapsedTimeMinutes,
    }
  }

  editTask(task: Task) {
    this.editedIndex = this.tasks.indexOf(task)
    this.editedTask = this.taskToEditedTask(task)
    this.dialog = true
  }

  deleteTask(task: Task) {
    const index = this.tasks.indexOf(task)
    const b = confirm('Are you sure you want to delete this task?')
    if (b) {
      this.tasks.splice(index, 1)
      this.sendDeleteTask(task)
    }
  }

  @Watch('dialog')
  OnDialogChange(val: boolean) {
    val || this.close()
  }

  close() {
    this.dialog = false
    this.$nextTick(() => {
      this.editedTask = Object.assign({}, this.defaultTask)
      this.editedIndex = -1
    })
  }

  save() {
    if (this.editedIndex > -1) {
      const t = this.editedTaskToTask(this.editedTask)
      Object.assign(this.tasks[this.editedIndex], t)
      console.log(t)
      this.sendEditedTask(t)
    } else {
      const t = this.editedTaskToTask(this.editedTask)
      this.tasks.push(t)
      this.sendInitialTask(t)
    }
    this.close()
  }

  async fetchTaskList(): Promise<Array<Task>> {
    try {
      const response = await this.$axios.$get(
        `${window.location.origin}/api/task/v1/list`
      )
      console.log(response.data)
      this.tasks = response.data
      return response.data
    } catch (error) {
      console.log(error)
      return []
    }
  }

  async sendInitialTask(task: Task) {
    try {
      const response = await this.$axios.$post(
        `${window.location.origin}/api/task/v1/add`,
        task
      )
      console.log(response)
    } catch (error) {
      console.log(error)
    }
  }

  async sendEditedTask(task: Task) {
    try {
      const response = await this.$axios.$put(
        `${window.location.origin}/api/task/v1/update${task.id}`,
        task
      )
      console.log(response)
    } catch (error) {
      console.log(error)
    }
  }

  async sendDeleteTask(task: Task) {
    try {
      const response = await this.$axios.$delete(
        `${window.location.origin}/api/task/v1/delete/${task.id}`
      )
      console.log(response)
    } catch (error) {
      console.log(error)
    }
  }
}
</script>
