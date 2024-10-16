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
              class="w-full flex items-center space-x-2 bg-gray-200 active:bg-gray-300 py-2 px-2 rounded-lg text-gray-800"
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
              class="w-full flex items-center space-x-2 hover:bg-gray-200 active:bg-gray-300 py-2 px-2 rounded-lg text-gray-500"
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
          <router-link to="/login">
            <button
                style="display: flex; justify-content: center; align-items: center; width: 100%; background-color: #007bff; color: white; padding: 8px; border-radius: 8px; margin-top: 65vh;"
            >
              <span style="text-align: center; font-size: small;font-weight: bold">退出登录</span>
            </button>
          </router-link>
        </nav>
      </aside>
      <main class="flex-grow p-6">
        <div class="overflow-auto" style="height: 95vh">
          <table class="table table-xs table-pin-rows table-pin-cols">
            <!-- head -->
            <thead>
            <tr>
              <th>用户ID</th>
              <th>用户名</th>
              <th>用户真实姓名</th>
              <th>用户年龄</th>
              <th>用户状态</th>
              <th>操作</th>
            </tr>
            </thead>
            <tbody>
            <!-- row 1 -->
            <tr v-for="user in users" :key="user.id">
              <td>{{ user.ID }}</td>
              <td>
                <div class="flex items-center gap-3">
                  <div class="avatar">
                    <div class="mask mask-squircle w-12 h-12">
                      <!-- Assume avatar is a path to the user's avatar image -->
                      <img :src="user.avatar" alt="Avatar"/>
                    </div>
                  </div>
                  <div>
                    <div class="font-bold">{{ user.username }}</div>
                  </div>
                </div>
              </td>
              <td>
                {{ user.name }}
              </td>
              <td>{{ user.age }}</td>
              <td>{{ showStatus(user.status) }}</td>
              <th>
                <button class="btn btn-ghost btn-xs" v-show="user.status"
                        @click="updateUserStatus(false,user.username)">停用账号
                </button>
                <button class="btn btn-ghost btn-xs" v-show="!user.status"
                        @click="updateUserStatus(true,user.username)">启用账号
                </button>
              </th>
            </tr>
            </tbody>
            <tfoot>
            <tr>
              <th>用户ID</th>
              <th>用户名</th>
              <th>用户真实姓名</th>
              <th>用户年龄</th>
              <th>用户状态</th>
              <th>操作</th>
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
  name: 'home',
  created() {
    this.getAllUser();
  },
  methods: {
    navigateToPackage(router) {
      this.$router.push(router);
    },
    // 获取用户信息
    getAllUser() {
      this.$axios.get("/main/getAllUserInfo").then((res) => {
        if (res.data.code == 200) {
          this.users = res.data.data
        }
        console.log(res)
      }).catch((res) => {
        console.log(res)
      })
    },
    showStatus(flag) {
      if (flag) {
        return "正常"
      } else {
        return "停用"
      }
    },
    updateUserStatus(status, username) {
      this.$axios.post("/main/updateUserStatus", {
        username: username,
        status: status
      }).then((res) => {
        if (res.data.code === 200) {
          this.$message.success("更新用户状态成功")
        } else {
          this.$message.error("更新用户状态失败: " + res.data.message)
        }
        console.log(res)
      }).catch((res) => {
        console.log(res)
      })
    }
  },
  data() {
    return {
      users: []
    }
  }
}
</script>

<style scoped>

</style>