<template>
  <div class="app h-screen w-full overflow-auto">
    <div class="grid h-screen min-h-screen w-full overflow-hidden lg:grid-cols-[280px_1fr]">
      <div class="hidden border-r bg-gray-100/40 lg:block dark:bg-gray-800/40">
        <div class="flex flex-col gap-2">
          <div class="flex h-[60px] items-center px-6"><a class="flex items-center gap-2 font-semibold"
                                                          href="#">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                 fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                 stroke-linejoin="round" class="h-6 w-6">
              <path d="M3 9h18v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9Z"></path>
              <path d="m3 9 2.45-4.9A2 2 0 0 1 7.24 3h9.52a2 2 0 0 1 1.8 1.1L21 9"></path>
              <path d="M12 3v6"></path>
            </svg>
            <span class="">区块链养老保险平台</span></a></div>
          <div class="flex-1">
            <nav class="grid items-start px-4 text-sm font-medium">
              <router-link to="/"><a
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-50"
                  href="#">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                     viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                     stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4">
                  <path d="m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
                  <polyline points="9 22 9 12 15 12 15 22"></polyline>
                </svg>
                主页
              </a></router-link>
              <router-link to="order"><a
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-50"
                  href="#">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                     viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                     stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4">
                  <circle cx="8" cy="21" r="1"></circle>
                  <circle cx="19" cy="21" r="1"></circle>
                  <path
                      d="M2.05 2.05h2l2.66 12.42a2 2 0 0 0 2 1.58h9.78a2 2 0 0 0 1.95-1.57l1.65-7.43H5.12">
                  </path>
                </svg>
                交易状况
              </a></router-link>
              <router-link to="product"><a
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-50"
                  href="#">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                     viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                     stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4">
                  <path d="m7.5 4.27 9 5.15"></path>
                  <path
                      d="M21 8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16Z">
                  </path>
                  <path d="m3.3 7 8.7 5 8.7-5"></path>
                  <path d="M12 22V12"></path>
                </svg>
                产品
              </a></router-link>
              <router-link to="customer"><a
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-50"
                  href="#">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                     viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                     stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4">
                  <path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"></path>
                  <circle cx="9" cy="7" r="4"></circle>
                  <path d="M22 21v-2a4 4 0 0 0-3-3.87"></path>
                  <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
                </svg>
                顾客
              </a></router-link>
              <router-link to="analytic"><a
                  class="flex items-center gap-3 rounded-lg bg-gray-100 px-3 py-2 text-gray-900  transition-all hover:text-gray-900 dark:bg-gray-800 dark:text-gray-50 dark:hover:text-gray-50"
                  href="#">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                     viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                     stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4">
                  <path d="M3 3v18h18"></path>
                  <path d="m19 9-5 5-4-4-3 3"></path>
                </svg>
                分析
              </a></router-link>
              <router-link to="nodeInfo"><a
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-50"
                  href="#">
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    class="h-4 w-4"
                >
                  <path d="m7.5 4.27 9 5.15"></path>
                  <path
                      d="M21 8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16Z"></path>
                  <path d="m3.3 7 8.7 5 8.7-5"></path>
                  <path d="M12 22V12"></path>
                </svg>
                节点信息
              </a></router-link>
            </nav>
          </div>
        </div>
      </div>
      <div class="flex flex-col">
        <header class="flex h-14 lg:h-[60px] items-center gap-4 border-b bg-gray-100/40 px-6 dark:bg-gray-800/40"><a
            class="lg:hidden" href="#">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
               viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
               stroke-linejoin="round" class="h-6 w-6">
            <path d="M3 9h18v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9Z"></path>
            <path d="m3 9 2.45-4.9A2 2 0 0 1 7.24 3h9.52a2 2 0 0 1 1.8 1.1L21 9"></path>
            <path d="M12 3v6"></path>
          </svg>
          <span class="sr-only">Analytic</span></a>
          <div class="flex-1">
            <h1 class="font-semibold text-lg">分析</h1>
          </div>
          <div class="flex flex-1 items-center gap-4 md:ml-auto md:gap-2 lg:gap-4">
            <form class="ml-auto flex-1 sm:flex-initial">
            </form>
            <div class="dropdown dropdown-end">
              <div tabindex="0" role="button" class="btn btn-ghost btn-circle avatar">
                <div class="w-12 rounded-full">
                  <img alt="Tailwind CSS Navbar component" :src="avatar" style="width: 100%"/>
                </div>
              </div>
              <ul tabindex="0"
                  class="mt-3 z-[1] p-2 shadow menu menu-sm dropdown-content bg-base-100 rounded-box w-52">
                <li @click="openVerified">
                  <a class="justify-between">
                    实名认证
                    <span class="badge" v-if="isUserAuthenticated">已认证</span>
                    <span class="badge" v-else>未认证</span>
                  </a>
                </li>
                <li @click="openUpdate"><a>修改用户信息</a></li>
                <li @click="logout"><a>登出</a></li>
                <li @click="userLogout"><a>账号注销</a></li>
              </ul>
            </div>
          </div>
        </header>
        <main class="flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-6">
          <div class="box">
            <Line :chart-data="lineChartData"></Line>
          </div>
          <div class="box">
            <RadarChart :chart-data="radarChartData"/>
          </div>
        </main>
      </div>
    </div>
  </div>
