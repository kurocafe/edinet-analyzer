export default defineEventHandler(async () => {
  const config = useRuntimeConfig()
  const res = await $fetch(`${config.public.apiBase}/api/v1/rankings/income-growth`)
  return res
})