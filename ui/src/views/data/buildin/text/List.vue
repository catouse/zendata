<template>
  <div>
    <div class="head">
      <div class="title">文本列表</div>
      <div class="buttons">
        <a-button type="primary" @click="create()">新建</a-button>
      </div>
    </div>

    <a-table :columns="columns" :data-source="models" rowKey="id">

      <span slot="action" slot-scope="record">
        <a @click="edit(record)">编辑</a> |

        <a-popconfirm
            title="确认删除？"
            ok-text="是"
            cancel-text="否"
            @confirm="remove(record)"
        >
          <a href="#">删除</a>
        </a-popconfirm>
      </span>
    </a-table>

  </div>
</template>

<script>

import {listText, removeText} from "../../../../api/manage";

const columns = [
  {
    title: '名称',
    dataIndex: 'title',
  },
  {
    title: '引用',
    dataIndex: 'name',
  },
  {
    title: '路径',
    dataIndex: 'path',
  },
  {
    title: '操作',
    key: 'action',
    scopedSlots: { customRender: 'action' },
  },
];

export default {
  name: 'TextList',
  components: {
  },
  data() {
    return {
      models: [],
      columns,

      designVisible: false,
      designModel: {},
      time: 0,
    };
  },
  computed: {

  },
  created () {
    this.loadData()
  },
  mounted () {
  },
  methods: {
    create() {
      this.$router.push({path: '/data/buildin/text/edit/0'});
    },
    loadData() {
      listText().then(json => {
        console.log('listText', json)
        this.models = json.data
      })
    },
    edit(record) {
      console.log(record)
      this.$router.push({path: `/data/buildin/text/edit/${record.id}`});
    },
    remove(record) {
      console.log(record)
      removeText(record.id).then(json => {
        console.log('removeText', json)
        this.loadData()
      })
    },

    handleDesignOk() {
      console.log('handleDesignOk')
      this.designVisible = false
    },
    handleDesignCancel() {
      console.log('handleDesignCancel')
      this.designVisible = false
    },
  }
}
</script>

<style scoped>

</style>