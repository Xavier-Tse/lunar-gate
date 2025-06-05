import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime"
import "dayjs/locale/zh-cn"

dayjs.extend(relativeTime)
dayjs.locale('zh-cn')

export function dateTimeFormat(date: string): string {
  return dayjs(date).format("YYYY-MM-DD HH:mm:ss")
}

export function dateFormat(date: string): string {
  return dayjs(date).format("YYYY-MM-DD")
}

export function relativeTimeFormat(date: string): string {
  return dayjs(date).fromNow()
}

type dateType = "datetime" | "date" | "relative"
export function dateFuncFormat(date: string, type?: dateType): string {
  if (type === "datetime") {
    return dateTimeFormat(date)
  }
  if (type === "date") {
    return dateFormat(date)
  }
  if (type === "relative") {
    return relativeTimeFormat(date)
  }
  return dateTimeFormat(date)
}
