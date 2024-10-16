<template>
  <canvas ref="pieChartRef"></canvas>
</template>

<script>
import {defineComponent, ref, onMounted, watch} from 'vue';
import {Chart, registerables} from 'chart.js';

Chart.register(...registerables);

export default defineComponent({
  name: 'PieChart',
  props: {
    chartData: {
      type: Object,
      required: true
    }
  },
  setup(props) {
    const pieChartRef = ref(null);
    let pieChartInstance = null;

    const createChart = () => {
      // 如果已存在实例，先销毁
      if (pieChartInstance) {
        pieChartInstance.destroy();
      }

      pieChartInstance = new Chart(pieChartRef.value, {
        type: 'pie',
        data: props.chartData,
        options: {
          responsive: true,
          maintainAspectRatio: false
        }
      });
    };

    onMounted(createChart);

    // 使用 watch 来监控 chartData 的变化
    watch(() => props.chartData, (newData) => {
      // 更新图表数据
      if (pieChartInstance) {
        pieChartInstance.data = newData;
        pieChartInstance.update();
      }
    }, {deep: true});

    return {
      pieChartRef
    };
  }
});
</script>

  