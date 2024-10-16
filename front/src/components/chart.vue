<template>
    <canvas ref="chartRef" style="width: 100%;margin-top:15%"></canvas>
  </template>
  
  <script>
  import { defineComponent, ref, onMounted, watch } from 'vue';
  import { Chart, registerables } from 'chart.js';
  Chart.register(...registerables);
  
  export default defineComponent({
    name: 'BarChart',
    props: {
      chartData: {
        type: Object,
        required: true
      },
      chartOptions: {
        type: Object,
        default: () => {}
      }
    },
    setup(props) {
      const chartRef = ref(null);
      let chartInstance = null;
  
      const createChart = () => {
        if (chartInstance) {
          chartInstance.destroy(); // 销毁旧的图表实例
        }
        chartInstance = new Chart(chartRef.value, {
          type: 'bar',
          data: props.chartData,
          options: props.chartOptions
        });
      };
  
      onMounted(createChart);
  
      // 使用watch来监控chartData和chartOptions的变化
      watch(() => props.chartData, createChart, { deep: true });
      watch(() => props.chartOptions, createChart, { deep: true });
  
      return {
        chartRef
      };
    },
  });
  </script>
  
  
  <style scoped>
  /* 您的样式可以放在这里 */
  </style>
  