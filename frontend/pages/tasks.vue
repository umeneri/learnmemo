<template>
  <v-app>
    <v-data-table
      :headers="headers"
      :items="tasks"
      sort-by="elapsedTime"
      class="elevation-1"
    >
      <template v-slot:top>
        <v-toolbar flat color="white">
          <v-toolbar-title>学習記録</v-toolbar-title>
          <v-divider class="mx-4" inset vertical></v-divider>
          <v-spacer></v-spacer>
          <v-dialog v-model="dialog" max-width="500px">
            <template v-slot:activator="{ on, attrs }">
              <v-btn color="primary" dark class="mb-2" v-bind="attrs" v-on="on"
                >新しい記録</v-btn>
            </template>
            <v-card>
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
                      ></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field
                        v-model="editedTask.elapsedTime"
                        label="学習時間"
                      ></v-text-field>
                    </v-col>
                  </v-row>
                </v-container>
              </v-card-text>

              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="blue darken-1" text @click="close">キャンセル</v-btn>
                <v-btn color="blue darken-1" text @click="save">保存</v-btn>
              </v-card-actions>
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
      <template v-slot:no-data>
        <v-btn color="primary" @click="initialize">リセット</v-btn>
      </template>
    </v-data-table>
  </v-app>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'nuxt-property-decorator'

interface Task {
  id: number
  title: string
  elapsedTime: number
  status: number
  createdAt: string
  updatedAt: string
}

@Component({
  components: {},
})
export default class DataTable extends Vue {
  dialog = false
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
  tasks: Array<Task> = []
  editedIndex = -1
  editedTask: Task = {
    id: 0,
    title: '',
    elapsedTime: 0,
    status: 0,
    createdAt: '2020/08/01',
    updatedAt: '2020/08/01',
  }
  defaultTask: Task = {
    id: 0,
    title: '',
    elapsedTime: 0,
    status: 0,
    createdAt: '2020/08/01',
    updatedAt: '2020/08/01',
  }

  get formTitle() {
    return this.editedIndex === -1 ? '新しい記録' : '編集'
  }

  @Watch('dialog')
  OnDialogChange(val: boolean) {
    val || this.close()
  }

  created() {
    this.initialize()
  }

  async initialize() {
    try {
      const response = await this.$axios.$get('/task/v1/list')
      console.log(response.data)
      this.tasks = response.data
    } catch (error) {
      console.log(error)
    }
  }

  editTask(task: Task) {
    console.log(task);

    this.editedIndex = this.tasks.indexOf(task)
    this.editedTask = Object.assign({}, task)
    this.dialog = true
  }

  deleteTask(task: Task) {
    const index = this.tasks.indexOf(task)
    confirm('Are you sure you want to delete this task?') &&
      this.tasks.splice(index, 1)
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
      Object.assign(this.tasks[this.editedIndex], this.editedTask)
    } else {
      this.tasks.push(this.editedTask)
    }
    this.close()
  }
}
</script>
