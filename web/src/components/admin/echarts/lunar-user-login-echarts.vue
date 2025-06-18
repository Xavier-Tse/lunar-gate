<script setup lang="ts">
import { theme } from '@/components/base/theme'
import * as echarts from 'echarts'
import {onMounted, watch} from "vue"

type EChartsOption = echarts.EChartsOption
let myChart: echarts.ECharts | null = null

watch(() => theme.value, () => {
  initEcharts()
})

onMounted(async() => {
  const chartDom = document.querySelector('.lunar-user-login-echarts') as HTMLDivElement
  myChart = echarts.init(chartDom)
  initEcharts()
})

window.onresize = () => {
  myChart?.resize()
}

function initEcharts() {
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
      data: ['登录',],
      textStyle: {
        color: textColor,
      },
    },
    grid: {
      left: '8%',
      right: '8%',
      bottom: '5%',
      containLabel: true,
    },
    xAxis: [
      {
        type: 'category',
        boundaryGap: false,
        data: [
          "2025-06-18",
          "2025-06-17",
          "2025-06-16",
          "2025-06-15",
          "2025-06-14",
          "2025-06-13",
          "2025-06-12",
        ],
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
        data: [11, 45, 14, 19, 19, 8, 10],
        smooth: true,
      },
    ],
  }
  option && myChart?.setOption(option)
}
</script>

<template>
  <div class="lunar-user-login-echarts"></div>
</template>

<style lang="less">
.lunar-user-login-echarts {
  width: 100%;
  height: 300px;
}
</style>