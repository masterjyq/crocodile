<template lang="html">
  <div :val="value_">
    <div>
      <el-radio v-model:value="type" label="1" size="mini" border
        >每月</el-radio
      >
    </div>
    <div>
      <el-radio v-model:value="type" label="5" size="mini" border
        >不指定</el-radio
      >
    </div>
    <div>
      <el-radio v-model:value="type" label="2" size="mini" border
        >周期</el-radio
      >
      <span style="margin-left: 0px; margin-right: 0px">从</span>
      <el-input-number
        @change="type = '2'"
        v-model:value="cycle.start"
        :min="1"
        :max="12"
        size="mini"
        style="width: 100px"
      ></el-input-number>
      <span style="margin-left: 0px; margin-right: 0px">至</span>
      <el-input-number
        @change="type = '2'"
        v-model:value="cycle.end"
        :min="2"
        :max="12"
        size="mini"
        style="width: 100px"
      ></el-input-number>
      月
    </div>
    <div>
      <el-radio v-model:value="type" label="3" size="mini" border
        >循环</el-radio
      >
      <span style="margin-left: 0px; margin-right: 0px">从</span>
      <el-input-number
        @change="type = '3'"
        v-model:value="loop.start"
        :min="1"
        :max="12"
        size="mini"
        style="width: 100px"
      ></el-input-number>
      <span style="margin-left: 0px; margin-right: 0px">月开始，每</span>
      <el-input-number
        @change="type = '3'"
        v-model:value="loop.end"
        :min="1"
        :max="12"
        size="mini"
        style="width: 100px"
      ></el-input-number>
      月执行一次
    </div>
    <div>
      <el-radio v-model:value="type" label="4" size="mini" border
        >指定</el-radio
      >
      <el-checkbox-group v-model:value="appoint" :min="1">
        <div v-for="i in 2" :key="i">
          <el-checkbox
            @change="type = '4'"
            v-for="j in 6"
            :key="j"
            :label="check((i - 1) * 6 + j, 'getVal')"
            >{{ check((i - 1) * 6 + j, 'getText') }}
          </el-checkbox>
        </div>
      </el-checkbox-group>
    </div>
  </div>
</template>

<script>
import { $on, $off, $once, $emit } from '../../utils/gogocodeTransfer'
export default {
  props: {
    value: {
      type: String,
      default: '*',
    },
  },
  data() {
    return {
      type: '1', // 类型
      cycle: {
        // 周期
        start: 0,
        end: 0,
      },
      loop: {
        // 循环
        start: 0,
        end: 0,
      },
      week: {
        // 指定周
        start: 0,
        end: 0,
      },
      work: 0,
      last: 0,
      appoint: [], // 指定
    }
  },
  computed: {
    value_() {
      let result = []
      switch (this.type) {
        case '1': // 每秒
          result.push('*')
          break
        case '2': // 年期
          result.push(`${this.cycle.start}-${this.cycle.end}`)
          break
        case '3': // 循环
          result.push(`${this.loop.start}/${this.loop.end}`)
          break
        case '4': // 指定
          result.push(this.appoint.join(','))
          break
        case '6': // 最后
          result.push(`${this.last === 0 ? '' : this.last}L`)
          break
        default:
          // 不指定
          result.push('?')
          break
      }
      $emit(this, 'update:value', result.join(''))
      return result.join('')
    },
  },
  watch: {
    value(a, b) {
      this.updateVal()
    },
  },
  methods: {
    updateVal() {
      if (!this.value) {
        return
      }
      if (this.value === '?') {
        this.type = '5'
      } else if (this.value.indexOf('-') !== -1) {
        // 2周期
        if (this.value.split('-').length === 2) {
          this.type = '2'
          this.cycle.start = this.value.split('-')[0]
          this.cycle.end = this.value.split('-')[1]
        }
      } else if (this.value.indexOf('/') !== -1) {
        // 3循环
        if (this.value.split('/').length === 2) {
          this.type = '3'
          this.loop.start = this.value.split('/')[0]
          this.loop.end = this.value.split('/')[1]
        }
      } else if (this.value.indexOf('*') !== -1) {
        // 1每
        this.type = '1'
      } else if (this.value.indexOf('L') !== -1) {
        // 6最后
        this.type = '6'
        this.last = this.value.replace('L', '')
      } else if (this.value.indexOf('#') !== -1) {
        // 7指定周
        if (this.value.split('#').length === 2) {
          this.type = '7'
          this.week.start = this.value.split('#')[0]
          this.week.end = this.value.split('#')[1]
        }
      } else if (this.value.indexOf('W') !== -1) {
        // 8工作日
        this.type = '8'
        this.work = this.value.replace('W', '')
      } else {
        // *
        this.type = '4'
        this.appoint = this.value.split(',')
      }
    },
    check(val, act) {
      if (act === 'getVal') {
        return val.toString()
      }
      if (val < 10) {
        return '0' + val
      }
      return val.toString()
    },
  },
  created() {
    this.updateVal()
  },
  emits: ['update:value'],
}
</script>

<style lang="css">
.el-checkbox + .el-checkbox {
  margin-left: 10px;
}
</style>
