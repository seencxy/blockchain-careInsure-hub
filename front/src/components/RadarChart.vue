<template>
  <canvas ref="radarChartRef"></canvas>
</template>

<script>
import { defineComponent, ref, onMounted, watch } from 'vue';
import { Chart, registerables} from 'chart.js';

// Register chart.js components
Chart.register(...registerables);

export default defineComponent({
  name: 'RadarChart',
  props: {
    chartData: {
      type: Object,
      required: true
    },
  },
  setup(props) {
    const radarChartRef = ref(null);
    let radarChartInstance = null;

    // Function to create the chart
    const createChart = () => {
      // If an instance already exists, destroy it before creating a new one
      if (radarChartInstance) {
        radarChartInstance.destroy();
        radarChartInstance = null;
      }

      // Make sure the ref is set before attempting to get the context
      if (radarChartRef.value) {
        const ctx = radarChartRef.value.getContext('2d');
        radarChartInstance = new Chart(ctx, {
          type: 'radar',
          data: props.chartData,
          options: {
            responsive: true,
            maintainAspectRatio: false,
            elements: {
              line: {
                borderWidth: 3
              }
            }
          }
        });
      }
    };

    // Create the chart after the component mounts
    onMounted(createChart);

    // Watch for changes in the props.chartData and update the chart accordingly
    watch(() => props.chartData, (newData) => {
      // Make sure to only update if the chart instance is available
      if (radarChartInstance) {
        radarChartInstance.data = newData;
        radarChartInstance.update();
      }
    }, {
      deep: true // Watch for nested changes within the data object
    });

    // Return the chart reference to be used in the template
    return {
      radarChartRef
    };
  }
});
</script>

<style>
/* Set the height of the canvas container to control the chart size */
.radar-chart-container {
  position: relative;
  height: 400px; /* Adjust this value as needed */
}
</style>
