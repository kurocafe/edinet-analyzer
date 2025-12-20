<template>
  <Bar
    id="my-chart-id"
    :options="chartOptions"
    :data="chartData"
  />
</template>

<script setup lang="ts">
  import { Bar } from 'vue-chartjs';
  import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale } from 'chart.js';
  import { computed } from 'vue'

  const props = defineProps<{
    data: number[],
    name: string
  }>()
  ChartJS.register(
    Title,
    Tooltip,
    Legend,
    BarElement,
    CategoryScale,
    LinearScale
  )

  const chartData = computed(() => ({
    labels: ['2023', '2024'],
    datasets: [
      {
        label: props.name,
        data: props.data,
        backgroundColor: ['#3b82f6', '#1e40af']
      }
    ]
  }))

  const chartOptions = {
    responsive: true
  }

  
</script>