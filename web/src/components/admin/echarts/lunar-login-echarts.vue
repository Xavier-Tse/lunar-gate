<script setup lang="ts">
import { dataUserApi, type dataUserResponse } from '@/api/data-api'
import { theme } from '@/components/base/theme'
import { Message } from '@arco-design/web-vue'
import * as echarts from 'echarts'
import {onMounted, reactive, watch} from "vue"

type EChartsOption = echarts.EChartsOption
let myChart: echarts.ECharts | null = null

const data = reactive<dataUserResponse>({
  dateList: [],
  loginCountList: [],
  registerCountList: [],
})

watch(() => theme.value, () => {
  initOptions()
})

onMounted(async() => {
  const chartDom = document.querySelector('.lunar-login-echarts') as HTMLDivElement
  myChart = echarts.init(chartDom)
  const res = await dataUserApi()
  if (res.code) {
    Message.error(res.message)
    return
  }
  Object.assign(data, res.data)
  initOptions()
})

window.onresize = () => {
  myChart?.resize()
}

function initOptions() {
  let option: EChartsOption
  const textColor = getComputedStyle(document.body).getPropertyValue("--color-text-1")
  const lineColor = getComputedStyle(document.body).getPropertyValue("--color-neutral-2")
  let themeColor = [
    '#1c5ae0',
    '#66ccff',
  ]

   option = {
    color: themeColor,
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'cross',
        label: {
          backgroundColor: '#6a7985',
        },
      },
    },
    legend: {
      data: ['登录', '注册',],
      textStyle: {
        color: textColor,
      },
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true,
    },
    xAxis: [
      {
        type: 'category',
        boundaryGap: false,
        data: data.dateList,
      },
    ],
    yAxis: [
      {
        type: 'value',
        splitLine: {
          lineStyle: {
            color: lineColor,
          },
        },
      },
    ],
    series: [
      {
        name: '登录',
        type: 'line',
        emphasis: {
          focus: 'series'
        },
        data: data.loginCountList,
        smooth: true,
      },
      {
        name: '注册',
        type: 'line',
        emphasis: {
          focus: 'series'
        },
        smooth: true,
        data: data.registerCountList,
      },
    ],
  }
  option && myChart?.setOption(option)
}
</script>

<template>
  <div class="lunar-login-echarts"></div>
</template>

<style lang="less">
.lunar-login-echarts {
  width: 100%;
  height: 300px;
}
</style>