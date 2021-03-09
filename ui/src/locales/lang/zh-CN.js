import antd from 'ant-design-vue/es/locale-provider/zh_CN'
import momentCN from 'moment/locale/zh-cn'

const components = {
  antLocale: antd,
  momentName: 'zh-cn',
  momentLocale: momentCN
}

const locale = {
  'site.title': 'ZenData',

  'menu.data.list': '我的数据',
  'menu.data.edit': '数据编辑',
  'menu.config.list': '字段列表',
  'menu.config.edit': '字段编辑',
  'menu.ranges.list': '序列列表',
  'menu.ranges.edit': '序列编辑',
  'menu.ranges.item.edit': '序列项编辑',
  'menu.instances.list': '实例列表',
  'menu.instances.edit': '实例编辑',
  'menu.instances.item.edit': '实例项编辑',
  'menu.excel.list': '表格列表',
  'menu.excel.edit': '表格编辑',
  'menu.text.list': '文本列表',
  'menu.text.edit': '文本编辑',

  'title.data.create': '数据创建',
  'title.config.create': '字段创建',
  'title.ranges.create': '序列创建',
  'title.instances.create': '实例创建',
  'title.excel.create': '表格创建',
  'title.text.create': '文本创建',

  'msg.mine': '我的数据',
  'msg.buildin': '內置数据',
  'msg.info': '信息',
  'msg.range': '区间',
  'msg.reference': '引用',
  'msg.preview': '预览',
  'msg.data.preview': '数据预览',

  'msg.workdir': '工作目录',
  'msg.help': '帮助',
  'msg.yes': '是',
  'msg.no': '否',

  'msg.data': '数据',
  'msg.config': '字段',
  'msg.ranges': '序列',
  'msg.instances': '实例',
  'msg.excel': 'Excel',
  'msg.excel.sheet': 'Excel表格',
  'msg.text': '文本',
  'msg.exec': '执行',
  'msg.file': '文件',

  'msg.design.title': '测试数据设计',
  'msg.design.create.brother': '创建同级',
  'msg.design.create.child': '创建子级',
  'msg.design.remove.node': '删除节点',

  'action.list': '列表',
  'action.create': '创建',
  'action.add': '添加',
  'action.edit': '编辑',
  'action.delete': '删除',
  'action.design': '设计',
  'action.back': '返回',
  'action.import.from.file': '刷新数据',

  'tips.refer': '引用',
  'tips.pls.select': '请选择',
  'tips.number': '数字',
  'tips.range.int': '整数或整数区间',
  'tips.number.or.letter': '整数或单个字母',

  'tips.value.each.line': '每行一个值',
  'tips.delete': '确认删除？',
  'tips.search': '输入关键字搜索',
  'tips.success.to.import': '成功导入数据。',
  'tips.zero': '0表示取所有记录',
  'tips.expr': '请输入数学运算表达式，由相同文件中的字段组成，如 "($field_step_negative * $field_nested_range) * -1 + 1000"。',

  'form.name': '名称',
  'form.file': '文件',
  'form.dir': '目录',
  'form.sub.dir': '子目录',
  'form.folder': '文件夹',
  'form.path': '路径',
  'form.file.name': '文件名',
  'form.desc': '描述',
  'form.content': '内容',
  'form.prefix': '前缀',
  'form.postfix': '后缀',
  'form.loop': '循环',
  'form.loopfix': '间隔符',
  'form.type': '类型',
  'form.type.list': '列表',
  'form.type.timestamp': '时间戳',
  'form.type.interval': '范围',
  'form.type.literal': '常量',
  'form.width': '宽度',
  'form.mode': '模式',
  'form.mode.parallel': '平行',
  'form.mode.recursive': '递归',
  'form.left.pad': '左填充',
  'form.right.pad': '右填充',
  'form.rand': '随机',
  'form.format': '格式',
  'form.function': '函数',
  'form.start': '开始',
  'form.end': '结束',
  'form.value': '取值',
  'form.repeat': '重复',
  'form.step': '步长',
  'form.col': '列名',
  'form.count': '数目',
  'form.expr': '表达式',
  'form.opt': '操作',
  'form.save': '保存',
  'form.reset': '重置',
  'form.cancel': '取消',

  'valid.required': '该字段不能为空。',
  'valid.loop.check': '需为整数或整数区间。',
  'valid.folder.users': '用户数据必须保存在users/目录下。',
  'valid.folder.yaml': '公用数据必须保存在/yaml目录下。',
  'valid.folder.data': 'Excel数据必须保存在data/目录下。',

}

export default {
  ...components,
  ...locale
}
