export default defineEventHandler(async () => {
  const config = useRuntimeConfig()
  const res = await $fetch(`${config.public.apiBase}/api/v1/rankings/profit-growth`)
  return res
})