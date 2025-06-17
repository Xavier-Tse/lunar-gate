<script setup lang="ts">
import { dataComputerApi, type dataComputerResponse } from '@/api/data-api'
import { theme } from '@/components/base/theme'
import { Message } from '@arco-design/web-vue'
import * as echarts from 'echarts'
import {onMounted, reactive, watch} from "vue"

type EChartsOption = echarts.EChartsOption
let myChart: echarts.ECharts | null = null

const data = reactive<dataComputerResponse>({
  cpu: 0,
  mem: 0,
  disk: 0,
})

watch(() => theme.value, () => {
  initEcharts()
})

onMounted(async() => {
  const chartDom = document.querySelector('.lunar-system-echarts') as HTMLDivElement
  myChart = echarts.init(chartDom)
  const res = await dataComputerApi()
  if (res.code) {
    Message.error(res.message)
    return
  }
  Object.assign(data, res.data)
  initEcharts()
})

window.onresize = () => {
  myChart?.resize()
}

function initEcharts() {
  let option: EChartsOption;
  const lineColor = getComputedStyle(document.body).getPropertyValue("--color-neutral-2")

  let themeColor = [
    '#1c5ae0',
    '#66ccff',
  ]
  if (theme.value === "dark") {
    themeColor = [
      '#1c5ae0',
      '#66ccff',
    ]
  }

  option = {
    color: themeColor,
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'shadow'
      },
      formatter: (params: any) => {
        const data = params[0]
        return `<div>
<div>${data.name}</div>
<div>${data.seriesName} ${(data.data as number).toFixed(1)}%</div>
</div>`
      }
    },
    grid: {
      left: '3%',
      right: '10%',
      bottom: '3%',
      top: "5%",
      containLabel: true
    },
    xAxis: {
      type: 'value',
      min: 0,
      max: 100,
      axisLabel: {
        formatter: '{value}%'
      },
      splitLine: {
        lineStyle: {
          color: lineColor
        }
      }
    },
    yAxis: {
      type: 'category',
      data: ['CPU', '内存', '磁盘']
    },
    series: [
      {
        name: '使用率',
        type: 'bar',
        data: [data.cpu, data.mem, data.disk],
        label: {
          show: true,
          formatter: function (params: any) {
            return (params.data as number).toFixed(1) + "%"
          }
        }
      },
    ]
  };
  option && myChart?.setOption(option);
}
</script>

<template>
  <div class="lunar-system-echarts"></div>
</template>

<style lang="less">
.lunar-system-echarts {
  width: 100%;
  height: 300px;
}
</style>