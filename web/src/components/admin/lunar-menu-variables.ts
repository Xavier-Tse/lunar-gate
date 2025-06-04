import { ref } from "vue";

export const collapsed = ref(false)

export interface MenuType {
  title: string
  name: string
  icon: string
  children?: MenuType[]
}
