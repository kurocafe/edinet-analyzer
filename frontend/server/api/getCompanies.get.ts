export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig(event)

  const res = await $fetch(
    `${config.public.apiBase}/api/v1/companies` 
  )
  

})