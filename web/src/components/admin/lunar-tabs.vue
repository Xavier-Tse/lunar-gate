<script setup lang="ts">
import router from '@/router';
import { onMounted, ref, watch } from 'vue';
import { useRoute } from 'vue-router';
import {Swiper, SwiperSlide} from 'swiper/vue';

interface TabType {
  title: string
  name: string
}

const tabList = ref<TabType[]>([
  {title: "数据统计", name: "data"},
])

const route = useRoute()

watch(() => route.path, () => {
  const obj = tabList.value.find((value) => value.name === route.name)
  if (obj) {
    return
  }
  tabList.value.push({
    title: route.meta.title as string,
    name: route.name as string,
  })
}, {immediate: true})

function save() {
  localStorage.setItem('lunar-tabs', JSON.stringify(tabList.value))
}

function load() {
  const data = localStorage.getItem('lunar-tabs')
  if (data) {
    tabList.value = JSON.parse(data)
  }
}
load()

watch(() => tabList.value.length, () => {
  save()
})

function select(tab: TabType) {
  router.push({name: tab.name})
}

function close(tab: TabType) {
  if (tab.name === 'data') {
    return
  }
  const index = tabList.value.findIndex((value) => value.name === route.name)
  if (index !== -1) {
    tabList.value.splice(index, 1)
    if (route.name === tab.name) {
      select(tabList.value[index - 1])
    }
  }
}

function closeAll() {
  tabList.value = [
    {title: "数据统计", name: "data"},
  ]
  select(tabList.value[0])
}

const slideCount = ref(12)
onMounted(() => {
  const swiperWidth = document.querySelector('.my-swiper')!.clientWidth
  const swiperScrollWidth = document.querySelector('.my-swiper .swiper-wrapper')!.scrollWidth
  if (swiperScrollWidth <= swiperWidth) {
    return
  }
  const slideList = document.querySelectorAll('.swiper-wrapper .swiper-slide')

  let sum = 0
  let count = 0
  for (const element of slideList) {
    sum += element.clientWidth
    if (sum > swiperWidth) {
      break
    }
    count++
  }
  slideCount.value = count
})
</script>

<template>
  <div class="lunar-tabs">
    <Swiper
      :slides-per-view="slideCount"
      class="my-swiper"
    >
      <swiper-slide v-for="item in tabList">
        <span @click.middle="close(item)" @click="select(item)" class="item"
          :class="{active: item.name === route.name}"
        >
          {{ item.title }}
          <IconClose @click.stop="close(item)" v-if="item.name != 'data'"></IconClose>
        </span>
      </swiper-slide>
    </Swiper>
    <span class="item close-all" @click="closeAll">
      全部关闭
    </span>
  </div>
</template>

<style lang="less">
.lunar-tabs {
  display: flex;
  justify-content: space-between;
  align-items: center;

  .my-swiper {
    width: calc(100% - 94px);
    overflow: hidden;
    white-space: nowrap;
    height: 100%;
    display: flex;
    align-items: center;

    .swiper-wrapper {
      display: flex;
      justify-content: start;
      width: 100%;

      .swiper-slide {
        width: auto !important;
      }
    }
  }

  .item {
    border: 1px solid var(--color-neutral-3);
    border-radius: 5px;
    padding: 3px 8px;
    cursor: pointer;
    font-size: 12px;
    margin-right: 10px;
    transition: all .3s;

    svg {
      font-size: 12px;
    }

    &.active {
      background-color: rgb(var(--arcoblue-6));
      color: white;
      border: 1px solid rgb(var(--arcoblue-6));
    }

    &:hover {
      background-color: rgb(var(--arcoblue-4));
      color: white;
      border: 1px solid rgb(var(--arcoblue-4));
    }
  }

  .close-all {
    margin-right: 0;
  }
}
</style>