<template>
  <div class="app">
    <div class="flex">
      <aside class="h-screen w-56 bg-gray-100 text-gray-800 p-4">
        <div class="flex items-center mb-4 space-x-1">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
               fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
               stroke-linejoin="round" class="h-6 w-6">
            <path d="M3 9h18v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9Z"></path>
            <path d="m3 9 2.45-4.9A2 2 0 0 1 7.24 3h9.52a2 2 0 0 1 1.8 1.1L21 9"></path>
            <path d="M12 3v6"></path>
          </svg>
          <h1 class="text-lg font-medium">管理后台</h1>
        </div>
        <nav class="space-y-2">
          <button
              class="w-full flex items-center space-x-2 hover:bg-gray-200 active:bg-gray-300 py-2 px-2 rounded-lg text-gray-500"
              @click="navigateToPackage('/admin/home')">
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
                class="w-4 h-4"
            >
              <path d="m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
              <polyline points="9 22 9 12 15 12 15 22"></polyline>
            </svg>
            <span class="text-sm font-medium">用户管理</span>
          </button>
          <button
              class="w-full flex items-center space-x-2 bg-gray-200 active:bg-gray-300 py-2 px-2 rounded-lg text-gray-800"
              @click="navigateToPackage('/admin/pay')">
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
                class="w-4 h-4"
            >
              <path d="M21 12V7H5a2 2 0 0 1 0-4h14v4"></path>
              <path d="M3 5v14a2 2 0 0 0 2 2h16v-5"></path>
              <path d="M18 12a2 2 0 0 0 0 4h4v-4Z"></path>
            </svg>
            <span class="text-sm font-medium">支付管理</span>
          </button>
          <button
              class="w-full flex items-center space-x-2 hover:bg-gray-200 active:bg-gray-300 py-2 px-2 rounded-lg text-gray-500"
              @click="navigateToPackage('/admin/package')">
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
                class="w-4 h-4"
            >
              <path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"></path>
              <circle cx="9" cy="7" r="4"></circle>
              <path d="M22 21v-2a4 4 0 0 0-3-3.87"></path>
              <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
            </svg>
            <span class="text-sm font-medium">套餐管理</span>
          </button>
          <div class="pt-4"> <!-- pt-4 是为了添加一些顶部间距，可以根据需要调整 -->
            <router-link to="/login">
              <button
                  style="display: flex; justify-content: center; align-items: center; width: 100%; background-color: #007bff; color: white; padding: 8px; border-radius: 8px; margin-top: 65vh;"
              >
                <span style="text-align: center; font-size: small;font-weight: bold">退出登录</span>
              </button>
            </router-link>
          </div>
        </nav>
      </aside>
      <main class="flex-grow p-6">
        <div class="overflow-auto" style="height: 95vh">
          <table class="table table-xs table-pin-rows table-pin-cols">
            <!-- head -->
            <thead>
            <tr>
              <th>支付ID</th>
              <th>支付方式</th>
              <th>订单创建时间</th>
              <th>支付账号</th>
              <th>套餐</th>
              <th>支付状态</th>
            </tr>
            </thead>
            <tbody>
            <!-- row 1 -->
            <tr v-for="pay in pays" :key="pay.id">
              <td>{{ pay.ID }}</td>
              <td>支付宝</td>
              <td>
                {{ pay.CreatedAt }}
              </td>
              <td>{{ pay.Account }}</td>
              <td>{{ pay.Choose }}</td>
              <td>{{ printStatus(pay.Status) }}</td>
            </tr>
            </tbody>
            <tfoot>
            <tr>
              <th>支付ID</th>
              <th>支付方式</th>
              <th>订单创建时间</th>
              <th>支付账号</th>
              <th>套餐</th>
              <th>支付状态</th>
            </tr>
            </tfoot>

          </table>
        </div>
      </main>
    </div>
  </div>
</template>

<script>
export default {
  name: 'pay',
  created() {
    this.getPayRecord();
  },
  methods: {
    navigateToPackage(router) {
      this.$router.push(router);
    },
    getPayRecord() {
      this.$axios.get("/main/getPayRecord").then((res) => {
        if (res.data.code === 200) {
          this.pays = res.data.data
        } else {
          this.$message.error("获取订单记录失败: " + res.data.message)
        }
        console.log(res)
      }).catch((res => {
        console.log(res)
      }))
    },
    printStatus(flag) {
      if (flag === 0) {
        return "等待支付"
      } else if (flag === 2) {
        return "交易成功"
      } else {
      }
      return "交易失败"
    },
    openCreate() {
      location.href('my_modal_8')
    }
  },
  data() {
    return {
      pays: []
    }
  }
}
</script>

<style scoped>
tr {
  height: 50px; /* 或者您想要的任何高度 */
}
</style>