</template>

<script>
import BarChart from '../../components/chart.vue';
import PieChart from '../../components/PieChart.vue';
import Line from '../../components/line.vue'
import RadarChart from '../../components/RadarChart.vue';

export default {
  name: "analytic",
  data() {
    return {
      avatar: '',
      isUserAuthenticated: false,
      lineChartData: {
        labels: ['0-20',
          '20-30',
          '30-40',
          '40-50',
          '50-60',
          '60-70',
          '70-80',
          '80+'],
        datasets: [
          {
            label: '系列一',
            backgroundColor: 'rgba(255, 99, 132, 0.2)',
            borderColor: 'rgba(255, 99, 132, 1)',
            data: [0, 10, 5, 2, 20],
          },
          {
            label: '系列二',
            backgroundColor: 'rgba(54, 162, 235, 0.2)',
            borderColor: 'rgba(54, 162, 235, 1)',
            data: [20, 30, 15, 10, 5],
          }
        ],
      },
      radarChartData: {
        labels: ['0-20',
          '20-30',
          '30-40',
          '40-50',
          '50-60',
          '60-70',
          '70-80',
          '80+'],
        datasets: [{
          label: 'My First Dataset',
          data: [65, 59, 90, 81, 56, 55, 40],
          fill: true,
          backgroundColor: 'rgba(255, 99, 132, 0.2)',
          borderColor: 'rgb(255, 99, 132)',
          pointBackgroundColor: 'rgb(255, 99, 132)',
          pointBorderColor: '#fff',
          pointHoverBackgroundColor: '#fff',
          pointHoverBorderColor: 'rgb(255, 99, 132)'
        }, {
          label: 'My Second Dataset',
          data: [28, 48, 40, 19, 96, 27, 100],
          fill: true,
          backgroundColor: 'rgba(54, 162, 235, 0.2)',
          borderColor: 'rgb(54, 162, 235)',
          pointBackgroundColor: 'rgb(54, 162, 235)',
          pointBorderColor: '#fff',
          pointHoverBackgroundColor: '#fff',
          pointHoverBorderColor: 'rgb(54, 162, 235)'
        }]
      }
    }
  },
  components: {
    BarChart,
    PieChart,
    Line,
    RadarChart
  },
  mounted() {
    this.$axios.get("/main/getStatistics").then((rep) => {
      if (rep.data.code === 200) {
        const responseData = rep.data.data;

        // 初始化datasets数组
        const newLineChartDatasets = [];
        const newRadarChartDatasets = [];

        // 动态生成datasets，这次使用套餐名作为标签
        Object.keys(responseData).forEach((packageName) => {
          const ageGroups = responseData[packageName];
          const ageGroupData = [
            ageGroups.age0To20,
            ageGroups.age20To30,
            ageGroups.age30To40,
            ageGroups.age40To50,
            ageGroups.age50To60,
            ageGroups.age60To70,
            ageGroups.age70To80,
            ageGroups.age80Plus
          ];

          // 使用套餐名作为每个数据集的label
          newLineChartDatasets.push({
            label: packageName,
            // 设定一种颜色方案，确保每个数据集颜色的唯一性
            backgroundColor: 'rgba(255, 99, 132, 0.2)',
            borderColor: 'rgba(255, 99, 132, 1)',
            data: ageGroupData,
          });

          newRadarChartDatasets.push({
            label: packageName,
            data: ageGroupData,
            fill: true,
            // 设定一种颜色方案，确保每个数据集颜色的唯一性
            backgroundColor: 'rgba(54, 162, 235, 0.2)',
            borderColor: 'rgb(54, 162, 235)',
            pointBackgroundColor: 'rgb(54, 162, 235)',
            pointBorderColor: '#fff',
            pointHoverBackgroundColor: '#fff',
            pointHoverBorderColor: 'rgb(54, 162, 235)'
          });
        });

        // 更新图表数据
        this.lineChartData.datasets = newLineChartDatasets;
        this.radarChartData.datasets = newRadarChartDatasets;

      } else {
        console.error("Error fetching data: ", rep.data.message);
      }
    }).catch((error) => {
      console.error("Error fetching data: ", error);
    });
  },


  created() {
    const authValue = localStorage.getItem('user_auth');
    this.isUserAuthenticated = authValue === 'true'; // 确保将字符串 'true' 转换为布尔值 true
    this.avatar = localStorage.getItem("avatar");
  },
  methods: {
    change() {
      this.advanced = !this.advanced
      this.base = !this.base
    },
    // 账号注销接口
    userLogout() {
      this.$axios.post('/user/logout', {}, {
        headers: {
          'Authorization': `Bearer ` + localStorage.getItem('auth_token')
        }
      }).then((res) => {
        console.log(res)
      }).catch((error) => {
        console.error(error)
      })

      // 删除所有缓存
      localStorage.removeItem("auth_token")
      localStorage.removeItem("user_auth")
      // 将路由切换到 /login
      this.$router.push('/login')
    },
    // 账号登出
    logout() {
      // 删除所有缓存
      localStorage.removeItem("auth_token")
      localStorage.removeItem("user_auth")
      // 将路由切换到 /login
      this.$router.push('/login')
    },
    openVerified() {
      if (this.isUserAuthenticated == false) {
        window.location.href = '#my_modal_8';
      }
    },
    openUpdate() {
      if (this.isUserAuthenticated == true) {
        window.location.href = '#updateUserInfo';
      }
    },
    getRandomColor() {
      const letters = '0123456789ABCDEF';
      let color = '#';
      for (let i = 0; i < 6; i++) {
        color += letters[Math.floor(Math.random() * 16)];
      }
      return color;
    },
    getAvatarStyle(user) {
      // 尝试从 localStorage 中获取颜色
      let colorStart = localStorage.getItem("colorStart");
      let colorEnd = localStorage.getItem("colorEnd");

      // 如果没有找到缓存的颜色，生成新的颜色
      if (!colorStart || !colorEnd) {
        colorStart = this.getRandomColor();
        colorEnd = this.getRandomColor();

        // 将生成的颜色存入 localStorage
        localStorage.setItem("colorStart", colorStart);
        localStorage.setItem("colorEnd", colorEnd);
      }

      // 通过指定色标的位置来增加渐变尺度
      // 比如从 10% 的位置开始到 90% 的位置结束
      return {
        background: `linear-gradient(to right, ${colorStart}, ${colorEnd})`,
        width: '32px',
        height: '32px',
        borderRadius: '50%',
        aspectRatio: '1 / 1',
        objectFit: 'cover'
      };
    },
  }

}
</script>

<style scoped>
.avatar {
  display: inline-block;
  overflow: hidden;
  /* 如果你要在里面放图标或者文字，建议设置 overflow */
  /* 其他样式 */
}

html, body {
  height: 100%;
  margin: 0;
}

.box {
  width: 100%;
  height: 40vh; /* 占视口高度的80% */
}

/* 如果父容器不是使用flex布局，可能需要添加额外的样式来确保它们填满整个屏幕 */

</style>