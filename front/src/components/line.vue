<template>
  <canvas ref="lineChartRef"></canvas>
</template>

<script>
import {defineComponent, ref, onMounted, watch} from 'vue';
import {Chart, registerables} from 'chart.js';

// 注册chart.js所有组件
Chart.register(...registerables);

export default defineComponent({
  name: 'LineChart',
  props: {
    chartData: {
      type: Object,
      required: true
    },
  },
  setup(props) {
    const lineChartRef = ref(null);
    let lineChartInstance = null;

    const createChart = () => {
      // 如果图表实例已存在，则销毁重建
      if (lineChartInstance) {
        lineChartInstance.destroy();
      }

      lineChartInstance = new Chart(lineChartRef.value, {
        type: 'line',
        data: props.chartData,
        options: {
          responsive: true,
          maintainAspectRatio: false,
          scales: {
            y: {
              beginAtZero: true
            }
          }
        }
      });
    };

    onMounted(createChart);

    // 监控chartData的变化，如果变化则更新图表
    watch(() => props.chartData, (newData) => {
      if (lineChartInstance) {
        lineChartInstance.data = newData;
        lineChartInstance.update();
      }
    }, {deep: true});

    return {
      lineChartRef
    };
  }
});
</script>
