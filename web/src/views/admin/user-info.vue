<script setup lang="ts">
import { userDetailApi, type userDetailResponse } from '@/api/user-api';
import LunarUserLoginEcharts from '@/components/admin/echarts/lunar-user-login-echarts.vue';
import LunarIcon from '@/components/base/lunar-icon.vue';
import LunarPointTitle from '@/components/base/lunar-point-title.vue';
import { useStore } from '@/stores';
import { dateTimeFormat } from '@/utils/date';
import { Message } from '@arco-design/web-vue';
import { reactive } from 'vue';

const store = useStore()

interface dateLineType {
  date: string
  title: string
  abstract: string
}

const dateLineList: dateLineType[] = [
  {date: "2025-06-17", title: "milestone V1.0.0", abstract: "milestone V1.0.0"},
]

const data = reactive<userDetailResponse>({
  userID: 0,
  nickname: '',
  username: '',
  email: '',
  avatar: '',
  createdAt: '',
  updatedAt: '',
  roleList: [],
  addr: '',
})

async function getData() {
  const res = await userDetailApi()
  if (res.code) {
    Message.error(res.message)
    return
  }
  Object.assign(data, res.data)
}
getData()
</script>

<template>
  <div class="user-info">
    <div class="head sector">
      <div class="banner"></div>
      <div class="info">
        <div class="nickname">
          <span>{{ data.nickname }}</span>
          <IconEdit />
        </div>
        <div class="other">
          <span>
            <LunarIcon icon="ic:sharp-location-on" />
            {{ data.addr }}
          </span>
          <span>
            <LunarIcon icon="mynaui:door-open-solid" />
            公司
          </span>
          <span>
            <LunarIcon icon="mynaui:package-solid" />
            部门
          </span>
        </div>
        <div class="avatar">
          <a-avatar :size="90" :image-url="data.avatar" />
        </div>
        <div class="actions">
          <a-button type="outline">关注用户</a-button>
          <a-button type="primary">发送消息</a-button>
        </div>
      </div>
    </div>
    <div class="body">
      <div>
        <div class="base-info sector">
          <LunarPointTitle>基本信息</LunarPointTitle>
          <div class="body">
            <a-form :model="{}" :wrapper-col-props="{span: 20}" :label-col-props="{span: 4}">
              <a-form-item label="用户名">{{ data.nickname }}</a-form-item>
              <a-form-item label="昵称">{{ data.nickname }}</a-form-item>
              <a-form-item label="注册时间">{{ dateTimeFormat(data.createdAt) }}</a-form-item>
              <a-form-item label="登陆时间">{{ dateTimeFormat(data.updatedAt) }}</a-form-item>
              <a-form-item label="手机号码">******</a-form-item>
              <a-form-item label="邮箱">{{ data.email }}</a-form-item>
              <a-form-item label="密码">
                已设置
                <a class="edit" href="javascript:void(0);">修改</a>
              </a-form-item>
            </a-form>
          </div>
        </div>
        <div class="my-role sector">
          <LunarPointTitle>我的角色</LunarPointTitle>
          <div class="body">
            <a-tag v-for="item in data.roleList">{{ item }}</a-tag>
          </div>
        </div>
      </div>
      <div class="sector">
        <div class="data-info">
          <LunarPointTitle>数据统计</LunarPointTitle>
          <div class="body">
            <div class="item">
              <div>114514</div>
              <div>登陆次数</div>
            </div>
            <div class="item">
              <div>1919</div>
              <div>代码提交</div>
            </div>
            <div class="item">
              <div>810</div>
              <div>任务完成</div>
            </div>
          </div>
        </div>
        <div class="login-info">
          <LunarPointTitle>登陆情况</LunarPointTitle>
          <div class="body">
            <LunarUserLoginEcharts />
          </div>
        </div>
      </div>
      <div class="sector">
        <div class="dynamic">
          <LunarPointTitle>我的动态</LunarPointTitle>
          <div class="body">
            <a-timeline>
              <a-timeline-item v-for="item in dateLineList" :label="item.date" line-type="dashed">
                {{ item.title }}
                <br/>
                <a-typography-text type="secondary">{{ item.abstract }}</a-typography-text>
              </a-timeline-item>
            </a-timeline>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="less">
.user-info {
  background-color: transparent!important;
  padding: 0!important;

  .head {
    .banner {
      height: 100px;
      width: 100%;
      background-image: url("@/assets/img/banner.png");
      background-repeat: no-repeat;
      background-size: cover;
    }

    .info {
      padding: 15px 200px 20px 180px;
      position: relative;

      .nickname {
        font-size: 18px;
        font-weight: 600;

        svg {
          margin-left: 10px;
          cursor: pointer;
        }
      }

      .other {
        display: flex;
        align-items: center;
        margin-top: 20px;
        color: var(--color-text-2);

        span {
          margin-right: 40px;
        }
      }

      .actions {
        position: absolute;
        right: 10px;
        top: 10px;

        >* {
          margin-left: 10px;
        }
      }

      .avatar {
        position: absolute;
        left: 80px;
        top: -40px;
      }
    }
  }

  >.body {
    margin-top: 20px;
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-column-gap: 20px;

    .base-info {
      padding: 20px;

      .body {
        margin-top: 10px;

        .arco-form-item {
          margin-bottom: 10px;
        }

        .edit {
          margin-left: 20px;
        }
      }
    }

    .my-role {
      margin-top: 10px;
      padding: 20px;

      .body {
        margin-top: 10px;

        >span {
          margin-right: 10px;
        }
      }
    }

    .data-info {
      padding: 20px;

      .body {
        margin-top: 10px;
        display: flex;
        grid-template-rows: repeat(3, 1fr);
        grid-column-gap: 50px;

        .item {
          width: 100%;
          height: 100px;
          border-radius: 5px;
          background-color: var(--color-fill-1);
          display: flex;
          justify-content: center;
          flex-direction: column;
          align-items: center;

          div:nth-child(1) {
            font-size: 30px;
            font-weight: 600;
            color: var(--color-text-1);
          }

          div:nth-child(2) {
            color: var(--color-text-2);
            margin-top: 5px;
          }

          &:nth-child(1) > div:nth-child(1) {
            color: rgba(20, 148, 240, 1);
          }

          &:nth-child(2) > div:nth-child(1) {
            color: rgba(242, 110, 110, 1);
          }

          &:nth-child(3) > div:nth-child(1) {
            color: rgba(33, 209, 108, 1);
          }
        }
      }
    }

    .login-info {
      padding: 20px;
    }

    .dynamic {
      padding: 20px;

      .body {
        margin-top: 10px;
      }
    }
  }

  .sector {
    background-color: var(--color-bg-1);
    border-radius: 5px;
  }
}
</style>